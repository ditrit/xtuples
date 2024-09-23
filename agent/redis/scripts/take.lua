local filter = cjson.decode(ARGV[1])
local istrue = true 
local randomKey = ""
 
local function split(str, sep) 
    local sep, fields = sep or ':', {} 
    local pattern = string.format('([^%s]+)', sep) 
    string.gsub(str, pattern, function(c) fields[#fields+1] = c end) 
    return fields 
end 

if filter.empty ~= nil and type(filter.empty) ~= 'table' then istrue = false end 
if filter.one_of ~= nil and type(filter.one_of) ~= 'string' then istrue = false end  
if type(filter.step) ~= 'string' then istrue = false end 

if istrue and redis.pcall('get', 'step') ~= filter.step 
then 
    istrue = false 
else 
    if filter.empty ~= nil then 
        for _, item in ipairs(filter.empty) do 
            if redis.pcall('scard', item) >0 then  
                istrue = false 
                break 
            end 
        end 
    end

    if istrue then 
        randomKey = redis.pcall('srandmember', filter.one_of) 
        if randomKey == false then 
            randomKey = ""
            istrue = false 
        end 
    end 
end 

if istrue and randomKey then 
    local timeout = 10 
    local state = redis.pcall('hget', 'retries', randomKey) 
    if state == false then 
        state = '0:0' 
        redis.pcall('hset', 'retries', randomKey, state) 
    end 
    local stateParts = split(state, ':') 
    local retry = tonumber(stateParts[1]) 
    local timestamp = tonumber(stateParts[2]) 
    local now = tonumber(redis.pcall('time')[1])  
    local newretrystate = retry+1 .. ':' .. now 
    if timestamp == 0 or ((now > (timestamp + timeout)) and retry < 4) then 
        redis.pcall('hset', 'retries', randomKey, newretrystate) 
    else 
        istrue = false 
        if retry >= 4 then 
            redis.pcall('smove', filter.one_of, 'error', randomKey) 
            randomKey = ""
        end 
    end 
end 
local numbool = 0;
if istrue then 
    numbool = 1 
end
local ret_val= { numbool, randomKey }
return ret_val

