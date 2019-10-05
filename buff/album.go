package buff

import (
	"dream/building"
)

// Album 相册
// 相册预设信息来源:http://www.paopaoche.net/sj/164396.html
// 相册加成是单调增长的，已经获得的加成则不会消失
// 翻译于百度......
type Album struct {
	globalBuff
	onlineBuff
	offlineBuff
	classBuff
}

func NewAlbum(gB, onB, offB int, cB classBuff) *Album {
	return &Album{
		globalBuff:  globalBuff(gB),
		onlineBuff:  onlineBuff(onB),
		offlineBuff: offlineBuff(offB),
		classBuff:   cB,
	}
}

// OnBuff 在线相册加成
func (a *Album) OnBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += a.globalBuff.Value()

	res += a.onlineBuff.Value()

	res += a.classBuff.Value(b.Class)

	return float64(res) / 100
}

// OffBuff 离线相册加成
func (a *Album) OffBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += a.globalBuff.Value()

	res += a.offlineBuff.Value()

	res += a.classBuff.Value(b.Class)

	return float64(res) / 100
}

// Composition 加成叠加
func (a *Album) Composition(newAlbum *Album) {
	a.globalBuff += newAlbum.globalBuff
	a.onlineBuff += newAlbum.onlineBuff
	a.offlineBuff += newAlbum.offlineBuff
	if a.classBuff == nil {
		a.classBuff = make(map[building.Class]int)
	}
	for k, v := range newAlbum.classBuff {
		a.classBuff[k] += v
	}
}

// 与逻辑无关的固有配置
// todo 供货

// 国家级单个相册
var (
	// 一带一路
	OneBeltOneRoad = NewAlbum(0, 0, 0, classBuff{building.Business: 60})
	// 一国两制
	OneCountryTwoSystems = NewAlbum(20, 0, 0, nil)
	//兴军强国
	MilitaryPower = NewAlbum(0, 0, 20, nil)
	//体育大国
	SportsPower = NewAlbum(0, 0, 20, nil)

	//改革开放
	ReformOpen = NewAlbum(20, 0, 0, nil)
	//中国制造
	MadeInChina = NewAlbum(0, 0, 0, classBuff{building.Industry: 60})
	//减贫奇迹
	PovertyReductionMiracle = NewAlbum(0, 0, 0, classBuff{building.Residence: 60})
	//美丽中国
	BeautifulChina = NewAlbum(0, 20, 0, nil)
	//文化自信
	CulturalConfidence = NewAlbum(0, 20, 0, nil)
	//获得感幸福感安全感
	GettingASenseOfHappinessAndSecurity = NewAlbum(20, 0, 0, nil)
)

// 区域级相册组
var (
	Chinese = NewAlbum(
		60, 40, 40,
		classBuff{
			building.Residence: 60,
			building.Business:  60,
			building.Industry:  60,
		},
	)

	// 黑龙江
	// 吉林
	// 辽宁
	// 东三省四日游
	ThreeEastern = NewAlbum(
		60, 50, 40,
		classBuff{
			building.Residence: 30,
			building.Business:  180,
			building.Industry:  150,
		},
	)

	//广东
	//福建
	//海南
	//广西
	//海滨热梦
	BeachDream = NewAlbum(
		70, 80, 70,
		classBuff{
			building.Residence: 210,
			building.Business:  150,
			building.Industry:  90,
		},
	)

	// 河南
	// 山东
	// 安徽
	// 平原古风
	PlainOldStyle = NewAlbum(
		50, 40, 60,
		classBuff{
			building.Residence: 150,
			building.Business:  60,
			building.Industry:  90,
		},
	)

	//上海
	//江苏
	//浙江
	//长江三角四日游
	YangtzeRiverDelta = NewAlbum(
		40, 40, 70,
		classBuff{
			building.Residence: 150,
			building.Business:  150,
			building.Industry:  60,
		},
	)
	// 重庆
	// 湖北
	// 湖南
	// 江西
	// 江水人家
	RiversAndLakes = NewAlbum(
		70, 70, 80,
		classBuff{
			building.Residence: 120,
			building.Business:  150,
			building.Industry:  150,
		},
	)

	// 北京
	// 天津
	// 河北
	// 京津冀四日游
	Beijing = NewAlbum(
		50, 80, 40,
		classBuff{
			building.Residence: 60,
			building.Business:  60,
			building.Industry:  150,
		},
	)
	//云南
	//贵州
	//四川
	//云贵川四日游
	YunGui = NewAlbum(
		40, 50, 60,
		classBuff{
			building.Residence: 150,
			building.Business:  90,
			building.Industry:  60,
		},
	)
	//陕西
	//山西
	//内蒙古
	//宁夏
	//黄土风情
	LoessStyle = NewAlbum(
		70, 50, 80,
		classBuff{
			building.Residence: 150,
			building.Business:  60,
			building.Industry:  150,
		},
	)
	//甘肃
	//青海
	//西藏
	//新疆
	//西部之旅
	WestTrip = NewAlbum(
		80, 80, 70,
		classBuff{
			building.Residence: 90,
			building.Business:  90,
			building.Industry:  150,
		},
	)
	//香港
	//澳门
	//台湾
	//港澳台四日游
	HongKong = NewAlbum(
		60, 70, 50,
		classBuff{
			building.Residence: 30,
			building.Business:  150,
			building.Industry:  90,
		},
	)
)
