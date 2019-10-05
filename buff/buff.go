package buff

import "dream/building"


type (
	// 全局buff
	globalBuff int
	// 在线buff
	onlineBuff int
	// 离线buff
	offlineBuff int

	// 建筑类型buff
	classBuff map[building.Class]int

	// 特定建筑buff
	fetterBuff map[building.Type]int
)

func (v globalBuff) Value() int { return int(v) }

func (v onlineBuff) Value() int { return int(v) }

func (v offlineBuff) Value() int { return int(v) }

func NewClassBuff(key building.Class, value int) classBuff {
	return map[building.Class]int{
		key: value,
	}
}

func (c classBuff) Value(class building.Class) int {
	if v, ok := c[class]; ok {
		return v
	}
	return 0
}

func (c classBuff) Add(key building.Class, value int) classBuff {
	if _, ok := c[key]; ok {
		c[key] += value
		return c
	}
	c[key] = value
	return c
}

func NewFetterBuff(key building.Type, value int) fetterBuff {
	return map[building.Type]int{
		key: value,
	}
}

func (b fetterBuff) Value(buildingType building.Type) int {
	if v, ok := b[buildingType]; ok {
		return v
	}
	return 0
}
func (b fetterBuff) Add(key building.Type, value int) fetterBuff {
	if _, ok := b[key]; ok {
		b[key] += value
		return b
	}
	b[key] = value
	return b
}
