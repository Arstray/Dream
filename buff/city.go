package buff

import (
	"dream/building"
)

// 由于城市任务不唯一，故不作预设封装
// 城市任务加成是不单调增长的，但更换建筑会重新构建组，所以加成是重新累加计算的
type City struct {
	globalBuff
	onlineBuff
	offlineBuff
	classBuff
	fetterBuff
}

func NewCity(gB, onB, offB int, cB classBuff, tB fetterBuff) *City {
	return &City{
		globalBuff:  globalBuff(gB),
		onlineBuff:  onlineBuff(onB),
		offlineBuff: offlineBuff(offB),
		classBuff:   cB,
		fetterBuff:  tB,
	}
}

// OnBuff 在线城市任务加成
func (c *City) OnBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += c.globalBuff.Value()

	res += c.onlineBuff.Value()

	res += c.classBuff.Value(b.Class)

	res += c.fetterBuff.Value(b.Type)

	return float64(res) / 100
}

func (c *City) OffBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += c.globalBuff.Value()

	res += c.offlineBuff.Value()

	res += c.classBuff.Value(b.Class)

	res += c.fetterBuff.Value(b.Type)

	return float64(res) / 100
}

func (c *City) Composition(newCity *City) {
	c.globalBuff += newCity.globalBuff
	c.onlineBuff += newCity.onlineBuff
	c.offlineBuff += newCity.offlineBuff
	for k, v := range newCity.classBuff {
		c.classBuff[k] += v
	}
	for k, v := range newCity.fetterBuff {
		c.fetterBuff[k] += v
	}
}
