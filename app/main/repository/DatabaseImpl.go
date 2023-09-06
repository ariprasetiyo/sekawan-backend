package repository

import (
	"context"
	baseModel "sekawan-backend/app/main/server"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	queryGetMerchantIdByUserIdChatat string = "select merchant_id  from merchant_users_chatat where user_id_chatat = ?"
	queryInsertToken                 string = "insert into auth_token values (user_id, token, created_at, expired_at returning token )"
)

func NewDatabase(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Database {
	return &databaseImpl{db: db, base: base}
}

type databaseImpl struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func (di databaseImpl) GetCount(ctx context.Context, userId string) string {
	var merchantId string
	err := di.db.WithContext(ctx).Raw(queryGetMerchantIdByUserIdChatat, userId).Scan(&merchantId)
	if err != nil {
		logrus.Errorln("error", err.Error)
	}
	return merchantId
}

func (di databaseImpl) SaveToken(ctx context.Context, userId string, token string, created_at string, expired_at string) string {
	var getToken string
	error := di.db.WithContext(ctx).Raw(queryInsertToken, userId, token, created_at, expired_at).Scan(getToken)
	if error != nil {
		logrus.Errorln("error", error.Error)
	}
	return getToken
}

func (di databaseImpl) GetUser(ctx context.Context, userId string, password string) string {
	var merchantId string
	err := di.db.WithContext(ctx).Raw(queryGetMerchantIdByUserIdChatat, userId).Scan(&merchantId)
	if err != nil {
		logrus.Errorln("error", err.Error)
	}
	return merchantId
}
