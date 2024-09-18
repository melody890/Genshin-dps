package bennett

import (
	"dps/src/core/action"
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	"dps/src/core/player/infusion"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
)

var burstFrames []int

const (
	burstStartFrame   = 34
	burstBuffDuration = 126
	burstKey          = "bennettburst"
	burstFieldKey     = "bennett-field"
)

func init() {
	burstFrames = frames.InitAbilSlice(53)
	burstFrames[action.ActionDash] = 49
	burstFrames[action.ActionJump] = 50
	burstFrames[action.ActionSwap] = 51
}

func (c Bennett) Burst(cfg *[4]character3.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "bennettActionBurst",
		Element:         attributes.Pyro,
		ElementQuantity: 2,
		Mult:            burst[cfg[charindex].Talents.Burst-1],
		MultType:        []string{"atk"},
	}
	cfg[charindex].Energy = 0
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Bennett) BurstMod(startframe int, cfg *[4]character3.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
	//假设所有人身上挂上火
	for i := range cfg {
		(*cfg)[i].ElementAttachment = attributes.Pyro
	}
	mod := modifier.Mods{
		Name:       "BennetBurst",
		Modifier:   [attributes.EndStatType]float64{},
		StartFrame: startframe,
		Dur:        780,
		CharIndex:  5,
	}
	pc := burstatk[cfg[charIndex].Talents.Burst-1]
	// 不考虑血量，默认全满血
	mod.Modifier[attributes.ATK] = pc * cfg[charIndex].BaseStats[attributes.BaseATK]
	if cfg[charIndex].Base.Cons >= 1 {
		mod.Modifier[attributes.ATK] += 0.2 * cfg[charIndex].BaseStats[attributes.BaseATK]
	}
	if cfg[charIndex].Base.Cons >= 6 {
		mod.Modifier[attributes.PyroP] = 0.15
	}
	for i := range cfg {
		(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
	}
	if cfg[charIndex].Base.Cons >= 6 {
		class := []string{"sword", "spear", "claymore"}
		element := attributes.Pyro
		tags := []attacks.AttackTag{attacks.AttackTagNormal, attacks.AttackTagExtra, attacks.AttackTagPlunge}
		inf := infusion.AddWeaponInfuse("BennettBurstInfusion", class, element, tags, true, startframe, 720, 5)
		for i := range cfg {
			(*cfg)[i].Infusion = append((*cfg)[i].Infusion, inf)
		}
	}
}

func (c Bennett) DotBurstElement(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}

func (c Bennett) BurstElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Bennett) DotBurstElementAttach(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}

func (c Bennett) DotBurstElementTrans(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers

}
