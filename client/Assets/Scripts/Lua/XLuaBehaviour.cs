using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using  XLua;
using Object = System.Object;

public class XLuaBehaviour : MonoBehaviour
{
        [CSharpCallLua]
        public delegate void LuaObjectAction(object lauObject, params object[] args);
        public string luaScript;
        private LuaObjectAction  luaStart;
        private LuaObjectAction  luaUpdate;
        private LuaObjectAction  luaOnDestroy;

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
            var data = scriptEnv.GetInPath<LuaTable>("data");

            foreach (var VARIABLE in data.GetKeys())
            {
                var value = data.Get<object>(VARIABLE);
                var t = value.GetType();
                if (t == typeof(XLua.LuaTable))
                {
                    print(t);
                }
                
                
            }
//            Action luaAwake = scriptEnv.Get<Action>("awake");
            scriptEnv.Get("start",out luaStart);
            scriptEnv.Get("update", out luaUpdate);
            scriptEnv.Get("ondestroy", out luaOnDestroy);
            
        }

        // Use this for initialization
        void Start()
        {
            if (luaStart != null)
            {
                luaStart(scriptEnv);
            }
        }

        // Update is called once per frame
        void Update()
        {
            if (luaUpdate != null)
            {
                luaUpdate(scriptEnv);
            }
        }
        
        void OnDestroy()
        {
            if (luaOnDestroy != null)
            {
                luaOnDestroy(scriptEnv);
            }
            luaOnDestroy = null;
            luaUpdate = null;
            luaStart = null;
            scriptEnv.Dispose();
        }
}
