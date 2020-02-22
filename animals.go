package main

import (
	"errors"
	"fmt"
	"strings"
	// "strings"
)

//Animal ...
type Animal struct {
	ID    int
	Name  string
	Voice string
	Leg   int
}

//AnimalList ...
type AnimalList struct {
	animals []Animal
}
//Animals ...
type Animals interface {
	Add(an Animal)
	PrintList(an Animal)
	Update(id int, an Animal) error
	Search(name string, leg int) error
	Delete(i int) error
}

// Errors
var (
	ErrorNotFound = errors.New("Animal Not Found")
)

//NewAnimalList ...
func NewAnimalList() AnimalList {
	animalList := AnimalList{}
	animalList.animals = []Animal{}

	return animalList
}
//Add ...
func (a *AnimalList) Add(an *Animal) []Animal {
	a.animals = append(a.animals, *an)

	return a.animals
}
//Delete ...
func (a *AnimalList) Delete(id int) {

	a.animals = append(a.animals[:id], a.animals[id+1:]...)

	return

}
//Update ...
func (a *AnimalList) Update(id int, an *Animal) error {
	var (
		isfound bool
		orderID int
	)
	for idd, contact := range a.animals {

		if contact.ID == id {

			orderID = idd
			isfound = true
		}
	}

	a.animals[orderID].Name = an.Name
	a.animals[orderID].Leg = an.Leg
	a.animals[orderID].Voice = an.Voice

	if isfound {
		return nil
	}

	return ErrorNotFound
}
//Search ...
func (a *AnimalList) Search(name string, leg int) []Animal {
	var animals []Animal

	for _, cl := range a.animals {
		if strings.Contains(strings.ToLower(cl.Name), strings.ToLower(name)) && cl.Leg == leg {
			animals = append(animals, cl)
		}

	}

	return animals
}
//PriList ...
func (a *AnimalList) PriList(anlist *AnimalList) {
	fmt.Println(*anlist)
}



func main() {
	al := NewAnimalList()

	al.Add(&Animal{ID: 0, Name: "Bobo", Leg: 4, Voice: "gav"})
	al.Add(&Animal{ID: 1, Name: "ohho", Leg: 2, Voice: "dog"})
	al.Add(&Animal{ID: 2, Name: "horse", Leg: 4, Voice: "ihaaa"})


	

	al.Delete(0) //  delete function number is id

	err := al.Update(1, &Animal{ID: 10, Name: "wolf", Leg: 4, Voice: "auu"})

	if err == ErrorNotFound {
		fmt.Println("Animal not found. Id")
	}
	fmt.Println(al.Search(" ",4))

	// al.PriList(&al)
		

}
