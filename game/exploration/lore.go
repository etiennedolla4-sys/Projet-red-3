package game

import (
	text "Projet-red-3/ui/text"
	"fmt"
	"math/rand"
)

func ShowEnemyLore(enemyName string) {
	switch enemyName {
	case "Slime":
		text.PrintSlow("Le slime tremble encore quelques secondes avant de se dissoudre.\n")

	case "Petit Gobelin":
		text.PrintSlow("Le petit gobelin serre une vieille pièce contre lui avant de tomber.\n")

	case "Gobelin des ruines":
		text.PrintSlow("Son équipement porte les mêmes symboles que ceux gravés sur les ruines.\n")

	case "Araignée ancienne":
		text.PrintSlow("La toile autour d'elle semble être là depuis des années.\n")

	case "Loup sauvage":
		text.PrintSlow("Son dernier hurlement se perd entre les arbres.\n")

	case "Prédateur de la forêt":
		text.PrintSlow("Même mort, sa silhouette reste difficile à distinguer dans l'obscurité.\n")

	case "Golem des cavernes":
		text.PrintSlow("Les pierres qui composent son corps cessent lentement de vibrer.\n")

	case "Chauve-souris géante":
		text.PrintSlow("Le battement de ses ailes cesse enfin de résonner dans la caverne.\n")

	case "Monstre des profondeurs":
		text.PrintSlow("Son cri mourant semble réveiller quelque chose plus bas.\n")

	case "Gardien abyssal":
		text.PrintSlow("Le gardien s'effondre lentement. Derrière lui, le passage semble enfin libre.\n")
	}
}

var ChestLore = []string{
	"Le coffre porte les armoiries d'un royaume disparu depuis des siècles.",
	"La serrure est couverte de traces de griffes.",
	"Une couche de poussière épaisse recouvre le coffre. Personne ne l'a ouvert depuis longtemps.",
	"À l'intérieur, vous trouvez des objets accompagnés d'un morceau de tissu ensanglanté.",
	"Le coffre semble avoir appartenu à un ancien explorateur.",
}

var RestLore = []string{
	"Un ancien feu de camp brûle encore faiblement.",
	"Quelqu'un a gravé sur le mur : \"Ne descendez pas plus bas.\"",
	"Pour quelques instants, les profondeurs semblent silencieuses.",
	"Des affaires abandonnées sont dispersées autour du camp.",
}

func ShowChestLore() {
	fmt.Println()
	text.PrintSlow(ChestLore[rand.Intn(len(ChestLore))])
}

func ShowRestLore() {
	fmt.Println()
	text.PrintSlow(RestLore[rand.Intn(len(RestLore))])
}
