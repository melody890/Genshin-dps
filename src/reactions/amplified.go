package reactions

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	enemy "dps/src/enemies"
	"dps/src/reactable"
	"fmt"
)

// CheckAmplified 增幅反应，其中AmpMult是增幅反应倍率，返回的是反应乘区
func CheckAmplified(ai *combat.AttackInfo, enemy *enemy.Enemy, em float64) float64 {
	if ai.Element == attributes.Pyro && enemy.Element == attributes.Cryo {
		//火打冰融化
		fmt.Println("火冰融化， 反应前剩余冰元素量：", enemy.ElementQuantity)
		ai.Amped = true
		ai.AmpMult = 2
		reactBonus := ai.ReactBonus[reactable.Melt]
		emBonus := (2.78 * em) / (1400 + em)
		UpdateElement(ai, enemy)
		return ai.AmpMult * (1 + emBonus + reactBonus)
	} else if ai.Element == attributes.Cryo && enemy.Element == attributes.Pyro {
		//冰打火融化
		fmt.Println("冰火融化， 反应前剩余火元素量：", enemy.ElementQuantity)
		ai.Amped = true
		ai.AmpMult = 1.5
		reactBonus := ai.ReactBonus[reactable.Melt]
		emBonus := (2.78 * em) / (1400 + em)
		UpdateElement(ai, enemy)
		return ai.AmpMult * (1 + emBonus + reactBonus)
	} else {
		return -1
	}
}
