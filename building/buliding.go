package building

// Building 建筑实体
type Building struct {
	Name  string
	Base  float64
	Type  Type
	Class Class
	Star  Star
}

// Class 建筑类别(商业,住宅,工业)
type Class int

// Type 建筑类型
type Type int

// Star 建筑星级
type Star int

func (s Star) Value() int {
	return int(s)
}

const (
	// Residence 住宅
	Residence Class = iota
	// Business 商业
	Business
	// Industry 工业
	Industry
)

const (
	//木屋
	WoodenHouse Type = iota

	//钢结构房
	SteelStructure

	//平房
	Bungalow

	//小型公寓
	StudioApartment

	// 居民楼
	ResidentialBuilding

	// 人才公寓
	TalentApartment

	// 花园洋房
	GardenHouse

	// 中式小楼
	ChineseSmallHouse

	// 空中别墅
	Villa

	// 复兴公馆
	RenaissanceMansion

	// 便利店
	ConvenienceStore

	// 学校
	School

	// 服装店
	Couture

	// 五金店
	HardwareStore

	// 菜市场
	FoodMarket

	// 图书城
	BookCity

	// 商贸中心
	BusinessCentre

	// 加油站
	GasStation

	// 民食斋
	FolkFood

	// 媒体之声
	MediaVoice

	// 木材厂
	WoodFactory

	// 造纸厂
	PaperMill

	// 水厂
	Waterworks

	// 电厂
	PowerPlant

	// 食品厂
	FoodManufacturer

	// 钢铁厂
	Ironworks

	// 纺织厂
	TextileMill

	// 零件厂
	SparePartsFactory

	// 企鹅机械
	PenguinMachinery

	// 人民石油
	Petroleum
)

// Residence
func NewWoodenHouse(star Star) Building {
	return Building{
		Name:  "木质小屋",
		Base:  1,
		Class: Residence,
		Type:  WoodenHouse,
		Star:  star,
	}
}
func NewSteelStructure(star Star) Building {
	return Building{
		Name:  "钢结构房",
		Base:  1,
		Class: Residence,
		Type:  SteelStructure,
		Star:  star,
	}
}
func NewBungalow(star Star) Building {
	return Building{
		Name:  "平凡住房",
		Base:  1.1,
		Class: Residence,
		Type:  Bungalow,
		Star:  star,
	}
}
func NewStudioApartment(star Star) Building {
	return Building{
		Name:  "小型公寓",
		Base:  1.18,
		Class: Residence,
		Type:  StudioApartment,
		Star:  star,
	}
}
func NewResidentialBuilding(star Star) Building {
	return Building{
		Name:  "居民大楼",
		Base:  1,
		Class: Residence,
		Type:  ResidentialBuilding,
		Star:  star,
	}
}
func NewTalentApartment(star Star) Building {
	return Building{
		Name:  "人才公寓",
		Base:  1.4,
		Class: Residence,
		Type:  TalentApartment,
		Star:  star,
	}
}
func NewGardenHouse(star Star) Building {
	return Building{
		Name:  "花园洋房",
		Base:  1.022,
		Class: Residence,
		Type:  GardenHouse,
		Star:  star,
	}
}
func NewChineseSmallHouse(star Star) Building {
	return Building{
		Name:  "中式小楼",
		Base:  1.4,
		Class: Residence,
		Type:  ChineseSmallHouse,
		Star:  star,
	}
}
func NewVilla(star Star) Building {
	return Building{
		Name:  "空中别墅",
		Base:  1.52,
		Class: Residence,
		Type:  Villa,
		Star:  star,
	}
}
func NewRenaissanceMansion(star Star) Building {
	return Building{
		Name:  "复兴公馆",
		Base:  1.672,
		Class: Residence,
		Type:  RenaissanceMansion,
		Star:  star,
	}
}

// Business
func NewConvenienceStore(star Star) Building {
	return Building{
		Name:  "便利之店",
		Base:  1,
		Class: Business,
		Type:  ConvenienceStore,
		Star:  star,
	}
}
func NewSchool(star Star) Building {
	return Building{
		Name:  "清华大学",
		Base:  1,
		Class: Business,
		Type:  School,
		Star:  star,
	}
}
func NewCouture(star Star) Building {
	return Building{
		Name:  "服装之店",
		Base:  1,
		Class: Business,
		Type:  Couture,
		Star:  star,
	}
}
func NewHardwareStore(star Star) Building {
	return Building{
		Name:  "五金之店",
		Base:  1,
		Class: Business,
		Type:  HardwareStore,
		Star:  star,
	}
}
func NewFoodMarket(star Star) Building {
	return Building{
		Name:  "菜市之场",
		Base:  1,
		Class: Business,
		Type:  FoodMarket,
		Star:  star,
	}
}
func NewBookCity(star Star) Building {
	return Building{
		Name:  "图书之城",
		Base:  1,
		Class: Business,
		Type:  BookCity,
		Star:  star,
	}
}
func NewBusinessCentre(star Star) Building {
	return Building{
		Name:  "商贸中心",
		Base:  1.022,
		Class: Business,
		Type:  BusinessCentre,
		Star:  star,
	}
}
func NewGasStation(star Star) Building {
	return Building{
		Name:  "加油之站",
		Base:  1.204,
		Class: Business,
		Type:  GasStation,
		Star:  star,
	}
}
func NewFolkFood(star Star) Building {
	return Building{
		Name:  "民食之斋",
		Base:  1.52,
		Class: Business,
		Type:  FolkFood,
		Star:  star,
	}
}
func NewMediaVoice(star Star) Building {
	return Building{
		Name:  "媒体之声",
		Base:  1.615,
		Class: Business,
		Type:  MediaVoice,
		Star:  star,
	}
}

// Industry
func NewWoodFactory(star Star) Building {
	return Building{
		Name:  "木材工厂",
		Base:  1,
		Class: Industry,
		Type:  WoodFactory,
		Star:  star,
	}
}
func NewPaperMill(star Star) Building {
	return Building{
		Name:  "造纸工厂",
		Base:  1,
		Class: Industry,
		Type:  PaperMill,
		Star:  star,
	}
}
func NewWaterworks(star Star) Building {
	return Building{
		Name:  "水力工厂",
		Base:  1.26,
		Class: Industry,
		Type:  Waterworks,
		Star:  star,
	}
}
func NewPowerPlant(star Star) Building {
	return Building{
		Name:  "电力工厂",
		Base:  1.18,
		Class: Industry,
		Type:  PowerPlant,
		Star:  star,
	}
}
func NewFoodManufacturer(star Star) Building {
	return Building{
		Name:  "食品工厂",
		Base:  1,
		Class: Industry,
		Type:  FoodManufacturer,
		Star:  star,
	}
}
func NewIronworks(star Star) Building {
	return Building{
		Name:  "钢铁工厂",
		Base:  1,
		Class: Industry,
		Type:  Ironworks,
		Star:  star,
	}
}
func NewTextileMill(star Star) Building {
	return Building{
		Name:  "纺织工厂",
		Base:  1,
		Class: Industry,
		Type:  TextileMill,
		Star:  star,
	}
}
func NewSparePartsFactory(star Star) Building {
	return Building{
		Name:  "零件工厂",
		Base:  1,
		Class: Industry,
		Type:  SparePartsFactory,
		Star:  star,
	}
}
func NewPenguinMachinery(star Star) Building {
	return Building{
		Name:  "企鹅机械",
		Base:  1.33,
		Class: Industry,
		Type:  PenguinMachinery,
		Star:  star,
	}
}
func NewPetroleum(star Star) Building {
	return Building{
		Name:  "人民石油",
		Base:  1,
		Class: Industry,
		Type:  Petroleum,
		Star:  star,
	}
}
