package viridescent

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	character2 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/reactable"
	"dps/src/template/artifact"
	"dps/src/template/modifier"
)

type Viridescent struct {
	*artifact.Artifact
}

func ApplyTwoSetStates(c character2.Character, cfg *[4]character2.CharWrapper, i int, framei int, ais *[]combat.AttackInfo) {
	AlreadyHaveBuff := false
	for j := range (*cfg)[i].Mods {
		if (*cfg)[i].Mods[j].Name == "Viridescent2Buff" {
			AlreadyHaveBuff = true
		}
	}
	if AlreadyHaveBuff == false {
		mod := modifier.Mods{
			Name:       "Viridescent2Buff",
			StartFrame: 0,
			Dur:        -1,
			CharIndex:  i,
		}
		mod.Modifier[attributes.AnemoP] = 0.15
		(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
	}
}

func ApplyFourSetStates(c character2.Character, cfg *[4]character2.CharWrapper, charIndex int, framei int, ais *[]combat.AttackInfo, enemy *enemy.Enemy) {
	//四件套：扩散反应造成的伤害提升60%。根据扩散的元素类型，降低收到影响的敌人40%的对应元素抗性，持续10秒。
	aiIndex := len(*ais) - 1
	exist4Buff := false
	for tagi := range (*ais)[aiIndex].Tags {
		if (*ais)[aiIndex].Tags[tagi] == "Viridescent4Buff" {
			exist4Buff = true
		}
	}
	if !exist4Buff {
		(*ais)[aiIndex].ReactBonus[reactable.Swirl] += 0.6
		(*ais)[aiIndex].Tags = append((*ais)[aiIndex].Tags, "Viridescent4Buff")
	}
	if (*ais)[aiIndex].Element == attributes.Anemo && (*ais)[aiIndex].ActorIndex == charIndex && (enemy.Element == attributes.Pyro || enemy.Element == attributes.Cryo || enemy.Element == attributes.Hydro || enemy.Element == attributes.Electro) && cfg[charIndex].Active == true {
		player.UseElementAttach(c, ais, enemy, len(*ais)-1)
		if (*ais)[aiIndex].Trigger == true {
			//触发扩散反应，减抗，如果有就先删除
			for j := len((*enemy).EnemyMod) - 1; j >= 0; j-- {
				modDel := (*enemy).EnemyMod[j]
				if (modDel.Name == "Viridescent4Buff" && modDel.Res[enemy.Element] != 0) || modDel.StartFrame+modDel.Dur < framei {
					(*enemy).EnemyMod = append((*enemy).EnemyMod[:j], (*enemy).EnemyMod[j+1:]...)
				}
			}
			enemyMod := modifier.EnemyMod{
				Name:       "Viridescent4Buff",
				StartFrame: framei - 1,
				Dur:        600,
			}
			enemyMod.Res[enemy.Element] = -0.4
			(*enemy).EnemyMod = append((*enemy).EnemyMod, enemyMod)
		}
	}
}

func (a *Viridescent) ApplyArtifactMod(c character2.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character2.CharWrapper, setIndex int, charIndex int, enemy *enemy.Enemy) {
	sets := cfg[charIndex].Equip.Sets
	//if framei == 52 {
	//	fmt.Println("看看为什么风套没触发")
	//}
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
			ApplyFourSetStates(c, cfg, charIndex, framei, attacks, enemy)
		}
	}
}
