package main
import "fmt"

func main(){
	isLoggedIn := true
	// hasSubscription := false
	isAdmin := true

	canOpenDashboard := isLoggedIn && isAdmin
	fmt.Println(canOpenDashboard)

	age :=18
	isAdult := age>25
	fmt.Println((isAdult))
	
}
  