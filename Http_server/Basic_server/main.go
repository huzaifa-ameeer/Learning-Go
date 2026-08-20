package main
import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		fmt.Println("Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
	_,_ = w.Write([]byte("Hello from GO"))
}

func main (){
	http.HandleFunc("/", handle)
	fmt.Println("Server running")
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
