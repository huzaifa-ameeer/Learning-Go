package main
import "fmt"

func main(){
	views1 := 500
	views2 := 1000
	totalViews := views1 + views2
	avgViews := totalViews/2

	fmt.Println(totalViews, avgViews)

	rating1 := 4.5
	rating2 := 5.2
	avgRatings := (rating1 + rating2)/2
	
	fmt.Println(avgRatings)
}