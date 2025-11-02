package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

type Sprite struct {
	Name string `json:"name"`
	FF   string `json:"ff"`
}

var (
	sprite     []Sprite
	program    string
	programDir string
	spriteDir  string
	spriteList string = filepath.Join(programDir, "sprite.json")

	// Flags
	listAll      *bool   = flag.Bool("a", false, "List all the sprites")
	debugPath    *bool   = flag.Bool("x", false, "Enable debug mode")
	selectByName *string = flag.String("name", "", "Select sprite by name")
	selectRandom *bool   = flag.Bool("r", false, "Select random sprites")
)

func main() {
	data, err := os.ReadFile(spriteList)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &sprite); err != nil {
		panic(err)
	}

	program, err = os.Executable()
	if err != nil {
		panic(err)
	}

	programDir = filepath.Dir(program)
	spriteDir = filepath.Join(programDir, "spritesDir")

	flag.Parse()

	switch {
	case *debugPath:
		PrintDebugPath(program, programDir, spriteDir, spriteList)

	case *listAll:
		ListAllSprites(sprite)

	case *selectRandom:
		DrawByRandom(sprite)

	case *selectByName != "":
		DrawByName(*selectByName, sprite)

	default:
		println("Usage : ./go-open flag  -h : for help")
	}
}

func DrawASCII(path string) {
	if path == "" {
		log.Fatal("Invalid Path")
	}
	sprite, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("cannot open file")
		return
	}
	fmt.Println(string(sprite))
}

func DrawByName(name string, sprites []Sprite) {
	for _, sprite := range sprites {
		if name == sprite.Name {
			path := filepath.Join(spriteDir, name)
			DrawASCII(path)
			fmt.Printf("%s - %s\n", sprite.Name, sprite.FF)
			return
		}
	}
	log.Println("Invalid name or non-existent sprite:", name)
}

func DrawByRandom(sprites []Sprite) {
	if len(sprites) == 0 {
		fmt.Println("No sprites available")
		return
	}

	index := rand.Intn(len(sprites))
	sprite := sprites[index]
	path := filepath.Join(spriteDir, sprite.Name)
	DrawASCII(path)
	fmt.Printf("%s - %s\n", sprites[index].Name, sprites[index].FF)
}

func ListAllSprites(sprites []Sprite) {
	for i, sprite := range sprites {
		fmt.Printf("%d - %s - %s\n", i, sprite.Name, sprite.FF)
	}
}

func PrintDebugPath(program, programDir, spriteDir, spriteList string) {
	fmt.Println("PROGRAM:", program)
	fmt.Println("PROGRAM_DIR:", programDir)
	fmt.Println("SPRITE_DIR:", spriteDir)
	fmt.Println("SPRITE_LIST:", spriteList)
}
