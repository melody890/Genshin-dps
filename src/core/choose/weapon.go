package choose

import (
	"dps/src/core/attribute"
	"dps/src/weapons/bow"
	"dps/src/weapons/spear"
	"dps/src/weapons/sword"
)

func Choose_weapon(weapon_name string) *attribute.Weapon_attri {
	var weapon *attribute.Weapon_attri
	switch weapon_name {
	case "bow_favonius":
		weapon = bow.Init_favonius()
	case "spear_engulfing":
		weapon = spear.Init_engulfing()
	case "sword_freedom":
		weapon = sword.Init_Freedom()
	case "sword_mistsplitte":
		weapon = sword.Init_mistsplitter()
	}

	return weapon
}
