package enemy

import (
	"dps/src/core/attributes"
	"dps/src/template/modifier"
)

type Enemy struct {
	Name  string
	Level int
	Hp    float64

	Element         attributes.Element       //身上附着元素
	ElementQuantity float64                  //残余元素量
	ElementType     attributes.ElementType   // 初始附着时候的元素强度，会决定元素衰减的速率
	ElementAttach   attributes.ElementAttach // 计时器计数器，用来判断是否有元素附着

	Res     []float64
	resists map[attributes.Element]float64

	// mods
	EnemyMod []modifier.EnemyMod
}
