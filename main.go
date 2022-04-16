package main

import (
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip39"
)

func main() {
	if os.Args[1] == "12" {
		entropy, _ := bip39.NewEntropy(128)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		fmt.Println(mnemonic)
	} else {
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		fmt.Println(mnemonic)
	}

}
