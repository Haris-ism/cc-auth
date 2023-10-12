package host

import (
	"cc-auth/hosts/callback"
)

type (
	host struct{
		callback	callback.CallbackInterface
	}
	HostInterface interface{
		Callback()callback.CallbackInterface
	}
)

func InitHost(callback callback.CallbackInterface) HostInterface {
	return &host{
		callback: callback,
	}
}

func (h *host) Callback()callback.CallbackInterface{
	return h.callback
}