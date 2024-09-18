package kazuha

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	character3 "dps/src/core/player/character"
	tmpl "dps/src/template/character"
	"fmt"
	"os"
	"strings"
)

const (
	Kazuha_hp           = 13348.0
	Kazuha_atk          = 297.0
	Kazuha_def          = 807.0
	Kazuha_breakthrough = 115.0
	Kazuha_element      = attributes.Anemo
	EnergyMax           = 60
)

type Kazuha struct {
	*tmpl.Character
}

func (c Kazuha) FindFrame(aipre combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	frame := 0
	//if aipre.Abil == "kazuhaActionPlungePress" {
	//	fmt.Println("debug")
	//}
	//fmt.Println("上一个动作", aipre.Abil)
	//fmt.Println("这个动作", Abil)
	// 如果上一个角色和这一个角色不一样，那么就换人
	if aipre.ActorIndex != ActorIndex {
		if aipre.Abil == "kazuhaActionSkillPress" {
			frame = skillPressFrames[0][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionSkillHold" {
			frame = skillHoldFrames[0][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionAttack" && cfg[ActorIndex].NormalCounter == 1 {
			frame = attackFrames[0][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionAttack" && cfg[ActorIndex].NormalCounter == 2 {
			frame = attackFrames[1][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionAttack" && cfg[ActorIndex].NormalCounter == 3 {
			frame = attackFrames[2][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionAttack" && cfg[ActorIndex].NormalCounter == 4 {
			frame = attackFrames[3][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionAttack" && cfg[ActorIndex].NormalCounter == 5 {
			frame = attackFrames[4][action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionBurst" {
			frame = burstFrames[action.ActionSwap]
		} else if aipre.Abil == "kazuhaActionPlungeHold" {
			frame = plungeHoldFrames[action.ActionSwap]
		}
	} else {
		if aipre.Abil == "kazuhaActionSkillPress" {
			if Abil == "kazuhaActionPlungePress" {
				frame = skillPressFrames[0][action.ActionHighPlunge]
			} else {
				frame = skillPressFrames[0][action.ActionDelay]
			}
		} else if aipre.Abil == "kazuhaActionSkillHold" {
			if Abil == "kazuhaActionPlungeHold" {
				frame = skillHoldFrames[0][action.ActionHighPlunge]
			} else {
				frame = skillHoldFrames[0][action.ActionDelay]
			}
		} else if aipre.Abil == "kazuhaActionAttack1" {
			if Abil == "kazuhaActionharge" {
				frame = attackFrames[0][action.ActionCharge]
			} else {
				frame = attackFrames[0][action.ActionDelay]
			}
		} else if aipre.Abil == "kazuhaActionAttack2" {
			frame = attackFrames[1][action.ActionDelay]
		} else if strings.Contains(aipre.Abil, "kazuhaActionAttack3") {
			if Abil == "kazuhaActionharge" {
				frame = attackFrames[2][action.ActionCharge] - attackHitmarks[2][1]
			} else {
				frame = attackFrames[2][action.ActionDelay] - attackHitmarks[2][1]
			}
		} else if aipre.Abil == "kazuhaActionAttack4" {
			if Abil == "kazuhaActionharge" {
				frame = attackFrames[3][action.ActionCharge]
			} else {
				frame = attackFrames[3][action.ActionDelay]
			}
		} else if strings.Contains(aipre.Abil, "kazuhaActionAttack5") {
			frame = attackFrames[4][action.ActionDelay] - attackHitmarks[4][2]
		} else if aipre.Abil == "kazuhaActionBurst" {
			if Abil == "kazuhaActionDash" {
				frame = burstFrames[action.ActionDash]
			} else if Abil == "kazuhaActionSkillPress" || Abil == "kazuhaActionSkillHold" {
				frame = burstFrames[action.ActionSkill]
			} else if Abil == "kazuhaActionAttack" {
				frame = burstFrames[action.ActionAttack]
			} else {
				frame = burstFrames[action.ActionDelay]
			}
		} else if aipre.Abil == "kazuhaActionPlungeHold" {
			if Abil == "kazuhaActionDash" {
				frame = plungeHoldFrames[action.ActionDash]
			} else if Abil == "kazuhaActionSkillPress" || Abil == "kazuhaActionSkillHold" {
				frame = plungeHoldFrames[action.ActionSkill]
			} else if Abil == "kazuhaActionBurst" {
				frame = plungeHoldFrames[action.ActionBurst]
			} else if Abil == "kazuhaActionJump" {
				frame = plungeHoldFrames[action.ActionJump]
			} else {
				frame = plungeHoldFrames[action.ActionDelay]
			}
		} else if aipre.Abil == "kazuhaActionPlungePress" {
			if Abil == "kazuhaActionDash" {
				frame = plungePressFrames[action.ActionDash]
			} else if Abil == "kazuhaActionJump" {
				frame = plungePressFrames[action.ActionJump]
			} else {
				frame = plungePressFrames[action.ActionDelay]
			}
		}
	}
	return frame
}

func (c Kazuha) FindAttackFrame(ai combat.AttackInfo) int {
	frame := 0
	switch ai.Abil {
	case "kazuhaActionSkillPress":
		frame = skillPressHitmark
	case "kazuhaActionSkillHold":
		frame = skillHoldHitmark
	case "kazuhaActionAttack1":
		frame = attackHitmarks[0][0]
	case "kazuhaActionAttack2":
		frame = attackHitmarks[1][0]
	case "kazuhaActionAttack31":
		frame = attackHitmarks[2][0]
	case "kazuhaActionAttack32":
		frame = attackHitmarks[2][1]
	case "kazuhaActionAttack4":
		frame = attackHitmarks[3][0]
	case "kazuhaActionAttack51":
		frame = attackHitmarks[4][0]
	case "kazuhaActionAttack52":
		frame = attackHitmarks[4][1]
	case "kazuhaActionAttack53":
		frame = attackHitmarks[4][2]
	case "kazuhaActionBurst":
		frame = burstHitmark
	case "kazuhaActionPlungeHold":
		frame = plungeHoldHitmark
	case "kazuhaActionPlungePress":
		frame = plungePressHitmark
	default:
		fmt.Println(ai.Abil)
		fmt.Println("kazuha.go 万叶找动作出现异常动作名")
		os.Exit(0)
	}
	return frame
}

func (c Kazuha) FindElement() attributes.Element {
	return attributes.Anemo
}

func (c Kazuha) InitEnergy() float64 {
	return EnergyMax
}

func (c Kazuha) NormalHitNum() int {
	return normalHitNum
}

func (c Kazuha) UpdateStates(cfg *[4]character3.CharWrapper, i int) {
	var state [attributes.EndStatType]float64
	state[attributes.BaseHP] += Kazuha_hp
	state[attributes.BaseATK] += Kazuha_atk
	state[attributes.BaseDEF] += Kazuha_def
	state[attributes.CR] += 0.05
	state[attributes.CD] += 0.5
	state[attributes.MaxEnergy] += 60
	state[attributes.EM] += Kazuha_breakthrough
	(*cfg)[i].BaseStats = state
}

//TODO:还有万叶的突破没写，那个扩散增伤
