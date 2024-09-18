package emblem

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/artifact"
	"dps/src/template/modifier"
	"strings"
)

type Emblem struct {
	*artifact.Artifact
}

//·····················
//········绝缘套········
//·····················

func ApplyTwoSetStates(c character.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	AlreadyHaveBuff := false
	for j := range (*cfg)[i].Mods {
		if (*cfg)[i].Mods[j].Name == "Emblem2Buff" {
			AlreadyHaveBuff = true
		}
	}
	// 没有的话加一个20充能
	if AlreadyHaveBuff == false {
		mod := modifier.Mods{
			Name:       "Emblem2Buff",
			StartFrame: 0,
			Dur:        -1,
			CharIndex:  i,
		}
		mod.Modifier[attributes.ER] = 0.2
		(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
	}
}

func ApplyFourSetStates(c character.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	//四件套：基于元素充能效率的25%，提高元素爆发造成的伤害。至多通过这种方式获得75%的提升。
	if strings.Contains((*ais)[len(*ais)-1].Abil, "Burst") && framei < ((*ais)[len(*ais)-1].AttackFrame+1) {
		// 判断现在的mod里不是已经有了
		AlreadyHaveBuff := false
		for j := range (*cfg)[i].Mods {
			if (*cfg)[i].Mods[j].Name == "Emblem4Buff" {
				AlreadyHaveBuff = true
			}
		}
		// 没有的话加一个大招的
		if AlreadyHaveBuff == false {
			mod := modifier.Mods{
				Name:       "Emblem4Buff",
				StartFrame: (*ais)[len(*ais)-1].Startframe - 1,
				Dur:        (*ais)[len((*ais))-1].AttackFrame - (*ais)[len((*ais))-1].Startframe + 2,
				CharIndex:  i,
				Condition:  attacks.AttackBurst,
			}
			mod.Modifier[attributes.DmgP] = cfg[i].BaseStats[attributes.ER]*0.25 + 0.2*0.25 // 最后这个是二件套给的
			//fmt.Println("这是在圣遗物里面的，看一下充能效率", cfg[i].BaseStats[attributes.ER])
			if mod.Modifier[attributes.DmgP] >= 0.75 {
				mod.Modifier[attributes.DmgP] = 0.75
			}
			(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
		}
	}
}

func (a Emblem) ApplyArtifactMod(c character.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character.CharWrapper, setIndex int, charIndex int, enemy *enemy.Enemy) {
	sets := cfg[charIndex].Equip.Sets
	// 删除过期的
	for i := range cfg {
		for j := len((*cfg)[i].Mods) - 1; j >= 0; j-- {
			if (*cfg)[i].Mods[j].StartFrame+(*cfg)[i].Mods[j].Dur <= framei && (*cfg)[i].Mods[j].Dur != -1 {
				//fmt.Println((*cfg)[i].Mods[j].StartFrame + (*cfg)[i].Mods[j].Dur)
				(*cfg)[i].Mods = append((*cfg)[i].Mods[:j], (*cfg)[i].Mods[j+1:]...)
			}
		}
	}

	for i := range sets.SetNum {
		if sets.SetNum[i] >= 2 {
			ApplyTwoSetStates(c, cfg, charIndex, framei, attacks)
		}
		if sets.SetNum[i] >= 4 {
			ApplyFourSetStates(c, cfg, charIndex, framei, attacks)
		}
	}
}
