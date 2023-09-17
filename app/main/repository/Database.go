package repository

import (
	"context"
	handlerAuth_model "sekawan-backend/app/main/handlerAuth/model"
)

type Database interface {
	GetCount(ctx context.Context, merchantId string) string
	GetUserIdByPhoneNo(ctx context.Context, phoneNumber string, password string) handlerAuth_model.AuthModel
	GetUserIdByEmail(ctx context.Context, phoneNumber string, password string) handlerAuth_model.AuthModel
	SaveToken(ctx context.Context, userId string, token string, created_at string, expired_at string) string
}
