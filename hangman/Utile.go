package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangmanStruct struct {
	File       string
	Try        int
	Blacklist  string
	Count      int
	Word       string
	HiddenWord string
	Input      string
}

var game HangmanStruct

var try int = 10
var blacklist string
var count int = 1

func ReadArgs() string {
	/*   Fonction Pour lire le contenu du du fichiers words.txt   */

	dos, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return " Erreur "
	}
	return string(dos)

}

func Read(fil string) string { // fonctionne

	/*   Fonction Pour lire le contenu du du fichiers words.txt   */

	dos, err := os.ReadFile(fil) //ouvre un fichier
	if err != nil {              // gere le cas d erreur
		fmt.Println(err)
		return " Erreur "
	}
	return string(dos) // retourne un string du fichier
}

func RanChoice(file string) string { // fonctionne

	/*  Fonction qui crée un tableau avec tout les mots du fichier  et en choisi un au hasard */

	rand.Seed(time.Now().Unix()) // génère une seed qui chance toutes les secondes
	tab := []string{}
	ajout := ""
	for i := 0; i < len(file)-1; i++ { // Boucle dans le fichier ouvert
		ajout += string(file[i])       // ajoute a la variable ajout les lettres d'un mot
		if string(file[i+1]) == "\n" { // conditions d ajout au tableau
			tab = append(tab, ajout) // ajout au tableau
			ajout = ""               // reinitialisation de la variable ajout
			i += 1                   // permet d'eviter un backslash
		}
	}
	//fmt.Println(ajout)
	ajout += string(file[len(file)-1]) // ajoute le dernier caractere du fichier
	tab = append(tab, ajout)           // ajoute le dernier mot
	//fmt.Println(tab)
	mot := tab[rand.Intn(len(tab))] // choisi le mot aleatoirement dans le tableau
	//fmt.Println(tab)
	return mot
}

func PrintWord(mot string) string { // fonctionne

	/*  Fonction qui affiche les tiret a la place du mot et les lettres qu'il faut  */

	rand.Seed(time.Now().Unix()) // génère une seed qui change toute les secondes
	bol := true
	tab := []string{}
	motf := ""
	for i := 0; i < len(mot); i++ { // boucle pour ajoute le nombre de tiret necessaire
		tab = append(tab, "_")
	}
	if len(mot) <= 3 { // si le mot a moins de 4 caractere bol passe a faux
		bol = false
	}
	if bol { // si bol est egal a vrai on affiche le nombre de lettre necessaire
		count := (len(mot) / 2) - 1 // formule donner
		//fmt.Println(count)
		if count%10 == 0 { // gere le cas de nombre impair
			count -= 1
		}
		//fmt.Println("len=", len(mot))
		//fmt.Println(count)
		for count != 0 { // tant que count est different de zero on revele une lettre au hasard
			a := rand.Intn(len(tab))
			tab[a] = string(mot[a])
			count -= 1
		}

	}
	for j := 0; j < len(tab); j++ { // ajoute a motf le mot avec les lettres affichée
		motf += tab[j]
	}

	return motf
}

func PrintAttempt(try int) { // fonctionne
	/*  Fonction D'affichage du nombre d'essai restant et eneve une chance si une erreur est faite  */
	if game.Count == 1 {
		fmt.Println("Good Luck, you have", game.Try, "attempt")
		game.Count += 1
	} else {
		fmt.Println("You have", game.Try, "attempt")
	}
	fmt.Println("\n")
	return
}

func Input() string {
	/*   Fonction pour prendre un input de lettre ou de mot    */
	fin := ""
	fmt.Println("Choisi une lettre UwU ")
	fmt.Scan(&fin) // le & sert a acceder la memoire de fin pour avoir l input du user
	//blacklist += fin
	return fin
}

func Compare(mot, lettre string) (bool, []int) { // fonctionne
	multiIndex := []int{}
	for i := 0; i < len(mot); i++ { // boucle dans le mot
		//fmt.Println("mot =", mot, "lettre =", lettre)
		if lettre == string(mot[i]) { // teste si la lettre est contenu dans le mot pour l ajouter ou non
			multiIndex = append(multiIndex, i)
		} else {
			continue
		}
	}
	if len(multiIndex) == 0 { // si il n'y a aucune lettre on return false sinon vrai avec un tableau contenant tout les index des lettres
		fmt.Print("\033[H\033[2J") // permet un affichange propre en effacant tout ce qu'il y a au dessus
		return false, multiIndex
	}
	fmt.Print("\033[H\033[2J")
	return true, multiIndex

}

func AjoutLetter(addLetter, mot string, index []int) string {
	motComplet := ""
	tab := []string{}
	//fmt.Println("user =", addLetter, "mot= ", mot)
	//fmt.Println("tab[index]=", tab[index])
	//fmt.Println(index)
	for _, j := range mot { // boucle dans le mot pour ajouter au tableau le mot
		tab = append(tab, string(j))
	}
	//fmt.Println(tab)
	for _, indice := range index { // affiche la lettre au bon emplacement
		tab[int(indice)] = addLetter
	}

	//fmt.Println(tab)
	for _, i := range tab { // permet d avoir un e variable contenant le mot avec les bonnes lettre afficher
		motComplet += string(i)
	}
	//fmt.Println(string(mot[index]))
	return motComplet
}

func IsBlacklisted(letter string, blacklist string) bool {
	/*   Fonction qui verifie si la lettre est deja utilier ou non     */
	for _, j := range blacklist { // boucle dans la variable blaclist et verifie si elle est deja utiliser
		if letter == string(j) {
			return true
		}
	}
	return false
}

func PrintHang(try int) {
	/*   Fonction qui affiche le pendu suivant le nombre d essai restant   */
	End := 69
	file := Read("./static/hangman.txt")
	for {
		if try == 10 {
			break
		}
		if try == 9 {
			for Start := 0; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 8 {
			End += 71
			for Start := 69; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 7 {
			End += 142
			for Start := 140; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 6 {
			End += 213
			for Start := 211; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 5 {
			End += 284
			for Start := 282; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 4 {
			End += 355
			for Start := 353; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 3 {
			End += 426
			for Start := 424; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 2 {
			End += 497
			for Start := 495; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 1 {
			End += 568
			for Start := 566; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		} else if try == 0 {
			End += 639
			for Start := 637; Start < End; Start++ {
				fmt.Print(string(file[Start]))
			}
			fmt.Print("\n")
			break
		}
	}
}

func Welcome() {
	/*  fonction qui affiche notre message de bienvenu    */
	Start := ""
	file := Read("./static/Welcome.txt")
	for i := 0; i < len(file); i++ {
		fmt.Print(string(file[i]))
	}
	fmt.Println()

	fmt.Println("Enter Start To Begin :")
	fmt.Scan(&Start)
	if Start == "Start" || Start == "start" || Start == "START" {
		fmt.Print("\033[H\033[2J")
	} else {
		Welcome()
	}
}

func Lose() {
	/*  fonction qui affiche le message de fin si perdu   */
	file := Read("./static/Lose.txt")
	for i := 0; i < len(file); i++ {
		fmt.Print(string(file[i]))
	}
	fmt.Println()

}

func Win() {
	/*  fonction qui affiche le message de fin si gagner   */
	file := Read("./static/Win.txt")
	for i := 0; i < len(file); i++ {
		fmt.Print(string(file[i]))
	}
	fmt.Println()

}

func Init(game *HangmanStruct) {
	game.File = Read("words.txt")
	game.Try = 10
	game.Blacklist = blacklist
	game.Count = 1
	game.Word = RanChoice(game.File)
	game.HiddenWord = PrintWord(game.Word)
	game.Input = ""
	fmt.Println(game.Try)
	fmt.Println(game.Blacklist)
	fmt.Println(game.Count)
	fmt.Println(game.Word)
	fmt.Println(game.HiddenWord)
	fmt.Println(game.Input)
}

func LaunchGame() {
	//file := ReadArgs()
	Welcome()
	Init(&game)
	//file := Read("words.txt")   // affecte a file les mots contenu dans le fichier
	//motr := RanChoice(file)     // affecte a motr un mot choisi au hasard entre tout les mots du fichier
	//motPrint := PrintWord(motr) // affect a motprint le mot avec les tiret et les lettre necessaire
	//hangman.AsciiArt("abcdefghijklmnopqrstuvwxyz")
	for {
		//fmt.Print("\033[H\033[2J") essai de clear le terminal a chaque round pour plus de clairte
		PrintAttempt(game.Try)       // affiche les essaie
		fmt.Println(game.HiddenWord) // en mettre que un seul pour que ca marche
		//AsciiArt(motr)
		fmt.Println("letter already used --->", game.Blacklist) // affiche les lettres deja utiliser
		game.Input = Input()                                    // recupere ce que rentre l utilisateur
		result, index := Compare(game.Word, game.Input)         // recupere le resultat et les indice si il y en a
		game.Count += 1
		if IsBlacklisted(game.Input, game.Blacklist) { // gere le cas d'une lettre deja utilier
			fmt.Println("vous avez déjà utilisé le", game.Input)
			continue
		}
		game.Blacklist += game.Input           // ajoute la lettre a la blacklist
		if result || game.Input == game.Word { // si la lettre est dans le mot on ajoute la lettre au mot et on affiche le pendu sinon on enleve a au nombre d essai
			//fmt.Print("test1")
			game.HiddenWord = AjoutLetter(game.Input, game.HiddenWord, index)
			//fmt.Print("test2")
			PrintHang(game.Try)
			//fmt.Print("test3")
		} else {
			if len(game.Input) > 1 {
				game.Try -= 2
			} else {
				game.Try -= 1
			}

			PrintHang(game.Try)
		}
		//fmt.Println("motprint =", motPrint, "motr =", motr)
		//fmt.Println(blacklist)
		if game.Try == 0 { // condition de lose
			fmt.Print("\033[H\033[2J")
			fmt.Println(" Le mot était ====>", game.Word)
			Lose()
			break
		}
		if game.HiddenWord == game.Word || game.Input == game.Word { // condition de win
			fmt.Println(game.Word)
			Win()
			break
		}
	}
}
