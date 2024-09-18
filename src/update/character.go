package update

import (
	"dps/src/characters/bennett"
	"dps/src/characters/diluc"
	"dps/src/characters/kaeya"
	"dps/src/characters/kazuha"
	"dps/src/core/player/character"
	"fmt"
	"log"
	"os"
)

func UpdateCharacterStatus(cfg *[4]character.CharWrapper) {
	for i := range *cfg {
		cfg[i].Index = i
		switch cfg[i].Base.CharName {
		case "bennett":
			bennettChar := &bennett.Bennett{}
			//bennettChar.
			(*cfg)[i].EnergyMax = bennettChar.InitEnergy()
			(*cfg)[i].Energy = bennettChar.InitEnergy()
			(*cfg)[i].NormalHitNum = bennettChar.NormalHitNum()
			(*cfg)[i].NormalCounter = 0
			bennettChar.UpdateStates(cfg, i)
		case "diluc":
			dilucChar := &diluc.Diluc{}
			(*cfg)[i].EnergyMax = dilucChar.InitEnergy()
			(*cfg)[i].Energy = dilucChar.InitEnergy()
			(*cfg)[i].NormalHitNum = dilucChar.NormalHitNum()
			(*cfg)[i].NormalCounter = 0
			(*cfg)[i].SkillCounter = 0
			dilucChar.UpdateStates(cfg, i)
		case "kaeya":
			kaeyaChar := &kaeya.Kaeya{}
			(*cfg)[i].EnergyMax = kaeyaChar.InitEnergy()
			(*cfg)[i].Energy = kaeyaChar.InitEnergy()
			(*cfg)[i].NormalHitNum = kaeyaChar.NormalHitNum()
			(*cfg)[i].NormalCounter = 0
			(*cfg)[i].SkillCounter = 0
			kaeyaChar.UpdateStates(cfg, i)
		case "kazuha":
			kazuhaChar := kazuha.Kazuha{}
			(*cfg)[i].EnergyMax = kazuhaChar.InitEnergy()
			(*cfg)[i].Energy = kazuhaChar.InitEnergy()
			(*cfg)[i].NormalHitNum = kazuhaChar.NormalHitNum()
			(*cfg)[i].NormalCounter = 0
			(*cfg)[i].SkillCounter = 0
			kazuhaChar.UpdateStates(cfg, i)
		case "":
		default:
			fmt.Println("ERROR!!!!  src/update/character.go 角色面板未更新")
			log.Println("ERROR!!!!  src/update/character.go 角色面板未更新")
			os.Exit(0)
		}
	}
}
