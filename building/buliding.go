package building

// Building 建筑实体
type Building struct {
	Name   string
	Base   float64
	Rarity Rarity
	Type   Type
	Class  Class
	Star   Star
}

// Class 建筑类别(商业,住宅,工业)
type Class int

// Type 建筑类型
type Type int

// Rarity 稀有度
type Rarity string

// Star 建筑星级
type Star int

func (s Star) Value() int {
	return int(s)
}

func (r Rarity) Value() string {
	return string(r)
}

const (
	// Common 普通
	Common Rarity = "普通"
	// Rare 稀有
	Rare Rarity = "稀有"
	// Epic 史诗
	Epic Rarity = "史诗"
)

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
