package main

import (
	"dps/src/characters/bennett"
	"dps/src/characters/diluc"
	"dps/src/characters/kaeya"
	"dps/src/characters/kazuha"
	"dps/src/core/attributes"
	character2 "dps/src/core/character"
	"dps/src/core/combat"
	"dps/src/core/info"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	"dps/src/damage"
	"dps/src/display"
	"dps/src/enemies"
	"dps/src/find"
	"dps/src/modifier"
	"dps/src/reactions"
	"dps/src/update"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type CharacterInfo struct {
	Char         character2.Base_attribute
	Breakthrough string
}

func Readactions() player.Actions {
	action := player.Actions{
		Characters: []string{"bennett", "diluc", "kaeya", "kazuha"},
		//Actions:    [][]string{{"bennett", "ActionBurst"}},
		Actions: [][]string{{"kaeya", "ActionSkill"}, {"kazuha", "ActionBurst"}, {"bennett", "ActionBurst"}, {"bennett", "ActionSkillPress"}, {"bennett", "ActionAttack"}, {"kazuha", "ActionSkillPress"}, {"kazuha", "ActionPlungePress"}, {"kaeya", "ActionBurst"}, {"diluc", "ActionSkill"}, {"diluc", "ActionAttack"}, {"diluc", "ActionSkill"}, {"diluc", "ActionAttack"}, {"diluc", "ActionSkill"}, {"diluc", "ActionBurst"}},
		//Actions: [][]string{{"kazuha", "ActionBurst"}, {"kazuha", "ActionSkillHold"}, {"kazuha", "ActionPlungeHold"}},
		//Actions: [][]string{{"kaeya", "ActionBurst"}},
		//Actions: [][]string{{"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}, {"diluc", "ActionSkill"}, {"diluc", "ActionSkill"}, {"diluc", "ActionSkill"}, {"diluc", "ActionBurst"}, {"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}, {"diluc", "ActionAttack"}},
		//Actions: [][]string{{"bennett", "ActionSkillPress"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionBurst"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}, {"bennett", "ActionAttack"}},
		Delay: []int{6, 23, 3, 9, 2, 0, 0, 15, 2, 17, 6, 0, 12, 0},
	}
	return action
}

func Readcfg() [4]character3.CharWrapper {
	cfg := [4]character3.CharWrapper{}
	//班尼特
	cfg[0].Index = 0
	cfg[0].Base = info.CharacterBase{
		CharName:  "bennett",
		Element:   attributes.Pyro,
		Level:     90,
		MaxLevel:  90,
		HP:        bennett.Bennett_hp,
		Atk:       bennett.Bennett_atk,
		Def:       bennett.Bennett_def,
		Cons:      6,
		Ascension: 6,
	}
	cfg[0].Weapon = info.WeaponProfile{
		Name:     "mistsplitter",
		Refine:   1,
		Level:    90,
		MaxLevel: 90,
		Class:    "sword",
	}
	cfg[0].Talents = info.TalentProfile{
		Attack: 1,
		Skill:  8,
		Burst:  13,
	}
	cfg[0].Equip.Sets = info.Sets{
		SetNum:  [2]int{5},
		SetName: [2]string{"noblesse"},
	}
	cfg[0].Equip.Sets.SetBonusState[attributes.ATK] += 45 + 311 + 33 + 45
	cfg[0].Equip.Sets.SetBonusState[attributes.ATKP] += 0
	cfg[0].Equip.Sets.SetBonusState[attributes.DEF] += 39
	cfg[0].Equip.Sets.SetBonusState[attributes.DEFP] += 0.073 + 0.109 + 0.102
	cfg[0].Equip.Sets.SetBonusState[attributes.HP] += 4780 + 687 + 538
	cfg[0].Equip.Sets.SetBonusState[attributes.HPP] += 0.152 + 0.466
	cfg[0].Equip.Sets.SetBonusState[attributes.CD] += 0.07
	cfg[0].Equip.Sets.SetBonusState[attributes.CR] += 0.066 + 0.074 + 0.097 + 0.027
	cfg[0].Equip.Sets.SetBonusState[attributes.ER] += 0.181 + 0.052 + 0.518 + 0.188
	cfg[0].Equip.Sets.SetBonusState[attributes.EM] += 19 + 23
	cfg[0].Equip.Sets.SetBonusState[attributes.Heal] += 0.359

	cfg[0].ElementAttachment = attributes.NoElement

	//迪卢克
	cfg[1].Index = 1
	cfg[1].Base = info.CharacterBase{
		CharName:  "diluc",
		Element:   attributes.Pyro,
		Level:     90,
		MaxLevel:  90,
		HP:        diluc.Diluc_hp,
		Atk:       diluc.Diluc_atk,
		Def:       diluc.Diluc_def,
		Ascension: 6,
		Cons:      4,
	}
	cfg[1].Weapon = info.WeaponProfile{
		Name:     "spine",
		Refine:   5,
		Level:    90,
		MaxLevel: 90,
		Class:    "claymore",
	}
	cfg[1].Talents = info.TalentProfile{
		Attack: 10,
		Skill:  13,
		Burst:  10,
	}
	cfg[1].Equip.Sets = info.Sets{
		SetNum:  [2]int{5},
		SetName: [2]string{"crimson"},
	}
	cfg[1].Equip.Sets.SetBonusState[attributes.ATK] += 33 + 311
	cfg[1].Equip.Sets.SetBonusState[attributes.ATKP] += 0.111 + 0.082 + 0.466 + 0.053
	cfg[1].Equip.Sets.SetBonusState[attributes.DEF] += 19
	cfg[1].Equip.Sets.SetBonusState[attributes.DEFP] += 0.117
	cfg[1].Equip.Sets.SetBonusState[attributes.HP] += 509 + 4780
	cfg[1].Equip.Sets.SetBonusState[attributes.HPP] += 0.058
	cfg[1].Equip.Sets.SetBonusState[attributes.CD] += 0.622 + 0.109 + 0.21 + 0.264 + 0.202
	cfg[1].Equip.Sets.SetBonusState[attributes.CR] += 0.101 + 0.007 + 0.062 + 0.074 + 0.066
	cfg[1].Equip.Sets.SetBonusState[attributes.ER] += 0.052
	cfg[1].Equip.Sets.SetBonusState[attributes.EM] += 40 + 51
	cfg[1].Equip.Sets.SetBonusState[attributes.Heal] += 0
	cfg[1].Equip.Sets.SetBonusState[attributes.PyroP] += 0.466

	cfg[1].ElementAttachment = attributes.NoElement

	//凯亚
	cfg[2].Index = 2
	cfg[2].Base = info.CharacterBase{
		CharName:  "kaeya",
		Element:   attributes.Cryo,
		Level:     90,
		MaxLevel:  90,
		HP:        kaeya.Kaeya_hp,
		Atk:       kaeya.Kaeya_atk,
		Def:       kaeya.Kaeya_def,
		Ascension: 6,
		Cons:      4,
	}
	cfg[2].Weapon = info.WeaponProfile{
		Name:     "skyward",
		Refine:   1,
		Level:    90,
		MaxLevel: 90,
		Class:    "sword",
	}
	cfg[2].Talents = info.TalentProfile{
		Attack: 2,
		Skill:  10,
		Burst:  8,
	}
	cfg[2].Equip.Sets = info.Sets{
		SetNum:  [2]int{5},
		SetName: [2]string{"emblem"},
	}
	cfg[2].Equip.Sets.SetBonusState[attributes.ATK] += 14 + 37 + 311 + 27
	cfg[2].Equip.Sets.SetBonusState[attributes.ATKP] += 0.466
	cfg[2].Equip.Sets.SetBonusState[attributes.DEF] += 58 + 23 + 19 + 35
	cfg[2].Equip.Sets.SetBonusState[attributes.DEFP] += 0
	cfg[2].Equip.Sets.SetBonusState[attributes.HP] += 4780
	cfg[2].Equip.Sets.SetBonusState[attributes.HPP] += 0.169
	cfg[2].Equip.Sets.SetBonusState[attributes.CD] += 0.132 + 0.179 + 0.148 + 0.117
	cfg[2].Equip.Sets.SetBonusState[attributes.CR] += 0.035 + 0.062 + 0.078 + 0.093
	cfg[2].Equip.Sets.SetBonusState[attributes.ER] += 0.058 + 0.24 + 0.518
	cfg[2].Equip.Sets.SetBonusState[attributes.EM] += 58 + 44
	cfg[2].Equip.Sets.SetBonusState[attributes.Heal] += 0
	cfg[2].Equip.Sets.SetBonusState[attributes.PhyP] += 0.583

	cfg[2].ElementAttachment = attributes.NoElement

	//枫原万叶
	cfg[3].Index = 3
	cfg[3].Base = info.CharacterBase{
		CharName:  "kazuha",
		Element:   attributes.Anemo,
		Level:     90,
		MaxLevel:  90,
		HP:        kazuha.Kazuha_hp,
		Atk:       kazuha.Kazuha_atk,
		Def:       kazuha.Kazuha_def,
		Ascension: 6,
		Cons:      2,
	}
	cfg[3].Weapon = info.WeaponProfile{
		Name:     "freedom",
		Refine:   1,
		Level:    90,
		MaxLevel: 90,
		Class:    "sword",
	}
	cfg[3].Talents = info.TalentProfile{
		Attack: 8,
		Skill:  9,
		Burst:  9,
	}
	cfg[3].Equip.Sets = info.Sets{
		SetNum:  [2]int{5},
		SetName: [2]string{"viridescent"},
	}
	cfg[3].Equip.Sets.SetBonusState[attributes.ATK] += 29 + 311 + 29
	cfg[3].Equip.Sets.SetBonusState[attributes.ATKP] += 0.099 + 0.041 + 0.087
	cfg[3].Equip.Sets.SetBonusState[attributes.DEF] += 53 + 42 + 19
	cfg[3].Equip.Sets.SetBonusState[attributes.DEFP] += 0
	cfg[3].Equip.Sets.SetBonusState[attributes.HP] += 4780 + 538 + 538 + 747
	cfg[3].Equip.Sets.SetBonusState[attributes.HPP] += 0.087
	cfg[3].Equip.Sets.SetBonusState[attributes.CD] += 0
	cfg[3].Equip.Sets.SetBonusState[attributes.CR] += 0.031
	cfg[3].Equip.Sets.SetBonusState[attributes.ER] += 0.052 + 0.097 + 0.518 + 0.097 + 0.11
	cfg[3].Equip.Sets.SetBonusState[attributes.EM] += 68 + 65 + 40 + 187 + 187

	cfg[3].ElementAttachment = attributes.NoElement

	//初始化技能帧数，全是-1
	for i := range cfg {
		cfg[i].SkillFrame = -1
		cfg[i].SkillGroupFrames = make([]int, 5)
		cfg[i].BurstFrame = -1
		cfg[i].BurstGroupFrames = make([]int, 5)
		for j := range cfg[i].SkillGroupFrames {
			cfg[i].SkillGroupFrames[j] = -1 // 将每个元素设置为 -1
			cfg[i].BurstGroupFrames[j] = -1
		}
	}
	return cfg
}

func Readenemy() enemy.Enemy {
	nowEnemy := enemy.Enemy{
		Name:    "enemy1",
		Level:   103,
		Hp:      524300,
		Element: attributes.NoElement, //身上附着元素
		Res:     make([]float64, attributes.EndEleType),
	}
	for i := 0; i < int(attributes.EndEleType); i++ {
		nowEnemy.Res[i] = 0.1
	}
	return nowEnemy
}

// findCharacterIndex 查找 CharName 在 cfg.Characters 中的索引
func findCharacterIndex(cfg [4]character3.CharWrapper, CharName string) int {
	for i := range cfg {
		if cfg[i].Base.CharName == CharName {
			return i
		}
	}
	return -1 // 如果没有找到，返回 -1
}

func Actionqueue(actions player.Actions, cfg *[4]character3.CharWrapper, enemy *enemy.Enemy, charDamage *[4]float64, SwirlDamage *float64) (float64, float64) {
	// 读出角色动作列表，总结所有攻击信息，并直接计算伤害
	var allAttack float64
	var allExpectationDamage float64
	var dotAction []combat.AttackInfo // 还有dot要算的动作
	var attacks []combat.AttackInfo   // 动作队列
	var preLen, afterLen int          //这里为了记录新来了多少dot动作
	framei := -1
	ActionNum := 0
	//加一下队伍的元素共鸣
	modifier.ElementResonance(cfg)
	for {
		framei++
		//if framei == 140 {
		//	fmt.Println("debug")
		//}
		//判断动作结束
		if len(attacks) != 0 && framei > attacks[len(attacks)-1].Endframe && ActionNum == len(actions.Actions) {
			breakTag := true
			for i := range dotAction {
				if framei <= dotAction[i].AttackFrame {
					breakTag = false
				}
			}
			if breakTag == true {
				break
			}
		}
		var charindex int
		// 对怪物的计时器进行更新
		attributes.UpdateTimer(&enemy.ElementAttach)
		// 这里处理怪物身上元素的衰减
		reactions.ElementDecay(enemy)
		// 这里检查所有需要逐帧检查的dot，比如万叶大招的染色
		DotByFrame(&dotAction, cfg, enemy, framei)
		// 没动作了就直接读进来
		addActionTag := false
		if len(attacks) == 0 {
			addActionTag = true
		}
		//找一下动作队列里最后一个不是dot的动作，并以此寻找下一个动作的开始时间
		var lastActionIndex int
		for i := len(attacks) - 1; i >= 0; i-- {
			if strings.Contains(attacks[i].Abil, "Action") {
				lastActionIndex = i
				break
			}
		}
		if len(attacks) != 0 && framei == attacks[lastActionIndex].Endframe {
			addActionTag = true
		}
		if addActionTag == true && ActionNum < len(actions.Actions) {
			CharName := actions.Actions[ActionNum][0]
			//找到角色索引
			charindex = findCharacterIndex(*cfg, CharName)
			ActionName := actions.Actions[ActionNum][1]
			var c character3.Character
			c = find.FindCharC(CharName)
			//if framei == 340 {
			//	fmt.Println("debug")
			//}
			preLen = len(dotAction)
			ais := player.UseAbility(c, ActionName, cfg, charindex, actions, framei, enemy, &dotAction)
			afterLen = len(dotAction)
			player.UseDotElementAttach(c, &dotAction, enemy, CharName, afterLen-preLen)
			//这里搞一个全加进来之后的版本，方便在里面求动作组最后一位的start和end
			attacksAll := append(attacks, ais...)
			//我假设这里每一个多段的普攻都没有mod添加。那么我可以在两个多段普攻里直接查询dot伤害并计算攻击伤害。目前思考是没问题的，不知道以后会不会有角色涉及到单次多段普攻添加不同mod的情况，这里先存疑
			for aiIndex := range ais {
				lastAiIndex := len(ais) - 1
				player.UseElementAttach(c, &ais, enemy, aiIndex)
				ai := ais[aiIndex]

				//如果是多段里的第一个小动作，那么startframe是attacks里上一个动作endframe+1，若也是attacks里第一个则为0
				if aiIndex == 0 {
					if len(attacks) == 0 {
						ai.Startframe = 0
					} else {
						ai.Startframe = attacks[len(attacks)-1].Endframe + actions.Delay[len(attacks)-1]
						ais[0].Startframe = attacks[len(attacks)-1].Endframe + actions.Delay[len(attacks)-1]
					}
					//如果是多段里的靠后的动作，那么startframe直接是前面动作的endframe+1，这里没做到！！！因为这个时候还没有endframe！！！
				} else {
					ai.Startframe = ais[aiIndex-1].Endframe + 1
				}
				//如果是动作组里的动作的前几个，那么直接定义其endframe为两个动作attackframe的中点
				//这里我单独处理单次多段普攻的情况，直接在单次多段普攻的第一个小动作的时候，就把所有段的普攻的attackframe都记录好，然后每次求endframe
				//如果不是单次多段普攻，那还是正常处理attackframe和endframe
				if lastAiIndex > 0 && aiIndex == 0 {
					for i := range ais {
						if i == 0 {
							ai.AttackFrame = c.FindAttackFrame(ais[i]) + ai.Startframe
							ais[i].AttackFrame = c.FindAttackFrame(ais[i]) + ai.Startframe
						} else {
							ais[i].AttackFrame = c.FindAttackFrame(ais[i]) + ai.Startframe
						}
					}
				}
				if aiIndex < lastAiIndex {
					ai.Endframe = (ai.AttackFrame + ais[aiIndex+1].AttackFrame) / 2
					for i := range ais {
						if i != len(ais)-1 {
							ais[i].Endframe = (ais[i].AttackFrame + ais[i+1].AttackFrame) / 2
						}
					}
					for i := range ais {
						if i != 0 {
							ais[i].Startframe = ais[i-1].Endframe + 1
						}
					}
				}
				if lastAiIndex == 0 {
					ai.AttackFrame = c.FindAttackFrame(ai) + ai.Startframe
				}
				//如果最后一个动作的endframe没定义，那么找一下endframe。
				//这里注意，如果有动作组，那么只找最后一个动作的endframe，其余动作的endframe在前面已经确定好了
				attacks = append(attacks, ai)
				if attacksAll[len(attacksAll)-1].Endframe == 0 && lastAiIndex > 0 {
					ais[len(ais)-1].Endframe = FindEndFrame(attacksAll, framei, actions, *cfg, ActionNum) + ais[len(ais)-1].AttackFrame
				} else if attacksAll[len(attacksAll)-1].Endframe == 0 && lastAiIndex == 0 {
					attacks[len(attacks)-1].Endframe = FindEndFrame(attacks, framei, actions, *cfg, ActionNum)
				}
				if aiIndex != lastAiIndex {
					sortedAttacks, count := CombineAndSortAttacks(dotAction, attacks, framei)
					//这里直接把attackframe传进去，不然不识别
					CheckAttackEvent(sortedAttacks, sortedAttacks[len(sortedAttacks)-1].AttackFrame, actions, cfg, &allAttack, enemy, count, &allExpectationDamage, charDamage, SwirlDamage)
				}
			}
			ActionNum++
		}
		//这里加入dot的动作，只加入那些已经发生了的dot
		sortedAttacks, count := CombineAndSortAttacks(dotAction, attacks, framei)
		//对于动作检查是不是要加modifier
		//if framei == 339 {
		//	fmt.Println("debug")
		//}
		CheckModEvent(&sortedAttacks, framei, actions, cfg, enemy, count, ActionNum-1)
		player.ApllyLockDot(cfg, &dotAction, charindex, afterLen-preLen, framei)
		CheckAttackEvent(sortedAttacks, framei, actions, cfg, &allAttack, enemy, count, &allExpectationDamage, charDamage, SwirlDamage)
	}
	return allAttack, allExpectationDamage
}

func DotByFrame(dotAction *[]combat.AttackInfo, cfg *[4]character3.CharWrapper, enemy *enemy.Enemy, frame int) {
	for i := range *dotAction {
		if (*dotAction)[i].Abil == "kazuhaDotBurstAttach" && (*dotAction)[i].Element == attributes.Anemo {
			//判断一下枫原万叶的染色问题，改变那些还没发生的，未染色的dot
			kazuhaElement := kazuha.CheckBurstElement(cfg, enemy)
			kazuha.ChangeDotColor(&(*dotAction)[i], kazuhaElement, cfg)
		} else if (*dotAction)[i].Abil == "kazuhaDotBurst" && (enemy.Element == attributes.Pyro || enemy.Element == attributes.Cryo || enemy.Element == attributes.Electro || enemy.Element == attributes.Hydro) && (*dotAction)[i].AttackFrame == frame {
			//每次扩散都触发枫原万叶的增伤被动
			charindex := 5
			for j := range cfg {
				if cfg[j].Base.CharName == "kazuha" {
					charindex = j
				}
			}
			kazuha.Ascension4(enemy, cfg, charindex, frame, (*dotAction)[i])
		}
	}
}

func CombineAndSortAttacks(dotAction []combat.AttackInfo, attacks []combat.AttackInfo, frame int) ([]combat.AttackInfo, int) {
	// 合并两个切片
	for i := range dotAction {
		if dotAction[i].AttackFrame == frame {
			attacks = append(attacks, dotAction[i])
		}
	}
	//看看有没有attack跟dot的出伤重叠的，重叠了要处理一下
	count := 1
	for i := len(attacks) - 1; i > 0; i-- {
		if attacks[i].AttackFrame == attacks[i-1].AttackFrame {
			count++
		} else {
			break
		}
	}

	return attacks, count
}

func CheckModEvent(attacks *[]combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, enemy *enemy.Enemy, count int, actionNum int) {
	var newAi []combat.AttackInfo
	for counti := 0; counti < count; counti++ {
		*attacks = (*attacks)[:len(*attacks)-counti]
		charIndex := (*attacks)[len(*attacks)-1].ActorIndex
		CharName := actions.Characters[charIndex]
		c := find.FindCharC(CharName)
		//每一帧都检查武器圣遗物状态
		modifier.Checkmodifier(attacks, framei, actions, cfg, enemy)
		// 先判断动作开始，实现动作带来的modifier
		if framei == (*attacks)[len(*attacks)-1].Startframe {
			player.ApplyCharMod(c, *attacks, framei, actions, cfg, enemy, actionNum)
		}
		//记录待会会删掉的ai，最后再给加回去
		newAi = append(newAi, (*attacks)[len(*attacks)-1])
	}
	//给删掉的再加回来
	for i := len(newAi) - 1; i >= 0; i-- {
		if i != len(newAi)-1 {
			*attacks = append(*attacks, newAi[i])
		}
	}
}

func CheckAttackEvent(attacks []combat.AttackInfo, framei int, actions player.Actions, cfg *[4]character3.CharWrapper, allAttack *float64, enemy *enemy.Enemy, count int, allExpectationDamage *float64, charDamage *[4]float64, SwirlDamage *float64) {
	var index int
	for counti := 0; counti < count; counti++ {
		attacks = attacks[:len(attacks)-counti]
		if framei == attacks[len(attacks)-1].AttackFrame {
			for charindex := range cfg {
				if strings.Contains(attacks[len(attacks)-1].Abil, cfg[charindex].Base.CharName) {
					index = charindex
				}
			}
			fmt.Println("...........................")
			print("第", framei, "帧,", attacks[len(attacks)-1].Abil, ",")
			damageBuffer, ExpectationDamage, SwirlDamage1 := damage.Damage(attacks, actions, cfg, enemy, count, framei)
			(*enemy).Hp -= damageBuffer
			*allAttack += damageBuffer
			*allExpectationDamage += ExpectationDamage
			(*charDamage)[index] += ExpectationDamage
			*SwirlDamage += SwirlDamage1
		}
	}
}

func CheckDotEvent(dotAction []combat.AttackInfo) {

}

// 输出动作信息（combat.AttackInfo）
func printStruct(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		switch field.Kind() {
		case reflect.String:
			log.Printf("%s: %s\n", fieldName, field.String())
		case reflect.Float64:
			log.Printf("%s: %.2f\n", fieldName, field.Float())
		case reflect.Int:
			log.Printf("%s: %d\n", fieldName, field.Int())
		case reflect.Bool:
			log.Printf("%s: %t\n", fieldName, field.Bool())
		default:
			log.Printf("%s: %v\n", fieldName, field.Interface())
		}
	}
}

func InputStates(cfg *[4]character3.CharWrapper) {
	//读角色基础信息
	update.UpdateCharacterStatus(cfg)
	//读武器词条
	update.UpdateWeaponStatus(cfg)
	//读圣遗物词条
	update.UpdateArtifactStatus(cfg)
}

// 找的是abil动作的起始frame
func FindStartFrame(attacks []combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	if len(attacks) == 0 {
		return 0
	} else {
		actorpre := cfg[attacks[len(attacks)-1].ActorIndex].Base.CharName
		var c character3.Character
		c = find.FindCharC(actorpre)
		//找到上一个动作用的frames
		return c.FindFrame(attacks[len(attacks)-1], ActorIndex, Abil, cfg) + attacks[len(attacks)-1].Startframe
	}
}

func FindLastStartFrame(attacks []combat.AttackInfo, ActorIndex int, Abil string, cfg [4]character3.CharWrapper) int {
	if len(attacks) == 0 {
		return 0
	} else {
		actorpre := cfg[attacks[len(attacks)-1].ActorIndex].Base.CharName
		var c character3.Character
		c = find.FindCharC(actorpre)
		//找到上一个动作用的frames
		return c.FindFrame(attacks[len(attacks)-1], ActorIndex, Abil, cfg) + attacks[len(attacks)-1].Startframe
	}
}

func FindEndFrame(attacks []combat.AttackInfo, framei int, actions player.Actions, cfg [4]character3.CharWrapper, ActionNum int) int {
	var endFrame int

	if len(actions.Actions) > ActionNum+1 {
		CharName := actions.Actions[ActionNum][0]
		//找到角色索引
		charIndex := findCharacterIndex(cfg, CharName)
		endFrame = FindStartFrame(attacks, charIndex, CharName+actions.Actions[ActionNum+1][1], cfg)
	} else {
		CharName := actions.Actions[ActionNum][0]
		//找到角色索引
		charIndex := findCharacterIndex(cfg, CharName)
		endFrame = FindLastStartFrame(attacks, charIndex, CharName+actions.Actions[ActionNum][1], cfg)
	}
	return endFrame
}

func main() {
	// 接着写
	//file, err := os.OpenFile("RunLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 覆盖
	file, err := os.OpenFile("RunLog.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// 设置日志输出目的地为文件
	log.SetOutput(file)

	log.Println("读入角色配置……")
	cfg := Readcfg()
	//log.Println("角色配置：", cfg)

	log.Println("读入动作序列……")
	actions := Readactions()
	//log.Println("动作序列：", actions)

	log.Println("读入怪物……")
	targetenemy := Readenemy()
	//log.Println("怪物：", targetenemy)

	log.Println("更新角色配置……")
	InputStates(&cfg)
	//log.Println("看一下states")
	display.PrintStates(cfg, 3)

	log.Println("录入配置……")
	//这里简直是屎山，但是我不知道怎么优化代码结构了，脑子瓦特了
	var charDamage [4]float64
	var SwirlDamage float64
	totaldamage, allExpectationDamage := Actionqueue(actions, &cfg, &targetenemy, &charDamage, &SwirlDamage)

	//AllAttackEvent = append(AllAttackEvent, ae)
	//totaldamage, _ := CalcByFrame(&AllAttackEvent, cfg, targetenemy, states)
	fmt.Println("各个角色伤害：", charDamage)
	fmt.Println("扩散伤害：", SwirlDamage)
	fmt.Println("totaldamage: ", totaldamage)
	fmt.Println("ExpectationDamage", allExpectationDamage)
	fmt.Println("Done!")
}
