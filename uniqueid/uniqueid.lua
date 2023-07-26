local members = redis.call('ZRANGEBYSCORE', KEYS[1], ARGV[1], ARGV[2])
local values = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
local machineID = ''

for _, value in ipairs(values) do
    local found = false
    for _, member in ipairs(members) do
        if tostring(value) == member then
            found = true
            break
        end
    end

    if not found then
        machineID = tostring(value)
        break
    end
end

if machineID == '' then
    return error("没有可分配的机器id")
else
    redis.call('ZADD', KEYS[1], ARGV[2], machineID)
end

redis.call('ZREMRANGEBYSCORE', KEYS[1], '-inf', ARGV[1])

return machineID
