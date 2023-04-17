package group

import (
	"context"
	"fmt"
	"github.com/t-kimura1988/japon-domain/src/domain/group"
	domainErr "github.com/t-kimura1988/japon-domain/src/exceptions"
	JaponService "github.com/t-kimura1988/japon-proto/pkg/grpc/messages/group"
	"github.com/t-kimura1988/japon-service/src/types"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JaponServer types.JaponServer

func NewGroupServer() *JaponServer {
	return &JaponServer{}
}

func (s *JaponServer) GetGroupDetail(c context.Context, req *JaponService.GetGroupDetailRequest) (*JaponService.GetGroupDetailResponse, error) {
	res, err := group.GroupDetailDomain(group.GroupDetailInput{GroupId: req.GetGroupId()})
	if me, ok := err.(*domainErr.DataNotFound); ok {
		stat := status.New(codes.NotFound, me.Message)
		stat, _ = stat.WithDetails(&errdetails.ErrorInfo{
			Reason: fmt.Sprintf("[GetGroupDetail Error]: %s", me.Message),
		})
		return nil, stat.Err()
	}
	return &JaponService.GetGroupDetailResponse{
		GroupId:               res.GroupID,
		GroupName:             res.GroupName,
		GroupDetail:           res.GroupDetail,
		GroupCreateFamilyName: res.GroupCreateFamilyName,
		GroupCreateGivenName:  res.GroupCreateGivenName,
		UserId:                res.UserId,
		GroupJoinNum:          res.GroupJoinNum,
	}, nil
}
