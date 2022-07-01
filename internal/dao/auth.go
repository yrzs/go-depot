package dao

import "go-depot/internal/model"

func (d Dao) GetApiAccessTokenInfoByAccessToken(at string) (*model.ApiAccessToken, error) {
	ac := model.ApiAccessToken{}
	return ac.GetByAccessToken(d.db, at)
}
