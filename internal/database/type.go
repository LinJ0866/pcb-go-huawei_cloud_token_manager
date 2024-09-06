package database

type huaweiTokens struct {
	Id         int    `gorm:"column:id;AUTO_INCREMENT"`
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}
