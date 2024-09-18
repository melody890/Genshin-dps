package bennett

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

var (
	skillFrames       [][]int                         //表示班尼特技能（按下和蓄力）的各个阶段的帧数。每一维对应不同的技能施放方式或状态。
	skillHoldHitmarks = [][]int{{45, 57}, {112, 121}} //表示班尼特技能蓄力攻击的命中时间点（帧数）。
	skillHoldHitboxes = [][]float64{{2.5}, {3, 3}}    //表示蓄力技能的攻击判定范围。
	skillHoldOffsets  = []float64{0.5, 0}             //表示蓄力技能的攻击偏移量。
)

const (
	skillPressHitmark   = 16                           //表示班尼特技能按下攻击的命中时间点（帧数），为第16帧。
	pressParticleICDKey = "bennett-press-particle-icd" //和下面那个分别表示班尼特技能按下和蓄力攻击的充能球生成冷却时间（Internal Cooldown）的标识符。
	holdParticleICDKey  = "bennett-hold-particle-icd"
)

func init() {
	skillFrames = make([][]int, 5)

	// skill (press) -> x
	skillFrames[0] = frames.InitAbilSlice(42)
	skillFrames[0][action.ActionDash] = 22
	skillFrames[0][action.ActionJump] = 23
	skillFrames[0][action.ActionSwap] = 41

	// skill (hold=1) -> x
	skillFrames[1] = frames.InitAbilSlice(98)
	skillFrames[1][action.ActionBurst] = 97
	skillFrames[1][action.ActionDash] = 65
	skillFrames[1][action.ActionJump] = 66
	skillFrames[1][action.ActionSwap] = 96

	// skill (hold=1,c4) -> x
	skillFrames[2] = frames.InitAbilSlice(107)
	skillFrames[2][action.ActionDash] = 95
	skillFrames[2][action.ActionJump] = 95
	skillFrames[2][action.ActionSwap] = 106

	// skill (hold=2) -> x
	skillFrames[3] = frames.InitAbilSlice(343)
	skillFrames[3][action.ActionSkill] = 339 // uses burst frames
	skillFrames[3][action.ActionBurst] = 339
	skillFrames[3][action.ActionDash] = 231
	skillFrames[3][action.ActionJump] = 340
	skillFrames[3][action.ActionSwap] = 337

	// skill (hold=2,a4) -> x
	skillFrames[4] = frames.InitAbilSlice(175)
	skillFrames[4][action.ActionDash] = 171
	skillFrames[4][action.ActionJump] = 174
	skillFrames[4][action.ActionSwap] = 175
}

func (c Bennett) SkillPress(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "bennettActionSkillPress",
		Element:         attributes.Pyro,
		ElementQuantity: 2,
		Mult:            skill[cfg[charindex].Talents.Skill-1],
		MultType:        []string{"atk"},
	}
	//计算产球
	count := 2
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.25 {
		count = 3
	}
	ParticleElement := attributes.Pyro
	RechargeType := "particle"
	elementInfo := character2.EnergyInfo{
		Count: count, ParticleElement: ParticleElement, RechargeType: RechargeType,
	}
	energy.EnergyRecharge(cfg, elementInfo)
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Bennett) SkillHold(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	var ais []combat.AttackInfo
	return ais
}

func (c Bennett) Skill(cfg *[4]character3.CharWrapper, charindex int, frame int) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Bennett) SkillMod(cfg *[4]character3.CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy) {

}

func (c Bennett) SkillPressMod() {
}

func (c Bennett) SkillHoldMod() {

}

func (c Bennett) SkillElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Bennett) SkillPressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Bennett) SkillHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}
