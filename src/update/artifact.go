package update

import (
	"dps/src/core/player/character"
)

func UpdateArtifactStatus(cfg *[4]character.CharWrapper) {
	for charIndex := range *cfg {
		for i := range (*cfg)[charIndex].Equip.Sets.SetBonusState {
			(*cfg)[charIndex].BaseStats[i] += (*cfg)[charIndex].Equip.Sets.SetBonusState[i]
		}
	}
}
