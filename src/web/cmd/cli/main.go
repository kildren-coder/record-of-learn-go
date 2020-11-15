package main

import (
	"fmt"
	"log"
	"os"
	poker "record-of-learn-go/src/web"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
