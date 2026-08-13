package main

import (
	"fmt"

	gonanoid "github.com/enjoyZhou/go-nanoid/v2"
)

func main() {
	// Simple usage
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom length
	id, err = gonanoid.New(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom alphabet
	id, err = gonanoid.Generate("abcdefg", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom non ascii alphabet
	id, err = gonanoid.Generate("こちんにабдежиклмнは你好喂שלום😯😪🥱😌😛äöüß", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	fmt.Printf("Generated id: %s\n", gonanoid.Must())
	fmt.Printf("Generated id: %s\n", gonanoid.MustGenerate("🚀💩🦄🤖", 4))
}
