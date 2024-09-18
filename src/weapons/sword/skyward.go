package sword

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
	Skyward_Base_atk          = 608
	Skyward_Hp_percent        = 0
	Skyward_Atk_percent       = 0
	Skyward_Def_percent       = 0
	Skyward_Elemental_mastery = 0
	Skyward_Crit_rate         = 0
	Skyward_Crit_damage       = 0
	Skyward_Energy_recharge   = 0.551
)

type Skyward struct {
	*weapon.Weapon
}

func (w Skyward) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	(*cfg)[i].BaseStats[attributes.BaseATK] += Skyward_Base_atk
	(*cfg)[i].BaseStats[attributes.ER] += Skyward_Energy_recharge
}

func (w Skyward) CheckWeapon(c character3.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, charIndex int, enemy *enemy.Enemy) {
	// 暴击率提升（4%/5%/6%/7%/8%）；施放元素爆发时，获得破空之势：移动速度提升10%，攻击速度提升10%，普通攻击和重击命中时，额外造成（20%/25%/30%/35%/40%）攻击力的伤害，持续12秒。
	attackIndex := len(*attacks) - 1
	if (*attacks)[attackIndex].ActorIndex != charIndex {
		return
	}
	mod := modifier.Mods{
		Name:       "SkywardPermanent",
		Modifier:   [attributes.EndStatType]float64{},
		StartFrame: framei,
		Dur:        -1,
		CharIndex:  charIndex,
	}
	mod.Modifier[attributes.CR] = 0.03 + float64(cfg[charIndex].Weapon.Refine)*0.01
}
