package bennett

import (
	character2 "dps/src/core/character"
	tmpl "dps/src/template/character"
)

const (
	Bennett_hp           = 12397.0
	Bennett_atk          = 191.0
	Bennett_def          = 771.0
	Bennett_breakthrough = 0.267
)

type Char struct {
	*tmpl.Character
}

// InitChar 初始化 Char 实例并返回 tmpl.Character 实例
// 里面是角色名字和初始三围
func (c *Char) InitChar(CharName string) *tmpl.Character {
	c.Character = &tmpl.Character{
		Name: CharName,
		Hp:   Bennett_hp,
		Atk:  Bennett_atk,
		Def:  Bennett_def,
	}
	return c.Character
}

func Init_bennett() (character2.Base_attribute, string) {
	char := character2.Base_attribute{
		Def: Bennett_def,
		Atk: Bennett_atk,
		Hp:  Bennett_hp,
	}
	breakthrough := "Energy_recharge"
	return char, breakthrough
}
