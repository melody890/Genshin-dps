package reactions

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	enemy "dps/src/enemies"
	"fmt"
)

func ElementDecay(enemy *enemy.Enemy) {
	var decayRate float64 // 一帧掉多少元素量
	if enemy.ElementType == attributes.Weak {
		//弱元素假定为1单位量，用9.5秒完全衰减
		decayRate = (1 * 0.8) / (9.5 * 60)
	} else if enemy.ElementType == attributes.Medium {
		//中元素假定为1.5单位量，用10.75秒完全衰减
		decayRate = (1.5 * 0.8) / (10.75 * 60)
	} else if enemy.ElementType == attributes.Strong {
		//强元素假定为2单位量，用12秒完全衰减
		decayRate = (2 * 0.8) / (12 * 60)
	} else if enemy.ElementType == attributes.SuperStrong {
		//超强元素假定为4单位量，用17秒完全衰减
		decayRate = (4 * 0.8) / (17 * 60)
	}
	enemy.ElementQuantity -= decayRate
	if enemy.ElementQuantity <= 0 {
		enemy.ElementQuantity = 0
		enemy.ElementType = attributes.Clear
		enemy.Element = attributes.NoElement
	}
}

func AddElement(ai combat.AttackInfo, enemy *enemy.Enemy) {
	if enemy.Element == attributes.NoElement && ai.Element != attributes.Physical && ai.Element != attributes.Anemo && ai.Element != attributes.Geo {
		enemy.Element = ai.Element
		enemy.ElementQuantity = ai.ElementQuantity * 0.8
		switch ai.ElementQuantity {
		case 1.0:
			enemy.ElementType = attributes.Weak
		case 1.5:
			enemy.ElementType = attributes.Medium
		case 2.0:
			enemy.ElementType = attributes.Strong
		case 4.0:
			enemy.ElementType = attributes.SuperStrong
		}
		fmt.Println(ai.Abil, "上", ai.Element, "元素", enemy.ElementQuantity, "元素量")
	}
}

func SupplyElement(ai combat.AttackInfo, enemy *enemy.Enemy) {
	if ai.Element == enemy.Element {
		if ai.ElementQuantity*0.8 > enemy.ElementQuantity {
			enemy.ElementQuantity = ai.ElementQuantity * 0.8
			fmt.Println("补充元素", enemy.Element, "至", enemy.ElementQuantity)
		}
	}
}

func UpdateElement(ai *combat.AttackInfo, enemy *enemy.Enemy) {
	if ai.Element == attributes.Pyro && enemy.Element == attributes.Cryo {
		enemy.ElementQuantity -= 2 * ai.ElementQuantity
		//额外元素不残留
		if enemy.ElementQuantity <= 0 {
			enemy.ElementQuantity = 0
			//刷新初始元素附着强度
			enemy.ElementType = attributes.Clear
			enemy.Element = attributes.NoElement
		}
		fmt.Println("反应后剩余冰元素量：", enemy.ElementQuantity)
	} else if ai.Element == attributes.Cryo && enemy.Element == attributes.Pyro {
		enemy.ElementQuantity -= 0.5 * ai.ElementQuantity
		if enemy.ElementQuantity <= 0 {
			enemy.ElementQuantity = 0
			enemy.ElementType = attributes.Clear
			enemy.Element = attributes.NoElement
		}
		fmt.Println("反应后剩余火元素量：", enemy.ElementQuantity)
	} else if ai.Element == attributes.Anemo && (enemy.Element == attributes.Pyro || enemy.Element == attributes.Cryo || enemy.Element == attributes.Electro || enemy.Element == attributes.Dendro) {
		elementBuffer := enemy.Element
		enemy.ElementQuantity -= 0.5 * ai.ElementQuantity
		if enemy.ElementQuantity <= 0 {
			enemy.ElementQuantity = 0
			enemy.ElementType = attributes.Clear
			enemy.Element = attributes.NoElement
		}
		if elementBuffer == attributes.Pyro {
			fmt.Println("扩散后剩余火元素量：", enemy.ElementQuantity)
		} else if elementBuffer == attributes.Cryo {
			fmt.Println("扩散后剩余冰元素量：", enemy.ElementQuantity)
		} else if elementBuffer == attributes.Electro {
			fmt.Println("扩散后剩余雷元素量：", enemy.ElementQuantity)
		} else if elementBuffer == attributes.Hydro {
			fmt.Println("扩散后剩余水元素量：", enemy.ElementQuantity)
		}

	}
}
