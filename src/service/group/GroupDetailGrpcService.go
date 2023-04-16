package group

import (
	"context"
	"fmt"
	domainErr "github.com/t-kimura1988/japon-domain/src/exceptions"
	JaponService "github.com/t-kimura1988/japon-proto/pkg/grpc/messages/group"
	serviceErr "github.com/t-kimura1988/japon-service/src/exceptions"
	"github.com/t-kimura1988/japon-service/src/types"

	"github.com/t-kimura1988/japon-domain/src/domain/group"
)

type JaponServer types.JaponServer

func NewGroupServer() *JaponServer {
	return &JaponServer{}
}

func (s *JaponServer) GetGroupDetail(c context.Context, req *JaponService.GetGroupDetailRequest) (*JaponService.GetGroupDetailResponse, error) {
	res, err := group.GroupDetailDomain(group.GroupDetailInput{GroupId: req.GetGroupId()})
	if me, ok := err.(*domainErr.DataNotFound); ok {
		return nil, &serviceErr.ServiceError{
			Code:       me.Code,
			Message:    me.Message,
			LogMessage: fmt.Sprintf("[Service: GetGroupDetail], %s", me.Message),
		}
	}
	return &JaponService.GetGroupDetailResponse{
		GroupId:     res.GroupID,
		GroupName:   res.GroupName,
		GroupDetail: res.GroupDetail,
	}, nil
}
