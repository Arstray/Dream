package industry

import "dream/building"

// Industry
func NewWoodFactory(star building.Star) building.Building {
	return building.Building{
		Name:   "木材工厂",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Industry,
		Type:   building.WoodFactory,
		Star:   star,
	}
}
func NewPaperMill(star building.Star) building.Building {
	return building.Building{
		Name:   "造纸工厂",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Industry,
		Type:   building.PaperMill,
		Star:   star,
	}
}
func NewWaterworks(star building.Star) building.Building {
	return building.Building{
		Name:   "水力工厂",
		Base:   1.26,
		Rarity: building.Common,
		Class:  building.Industry,
		Type:   building.Waterworks,
		Star:   star,
	}
}
func NewPowerPlant(star building.Star) building.Building {
	return building.Building{
		Name:   "电力工厂",
		Base:   1.18,
		Rarity: building.Common,
		Class:  building.Industry,
		Type:   building.PowerPlant,
		Star:   star,
	}
}
func NewFoodManufacturer(star building.Star) building.Building {
	return building.Building{
		Name:   "食品工厂",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Industry,
		Type:   building.FoodManufacturer,
		Star:   star,
	}
}
func NewIronworks(star building.Star) building.Building {
	return building.Building{
		Name:   "钢铁工厂",
		Base:   1,
		Rarity: building.Rare,
		Class:  building.Industry,
		Type:   building.Ironworks,
		Star:   star,
	}
}
func NewTextileMill(star building.Star) building.Building {
	return building.Building{
		Name:   "纺织工厂",
		Base:   1,
		Rarity: building.Rare,
		Class:  building.Industry,
		Type:   building.TextileMill,
		Star:   star,
	}
}
func NewSparePartsFactory(star building.Star) building.Building {
	return building.Building{
		Name:   "零件工厂",
		Base:   1,
		Rarity: building.Rare,
		Class:  building.Industry,
		Type:   building.SparePartsFactory,
		Star:   star,
	}
}
func NewPenguinMachinery(star building.Star) building.Building {
	return building.Building{
		Name:   "企鹅机械",
		Base:   1.33,
		Rarity: building.Epic,
		Class:  building.Industry,
		Type:   building.PenguinMachinery,
		Star:   star,
	}
}
func NewPetroleum(star building.Star) building.Building {
	return building.Building{
		Name:   "人民石油",
		Base:   1,
		Rarity: building.Epic,
		Class:  building.Industry,
		Type:   building.Petroleum,
		Star:   star,
	}
}
