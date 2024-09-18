package damage

import (
	attacks2 "dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/reactions"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func IsCrit(cfg character3.CharWrapper, attacks []combat.AttackInfo) (bool, float64) {
	var isCrit bool
	var critrate = cfg.BaseStats[attributes.CR]
	attackIndex := len(attacks) - 1
	for i := range cfg.Mods {
		if cfg.Mods[i].Condition == attacks2.AttackTagNone {
			critrate += cfg.Mods[i].Modifier[attributes.CR]
		} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg.Mods[i].Condition == attacks2.AttackBurst {
			critrate += cfg.Mods[i].Modifier[attributes.CR]
		} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg.Mods[i].Condition == attacks2.AttackTagNormal {
			critrate += cfg.Mods[i].Modifier[attributes.CR]
		} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg.Mods[i].Condition == attacks2.AttackTagPlunge {
			critrate += cfg.Mods[i].Modifier[attributes.CR]
		}
	}
	if critrate < 0 {
		critrate = 0
	}
	if critrate > 1 {
		critrate = 1
	}
	// 初始化随机数种子,生成一个0到1之间的随机数
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Float64()
	if randomValue < critrate {
		isCrit = true
	} else {
		isCrit = false
	}

	return isCrit, critrate
}

func Damage(attacks []combat.AttackInfo, actions player.Actions, cfg *[4]character3.CharWrapper, enemy *enemy.Enemy, count int, frame int) (float64, float64, float64) {
	charIndex := attacks[len(attacks)-1].ActorIndex
	attackIndex := len(attacks) - 1
	ai := attacks[attackIndex]
	st := attributes.EleToDmgP(attacks[attackIndex].Element)
	elementDamageBonus := 0.0
	if st > -1 {
		elementDamageBonus += cfg[charIndex].BaseStats[st]
		if attacks[attackIndex].Lock == false {
			for i := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if charIndex == cfg[charIndex].Mods[i].CharIndex || (cfg[charIndex].Mods[i].CharIndex == 5 && cfg[charIndex].Active) || cfg[charIndex].Mods[i].CharIndex == 6 {
						elementDamageBonus += cfg[charIndex].Mods[i].Modifier[st]
					}
				}
			}
		} else {
			for i := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if charIndex == attacks[attackIndex].LockMod[i].CharIndex || (attacks[attackIndex].LockMod[i].CharIndex == 5 && cfg[charIndex].Active) || attacks[attackIndex].LockMod[i].CharIndex == 6 {
						elementDamageBonus += attacks[attackIndex].LockMod[i].Modifier[st]
					}
				}
			}
		}
	}

	//这里结算的元素伤害加成，由单独的元素伤害加成和总的伤害加成构成
	dmgBonus := elementDamageBonus + cfg[charIndex].BaseStats[attributes.DmgP]

	//fmt.Println("dmgBonus", dmgBonus)
	//print("buff来源：  ")
	if attacks[attackIndex].Lock == false {
		for i := range cfg[charIndex].Mods {
			//fmt.Println(cfg[charIndex].Mods[i].Name)
			if cfg[charIndex].Mods[i].StartFrame <= frame {
				if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
					dmgBonus += cfg[charIndex].Mods[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
					dmgBonus += cfg[charIndex].Mods[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
					dmgBonus += cfg[charIndex].Mods[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
					dmgBonus += cfg[charIndex].Mods[i].Modifier[attributes.DmgP]
				}
			}
		}
	} else {
		for i := range attacks[attackIndex].LockMod {
			//fmt.Println(attacks[attackIndex].LockMod[i].Name)
			if attacks[attackIndex].LockMod[i].StartFrame <= frame {
				if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
					dmgBonus += attacks[attackIndex].LockMod[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
					dmgBonus += attacks[attackIndex].LockMod[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
					dmgBonus += attacks[attackIndex].LockMod[i].Modifier[attributes.DmgP]
					//fmt.Println("DMP来源", cfg[charIndex].Mods[i].Name, "加成值", cfg[charIndex].Mods[i].Modifier[attributes.DmgP])
				} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
					dmgBonus += attacks[attackIndex].LockMod[i].Modifier[attributes.DmgP]
				}
			}
		}
	}
	//print("\n")
	//fmt.Print("dmgBonus:", dmgBonus, "\n")
	// calculate using attack or def
	var a float64
	if ai.UseHP {
		var HPP float64
		a = cfg[charIndex].BaseStats[attributes.BaseHP]
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						HPP += mod.Modifier[attributes.HPP]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						HPP += mod.Modifier[attributes.HPP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						HPP += mod.Modifier[attributes.HPP]
					}
				}
			}
		}
		a *= 1 + HPP + cfg[charIndex].BaseStats[attributes.HPP]
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.HP]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.HP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.HP]
					}
				}
			}
		}
		a += cfg[charIndex].BaseStats[attributes.HP]
	} else if ai.UseDef {
		var DEFP float64
		a = cfg[charIndex].BaseStats[attributes.BaseDEF]
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						DEFP += mod.Modifier[attributes.DEFP]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						DEFP += mod.Modifier[attributes.DEFP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						DEFP += mod.Modifier[attributes.DEFP]
					}
				}
			}
		}
		a *= 1 + DEFP + cfg[charIndex].BaseStats[attributes.DEFP]
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.DEF]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.DEF]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.DEF]
					}
				}
			}
		}
		a += cfg[charIndex].BaseStats[attributes.DEF]
	} else {
		a = cfg[charIndex].BaseStats[attributes.BaseATK]
		var ATKP float64
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						ATKP += mod.Modifier[attributes.ATKP]
						//fmt.Println(mod.Name, "ATKP: ", mod.Modifier[attributes.ATKP])
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						ATKP += mod.Modifier[attributes.ATKP]
						//fmt.Println(mod.Name, "ATKP: ", mod.Modifier[attributes.ATKP])
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						ATKP += mod.Modifier[attributes.ATKP]
						//fmt.Println(mod.Name, "ATKP: ", mod.Modifier[attributes.ATKP])
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						ATKP += mod.Modifier[attributes.ATKP]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						ATKP += mod.Modifier[attributes.ATKP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						ATKP += mod.Modifier[attributes.ATKP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						ATKP += mod.Modifier[attributes.ATKP]
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						ATKP += mod.Modifier[attributes.ATKP]
					}
				}
			}
		}
		a *= 1 + ATKP + cfg[charIndex].BaseStats[attributes.ATKP]
		if attacks[attackIndex].Lock == false {
			for i, mod := range cfg[charIndex].Mods {
				if cfg[charIndex].Mods[i].StartFrame <= frame {
					if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.ATK]
					}
				}
			}
		} else {
			for i, mod := range attacks[attackIndex].LockMod {
				if attacks[attackIndex].LockMod[i].StartFrame <= frame {
					if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
						a += mod.Modifier[attributes.ATK]
						//fmt.Println(mod.Name, "ATK: ", mod.Modifier[attributes.ATK])
					} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
						a += mod.Modifier[attributes.ATK]
					}
				}
			}
		}
		a += cfg[charIndex].BaseStats[attributes.ATK]
		//fmt.Println("ArtifactATK", cfg[charIndex].BaseStats[attributes.ATK])
		//log.Println("BaseAtk:", ae.Snapshot.BaseAtk, "ATKP: ", ae.Snapshot.Stats[attributes.ATKP], "ATK: ", ae.Snapshot.Stats[attributes.ATK])
	}
	//base是倍率+FlatDmg（激化，钟离，闲云等）
	base := ai.Mult*a + ai.FlatDmg
	//if frame == 453 {
	//	fmt.Println("BaseAtk", cfg[charIndex].BaseStats[attributes.BaseATK])
	//	//fmt.Println("ATKP", cfg[charIndex].BaseStats[attributes.ATKP])
	//	fmt.Println("裸面板ATK", cfg[charIndex].BaseStats[attributes.BaseATK]*(1+cfg[charIndex].BaseStats[attributes.ATKP])+cfg[charIndex].BaseStats[attributes.ATK])
	//	fmt.Println("atk:", a)
	//	fmt.Println("FlatDmg:", ai.FlatDmg)
	//	fmt.Println("mult: ", ai.Mult)
	//	//fmt.Println("base:", base)
	//}
	//log.Println("Mult:", ai.Mult, "Base: ", base)
	damage := base * (1 + dmgBonus)
	//fmt.Println("元素伤害：", damage)
	res := enemy.Res[ai.Element]
	//fmt.Println("抗性", res)
	//fmt.Println("攻击元素", ai.Element)

	for i := range enemy.EnemyMod {
		if enemy.EnemyMod[i].StartFrame <= frame && enemy.EnemyMod[i].StartFrame+enemy.EnemyMod[i].Dur > frame {
			res += enemy.EnemyMod[i].Res[ai.Element]
		}
	}
	//fmt.Println("应用减抗后的抗性:", res)
	resmod := 1 - res/2
	if res >= 0 && res < 0.75 {
		resmod = 1 - res
	} else if res > 0.75 {
		resmod = 1 / (4*res + 1)
	}
	damage *= resmod
	//fmt.Println("resmod:", resmod)
	//fmt.Println("res dmg: ", damage)
	//fmt.Println("ignore: ", ai.IgnoreDefPercent)

	defmod := float64(cfg[charIndex].Base.Level+100) /
		(float64(cfg[charIndex].Base.Level+100) +
			float64(enemy.Level+100)*(1-ai.IgnoreDefPercent))
	damage *= defmod
	//fmt.Println("defmod: ", defmod)
	//fmt.Println("dmg def: ", damage)

	//preampdmg := damage

	// 精通区
	// calculate em bonus
	em := cfg[charIndex].BaseStats[attributes.EM]
	//fmt.Println("加buff前元素精通：", em)
	if attacks[attackIndex].Lock == false {
		for i, mod := range cfg[charIndex].Mods {
			if cfg[charIndex].Mods[i].StartFrame <= frame {
				if cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNone {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && cfg[charIndex].Mods[i].Condition == attacks2.AttackBurst {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagNormal {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
					em += mod.Modifier[attributes.EM]
				}
			}
		}
	} else {
		for i, mod := range attacks[attackIndex].LockMod {
			if attacks[attackIndex].LockMod[i].StartFrame <= frame {
				if attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNone {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Burst") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackBurst {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Attack") && attacks[attackIndex].LockMod[i].Condition == attacks2.AttackTagNormal {
					em += mod.Modifier[attributes.EM]
				} else if strings.Contains(attacks[attackIndex].Abil, "Plunge") && cfg[charIndex].Mods[i].Condition == attacks2.AttackTagPlunge {
					em += mod.Modifier[attributes.EM]
				}
			}
		}
	}
	//fmt.Println("加buff后元素精通：", em)

	actionMult := -1.0
	var swirlDamage float64
	var swirlTag bool
	if ai.Trigger == true {
		// 毫无反应，单纯上元素
		reactions.AddElement(ai, enemy)
		// 同类元素补充机制
		reactions.SupplyElement(ai, enemy)
		// 扩散反应
		swirlDamage, swirlTag = reactions.DamageSwirl(&ai, enemy, em, cfg[charIndex].Base.Level, frame)
		// 增幅反应，其中AmpMult是增幅反应倍率，返回反应乘区
		actionMult = reactions.CheckAmplified(&ai, enemy, em)
	}

	if ai.Amped {
		//反应伤害加成
		damage *= actionMult
	}

	precritdmg := damage
	//TODO:这里我全都输出
	fmt.Println("未暴击伤害：", damage)
	fmt.Println("暴击后伤害：", damage*(1+cfg[charIndex].BaseStats[attributes.CD]))
	iscrit, critRate := IsCrit(cfg[charIndex], attacks)
	ExpectationDamage := (1-critRate)*damage + critRate*damage*(1+cfg[charIndex].BaseStats[attributes.CD])
	if iscrit {
		damage *= 1 + cfg[charIndex].BaseStats[attributes.CD]
	}

	if swirlTag {
		fmt.Printf("   造成扩散伤害：%.3f", swirlDamage)
	}
	print("\n")
	//}
	log.Println("precritdmg", precritdmg, "critdamage", precritdmg*(1+cfg[charIndex].BaseStats[attributes.CD]))
	//log.Println("preampdmg", preampdmg)
	log.Println("damage", damage)
	//fmt.Println("damage: ", damage)
	return damage, ExpectationDamage, swirlDamage
}
