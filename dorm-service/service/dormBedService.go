package service

import (
	"context"
	"github.com/cuit9622/dms/common/entity"
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DormBedService struct {
	pb.UnimplementedDormBedServiceServer
}

func (d DormBedService) GetStudentCount(context.Context, *wrapperspb.Int64Value) (*pb.StudentCounts, error) {
	result := pb.StudentCounts{}
	global.GLO_DB.Raw(`select id,name,
(select count(*) from dorms 
inner join dorm_beds on dorms.id=dorm_beds.dorm_id
where dorms.dorm_building_id=dorm_buildings.id and dorm_beds.student_id!=0) 'count'
from dorm_buildings`).Scan(&result.StudentCounts)
	return &result, nil
}

func (d DormBedService) Update(_ context.Context, bed *pb.DormBed) (*wrapperspb.Int32Value, error) {
	dorm := entity.Dorm{}
	global.GLO_DB.Select("size").First(&dorm, bed.DormID)
	var count int64
	global.GLO_DB.Model(&entity.DormBed{}).Find(nil, &entity.DormBed{DormID: bed.DormID}).Count(&count)
	if int32(count) >= dorm.Size && bed.Id == 0 {
		return wrapperspb.Int32(-1), nil
	}
	r := global.GLO_DB.Save(bed)
	return wrapperspb.Int32(int32(r.RowsAffected)), nil
}
func (d DormBedService) Delete(_ context.Context, id *wrapperspb.Int64Value) (*wrapperspb.Int32Value, error) {
	r := global.GLO_DB.Delete(&entity.DormBed{}, id.Value)
	return &wrapperspb.Int32Value{Value: int32(r.RowsAffected)}, nil
}
