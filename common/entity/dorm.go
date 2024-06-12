package entity

//go:generate gomodifytags -all -add-tags form,json -add-options json=omitempty -transform camelcase --skip-unexported -w -file $GOFILE
type Dorm struct {
	ID             int64     `gorm:"primaryKey;comment:寝室主键" form:"id" json:"id,omitempty"`
	Name           string    `gorm:"not null;comment:寝室名字" form:"name" json:"name,omitempty"`
	Size           int32     `gorm:"not null;comment:寝室人数" form:"size" json:"size,omitempty"`
	Floor          int32     `gorm:"not null;comment:寝室所在楼层" form:"floor" json:"floor,omitempty"`
	DormBuildingID int64     `gorm:"not null;comment:寝室楼ID" form:"dormBuildingID" json:"dormBuildingID,omitempty"`
	DormBeds       []DormBed `form:"dormBeds" json:"dormBeds,omitempty"`
}
type DormBed struct {
	ID        int64 `gorm:"primaryKey;comment:床ID" form:"id" json:"id,omitempty"`
	StudentID int64 `gorm:"not null;comment:床绑定的学生ID" form:"studentID" json:"studentID,omitempty"`
	DormID    int64 `gorm:"not null;comment:寝室ID" form:"dormID" json:"dormID,omitempty"`
}
type DormBuilding struct {
	ID    int64  `gorm:"primaryKey;comment:寝室楼ID" form:"id" json:"id,omitempty"`
	Name  string `gorm:"not null;comment:寝室名字" form:"name" json:"name,omitempty"`
	Sex   int8   `gorm:"not null;comment:寝室性别 0-女,1-男" form:"sex" json:"sex"`
	Floor int32  `gorm:"not null;comment:寝室楼层" form:"floor" json:"floor,omitempty"`
}
