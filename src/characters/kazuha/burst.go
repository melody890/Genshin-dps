package kazuha

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
)

var burstFrames []int

const (
	burstHitmark   = 82
	burstFirstTick = 135
)

const burstStatus = "kazuha-q"

func init() {
	burstFrames = frames.InitAbilSlice(93) // Q -> J
	burstFrames[action.ActionAttack] = 92  // Q -> N1
	burstFrames[action.ActionSkill] = 92   // Q -> E
	burstFrames[action.ActionDash] = 92    // Q -> D
	burstFrames[action.ActionSwap] = 90    // Q -> Swap
}

func (c Kazuha) Burst(cfg *[4]character3.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kazuhaActionBurst",
		Element:         attributes.Anemo,
		ElementQuantity: 2,
		Mult:            burstSlash[cfg[charindex].Talents.Burst-1],
		MultType:        []string{"atk"},
	}
	Ascension4(enemy, cfg, charindex, frame, ai)
	cfg[charindex].Energy = 0
	offset := 117 //大招的dot每117帧生效一次，一共5次
	for i := range 5 {
		dotAI := combat.AttackInfo{
			ActorIndex:      charindex,
			Abil:            "kazuhaDotBurst",
			Element:         attributes.Anemo,
			ElementQuantity: 1,
			Mult:            burstDot[cfg[charindex].Talents.Burst-1],
			MultType:        []string{"atk"},
			AttackFrame:     burstFirstTick + offset*i + frame,
			Lock:            true,
		}
		*dotAction = append(*dotAction, dotAI)
	}
	//附着伤害，如果怪物身上有元素，那么是有一段附着伤害的，而且之后的dot会染色。
	burstElement := CheckBurstElement(cfg, enemy)
	//染色以及附着伤害
	for i := range 5 {
		AiAttach := combat.AttackInfo{
			ActorIndex:      charindex,
			Abil:            "kazuhaDotBurstAttach",
			Element:         burstElement,
			ElementQuantity: 1,
			Mult:            burstEleDot[cfg[charindex].Talents.Burst-1],
			MultType:        []string{"atk"},
			AttackFrame:     burstFirstTick + offset*i + frame,
			Lock:            true,
		}
		if burstElement == attributes.Anemo {
			AiAttach.Mult = 0
		}
		*dotAction = append(*dotAction, AiAttach)
	}

	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func ChangeDotColor(ai *combat.AttackInfo, element attributes.Element, cfg *[4]character3.CharWrapper) {
	(*ai).Element = element
	charindex := ai.ActorIndex
	if element != attributes.Anemo {
		(*ai).Mult = burstDot[cfg[charindex].Talents.Burst-1]
	}
}

func (c Kazuha) BurstMod(startframe int, cfg *[4]character3.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
	if cfg[charIndex].Base.Cons >= 2 {
		mod := modifier.Mods{
			Name:       "KazuhaC2Burst",
			StartFrame: startframe,
			Dur:        12 * 60,
			CharIndex:  6,
		}
		mod.Modifier[attributes.EM] = 200
		for i := range cfg {
			(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
		}
	}
}

func CheckBurstElement(cfg *[4]character3.CharWrapper, enemy *enemy.Enemy) attributes.Element {
	var EnvironmentElements []attributes.Element

	for i := range cfg {
		EnvironmentElements = append(EnvironmentElements, cfg[i].ElementAttachment)
	}
	EnvironmentElements = append(EnvironmentElements, enemy.Element)
	for i := range EnvironmentElements {
		if EnvironmentElements[i] == attributes.Pyro {
			return attributes.Pyro
		}
	}
	for i := range EnvironmentElements {
		if EnvironmentElements[i] == attributes.Hydro {
			return attributes.Hydro
		}
	}
	for i := range EnvironmentElements {
		if EnvironmentElements[i] == attributes.Electro {
			return attributes.Electro
		}
	}
	for i := range EnvironmentElements {
		if EnvironmentElements[i] == attributes.Cryo {
			return attributes.Cryo
		}
	}
	return attributes.Anemo
}

func (c Kazuha) DotBurstElement(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	for _ = range timer {
		triggers = append(triggers, true)
	}
	return triggers
}

func (c Kazuha) BurstElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kazuha) DotBurstElementAttach(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	for _ = range timer {
		triggers = append(triggers, true)
	}
	return triggers
}

func (c Kazuha) DotBurstElementTrans(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}
