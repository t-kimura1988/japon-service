package group

import (
	"context"
	japonService "github.com/t-kimura1988/japon-proto/pkg/grpc"
	"github.com/t-kimura1988/japon-service/src/types"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type groupCreateServer types.JaponServer

func (s *groupCreateServer) CreateGroup(context.Context, *japonService.CreateGroupRequest) (*emptypb.Empty, error) {
	log.Println("create group!!!!")
	return nil, nil
}
