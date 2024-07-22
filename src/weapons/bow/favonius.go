package bow

import "dps/src/core/attribute"

func Init_favonius() *attribute.Weapon_attri {
	var weapon *attribute.Weapon_attri
	weapon = new(attribute.Weapon_attri)
	weapon.Base_atk = 454
	weapon.Energy_recharge = 0.613
	return weapon
}
