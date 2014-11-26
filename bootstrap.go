package main

import (
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/factory"
)

func bootstrap() {
	iface.NewUserRepo = factory.NewUserRepo
	iface.NewClientRepo = factory.NewClientRepo
	iface.NewPersonRepo = factory.NewPersonRepo
	iface.NewActivityRepo = factory.NewActivityRepo
	iface.NewBranchRepo = factory.NewBranchRepo
	iface.NewCityRepo = factory.NewCityRepo
	iface.NewDistrictRepo = factory.NewDistrictRepo
	iface.NewRegionRepo = factory.NewRegionRepo
}
