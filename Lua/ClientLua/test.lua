---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by anzenghui.
--- DateTime: 2019/12/5 16:28
---
local V = {}

function V:start()
    print("start")

    print(RegisterEvent)
    RegisterEvent("EVENT.NET.MESSAGE.102", function(...)
        local args = ...

        print(type(args),args)
    end)
end
return V