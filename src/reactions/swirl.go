package reactions

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	enemy "dps/src/enemies"
	"dps/src/reactable"
	"fmt"
)

func DamageSwirl(ai *combat.AttackInfo, enemy *enemy.Enemy, em float64, level int, frame int) (float64, bool) {
	var damage float64
	var tag bool

	if ai.Element == attributes.Anemo && swirlYes(enemy.Element) {
		tag = true
		if enemy.Element == attributes.Pyro {
			fmt.Println("火扩散，扩散前火元素量：", enemy.ElementQuantity)
		} else if enemy.Element == attributes.Cryo {
			fmt.Println("冰扩散，扩散前冰元素量：", enemy.ElementQuantity)
		} else if enemy.Element == attributes.Electro {
			fmt.Println("雷扩散，扩散前雷元素量：", enemy.ElementQuantity)
		} else if enemy.Element == attributes.Hydro {
			fmt.Println("水扩散，扩散前水元素量：", enemy.ElementQuantity)
		}
		res := enemy.Res[enemy.Element]
		for i := range enemy.EnemyMod {
			if enemy.EnemyMod[i].StartFrame < frame && enemy.EnemyMod[i].StartFrame+enemy.EnemyMod[i].Dur > frame {
				res += enemy.EnemyMod[i].Res[enemy.Element]
			}
		}
		//fmt.Println("应用减抗后的抗性:", res)
		resmod := 1 - res/2
		if res >= 0 && res < 0.75 {
			resmod = 1 - res
		} else if res > 0.75 {
			resmod = 1 / (4*res + 1)
		}
		ai.Trans = true
		levelMult := reactionLvlBase[level-1]
		emBonus := (16.0 * em) / (2000 + em)
		damage = 0.6 * levelMult * (1 + emBonus + ai.ReactBonus[reactable.Swirl]) * resmod
		UpdateElement(ai, enemy)
	}
	return damage, tag
}

func swirlYes(element attributes.Element) bool {
	if element == attributes.Pyro || element == attributes.Cryo || element == attributes.Hydro || element == attributes.Electro {
		return true
	}
	return false
}
