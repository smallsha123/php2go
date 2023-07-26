package uniqueid

import (
	"fmt"
	"github.com/smallsha123/php2go/globalkey"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenId() (int64, string) {
	for i := 0; i <= 5; i++ {
		id, _ := flake.NextID()
		if int64(id) > 0 {
			return int64(id), time.Now().Format("200601")
		}
	}
	return 0, ""
}

// TimeUnique 根据时间生成唯一id 每分钟最多生成 99999 个唯一值
type TimeUnique struct {
	mu        sync.Mutex
	cache     map[string]struct{}
	machineID int    // 机器ID
	dateTime  string // 基础key
}

func NewTimeUnique(redis *redis.Redis) (*TimeUnique, error) {
	// 获取当前正在使用机器id
	newTime := time.Now()

	var exceedTime = time.Minute * 3
	var renovateTime = time.Minute * 1

	script := `
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
`
	machineID, err := redis.Eval(script, []string{globalkey.RedisOrderNumGenerate}, []interface{}{newTime.Add(-exceedTime).Unix(), newTime.Unix()})
	if err != nil {
		return nil, err
	}
	newMachineID, _ := strconv.Atoi(machineID.(string))

	// 启动一个goroutine，刷新当前id的过期时间
	go func() {
		for true {
			select {
			case <-time.Tick(renovateTime):
				redis.Zadd(globalkey.RedisOrderNumGenerate, time.Now().Unix(), strconv.Itoa(newMachineID))
			}
		}
	}()

	return &TimeUnique{
		cache:     make(map[string]struct{}),
		machineID: newMachineID,
		dateTime:  time.Now().Format("200601021504"),
	}, err
}

// Generate 生成一个新的订单号，格式为：日期(8) + 时分(4) + 机器ID(1) + 随机数(5)
func (g *TimeUnique) Generate() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	var orderNumber string
	for true {
		// 获取当前的日期和时间
		now := time.Now()

		var dateTime = now.Format("200601021504")
		if dateTime != g.dateTime {
			g.cache = make(map[string]struct{})
		}

		// 生成随机数
		rand.Seed(now.UnixNano())
		random := rand.Intn(100000)
		orderNumber = fmt.Sprintf("%s%01d%05d", dateTime, g.machineID, random)

		// 检查订单号是否已经在缓存中
		if _, ok := g.cache[orderNumber]; !ok {
			g.cache[orderNumber] = struct{}{}
			break
		}
	}

	return orderNumber
}
