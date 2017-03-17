package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*
How many choices allowed?
Output format?
Input: DB? CSV? txt?
Categories:
Colors
Verbs
Adjectives
First names?
US President names
US States
Countries
Weather types
Natural objs (tree, rock, mountain)
Scientists
Villans
animals
fruit
jobs/roles
electronic parts (transistor, resistor, battery)
materials (steel, pine wood, cf)
time of day descriptors (dawn,evening, night)
seasons (spring, autumn)

*/
var (
	colors      = []string{"blue", "black", "green", "grey", "purple", "red", "sage", "yellow"}
	actionVerbs = []string{"acting", "breaking", "eating", "laughing", "running", "standing", "talking", "walking"}
	animals     = []string{"bear", "beaver", "bison", "badger", "bird", "buffalo", "cat", "cow", "dog", "duck", "fish", "goat", "horse", "pig", "rabbit", "rat", "sheep", "unicorn"}
	fruit       = []string{"apple", "banana", "orange", "strawberry", "lemon", "lime"}
	all         [][]string
)

type Options struct {
	color      bool
	animal     bool
	actionVerb bool
	fruit      bool
	random     bool
}

func newCodeName(opt Options) string {
	var codename string
	var (
		arr1 []string
		arr2 []string
	)
	rand.Seed(time.Now().UnixNano())
	// Random takes priority if selected
	if opt.random {
		// Keep getting arrays until they are not the same
		for {
			arr1 = all[rand.Intn(len(all))]
			arr2 = all[rand.Intn(len(all))]
			if arr1[0] != arr2[0] {
				break
			}
			fmt.Println("They were the same, trying again")
		}
		codename = arr1[rand.Intn(len(arr1))] + " " + arr2[rand.Intn(len(arr2))]
		return codename

	}
	r1 := rand.Intn(len(colors))
	r2 := rand.Intn(len(animals))
	r3 := rand.Intn(len(colors))
	fmt.Println(r1, r2, r3)
	codename = colors[r1] + " " + animals[r2]
	return codename
}
func init() {
	// Nest string arrays into array of all must access all[i][x]
	all = append(all, colors)
	all = append(all, actionVerbs)
	all = append(all, animals)
	all = append(all, fruit)
}
func Generate() {
	start := time.Now()
	opt := Options{color: true, animal: true, random: true}
	codename := newCodeName(opt)
	fmt.Printf("Codename: %v\n", codename)
	fmt.Println("Execution time: ", time.Since(start))

	err := errTest()
	if err.Error() == "Fail" {
		fmt.Println("expected")
	}
}
func errTest() error {
	return errors.New("Fail")
}
