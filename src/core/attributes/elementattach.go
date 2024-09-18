package attributes

import "strings"

type ElementAttach struct {
	ActionName []string
	Counter    []int
	Timer      []int
}

func UpdateTimer(ElementAttach *ElementAttach) {
	for i := len(ElementAttach.Timer) - 1; i >= 0; i-- {
		if !strings.Contains(ElementAttach.ActionName[i], "Dot") {
			(*ElementAttach).Timer[i] += 1
			//到2.5秒清空该攻击的记录，相当于计数器刷新
			if (*ElementAttach).Timer[i] == 150 {
				(*ElementAttach).ActionName = append((*ElementAttach).ActionName[:i], (*ElementAttach).ActionName[i+1:]...)
				(*ElementAttach).Counter = append((*ElementAttach).Counter[:i], (*ElementAttach).Counter[i+1:]...)
				(*ElementAttach).Timer = append((*ElementAttach).Timer[:i], (*ElementAttach).Timer[i+1:]...)
			}
		}
	}
}
