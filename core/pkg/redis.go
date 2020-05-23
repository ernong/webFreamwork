package pkg

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"

	"time"

	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

var FriendshipRedisPrefix = "friendship_"            //从聊天db拉取的好友关系列表，好友圈查询
var UserRedisPrefix = "user_"                        // 缓存通过session获取token的用户信息，用于uid到名称，头像的转换
var UserFrontCoverPraisePrefix = "front_cover_"      //缓存封面点赞信息，cover_id-uid-0
var CycleOffsetRedisPrefix = "world_cycle_offset_"   //每个用户对应刷世界圈的时间偏移位置 world_cycle_offset_[user_id]-tableIndex-timestampOffset
var FriendOffsetRedisPrefix = "friend_cycle_offset_" //每个用户对应刷朋友圈的时间偏移位置 friend_cycle_offset_[user_id]-tableIndex-timestampOffset
var UserNameRedis = "username"                       //UserRedisPrefix中保存用户名
var UserAvatarRedis = "avatar"                       //UserRedisPrefix中保存头像

func LoadRedisConfig(viper *viper.Viper) *RedisConfig {
	cfg := &RedisConfig{
		Addr:     viper.GetString("host"),
		Password: viper.GetString("pass"),
		DB:       viper.GetInt("index"),
		PoolSize: viper.GetInt("pool_size"),
	}
	log.Info("LoadRedisConfig:%#v\n", cfg)
	return cfg
}

func (cfg *RedisConfig) InitRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	if redisCli == nil {
		log.Error("InitRedis RedisCli is nil")
		return nil
	}
	_, err := redisCli.Ping().Result()
	if err != nil {
		log.Error("InitRedis error:%v", err)
		return nil
	}
	return redisCli
}

var MaxErrCnt = 5

func SetRedisVal(key string, val interface{}, expiration time.Duration) error {
	buf, err := json.Marshal(val)
	if err != nil {
		log.Error("json.Marshal error", err)
		return err
	}

	cmd := RedisCli.Set(key, string(buf), expiration)
	if cmd.Err() != nil {
		log.Error("SetRedisVal error", cmd.Err())
		return cmd.Err()
	}
	return nil
}

func GetRedisVal(key string) string {
	tryCnt := 0
	for {
		value, err := RedisCli.Get(key).Result()
		if err == redis.Nil {
			return ""
		} else if err != nil {
			tryCnt++
			if tryCnt > MaxErrCnt {
				log.Error("RedisCli.Get error", err)
				return ""
			}
			continue
		}
		return value
	}
}

func RedisExists(key string) bool {
	errCnt := 0
	for {
		ret := RedisCli.Exists(key)
		if ret == nil || ret.Err() != nil {
			errCnt++
		} else {
			return ret.Val()
		}

		if errCnt > MaxErrCnt {
			log.Error("redisExists error", nil)
			break
		}
	}

	return false
}

/*
Redis Hash 操作
*/

//redisDel 删除一个KEY
func RedisHDel(strKey string, field ...string) bool {
	errCnt := 0
	for {
		_, err := RedisCli.HDel(strKey, field...).Result()
		if err == redis.Nil {
			return false
		} else if err != nil {
			errCnt++
			if errCnt > MaxErrCnt {
				log.Error("RedisCli.HDel error", err)
				return false
			}
			continue
		}
		return true
	}
}

//RedisHSet 设置redis值
func RedisHSet(strKey, strSubKey, strValue string) bool {
	errCnt := 0
	for {
		_, err := RedisCli.HSet(strKey, strSubKey, strValue).Result()
		if err == redis.Nil {
			log.Error("RedisCli.Hset error", err)
			return false
		} else if err != nil {
			errCnt++
			if errCnt > MaxErrCnt {
				log.Error("RedisCli.HDel error", err)
				return false
			}
			continue
		}
		return true
	}

}

//RedisHGet 获取redis值
func RedisHGet(strKey, strSubKey string) string {
	errCnt := 0
	for {
		strSubVal, err := RedisCli.HGet(strKey, strSubKey).Result()
		if err == redis.Nil {
			return ""
		} else if err != nil {
			errCnt++
			if errCnt > MaxErrCnt {
				log.Error("RedisCli.HDel error", err)
				return ""
			}
			continue
		}
		return strSubVal
	}

}

//RedisHGet 获取redis值
func RedisHGetAll(strKey string) map[string]string {
	errCnt := 0
	for {
		strSubVal, err := RedisCli.HGetAll(strKey).Result()
		if err == redis.Nil {
			return nil
		} else if err != nil {
			errCnt++
			if errCnt > MaxErrCnt {
				log.Error("RedisCli.HDel error", err)
				return nil
			}
			continue
		}
		return strSubVal
	}

}

//RedisHMSet 设置redis值
func RedisHMSet(strKey string, value map[string]string) bool {
	errCnt := 0
	for {
		_, err := RedisCli.HMSet(strKey, value).Result()
		if err == redis.Nil {
			return false
		} else if err != nil {
			errCnt++
			if errCnt > MaxErrCnt {
				log.Error("RedisCli.HDel error", err)
				return false
			}
			continue
		}
		return true
	}
}

//printHMSet 输出
func printHMSet(key string, fields map[string]string) {
	args := make([]interface{}, 2+len(fields)*2)
	args[0] = "hmset"
	args[1] = key
	i := 2
	for k, v := range fields {
		args[i] = k
		args[i+1] = v
		i += 2
	}

	log.Debugln(args)
}
