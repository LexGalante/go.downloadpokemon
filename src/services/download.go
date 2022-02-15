package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lexgalante/go.downloadpokemon/src/schemas"
)

//DownloadPokemonSpriteURL -> download pokemon json
func DownloadPokemonSpriteURL(id int, pokemons chan<- schemas.Pokemon) {
	url := fmt.Sprintf("%s/pokemon/%v", os.Getenv("URL_BASE"), id)

	fmt.Println("DownloadPokemonSpriteURL: starting download pokemon: ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("DownloadPokemonSpriteURL: unexpected error on GET: ", url, " error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("DownloadPokemonSpriteURL: download pokemon: ", url, " result: ", resp.StatusCode)

	if resp.StatusCode != 200 {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("DownloadPokemonSpriteURL: unexpected error on read body: ", url, " error: ", err.Error())
		return
	}

	var pokemon schemas.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("DownloadPokemonSpriteURL: unexpected error on unmarshal: ", url, " error: ", err.Error())
		return
	}

	fmt.Println("DownloadPokemonSpriteURL: ok pokemon: ", pokemon.Name)

	pokemons <- pokemon
}

//DownloadPokemonSpritePNG -> download pokemon png
func DownloadPokemonSpritePNG(pokemon schemas.Pokemon) {
	fmt.Println("DownloadPokemonSpritePNG: starting download sprite png:", pokemon.Sprites.Default)

	resp, err := http.Get(pokemon.Sprites.Default)
	if err != nil {
		fmt.Println("DownloadPokemonSpritePNG: unexpected error on GET:", pokemon.Sprites.Default, "error:", err.Error())
		return
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("pokemons/%s.png", pokemon.Name)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("DownloadPokemonSpritePNG: unexpected error on create file:", fileName, " error:", err.Error())
		return
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("DownloadPokemonSpritePNG: unexpected error on copy content:", fileName, "error:", err.Error())
		return
	}

	fmt.Println("DownloadPokemonSpritePNG: download", pokemon.Name, "sprite finished")
}
