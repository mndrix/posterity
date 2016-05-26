package main

import "fmt"

type Family struct {
	People []*Person
}

// Tick iterates a single person forward in time by one unit. The resulting
// slice of people represents the people still alive at the next time interval.
func Tick(x *Person) []*Person {
	var people []*Person

	if y, ok := x.HasNewChild(); ok {
		people = append(people, y)
	}
	if !x.Dies() {
		x.Age++
		people = append(people, x)
	}

	return people
}

// Next adjusts the family one tick of the clock.
func (f *Family) Next() {
	// handle births and deaths
	var people []*Person
	for _, x := range f.People {
		xs := Tick(x)
		people = append(people, xs...)
	}
	f.People = people
}

func main() {
	const years = 20
	const iterations = 1000

	smallestSize := 999
	largestSize := 0
	sizes := make(map[int]int)
	for i := 0; i < iterations; i++ {

		family := &Family{
			People: []*Person{
				{Male, 37, 5},   // me
				{Female, 12, 0}, // Ara
				{Male, 11, 0},   // Jericho
				{Female, 9, 0},  // Haven
				{Male, 6, 0},    // Gideon
				{Male, 2, 0},    // Brigham
			},
		}

		for y := 0; y < years; y++ {
			family.Next()
		}

		size := len(family.People)
		sizes[size]++
		if size > largestSize {
			largestSize = size
		}
		if size < smallestSize {
			smallestSize = size
		}
	}

	for i := smallestSize; i <= largestSize; i++ {
		fmt.Printf("%2d %.1f\n", i, float64(sizes[i])/float64(iterations)*100)
	}
}
