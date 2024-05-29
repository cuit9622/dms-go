package util

import (
	"fmt"

	"github.com/cuit9622/dms/common/global"

	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func GetInstance(name string) *model.Instance {
	instance, err := global.GLO_NACOS.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
	})
	if err != nil {
		global.GLO_LOG.Panic(fmt.Sprintf("failed to get %s instance: %s", name, err.Error()))
	}
	return instance
}
