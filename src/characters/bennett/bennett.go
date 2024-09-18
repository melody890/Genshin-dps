package bennett

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/character"
	"dps/src/core/combat"
	character3 "dps/src/core/player/character"
	tmpl "dps/src/template/character"
	"fmt"
	"os"
)

const (
	Bennett_hp           = 12397.0
	Bennett_atk          = 191.0
	Bennett_def          = 771.0
	Bennett_breakthrough = 0.267
	Bennett_element      = attributes.Pyro
	EnergyMax            = 60
)

type Bennett struct {
	*tmpl.Character
}

func (c Bennett) FindFrame(aipre combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	frame := 0
	//fmt.Println("上一个动作", aipre.Abil)
	//fmt.Println("这个动作", Abil)
	// 如果上一个角色和这一个角色不一样，那么就换人
	if aipre.ActorIndex != ActorIndex {
		if aipre.Abil == "bennettActionSkillPress" {
			frame = skillFrames[0][action.ActionSwap]
		} else if aipre.Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 1 {
			frame = attackFrames[0][action.ActionSwap]
		} else if aipre.Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
			frame = attackFrames[1][action.ActionSwap]
		} else if aipre.Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
			frame = attackFrames[2][action.ActionSwap]
		} else if aipre.Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 4 {
			frame = attackFrames[3][action.ActionSwap]
		} else if aipre.Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 5 {
			frame = attackFrames[4][action.ActionSwap]
		} else if aipre.Abil == "bennettActionBurst" {
			frame = burstFrames[action.ActionSwap]
		}
	} else {
		if aipre.Abil == "bennettActionSkillPress" {
			if Abil == "bennettActionDash" {
				frame = skillFrames[0][action.ActionDash]
			} else if Abil == "bennettActionJump" {
				frame = skillFrames[0][action.ActionJump]
			} else {
				frame = skillFrames[0][action.ActionDelay]
			}
		} else if aipre.Abil == "bennettActionAttack1" {
			if Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
				frame = attackFrames[0][action.ActionAttack]
			} else {
				frame = attackFrames[0][action.ActionDelay]
				print()
			}
		} else if aipre.Abil == "bennettActionAttack2" {
			if Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
				frame = attackFrames[1][action.ActionAttack]
			} else {
				frame = attackFrames[1][action.ActionDelay]
			}
		} else if aipre.Abil == "bennettActionAttack3" {
			if Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 4 {
				frame = attackFrames[2][action.ActionAttack]
			} else {
				frame = attackFrames[2][action.ActionDelay]
			}
		} else if aipre.Abil == "bennettActionAttack4" {
			if Abil == "bennettActionAttack" && cfg[ActorIndex].NormalCounter == 5 {
				frame = attackFrames[3][action.ActionAttack]
			} else {
				frame = attackFrames[3][action.ActionDelay]
			}
		} else if aipre.Abil == "bennettActionAttack5" {
			frame = attackFrames[4][action.ActionAttack]
		} else if aipre.Abil == "bennettActionBurst" {
			if Abil == "bennettActionDash" {
				frame = burstFrames[action.ActionDash]
			} else if Abil == "bennettActionJump" {
				frame = burstFrames[action.ActionJump]
			} else {
				frame = burstFrames[action.ActionDelay]
			}
		}
	}
	//fmt.Println("帧数：", frame)
	return frame
}

func (c Bennett) FindAttackFrame(ai combat.AttackInfo) int {
	frame := 0
	switch ai.Abil {
	case "bennettActionSkillPress":
		frame = skillPressHitmark
	case "bennettActionAttack1":
		frame = attackHitmarks[0]
	case "bennettActionAttack2":
		frame = attackHitmarks[1]
	case "bennettActionAttack3":
		frame = attackHitmarks[2]
	case "bennettActionAttack4":
		frame = attackHitmarks[3]
	case "bennettActionAttack5":
		frame = attackHitmarks[4]
	case "bennettActionBurst":
		frame = burstStartFrame
	case "":
	default:
		fmt.Println(ai.Abil)
		fmt.Println("bennett.go 班尼特找动作出现异常动作名")
		os.Exit(0)
	}
	return frame
}

//// InitChar 初始化 Char 实例并返回 tmpl.Character 实例
//// 里面是角色名字和初始三围
//func (c *Bennett) InitChar(CharName string) *tmpl.Character {
//	c.Character = &tmpl.Character{
//		Name: CharName,
//		Hp:   Bennett_hp,
//		Atk:  Bennett_atk,
//		Def:  Bennett_def,
//	}
//	return c.Character
//}

func (c Bennett) InitEnergy() float64 {
	return EnergyMax
}

func Init_bennett() (character.Base_attribute, string) {
	char := character.Base_attribute{
		Def: Bennett_def,
		Atk: Bennett_atk,
		Hp:  Bennett_hp,
	}
	breakthrough := "Energy_recharge"
	return char, breakthrough
}

func (c *Bennett) FindElement() attributes.Element {
	return attributes.Pyro
}

func SearchCharIndex(cfg character.Charcfg) int {
	//搜索班尼特的index
	CharIndex := -1
	for i, char := range cfg.Characters {
		if char == "bennett" {
			CharIndex = i
			break
		}
	}
	if CharIndex == -1 {
		_ = fmt.Errorf("character %s not found", "bennett")
	}
	return CharIndex
}

func (c *Bennett) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	var state [attributes.EndStatType]float64
	state[attributes.BaseHP] += Bennett_hp
	state[attributes.BaseATK] += Bennett_atk
	state[attributes.BaseDEF] += Bennett_def
	state[attributes.CR] += 0.05
	state[attributes.CD] += 0.5
	state[attributes.MaxEnergy] += 60
	state[attributes.ER] = 1 + Bennett_breakthrough
	(*cfg)[i].BaseStats = state
}

func (c Bennett) NormalHitNum() int {
	return normalHitNum
}
