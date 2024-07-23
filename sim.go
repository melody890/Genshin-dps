package main

import (
	"dps/src/characters"
	"dps/src/characters/bennett"
	character2 "dps/src/core/character"
	"dps/src/core/combat"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	"fmt"
)

type Actions struct {
	Characters []string
	actions    [][]string
}

type CharacterInfo struct {
	Char         character2.Base_attribute
	Breakthrough string
}

func Readactions() Actions {
	action := Actions{
		Characters: []string{"bennett"},
		actions:    [][]string{{"bennett", "ActionSkill"}},
	}
	return action
}

func Readcfg() character2.Charcfg {
	cfg := character2.Charcfg{
		Characters:     []string{"bennett"},
		Weapons:        []string{"mistsplitter"},
		Artifacts:      []string{"Artifact1", "Artifact2", "Artifact1", "Artifact1", "Artifact1"},
		CharacterLevel: 10,
		WeaponLevel:    5,
		ArtifactLevel:  []int{20, 20, 20, 20, 20},
		TalentLevel:    [][3]int{{1, 2, 2}},
	}
	return cfg
}

func Actionqueue(actions Actions, cfg character2.Charcfg) []combat.AttackInfo {
	var charbaseinfo []CharacterInfo // 初始化一个 CharacterInfo 的切片，存储各个角色的初始信息
	//var enemy enemies.Enemies        //初始化怪物结构体

	// 读出角色三围
	for _, charName := range actions.Characters {
		char, breakthrough, err := characters.InitCharacter(charName)
		if err != nil {
			fmt.Printf("Error initializing character %s: %v\n", charName, err)
			continue
		}
		charbaseinfo = append(charbaseinfo, CharacterInfo{Char: char, Breakthrough: breakthrough})
	}

	// 读出角色动作列表，总结所有攻击信息
	AllAttackInfo := make([]combat.AttackInfo, 0)
	for _, ActionContext := range actions.actions {
		CharName := ActionContext[0]
		ActionName := ActionContext[1]
		var c character3.Character

		switch CharName {
		case "bennett":
			bennettChar := &bennett.Char{}
			bennettChar.InitChar(CharName) // 初始化 Char 实例.初始化三围。
			c = bennettChar
		}
		ai := player.UseAbility(c, ActionName, cfg)
		AllAttackInfo = append(AllAttackInfo, ai)
		//fmt.Println(ai)
		//fmt.Println(CharName, ActionName)
	}
	return AllAttackInfo
}

func CalcByFrame(info []combat.AttackInfo) {
	fmt.Println(info[0])
}

func main() {
	fmt.Println("读入配置")
	cfg := Readcfg()
	fmt.Println("角色配置：", cfg)
	fmt.Println("读入动作序列")
	actions := Readactions()
	fmt.Println("动作序列：", actions)
	AllAttackInfo := Actionqueue(actions, cfg)
	CalcByFrame(AllAttackInfo)
}
