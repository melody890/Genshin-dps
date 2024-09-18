package display

import (
	"dps/src/core/attributes"
	"dps/src/core/player/character"
	"dps/src/template/modifier"
	"fmt"
	"log"
)

var statNames = map[attributes.Stat]string{
	attributes.NoStat:        "NoStat",
	attributes.DEF:           "DEF",
	attributes.DEFP:          "DEFP",
	attributes.HP:            "HP",
	attributes.HPP:           "HPP",
	attributes.ATK:           "ATK",
	attributes.ATKP:          "ATKP",
	attributes.ER:            "ER",
	attributes.EM:            "EM",
	attributes.CR:            "CR",
	attributes.CD:            "CD",
	attributes.Heal:          "Heal",
	attributes.PyroP:         "PyroP",
	attributes.HydroP:        "HydroP",
	attributes.CryoP:         "CryoP",
	attributes.ElectroP:      "ElectroP",
	attributes.AnemoP:        "AnemoP",
	attributes.GeoP:          "GeoP",
	attributes.DendroP:       "DendroP",
	attributes.PhyP:          "PhyP",
	attributes.AtkSpd:        "AtkSpd",
	attributes.DmgP:          "DmgP",
	attributes.DelimBaseStat: "DelimBaseStat",
	attributes.BaseHP:        "BaseHP",
	attributes.BaseATK:       "BaseATK",
	attributes.BaseDEF:       "BaseDEF",
	attributes.MaxEnergy:     "MaxEnergy",
}

func PrintStates(cfg [4]character.CharWrapper, charIndex int) {
	if charIndex < 1 || charIndex > 5 {
		log.Println("输出角色序号不对")
		return
	}
	if charIndex == 5 {
		for i, charwrapper := range cfg {
			log.Printf("Char %d:\n", i+1)
			for statType, value := range charwrapper.BaseStats {
				if name, ok := statNames[attributes.Stat(statType)]; ok {
					log.Printf("  %s: %.4f\n", name, value)
				} else {
					log.Printf("  Unknown Stat (%d): %.4f\n", statType, value)
				}
			}
		}
	} else {
		for i, charwrapper := range cfg {
			if i+1 == charIndex {
				log.Printf("Char %d:\n", i+1)
				for statType, value := range charwrapper.BaseStats {
					if name, ok := statNames[attributes.Stat(statType)]; ok {
						log.Printf("  %s: %.4f\n", name, value)
					} else {
						log.Printf("  Unknown Stat (%d): %.4f\n", statType, value)
					}
				}
			}
		}
	}
}

func PrintSingleStates(states [attributes.EndStatType]float64) {
	log.Printf("state输出")
	for statType, value := range states {
		if name, ok := statNames[attributes.Stat(statType)]; ok {
			log.Printf("  %s: %.4f\n", name, value)
		} else {
			log.Printf("  Unknown Stat (%d): %.4f\n", statType, value)
		}
	}
}

func PrintMods(mods []modifier.Mods) {
	for i, mod := range mods {
		log.Printf("Mod %d:\n", i+1)
		for statType, value := range mod.Modifier {
			if name, ok := statNames[attributes.Stat(statType)]; ok {
				log.Printf("  %s: %.2f\n", name, value)
			} else {
				log.Printf("  Unknown Stat (%d): %.2f\n", statType, value)
			}
		}
		log.Println("Name:", mod.Name)
		log.Println("StartFrame:", mod.StartFrame)
		log.Println("CharIndex:", mod.CharIndex)
		log.Println("Dur:", mod.Dur)
	}
}

func FmtPrintMods(mods []modifier.Mods) {
	for i, mod := range mods {
		fmt.Printf("Mod %d:\n", i+1)
		for statType, value := range mod.Modifier {
			if name, ok := statNames[attributes.Stat(statType)]; ok {
				fmt.Printf("  %s: %.2f\n", name, value)
			} else {
				fmt.Printf("  Unknown Stat (%d): %.2f\n", statType, value)
			}
		}
		fmt.Println("Name:", mod.Name)
		fmt.Println("StartFrame:", mod.StartFrame)
		fmt.Println("CharIndex:", mod.CharIndex)
		fmt.Println("Dur:", mod.Dur)
	}
}
