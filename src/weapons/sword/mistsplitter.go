package sword

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
	"dps/src/template/weapon"
	"strings"
)

const (
	Mistsplitter_Base_atk          = 674
	Mistsplitter_Hp_percent        = 0
	Mistsplitter_Atk_percent       = 0
	Mistsplitter_Def_percent       = 0
	Mistsplitter_Elemental_mastery = 0
	Mistsplitter_Crit_rate         = 0
	Mistsplitter_Crit_damage       = 0.441
	Mistsplitter_Energy_recharge   = 0
)

type Mistsplitter struct {
	*weapon.Weapon
}

func (w Mistsplitter) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	(*cfg)[i].BaseStats[attributes.BaseATK] += Mistsplitter_Base_atk
	(*cfg)[i].BaseStats[attributes.CD] += Mistsplitter_Crit_damage
}

// CheckHasBuffAttack 检查一下普攻给的雾切之巴印还在不在
func CheckHasBuffAttack(cfg *character3.CharWrapper) bool {
	for _, mod := range cfg.Mods {
		if mod.Name == "MistsplitterAttack" {
			return true
		}
	}
	return false
}

func CheckHasBuffBurst(cfg *character3.CharWrapper) bool {
	for _, mod := range cfg.Mods {
		if mod.Name == "MistsplitterBurst" {
			return true
		}
	}
	return false
}

// CheckWeapon 获得12%的全元素伤害加成，并获得「雾切之巴印」的力量。在1/2/3层时，「雾切之巴印」将为角色的元素类型提供8%/16%/28%的元素伤害加成。
// 角色在以下情况中会获得1层「雾切之巴印」：普通攻击造成元素伤害时（持续5秒）、释放元素爆发时（持续10秒）、能量低于100%时（能量充满时消失）。
// 每层效果的持续时间独立计算。
// 前台才能触发
func (w Mistsplitter) CheckWeapon(c character3.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, charIndex int, enemy *enemy.Enemy) {
	// 注意：雾切所有的buff都加在前台角色身上
	attackIndex := len(*attacks) - 1
	if (*attacks)[attackIndex].ActorIndex != charIndex {
		return
	}
	//全元素12%伤害加成，常驻mod
	//删除现有的buffmod
	for i := len((*cfg)[charIndex].Mods) - 1; i >= 0; i-- {
		modDel := (*cfg)[charIndex].Mods[i]
		if modDel.Name == "MistsplitterPermanent" {
			(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[i+1:]...)
		}
	}
	mod := modifier.Mods{
		Name:       "MistsplitterPermanent",
		Modifier:   [attributes.EndStatType]float64{},
		StartFrame: framei,
		Dur:        -1,
		CharIndex:  5,
	}
	permanentBuff := float64(cfg[charIndex].Weapon.Refine)*0.03 + 0.09
	mod.Modifier[attributes.PyroP] = permanentBuff
	mod.Modifier[attributes.HydroP] = permanentBuff
	mod.Modifier[attributes.CryoP] = permanentBuff
	mod.Modifier[attributes.ElectroP] = permanentBuff
	mod.Modifier[attributes.AnemoP] = permanentBuff
	mod.Modifier[attributes.GeoP] = permanentBuff
	mod.Modifier[attributes.DendroP] = permanentBuff
	(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)

	// 判断是不是普攻且带元素伤害
	if strings.Contains((*attacks)[attackIndex].Abil, "Attack") && (*attacks)[attackIndex].Element != attributes.Physical && !CheckHasBuffAttack(&(*cfg)[charIndex]) && framei == (*attacks)[attackIndex].AttackFrame {
		startFrame := c.FindAttackFrame((*attacks)[attackIndex]) + framei
		character3.AddStatus("MistsplitterAttack", 300, startFrame, cfg, charIndex, 0, 1)
	}

	// 判断是不是释放元素爆发时
	if strings.Contains((*attacks)[attackIndex].Abil, "Burst") && !CheckHasBuffBurst(&(*cfg)[charIndex]) {
		startFrame := c.FindAttackFrame((*attacks)[attackIndex])
		character3.AddStatus("MistsplitterBurst", 600, startFrame, cfg, charIndex, 0, 1)
	}

	count := 0
	for _, statusMod := range cfg[charIndex].StatusMods {
		if framei >= statusMod.StartFrame && framei < statusMod.StartFrame+statusMod.Dur {
			count++
		}
	}
	if cfg[charIndex].Energy < cfg[charIndex].EnergyMax && !(strings.Contains((*attacks)[attackIndex].Abil, "Burst") && (framei == (*attacks)[attackIndex].Startframe)) {
		count++
	}

	//fmt.Println("count ", count, "len StatusMods ", len(cfg[charIndex].StatusMods))
	var modBuff modifier.Mods
	var emptyMods modifier.Mods
	//删除现有的buffmod
	for i := len((*cfg)[charIndex].Mods) - 1; i >= 0; i-- {
		modDel := (*cfg)[charIndex].Mods[i]
		if modDel.Name == "MistsplitterBuff" {
			(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[i+1:]...)
		}
	}
	if count == 1 {
		modBuff = modifier.Mods{
			Name:       "MistsplitterBuff",
			Modifier:   [attributes.EndStatType]float64{},
			StartFrame: framei,
			Dur:        -1,
			CharIndex:  5,
		}
		buff1 := float64(cfg[charIndex].Weapon.Refine)*0.02 + 0.06
		modBuff.Modifier[attributes.EleToDmgP(cfg[charIndex].Base.Element)] += buff1
	} else if count == 2 {
		modBuff = modifier.Mods{
			Name:       "MistsplitterBuff",
			Modifier:   [attributes.EndStatType]float64{},
			StartFrame: framei,
			Dur:        -1,
			CharIndex:  5,
		}
		buff2 := float64(cfg[charIndex].Weapon.Refine)*0.04 + 0.12
		modBuff.Modifier[attributes.EleToDmgP(cfg[charIndex].Base.Element)] += buff2
	} else if count == 3 {
		modBuff = modifier.Mods{
			Name:       "MistsplitterBuff",
			Modifier:   [attributes.EndStatType]float64{},
			StartFrame: framei,
			Dur:        -1,
			CharIndex:  5,
		}
		buff3 := float64(cfg[charIndex].Weapon.Refine)*0.07 + 0.21
		modBuff.Modifier[attributes.EleToDmgP(cfg[charIndex].Base.Element)] += buff3
	}
	if modBuff != emptyMods {
		(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, modBuff)
	}
	//if framei == 424 {
	//	fmt.Println(len((*cfg)[charIndex].Mods))
	//	fmt.Println((*cfg)[charIndex].Mods)
	//}

}
