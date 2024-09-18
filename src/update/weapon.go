package update

import (
	character2 "dps/src/core/player/character"
	"dps/src/weapons/claymore"
	"dps/src/weapons/sword"
	"log"
	"os"
)

func UpdateWeaponStatus(cfg *[4]character2.CharWrapper) {
	for i := range *cfg {
		WeaponName := cfg[i].Weapon.Name
		switch WeaponName {
		case "mistsplitter":
			mistsplitterWeapon := &sword.Mistsplitter{}
			mistsplitterWeapon.UpdateStates(cfg, i)
		case "spine":
			spineWeapon := &claymore.Spine{}
			spineWeapon.UpdateStates(cfg, i)
		case "skyward":
			skywardWeapon := &sword.Skyward{}
			skywardWeapon.UpdateStates(cfg, i)
		case "freedom":
			freedomWeapon :=&sword.Freedom{}
			freedomWeapon.UpdateStates(cfg, i)
		case "":
		default:
			log.Println(WeaponName)
			log.Println("ERROR!src/update/weapon.go读取武器名称遇到未知武器")
			os.Exit(0)
		}
	}
}

//func UpdateWeaponRefineStatus(cfg character.Charcfg) [4][attributes.EndStatType]float64 {
//	var states [4][attributes.EndStatType]float64
//	var state [attributes.EndStatType]float64
//	for i, Weapons := range cfg.Weapons {
//		WeaponName := Weapons
//		switch WeaponName {
//		case "mistsplitter":
//			mistsplitterWeapon := &sword.Mistsplitter{}
//			mistsplitterWeapon.InitWeapon()
//			//state = mistsplitterWeapon.UpdateStates()
//			states[i] = state
//		default:
//			log.Println("ERROR!src/update/weapon.go读取武器名称遇到未知武器")
//		}
//	}
//	return states
//}
