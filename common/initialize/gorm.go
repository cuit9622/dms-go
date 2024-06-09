package initialize

import (
	"fmt"
	"github.com/cuit9622/dms/common/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() {
	dsn := fmt.Sprintf("root:root@tcp(%s)/dms?charset=utf8mb4&parseTime=True&loc=Local", global.GLO_VP.GetString("MYSQL_SERVER"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("error create gorm: %s", err))
	}
	global.GLO_DB = db
	// err = db.Migrator().CreateTable(&entity.DormBed{}, &entity.Dorm{}, &entity.DormBuilding{})
	global.GLO_LOG.Info("Gorm initialization complete")
}
