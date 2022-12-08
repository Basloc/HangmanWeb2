package main

/*
Pour faire les struct initialiser une struct avec les valeur du prog hangman donc
1* Réadapter le prog de base pour le faire marcher avec des struct ✅
2.1* Faire le PAD du HandUnit Fnaf avec chaque bouton qui envoi une lettre et possibilité de recup la lettre
2* Créer les templates pour recup les infos
3* Faire marcher le hangman sans aucun style ou quoi
4* Faire le css/html pour passer d'un écran d'accueil à un écran de jeu
5* Pense faire des bonus genre statistique ou des ptite anim stylée
*/

import (
	"fmt"
	"hangman/hangman"
	"log"
	"net/http"
	"text/template"
)

// faire une struct ici
type HangData struct {
	Try        int
	Input      string
	Word       string
	HiddenWord string
	count      int
	Blacklist  string
}

func Home(rw http.ResponseWriter, r *http.Request, str *HangData) {
	template, err := template.ParseFiles("./home.html", "./template/header.html", "./template/footer.html", "./static/home.css")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, str)

}

func SalleDeJeu(rw http.ResponseWriter, r *http.Request, str *HangData) {
	template, err := template.ParseFiles("./salledejeu.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, str)
}

func Victoire(rw http.ResponseWriter, r *http.Request, str *HangData) {
	template, err := template.ParseFiles("./victory.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, str)
}

func Lose(rw http.ResponseWriter, r *http.Request, str *HangData) {
	template, err := template.ParseFiles("./lose.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, str)
}

func main() {
	HangPts := HangData{10, "", "", "", 1, ""}
	Pt := &HangPts

	InitialiseStruc(Pt)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r, Pt)
	})

	http.HandleFunc("/jeux", func(rw http.ResponseWriter, r *http.Request) {
		SalleDeJeu(rw, r, Pt)
	})
	http.HandleFunc("/victoire", func(rw http.ResponseWriter, r *http.Request) {
		Victoire(rw, r, Pt)
	})
	http.HandleFunc("/lose", func(rw http.ResponseWriter, r *http.Request) {
		Lose(rw, r, Pt)
	})

	http.HandleFunc("/calcul", func(rw http.ResponseWriter, r *http.Request) {

		Testing(rw, r, Pt)

		fmt.Println(Pt.Input)

		Pt.Input = ""
		Pt.Input = r.FormValue("letter")
		result, index := hangman.Compare(Pt.Word, Pt.Input)
		if hangman.IsBlacklisted(Pt.Input, Pt.Blacklist) {
			//afficher que c est deja utiliser
			fmt.Println("blacklist avant :", Pt.Blacklist)
			fmt.Println("test blacklist")

		}
		Pt.Blacklist += Pt.Input
		fmt.Println("blacklist apres :", Pt.Blacklist)
		if result || Pt.Input == Pt.Word {
			Pt.HiddenWord = hangman.AjoutLetter(Pt.Input, Pt.HiddenWord, index)
			// afficher le pendu
			fmt.Println("test affichage pendu n*", Pt.Try)
		} else {
			if len(Pt.Input) > 1 {
				Pt.Try -= 2
			} else {
				Pt.Try -= 1
			}
			// afficher pendu
			fmt.Println("test pendu si loose n* :", Pt.Try)
		}
		if Pt.Try == 0 { // condition de lose
			//fmt.Print("\033[H\033[2J")
			fmt.Println(" Le mot était ====>", Pt.Word)
			http.Redirect(rw, r, "/lose", http.StatusFound)
			fmt.Println("test loose")

		}
		if Pt.HiddenWord == Pt.Word || Pt.Input == Pt.Word { // condition de win
			fmt.Println(Pt.Word)
			//faire lancer le bruit OOOOUUUUUAAAAAIIIII de fnaf
			http.Redirect(rw, r, "/victoire", http.StatusFound)
			//break
		}
		http.Redirect(rw, r, "/jeux", http.StatusFound)

	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func Testing(rw http.ResponseWriter, r *http.Request, games *HangData) {
	template, _ := template.ParseFiles("./calcul.html")
	template.Execute(rw, games)
}

func InitialiseStruc(Pt *HangData) {
	Pt.Try = 10
	Pt.count = 1
	Pt.Input = ""
	Pt.Word = hangman.RanChoice(hangman.Read("words.txt"))
	Pt.HiddenWord = hangman.PrintWord(Pt.Word)
	Pt.Blacklist = ""

}
