package repository

import (
	"context"
	handlerAuth_model "sekawan-backend/app/main/handlerAuth/model"
	baseModel "sekawan-backend/app/main/server"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	queryGetMerchantIdByUserIdChatat string = "select merchant_id  from merchant_users_chatat where user_id_chatat = ?"
	queryGetUsersByPhoneNumber       string = "select user_id, full_name, phone_no, email, acl from auth_users where phone_no_hash = ? and password_hash=? and is_active = true and expired_at >= now()"
	queryGetUsersByEmail             string = "select user_id, full_name, phone_no, email, acl from auth_users where email = ? and password_hash=? and is_active = true and expired_at >= now()"
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

// todo here
func (di databaseImpl) GetUserIdByPhoneNo(ctx context.Context, phoneNumbeMd5 string, passwordMd5 string) handlerAuth_model.AuthModel {
	var modelAuth handlerAuth_model.AuthModel
	result := di.db.WithContext(ctx).Raw(queryGetUsersByPhoneNumber, phoneNumbeMd5, passwordMd5).Scan(&modelAuth)
	if result.Error != nil {
		logrus.Errorln("error "+phoneNumbeMd5+" password "+passwordMd5, result.Error)
	}
	return modelAuth
}

// todo here
func (di databaseImpl) GetUserIdByEmail(ctx context.Context, phoneNumbeMd5 string, passwordMd5 string) handlerAuth_model.AuthModel {
	var modelAuth handlerAuth_model.AuthModel
	result := di.db.WithContext(ctx).Raw(queryGetUsersByEmail, phoneNumbeMd5, passwordMd5).Scan(&modelAuth)
	if result.Error != nil {
		logrus.Errorln("error", result.Error)
	}
	return modelAuth
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
