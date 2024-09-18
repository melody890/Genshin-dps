package character

import (
	"dps/src/core/attribute"
	"dps/src/core/attributes"
)

type Charcfg struct {
	Characters     [4]string
	Weapons        [4]string
	Artifacts      [4][5]string
	CharacterLevel [4]int
	WeaponLevel    [4]int
	WeaponRefine   [4]int
	ArtifactLevel  [4][5]int
	ArtifactBonus  [4][attributes.EndStatType]float64
	TalentLevel    [4][3]int
	Constellation  [4]int
	Element        [4]attributes.Element
}

type Char struct {
	Name          string
	Constellation int
	Talent        []int
	Artifact      string
	Attribute     attribute.Attri
}

type Base_attribute struct {
	Atk float64
	Hp  float64
	Def float64
}
