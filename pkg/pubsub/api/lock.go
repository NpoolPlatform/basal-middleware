package api

import (
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	prefix "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func key(servicename, protocol string) string {
	return fmt.Sprintf("%v:%v%v", prefix.Prefix_PrefixAPIRegister, servicename, protocol)
}

func Lock(servicename, protocol string) error {
	return redis2.TryLock(key(servicename, protocol), 0)
}

func Unlock(servicename, protocol string) error {
	return redis2.Unlock(key(servicename, protocol))
}
