package main

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/tscholl2/isitrandom"
)

//
// func TestStraightDeckRNG(t *testing.T) {
// 	rand.Seed(42)
// 	isitrandom.TestRNG(t, bytes.NewBuffer(BoolToByte(Sign(Diff(GenerateDeck())))))
// }

func TestRandomlyShuffledDeckRNG(t *testing.T) {
	rand.Seed(15)
	shuffled := EmpiricallyShuffle(3)
	//a2 := Sign(Diff(shuffled))
	a2 := Sign(SplitDiff(shuffled))
	isitrandom.TestRNG(t, bytes.NewBuffer(BoolToByte(a2)))
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
