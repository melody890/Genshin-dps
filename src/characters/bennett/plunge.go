package bennett

import (
	"dps/src/core/combat"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
)

func (c Bennett) HighHoldPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Bennett) HighPressPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Bennett) PlungeHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Bennett) PlungePressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Bennett) PlungeHighMod(cfg *[4]character.CharWrapper, charindex int, frame int, enemy *enemy.Enemy, attacks []combat.AttackInfo) {
}
