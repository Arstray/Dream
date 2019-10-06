package dream

import (
	"dream/buff"
	"dream/building"
	"dream/building/business"
	"dream/building/industry"
	"dream/building/residence"
	"testing"
)

func TestNewGroupEmpty(t *testing.T) {
	g := NewGroup([9]building.Building{
		residence.NewWoodenHouse(5),
		residence.NewGardenHouse(3),
		residence.NewStudioApartment(5),

		business.NewFolkFood(3),
		business.NewBusinessCentre(4),
		business.NewMediaVoice(2),

		industry.NewWoodFactory(5),
		industry.NewPenguinMachinery(2),
		industry.NewPetroleum(2),
	})
	g.SetProfit(buff.NewPolicy(90, 100, 100, nil))

	g.Profit()
}
func TestNewGroup(t *testing.T) {
	g := NewGroup([9]building.Building{

		industry.NewIronworks(5),
		industry.NewPowerPlant(5),
		industry.NewTextileMill(5),

		residence.NewSteelStructure(5),
		residence.NewTalentApartment(4),
		residence.NewChineseSmallHouse(4),

		business.NewFolkFood(3),
		business.NewCouture(5),
		business.NewBookCity(4),
	})
	// 设置已经完成的相册(预设)
	//江水人家套
	g.SetAlbum(buff.RiversAndLakes)
	//东三省套
	g.SetAlbum(buff.ThreeEastern)
	//中国套
	g.SetAlbum(buff.Chinese)
	// 设置已经完成的相册(单个)

	// 设置完成的政策组(预设)
	// 阶段一
	g.SetProfit(buff.Stage1)
	// 阶段二
	g.SetProfit(buff.Stage2)
	// 设置已经完成的单个政策
	g.SetProfit(buff.NewPolicy(300, 0, 0,
		buff.NewClassBuff(building.Business, 1200).Add(building.Industry, 1200),
	))
	// 设置家国之光，家园之光活动buff(属于政策)
	g.SetProfit(buff.NewPolicy(95, 100, 100, nil))
	// 设置正在进行的城市任务
	g.SetCity(buff.NewCity(0, 0, 0, nil, buff.NewFetterBuff(building.School, 200).Add(building.WoodenHouse, 200)))
	// 显示目前搭配倍率
	g.Profit()
}
