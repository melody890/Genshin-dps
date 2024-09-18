package diluc

import (
	"dps/src/core/combat"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
)

func (c Diluc) HighHoldPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Diluc) HighPressPlunge(cfg *[4]character.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Diluc) PlungeHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Diluc) PlungePressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Diluc) PlungeHighMod(cfg *[4]character.CharWrapper, charindex int, frame int, enemy *enemy.Enemy, attacks []combat.AttackInfo) {
}
