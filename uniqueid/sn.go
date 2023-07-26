package uniqueid

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"sync"
	"time"

	"github.com/smallsha123/php2go/tool"
)

// 生成sn单号
type SnPrefix string

const (
	SnPrefixHomestayOrder      SnPrefix = "HSO" // 民宿订单前缀 looklook_order/homestay_order
	SN_PREFIX_THIRD_PAYMENT    SnPrefix = "PMT" // 第三方支付流水记录前缀 third_payment
	SN_PREFIX_ORDER            SnPrefix = "TKO" // 平台订单
	SN_PREFIX_PAY              SnPrefix = "TKP" // 平台支付单
	SN_PREFIX_PAY_SPLIT        SnPrefix = "PS"  // 平台支付单分账
	SN_AD_PATH                 SnPrefix = "AD"
	SnPrefixSplitCreateOrder   SnPrefix = "SC" // 创建分账订单
	SnPrefixSplitUnfreezeOrder SnPrefix = "SU" // 解冻分账订单剩余资金
	SnPrefixSplitReturnOrder   SnPrefix = "SR" // 订单金额回退

)

// GenSn 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(8, tool.KC_RAND_KIND_NUM))
}

// RefundSn 生成售后单号
func RefundSn() string {

	return fmt.Sprintf("%s%s", time.Now().Format("200601021504"), tool.Krand(6, tool.KC_RAND_KIND_NUM))
}

func OrderSn(shopId int64) string {
	return fmt.Sprintf("%s%d%d%s", time.Now().Format("20060102150406"), time.Now().UnixMilli()%1000, shopId, tool.Krand(6, tool.KC_RAND_KIND_NUM))
}

func GenProductSn() string {
	return fmt.Sprintf("%s%s", time.Now().Format("200601021504"), RandProductStr(4))
}

func GetAdSn() string {
	return fmt.Sprintf("%s%v", time.Now().Format("010205"), GetRandSn(9999, 1000))
}

// GetRandSn 生成随机数
func GetRandSn(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandProductStr(length int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GenShopSN 店铺ID为创建日期（20221208）+纳税人识别码后5位+数字自增
func GenShopSN(s string, length int) string {
	if len(s) < length {
		newId, _ := GenId()
		return cast.ToString(newId)
	}
	num := s[len(s)-length : len(s)]
	return fmt.Sprintf("%s%s", time.Now().Format("20060102"), num)
}

const (
	epoch             = 1288834974657
	workerBits        = 6
	datacenterBits    = 6
	maxWorkerID       = -1 ^ (-1 << workerBits)
	maxDatacenterID   = -1 ^ (-1 << datacenterBits)
	sequenceBits      = 12
	workerIDShift     = sequenceBits
	datacenterIDShift = sequenceBits + workerBits
	timestampShift    = sequenceBits + workerBits + datacenterBits
	sequenceMask      = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	datacenterID  int64
	sequence      int64
}

func NewSnowflake(workerID, datacenterID int64) *Snowflake {
	if workerID < 0 || workerID > maxWorkerID {
		panic(fmt.Sprintf("Worker ID must be between 0 and %d", maxWorkerID))
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		panic(fmt.Sprintf("Datacenter ID must be between 0 and %d", maxDatacenterID))
	}
	return &Snowflake{
		lastTimestamp: -1,
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
	}
}

func (s *Snowflake) NextID() (int64, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano()/1000000 - epoch

	if timestamp < s.lastTimestamp {
		panic(fmt.Sprintf("Clock is moving backwards. Rejecting requests until %d", s.lastTimestamp))
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			timestamp = s.waitNextMillis(timestamp)
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	return (timestamp << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence, time.Now().Format("200601")
}

func (s *Snowflake) waitNextMillis(timestamp int64) int64 {
	for timestamp == s.lastTimestamp {
		time.Sleep(1 * time.Millisecond)
		timestamp = time.Now().UnixNano()/1000000 - epoch
	}
	return timestamp
}
