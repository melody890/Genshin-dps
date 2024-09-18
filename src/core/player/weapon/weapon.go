package weapon

import (
	"dps/src/core/combat"
	"dps/src/core/player"
	character2 "dps/src/core/player/character"
	enemy "dps/src/enemies"
)

type Weapon interface {
	//InitWeapon() *weapon.Weapon
	UpdateStates(cfg *[4]character2.CharWrapper, i int)
	CheckWeapon(c character2.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character2.CharWrapper, charIndex int, enemy *enemy.Enemy)
}
