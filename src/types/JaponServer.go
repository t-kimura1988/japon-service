package types

import JaponService "github.com/t-kimura1988/japon-proto/pkg/grpc"

type JaponServer struct {
	JaponService.UnimplementedGroupServiceServer
}
