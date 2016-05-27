package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Person struct {
	Gender           Gender
	Age              int
	NumberOfChildren int
}

type Gender int

const (
	Male   Gender = iota
	Female Gender = iota
)

func (self Gender) String() string {
	if self == Male {
		return "M"
	}
	if self == Female {
		return "F"
	}
	panic("unexpected gender")
}

func (self *Person) String() string {
	return fmt.Sprintf("%s %d %d", self.Gender, self.Age, self.NumberOfChildren)
}

// HasNewChild returns true if this person has a child during the transition to
// the next time interval.  The invocant's child count is incremented and the
// new child itself is returned.
func (self *Person) HasNewChild() (*Person, bool) {
	var P float64

	// probability of having a baby
	switch self.Gender {
	case Male:
		if self.Age < 15 {
			P = 0
		} else if self.Age <= 19 {
			P = 11.3 / 1000
		} else if self.Age <= 24 {
			P = 53.9 / 1000
		} else if self.Age <= 29 {
			P = 89.7 / 1000
		} else if self.Age <= 34 {
			P = 103.9 / 1000
		} else if self.Age <= 39 {
			P = 68.8 / 1000
		} else if self.Age <= 44 {
			P = 27.9 / 1000
		} else if self.Age <= 49 {
			P = 9.3 / 1000
		} else if self.Age <= 54 {
			P = 2.8 / 1000
		} else {
			P = 0.4 / 1000
		}
	case Female:
		if self.Age < 10 {
			P = 0
		} else if self.Age <= 14 {
			P = 0.2 / 1000
		} else if self.Age <= 17 {
			switch self.NumberOfChildren {
			case 0:
				P = 9.4 / 1000
			case 1:
				P = 0.7 / 1000
			case 2:
				P = 0
			case 3:
				P = 0
			case 4:
				P = 0
			case 5, 6:
				P = 0
			default:
				P = 0
			}
		} else if self.Age <= 19 {
			switch self.NumberOfChildren {
			case 0:
				P = 33.7 / 1000
			case 1:
				P = 7.2 / 1000
			case 2:
				P = 0.9 / 1000
			case 3:
				P = 0.1 / 1000
			case 4:
				P = 0 / 1000
			case 5, 6:
				P = 0 / 1000
			default:
				P = 0 / 1000
			}
		} else if self.Age <= 24 {
			switch self.NumberOfChildren {
			case 0:
				P = 39.7 / 1000
			case 1:
				P = 25.7 / 1000
			case 2:
				P = 8.9 / 1000
			case 3:
				P = 2.2 / 1000
			case 4:
				P = 0.5 / 1000
			case 5, 6:
				P = 0.1 / 1000
			default:
				P = 0 / 1000
			}
		} else if self.Age <= 29 {
			switch self.NumberOfChildren {
			case 0:
				P = 41.4 / 1000
			case 1:
				P = 36.7 / 1000
			case 2:
				P = 19.7 / 1000
			case 3:
				P = 7.4 / 1000
			case 4:
				P = 2.4 / 1000
			case 5, 6:
				P = 1.0 / 1000
			default:
				P = 0.1 / 1000
			}
		} else if self.Age <= 34 {
			switch self.NumberOfChildren {
			case 0:
				P = 31.4 / 1000
			case 1:
				P = 35.6 / 1000
			case 2:
				P = 21.1 / 1000
			case 3:
				P = 9.5 / 1000
			case 4:
				P = 3.6 / 1000
			case 5, 6:
				P = 2.1 / 1000
			default:
				P = 0.5 / 1000
			}
		} else if self.Age <= 39 {
			switch self.NumberOfChildren {
			case 0:
				P = 11.6 / 1000
			case 1:
				P = 16.3 / 1000
			case 2:
				P = 11.6 / 1000
			case 3:
				P = 6.1 / 1000
			case 4:
				P = 2.8 / 1000
			case 5, 6:
				P = 2.0 / 1000
			default:
				P = 0.8 / 1000
			}
		} else if self.Age <= 44 {
			switch self.NumberOfChildren {
			case 0:
				P = 2.2 / 1000
			case 1:
				P = 2.8 / 1000
			case 2:
				P = 2.1 / 1000
			case 3:
				P = 1.3 / 1000
			case 4:
				P = 0.7 / 1000
			case 5, 6:
				P = 0.6 / 1000
			default:
				P = 0.4 / 1000
			}
		} else if self.Age <= 49 {
			switch self.NumberOfChildren {
			case 0:
				P = 0.2 / 1000
			case 1:
				P = 0.2 / 1000
			case 2:
				P = 0.1 / 1000
			case 3:
				P = 0.1 / 1000
			case 4:
				P = 0 / 1000
			case 5, 6:
				P = 0.1 / 1000
			default:
				P = 0 / 1000
			}
		} else {
			P = 0
		}
	}

	// having a baby?
	if rand.Float64() < P {
		self.NumberOfChildren++
		baby := &Person{}
		if rand.Float64() < 0.48 {
			baby.Gender = Male
		} else {
			baby.Gender = Female
		}
		return baby, true
	}

	return nil, false
}

// Dies returns true if this person should die during the transition to the next
// time interval.
func (self *Person) Dies() bool {
	var P float64

	switch self.Gender {
	case Male:
		if self.Age < 1 {
			P = 0.006575
		} else if self.Age < 4 {
			P = 0.0003
		} else if self.Age < 13 {
			P = 0.0001
		} else if self.Age < 20 {
			P = 0.0005
		} else if self.Age < 40 {
			P = 0.0015
		} else if self.Age < 50 {
			P = 0.0035
		} else if self.Age < 60 {
			P = 0.0075
		} else if self.Age < 70 {
			P = 0.015
		} else if self.Age < 80 {
			P = 0.035
		} else if self.Age < 90 {
			P = 0.10
		} else if self.Age < 100 {
			P = 0.25
		} else {
			P = 1.0
		}
	case Female:
		if self.Age < 1 {
			P = 0.005516
		} else if self.Age < 4 {
			P = 0.0002
		} else if self.Age < 13 {
			P = 0.0001
		} else if self.Age < 20 {
			P = 0.0002
		} else if self.Age < 40 {
			P = 0.0006
		} else if self.Age < 50 {
			P = 0.0017
		} else if self.Age < 60 {
			P = 0.0035
		} else if self.Age < 70 {
			P = 0.01
		} else if self.Age < 80 {
			P = 0.025
		} else if self.Age < 90 {
			P = 0.075
		} else if self.Age < 100 {
			P = 0.20
		} else {
			P = 1.0
		}
	}

	return rand.Float64() < P
}
