package character

import "dps/src/core/attribute"

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
