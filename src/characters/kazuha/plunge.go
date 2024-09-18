package kazuha

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
)

var plungePressFrames []int
var plungeHoldFrames []int

// a1 is 1 frame before this
// collision is 6 frame before this
const plungePressHitmark = 36
const plungeHoldHitmark = 41

var highPlungeFrames []int
var lowPlungeFrames []int

const lowPlungeHitmark = 46
const highPlungeHitmark = 47
const collisionHitmark = lowPlungeHitmark - 6

const lowPlungePoiseDMG = 100.0
const lowPlungeRadius = 3.0

const highPlungePoiseDMG = 150.0
const highPlungeRadius = 5.0

// TODO: missing plunge -> skill
func init() {
	// skill (press) -> high plunge -> x
	plungePressFrames = frames.InitAbilSlice(55) // max
	plungePressFrames[action.ActionDash] = 43
	plungePressFrames[action.ActionJump] = 50
	plungePressFrames[action.ActionSwap] = 50

	// skill (hold) -> high plunge -> x
	plungeHoldFrames = frames.InitAbilSlice(61) // max
	plungeHoldFrames[action.ActionSkill] = 60   // uses burst frames
	plungeHoldFrames[action.ActionBurst] = 60
	plungeHoldFrames[action.ActionDash] = 48
	plungeHoldFrames[action.ActionJump] = 55
	plungeHoldFrames[action.ActionSwap] = 54

	// low_plunge -> x
	lowPlungeFrames = frames.InitAbilSlice(73)
	lowPlungeFrames[action.ActionAttack] = 52
	lowPlungeFrames[action.ActionSkill] = 52
	lowPlungeFrames[action.ActionBurst] = 51
	lowPlungeFrames[action.ActionDash] = lowPlungeHitmark
	lowPlungeFrames[action.ActionJump] = 69
	lowPlungeFrames[action.ActionSwap] = 53

	// high_plunge -> x
	highPlungeFrames = frames.InitAbilSlice(73)
	highPlungeFrames[action.ActionAttack] = 54
	highPlungeFrames[action.ActionSkill] = 53
	highPlungeFrames[action.ActionBurst] = 53
	highPlungeFrames[action.ActionDash] = highPlungeHitmark
	highPlungeFrames[action.ActionJump] = 69
	highPlungeFrames[action.ActionSwap] = 55
}

func (c Kazuha) HighHoldPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	var ais []combat.AttackInfo
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionPlungeHold",
		Element:         attributes.Anemo,
		ElementQuantity: 1,
		Mult:            highPlunge[cfg[charindex].Talents.Attack-1],
		MultType:        []string{"atk"},
		Startframe:      frame,
	}
	ais = append(ais, ai)
	//元素转化
	//突破1：千早振在施放时，如果接触了水元素/火元素/冰元素/雷元素，则会使这次千早振的下落攻击·乱岚拨止，发生元素转化，将附加200%攻击力的对应元素伤害，该伤害视为下落攻击伤害。每次千早振的技能效果中，元素转化仅会发生一次。
	plungeElement := CheckBurstElement(cfg, enemy)
	aiChange := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionPlungeHoldChange",
		Element:         plungeElement,
		ElementQuantity: 1,
		Mult:            2.0,
		MultType:        []string{"atk"},
		AttackFrame:     plungeHoldHitmark + frame + 1,
		Trigger:         true,
	}
	*dotAction = append(*dotAction, aiChange)
	//ais = append(ais, aiChange)
	return ais
}

func (c Kazuha) HighPressPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	var ais []combat.AttackInfo
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionPlungePress",
		Element:         attributes.Anemo,
		ElementQuantity: 1,
		Mult:            highPlunge[cfg[charindex].Talents.Attack-1],
		MultType:        []string{"atk"},
		Startframe:      frame,
	}
	ais = append(ais, ai)
	//元素转化
	//突破1：千早振在施放时，如果接触了水元素/火元素/冰元素/雷元素，则会使这次千早振的下落攻击·乱岚拨止，发生元素转化，将附加200%攻击力的对应元素伤害，该伤害视为下落攻击伤害。每次千早振的技能效果中，元素转化仅会发生一次。
	plungeElement := CheckBurstElement(cfg, enemy)
	aiChange := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionPlungePressChange",
		Element:         plungeElement,
		ElementQuantity: 1,
		Mult:            2.0,
		MultType:        []string{"atk"},
		AttackFrame:     plungeHoldHitmark + frame,
		Trigger:         true,
	}
	*dotAction = append(*dotAction, aiChange)
	//ais = append(ais, aiChange)
	return ais
}

func (c Kazuha) PlungeHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kazuha) PlungePressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kazuha) PlungeHighMod(cfg *[4]character.CharWrapper, charindex int, frame int, enemy *enemy.Enemy, attacks []combat.AttackInfo) {
	ai := attacks[len(attacks)-1]
	plungeElement := CheckBurstElement(cfg, enemy)
	(*enemy).Element = plungeElement
	Ascension4(enemy, cfg, charindex, frame, ai)
}
