package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func SafeConvertTime(t *gtime.Time) *timestamppb.Timestamp {
	if t == nil || t.IsZero() {
		return nil
	}
	return timestamppb.New(t.Time)
}

// GenerateOrderNumber 生成订单编号
func GenerateOrderNumber() string {
	return fmt.Sprintf("ORD%s%04d", time.Now().Format("20060102150405"), rand.Intn(9999))
}
