package grpc_client

import (
	grpc_callback "cc-auth/grpc/client/callback"
	grpc_transaction "cc-auth/grpc/client/transaction"
)

type (
	host struct{
		callback	grpc_callback.CallbackInterface
		transaction	grpc_transaction.TransactionInterface
	}
	GrpcInterface interface{
		Callback()grpc_callback.CallbackInterface
		Transaction()grpc_transaction.TransactionInterface
	}
)

func InitGrpcClient(callback grpc_callback.CallbackInterface, transaction grpc_transaction.TransactionInterface) GrpcInterface {
	
	return &host{
		callback: callback,
		transaction: transaction,
	}
}

func (g *host) Callback()grpc_callback.CallbackInterface{
	return g.callback
}
func (g *host) Transaction()grpc_transaction.TransactionInterface{
	return g.transaction
}