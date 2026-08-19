package main
import "fmt"

func main (){
	// age := map[string]int {
	// 	"huzaifa" : 22,
	// 	"ghufran" : 24,
	// }

	// fmt.Println(age["ghufran"])

	// scores := make(map[string]int , 10)
	// scores["a"] = 21
	// scores["2"] = 21
	// scores["3"] = 21
	// scores["2"] = 21

	// fmt.Println(len(scores))

	users := map[string] string {
		"u1" : "Huzaifa",
		"u2" : "Ali",
		"u3" : "Zubair",
		"u4" : "Ahmad",
	}

	for i, name := range(users){
		fmt.Println(i,name)
	}

	delete(users, "u3")

	fmt.Println(users)

}