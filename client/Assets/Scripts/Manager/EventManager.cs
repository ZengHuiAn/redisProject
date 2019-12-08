using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EventManager
{
    static EventManager _instance;

    public static EventManager Instance
    {
        get
        {
            if (_instance == null)
            {
                _instance = new EventManager();
            }

            return _instance;
        }
    }

    Dictionary<string, List<Action<object>>> _dic = new Dictionary<string, List<Action<object>>>();

    public void AddEventAction(string eventName, Action<object> callBack)
    {
        if (!_dic.ContainsKey(eventName))
        {
            _dic[eventName] = new List<Action<object>>();
        }

        _dic[eventName].Add(callBack);
    }

    public void RemoveEventAction(string eventName, Action<object> callBack)
    {
        if (_dic.ContainsKey(eventName))
        {
            if (_dic[eventName] != null)
            {
                _dic[eventName].Remove(callBack);    
            }
        }
    }

    public void Call(string eventName, object data)
    {
        if (_dic.ContainsKey(eventName))
        {
            if (_dic[eventName] != null)
            {
                for (int i = 0; i < _dic[eventName].Count; i++)
                {
                    _dic[eventName][i].Invoke(data);
                } 
            }
        }
    }
}
