package usecase

import (
	cons "cc-auth/controllers/models"
	postgre "cc-auth/databases/postgresql"
	trans "cc-auth/databases/postgresql/models"
	redis_db "cc-auth/databases/redis"
	grpcClient "cc-auth/grpc/client"
	host "cc-auth/hosts"
	cb "cc-auth/hosts/callback/models"
	tModels "cc-auth/hosts/transaction/models"
	"cc-auth/models"
)

type (
	usecase struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
		host	host.HostInterface
		hostGrpc grpcClient.GrpcInterface
	}
	UsecaseInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
		InsertDB(req models.ItemList) error
		QueryDB() ([]models.ItemList, error)
		Register(req cons.Credentials) error
		Login (req cons.Credentials)(string,error)
		AddCC(req cons.ReqCC)error
		TopUpCC(req cons.ReqCC)error
		GetCC()([]trans.CreditCards,error)
		DelCC(id string)error
		InquiryItems()([]cb.InquiryItems,error)
		InquiryDiscounts()([]cb.InquiryDiscounts,error)
		TransItem(req tModels.TransactionItems)(tModels.DecTransItem,error)
	}
)

func InitUsecase(postgre postgre.PostgreInterface, redis redis_db.RedisInterface, host host.HostInterface, hostGrpc grpcClient.GrpcInterface) UsecaseInterface {
	return &usecase{
		postgre: postgre,
		redis:   redis,
		host: host,
		hostGrpc:hostGrpc,
	}
}
