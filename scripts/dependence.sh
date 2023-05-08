# gin 框架依赖
go get -u github.com/gin-gonic/gin

# gorm 框架依赖
go get -u gorm.io/gorm

# gorm 集成的 mysql 数据库驱动依赖
go get -u gorm.io/driver/mysql

# gorm 自动生成
go get -u gorm.io/gen


# viper 配置文件解析库依赖
go get -u github.com/spf13/viper

# validator 参数校验库依赖
go get github.com/go-playground/validator/v10

# zap 日志框架依赖
go get -u go.uber.org/zap

# lumberjack 日志切割库依赖, 与 zap 配合使用
go get -u github.com/natefinch/lumberjack

# gin-swagger gin 框架中使用 swagger 的依赖
go get -u github.com/swaggo/gin-swagger

# swagger 库使用的 swag 命令依赖
go get -u github.com/swaggo/files

# gin-contrib gin 框架跨域中间件依赖
go get github.com/gin-contrib/cors

# jwt 生成和解析框架
go get github.com/golang-jwt/jwt/v4

# redis 客户端框架
# go get -u github.com/go-redis/redis (redis5对应版本)
# go get github.com/go-redis/redis/v8 (redis6对应版本)
# go get github.com/go-redis/redis/v9 (redis7对应版本)
go get github.com/go-redis/redis/v8

# Golang WebSocket 第三方库
# go get github.com/gorilla/websocket

# decimal 数据类型支持, 数据库中的金额对象必须使用该类型的数据
go get github.com/shopspring/decimal

# 微信, 支付宝, PayPal, QQ 支付的 Go 语言 SDK(第三方)
# go get -u github.com/go-pay/gopay