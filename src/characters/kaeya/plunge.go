package kaeya

import (
	"dps/src/core/combat"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
)

func (c Kaeya) HighHoldPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Kaeya) HighPressPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Kaeya) PlungeHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kaeya) PlungePressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kaeya) PlungeHighMod(cfg *[4]character.CharWrapper, charindex int, frame int, enemy *enemy.Enemy, attacks []combat.AttackInfo) {

}
