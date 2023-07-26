package auth

import (
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	str := GenToken("smartkefu", time.Now().Unix(), 86400, nil)
	t.Logf("str:%s", str)
}
