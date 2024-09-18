package bennett

import (
	"dps/src/core/action"
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"slices"
)

var (
	attackFrames           [][]int
	attackHitmarks         = []int{13, 9, 13, 25, 24}
	attackHitlagHaltFrames = []float64{0.03, 0.03, 0.06, 0.09, 0.12}
	attackHitboxes         = [][]float64{{1.2}, {1.2}, {2}, {1, 3.5}, {2}}
	attackOffsets          = []float64{0.8, 0.8, 0.6, 0.3, 0.8}
	attackFanAngles        = []float64{360, 360, 30, 360, 360}
)

const normalHitNum = 5

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0], 33)
	attackFrames[0][action.ActionAttack] = 20

	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1], 27)
	attackFrames[1][action.ActionAttack] = 17

	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2], 46)
	attackFrames[2][action.ActionAttack] = 37

	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3], 48)
	attackFrames[3][action.ActionAttack] = 44

	attackFrames[4] = frames.InitNormalCancelSlice(attackHitmarks[4], 60)
	attackFrames[4][action.ActionCharge] = 500 //TODO: this action is illegal; need better way to handle it
}

func (c Bennett) Attack(cfg *[4]character.CharWrapper, charindex int, startFrame int, enemy *enemy.Enemy) []combat.AttackInfo {
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
		Mult:            attack[cfg[charindex].NormalCounter-1][cfg[charindex].Talents.Attack-1],
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
		ai.Abil = "bennettActionAttack1"
	case 2:
		ai.Abil = "bennettActionAttack2"
	case 3:
		ai.Abil = "bennettActionAttack3"
	case 4:
		ai.Abil = "bennettActionAttack4"
	case 5:
		ai.Abil = "bennettActionAttack5"
	default:
		ai.Abil = "bennettActionAttack"
	}

	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Bennett) AttackMod(startframe int, cfg *[4]character.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
}

func (c Bennett) AttackElementAttach(enemy *enemy.Enemy) bool {
	var elementTrigger bool
	var exist bool
	for i := range enemy.ElementAttach.ActionName {
		if enemy.ElementAttach.ActionName[i] == "bennettActionAttack" {
			exist = true
			(*enemy).ElementAttach.Counter[i]++
			if (*enemy).ElementAttach.Counter[i]%3 == 1 {
				elementTrigger = true
			}
		}
	}
	if exist == false {
		(*enemy).ElementAttach.ActionName = append((*enemy).ElementAttach.ActionName, "bennettActionAttack")
		(*enemy).ElementAttach.Counter = append((*enemy).ElementAttach.Counter, 1)
		(*enemy).ElementAttach.Timer = append((*enemy).ElementAttach.Timer, 0)
		elementTrigger = true
	}
	return elementTrigger
}
