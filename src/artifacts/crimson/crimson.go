package crimson

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/reactable"
	"dps/src/template/artifact"
	"dps/src/template/modifier"
	"strings"
)

type Crimson struct {
	*artifact.Artifact
}

func ApplyTwoSetStates(c character.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	// 判断现在的mod里不是已经有了
	AlreadyHaveBuff := false
	for j := range (*cfg)[i].Mods {
		if (*cfg)[i].Mods[j].Name == "Crimson2Buff" {
			AlreadyHaveBuff = true
		}
	}
	// 没有的话加一个15火伤
	if AlreadyHaveBuff == false {
		mod := modifier.Mods{
			Name:       "Crimson2Buff",
			StartFrame: 0,
			Dur:        -1,
			CharIndex:  i,
		}
		mod.Modifier[attributes.PyroP] = 0.15
		(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
	}
}

func ApplyFourSetStates(c character.Character, cfg *[4]character.CharWrapper, i int, framei int, ais *[]combat.AttackInfo, enemy *enemy.Enemy) {
	//超载、燃烧、烈绽放反应造成的伤害提升40%，蒸发、融化反应的加成系数提高15%。施放元素战技后的10秒内，2件套的效果提高50%，该效果最多叠加3次。
	//TODO:超载燃烧烈绽放先稍等一下，先写增幅反应
	attackIndex := len(*ais) - 1
	if (*ais)[attackIndex].Trigger == true && (*ais)[attackIndex].Element == attributes.Pyro && (enemy.Element == attributes.Cryo || enemy.Element == attributes.Hydro) {
		if (*ais)[attackIndex].ReactBonus[reactable.Melt] != 0.15 {
			(*ais)[attackIndex].ReactBonus[reactable.Melt] += 0.15
			(*ais)[attackIndex].ReactBonus[reactable.Vaporize] += 0.15
		}
	}
	if framei == (*ais)[attackIndex].Startframe && strings.Contains((*ais)[attackIndex].Abil, "Skill") && (*ais)[attackIndex].ActorIndex == i {
		mod := modifier.Mods{
			Name:       "Crimson4Buff",
			StartFrame: (*ais)[len(*ais)-1].Startframe + 1,
			Dur:        600,
			CharIndex:  i,
		}
		mod.Modifier[attributes.PyroP] = 0.075
		//检查一下现在的buff有几个，如果已经有3个就不加了，虽然好像没有谁能叠三个
		count := 0
		for j := range (*cfg)[i].Mods {
			if (*cfg)[i].Mods[j].Name == "Crimson4Buff" {
				count++
			}
		}
		if count < 3 {
			(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
		}
	}
}

func (a Crimson) ApplyArtifactMod(c character.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character.CharWrapper, setIndex int, charIndex int, enemy *enemy.Enemy) {
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
			ApplyFourSetStates(c, cfg, charIndex, framei, attacks, enemy)
		}
	}
}
