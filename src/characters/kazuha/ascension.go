package kazuha

import (
	attacks2 "dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
	"strings"
)

func Ascension4(enemy *enemy.Enemy, cfg *[4]character.CharWrapper, charindex int, frame int, ai combat.AttackInfo) {
	//突破4：枫原万叶触发扩散反应后，枫原万叶的每点元素精通，会为队伍中所有角色提供0.04%对应元素伤害加成，持续8秒。通过这种方式获得的不同元素伤害加成可以共存。
	//值得注意的是，这个天赋在后台也可以生效

	if enemy.Element == attributes.Pyro || enemy.Element == attributes.Cryo || enemy.Element == attributes.Hydro || enemy.Element == attributes.Electro {
		for charIndex := range len(*cfg) {
			for i := len((*cfg)[charIndex].Mods) - 1; i >= 0; i-- {
				modDel := (*cfg)[charIndex].Mods[i]
				if modDel.Name == "kazuhaAscension4Buff" && modDel.Modifier[attributes.EleToDmgP(enemy.Element)] != 0 {
					(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[i+1:]...)
				}
			}
		}
		if !existA4Buff(cfg, enemy.Element) {
			A4mod := modifier.Mods{
				Name:       "kazuhaAscension4Buff",
				StartFrame: frame + 1,
				Dur:        480,
				CharIndex:  6,
			}
			var em float64
			for i, mod := range cfg[charindex].Mods {
				if cfg[charindex].Mods[i].Condition == attacks2.AttackTagNone {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(ai.Abil, "Burst") && cfg[charindex].Mods[i].Condition == attacks2.AttackBurst {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(ai.Abil, "Attack") && cfg[charindex].Mods[i].Condition == attacks2.AttackTagNormal {
					em += mod.Modifier[attributes.EM]
				}
			}
			A4mod.Modifier[attributes.EleToDmgP(enemy.Element)] = (cfg[charindex].BaseStats[attributes.EM] + em) * 0.0004
			if cfg[charindex].Base.Cons >= 2 && ai.Abil == "kazuhaActionBurst" {
				A4mod.Modifier[attributes.EleToDmgP(enemy.Element)] += 0.08
			}
			for i := range cfg {
				(*cfg)[i].Mods = append((*cfg)[i].Mods, A4mod)
			}
		}
	}
}

func existA4Buff(cfg *[4]character.CharWrapper, element attributes.Element) bool {
	for i := range cfg {
		modsi := cfg[i].Mods
		for j := range modsi {
			if modsi[j].Name == "kazuhaAscension4Buff" {
				if element == attributes.Pyro && (modsi[j].Modifier[attributes.PyroP] != 0) {
					return true
				} else if element == attributes.Cryo && (modsi[j].Modifier[attributes.CryoP] != 0) {
					return true
				} else if element == attributes.Electro && (modsi[j].Modifier[attributes.ElectroP] != 0) {
					return true
				} else if element == attributes.Hydro && (modsi[j].Modifier[attributes.HydroP] != 0) {
					return true
				}
			}
		}
	}
	return false
}
