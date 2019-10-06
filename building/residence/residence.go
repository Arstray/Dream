package residence

import "dream/building"

// Residence
func NewWoodenHouse(star building.Star) building.Building {
	return building.Building{
		Name:   "木质小屋",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Residence,
		Type:   building.WoodenHouse,
		Star:   star,
	}
}
func NewSteelStructure(star building.Star) building.Building {
	return building.Building{
		Name:   "钢结构房",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Residence,
		Type:   building.SteelStructure,
		Star:   star,
	}
}
func NewBungalow(star building.Star) building.Building {
	return building.Building{
		Name:   "平凡住房",
		Base:   1.1,
		Rarity: building.Common,
		Class:  building.Residence,
		Type:   building.Bungalow,
		Star:   star,
	}
}
func NewStudioApartment(star building.Star) building.Building {
	return building.Building{
		Name:   "小型公寓",
		Base:   1.18,
		Rarity: building.Common,
		Class:  building.Residence,
		Type:   building.StudioApartment,
		Star:   star,
	}
}
func NewResidentialBuilding(star building.Star) building.Building {
	return building.Building{
		Name:   "居民大楼",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Residence,
		Type:   building.ResidentialBuilding,
		Star:   star,
	}
}
func NewTalentApartment(star building.Star) building.Building {
	return building.Building{
		Name:   "人才公寓",
		Base:   1.4,
		Rarity: building.Rare,
		Class:  building.Residence,
		Type:   building.TalentApartment,
		Star:   star,
	}
}
func NewGardenHouse(star building.Star) building.Building {
	return building.Building{
		Name:   "花园洋房",
		Base:   1.022,
		Rarity: building.Rare,
		Class:  building.Residence,
		Type:   building.GardenHouse,
		Star:   star,
	}
}
func NewChineseSmallHouse(star building.Star) building.Building {
	return building.Building{
		Name:   "中式小楼",
		Base:   1.4,
		Rarity: building.Rare,
		Class:  building.Residence,
		Type:   building.ChineseSmallHouse,
		Star:   star,
	}
}
func NewVilla(star building.Star) building.Building {
	return building.Building{
		Name:   "空中别墅",
		Base:   1.52,
		Rarity: building.Epic,
		Class:  building.Residence,
		Type:   building.Villa,
		Star:   star,
	}
}
func NewRenaissanceMansion(star building.Star) building.Building {
	return building.Building{
		Name:   "复兴公馆",
		Base:   1.672,
		Rarity: building.Epic,
		Class:  building.Residence,
		Type:   building.RenaissanceMansion,
		Star:   star,
	}
}
