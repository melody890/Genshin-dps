package info

import (
	"dps/src/core/attributes"
)

type CharacterProfile struct {
	Base    CharacterBase
	Weapon  WeaponProfile
	Talents TalentProfile
	Stats   []float64
	//StatsByLabel map[string][]float64
	Sets Sets
	//SetParams    map[keys.Set]map[string]int
	//Params       map[string]int
}

type CharacterBase struct {
	CharName string
	//Rarity    int
	Element   attributes.Element
	Level     int
	MaxLevel  int
	Ascension int //突破次数
	HP        float64
	Atk       float64
	Def       float64
	Cons      int
}
