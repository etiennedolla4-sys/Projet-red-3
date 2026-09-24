package game

type Layer struct {
	Depth       int
	Name        string
	Description string
}

func GetLayer(depth int) Layer {
	switch depth {
	case 1:
		return Layer{
			Depth:       1,
			Name:        "La lisière",
			Description: "Une zone relativement sûre où la lumière atteint encore les profondeurs.",
		}

	case 2:
		return Layer{
			Depth:       2,
			Name:        "Les ruines oubliées",
			Description: "D'anciennes ruines recouvertes par une végétation étrange.",
		}

	case 3:
		return Layer{
			Depth:       3,
			Name:        "La forêt silencieuse",
			Description: "Une immense forêt où chaque bruit semble attirer quelque chose.",
		}

	case 4:
		return Layer{
			Depth:       4,
			Name:        "Les cavernes obscures",
			Description: "Un réseau de cavernes où la lumière disparaît presque totalement.",
		}

	case 5:
		return Layer{
			Depth:       5,
			Name:        "Les profondeurs",
			Description: "Une région hostile où peu d'explorateurs reviennent indemnes.",
		}

	default:
		return Layer{
			Depth:       depth,
			Name:        "Territoire inconnu",
			Description: "Personne ne sait réellement ce qui se trouve ici.",
		}
	}
}
