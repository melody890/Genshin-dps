package kazuha

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
)

var (
	attackFrames          [][]int
	attackHitmarks        = [][]int{{13}, {11}, {16, 25}, {15}, {15, 19, 27}}
	attackHitlagHaltFrame = [][]float64{{.03}, {.03}, {.01, .05}, {.06}, {0, 0, 0}}
	attackHitlagFactor    = [][]float64{{.01}, {.01}, {.01, .01}, {.01}, {.05, .05, .05}}
	attackDefHalt         = [][]bool{{true}, {true}, {false, true}, {true}, {true, false, true}}
	attackRadius          = []float64{1.5, 1.5, 1.5, 1.5, 2.2}
	attackOffsets         = [][]float64{{1}, {0.7}, {0.8, 0.7}, {1}, {0, 0, 0}}
)

const normalHitNum = 5

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0][0], 21)
	attackFrames[0][action.ActionCharge] = 21

	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1][0], 20)

	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2][1], 33)
	attackFrames[2][action.ActionCharge] = 31

	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3][0], 38)
	attackFrames[3][action.ActionCharge] = 35

	attackFrames[4] = frames.InitNormalCancelSlice(attackHitmarks[4][2], 72)
	attackFrames[4][action.ActionCharge] = 500 //TODO: this action is illegal; need better way to handle it
}

func (c Kazuha) Attack(cfg *[4]character.CharWrapper, charindex int, startFrame int, enemy *enemy.Enemy) []combat.AttackInfo {
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
		//Mult:       attack[cfg[charindex].NormalCounter-1][cfg[charindex].Talents.Attack-1],
		MultType: []string{"atk"},
	}
	var ais []combat.AttackInfo
	if cfg[charindex].NormalCounter == 1 {
		ai.Abil = "kazuhaActionAttack1"
		ai.Mult = attack[0][0][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai)
	} else if cfg[charindex].NormalCounter == 2 {
		ai.Abil = "kazuhaActionAttack2"
		ai.Mult = attack[1][0][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai)
	} else if cfg[charindex].NormalCounter == 3 {
		ai1 := combat.AttackInfo{
			Abil:       "kazuhaActionAttack31",
			ActorIndex: charindex,
			Element:    attributes.Physical,
			MultType:   []string{"atk"},
		}
		ai1.Mult = attack[2][0][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai1)
		ai2 := combat.AttackInfo{
			Abil:       "kazuhaActionAttack32",
			ActorIndex: charindex,
			Element:    attributes.Physical,
			MultType:   []string{"atk"},
		}
		ai2.Mult = attack[2][1][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai2)
	} else if cfg[charindex].NormalCounter == 4 {
		ai.Abil = "kazuhaActionAttack4"
		ai.Mult = attack[3][0][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai)
	} else if cfg[charindex].NormalCounter == 5 {
		ai1 := combat.AttackInfo{
			Abil:       "kazuhaActionAttack51",
			ActorIndex: charindex,
			Element:    attributes.Physical,
			MultType:   []string{"atk"},
		}
		ai1.Mult = attack[4][0][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai1)
		ai2 := combat.AttackInfo{
			Abil:       "kazuhaActionAttack52",
			ActorIndex: charindex,
			Element:    attributes.Physical,
			MultType:   []string{"atk"},
		}
		ai2.Mult = attack[4][1][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai2)
		ai3 := combat.AttackInfo{
			Abil:       "kazuhaActionAttack53",
			ActorIndex: charindex,
			Element:    attributes.Physical,
			MultType:   []string{"atk"},
		}
		ai3.Mult = attack[4][2][cfg[charindex].Talents.Attack-1]
		ais = append(ais, ai3)
	}
	return ais
}

func (c Kazuha) AttackMod(startframe int, cfg *[4]character.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
}

func (c Kazuha) AttackElementAttach(enemy *enemy.Enemy) bool {
	var elementTrigger bool
	var exist bool
	for i := range enemy.ElementAttach.ActionName {
		if enemy.ElementAttach.ActionName[i] == "kazuhaActionAttack" {
			exist = true
			(*enemy).ElementAttach.Counter[i]++
			if (*enemy).ElementAttach.Counter[i]%3 == 1 {
				elementTrigger = true
			}
		}
	}
	if exist == false {
		(*enemy).ElementAttach.ActionName = append((*enemy).ElementAttach.ActionName, "kazuhaActionAttack")
		(*enemy).ElementAttach.Counter = append((*enemy).ElementAttach.Counter, 1)
		(*enemy).ElementAttach.Timer = append((*enemy).ElementAttach.Timer, 0)
		elementTrigger = true
	}
	return elementTrigger
}
