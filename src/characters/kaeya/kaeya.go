package kaeya

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	character3 "dps/src/core/player/character"
	tmpl "dps/src/template/character"
	"fmt"
	"os"
)

const (
	Kaeya_hp           = 11636.0
	Kaeya_atk          = 223.0
	Kaeya_def          = 792.0
	Kaeya_breakthrough = 0.267
	Kaeya_element      = attributes.Cryo
	EnergyMax          = 60
)

type Kaeya struct {
	*tmpl.Character
}

func (c Kaeya) FindElement() attributes.Element {
	return Kaeya_element
}

func (c Kaeya) InitEnergy() float64 {
	return EnergyMax
}

func (c Kaeya) NormalHitNum() int {
	return normalHitNum
}

func (c Kaeya) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	var state [attributes.EndStatType]float64
	state[attributes.BaseHP] += Kaeya_hp
	state[attributes.BaseATK] += Kaeya_atk
	state[attributes.BaseDEF] += Kaeya_def
	state[attributes.CR] += 0.05
	state[attributes.CD] += 0.5
	state[attributes.MaxEnergy] += 60
	state[attributes.ER] = 1 + Kaeya_breakthrough
	(*cfg)[i].BaseStats = state
}

func (c Kaeya) FindFrame(aipre combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	frame := 0
	//fmt.Println("上一个动作", aipre.Abil)
	//fmt.Println("这个动作", Abil)
	// 如果上一个角色和这一个角色不一样，那么就换人
	if aipre.ActorIndex != ActorIndex {
		if aipre.Abil == "kaeyaActionSkill" {
			frame = skillFrames[action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 1 {
			frame = attackFrames[0][action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
			frame = attackFrames[1][action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
			frame = attackFrames[2][action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 4 {
			frame = attackFrames[3][action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 5 {
			frame = attackFrames[4][action.ActionSwap]
		} else if aipre.Abil == "kaeyaActionBurst" {
			frame = burstFrames[action.ActionSwap]
		}
	} else {
		if aipre.Abil == "kaeyaActionSkill" {
			if Abil == "kaeyaActionDash" {
				frame = skillFrames[action.ActionDash]
			} else if Abil == "kaeyaActionJump" {
				frame = skillFrames[action.ActionJump]
			} else if Abil == "kaeyaActionBurst" {
				frame = skillFrames[action.ActionBurst]
			} else {
				frame = skillFrames[action.ActionDelay]
			}
		} else if aipre.Abil == "kaeyaActionAttack1" {
			if Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
				frame = attackFrames[0][action.ActionAttack]
			} else {
				frame = attackFrames[0][action.ActionDelay]
				print()
			}
		} else if aipre.Abil == "kaeyaActionAttack2" {
			if Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
				frame = attackFrames[1][action.ActionAttack]
			} else {
				frame = attackFrames[1][action.ActionDelay]
			}
		} else if aipre.Abil == "kaeyaActionAttack3" {
			if Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 4 {
				frame = attackFrames[2][action.ActionAttack]
			} else {
				frame = attackFrames[2][action.ActionDelay]
			}
		} else if aipre.Abil == "kaeyaActionAttack4" {
			if Abil == "kaeyaActionAttack" && cfg[ActorIndex].NormalCounter == 5 {
				frame = attackFrames[3][action.ActionAttack]
			} else {
				frame = attackFrames[3][action.ActionDelay]
			}
		} else if aipre.Abil == "kaeyaActionAttack5" {
			frame = attackFrames[4][action.ActionDelay]
		} else if aipre.Abil == "kaeyaActionBurst" {
			if Abil == "kaeyaActionDash" {
				frame = burstFrames[action.ActionDash]
			} else if Abil == "kaeyaActionJump" {
				frame = burstFrames[action.ActionJump]
			} else if Abil == "kaeyaActionAttack" {
				frame = burstFrames[action.ActionAttack]
			} else {
				frame = burstFrames[action.ActionDelay]
			}
		}
	}
	return frame
}

func (c Kaeya) FindAttackFrame(ai combat.AttackInfo) int {
	frame := 0
	switch ai.Abil {
	case "kaeyaActionSkill":
		frame = skillHitmark
	case "kaeyaActionAttack1":
		frame = attackHitmarks[0]
	case "kaeyaActionAttack2":
		frame = attackHitmarks[1]
	case "kaeyaActionAttack3":
		frame = attackHitmarks[2]
	case "kaeyaActionAttack4":
		frame = attackHitmarks[3]
	case "kaeyaActionAttack5":
		frame = attackHitmarks[4]
	case "kaeyaActionBurst":
		frame = burstHitmark
	case "":
	default:
		fmt.Println(ai.Abil)
		fmt.Println("kaeya.go 凯亚找动作出现异常动作名")
		os.Exit(0)
	}
	return frame
}
