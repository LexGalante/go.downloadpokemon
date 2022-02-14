package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/lexgalante/go.downloadpokemon/src/services"
)

func main() {
	now := time.Now()

	readDotEnv()

	fmt.Println("starting pokemon download's")

	ensureDir("pokemons")

	numberOfDownloads := getNumberOfDownloads()

	var wg sync.WaitGroup

	for i := 1; i <= numberOfDownloads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			services.DownloadPokemon(i)
		}(i)
	}

	wg.Wait()

	fmt.Println("stopped pokemon download, elapsed time: ", time.Now().Sub(now))
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

func getNumberOfDownloads() int {
	args := os.Args[1:]

	var downloads int64

	downloads, err := strconv.ParseInt(args[0], 4, 4)
	if err != nil {
		fmt.Println("cannot read arg, consider default 100 downloads")
		downloads = 100
	}

	return int(downloads)
}
