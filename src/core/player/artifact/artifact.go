package artifact

import (
	"dps/src/core/combat"
	"dps/src/core/player"
	character2 "dps/src/core/player/character"
	enemy "dps/src/enemies"
)

type Artifact interface {
	ApplyArtifactMod(c character2.Character, attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character2.CharWrapper, setIndex int, charIndex int, enemy *enemy.Enemy)
}
