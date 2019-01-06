// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values,
// ranging over the elements in the slice which stores the favorite flavors.

// Hands-on exercise #2
// Take the code from the previous exercise,
// then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.

// Hands-on exercise #3
// Create a new type: vehicle.
// The underlying type is a struct.
// The fields:
// doors
// color
// Create two new types: truck & sedan.
// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool. solution
// Using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// Print out each of these values.
// Print out a single field from each of these values.

// Hands-on exercise #4
// Create and use an anonymous struct.

package main

import (
	"fmt"
)

type foodies struct {
	firstName      string
	lastName       string
	favoriteFlavor string
}

type vehicle struct {
	doors string
	color string
}

type truck struct {
	fourwheel bool
	vehicle
}

type sedan struct {
	luxury bool
	vehicle
}

func main() {
	// exersice 1

	favFlavor := []string{`Chocolate`, `Vanilla`, `Supreme Sundae`}

	person1 := foodies{
		firstName:      "Andrew",
		lastName:       "Jash",
		favoriteFlavor: favFlavor[1],
	}

	person2 := foodies{
		firstName:      "Marry",
		lastName:       "Jash",
		favoriteFlavor: favFlavor[0],
	}
	fmt.Println(person1)
	fmt.Println(person2)

	// exersice 2
	identity3 := map[string]foodies{
		person1.firstName: person1,
		person2.firstName: person2,
	}
	fmt.Println(identity3)

	// exersice 3

	sedanMercy := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: "Two",
			color: "Black",
		},
	}

	volvoTruck := truck{
		fourwheel: false,
		vehicle: vehicle{
			doors: "Two",
			color: "White",
		},
	}
	fmt.Println(sedanMercy)
	fmt.Println(volvoTruck)

	// exersice 4
	ex4 := struct {
		name   string
		club   string
		number int
	}{
		name:   "Andrea Pirlo",
		club:   "AC Milan",
		number: 21,
	}
	fmt.Println(ex4)
}
