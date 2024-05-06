package rediser

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	// 对应的可使用任务数量减1
	CompanyTaskDeductionLua = `
	local test = tonumber(redis.call('GET', KEYS[1]))
	if test and test > 0 then
		redis.call('DECR', KEYS[1])
		return 1
	else
		return 0
	end`

	// 对应脚本的 sha
	CompanyTaskDeductionLuaSHA = ""
)

// UpdateCacheHash 执行 lua 扣减脚本
func EvalLua(ctx context.Context, rdb *redis.Client, pre string) error {
	res, err := rdb.Eval(ctx, CompanyTaskDeductionLua, []string{pre}).Result()

	if err != nil {
		return err
	}

	b, err := json.Marshal(res)

	if err != nil {
		return err
	}

	if string(b) == "0" {
		return errors.New("decr task err")
	}

	return nil
}

// 加载脚本，保存对应 sha ,先检查该 sha 的脚本是否存在
func LoadLua(ctx context.Context, rdb *redis.Client) error {
	exist, err := rdb.ScriptExists(ctx, CompanyTaskDeductionLuaSHA).Result()

	if err != nil {
		return err
	}

	log.Println("exists:", exist)

	if len(exist) > 0 && exist[0] {
		return errors.New("exist")
	}

	res, err := rdb.ScriptLoad(ctx, CompanyTaskDeductionLua).Result()

	if err != nil {
		return err
	}

	CompanyTaskDeductionLuaSHA = res

	log.Println("CompanyTaskDeductionLuaSHA:", CompanyTaskDeductionLuaSHA)
	return nil
}

// 执行对应 sha 的脚本
// 一般使用这种方法，先把脚本传上去，再执行对应的sha 就好
func EvalSha(ctx context.Context, rdb *redis.Client, key string) error {
	res, err := rdb.EvalSha(ctx, CompanyTaskDeductionLuaSHA, []string{key}).Result()

	if err != nil {
		return err
	}

	b, err := json.Marshal(res)

	if err != nil {
		return err
	}

	log.Println("lua:", err, string(b))
	log.Println("CompanyTaskDeductionLuaSHA:", CompanyTaskDeductionLuaSHA)
	return nil
}
