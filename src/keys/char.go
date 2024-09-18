package keys

import "dps/src/core/attributes"

type Char int

const (
	NoChar Char = iota
	AetherAnemo
	LumineAnemo
	AetherGeo
	LumineGeo
	AetherElectro
	LumineElectro
	AetherDendro
	LumineDendro
	AetherHydro
	LumineHydro
	AetherPyro
	LuminePyro
	AetherCryo
	LumineCryo
	Aether
	Lumine
	TravelerDelim // delim
)

var charNames = [EndCharKeys]string{
	"",
	"aetheranemo",
	"lumineanemo",
	"aethergeo",
	"luminegeo",
	"aetherelectro",
	"lumineelectro",
	"aetherdendro",
	"luminedendro",
	"aetherhydro",
	"luminehydro",
	"aetherpyro",
	"luminepyro",
	"aethercryo",
	"luminecryo",
	"aether",
	"lumine",
	"", // delim for traveler
	"test_char_do_not_use",
}

var charPrettyName = [EndCharKeys]string{
	"Invalid",
	"Aether (Anemo)",
	"Lumine (Anemo)",
	"Aether (Geo)",
	"Lumine (Geo)",
	"Aether (Electro)",
	"Lumine (Electro)",
	"Aether (Dendro)",
	"Lumine (Dendro)",
	"Aether (Hydro)",
	"Lumine (Hydro)",
	"Aether (Pyro)",
	"Lumine (Pyro)",
	"Aether (Cryo)",
	"Lumine (Cryo)",
	"Aether",
	"Lumine",
	"Invalid",
	"!!!TEST CHAR DO NOT USE!!!",
}

var CharKeyToEle = map[Char]attributes.Element{
	AetherAnemo:      attributes.Anemo,
	LumineAnemo:      attributes.Anemo,
	AetherGeo:        attributes.Geo,
	LumineGeo:        attributes.Geo,
	AetherElectro:    attributes.Electro,
	LumineElectro:    attributes.Electro,
	AetherDendro:     attributes.Dendro,
	LumineDendro:     attributes.Dendro,
	AetherHydro:      attributes.Hydro,
	LumineHydro:      attributes.Hydro,
	AetherPyro:       attributes.Pyro,
	LuminePyro:       attributes.Pyro,
	AetherCryo:       attributes.Cryo,
	LumineCryo:       attributes.Cryo,
	TestCharDoNotUse: attributes.Geo,
}
