using System;
using System.Collections;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Runtime.Remoting.Messaging;
using System.Threading;
using UnityEngine;

public class network : MonoBehaviour
{
    public Socket client;
    public string server_ip;
    public int port;
    private const int RECV_LEN = 8192;
    private Thread readThread;
    private Thread sendThread;
    private int receved = 0;
    private bool is_Connect = false;
    private byte[] recv_buffer = new byte[RECV_LEN];

    private byte[] long_pkg = null;
    private Action<SMessage> on_recv_tcp_data;
    private int long_pkg_size = 0;

    private Queue<byte[]> sendQueue = new Queue<byte[]>();


    public static network instance;
    
    private void Awake()
    {
        instance = this;
        DontDestroyOnLoad(this);
    }

    // Start is called before the first frame update
    void Start()
    {
        EventManager.Instance.AddEventAction("OnLogin", onLogin);
        this.on_recv_tcp_data = message => { LogTool.Instance.ToStringAll(message.Header); };
        connect2Server();
    }


    void onLogin(object data)
    {
        Debug.Log("data,login------------>>>");

        EventManager.Instance.RemoveEventAction("OnLogin", onLogin);
        print("----------->>>>");
    }

    // Update is called once per frame
    void Update()
    {
    }

    private void OnDestroy()
    {
        this.close();
        on_log_msg("释放完成...");
    }

    void on_connect_timeout()
    {
        this.on_log_msg("timeout");
    }

    void connect2Server()
    {
        try
        {
            this.client = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            IPAddress ipAddress = IPAddress.Parse(this.server_ip);
            IPEndPoint ipEndPoint = new IPEndPoint(ipAddress, this.port);

            IAsyncResult result =
                this.client.BeginConnect(ipEndPoint, new AsyncCallback(this.on_connected), this.client);
            bool success = result.AsyncWaitHandle.WaitOne(1000, true);
            if (!success)
            {
                this.on_connect_timeout();
            }
        }
        catch (System.Exception e)
        {
            this.on_log_msg(e.ToString());
        }
    }

    private void on_log_msg(string msg)
    {
        Debug.Log(msg);
    }

    public void start_send_data(ClientHeader header, byte[] body)
    {
        if (this.is_Connect == false)
        {
            return;
        }

        if (!this.client.Connected)
        {
            return;
        }

        var data = PackTools.PackObject(header);
        List<byte> buffer = new List<byte>(data);
        
        buffer.AddRange(body);
        this.sendQueue.Enqueue(buffer.ToArray());
    }

    void start_recv_data()
    {
        if (this.is_Connect == false)
        {
            return;
        }

        while (true)
        {
            if (!this.client.Connected)
            {
                break;
            }

            try
            {
                Debug.Log("读取数据中...");
                int recv_len = 0;
                // small
                if (this.receved < RECV_LEN)
                {
                    recv_len = this.client.Receive(this.recv_buffer, this.receved, RECV_LEN - this.receved,
                        SocketFlags.None);
                }
                else
                {
                    if (this.long_pkg == null)
                    {
                        var out_header = new byte[PackTools.PACK_LENGTH];
                        Array.Copy(this.recv_buffer, out_header, out_header.Length);
                        ClientHeader header = PackTools.UnPackObject(out_header);
                        this.long_pkg_size = Convert.ToInt32(header.Length);
                        this.long_pkg = new byte[header.Length];
                        //将上次的数据拷贝到这次的大包数组里
                        Array.Copy(this.recv_buffer, 0, this.long_pkg, 0, this.receved);
                    }

                    recv_len = this.client.Receive(this.recv_buffer, this.receved, RECV_LEN - this.receved,
                        SocketFlags.None);
                }

                if (recv_len > 0)
                {
                    this.receved += recv_len;
                    tcp_Func();
                }
            }
            catch (Exception e)
            {
                Debug.Log(e.Message);
                break;
            }
        }
    }

    private byte[] sendBuffer;
    void start_tcp_send_data()
    {
        if (this.is_Connect == false)
        {
            return;
        }
        
        while (true)
        {
            if (!this.client.Connected)
            {
                break;
            }

            try
            {

                if (this.sendQueue.Count == 0)
                {
                    
                }
                else
                {
                    List<byte> buffer = new List<byte>();
                    while (this.sendQueue.Count != 0)
                    {
                        Debug.Log("send data");
                        var sendBuffer = this.sendQueue.Peek();
                        if (buffer.Count + sendBuffer.Length > RECV_LEN)
                        {
                            break;
                        }

                        buffer.AddRange(sendBuffer);
                        this.sendQueue.Dequeue();
                    }
                
                    this.sendBuffer = buffer.ToArray();
                    IAsyncResult ar = this.client.BeginSend(this.sendBuffer,0,this.sendBuffer.Length,  SocketFlags.None,new AsyncCallback(this.on_send_end), this.client);
                    ar.AsyncWaitHandle.WaitOne(1000);
                }
            }
            catch (Exception e)
            {
                Debug.Log(e.Message);
                break;
            }
        }
    }

    void on_send_end(IAsyncResult ar)
    {
        try
        {
            Socket client = (Socket) ar.AsyncState;
            client.EndConnect(ar);
            this.on_log_msg("send success");
        }
        catch (Exception e)
        {
            this.on_log_msg(e.ToString());
        }
    }

    void tcp_Func()
    {
        var tcp_data = this.long_pkg != null ? this.long_pkg : this.recv_buffer;


        while (this.receved > 0)
        {
            ClientHeader header;
            try
            {
                var out_header = new byte[PackTools.PACK_LENGTH];
                Array.Copy(tcp_data, out_header, out_header.Length);
                header = PackTools.UnPackObject(out_header);
            }
            catch (Exception e)
            {
                Debug.Log(e.Message);
                break;
            }

            if (this.receved < header.Length)
            {
                break;
            }

            int pkgSize = Convert.ToInt32(header.Length);
            int raw_data_start = PackTools.PACK_LENGTH;
            int raw_data_len = pkgSize - PackTools.PACK_LENGTH;
            this.on_recv_tcp_cmd(tcp_data, raw_data_start, raw_data_len);

            if (this.receved > pkgSize)
            {
                this.recv_buffer = new byte[RECV_LEN];
                Array.Copy(tcp_data, pkgSize, this.recv_buffer, 0, this.receved - pkgSize);
            }

            this.receved -= pkgSize;
            if (this.receved == 0 && this.long_pkg != null)
            {
                this.long_pkg = null;
                this.long_pkg_size = 0;
            }
        }
    }

    void on_recv_tcp_cmd(byte[] tcp_data, int start, int data_len)
    {
        var out_header = new byte[PackTools.PACK_LENGTH];
        Array.Copy(tcp_data, out_header, out_header.Length);
        var header = PackTools.UnPackObject(out_header);

        var out_msg = new byte[data_len];
        Array.Copy(tcp_data, start, out_msg, 0, data_len);

        SMessage message = new SMessage(header, out_msg);
        on_recv_tcp_data(message);
    }


    void close()
    {
        if (!this.is_Connect)
        {
            return;
        }

        if (this.readThread != null)
        {
            this.readThread.Abort();
        }


        if (this.client != null && this.client.Connected)
        {
            this.client.Close();
        }
    }


    private void on_connected(IAsyncResult ar)
    {
        try
        {
            Socket client = (Socket) ar.AsyncState;
            client.EndConnect(ar);
            this.on_log_msg("success");

            if (this.readThread != null)
            {
                this.readThread.Abort();
            }

            this.is_Connect = true;
            EventManager.Instance.Call("OnLogin", null);
            this.readThread = new Thread(new ThreadStart(this.start_recv_data));
            this.readThread.Start();
            this.sendThread = new Thread(new ThreadStart(this.start_tcp_send_data));
            this.sendThread.Start();
        }
        catch (Exception e)
        {
            this.on_log_msg(e.ToString());
        }
    }
}