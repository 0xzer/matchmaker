package util

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/nu7hatch/gouuid"
)

var Charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStr(length int) string {
	b := make([]rune, length)
    for i := range b {
        b[i] = Charset[rand.Intn(len(Charset))]
    }
    return string(b)
}

func RandomHex(n int) string {
	bytes := make([]byte, n)
	crand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func RandomUUIDv4() string {
	id, _ := uuid.NewV4()
	return id.String()
}

func NewDeviceTemplate() *Device {
	return &Device{
		DeviceId: RandStr(11),
		PersistentDeviceId: RandomHex(8),
		OsVersion: "31",
		Manufacturer: "samsung",
		Model: "SM-A217F",
		DeviceOsVersion: "Android 12",
		DataProvider: "wifi",
		AdvertisingId: RandomUUIDv4(),
		AuthSession: NewSession(),
	}
}

type Ts struct {
	Unix int64
	UnixMilli int64
}

func GetTimestamp() Ts {
	t := time.Now()
	return Ts{
		Unix: t.Unix(),
		UnixMilli: t.UnixMilli(),
	}
}

func FormatDateActivity(t time.Time) string {
	t = t.UTC()
	date := t.Format("2006-01-02T15:04:05.000Z")
	return date
}