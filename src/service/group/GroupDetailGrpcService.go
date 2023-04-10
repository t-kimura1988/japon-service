package group

import (
	"context"
	JaponService "github.com/t-kimura1988/japon-proto/pkg/grpc"
	"jpon-service/src/types"
	"log"
)

type JaponServer types.JaponServer

func NewGroupServer() *JaponServer {
	return &JaponServer{}
}

func (s *JaponServer) GetGroupDetail(context.Context, *JaponService.GetGroupDetailRequest)  (*JaponService.GetGroupDetailResponse, error){
	log.Println("detail group!!!!")
	return &JaponService.GetGroupDetailResponse{}, nil
}