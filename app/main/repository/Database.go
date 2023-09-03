package repository

import "context"

type Database interface {
	GetCount(ctx context.Context, merchantId string) string
	SaveToken(ctx context.Context, userId string, token string, created_at string, expired_at string) string
}
