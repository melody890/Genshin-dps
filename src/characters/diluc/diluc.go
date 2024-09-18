package diluc

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
	Diluc_hp           = 12981.0
	Diluc_atk          = 335.0
	Diluc_def          = 784.0
	Diluc_breakthrough = 0.242
	Diluc_element      = attributes.Pyro
	EnergyMax          = 40
)

type Diluc struct {
	*tmpl.Character
}

// func (c Diluc) InitChar(CharName) *tmpl.Character
func (c Diluc) InitEnergy() float64 {
	return 40
}

func (c Diluc) NormalHitNum() int {
	return normalHitNum
}

func (c *Diluc) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	var state [attributes.EndStatType]float64
	state[attributes.BaseHP] += Diluc_hp
	state[attributes.BaseATK] += Diluc_atk
	state[attributes.BaseDEF] += Diluc_def
	state[attributes.CR] += 0.05
	state[attributes.CD] += 0.5
	state[attributes.MaxEnergy] += 40
	state[attributes.CR] += Diluc_breakthrough
	(*cfg)[i].BaseStats = state
}

func (c Diluc) FindFrame(aipre combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	frame := 0
	//fmt.Println("上一个动作", aipre.Abil)
	//fmt.Println("这个动作", Abil)
	// 如果上一个角色和这一个角色不一样，那么就换人
	if aipre.ActorIndex != ActorIndex {
		if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 1 {
			frame = skillFrames[0][action.ActionSwap]
		} else if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 2 {
			frame = skillFrames[1][action.ActionSwap]
		} else if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 3 {
			frame = skillFrames[2][action.ActionSwap]
		} else if aipre.Abil == "dilucActionAttack" && cfg[ActorIndex].NormalCounter == 0 {
			frame = attackFrames[0][action.ActionSwap]
		} else if aipre.Abil == "dilucActionAttack" && cfg[ActorIndex].NormalCounter == 1 {
			frame = attackFrames[1][action.ActionSwap]
		} else if aipre.Abil == "dilucActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
			frame = attackFrames[2][action.ActionSwap]
		} else if aipre.Abil == "dilucActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
			frame = attackFrames[3][action.ActionSwap]
		} else if aipre.Abil == "dilucActionBurst" {
			frame = burstFrames[action.ActionSwap]
		}
	} else {
		if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 1 || aipre.Abil == "dilucActionSkill1" {
			if Abil == "dilucActionDash" {
				frame = skillFrames[0][action.ActionDash]
			} else if Abil == "dilucActionJump" {
				frame = skillFrames[0][action.ActionJump]
			} else if Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 2 || aipre.Abil == "dilucActionSkill2" {
				frame = skillFrames[0][action.ActionSkill]
			} else {
				frame = skillFrames[0][action.ActionDelay]
			}
		} else if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 2 || aipre.Abil == "dilucActionSkill2" {
			if Abil == "dilucActionDash" {
				frame = skillFrames[1][action.ActionDash]
			} else if Abil == "dilucActionJump" {
				frame = skillFrames[1][action.ActionJump]
			} else if Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 3 || aipre.Abil == "dilucActionSkill3" {
				frame = skillFrames[1][action.ActionSkill]
			} else if Abil == "dilucActionBurst" {
				frame = skillFrames[1][action.ActionBurst]
			} else {
				frame = skillFrames[1][action.ActionDelay]
			}
		} else if aipre.Abil == "dilucActionSkill" && cfg[ActorIndex].SkillCounter == 3 || aipre.Abil == "dilucActionSkill3" {
			if Abil == "dilucActionDash" {
				frame = skillFrames[2][action.ActionDash]
			} else if Abil == "dilucActionJump" {
				frame = skillFrames[2][action.ActionJump]
			} else if Abil == "dilucActionBurst" {
				frame = skillFrames[2][action.ActionBurst]
			} else if Abil == "dilucActionAttack" {
				frame = skillFrames[2][action.ActionSkill]
			} else {
				frame = skillFrames[2][action.ActionDelay]
			}
		} else if aipre.Abil == "dilucActionAttack1" {
			frame = attackFrames[0][action.ActionDelay]
		} else if aipre.Abil == "dilucActionAttack2" {
			frame = attackFrames[1][action.ActionDelay]
		} else if aipre.Abil == "dilucActionAttack3" {
			frame = attackFrames[2][action.ActionDelay]
		} else if aipre.Abil == "dilucActionAttack4" {
			frame = attackFrames[3][action.ActionDelay]
		} else if aipre.Abil == "dilucActionBurst" {
			if Abil == "dilucActionDash" {
				frame = burstFrames[action.ActionDash]
			} else if Abil == "dilucActionJump" {
				frame = burstFrames[action.ActionJump]
			} else {
				frame = burstFrames[action.ActionDelay]
			}
		}
	}
	//fmt.Println("帧数：", frame)
	return frame
}

//func (c Diluc) FindEndFrame(aepre combat.AttackEvent) int {
//	frame := 0
//	if aepre.Info.Abil == "skill press" {
//		frame = skillFrames[0][action.InvalidAction]
//	} else if aepre.Info.Abil == "attack1" {
//		frame = attackFrames[0][action.InvalidAction]
//	} else if aepre.Info.Abil == "burst" {
//		frame = burstFrames[action.InvalidAction]
//	} else {
//		log.Println("ERROR！bennet.go寻找动作帧数时动作名称错误")
//		os.Exit(0)
//	}
//	return frame
//}

func (c Diluc) FindAttackFrame(ai combat.AttackInfo) int {
	frame := 0
	switch ai.Abil {
	case "dilucActionSkill1":
		frame = skillHitmarks[0]
	case "dilucActionSkill2":
		frame = skillHitmarks[1]
	case "dilucActionSkill3":
		frame = skillHitmarks[2]
	case "dilucActionAttack1":
		frame = attackHitmarks[0]
	case "dilucActionAttack2":
		frame = attackHitmarks[1]
	case "dilucActionAttack3":
		frame = attackHitmarks[2]
	case "dilucActionAttack4":
		frame = attackHitmarks[3]
	case "dilucActionBurst":
		frame = burstHitmark
	case "":
	default:
		fmt.Println(ai.Abil)
		fmt.Println("diluc.go 迪卢克找动作出现异常动作名")
		os.Exit(0)
	}
	return frame
}

func (c Diluc) FindElement() attributes.Element {
	return attributes.Pyro
}
