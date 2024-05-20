package entity

type Dorm struct {
	Id             int64  `gorm:"primaryKey;comment:寝室主键"`
	Name           string `gorm:"not null;comment:寝室名字"`
	Size           int    `gorm:"not null;comment:寝室人数"`
	Floor          int    `gorm:"not null;comment:寝室所在楼层"`
	DormBuildingId int64  `gorm:"not null;comment:寝室楼ID"`
	DormBeds       []DormBed
}
type DormBed struct {
	Id        int64 `gorm:"primaryKey;comment:床ID"`
	StudentId int64 `gorm:"not null;comment:床绑定的学生ID"`
	DormId    int64 `gorm:"not null;comment:寝室ID"`
}
type DormBuilding struct {
	Id    int64  `gorm:"primaryKey;comment:寝室楼ID"`
	Name  string `gorm:"not null;comment:寝室名字"`
	Sex   int8   `gorm:"not null;comment:寝室性别 0-女,1-男"`
	Dorms []Dorm
}
