package sword

import "dps/src/core/attribute"

func Init_mistsplitter() *attribute.Weapon_attri {
	var weapon *attribute.Weapon_attri
	weapon = new(attribute.Weapon_attri)
	weapon.Base_atk = 674
	weapon.Crit_damage = 0.441
	return weapon
}
