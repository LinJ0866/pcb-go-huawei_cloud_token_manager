package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/clause"
)

func GetAvailableToken() (int, huaweiTokens) {
	var tokenInfo huaweiTokens
	result := instance.Select("token", "expire_time").Where("expire_time > ?", fmt.Sprintf("%d", time.Now().Unix())).First(&tokenInfo)
	if result.Error != nil {
		log.Println("err:", result.Error)
		return -1, tokenInfo
	}
	return 0, tokenInfo
}

func UpdateToken(token string) (int, huaweiTokens) {
	tokenInfo := huaweiTokens{
		Id:         1,
		Token:      token,
		ExpireTime: fmt.Sprintf("%d", time.Now().AddDate(0, 0, 1).Unix()),
	}
	err := instance.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&tokenInfo).Error
	if err != nil {
		log.Println("update token to mysql err: ", err)
		return -1, tokenInfo
	}

	return 0, tokenInfo
}
