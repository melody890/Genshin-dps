package raiden

import (
	character2 "dps/src/core/character"
	"fmt"
)

const (
	Raiden_hp           = 12907.0
	Raiden_atk          = 337.0
	Raiden_def          = 789.0
	Raiden_breakthrough = 0.32
)

func Init_raiden() (*character2.Base_attribute, string) {
	var char *character2.Base_attribute
	char = new(character2.Base_attribute)
	char.Def = Raiden_def
	char.Atk = Raiden_atk
	char.Hp = Raiden_hp
	breakthrough := "Energy_recharge"
	return char, breakthrough
}

func main() {
	fmt.Println("此刻，寂灭之时")
}
