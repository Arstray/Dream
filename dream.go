package dream

import (
	"dream/buff"
	"dream/building"
	"fmt"
)

// Group 实际使用的建筑组合
// 通过与四类加成组合，计算当前倍率
type Group struct {
	buildings [9]building.Building
	album     *buff.Album
	city      *buff.City
	fetters   *buff.Fetters
	policy    *buff.Policy
}

// 返回倍率 a=x*y*z*c
// x建筑 y政策 z相册 c城市任务
// 离线时，失去所有在线加成，获得所有离线加成，最终结果*0.5
func (g Group) Profit() {
	fmt.Println("在线情况")
	fmt.Printf("建筑名称 \t 星级 \t 关联 \t 相册 \t 政策 \t 城市任务 \t 外界倍率 \t 自身倍率 \t 相对生产值 \n")
	for _, b := range g.buildings {
		fmt.Printf(
			"%s \t %d \t\t %.2f \t %.2f \t %.2f \t %.2f \t\t %6.2f \t %6.2f \t %.2f \n",
			b.Name, b.Star.Value(), g.fetters.OnBuff(b), g.album.OnBuff(b), g.policy.OnBuff(b), g.city.OnBuff(b),
			g.fetters.OnBuff(b)*g.album.OnBuff(b)*g.policy.OnBuff(b)*g.city.OnBuff(b),
			BaseGrowUp(b.Base, b.Star.Value()),
			g.fetters.OnBuff(b)*g.album.OnBuff(b)*g.policy.OnBuff(b)*g.city.OnBuff(b)*BaseGrowUp(b.Base, b.Star.Value()),
		)
	}

	fmt.Println("离线情况")
	fmt.Printf("名称 \t\t 星级 \t 关联 \t 相册 \t 政策 \t 城市任务 \t 外界倍率 \t 自身倍率 \t 相对生产值 \n")
	for _, b := range g.buildings {
		fmt.Printf(
			"%s \t %d \t\t %.2f \t %.2f \t %.2f \t %.2f \t\t %6.2f \t %6.2f \t %.2f \n",
			b.Name, b.Star.Value(), g.fetters.OffBuff(b), g.album.OffBuff(b), g.policy.OffBuff(b), g.city.OffBuff(b),
			g.fetters.OffBuff(b)*g.album.OffBuff(b)*g.policy.OffBuff(b)*g.city.OffBuff(b)*0.5,
			BaseGrowUp(b.Base, b.Star.Value()),
			g.fetters.OnBuff(b)*g.album.OnBuff(b)*g.policy.OnBuff(b)*g.city.OnBuff(b)*BaseGrowUp(b.Base, b.Star.Value())*0.5,
		)
	}
}

// NewGroup 新建一个建筑组合
func NewGroup(buildings [9]building.Building) Group {
	fetters := buff.NewFetterBuffs()
	for _, building := range buildings {
		fetters.Composition(building)
	}

	return Group{
		buildings: buildings,
		fetters:   fetters,
		album:     buff.NewAlbum(0, 0, 0, nil),
		city:      buff.NewCity(0, 0, 0, nil, nil),
		policy:    buff.NewPolicy(0, 0, 0, nil),
	}
}

// SetProfit 设置政策或活动
func (g Group) SetProfit(policy *buff.Policy) {
	g.policy.Composition(policy)
}

// SetCity 设置城市任务
func (g Group) SetCity(city *buff.City) {
	g.city.Composition(city)
}

// SetAlbum 设置相册
func (g Group) SetAlbum(album *buff.Album) {
	g.album.Composition(album)
}

// BaseGrowUp 基准升级
func BaseGrowUp(base float64, star int) float64 {
	var res = base
	for i := 1; i <= star; i++ {
		res *= float64(i)
	}
	return res
}
