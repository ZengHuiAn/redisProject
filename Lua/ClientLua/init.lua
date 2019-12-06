print("init--------->>>>")
print("----->>>>")
local timer = 10
function main()
    --setmetatable(_G, {__index=function(_, k)
    --    print("GLOBAL NAME:".. k, "NOT EXISTS",debug.traceback())
    --end, __newindex = function(t, k, v)
    --    print("SET GLOBAL NAME", k, v, debug.traceback())
    --    rawset(t, k, v);
    --end})
    print("----->>>>")
    for i = 1, 100 do
        timer = timer +1
    end

    --UnityEngine.SceneManagement.SceneManager.LoadScene("LoginScene");
end

lua_class = require("class")




-- 打印表的格式的方法
local function _sprinttb(tb, tabspace)
    tabspace =tabspace or ''
    local str =string.format(tabspace .. '{\n' )
    for k,v in pairs(tb or {}) do
        if type(v)=='table' then
            if type(k)=='string' then
                str =str .. string.format("%s%s =\n", tabspace..'  ', k)
                str =str .. _sprinttb(v, tabspace..'  ')
            elseif type(k)=='number' then
                str =str .. string.format("%s[%d] =\n", tabspace..'  ', k)
                str =str .. _sprinttb(v, tabspace..'  ')
            end
        else
            if type(k)=='string' then
                str =str .. string.format("%s%s = %s,\n", tabspace..'  ', tostring(k), tostring(v))
            elseif type(k)=='number' then
                str =str .. string.format("%s[%s] = %s,\n", tabspace..'  ', tostring(k), tostring(v))
            end
        end
    end
    str =str .. string.format(tabspace .. '},\n' )
    return str
end

function sprinttb(tb, tabspace)
    if CS.UnityEngine.Application.isEditor then
        local function ss()
            return _sprinttb(tb, tabspace);
        end
        return setmetatable({}, {
            __concat = ss,
            __tostring = ss,
        });
    else
        return "";
    end
end