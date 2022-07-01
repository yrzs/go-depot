package model

import "github.com/jinzhu/gorm"

type ApiAccessToken struct {
	ID           int    `gorm:"primary_key" json:"id"` //
	MerchantId   int    `json:"merchant_id"`           //商户id
	RefreshToken string `json:"refresh_token"`         //刷新令牌
	AccessToken  string `json:"access_token"`          //授权令牌
	MemberId     int    `json:"member_id"`             //用户id
	Openid       string `json:"openid"`                //授权对象openid
	Group        string `json:"group"`                 //组别
	Status       int    `json:"status"`                //状态[-1:删除;0:禁用;1启用]
	CreatedAt    int    `json:"created_at"`            //创建时间
	UpdatedAt    int    `json:"updated_at"`            //修改时间
}

func (a ApiAccessToken) TableName() string {
	return "api_access_token"
}

/**
get info by access_token
*/
func (a ApiAccessToken) GetByAccessToken(db *gorm.DB, at string) (*ApiAccessToken, error) {
	var ac = &ApiAccessToken{}
	err := db.Where("access_token = ? and status = 1", at).First(&ac).Error
	return ac, err
}
