package huawei_cloud_controller

import (
	"bytes"
	"fmt"
	"go-huawei_cloud_token_manager/config"
	"net/http"
	"time"
)

func GetNewToken() (int, string) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	body := fmt.Sprintf(`
	{
		"auth": {
			"identity": {
				"methods": [
					"password"
				],
				"password": {
					"user": {
						"domain": {
							"name": "%s"
						},
						"name": "%s",
						"password": "%s"
					}
				}
			},
			"scope": {
				"project": {
					"name": "cn-southwest-2"
				}
			}
		}
	}
	`, config.Cfg.Huawei.User, config.Cfg.Huawei.IAMUser, config.Cfg.Huawei.IAMPass)
	rqst, err := client.Post("https://iam.cn-southwest-2.myhuaweicloud.com/v3/auth/tokens", "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil || rqst.StatusCode != 201 {
		fmt.Println("New request failed:", err)
		fmt.Println(rqst)
		return -1, ""
	}

	return 0, rqst.Header["X-Subject-Token"][0]
}
