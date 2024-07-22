package choose

import (
	"dps/src/characters/bennett"
	"dps/src/characters/kazuha"
	"dps/src/characters/raiden"
	"dps/src/characters/sara"
	character2 "dps/src/core/character"
)

// 返回角色基础属性和突破加成
func Choose_character(character_name string) (*character2.Base_attribute, string, float64) {
	var character *character2.Base_attribute
	var breakthrough string
	var breakthrough_value float64
	switch character_name {
	case "raiden":
		{
			character, breakthrough = raiden.Init_raiden()
			breakthrough_value = raiden.Raiden_breakthrough
		}
	case "sara":
		{
			character, breakthrough = sara.Init_sara()
			breakthrough_value = sara.Sara_breakthrough
		}
	case "bennett":
		{
			character, breakthrough = bennett.Init_bennett()
			breakthrough_value = bennett.Bennett_breakthrough
		}
	case "kazuha":
		{
			character, breakthrough = kazuha.Init_kazuha()
			breakthrough_value = kazuha.Kazuha_breakthrough
		}
	}
	return character, breakthrough, breakthrough_value
}
