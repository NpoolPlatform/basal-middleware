package usercode

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/go-redis/redis/v8"
)

const (
	redisTimeout = 5 * time.Second
)

func key(c *npool.UserCode) string {
	return fmt.Sprintf(
		"%v:%v:%v:%v:%v",
		c.GetPrefix(),
		c.GetAppID(),
		c.GetAccountType(),
		c.GetAccount(),
		c.GetUsedFor())
}

func CreateUserCode(
	ctx context.Context,
	prefix, appID, account string,
	accountType basetypes.SignMethod,
	usedFor basetypes.UsedFor,
) (*npool.UserCode, error) {
	cli, err := redis2.GetClient()
	if err != nil {
		return nil, fmt.Errorf("fail get redis client: %v", err)
	}

	const codeLen = 6
	const expireMins = 10

	vCode := generate(codeLen)
	code := &npool.UserCode{
		Prefix:      prefix,
		AppID:       appID,
		Account:     account,
		AccountType: accountType,
		UsedFor:     usedFor,
		Code:        vCode,
		NextAt:      uint32(time.Now().Add(time.Minute).Unix()),
		ExpireAt:    uint32(time.Now().Add(expireMins * time.Minute).Unix()),
	}

	yes, err := nextable(
		ctx,
		prefix,
		appID,
		account,
		accountType,
		usedFor,
	)
	if err != nil {
		return nil, err
	}
	if !yes {
		return nil, fmt.Errorf("wait for next code generation")
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	body, err := json.Marshal(code)
	if err != nil {
		return nil, fmt.Errorf("fail marshal code: %v", err)
	}

	err = cli.Set(ctx, key(code), body, time.Until(time.Unix(int64(code.ExpireAt), 0))).Err()
	if err != nil {
		return nil, fmt.Errorf("fail create code cache: %v", err)
	}

	return code, nil
}

func VerifyUserCode(
	ctx context.Context,
	prefix, appID, account, vCode string,
	accountType basetypes.SignMethod,
	usedFor basetypes.UsedFor,
) error {
	cli, err := redis2.GetClient()
	if err != nil {
		return fmt.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	code := &npool.UserCode{
		Prefix:      prefix,
		AppID:       appID,
		Account:     account,
		AccountType: accountType,
		UsedFor:     usedFor,
		Code:        vCode,
	}

	val, err := cli.Get(ctx, key(code)).Result()
	if err == redis.Nil {
		return fmt.Errorf("code not found %v in redis", key(code))
	} else if err != nil {
		return fmt.Errorf("fail get code: %v", err)
	}

	user := npool.UserCode{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return fmt.Errorf("fail unmarshal val: %v", err)
	}

	if user.Code != code.Code {
		return fmt.Errorf("invalid code")
	}

	if time.Now().After(time.Unix(int64(user.ExpireAt), 0)) {
		return fmt.Errorf("code expired")
	}

	err = deleteUserCode(
		ctx,
		prefix,
		appID,
		account,
		accountType,
		usedFor,
	)
	if err != nil {
		return fmt.Errorf("fail delete code: %v", err)
	}

	return nil
}

func nextable(ctx context.Context,
	prefix, appID, account string,
	accountType basetypes.SignMethod,
	usedFor basetypes.UsedFor,
) (bool, error) {
	cli, err := redis2.GetClient()
	if err != nil {
		return false, fmt.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	code := &npool.UserCode{
		Prefix:      prefix,
		AppID:       appID,
		Account:     account,
		AccountType: accountType,
		UsedFor:     usedFor,
	}

	val, err := cli.Get(ctx, key(code)).Result()
	if err == redis.Nil {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("fail get code: %v", err)
	}

	err = json.Unmarshal([]byte(val), &code)
	if err != nil {
		return false, fmt.Errorf("fail unmarshal val: %v", err)
	}

	if !time.Now().After(time.Unix(int64(code.NextAt), 0)) {
		return false, nil
	}

	return true, nil
}

func deleteUserCode(ctx context.Context,
	prefix, appID, account string,
	accountType basetypes.SignMethod,
	usedFor basetypes.UsedFor,
) error {
	cli, err := redis2.GetClient()
	if err != nil {
		return fmt.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	code := &npool.UserCode{
		Prefix:      prefix,
		AppID:       appID,
		Account:     account,
		AccountType: accountType,
		UsedFor:     usedFor,
	}

	err = cli.Del(ctx, key(code)).Err()
	if err != nil {
		return fmt.Errorf("fail delete code: %v", err)
	}

	return nil
}
