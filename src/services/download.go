package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lexgalante/go.downloadpokemon/src/schemas"
)

//DownloadPokemon -> download pokemon json
func DownloadPokemon(id int) {
	url := fmt.Sprintf("%s/pokemon/%v", os.Getenv("URL_BASE"), id)

	fmt.Println("DownloadPokemon: starting download pokemon: ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on GET: ", url, " error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("DownloadPokemon: download pokemon: ", url, " result: ", resp.StatusCode)

	if resp.StatusCode != 200 {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on read body: ", url, " error: ", err.Error())
		return
	}

	var pokemon schemas.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on unmarshal: ", url, " error: ", err.Error())
		return
	}

	fmt.Println("DownloadPokemon: starting download sprite: ", pokemon.Sprites.Default)

	resp, err = http.Get(pokemon.Sprites.Default)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on GET: ", pokemon.Sprites.Default, " error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("pokemons/%s.png", pokemon.Name)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on create file: ", fileName, " error: ", err.Error())
		return
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("DownloadPokemon: unexpected error on copy content: ", fileName, " error: ", err.Error())
		return
	}

	fmt.Println("DownloadPokemon: download ", id, " success...")
}
