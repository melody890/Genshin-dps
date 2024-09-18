package find

import (
	"dps/src/core/player/weapon"
	"dps/src/weapons/claymore"
	"dps/src/weapons/sword"
)

func FindWeapon(weaponName string) weapon.Weapon {
	var w weapon.Weapon

	switch weaponName {
	case "mistsplitter":
		mistsplitter := &sword.Mistsplitter{}
		w = mistsplitter
	case "spine":
		spine := &claymore.Spine{}
		w = spine
	case "skyward":
		skyward := &sword.Skyward{}
		w = skyward
	case "freedom":
		freedom := &sword.Freedom{}
		w = freedom
	}
	return w
}
