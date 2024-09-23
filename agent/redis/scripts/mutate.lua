
local success = tonumber(ARGV[1])
local onsuccess = cjson.decode(ARGV[2])
local key = ARGV[3]



if success > 0
then 
    if onsuccess ~= nil and type(onsuccess) == 'table' then 
        if onsuccess.step ~= nil and type(onsuccess.step) == 'string' then
            redis.pcall('set', 'step', onsuccess.step) 
        end 
        if onsuccess.state ~= nil and type(onsuccess.state  ) == 'string' and  
           onsuccess.old_state ~= nil and type(onsuccess.old_state  ) == 'string'
        then 
            redis.pcall('smove', onsuccess.old_state, onsuccess.state, key)
            redis.pcall('hset', 'retries', key, '0:0')
        end
    end
end
return success
