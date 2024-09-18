package infusion

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
)

type WeaponInfusion struct {
	Name            string
	Class           []string
	Ele             attributes.Element
	Tags            []attacks.AttackTag
	CanBeOverridden bool

	StartFrame int
	Dur        int
	CharIndex  int
}

func AddWeaponInfuse(name string, class []string, element attributes.Element, tags []attacks.AttackTag, can bool, start int, dur int, charindex int) WeaponInfusion {
	inf := WeaponInfusion{
		Name:            name,
		Class:           class,
		Ele:             element,
		Tags:            tags,
		CanBeOverridden: can,
		StartFrame:      start,
		Dur:             dur,
		CharIndex:       charindex,
	}
	return inf
}
