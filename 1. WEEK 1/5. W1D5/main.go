package main

import "fmt"

//1/ Deklarasi Interface animal
type Animal interface{
	Sound() string
}

//2. Define struktur untuk cat dan Dog
type Cat struct{
	Name string
	Years int
}
type Dog struct{}

//3. Implementasi methode for both strutures
func (c Cat) Sound() string {
	//return "Meow!"
	return c.Name + " says: Meow!"
}

func (c Cat) Age() int {
	return c.Years
	
}

func (d Dog) Sound() string {
	return "Woof!"
}

//4. Define func untuk mencetak suara animal
func MakeSound(a Animal){
	fmt.Println(a.Sound())
}

// func main(){
// 	cat := Cat{}
// 	dog := Dog{}

// 	MakeSound(cat) // Meow!
// 	MakeSound(dog) // Woof!

// }

func main(){
	// cat := Cat{Name : "Whiskers"}
	// var a Animal = cat

	//typeassertion
	// value, ok := a.(Cat)
	// if ok {
	// 	fmt.Println("This is a cat name:", value.Name)
	// } else {
	// 	fmt.Println("This is not a cat")
	// }

	//type assertion2
	// value, ok := a.(Cat)
	// if ok {
	// 	fmt.Println("This is a cat name", value.Name, "and it id", value.Age(), "Yeard old")
	// } else {
	// 	fmt.Println("This is not a cat")
	// }

	cat1 := Cat{Name: "Whiskar"}
	cat2 := Cat{Name: "Felix"}
	dog := Dog{}

	//Interface Slice
	animals := []Animal{cat1, cat2, dog}
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}

	//Interface Map
	animalMap := map[string]Animal{
		"Whisker": cat1,
		"felixL": cat2,
		"Doggy": dog,
	}

	for name, animal := range animalMap {
		fmt.Println("%s: %s\n", name, animal.Sound())
	}


}