package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func main() {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(mnemonic)
}
