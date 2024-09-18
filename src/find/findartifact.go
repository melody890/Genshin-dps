package find

import (
	"dps/src/artifacts/crimson"
	"dps/src/artifacts/emblem"
	"dps/src/artifacts/noblesse"
	"dps/src/artifacts/viridescent"
	"dps/src/core/player/artifact"
)

func FindArtifact(artifactName string) artifact.Artifact {
	var a artifact.Artifact

	switch artifactName {
	case "noblesse":
		ArtNoblesse := &noblesse.Noblesse{}
		a = ArtNoblesse
	case "crimson":
		ArtCrimson := &crimson.Crimson{}
		a = ArtCrimson
	case "emblem":
		ArtEmblem := &emblem.Emblem{}
		a = ArtEmblem
	case "viridescent":
		ArtViridescent := &viridescent.Viridescent{}
		a = ArtViridescent
	}
	return a
}
