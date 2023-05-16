package api

import (
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	prefix "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func key(servicename string) string {
	return fmt.Sprintf("%v:%v", prefix.Prefix_PrefixAPIRegister, servicename)
}

func Lock(servicename string) error {
	return redis2.TryLock(key(servicename), 0)
}

func Unlock(servicename string) error {
	return redis2.Unlock(key(servicename))
}
