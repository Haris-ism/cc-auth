package grpc_client

import (
	grpc_callback "cc-auth/grpc/client/callback"
)

type (
	host struct{
		callback	grpc_callback.CallbackInterface
		// transaction	transaction.TransactionInterface
		// conn		*grpc.ClientConn
	}
	GrpcInterface interface{
		Callback()grpc_callback.CallbackInterface
		// Transaction()transaction.TransactionInterface
	}
)

func InitGrpcClient(callback grpc_callback.CallbackInterface) GrpcInterface {
	
	return &host{
		callback: callback,
		// transaction: transaction,
		// conn:conn,
	}
}

func (g *host) Callback()grpc_callback.CallbackInterface{
	return g.callback
}
// func (g *host) Transaction()transaction.TransactionInterface{
// 	return g.transaction
// }