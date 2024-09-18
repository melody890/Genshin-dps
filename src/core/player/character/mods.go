package character

import (
	"dps/src/template/modifier"
)

// Add.

func AddStatus(name string, dur int, startFrame int, cfg *[4]CharWrapper, charIndex int, CD int, MaxNum int) {
	// 清理
	for i := len(cfg[charIndex].StatusMods) - 1; i >= 0; i-- {
		mod := cfg[charIndex].StatusMods[i]
		if mod.Dur != -1 && mod.StartFrame+mod.Dur < startFrame {
			cfg[charIndex].StatusMods = append(cfg[charIndex].StatusMods[:i], cfg[charIndex].StatusMods[i+1:]...)
		}
	}
	if len(cfg[charIndex].StatusMods) == 0 {
		mod := modifier.Base{
			ModName:    name,
			Dur:        dur,
			StartFrame: startFrame,
			CD:         CD,
		}
		(*cfg)[charIndex].StatusMods = append((*cfg)[charIndex].StatusMods, mod)
	} else {
		NumBuff := 0
		for _, statusMod := range cfg[charIndex].StatusMods {
			if statusMod.ModName == name {
				NumBuff++
			}
		}
		if NumBuff >= MaxNum && MaxNum != -1 {
			return
		}
		mod := modifier.Base{
			ModName:    name,
			Dur:        dur,
			StartFrame: startFrame,
			CD:         CD,
		}
		(*cfg)[charIndex].StatusMods = append((*cfg)[charIndex].StatusMods, mod)
	}
}

func RemoveStatus(name string, cfg *[4]CharWrapper, charIndex int) {
	for i := len(cfg[charIndex].Mods) - 1; i >= 0; i-- {
		if cfg[charIndex].Mods[i].Name == name {
			(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods[:i], (*cfg)[charIndex].Mods[:i+1]...)
		}
	}
}

func CDReady(name string, cfg *[4]CharWrapper, charIndex int, frame int) bool {
	for _, statusMod := range cfg[charIndex].StatusMods {
		if statusMod.ModName == name && statusMod.CD+statusMod.StartFrame > frame {
			return false
		}
	}
	return true
}
