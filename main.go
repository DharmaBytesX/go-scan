package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"sync"
	"time"
)

// Utilisation d'une regex pour valider une adresse IPv4
func validationIP(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(ip) // Vérifie si l'adresse IP correspond à la regex
}

// Fonction pour scanner un port spécifique
func scanPort(protocol, port, ip string, wg *sync.WaitGroup) {
	defer wg.Done() // Décrémenter le compteur du WaitGroup quand la fonction se termine

	address := ip + ":" + port

	var conn net.Conn
	var err error

	conn, err = net.Dial(protocol, address) // Tente de se connecter à l'adresse IP et au port spécifié

	if err != nil { // Si une erreur survient, le port est considéré comme fermé
		fmt.Printf("Port %s/%s FERMÉ\n", port, protocol)
		return
	}
	defer conn.Close()                                // Fermer la connexion quand le scan est terminé
	fmt.Printf("Port %s/%s OUVERT\n", port, protocol) // Si la connexion réussit, le port est ouvert
}

func main() {
	// Commencer à mesurer le temps d'exécution
	start := time.Now()

	banner := `
   ____  ___       ____   ____    _    _   _ 
  / ___|/ _ \     / ___| / ___|  / \  | \ | |
 | |  _| | | |____\___ \| |     / _ \ |  \| |
 | |_| | |_| |_____|__) | |___ / ___ \| |\  |
  \____|\___/     |____/ \____/_/   \_\_| \_|										   
`

	fmt.Println(banner)
	fmt.Println("@simonaurascussel\n")

	// Récupérer les arguments de la ligne de commande, sans le nom du programme
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("Veuillez fournir une adresse IP à scanner.")
		os.Exit(1)
	}

	// Récupérer l'adresse IP fournie en argument
	var ip string = argsWithoutProg[0]
	// Valider l'adresse IP à l'aide d'une regex
	var ipValide bool = validationIP(ip)

	if !ipValide { // Si l'adresse IP n'est pas valide, afficher un message d'erreur et quitter le programme
		fmt.Println("La valeur [" + ip + "] n'est pas une adresse IPv4 valide.")
		os.Exit(3)
	}

	// Déclarer la carte des ports
	ports := make(map[string][]string)

	// Liste des ports TCP à scanner
	ports["tcp"] = []string{
		"20", "21", "22", "23", "25", "53", "80", "110", "143", "443", "465", "587", "993", "995",
		"1433", "1521", "1723", "3306", "3389", "5432", "5900", "6379", "8000", "8080", "8443", "9000",
		"9200", "11211", "27017",
	}

	// Protocole à utiliser pour le scan des ports (TCP)
	protocols := []string{"tcp"}

	// Création d'un WaitGroup pour synchroniser les goroutines
	var wg sync.WaitGroup

	fmt.Printf("Démarrage du scan de ports TCP vers l'adresse IPV4 %s \n\n", ip)

	for _, protocol := range protocols {
		for _, port := range ports[protocol] {
			wg.Add(1)                            // Incrémenter le compteur du WaitGroup
			go scanPort(protocol, port, ip, &wg) // Exécuter scanPort comme une goroutine
		}
	}

	wg.Wait() // Attendre que toutes les goroutines soient terminées

	// Mesurer la durée d'exécution et afficher la durée
	duration := time.Since(start)
	fmt.Printf("Temps d'exécution : %s\n", duration)
}
