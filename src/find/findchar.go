package find

import (
	"dps/src/characters/bennett"
	"dps/src/characters/diluc"
	"dps/src/characters/kaeya"
	"dps/src/characters/kazuha"
	"dps/src/core/player/character"
	tmpl "dps/src/template/character"
)

func FindCharC(charName string) character.Character {
	var c character.Character

	switch charName {
	case "bennett":
		bennettChar := &bennett.Bennett{
			Character: &tmpl.Character{},
		}
		c = bennettChar
	case "diluc":
		dilucChar := &diluc.Diluc{
			Character: &tmpl.Character{},
		}
		c = dilucChar
	case "kaeya":
		kaeyaChar := &kaeya.Kaeya{
			Character: &tmpl.Character{},
		}
		c = kaeyaChar
	case "kazuha":
		kazuhaChar := &kazuha.Kazuha{
			Character: &tmpl.Character{},
		}
		c = kazuhaChar
	}
	return c
}
