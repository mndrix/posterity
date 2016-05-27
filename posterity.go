package main

import "fmt"

type Family struct {
	Assets    int     // family assets in dollars
	Inflation float64 // inflation rate
	Return    float64 // return on assets

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

	// handle family assets
	income := float64(f.Assets) * (f.Return - f.Inflation)
	for _, x := range f.People {
		if x.Age == 13 || x.Age == 14 {
			income -= 1200 // $100 per month
		}
		if x.Age == 15 || x.Age == 16 {
			income -= 2400 // $200 per month
		}
		if x.Age == 17 {
			income -= 4800 // $400 per month
		}
		if x.Age >= 18 && x.Age <= 35 {
			income -= 12000 // $1,000 per month
		}
		if income < 0 { // or oldest ones get no money
			income = 0
		}
	}
	f.Assets += int(income)
}

func main() {
	const years = 20
	const iterations = 1000

	smallestSize := 999
	largestSize := 0
	sizes := make(map[int]int)

	smallestValue := 999
	largestValue := 0
	values := make(map[int]int)

	for i := 0; i < iterations; i++ {

		family := &Family{
			Assets:    1000000,
			Inflation: 0.01,
			Return:    0.05,

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

		value := family.Assets / 100000
		values[value]++
		if value > largestValue {
			largestValue = value
		}
		if value < smallestValue {
			smallestValue = value
		}
	}

	fmt.Println("\nAssets:")
	for i := smallestValue; i <= largestValue; i++ {
		fmt.Printf("%.1f %.1f\n", float64(i)/10, float64(values[i])/float64(iterations)*100)
	}

	fmt.Println("\nFamily size:")
	for i := smallestSize; i <= largestSize; i++ {
		fmt.Printf("%2d %.1f\n", i, float64(sizes[i])/float64(iterations)*100)
	}
}
