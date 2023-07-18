package facades

import (
	"github.com/mewway/go-laravel/contracts/grpc"
)

func Grpc() grpc.Grpc {
	return App().MakeGrpc()
}
