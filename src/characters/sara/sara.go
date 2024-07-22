package sara

import (
	character2 "dps/src/core/character"
)

const (
	Sara_hp           = 9570.0
	Sara_atk          = 195.0
	Sara_def          = 628.0
	Sara_breakthrough = 0.24
)

func Init_sara() (*character2.Base_attribute, string) {
	var char *character2.Base_attribute
	char = new(character2.Base_attribute)
	char.Def = Sara_def
	char.Atk = Sara_atk
	char.Hp = Sara_hp
	breakthrough := "Atk_percent"
	return char, breakthrough
}
