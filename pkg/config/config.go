package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//
// ServerConfig 服务器相关配置, 配置文件中的 internal 子树
//
type ServerConfig struct {
	Port string `json:"port" mapstructure:"port"` // 服务器监听端口
}

//
// MySqlConfig MySql 数据库配置信息, 配置文件中 mysql 子树对应的结构体
//
type MySqlConfig struct {
	Host               string `json:"host" mapstructure:"host"`                               // 数据库所在服务器 IP
	Port               int    `json:"port" mapstructure:"port"`                               // mysql 端口
	DatabaseName       string `json:"databaseName" mapstructure:"databaseName"`               // 数据库名称
	Username           string `json:"username" mapstructure:"username"`                       // 用户名
	Password           string `json:"password" mapstructure:"password"`                       // 密码
	MaxOpenConnections int    `json:"max_open_connections" mapstructure:"maxOpenConnections"` // 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"maxIdleConnections"` // 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	ConnMaxLifeTime    int    `json:"conn_max_life_time" mapstructure:"connMaxLifeTime"`      // 设置最大连接超时
}

//
// RedisConfig redis 相关配置, 配置文件中 redis 子树对应的结构体
//
type RedisConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`         // redis 主机
	Password string `json:"password" mapstructure:"password"` // redis 密码
	Database int    `json:"database" mapstructure:"database"` // 过期时间(单位:秒)
	Timeout  int64  `json:"timeout" mapstructure:"timeout"`   // 连接超时时间
}

//
// ZapConfig zap 日志库配置, 配置文件中 logging 子树对应的结构体
//
type ZapConfig struct {
	Filename   string `json:"filename" mapstructure:"filename"`      // 日志文件名称
	MaxSize    int    `json:"max_size" mapstructure:"maxSize"`       // 日志文件所占空间大小(单位:M)
	MaxBackups int    `json:"max_backups" mapstructure:"maxBackups"` // 备份的最大数量
	MaxAge     int    `json:"max_age" mapstructure:"maxAge"`         // 日志文件存活时长(单位:天)
	Compress   bool   `json:"compress" mapstructure:"compress"`      // 日志备份的时候, 最大备份数量
	LocalTime  bool   `json:"local_time" mapstructure:"localTime"`   // 日志备份的时候, 是否使用本地时间作为备份文件名, 如果不使用本地时间的话, 会有 8 个小时的时差
}

//
// JwtConfig Jwt 相关配置信息, 配置文件中 jwt 子树对应的结构体
//
type JwtConfig struct {
	TokenHeader string `json:"token_header" mapstructure:"tokenHeader"` // 请求头 key 的名称
	Secret      string `json:"secret" mapstructure:"secret"`            // 用于 Jwt 计算签名的字符串
	Expire      int64  `json:"expire" mapstructure:"expire"`            // 过期时长(单位:秒)
	TokenHead   string `json:"token_head" mapstructure:"tokenHead"`     // 请求头内容中应该携带的头信息,本程序要求内容以 Bearer+空格 开头
}

// ApplicationConfig 整个配置文件对应的结构体
type ApplicationConfig struct {
	ServerConfig ServerConfig `json:"server" mapstructure:"server"`   // http 服务器相关配置
	MySqlConfig  MySqlConfig  `json:"mysql" mapstructure:"mysql"`     // mysql 子树对应的结构体
	LogConfig    ZapConfig    `json:"logging" mapstructure:"logging"` // logging 子树对应的结构体
	RedisConfig  RedisConfig  `json:"redis" mapstructure:"redis"`     // redis 子树对应的结构体
	JwtConfig    JwtConfig    `json:"jwt" mapstructure:"jwt"`         // jwt 子树对应的结构体
}

//
// Config 另外一种定义整个配置文件对应的结构体的方式(使用匿名结构体的方式)
//
type Config struct {
	ServerConfig struct {
		Port int `json:"port" mapstructure:"port"` // 服务器监听端口
	} `json:"server" mapstructure:"server"` // http 服务器相关配置

	MySqlConfig struct {
		Host               string `json:"host" mapstructure:"host"`                               // 数据库所在服务器 IP
		Port               int    `json:"port" mapstructure:"port"`                               // mysql 端口
		DatabaseName       string `json:"databaseName" mapstructure:"databaseName"`               // 数据库名称
		Username           string `json:"username" mapstructure:"username"`                       // 用户名
		Password           string `json:"password" mapstructure:"password"`                       // 密码
		MaxOpenConnections int    `json:"max_open_connections" mapstructure:"maxOpenConnections"` // 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
		MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"maxIdleConnections"` // 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
		ConnMaxLifeTime    int    `json:"conn_max_life_time" mapstructure:"connMaxLifeTime"`      // 设置最大连接超时
	} `json:"mysql" mapstructure:"mysql"` // mysql 子树对应的结构体

	LogConfig struct {
		Filename   string `json:"filename" mapstructure:"filename"`      // 日志文件名称
		MaxSize    int    `json:"max_size" mapstructure:"maxSize"`       // 日志文件所占空间大小(单位:M)
		MaxBackups int    `json:"max_backups" mapstructure:"maxBackups"` // 备份的最大数量
		MaxAge     int    `json:"max_age" mapstructure:"maxAge"`         // 日志文件存活时长(单位:天)
		Compress   bool   `json:"compress" mapstructure:"compress"`      // 日志备份的时候, 最大备份数量
		LocalTime  bool   `json:"local_time" mapstructure:"localTime"`   // 日志备份的时候, 是否使用本地时间作为备份文件名, 如果不使用本地时间的话, 会有 8 个小时的时差
	} `json:"logging" mapstructure:"logging"` // logging 子树对应的结构体

	RedisConfig struct {
		Addr     string `json:"addr" mapstructure:"addr"`         // redis 主机
		Password string `json:"password" mapstructure:"password"` // redis 密码
		Database int    `json:"database" mapstructure:"database"` // 过期时间(单位:秒)
		Timeout  int64  `json:"timeout" mapstructure:"timeout"`   // 连接超时时间
	} `json:"redis" mapstructure:"redis"` // redis 子树对应的结构体

	JwtConfig struct {
		TokenHeader string `json:"token_header" mapstructure:"tokenHeader"` // 请求头 key 的名称
		Secret      string `json:"secret" mapstructure:"secret"`            // 用于 Jwt 计算签名的字符串
		Expire      int64  `json:"expire" mapstructure:"expire"`            // 过期时长(单位:秒)
		TokenHead   string `json:"token_head" mapstructure:"tokenHead"`     // 请求头内容中应该携带的头信息,本程序要求内容以 Bearer+空格 开头
	} `json:"jwt" mapstructure:"jwt"` // jwt 子树对应的结构体
}

const (
	defaultConfigFileName = "./configs/config.yaml" // 默认的配置文件名称
	defaultConfigFileType = "yaml"                  // 默认的配置文件类型
	defaultConfigPath     = "./configs/"            // 默认的配置文件存放路径
)

const (
	DigitDateTimeLayout = "20060102150405"
	DateTimeLayout      = "2006-01-02 15:04:05"
	DateLayout          = "2006-01-02"
	TimeLayout          = "15:04:05"
)

const UploadsPath = "./uploads/" // 上传文件存放路径

// appConfig 应用程序配置, 对应整个配置文件
var appConfig = &ApplicationConfig{}

//
// init 包级别的初始化函数
//
func init() {
	if err := initApplicationConfig(); err != nil {
		_ = fmt.Errorf("读取配置文件失败: %s\n", err)
	}
	fmt.Println("init configs file...")
}

//
// initApplicationConfig 初始化整个应用程序的配置, 读取配置文件
//
func initApplicationConfig() error {
	viper.SetConfigFile(defaultConfigFileName)     // 设置配置文件名称
	viper.SetConfigType(defaultConfigFileType)     // 设置配置文件类型
	viper.AddConfigPath(defaultConfigPath)         // 添加配置文件路径
	viper.WatchConfig()                            // 启动 viper 配置文件监控
	viper.OnConfigChange(func(in fsnotify.Event) { // 当配置文件改变时候的回调函数定义
		err := readConfigFile() // 当配置文件改变的时候, 重新读取配置文件, 并重新序列化各个子树
		if err != nil {
			fmt.Println("读取 configs 文件失败:", err)
		}
	})

	err := readConfigFile() // 读取配置文件, 并重新序列化各个子树

	if err != nil {
		return err
	}
	return nil
}

//
// Get 获取整个配置
//
func Get() *ApplicationConfig {
	return appConfig
}

//
// readConfigFile 读取配置文件中的内容, 并将各个子树序列化成对应的结构体
//
func readConfigFile() error {
	// 读取配置文件内容
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("配置文件没找到", err)
		} else {
			fmt.Println("配置文件找到了", err)
		}
		return err
	}

	// 将整个配置文件反序列化到结构体
	if err := viper.Unmarshal(appConfig); err != nil {
		return err
	}

	return nil
}
