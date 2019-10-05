package buff

import (
	"dream/building"
)

// Policy 政策
// 政策预设信息来自游戏
// 政策加成是不单调增长的，但更换建筑会重新构建组，所以加成是重新累加计算的
// 家园之光Buff属于政策加成
type Policy struct {
	globalBuff
	onlineBuff
	offlineBuff
	classBuff
}

func NewPolicy(gB, onB, offB int, cB classBuff) *Policy {
	return &Policy{
		globalBuff:  globalBuff(gB),
		onlineBuff:  onlineBuff(onB),
		offlineBuff: offlineBuff(offB),
		classBuff:   cB,
	}
}

// OnBuff 在线政策加成
func (p *Policy) OnBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += p.globalBuff.Value()

	res += p.onlineBuff.Value()

	res += p.classBuff.Value(b.Class)

	return float64(res) / 100
}

// OffBuff 离线政策加成
func (p *Policy) OffBuff(b building.Building) float64 {
	// 基础倍率
	var res = 100

	res += p.globalBuff.Value()

	res += p.offlineBuff.Value()

	res += p.classBuff.Value(b.Class)

	return float64(res) / 100
}
// Composition 加成叠加
func (p *Policy) Composition(newPolicy *Policy) {
	p.globalBuff += newPolicy.globalBuff
	p.onlineBuff += newPolicy.onlineBuff
	p.offlineBuff += newPolicy.offlineBuff
	if p.classBuff == nil {
		p.classBuff = make(map[building.Class]int)
	}
	for k, v := range newPolicy.classBuff {
		p.classBuff[k] += v
	}
}

var (
	Stage1 = NewPolicy(
		100, 0, 0,
		map[building.Class]int{
			building.Residence: 300,
			building.Business:  300,
		},
	)

	Stage2 = NewPolicy(
		200, 200, 200,
		map[building.Class]int{
			building.Industry: 600,
		},
	)
	// 由于还完全未解锁第三阶段，不能给出整体预设
)
