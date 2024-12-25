package exercise_practice

import "fmt"

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
func PracticeExercise() {
	//1 Solution
	hobbies := [3]string{"Programming", "Cricket", "Adventure"}
	fmt.Println("My top 3 hobbies are", hobbies)

	//2 Solution
	fmt.Println("My 1st hobby is: ", hobbies[0])
	fmt.Println("My second and third hobbies are: ", hobbies[1:3])

	//3 Solution
	//3.1
	hobbiesSubSlice := hobbies[0:2]
	fmt.Println("Hobbies sub slice is: ", hobbiesSubSlice)

	//3.2

	//4 Solution
	hobbiesReslice := hobbiesSubSlice[1:3]
	fmt.Println("Hobbies reslice is: ", hobbiesReslice)

	//5 Solution
	courseGoals := []string{"Finish the GO course", "Focus on Learning not on finishing in GO course"}
	fmt.Println("Two course goals are:", courseGoals)

	//6 Solution
	courseThreeGoals := append(courseGoals, "Consistent on learning the course")
	fmt.Println("Course three goals are: ", courseThreeGoals)

	//7 Solution
	type Product struct {
		id    int64
		title string
		price float64
	}
	dynamicListOfProducts := []Product{{1, "Mobile", 26.00}, {2, "Tablet", 99.99}}
	fmt.Println("Dynamic product list is: ", dynamicListOfProducts)

	dynamicListOfThreeProducts := append(dynamicListOfProducts, Product{3, "Monitor", 89.99})
	fmt.Println("Dynamic three product list is: ", dynamicListOfThreeProducts)
}
