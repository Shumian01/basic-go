local key=KEYS[1]
--用户输入的code

local expectedCode=ARGV[1]
local cntKey=key..":cnt"

--转成数字
local cnt=tonumber(redis.call("get",cntKey))

if cnt<=0 then
--    //说明用户一直输错
    -- 或者已经用过
    return -1
elseif expectedCode == code then
    --输入了
    --用完了 不能再用了
    redis.call("set",cntKey,-1)
    --redis.call("del",key)
    return 0
else
    --输错了
    --可验证次数减一
    redis.call("decr",cntKey,-1)
    return -2
end