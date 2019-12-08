using Boo.Lang;
using UnityEngine;
using XLua;

namespace Util
{
    public class LuaConvert 
    {
        public static object ConvertToObject(LuaTable target)
        {
            var resullt = new System.Collections.Generic.Dictionary<object,object>();
            
            foreach (var VARIABLE in target.GetKeys())
            {
                var value = target.Get<object,object>(VARIABLE);
                
                
                var t = value.GetType();
                if (t == typeof(XLua.LuaTable))
                {
                    resullt[VARIABLE] = ConvertToObject(value as LuaTable);
                }
                else
                {
                    resullt[VARIABLE] = value;
                }
            }

            return resullt;
        }
        

        public static object NetConvertToObject(LuaTable target)
        {
            return ConvertSendToObject(target);
        }
        
        public static object ConvertSendToObject(LuaTable target)
        {
            var result = new System.Collections.Generic.List<object>();
            
            foreach (var VARIABLE in target.GetKeys())
            {
                var value = target.Get<object,object>(VARIABLE);
                
                
                var t = value.GetType();
                if (t == typeof(XLua.LuaTable))
                {
                    result.Add(ConvertSendToObject(value as LuaTable));
                }
                else
                {
                    result.Add(value);
                }
            }

            return result.ToArray();
        }
    }
}