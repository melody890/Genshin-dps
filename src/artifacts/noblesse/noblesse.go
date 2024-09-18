package noblesse

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	"dps/src/core/player/character"
	character2 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/artifact"
	"dps/src/template/modifier"
	"strings"
)

type Noblesse struct {
	*artifact.Artifact
}

func ApplyTwoSetStates(c character2.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	//大招伤害增加
	if strings.Contains((*ais)[len((*ais))-1].Abil, "Burst") && framei < ((*ais)[len((*ais))-1].AttackFrame+1) {
		// 判断现在的mod里不是已经有了
		AlreadyHaveBuff := false
		for j := range (*cfg)[i].Mods {
			if (*cfg)[i].Mods[j].Name == "Noblesse2Buff" {
				AlreadyHaveBuff = true
			}
		}
		// 没有的话加一个大招的
		if AlreadyHaveBuff == false {
			mod := modifier.Mods{
				Name:       "Noblesse2Buff",
				StartFrame: 0,
				Dur:        -1,
				CharIndex:  i,
				Condition:  attacks.AttackBurst,
			}
			mod.Modifier[attributes.DmgP] = 0.2
			(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
		}
	}
}

func ApplyFourSetStates(c character2.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	if strings.Contains((*ais)[len((*ais))-1].Abil, "Burst") && framei == ((*ais)[len((*ais))-1].Startframe) {
		// 判断一下是不是穿宗室的这个人开大了
		charIndex := (*ais)[len(*ais)-1].ActorIndex
		if cfg[charIndex].Equip.Sets.SetName[0] != "noblesse" {
			return
		}
		// 判断现在的mod里不是已经有了
		AlreadyHaveBuff := false
		for j := range (*cfg)[i].Mods {
			if (*cfg)[i].Mods[j].Name == "Noblesse4Buff" {
				AlreadyHaveBuff = true
			}
		}
		// 没有的话加一个所有人攻击力
		if AlreadyHaveBuff == false {
			mod := modifier.Mods{
				Name:       "Noblesse4Buff",
				StartFrame: (*ais)[len(*ais)-1].AttackFrame + 2,
				Dur:        720,
				CharIndex:  6,
			}
			mod.Modifier[attributes.ATKP] = 0.2
			for charIndex := range cfg {
				(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
			}

		}
	}
}

func (a *Noblesse) ApplyArtifactMod(c character2.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character2.CharWrapper, setIndex int, charIndex int, enemy *enemy.Enemy) {
	sets := cfg[charIndex].Equip.Sets
	// 删除过期的
	for i := range cfg {
		for j := len((*cfg)[i].Mods) - 1; j >= 0; j-- {
			//for j := range (*cfg)[i].Mods{
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
