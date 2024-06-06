package service

import (
	"context"
	"github.com/cuit9622/dms/common/entity"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/common/util/gormUtil"
)

type DormBuildingService struct {
	pb.UnimplementedDormBuildingServiceServer
}

func (i DormBuildingService) Create(ctx context.Context, in *pb.DormBuilding) (*wrapperspb.Int32Value, error) {
	result := global.GLO_DB.Create(&in)
	return &wrapperspb.Int32Value{Value: int32(result.RowsAffected)}, nil
}

func (i DormBuildingService) Get(ctx context.Context, p *pb.PageRequest) (*pb.PageResult, error) {
	result := pb.DormBuildings{}
	global.GLO_DB.Scopes(gormUtil.Paginate(p.Page, p.PageSize)).Model(&entity.DormBuilding{}).Find(&result.DormBuildings)
	r, err := anypb.New(&result)
	if err != nil {
		return nil, err
	}
	return &pb.PageResult{Total: 0, Data: r}, nil
}

func (i DormBuildingService) Update(ctx context.Context, building *pb.DormBuilding) (*wrapperspb.Int32Value, error) {
	r := global.GLO_DB.Save(&building)
	return &wrapperspb.Int32Value{Value: int32(r.RowsAffected)}, nil
}

func (i DormBuildingService) Delete(ctx context.Context, in *wrapperspb.Int64Value) (*wrapperspb.Int32Value, error) {
	r := global.GLO_DB.Delete(&entity.DormBuilding{}, in.Value)
	return &wrapperspb.Int32Value{Value: int32(r.RowsAffected)}, nil
}
