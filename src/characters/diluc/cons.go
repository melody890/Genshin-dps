package diluc

import (
	"dps/src/core/attributes"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
)

// 命座一：对生命值高于50%的敌人，造成的伤害提高15%
func c1(cfg *[4]character3.CharWrapper, enemy *enemy.Enemy, startFrame int, charIndex int) {
	hasC1Buff := false
	for i := range (*cfg)[charIndex].Mods {
		if enemy.Hp >= enemy.Hp/2 {
			if (*cfg)[charIndex].Mods[i].Name == "DilucC1" {
				hasC1Buff = true
			}
		}
	}
	if enemy.Hp >= enemy.Hp/2 && hasC1Buff == false {
		mod := modifier.Mods{
			Name:       "DilucC1",
			Modifier:   [attributes.EndStatType]float64{},
			StartFrame: 0,
			Dur:        -1,
			CharIndex:  charIndex,
		}
		mod.Modifier[attributes.DmgP] = 0.15
		(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
	}
}

// 二命要挨打，我先不写了
func c2() {}

func c4() {}
