package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"myblog/cmd/gorm/model"
	"myblog/pkg/config"

	"time"
)


const (
	LoginUserKey             = "prj:login:user:" // 已登陆用户在 redis 中存放的 key, 该 key + username 构成完整的 key
	LoginUserExpiredDuration = time.Hour * 24    // 用户登录信息在缓存中存在的时长
)

var rdb *redis.Client

func New() error {
	cfg := config.Get().RedisConfig

	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		return err
	}
	return nil
}

//
// Set 将键值对存入 redis 中
//	@parameter	key 键
//	@parameter	value 值
//	@parameter	ttl 过期时间,这里是一个时间段,也就是这个键值对儿在 redis 中的存活时长
//	@return		error 错误信息
//
func Set(key, value string, ttl time.Duration) error {
	if err := rdb.Set(context.Background(), key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}
	return nil
}

//
// Get 从 redis 中获取 key 对应的 value
//	@parameter	key 键
//	@return		string key 在 redis 中对应的 value
//	@return		error 错误信息
//
func Get(key string) (string, error) {
	value, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}
	return value, nil
}

//
// SaveLoginUser 将登录成功的用户保存到 redis
//	@parameter	user 当前登录用户对象
//	@return		error
//
func SaveLoginUser(user *model.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		zap.S().Errorf("将 user 对象序列化成 json 失败: %s", err)
		return err
	}
	if err = Set(LoginUserKey+user.Username, string(bytes), LoginUserExpiredDuration); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

//
// GetLoginUser 获取存放到 redis 中的已登录用户信息
//	@parameter	username
//	@return		*model.User
//
func GetLoginUser(username string) (*model.User, error) {
	var u model.User
	result, err := Get(LoginUserKey + username)
	if err != nil {
		if err == redis.Nil {
			zap.S().Errorf("从缓存中获取 %s 的值失败:", LoginUserKey+username)
			return nil, err
		}
	}
	if err = json.Unmarshal([]byte(result), &u); err != nil {
		zap.S().Errorf("将缓存中读取出的数据反序列化成 user 对象失败: %s", err)
		return nil, err
	}
	return &u, nil
}