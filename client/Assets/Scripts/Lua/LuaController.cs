using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using XLua;

public class LuaController : MonoBehaviour
{
    private LuaEnv _luaEnv;

    public LuaEnv L
    {
        get { return _luaEnv; }
    }

    private static LuaController instance;
    
    
    public event Action onAwake ;
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


        var MainAction = L.Global.Get<System.Action>("main");
        if (MainAction != null)
        {
            MainAction();
        }
    }
}
