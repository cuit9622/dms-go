package service

import (
	"context"
	"github.com/cuit9622/dms/common/entity"
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/common/util/gormUtil"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DormService struct {
	pb.UnimplementedDormServiceServer
}

func (d DormService) Get(_ context.Context, request *pb.DormGetRequest) (*pb.PageResult, error) {
	result := pb.Dorms{}
	query := entity.Dorm{
		DormBuildingID: request.DormBuildingID,
		Floor:          request.Floor,
	}
	var total int64
	global.GLO_DB.Scopes(gormUtil.Paginate(request.Page.Page, request.Page.PageSize)).
		Model(&entity.Dorm{}).
		Find(&result.Dorms, &query)
	global.GLO_DB.Model(&entity.Dorm{}).Find(nil, &query).Count(&total)

	for _, item := range result.Dorms {
		global.GLO_DB.Model(&entity.DormBed{}).Select("id", "student_id").Find(&item.DormBeds, map[string]interface{}{"dorm_id": item.Id})
	}

	r, err := anypb.New(&result)
	if err != nil {
		return nil, err
	}
	return &pb.PageResult{Total: total, Records: r}, nil
}

func (d DormService) Update(_ context.Context, dorm *pb.Dorm) (*wrapperspb.Int32Value, error) {
	r := global.GLO_DB.Save(&dorm)
	return &wrapperspb.Int32Value{Value: int32(r.RowsAffected)}, nil
}

func (d DormService) Delete(_ context.Context, value *wrapperspb.Int64Value) (*wrapperspb.Int32Value, error) {
	r := global.GLO_DB.Delete(&entity.Dorm{}, value.Value)
	return &wrapperspb.Int32Value{Value: int32(r.RowsAffected)}, nil
}
