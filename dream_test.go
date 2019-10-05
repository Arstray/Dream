package dream

import (
	"dream/buff"
	"dream/building"
	"testing"
)

func TestNewGroupEmpty(t *testing.T) {
	g := NewGroup([9]building.Building{
		building.NewWoodenHouse(5),
		building.NewWoodFactory(5),
		building.NewGardenHouse(3),
		building.NewStudioApartment(5),
		building.NewFolkFood(3),
		building.NewBusinessCentre(4),
		building.NewMediaVoice(2),
		building.NewPenguinMachinery(2),
		building.NewPetroleum(2),
	})
	g.SetProfit(buff.NewPolicy(90, 100, 100, nil))

	g.Profit()
}
func TestNewGroup(t *testing.T) {
	g := NewGroup([9]building.Building{
		building.NewWoodenHouse(5),
		building.NewWoodFactory(5),
		building.NewGardenHouse(3),
		building.NewStudioApartment(5),
		building.NewFolkFood(3),
		building.NewBusinessCentre(4),
		building.NewMediaVoice(2),
		building.NewPenguinMachinery(2),
		building.NewPetroleum(2),
	})
	// 设置已经完成的相册(预设)
	g.SetAlbum(buff.RiversAndLakes)
	g.SetAlbum(buff.ThreeEastern)
	g.SetAlbum(buff.OneBeltOneRoad)
	g.SetAlbum(buff.SportsPower)
	g.SetAlbum(buff.MilitaryPower)
	g.SetAlbum(buff.OneCountryTwoSystems)
	// 设置已经完成的相册(单个)

	// 设置完成的政策组(预设)
	g.SetProfit(buff.Stage1)
	g.SetProfit(buff.Stage2)
	// 设置已经完成的单个政策
	g.SetProfit(buff.NewPolicy(300, 0, 0,
		buff.NewClassBuff(building.Business, 900).Add(building.Industry, 900),
	))
	// 设置家国之光，家园之光活动buff(属于政策)
	g.SetProfit(buff.NewPolicy(90, 100, 100, nil))
	// 设置正在进行的城市任务

	// 显示目前搭配倍率
	g.Profit()
}
