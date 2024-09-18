package modifier

import (
	"dps/src/core/attributes"
	"dps/src/core/character"
	"dps/src/core/combat"
	"dps/src/core/player"
	character2 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/find"
	"dps/src/template/modifier"
)

func Checkmodifier(attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character2.CharWrapper, enemy *enemy.Enemy) {
	//var modsList []modifier.Mods
	//武器
	for charIndex := range cfg {
		CharName := actions.Characters[charIndex]
		c := find.FindCharC(CharName)
		weaponName := cfg[charIndex].Weapon.Name
		weapon := find.FindWeapon(weaponName)
		weapon.CheckWeapon(c, attacks, framei, actions, cfg, charIndex, enemy)
	}
	//圣遗物,四个人的都要check
	for charIndex := range cfg {
		set := cfg[charIndex].Equip.Sets
		for setIndex := range set.SetName {
			artifactName := set.SetName[setIndex]
			if artifactName != "" {
				CharArtifact := find.FindCharC(cfg[charIndex].Base.CharName)
				artifact1 := find.FindArtifact(artifactName)
				artifact1.ApplyArtifactMod(CharArtifact, attacks, framei, actions, cfg, setIndex, charIndex, enemy)
			}
		}
	}
}

func ApplyMod(mods []modifier.Mods, ae *combat.AttackEvent, cfg character.Charcfg, activeChar string, nowFrame int) {
	//display.FmtPrintMods(mods)
	//fmt.Println(nowFrame)
	//os.Exit(0)
	for _, mod := range mods {
		if mod.StartFrame <= nowFrame {
			//对应modifier
			if mod.CharIndex == ae.Info.ActorIndex {
				for i := 0; i < int(attributes.EndStatType); i++ {
					ae.Snapshot.Stats[i] = ae.Snapshot.Stats[i] + mod.Modifier[i]
				}
			} else if mod.CharIndex == 6 { // 所有人的modifier
				for i := 0; i < int(attributes.EndStatType); i++ {
					ae.Snapshot.Stats[i] = ae.Snapshot.Stats[i] + mod.Modifier[i]
				}
			} else if mod.CharIndex == 5 && ae.Info.ActorIndex == find.FindIndex(cfg.Characters[:], activeChar) { // 前台角色的modifier
				for i := 0; i < int(attributes.EndStatType); i++ {
					ae.Snapshot.Stats[i] = ae.Snapshot.Stats[i] + mod.Modifier[i]
				}
			}
		}
	}
}
