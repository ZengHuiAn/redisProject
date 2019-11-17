using System;
using System.Collections;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using UnityEngine;

public class network : MonoBehaviour
{
    public Socket client;
    public string server_ip;
    public int port;

    private Thread readThread;
    private bool is_Connect  = false;
    private byte[] recv_buffer = new byte[8296];
    // Start is called before the first frame update
    void Start()
    {
        connect2Server();
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

            IAsyncResult result = this.client.BeginConnect(ipEndPoint, new AsyncCallback(this.on_connected), this.client);
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
                int recv_len = this.client.Receive(this.recv_buffer);

                // 接受数据
                if (recv_len >0)
                {
                    
                }
            }
            catch (Exception e)
            {
                
                break;
            }

        }
    }

    void close()
    {
        if (!this.is_Connect)
        {
            return;
        }

        if (this.readThread !=null)
        {
            this.readThread.Abort();
        }


        if (this.client!=null && this.client.Connected)
        {
            this.client.Close();
        }
    }
    
    
    private void on_connected(IAsyncResult ar)
    {
        try
        {
            Socket client = (Socket)ar.AsyncState;
            client.EndConnect(ar);
            this.on_log_msg("success");

            if (this.readThread != null)
            {
                this.readThread.Abort();
            }
            this.is_Connect = true;
            this.readThread = new Thread(new ThreadStart(this.start_recv_data));
            this.readThread.Start();
        }
        catch (Exception e)
        {
            this.on_log_msg(e.ToString());
        }
    }
}
