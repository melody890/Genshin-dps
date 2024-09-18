package kaeya

import (
	"dps/src/core/action"
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
	"slices"
)

var (
	attackFrames          [][]int
	attackHitmarks        = []int{14, 9, 14, 23, 30}
	attackHitlagHaltFrame = []float64{.03, .03, .06, .06, 0.1}
	attackHitboxes        = [][]float64{{1.7}, {1.5}, {1, 2.6}, {1, 3.5}, {1.8}}
	attackOffsets         = []float64{0.8, 1.2, -0.2, 0.3, 0.5}
)

const normalHitNum = 5

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0], 30) // N1 -> CA
	attackFrames[0][action.ActionAttack] = 21                             // N1 -> N2

	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1], 25) // N2 -> CA
	attackFrames[1][action.ActionAttack] = 21                             // N2 -> N3

	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2], 47) // N3 -> CA
	attackFrames[2][action.ActionAttack] = 39                             // N3 -> N4

	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3], 46) // N4 -> CA
	attackFrames[3][action.ActionAttack] = 38                             // N4 -> N5

	attackFrames[4] = frames.InitNormalCancelSlice(attackHitmarks[4], 64) // N5 -> N1
	attackFrames[4][action.ActionCharge] = 500                            // N5 -> CA, TODO: this action is illegal; need better way to handle it
}

func (c Kaeya) Attack(cfg *[4]character.CharWrapper, charindex int, startFrame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Element:         attributes.Physical,
		ElementQuantity: 1,
		Mult:            auto[cfg[charindex].NormalCounter-1][cfg[charindex].Talents.Attack-1],
		MultType:        []string{"atk"},
	}
	// 看附魔
	if len((*cfg)[charindex].Infusion) > 0 {
		for i := range (*cfg)[charindex].Infusion {
			if slices.Contains((*cfg)[charindex].Infusion[i].Class, "sword") && slices.Contains((*cfg)[charindex].Infusion[i].Tags, attacks.AttackTagNormal) && (*cfg)[charindex].Infusion[i].StartFrame < startFrame && (*cfg)[charindex].Infusion[i].Dur+(*cfg)[charindex].Infusion[i].StartFrame >= startFrame && (charindex == (*cfg)[charindex].Infusion[i].CharIndex || (*cfg)[charindex].Infusion[i].CharIndex == 5 || (*cfg)[charindex].Infusion[i].CharIndex == 6) {
				ai.Element = (*cfg)[charindex].Infusion[i].Ele
			}
		}
	}
	switch cfg[charindex].NormalCounter {
	case 1:
		ai.Abil = "kaeyaActionAttack1"
	case 2:
		ai.Abil = "kaeyaActionAttack2"
	case 3:
		ai.Abil = "kaeyaActionAttack3"
	case 4:
		ai.Abil = "kaeyaActionAttack4"
	case 5:
		ai.Abil = "kaeyaActionAttack5"
	default:
		ai.Abil = "kaeyaActionAttack"
	}
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Kaeya) AttackMod(startframe int, cfg *[4]character.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
	if cfg[charIndex].Base.Cons >= 1 {
		// 一命：受到冰元素影响的敌人，凯亚的普通攻击与重击暴击率提升15%
		//先删除已有的buff
		for j := len((*cfg)[charIndex].Mods) - 1; j >= 0; j-- {
			if (*cfg)[charIndex].Mods[j].Name == "kaeyaC1" {
				(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:j], (*cfg)[charIndex].Mods[j+1:]...)
			}
		}
		if enemy.Element == attributes.Cryo {
			mod := modifier.Mods{
				Name:       "kaeyaC1",
				StartFrame: startframe,
				Dur:        -1,
				CharIndex:  charIndex,
				Condition:  attacks.AttackTagNormal,
			}
			mod.Modifier[attributes.CR] += 0.15
			(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
		}
	}
}

func (c Kaeya) AttackElementAttach(enemy *enemy.Enemy) bool {
	var elementTrigger bool
	var exist bool
	for i := range enemy.ElementAttach.ActionName {
		if enemy.ElementAttach.ActionName[i] == "kaeyaActionAttack" {
			exist = true
			(*enemy).ElementAttach.Counter[i]++
			if (*enemy).ElementAttach.Counter[i]%3 == 1 {
				elementTrigger = true
			}
		}
	}
	if exist == false {
		(*enemy).ElementAttach.ActionName = append((*enemy).ElementAttach.ActionName, "kaeyaActionAttack")
		(*enemy).ElementAttach.Counter = append((*enemy).ElementAttach.Counter, 1)
		(*enemy).ElementAttach.Timer = append((*enemy).ElementAttach.Timer, 0)
		elementTrigger = true
	}
	return elementTrigger
}
