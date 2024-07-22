package kazuha

import (
	character2 "dps/src/core/character"
)

const (
	Kazuha_hp           = 13348.0
	Kazuha_atk          = 297.0
	Kazuha_def          = 807.0
	Kazuha_breakthrough = 115.0
)

func Init_kazuha() (character2.Base_attribute, string) {
	char := character2.Base_attribute{
		Def: Kazuha_def,
		Atk: Kazuha_atk,
		Hp:  Kazuha_hp,
	}
	breakthrough := "Elemental_mastery"
	return char, breakthrough
}
