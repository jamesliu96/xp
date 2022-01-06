package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/jamesliu96/xp"
)

const p = "p"
const x = "x"

func usage() {
	fmt.Fprintf(os.Stderr, "usage: xp %s\n       xp %s <scalar> <point>\n", p, x)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	directive := os.Args[1]
	if directive == p {
		priv, pub, err := xp.P()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-4s %s\n", "priv", hex.EncodeToString(priv))
		fmt.Printf("%-4s %s\n", "pub", hex.EncodeToString(pub))
	} else if directive == x {
		if len(os.Args) < 4 {
			usage()
			return
		}
		scalar, err := hex.DecodeString(os.Args[2])
		if err != nil {
			panic(err)
		}
		point, err := hex.DecodeString(os.Args[3])
		if err != nil {
			panic(err)
		}
		product, err := xp.X(scalar, point)
		if err != nil {
			panic(err)
		}
		fmt.Println(hex.EncodeToString(product))
	}
}
