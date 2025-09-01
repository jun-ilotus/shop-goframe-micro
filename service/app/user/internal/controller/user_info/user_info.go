package user_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "service/app/user/api/user_info/v1"
	"service/app/user/internal/consts"
	"service/app/user/internal/logic/user_info"
	"service/app/user/internal/model/entity"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedUserInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserInfoServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	// 调用 login 层
	token, expireIn, userInfo, err := user_info.Login(ctx, req.Name, req.Password)
	infoError := consts.InfoError(consts.UserInfo, consts.LoginFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireIn) * time.Second)
	expireProto := timestamppb.New(expireTime)
	if err := expireProto.CheckValid(); err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.UserInfoLoginRes{
		Type:     "Bearer",
		Token:    token,
		ExpireIn: uint32(expireIn),
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}

func (*Controller) Register(ctx context.Context, req *v1.UserInfoRegisterReq) (res *v1.UserInfoRegisterRes, err error) {
	var registerData *entity.UserInfo

	// 将请求参数 req 转换为实体对象 consigneeInfo
	if err := gconv.Struct(req, &registerData); err != nil {
		return nil, err
	}

	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.RegisterFail)
	// 调用logic 层注册
	userInfo, err := user_info.Register(ctx, registerData)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoRegisterRes{
		Id: uint32(userInfo.Id),
	}, nil
}

func (*Controller) UpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (*v1.UserInfoUpdatePasswordRes, error) {
	// 调用logic层修改密码
	err := user_info.UpdatePassword(ctx, int(req.Id), req.Password, req.SecretAnswer)
	infoError := consts.InfoError(consts.UserInfo, consts.UpdatePasswordFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.UserInfoUpdatePasswordRes{
		Id: req.Id,
	}, nil
}

func (*Controller) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	userInfo, err := user_info.GetUserInfo(ctx, int(req.Id))
	infoError := consts.InfoError(consts.UserInfo, consts.GetUserInfoFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.UserInfoRes{
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}
