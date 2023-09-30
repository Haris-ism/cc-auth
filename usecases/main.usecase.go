package usecase

import (
	auth "cc-auth/controllers/models"
	postgre "cc-auth/databases/postgresql"
	redis_db "cc-auth/databases/redis"
	"cc-auth/models"
)

type (
	usecase struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
	}
	UsecaseInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
		InsertDB(req models.ItemList) error
		QueryDB() ([]models.ItemList, error)
		Register(req auth.Credentials) error
	}
)

func InitUsecase(postgre postgre.PostgreInterface, redis redis_db.RedisInterface) UsecaseInterface {
	return &usecase{
		postgre: postgre,
		redis:   redis,
	}
}
