package main

import (
	"bufio"
	"dps/src/characters/raiden"
	"dps/src/core/attribute"
	character2 "dps/src/core/character"
	"dps/src/core/choose"
	"fmt"
	"os"
)

const TEAMNUM = 4

type input_queue struct {
	character      []string
	constellation  []int
	talent         [][]int
	artifact       []string
	artifact_bonus [][]float64
	action         string
	enemy          []string
}

func get_input() ([]string, []int, [][]int, []string, []string, [][]float64, string) {
	//默认四人队
	character := make([]string, 0)
	constellation := make([]int, 0)
	talent := make([][]int, 0)
	artifact := make([]string, 0)
	weapon := make([]string, 0)
	artifact_bonus := make([][]float64, 0)
	var action string

	//读入角色
	for i := 0; i < TEAMNUM; i++ {
		//创建一个reader对象
		reader := bufio.NewReader(os.Stdin)

		res, err := reader.ReadBytes('\n')
		if err != nil {
			panic("角色输入出错")
		} else {
			character = append(character, string(res))
		}
	}
	//读入命座
	for i := 0; i < TEAMNUM; i++ {
		var res int
		_, err := fmt.Scan(&res)
		if err != nil {
			panic("读入命座出错")
		} else {
			constellation = append(constellation, res)
		}
	}
	//读入天赋
	for i := 0; i < TEAMNUM; i++ {
		var res []int
		var tmp int
		for j := 0; j < 3; j++ {
			_, err := fmt.Scan(&tmp)
			if err != nil {
				panic("读入天赋出错")
			} else {
				res = append(res, tmp)
			}
		}
		talent = append(talent, res)
	}
	//读入圣遗物
	for i := 0; i < TEAMNUM; i++ {
		var res string
		_, err := fmt.Scan(&res)
		if err != nil {
			panic("读入圣遗物出错")
		} else {
			artifact = append(artifact, res)
		}
	}
	//读入武器
	for i := 0; i < TEAMNUM; i++ {
		var res string
		_, err := fmt.Scan(&res)
		if err != nil {
			panic("读入武器出错")
		} else {
			weapon = append(weapon, res)
		}
	}
	//读入圣遗物面板
	for i := 0; i < TEAMNUM; i++ {
		var res []float64
		var tmp float64
		for j := 0; j < int(attribute.End_artifact_attribute); j++ {
			//for j := 0; j < 3; j++ {
			_, err := fmt.Scan(&tmp)
			if err != nil {
				panic("读入圣遗物面板出错")
			} else {
				res = append(res, tmp)
			}
		}
		artifact_bonus = append(artifact_bonus, res)
	}
	//读动作序列，如：1e2q3q3e4q4e2e1q1a1a%
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadBytes('%')
	if err != nil {
		panic("读入动作序列出错")
	} else {
		action = string(res)
	}
	return character, constellation, talent, artifact, weapon, artifact_bonus, action
	//fmt.Println(character, constellation, talent, artifact, weapon, artifact_bonus, action)
}

// 将读入的信息写入角色信息
func deploy(character []string, constellation []int, talent [][]int, artifact []string, weapon []string, artifact_bonus [][]float64) (*character2.Char, *character2.Char, *character2.Char, *character2.Char) {
	var char1, char2, char3, char4 *character2.Char
	char1 = new(character2.Char)
	char2 = new(character2.Char)
	char3 = new(character2.Char)
	char4 = new(character2.Char)
	deploy_one(character[0], constellation[0], talent[0], artifact[0], weapon[0], artifact_bonus[0], char1)
	deploy_one(character[1], constellation[1], talent[1], artifact[1], weapon[1], artifact_bonus[1], char2)
	deploy_one(character[2], constellation[2], talent[2], artifact[2], weapon[2], artifact_bonus[2], char3)
	deploy_one(character[3], constellation[3], talent[3], artifact[3], weapon[3], artifact_bonus[3], char4)
	return char1, char2, char3, char4
}

func deploy_one(character string, constellation int, talent []int, artifact string, weapon string, artifact_bonus []float64, char *character2.Char) {
	char.Name = character
	char.Constellation = constellation
	char.Talent = talent
	char.Artifact = artifact
	//部署武器词条加成，角色突破加成
	deploy_weapon(char, weapon)
	//部署圣遗物面板加成
	deploy_artifact_bonus(char, artifact_bonus)
}

func deploy_artifact_bonus(char *character2.Char, artifact_bonus []float64) {
	char.Attribute.Max_hp = char.Attribute.Max_hp + raiden.Raiden_hp*artifact_bonus[attribute.Hp_percent] + artifact_bonus[attribute.Hp]
	char.Attribute.Atk = char.Attribute.Atk + char.Attribute.Atk*artifact_bonus[attribute.Atk_percent] + artifact_bonus[attribute.Atk]
	char.Attribute.Def = char.Attribute.Def + raiden.Raiden_def*artifact_bonus[attribute.Def_percent] + artifact_bonus[attribute.Def]
	char.Attribute.Elemental_mastery = int(artifact_bonus[attribute.Elemental_mastery])
	char.Attribute.Crit_rate = 0.05 + artifact_bonus[attribute.Crit_rate]
	char.Attribute.Crit_dmg = 0.50 + artifact_bonus[attribute.Crit_damage]
	char.Attribute.Energy_recharge = 1 + artifact_bonus[attribute.Energy_recharge]
	char.Attribute.Pyro_dmg = artifact_bonus[attribute.Pyro_dmg]
	char.Attribute.Hydro_dmg = artifact_bonus[attribute.Hydro_dmg]
	char.Attribute.Dendro_dmg = artifact_bonus[attribute.Dendro_dmg]
	char.Attribute.Electro_dmg = artifact_bonus[attribute.Electro_dmg]
	char.Attribute.Anemo_dmg = artifact_bonus[attribute.Anemo_dmg]
	char.Attribute.Cryo_dmg = artifact_bonus[attribute.Cryo_dmg]
	char.Attribute.Geo_dmg = artifact_bonus[attribute.Geo_dmg]
	char.Attribute.Phisycal_dmg = artifact_bonus[attribute.Phisycal_dmg]
}

func deploy_weapon(char *character2.Char, weapon_name string) {
	var weapon *attribute.Weapon_attri
	var character *character2.Base_attribute
	var breakthrough string
	var breakthrough_value float64
	weapon = choose.Choose_weapon(weapon_name)
	character, breakthrough, breakthrough_value = choose.Choose_character(char.Name)
	char.Attribute.Atk = character.Atk + weapon.Base_atk
	char.Attribute.Atk = char.Attribute.Atk * (1 + weapon.Atk_percent)
	char.Attribute.Max_hp = character.Hp * (1 + weapon.Hp_percent)
	char.Attribute.Def = character.Def * (1 + weapon.Def_percent)
	char.Attribute.Elemental_mastery = weapon.Elemental_mastery
	char.Attribute.Energy_recharge = weapon.Energy_recharge
	char.Attribute.Crit_rate = weapon.Crit_rate
	char.Attribute.Crit_dmg = weapon.Crit_damage
	//部署角色突破加成
	//TODO
	switch breakthrough {
	case "Atk_percent":
		char.Attribute.Atk = char.Attribute.Atk + (character.Atk+weapon.Base_atk)*(1+breakthrough_value)
	case "Energy_recharge":
		char.Attribute.Energy_recharge = char.Attribute.Energy_recharge + breakthrough_value
	case "Elemental_mastery":
		char.Attribute.Elemental_mastery = char.Attribute.Elemental_mastery + int(breakthrough_value)
	}
}

// 处理动作序列
func action_queue(action string) []string {
	//TODO
	var queue []string

	return queue
}

// 计算总伤
func total_attack(action_queue []string) float64 {
	var atk float64
	return atk
}

func main() {
	//武器默认90级，圣遗物默认满级
	character, constellation, talent, artifact, weapon, artifact_bonus, action := get_input()
	deploy(character, constellation, talent, artifact, weapon, artifact_bonus)
	action_queue := action_queue(action)
	//逐帧查看，每一帧检查角色天赋圣遗物武器效果
	total_attack := total_attack(action_queue)
	fmt.Println(total_attack)
}
