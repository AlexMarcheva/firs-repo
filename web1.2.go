package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
//func main() {
//
//	fs := http.FileServer(http.Dir("static"))
//	http.Handle("/", fs)
//
//	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
//		fmt.Fprint(w, "About Page")
//	})
//	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
//		fmt.Fprint(w, "Contact Page")
//	})
//	fmt.Println("Server is listening...")
//	http.ListenAndServe(":8181", nil)
//}

func productsHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w,response)
}