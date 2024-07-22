package bennett

import (
	character2 "dps/src/core/character"
)

const (
	Bennett_hp           = 12397.0
	Bennett_atk          = 191.0
	Bennett_def          = 771.0
	Bennett_breakthrough = 0.267
)

func Init_bennett() (character2.Base_attribute, string) {
	char := character2.Base_attribute{
		Def: Bennett_def,
		Atk: Bennett_atk,
		Hp:  Bennett_hp,
	}
	breakthrough := "Energy_recharge"
	return char, breakthrough
}
