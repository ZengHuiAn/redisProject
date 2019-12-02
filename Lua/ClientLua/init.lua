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