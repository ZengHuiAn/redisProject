using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LogTool 
{
    public string BytesToString(byte[] buffer)
    {

        if (buffer == null)
        {
            return "";
        }
        string str = "";
        for (int i = 0; i < buffer.Length; i++)
        {
            str += (buffer[i]+" ");
        }

        return str;
    }

    public void LogBytes(byte[] buffer)
    {
        Debug.Log(this.BytesToString(buffer));
    }

    private static LogTool instance;
    public static LogTool Instance
    {
        get
        {
            if (instance == null)
            {
                instance = new LogTool();
            }

            return instance;
        }
    }
}
