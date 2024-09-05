# go-scan

Un scanner de ports TCP développé en Go.

## Description

`go-scan` est un outil de ligne de commande pour scanner les ports TCP d'une adresse IP donnée. Conçu pour être simple et efficace, il utilise la concurrence pour améliorer les performances et réduire le temps de scan.

## Performance

Le scanner utilise des goroutines pour paralléliser les opérations de scan, permettant de vérifier les ports de manière rapide, en moins de 2 secondes.

## Installation

Assurez-vous d'avoir Go installé sur votre système. Vous pouvez télécharger Go depuis [le site officiel](https://golang.org/dl/).

Clonez ce dépôt sur votre machine :

```bash
git clone https://github.com/votre-utilisateur/go-scan.git
cd go-scan
```

## Utilisation

Pour exécuter le scanner, utilisez la commande suivante :

```bash
go run main.go @adresseIPV4
```

Remplacez `@adresseIPV4` par l'adresse IP de la cible que vous souhaitez scanner.

## Exemples

Pour scanner les ports de l'adresse IP `192.168.1.1`, exécutez :

```bash
go run main.go 192.168.1.1
```
