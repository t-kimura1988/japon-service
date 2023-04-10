package group

import (
	"context"
	japonService "github.com/t-kimura1988/japon-proto/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"jpon-service/src/types"
	"log"
)

type groupCreateServer types.JaponServer

func (s *groupCreateServer) CreateGroup(context.Context, *japonService.CreateGroupRequest)  (*emptypb.Empty, error){
	log.Println("create group!!!!")
	return nil, nil
}