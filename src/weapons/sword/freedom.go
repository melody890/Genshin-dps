package sword

import "dps/src/core/attribute"

func Init_Freedom() *attribute.Weapon_attri {
	var weapon *attribute.Weapon_attri
	weapon = new(attribute.Weapon_attri)
	weapon.Base_atk = 608
	weapon.Elemental_mastery = 198
	return weapon
}
