package claymore

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
	"dps/src/template/weapon"
)

const (
	Spine_Base_atk          = 510
	Spine_Hp_percent        = 0
	Spine_Atk_percent       = 0
	Spine_Def_percent       = 0
	Spine_Elemental_mastery = 0
	Spine_Crit_rate         = 0.276
	Spine_Crit_damage       = 0
	Spine_Energy_recharge   = 0
)

type Spine struct {
	*weapon.Weapon
}

func (w Spine) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	(*cfg)[i].BaseStats[attributes.BaseATK] += Spine_Base_atk
	(*cfg)[i].BaseStats[attributes.CR] += Spine_Crit_rate
}

func (w Spine) CheckWeapon(c character3.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, charIndex int, enemy *enemy.Enemy) {
	attackIndex := len(*attacks) - 1
	if (*attacks)[attackIndex].ActorIndex != charIndex {
		return
	}
	//charIndex := attacks[attackIndex].ActorIndex
	// 删除现有buff
	for i := len((*cfg)[charIndex].Mods) - 1; i >= 0; i-- {
		modDel := (*cfg)[charIndex].Mods[i]
		if modDel.Name == "SpinePermanent" {
			(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[i+1:]...)
		}
	}
	// 角色在场上时，每4秒提升（6%/7%/8%/9%/10%）造成的伤害，（3%/2.7%/2.4%/2.2%/2.0%）受到的伤害。该效果最多叠加5层，不随角色退场重置，受到伤害后会减少1层效果。
	mod := modifier.Mods{
		Name:       "SpinePermanent",
		Modifier:   [attributes.EndStatType]float64{},
		StartFrame: framei,
		Dur:        -1,
		CharIndex:  charIndex,
	}
	permanentBuff := float64(cfg[charIndex].Weapon.Refine)*0.05 + 0.25
	mod.Modifier[attributes.DmgP] = permanentBuff
	(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
}
