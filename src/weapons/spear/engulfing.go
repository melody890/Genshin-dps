package spear

import "dps/src/core/attribute"

func Init_engulfing() *attribute.Weapon_attri {
	var weapon *attribute.Weapon_attri
	weapon = new(attribute.Weapon_attri)
	weapon.Base_atk = 608
	weapon.Energy_recharge = 0.551
	return weapon
}
