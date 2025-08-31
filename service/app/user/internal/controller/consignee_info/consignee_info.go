package consignee_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/user/api/consignee_info/v1"
	"service/app/user/api/pbentity"
	"service/app/user/internal/consts"
	"service/app/user/internal/dao"
	"service/app/user/internal/model/entity"
	"service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedConsigneeInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterConsigneeInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.ConsigneeInfoGetListResponse{
		List:  make([]*pbentity.ConsigneeInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.GetListFail)

	// 查询总数
	total, err := dao.ConsigneeInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	consigneeRecords, err := dao.ConsigneeInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		return &v1.ConsigneeInfoGetListRes{Data: response}, nil
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range consigneeRecords {
		var consignee entity.ConsigneeInfo
		if err := record.Struct(&consignee); err != nil {
			continue
		}

		var pbConsignee pbentity.ConsigneeInfo
		if err := gconv.Struct(consignee, &pbConsignee); err != nil {
			continue
		}

		pbConsignee.CreatedAt = utility.SafeConvertTime(consignee.CreatedAt)
		pbConsignee.UpdatedAt = utility.SafeConvertTime(consignee.UpdatedAt)
		pbConsignee.DeletedAt = utility.SafeConvertTime(consignee.DeletedAt)

		response.List = append(response.List, &pbConsignee)
	}

	return &v1.ConsigneeInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.CreateFail)

	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.ConsigneeInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.ConsigneeInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.UpdateFail)

	// 根据ID更新数据库中的信息
	_, err = dao.ConsigneeInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.ConsigneeInfoUpdateRes{Id: uint32(req.Id)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.ConsigneeInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.ConsigneeInfoDeleteRes{}, nil // 返回空结构体
}
