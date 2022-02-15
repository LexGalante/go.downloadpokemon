package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lexgalante/go.downloadpokemon/src/schemas"
	"github.com/lexgalante/go.downloadpokemon/src/services"
)

const limitDownloads = 100

func main() {
	now := time.Now()

	readDotEnv()

	fmt.Println("starting pokemon download's")

	ensureDir("pokemons")

	pokemons := make(chan schemas.Pokemon, limitDownloads)

	for i := 1; i <= limitDownloads; i++ {
		go services.DownloadPokemonSpriteURL(i, pokemons)
	}

	for i := 1; i <= limitDownloads; i++ {
		go services.DownloadPokemonSpritePNG(<-pokemons)
	}

	close(pokemons)

	fmt.Println("stopped pokemon download, elapsed time:", time.Now().Sub(now))
}

func readDotEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load .env file")
	}
}

func ensureDir(dirName string) {
	os.RemoveAll(dirName)
	os.Mkdir(dirName, 0755)
}
