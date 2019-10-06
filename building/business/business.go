package business

import "dream/building"

// Business
func NewConvenienceStore(star building.Star) building.Building {
	return building.Building{
		Name:   "便利之店",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Business,
		Type:   building.ConvenienceStore,
		Star:   star,
	}
}
func NewSchool(star building.Star) building.Building {
	return building.Building{
		Name:   "清华大学",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Business,
		Type:   building.School,
		Star:   star,
	}
}
func NewCouture(star building.Star) building.Building {
	return building.Building{
		Name:   "服装之店",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Business,
		Type:   building.Couture,
		Star:   star,
	}
}
func NewHardwareStore(star building.Star) building.Building {
	return building.Building{
		Name:   "五金之店",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Business,
		Type:   building.HardwareStore,
		Star:   star,
	}
}
func NewFoodMarket(star building.Star) building.Building {
	return building.Building{
		Name:   "菜市之场",
		Base:   1,
		Rarity: building.Common,
		Class:  building.Business,
		Type:   building.FoodMarket,
		Star:   star,
	}
}
func NewBookCity(star building.Star) building.Building {
	return building.Building{
		Name:   "图书之城",
		Base:   1,
		Rarity: building.Rare,
		Class:  building.Business,
		Type:   building.BookCity,
		Star:   star,
	}
}
func NewBusinessCentre(star building.Star) building.Building {
	return building.Building{
		Name:   "商贸中心",
		Base:   1.022,
		Rarity: building.Rare,
		Class:  building.Business,
		Type:   building.BusinessCentre,
		Star:   star,
	}
}
func NewGasStation(star building.Star) building.Building {
	return building.Building{
		Name:   "加油之站",
		Base:   1.204,
		Rarity: building.Rare,
		Class:  building.Business,
		Type:   building.GasStation,
		Star:   star,
	}
}
func NewFolkFood(star building.Star) building.Building {
	return building.Building{
		Name:   "民食之斋",
		Base:   1.52,
		Rarity: building.Epic,
		Class:  building.Business,
		Type:   building.FolkFood,
		Star:   star,
	}
}
func NewMediaVoice(star building.Star) building.Building {
	return building.Building{
		Name:   "媒体之声",
		Base:   1.615,
		Rarity: building.Epic,
		Class:  building.Business,
		Type:   building.MediaVoice,
		Star:   star,
	}
}
