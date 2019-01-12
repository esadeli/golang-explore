package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

// type mobilephone struct {
// 	Brand string // the name of parameter should start with capital letter
// 	// otherwise you will get empty array of object
// 	Price float64
// }

type mobilephone struct {
	Brand string  `json: "brand"`
	Price float64 `json: "price`
}

// implement sort object
type mobArr []mobilephone // new interface

// just change the name of input variable and the parameter to compare
// other than that, keep it as it is
func (data mobArr) Len() int           { return len(data) }
func (data mobArr) Swap(i, j int)      { data[i], data[j] = data[j], data[i] }
func (data mobArr) Less(i, j int) bool { return data[i].Price < data[j].Price }

func main10() {
	// arrPhone := []mobilephone{}

	phone1 := mobilephone{
		Brand: "Samsung",
		Price: 31000.33,
	}
	// arrPhone = append(arrPhone, phone1)

	phone2 := mobilephone{
		Brand: "Xiao Mi",
		Price: 42000.66,
	}

	phone3 := mobilephone{
		Brand: "Nokia",
		Price: 10000.11,
	}

	phone4 := mobilephone{
		Brand: "Oppo",
		Price: 51000.22,
	}

	// arrPhone = append(arrPhone, phone2, phone3, phone4)
	arrPhone := []mobilephone{phone1, phone2, phone3, phone4}
	fmt.Println("new array", arrPhone)

	// Marshal JSON
	arrMars, err := json.Marshal(arrPhone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arrMars))

	// unmarshal JSON
	s := `[{"Brand":"Samsung","Price":30000.03},{"Brand":"Xiao Mi","Price":40000.9}]`
	bs := []byte(s)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	var mobphone []mobilephone // define variable mobphone as an array of object /type mobilephone

	err2 := json.Unmarshal(bs, &mobphone) // should input byte type variable and a pointer
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println("All of mobile phone data ", mobphone)

	// interface
	fmt.Fprintln(os.Stdout, "Hello OS")
	io.WriteString(os.Stdout, "Hello OS IO\n")

	// sort demo
	sortIntEx := []int{3, 200, 2, 1000}
	sortStringEx := []string{"Arsenal", "AS Roma", "FC Barcelona", "FC Bayern"}

	sort.Ints(sortIntEx)
	fmt.Println(sortIntEx)
	sort.Strings(sortStringEx)
	fmt.Println(sortStringEx)

	// sorting and array of object
	sort.Sort(mobArr(arrPhone))
	fmt.Println("New sort ", arrPhone)

	// let's try to bcrypt something
	passwordToHash := "Niceone"
	bs, err3 := bcrypt.GenerateFromPassword([]byte(passwordToHash), bcrypt.MinCost)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(passwordToHash)
	fmt.Println(bs)

	// Bcrypt in GoLang !!!
	loginPass := "Niceone"
	err4 := bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err4 != nil {
		// fmt.Println(err4)
		fmt.Println("You can't Login")
		return
	} else {
		fmt.Println("Login successful")
	}

}
