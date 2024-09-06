package v1

import (
	"fmt"
	"go-huawei_cloud_token_manager/config"
	huawei_cloud_controller "go-huawei_cloud_token_manager/internal/controller"
	"go-huawei_cloud_token_manager/internal/database"
	"go-huawei_cloud_token_manager/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.Cfg.AppMode)
	router := gin.Default()
	router.Use(cors.Default())

	TokenRouter := router.Group("/huawei")
	{
		TokenRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, utils.SendResult(200, fmt.Sprintf("server@%s start successful!", config.Version), nil))
		})

		TokenRouter.GET("/getToken", func(c *gin.Context) {
			code, tokenInfo := database.GetAvailableToken()
			if code == 0 {
				c.JSON(200, utils.SendResult(200, "获取成功", tokenInfo))
				return
			}
			// 缓存的token已过期，需更新
			code, newToken := huawei_cloud_controller.GetNewToken()

			// 更新成功
			if code == 0 {
				_, tokenInfo := database.UpdateToken(newToken)
				c.JSON(200, utils.SendResult(200, "获取成功", tokenInfo))
				return
			}
			c.JSON(200, utils.SendResult(4001, "获取失败", nil))
		})

	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, utils.SendResult(404, fmt.Sprintf("go-huawei_cloud_token@%s server: not found", config.Version), nil))
	})

	router.Run(config.Cfg.Port)
}

// func GetNowTime(c *gin.Context) {
// 	dateTime := time.Now()
// 	c.JSON(200, utils.SendResult(200, fmt.Sprintf("nowTime: %v", dateTime), nil))
// }
