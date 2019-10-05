package buff

import "dream/building"

// Fetters 关联(羁绊)
// 关联预设信息来自游戏
// 关联加成是不单调增长的，但更换建筑会重新构建组，所以加成是重新累加计算的
type Fetters struct {
	globalBuff
	onlineBuff
	offlineBuff
	classBuff
	fetterBuff
}

func NewFetterBuffs() *Fetters {
	return &Fetters{
		globalBuff:  0,
		onlineBuff:  0,
		offlineBuff: 0,
		classBuff:   make(map[building.Class]int),
		fetterBuff:  make(map[building.Type]int),
	}
}
// OnBuff 在线关联加成
func (f *Fetters) OnBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += f.globalBuff.Value()

	res += f.onlineBuff.Value()

	res += f.classBuff.Value(b.Class)

	res += f.fetterBuff.Value(b.Type)

	return float64(res) / 100
}
// OffBuff 离线关联加成
func (f *Fetters) OffBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += f.globalBuff.Value()

	res += f.offlineBuff.Value()

	res += f.classBuff.Value(b.Class)

	res += f.fetterBuff.Value(b.Type)

	return float64(res) / 100
}

// Composition 加成叠加
func (f *Fetters) Composition(b building.Building) {
	newFetters := typeBuff[b.Type](b.Star.Value())
	f.globalBuff += newFetters.globalBuff
	f.onlineBuff += newFetters.onlineBuff
	f.offlineBuff += newFetters.offlineBuff
	for k, v := range newFetters.classBuff {
		f.classBuff[k] += v
	}
	for k, v := range newFetters.fetterBuff {
		f.fetterBuff[k] += v
	}
}

// typeBuff 与逻辑无关的固有配置
// 根据类型和星级确定关联加成
var typeBuff = map[building.Type]func(star int) Fetters{
	// 平房->住宅
	building.Bungalow: func(star int) Fetters {
		return Fetters{classBuff: NewClassBuff(building.Residence, star*20)}
	},
	// 水厂->离线
	building.Waterworks: func(star int) Fetters {
		return Fetters{offlineBuff: offlineBuff(star * 5)}
	},
	// 电厂->在线
	building.PowerPlant: func(star int) Fetters {
		return Fetters{onlineBuff: onlineBuff((star * 30) - 10)}
	},

	// 木屋->伐木场
	building.WoodenHouse: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.WoodFactory, star*100)}
	},
	// 伐木场->木屋
	building.WoodFactory: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.WoodenHouse, star*100)}
	},

	// 居民楼->便利店
	building.ResidentialBuilding: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.ConvenienceStore, star*100)}
	},
	// 便利店->居民楼
	building.ConvenienceStore: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.ResidentialBuilding, star*100)}
	},

	// 菜市场->食品厂
	building.FoodMarket: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.FoodManufacturer, star*100)}
	},
	// 食品厂->菜市场
	building.FoodManufacturer: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.FoodMarket, star*100)}
	},

	// 钢结构房->炼钢厂
	building.SteelStructure: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.Ironworks, star*100)}
	},
	// 炼钢厂->钢结构房
	// 炼钢厂->工业
	building.Ironworks: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.SteelStructure, star*100),
			classBuff:  NewClassBuff(building.Industry, star*15),
		}
	},

	// 服装店->纺织厂
	building.Couture: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.TextileMill, star*100)}
	},
	// 纺织厂->服装店
	// 纺织厂->商业
	building.TextileMill: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.Couture, star*100),
			classBuff:  NewClassBuff(building.Business, star*15),
		}
	},

	// 图书城->造纸厂
	// 图书城->学校
	building.BookCity: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.PaperMill, star*100).Add(building.School, star*100),
		}
	},
	// 造纸厂->图书城
	building.PaperMill: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.BookCity, star*100)}
	},
	// 学校->图书城
	building.School: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.BookCity, star*100)}
	},

	// 加油站->人民石油
	// 加油站->离线
	building.GasStation: func(star int) Fetters {
		return Fetters{
			fetterBuff:  NewFetterBuff(building.Petroleum, star*50),
			offlineBuff: offlineBuff(star * 10),
		}
	},
	// 人民石油->加油站
	// 人民石油->离线
	building.Petroleum: func(star int) Fetters {
		return Fetters{
			fetterBuff:  NewFetterBuff(building.GasStation, star*100),
			offlineBuff: offlineBuff(star * 10),
		}
	},

	// 零件厂->企鹅机械
	// 零件厂->五金店
	building.SparePartsFactory: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.PenguinMachinery, star*50).Add(building.HardwareStore, star*100),
		}
	},
	// 企鹅机械->零件厂
	// 企鹅机械->全部
	building.PenguinMachinery: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.SparePartsFactory, star*100),
			globalBuff: globalBuff(star * 10),
		}
	},
	// 五金店->零件厂
	building.HardwareStore: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.SparePartsFactory, star*100)}
	},

	// 商贸中心->洋房花园
	// todo 商贸中心->供货
	building.BusinessCentre: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.GardenHouse, star*100)}
	},

	//	洋房花园->商贸中心
	// todo 洋房花园->供货
	building.GardenHouse: func(star int) Fetters {
		return Fetters{fetterBuff: NewFetterBuff(building.BusinessCentre, star*100)}
	},
	// 民食斋->在线
	// 民食斋->空中别墅
	building.FolkFood: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.Villa, star*100),
			onlineBuff: onlineBuff(star * 20),
		}
	},
	// 空中别墅->在线
	// 空中别墅->民食斋
	building.Villa: func(star int) Fetters {
		return Fetters{
			fetterBuff: NewFetterBuff(building.FolkFood, star*100),
			onlineBuff: onlineBuff(star * 20),
		}
	},
	// 媒体之声->离线
	// 媒体之声->全部
	building.MediaVoice: func(star int) Fetters {
		return Fetters{
			globalBuff:  globalBuff(star * 5),
			offlineBuff: offlineBuff(star * 10),
		}
	},
	// 复兴公馆->离线
	// todo 复兴公馆->供货
	building.RenaissanceMansion: func(star int) Fetters {
		return Fetters{offlineBuff: offlineBuff(star * 10)}
	},

	// 中式小楼->在线
	// 中式小楼->住宅
	building.ChineseSmallHouse: func(star int) Fetters {
		return Fetters{
			onlineBuff: onlineBuff(star * 20),
			classBuff:  NewClassBuff(building.Residence, 15*star),
		}
	},
	// 人才公寓->在线
	// 人才公寓->工业
	building.TalentApartment: func(star int) Fetters {
		return Fetters{
			onlineBuff: onlineBuff(star * 20),
			classBuff:  NewClassBuff(building.Industry, 15*star),
		}
	},
	// todo 小型公寓->供货
	building.StudioApartment: func(star int) Fetters {
		return Fetters{}
	},
}
