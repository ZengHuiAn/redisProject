using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using  XLua;
public class XLuaBehaviour : MonoBehaviour
{
        public string luaScript;
        private Action luaStart;
        private Action luaUpdate;
        private Action luaOnDestroy;

        private LuaTable scriptEnv;

        private LuaEnv luaEnv
        {
            get { return LuaController.Instance.L; }
        }
        void Awake()
        {

            
            // 为每个脚本设置一个独立的环境，可一定程度上防止脚本间全局变量、函数冲突
            

            object[] obj =  luaEnv.DoString(FileUtils.ReadFileContent(luaScript) , luaScript);
            if (obj !=null && obj.Length >0)
            {
                scriptEnv = obj[0] as LuaTable;
            }

            if (scriptEnv == null)
            {
                return;
            }
            print(obj.Length);
//            Action luaAwake = scriptEnv.Get<Action>("awake");
            scriptEnv.Get("start", out luaStart);
            scriptEnv.Get("update", out luaUpdate);
            scriptEnv.Get("ondestroy", out luaOnDestroy);
            
        }

        // Use this for initialization
        void Start()
        {
            if (luaStart != null)
            {
                luaStart();
            }
        }

        // Update is called once per frame
        void Update()
        {
            if (luaUpdate != null)
            {
                luaUpdate();
            }
        }
        
        void OnDestroy()
        {
            if (luaOnDestroy != null)
            {
                luaOnDestroy();
            }
            luaOnDestroy = null;
            luaUpdate = null;
            luaStart = null;
            scriptEnv.Dispose();
        }
}
