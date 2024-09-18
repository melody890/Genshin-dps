package diluc

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"slices"
)

var (
	attackFrames          [][]int
	attackHitmarks        = []int{24, 39, 26, 49}
	attackPoiseDMG        = []float64{108.1, 105.57, 119.03, 161.46}
	attackHitlagHaltFrame = []float64{.1, .09, .09, .12}
	attackHitboxes        = [][]float64{{2}, {2, 3}, {2}, {2, 3}}
	attackOffsets         = []float64{0.5, -1, 0.5, -0.5}
	attackFanAngles       = []float64{300, 360, 300, 360}
)

const normalHitNum = 4

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0], 32)
	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1], 46)
	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2], 34)
	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3], 99)
}

func (c Diluc) Attack(cfg *[4]character.CharWrapper, charindex int, startFrame int, enemy *enemy.Enemy) []combat.AttackInfo {
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
	if len((*cfg)[charindex].Infusion) > 0 {
		for i := range (*cfg)[charindex].Infusion {
			if slices.Contains((*cfg)[charindex].Infusion[i].Class, "claymore") && slices.Contains((*cfg)[charindex].Infusion[i].Tags, attacks.AttackTagNormal) && (*cfg)[charindex].Infusion[i].StartFrame < startFrame && (*cfg)[charindex].Infusion[i].Dur+(*cfg)[charindex].Infusion[i].StartFrame >= startFrame && (charindex == (*cfg)[charindex].Infusion[i].CharIndex || (*cfg)[charindex].Infusion[i].CharIndex == 5 || (*cfg)[charindex].Infusion[i].CharIndex == 6) {
				ai.Element = (*cfg)[charindex].Infusion[i].Ele
			}
		}
	}
	switch cfg[charindex].NormalCounter {
	case 1:
		ai.Abil = "dilucActionAttack1"
	case 2:
		ai.Abil = "dilucActionAttack2"
	case 3:
		ai.Abil = "dilucActionAttack3"
	case 4:
		ai.Abil = "dilucActionAttack4"
	case 5:
		ai.Abil = "dilucActionAttack5"
	default:
		ai.Abil = "dilucActionAttack"
	}
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Diluc) AttackMod(startframe int, cfg *[4]character.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
	c1(cfg, enemy, startframe, charIndex)
}

func (c Diluc) AttackElementAttach(enemy *enemy.Enemy) bool {
	var elementTrigger bool
	var exist bool
	for i := range enemy.ElementAttach.ActionName {
		if enemy.ElementAttach.ActionName[i] == "dilucActionAttack" {
			exist = true
			(*enemy).ElementAttach.Counter[i]++
			if (*enemy).ElementAttach.Counter[i]%3 == 1 {
				elementTrigger = true
			}
		}
	}
	if exist == false {
		(*enemy).ElementAttach.ActionName = append((*enemy).ElementAttach.ActionName, "dilucActionAttack")
		(*enemy).ElementAttach.Counter = append((*enemy).ElementAttach.Counter, 1)
		(*enemy).ElementAttach.Timer = append((*enemy).ElementAttach.Timer, 0)
		elementTrigger = true
	}
	return elementTrigger
}
