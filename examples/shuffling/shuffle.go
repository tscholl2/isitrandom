package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
)

// EmpiricalShuffling is the Empirical probability of riffling and getting
// packets with either                 1       2        3      4       5    cards
var EmpiricalShuffling = []float64{0.5117, 0.8729, 0.9565, 0.9799, 1.000}

func main() {
	rand.Seed(43)
	randomDecks := []byte{}
	for i := 0; i < 100; i++ {
		a := ShuffleRandomly(GenerateDeck())
		b := BoolToByte(Sign(Diff(a)))
		randomDecks = append(randomDecks, b...)
	}
	err := ioutil.WriteFile("test.dat", randomDecks, 0644)
	if err != nil {
		log.Fatal(err)
	}

	randomDecks = []byte{}
	for i := 0; i < 100; i++ {
		b := BoolToByte(Sign(Diff(EmpiricallyShuffle(11))))
		randomDecks = append(randomDecks, b...)
	}
	err = ioutil.WriteFile("test2.dat", randomDecks, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func EmpiricallyShuffle(numberOfShuffles int) []int {
	deck := GenerateDeck()
	// Generate the packets
	for shuffleNum := 0; shuffleNum < numberOfShuffles; shuffleNum++ {
		packets := []int{}
		topDeck := 0
		for i := range deck {
			r := rand.Float64()
			cardsToAdd := 0
			for j, val := range EmpiricalShuffling {
				if r < val {
					cardsToAdd = j + 1
					break
				}
			}
			// Add the packets onto the deck
			if Sum(packets)+cardsToAdd > len(deck) {
				packets = append(packets, len(deck)-Sum(packets))
			} else {
				packets = append(packets, cardsToAdd)
			}

			// Figure out how many cards are in the cut
			if i%2 == 0 {
				topDeck = topDeck + packets[len(packets)-1]
			}

			// Finish once we reach 52
			if Sum(packets) == 52 {
				break
			}
		}

		fmt.Println(packets)
		fmt.Println(topDeck)
		// Now generate shuffled deck based on alternating packets
		var packet1 []int
		var packet2 []int
		if rand.Float64() < 0.5 {
			for i := 0; i < len(deck); i++ {
				if i < topDeck {
					packet1 = append(packet1, i)
				} else {
					packet2 = append(packet2, i)
				}
			}
		} else {
			for i := 0; i < len(deck); i++ {
				if i < len(deck)-topDeck {
					packet2 = append(packet2, i)
				} else {
					packet1 = append(packet1, i)
				}
			}
		}
		fmt.Println(packet1, packet2)
		packet1i := -1
		packet2i := -1
		cardi := -1
		cardNums := make([]int, len(deck))
		for i, val := range packets {
			if i%2 == 0 {
				for j := 0; j < val; j++ {
					cardi++
					packet1i++
					cardNums[cardi] = packet1[packet1i]
				}
			} else {
				for j := 0; j < val; j++ {
					cardi++
					packet2i++
					cardNums[cardi] = packet2[packet2i]
				}
			}
		}

		// Make a new deck
		newdeck := make([]int, len(deck))
		for i := range deck {
			newdeck[i] = deck[cardNums[i]]
		}
		fmt.Println(newdeck)
		deck = newdeck
	}
	return deck
}

func Sum(src []int) int {
	sum := 0
	for _, val := range src {
		sum += val
	}
	return sum
}

// BoolToByte takes an slice of bools and converts
// to a byte, TRUNCATING ANYTHING THAT DOESN'T FIT
func BoolToByte(src []bool) []byte {
	current := 0
	bytes := []byte{}
	for current+8 < len(src) {
		val := float64(0)
		for i := 0; i <= 7; i++ {
			if src[current+i] {
				val = val + math.Pow(2, float64(7-i))
			}
		}
		current += 8
		bytes = append(bytes, byte(val))
	}
	return bytes
}

// Sign takes a slice of integers and returns
// slice of bools that is true if >= 0 and false otherwise
func Sign(src []int) []bool {
	dest := make([]bool, len(src))
	for i, val := range src {
		if val >= 0 {
			dest[i] = true
		} else {
			dest[i] = false
		}
	}
	return dest
}

// Diff takes a slice of N integers and returns
// a slice of N-1 integers that is the difference
// between consecutive integers
func Diff(src []int) []int {
	diffVector := make([]int, len(src)-1)
	for i := 1; i < len(src); i++ {
		diffVector[i-1] = src[i] - src[i-1]
	}
	return diffVector
}

// ShuffleRandomly returns a shuffled list of ints
func ShuffleRandomly(src []int) []int {
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

// Generates a list of ints from 1 to 52
func GenerateDeck() []int {
	deck := make([]int, 52)
	for i := 0; i < 52; i++ {
		deck[i] = i
	}
	return deck
}
