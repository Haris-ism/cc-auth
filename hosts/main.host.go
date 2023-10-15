package host

import (
	"cc-auth/hosts/callback"
	"cc-auth/hosts/transaction"
)

type (
	host struct{
		callback	callback.CallbackInterface
		transaction	transaction.TransactionInterface
	}
	HostInterface interface{
		Callback()callback.CallbackInterface
		Transaction()transaction.TransactionInterface
	}
)

func InitHost(callback callback.CallbackInterface,transaction transaction.TransactionInterface) HostInterface {
	return &host{
		callback: callback,
		transaction: transaction,
	}
}

func (h *host) Callback()callback.CallbackInterface{
	return h.callback
}
func (h *host) Transaction()transaction.TransactionInterface{
	return h.transaction
}