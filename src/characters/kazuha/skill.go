package kazuha

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/energy"
	character2 "dps/src/template/character"
)

var skillPressFrames [][]int
var skillHoldFrames [][]int

const (
	skillPressHitmark = 10
	skillPressRadius  = 5
	skillPressCDStart = 8
	skillHoldHitmark  = 33
	skillHoldRadius   = 9
	skillHoldCDStart  = 31
	particleICDKey    = "kazuha-particle-icd"
)

func init() {
	// Tap E
	skillPressFrames = make([][]int, 2)
	// Tap E -> X
	skillPressFrames[0] = frames.InitAbilSlice(77) // averaged all abils
	skillPressFrames[0][action.ActionHighPlunge] = 24
	// Tap E (Glide Cancel) -> X
	skillPressFrames[1] = frames.InitAbilSlice(69)
	skillPressFrames[1][action.ActionBurst] = 61
	skillPressFrames[1][action.ActionDash] = 61
	skillPressFrames[1][action.ActionJump] = 59
	skillPressFrames[1][action.ActionSwap] = 60

	// Hold E
	skillHoldFrames = make([][]int, 2)
	// Hold E -> X
	skillHoldFrames[0] = frames.InitAbilSlice(175) // averaged all abils
	skillHoldFrames[0][action.ActionHighPlunge] = 58
	// Hold E (Glide Cancel) -> X
	skillHoldFrames[1] = frames.InitAbilSlice(160)
	skillHoldFrames[1][action.ActionAttack] = 158
	skillHoldFrames[1][action.ActionBurst] = 159
	skillHoldFrames[1][action.ActionSwap] = 155
}

func (c Kazuha) SkillPress(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionSkillPress",
		Element:         attributes.Anemo,
		ElementQuantity: 1,
		Mult:            skill[cfg[charindex].Talents.Skill-1],
		MultType:        []string{"atk"},
	}
	count := 3
	ParticleElement := attributes.Anemo
	RechargeType := "particle"
	elementInfo := character2.EnergyInfo{
		Count: count, ParticleElement: ParticleElement, RechargeType: RechargeType,
	}
	energy.EnergyRecharge(cfg, elementInfo)
	Ascension4(enemy, cfg, charindex, frame, ai)
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Kazuha) SkillHold(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionSkillHold",
		Element:         attributes.Anemo,
		ElementQuantity: 2,
		Mult:            skillHold[cfg[charindex].Talents.Skill-1],
		MultType:        []string{"atk"},
	}
	count := 4
	ParticleElement := attributes.Anemo
	RechargeType := "particle"
	elementInfo := character2.EnergyInfo{
		Count: count, ParticleElement: ParticleElement, RechargeType: RechargeType,
	}
	energy.EnergyRecharge(cfg, elementInfo)
	Ascension4(enemy, cfg, charindex, frame, ai)
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Kazuha) Skill(cfg *[4]character3.CharWrapper, charindex int, frame int) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Kazuha) SkillMod(cfg *[4]character3.CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy) {

}

func (c Kazuha) SkillPressMod() {
}

func (c Kazuha) SkillHoldMod() {

}

func (c Kazuha) SkillElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kazuha) SkillPressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kazuha) SkillHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}
