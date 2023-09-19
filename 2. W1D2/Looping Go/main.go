// Case : Pencarian Data

// array slice : apple, banana, cherry, grappe, orange

// package main

// import "fmt"

// func main(){
// 	data := []string{"Apple", "Banana", "Cherry", "Garappe", "Orange"}
// 	target := "Cherry"


// 	found := false
// 	for _, item := range data{
// 		if item == target {
// 		found = true
// 		break
// 		} 
// 	} 

// 	if found {
// 		fmt.Printf("Item %s found in the slice.\n", target)
//  	} else {
// 		fmt.Printf("Item %s not found in the slice.\n", target)
// 	}

// }

// Case 2 : Menghitung jumlah kemunculan Angka

// Misal ada data 1,2,3,4,3,2,1,2,3,4,3,2,1

// package main

// import "fmt"



// func main(){
// 	data := []int{1,2,3,4,3,2,1,2,3,4,3,2,1}
// 	target := 2

// 	count := 0
// 	for _, item := range data {
// 		if item == target {
// 			count++
// 		}
// 	}
// 	fmt.Printf("Item %d appears %d times in the slixe.\n", target, count)
// }

// case 3 : Mengitung rata rata angka

// Misal ada data 1,2,3,4,3,2,1,2,3,4,3,2,1
// coba buat program untuk pengolahan data : menghitung statistik dari kumpuklan slice AVG/Rata2

package main



func main(){
	loop1()
	loop2()
	loop3()
	loop4()
	loop5()
	loop6()
	loop7()
	loop8()
	loop9()
	// data := []int{1,2,3,4,3,2,1,2,3,4,3,2,1}
	// sum := 0 // Untuk inisialisai variable jumlah dengan nilai 0

	// //Itearasi slice
	// for _, value := range data {
	// 	sum += value 
	// }

	// avg := float64(sum) / float64(len(data))
	// fmt.Printf("Rata-rata: %.2f\n", avg)
}