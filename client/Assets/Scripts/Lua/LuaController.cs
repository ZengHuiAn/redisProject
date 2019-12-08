using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using XLua;

public class LuaController : MonoBehaviour
{
    private LuaEnv _luaEnv;
    internal static float lastGCTime = 0;
    internal const float GCInterval = 1;//1 second 
    public LuaEnv L
    {
        get { return _luaEnv; }
    }

    private static LuaController instance;

    public static LuaController Instance
    {
        get { return instance; }
    }

    void Awake()
    {
        if (instance !=null)
        {
            DestroyImmediate(gameObject);
            return;
        }

        instance = this;
        
        _luaEnv = new LuaEnv();
        
        L.AddLoader(FileUtils.LuaLoader);
        
        L.DoString("require 'init'");
        
        L.Global.Set<string,Action<string,Action<object>>>("RegisterEvent",EventManager.Instance.AddEventAction);

        var MainAction = L.Global.Get<System.Action>("main");
        if (MainAction != null)
        {
            MainAction();
        }
    }
    
    void Update()
    {
        if (Time.time - LuaController.lastGCTime > GCInterval)
        {
            L.Tick();
            LuaController.lastGCTime = Time.time;
        }
    }
}
