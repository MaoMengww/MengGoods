package cache

import (
	"context"
	"fmt"
)

type ScriptKey string

const (
	GetTTLAndDelPaymentTokenScriptKey ScriptKey = "getTTLAndDel"
	SetOrIncrRefundKeyScriptKey       ScriptKey = "setOrIncr"
)

type Script struct {
	Hash   string
	Script string
}

var scripts = map[ScriptKey]*Script{
	GetTTLAndDelPaymentTokenScriptKey: {
		Hash: string(GetTTLAndDelPaymentTokenScriptKey),	
		Script: `
			local exists = redis.call("EXISTS", KEYS[1])

			if exists == 1 then
   		 		if redis.call("GET", KEYS[1]) == ARGV[1] then
        			local ttl = redis.call("TTL", KEYS[1])
        			redis.call("DEL", KEYS[1])
        			return {ttl, 1} -- 返回 ttl 1代表成功删除
    			end
			end
			return {-1, 0}
		`,
	},
	SetOrIncrRefundKeyScriptKey: {
		Hash: string(SetOrIncrRefundKeyScriptKey),	
		Script: `
    			local current = redis.call("INCR", KEYS[1])
				if current == 1 then
    				redis.call("EXPIRE", KEYS[1], 60)
				end
				return current
		`,
	},
}

func (p *paymentRedis) loadScript() (err error) {
	ctx := context.Background()
	for key, value := range scripts {
		hash, err := p.client.ScriptLoad(ctx, value.Script).Result()
		if err != nil {
			return fmt.Errorf("load script %s failed: %w", key, err)
		}
		value.Hash = hash
	}
	return nil
}
