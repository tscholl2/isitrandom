package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"github.com/tscholl2/isitrandom"
)

func TestShuffleDeckRNG(t *testing.T) {
	_, deck := GenerateDeck()
	isitrandom.TestRNG(t, bytes.NewBuffer(deck))
}

// func main() {
// 	fmt.Println("Hello, playground")
//
// 	deck := GenerateDeck()
// 	WriteBits(deck)
// 	shuffledDeck := ShuffleDeckRandomly(deck)
// 	WriteBits(shuffledDeck)
//
// }

func ShuffleDeckRandomly(src []int) []int {
	// Randomly shuffle deck
	rand.Seed(42)
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func GenerateDeck() ([]int, []byte) {
	// Generate 4 sets of 13 numbers that are ordinally dissimilar
	deck := make([]int, 52)
	current := 0
	start := int(-128)
	for i := start; i < start+13; i++ {
		deck[current] = i
		current++
	}
	start = int(-63)
	for i := start; i < start+13; i++ {
		deck[current] = i
		current++
	}
	start = int(50)
	for i := start; i < start+13; i++ {
		deck[current] = i
		current++
	}
	start = int(116)
	for i := start; i < start+13; i++ {
		deck[current] = i
		current++
	}

	buf := new(bytes.Buffer)
	for _, val := range deck {
		err := binary.Write(buf, binary.LittleEndian, int8(val))
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}

	return deck, buf.Bytes()
}

// func WriteBits(deck []int) {
// 	buf := new(bytes.Buffer)
//
// 	for _, val := range deck {
// 		err := binary.Write(buf, binary.LittleEndian, int8(val))
// 		if err != nil {
// 			fmt.Println("binary.Write failed:", err)
// 		}
// 	}
//
// 	fmt.Printf("% x\n", buf.Bytes())
//
// 	data := buf.Bytes()
// 	bitString := ""
// 	r := isitrandom.New(bytes.NewBuffer(data))
// 	for i := 0; i < len(data); i++ {
// 		for j := 0; j < 8; j++ {
// 			bit, _ := r.ReadBit()
// 			if bit {
// 				bitString += "1"
// 			} else {
// 				bitString += "0"
// 			}
// 		}
// 	}
// 	fmt.Println(bitString)
// }
