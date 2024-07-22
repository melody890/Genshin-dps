package main

import (
	"dps/src/characters"
	character2 "dps/src/core/character"
	"dps/src/core/player"
	character3 "dps/src/core/player/character"
	"fmt"
)

type Charcfg struct {
	Characters     []string
	Weapons        []string
	Artifacts      []string
	CharacterLevel int
	WeaponLevel    int
	ArtifactLevel  []int
	TalentLevel    [][3]int
}

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
		actions:    [][]string{{"1", "ActionSkill"}},
	}
	return action
}

func Readcfg() Charcfg {
	cfg := Charcfg{
		Characters:     []string{"bennett"},
		Weapons:        []string{"mistsplitter"},
		Artifacts:      []string{"Artifact1", "Artifact2", "Artifact1", "Artifact1", "Artifact1"},
		CharacterLevel: 10,
		WeaponLevel:    5,
		ArtifactLevel:  []int{3, 4},
		TalentLevel:    [][3]int{{1, 2, 2}, {1, 1, 1}, {1, 1, 1}},
	}
	return cfg
}

func FindActionQueue(actions Actions) {
	var charbaseinfo []CharacterInfo // 初始化一个 CharacterInfo 的切片，存储各个角色的初始信息
	for _, charName := range actions.Characters {
		char, breakthrough, err := characters.InitCharacter(charName)
		if err != nil {
			fmt.Printf("Error initializing character %s: %v\n", charName, err)
			continue
		}
		charbaseinfo = append(charbaseinfo, CharacterInfo{Char: char, Breakthrough: breakthrough})
	}
	for _, ActionContext := range actions.actions {
		CharIndex := ActionContext[0]
		ActionName := ActionContext[1]
		switch ActionName {
		case "ActionSkill":
			player.UseAbility(character3.Character.Skill)
		}

		fmt.Println(CharIndex, ActionName)
	}

}

func main() {
	fmt.Println("读入配置")
	cfg := Readcfg()
	fmt.Println(cfg)
	fmt.Println("读入动作序列")
	actions := Readactions()
	fmt.Println(actions)
	FindActionQueue(actions)
}
