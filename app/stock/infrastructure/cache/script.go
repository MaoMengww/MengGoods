package cache

import (
	"context"
	"fmt"
)

type ScriptKey string

const (
	ReduceKey ScriptKey = "reduceStock"
)

type Script struct {
	Hash   string
	Script string
}

var scripts = map[ScriptKey]*Script{
	ReduceKey: {
		Hash: string(ReduceKey),	
		Script: `
			local failed_status = -1
			for i = 1, #KEYS do
				local stock = tonumber(redis.call('get', KEYS[i]))
				local count = tonumber(ARGV[i])
				if stock < count then
					failed_status = i
					break
				end
			end
			if failed_status ~= -1 then
				return -1
			end
			for i = 1, #KEYS do
				redis.call('decrby', KEYS[i], ARGV[i])
			end
			return 1
			`,
		},
}

func (s *StockCache) loadScript() (err error) {
	ctx := context.Background()
	for key, value := range scripts {
		hash, err := s.Client.ScriptLoad(ctx, value.Script).Result()
		if err != nil {
			return fmt.Errorf("load script %s failed: %w", key, err)
		}
		value.Hash = hash
	}
	return nil
}