package bennett

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/character"
	"dps/src/core/combat"
	"fmt"
)

var (
	skillFrames       [][]int
	skillHoldHitmarks = [][]int{{45, 57}, {112, 121}}
)

const (
	skillPressHitmark   = 16                           //点按出伤
	pressParticleICDKey = "bennett-press-particle-icd" //TODO：粒子冷却，难道是充能CD？
	holdParticleICDKey  = "bennett-hold-particle-icd"
)

func SearchCharIndex(cfg character.Charcfg) int {
	//搜索班尼特的index
	CharIndex := -1
	for i, char := range cfg.Characters {
		if char == "bennett" {
			CharIndex = i
			break
		}
	}
	if CharIndex == -1 {
		fmt.Errorf("character %s not found", "bennett")
	}
	return CharIndex
}

func (c Char) Skill(cfg character.Charcfg) combat.AttackInfo {
	CharIndex := SearchCharIndex(cfg)
	ai := combat.AttackInfo{
		ActorName:          "bennett",
		Abil:               "Passion Overload (Press)",
		StrikeType:         attacks.StrikeTypeSlash,
		Element:            attributes.Pyro,
		Durability:         50,
		HitlagHaltFrames:   0.09 * 60,
		CanBeDefenseHalted: true,
		Mult:               skill[CharIndex],
	}
	return ai
}
