print("init--------->>>>")

function main()
    fmt.Println("main")
    setmetatable(_G, {__index=function(_, k)
        fmt.Error("GLOBAL NAME:".. k, "NOT EXISTS",debug.traceback())
    end, __newindex = function(t, k, v)
        fmt.Error("SET GLOBAL NAME", k, v, debug.traceback())
        rawset(t, k, v);
    end})
    UnityEngine.SceneManagement.SceneManager.LoadScene("LoginScene");
end