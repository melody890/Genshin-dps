package sword

import (
	attacks2 "dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/reactions"
	"dps/src/template/modifier"
	"dps/src/template/weapon"
)

const (
	Freedom_Base_atk          = 608
	Freedom_Hp_percent        = 0
	Freedom_Atk_percent       = 0
	Freedom_Def_percent       = 0
	Freedom_Elemental_mastery = 198
	Freedom_Crit_rate         = 0
	Freedom_Crit_damage       = 0
	Freedom_Energy_recharge   = 0
)

type Freedom struct {
	*weapon.Weapon
}

func (w Freedom) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	(*cfg)[i].BaseStats[attributes.BaseATK] += Freedom_Base_atk
	(*cfg)[i].BaseStats[attributes.EM] += Freedom_Elemental_mastery
}

func (w Freedom) CheckWeapon(c character3.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, charIndex int, enemy *enemy.Enemy) {
	//飘游风中的【千年的大乐章】的一部分。造成的伤害提高（10%/12.5%/15%/17.5%/20%）；触发元素反应时，角色获得一枚奋起之符，每0.5秒最多触发一次，角色处于队伍后台也可以触发。拥有2枚奋起之符时，将消耗所有奋起之符，使附近的队伍中所有角色获得持续12秒的【千年的大乐章·抗争之歌】效果：普通攻击、重击、下落攻击造成的伤害提高（16%/20%/24%/28%/32%），攻击力提升（20%/25%/30%/35%/40%）。触发后20秒内，无法再次获得奋起之符。【千年的大乐章】触发的多种数值效果中，同类数值效果不可叠加。
	//if framei == 140 {
	//	fmt.Println("debug")
	//}
	var atkFrame int
	var lenAttacks int
	for i := len(*attacks) - 1; i >= 0; i-- {
		if i == len(*attacks)-1 {
			atkFrame = (*attacks)[i].AttackFrame
			lenAttacks++
		} else {
			if (*attacks)[i].AttackFrame != atkFrame {
				break
			} else {
				lenAttacks++
			}
		}
	}
	for attackIndex := len(*attacks) - 1; attackIndex >= len(*attacks)-lenAttacks; attackIndex-- {
		if (*attacks)[attackIndex].ActorIndex != charIndex {
			return
		}
		for i := len((*cfg)[charIndex].Mods) - 1; i >= 0; i-- {
			modDel := (*cfg)[charIndex].Mods[i]
			if modDel.Name == "FreedomPermanent" {
				(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[i+1:]...)
			}
		}
		mod := modifier.Mods{
			Name:       "FreedomPermanent",
			StartFrame: framei,
			Dur:        -1,
			CharIndex:  charIndex,
		}
		mod.Modifier[attributes.DmgP] = 0.075 + float64(cfg[charIndex].Weapon.Refine)*0.025
		(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
		//如果这个动作是持有苍古的人触发的，而且发生了元素反应
		if (*attacks)[attackIndex].ActorIndex == charIndex && reactions.IsElementReaction((*attacks)[attackIndex], enemy) && (*attacks)[attackIndex].AttackFrame == framei {
			if character3.CDReady("FreedomBuff", cfg, charIndex, framei) {
				character3.AddStatus("FreedomBuff", -1, framei, cfg, charIndex, 30, -1)
			}
		}
		count := 0
		for _, statusMod := range cfg[charIndex].StatusMods {
			if statusMod.ModName == "FreedomBuff" {
				count++
			}
		}
		if count == 2 {
			modSymphony1 := modifier.Mods{
				Name:       "FreedomSymphonyATK",
				StartFrame: framei,
				Dur:        12 * 60,
				CharIndex:  6,
			}
			modSymphony1.Modifier[attributes.ATKP] = 0.15 + float64(cfg[charIndex].Weapon.Refine)*0.05
			modSymphony2 := modifier.Mods{
				Name:       "FreedomSymphonyDMG",
				StartFrame: framei,
				Dur:        12 * 60,
				CharIndex:  6,
				Condition:  attacks2.AttackTagNormal,
			}
			modSymphony2.Modifier[attributes.DmgP] = 0.12 + float64(cfg[charIndex].Weapon.Refine)*0.04
			modSymphony3 := modifier.Mods{
				Name:       "FreedomSymphonyDMG",
				StartFrame: framei,
				Dur:        12 * 60,
				CharIndex:  6,
				Condition:  attacks2.AttackTagPlunge,
			}
			modSymphony3.Modifier[attributes.DmgP] = 0.12 + float64(cfg[charIndex].Weapon.Refine)*0.04
			modSymphony4 := modifier.Mods{
				Name:       "FreedomSymphonyDMG",
				StartFrame: framei,
				Dur:        12 * 60,
				CharIndex:  6,
				Condition:  attacks2.AttackCharge,
			}
			modSymphony4.Modifier[attributes.DmgP] = 0.12 + float64(cfg[charIndex].Weapon.Refine)*0.04
			//给队伍角色加上千年大乐章效果
			for i := range cfg {
				for j := len((*cfg)[i].Mods) - 1; j >= 0; j-- {
					modDel := (*cfg)[i].Mods[j]
					if modDel.Name == "FreedomSymphonyDMG" {
						(*cfg)[i].Mods = append((*cfg)[i].Mods[:j], (*cfg)[i].Mods[j+1:]...)
					}
				}
				for j := len((*cfg)[i].Mods) - 1; j >= 0; j-- {
					modDel := (*cfg)[i].Mods[j]
					if modDel.Name == "FreedomSymphonyATK" {
						(*cfg)[i].Mods = append((*cfg)[i].Mods[:j], (*cfg)[i].Mods[j+1:]...)
					}
				}
				(*cfg)[i].Mods = append((*cfg)[i].Mods, modSymphony1)
				(*cfg)[i].Mods = append((*cfg)[i].Mods, modSymphony2)
				(*cfg)[i].Mods = append((*cfg)[i].Mods, modSymphony3)
				(*cfg)[i].Mods = append((*cfg)[i].Mods, modSymphony4)
			}
			//把现有的奋起之符撤下来，然后塞一个CD20的当缓冲
			character3.RemoveStatus("FreedomBuff", cfg, charIndex)
			character3.AddStatus("FreedomBuff", 20*60, framei, cfg, charIndex, 0, 1)
		}
	}
}
