package kaeya

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/energy"
	character2 "dps/src/template/character"
	"math/rand"
	"time"
)

var skillFrames []int

const (
	skillHitmark   = 28
	particleICDKey = "kaeya-particle-icd"
)

func init() {
	skillFrames = frames.InitAbilSlice(53) // E -> N1
	skillFrames[action.ActionBurst] = 52   // E -> Q
	skillFrames[action.ActionDash] = 25    // E -> D
	skillFrames[action.ActionJump] = 26    // E -> J
	skillFrames[action.ActionSwap] = 49    // E -> Swap
}

func (c Kaeya) Skill(cfg *[4]character3.CharWrapper, charindex int, frame int) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	cfg[charindex].Active = true
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kaeyaActionSkill",
		Element:         attributes.Cryo,
		ElementQuantity: 2,
		Mult:            skill[cfg[charindex].Talents.Skill-1],
		MultType:        []string{"atk"},
	}
	//计算产球
	count := 2
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.67 {
		count = 3
	}
	ParticleElement := attributes.Cryo
	RechargeType := "particle"
	elementInfo := character2.EnergyInfo{
		Count: count, ParticleElement: ParticleElement, RechargeType: RechargeType,
	}
	energy.EnergyRecharge(cfg, elementInfo)
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Kaeya) SkillPress(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Kaeya) SkillHold(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	var ais []combat.AttackInfo
	return ais
}

func (c Kaeya) SkillMod(cfg *[4]character3.CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy) {
}

func (c Kaeya) SkillPressMod() {
}

func (c Kaeya) SkillHoldMod() {

}

func (c Kaeya) SkillElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kaeya) SkillPressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kaeya) SkillHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}
