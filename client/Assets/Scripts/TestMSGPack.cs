using System;
using MessagePack;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using UnityEngine;


public class TestMSGPack : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
//        var obj = new object[]
//        {
//            new char[]
//            {
//                'a',
//                '增',
//                '怕',
//                '%'
//            },
//        };
        LogTool.Instance.ToStringAll(BitConverter.GetBytes((Int16) (50)));
        ClientHeader header = new ClientHeader();
        SerializableStructAttribute[] objects =
            (SerializableStructAttribute[]) header.GetType()
                .GetCustomAttributes(typeof(SerializableStructAttribute), false);

        foreach (var VARIABLE in objects)
        {
            Debug.Log(VARIABLE);
        }
        EventManager.Instance.AddEventAction("EVENT.NET.MESSAGE", (msg) =>
        {
            SMessage m = msg as SMessage;
            print(m.Header.MessageID);
            EventManager.Instance.Call($"EVENT.NET.MESSAGE.{m.Header.MessageID}",1.5);
        });
//
//        var bs = NetPackData.pack_all(obj);
////        
////        LogTool.Instance.ToStringAll((byte)('a'));
////        LogTool.Instance.ToStringAll(BitConverter.GetBytes(false));
////        LogTool.Instance.ToStringAll(bs);
//
//        var obj_oo = NetUnPackData.unpack_all(bs); 
//        LogTool.Instance.ToStringAll(obj_oo);
//        .Instance.LogBytes(bs);
    }

    public UnityEngine.UI.InputField[] InputFields;

    public void testEvent()
    {
        EventManager.Instance.Call("OnLogin",null);
    }

    private static int sn = 0;

    public void sendNetData()
    {

        var obj = new object[]
        {
            sn+=1,
            InputFields[0].text,
            InputFields[1].text
        };
        var buffer =  NetPackData.pack_all(obj);

        var header = new ClientHeader()
        {
            Length = Convert.ToUInt32(PackTools.PACK_LENGTH) + Convert.ToUInt32(buffer.Length),
            Flag = 2,
            MessageID = 101,
            ProtoType = 1,
        };
        
        network.instance.start_send_data(header,buffer);
    }

}

