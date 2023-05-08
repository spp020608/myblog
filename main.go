package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	_ "myblog/docs"
	"myblog/internal/router"
	"myblog/pkg/color"
	"myblog/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=weichuang-gin-api
const logo = `
██╗    ██╗███████╗██╗ ██████╗██╗  ██╗██╗   ██╗ █████╗ ███╗   ██╗ ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██║    ██║██╔════╝██║██╔════╝██║  ██║██║   ██║██╔══██╗████╗  ██║██╔════╝       ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║ █╗ ██║█████╗  ██║██║     ███████║██║   ██║███████║██╔██╗ ██║██║  ███╗█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║███╗██║██╔══╝  ██║██║     ██╔══██║██║   ██║██╔══██║██║╚██╗██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚███╔███╔╝███████╗██║╚██████╗██║  ██║╚██████╔╝██║  ██║██║ ╚████║╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚══╝╚══╝ ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝
`

func main() {

	// swagger 页面地址:   http://localhost:8080/swagger/index.html

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println(color.Blue(logo))

	s := router.NewHttpServer()

	server := &http.Server{
		Addr:    ":" + config.Get().ServerConfig.Port,
		Handler: s,
	}

	// 开启协程监听网络请求
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Errorf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)

	// kill
	// kill -2 syscall.SIGINT
	// kill -9 syscall.SIGKILL
	// kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.S().Errorf("shut down err: %s", err)
	}

	select {
	case <-ctx.Done():
		zap.L().Info("5秒超时退出")
	}

	zap.L().Info("服务器关闭")

}
