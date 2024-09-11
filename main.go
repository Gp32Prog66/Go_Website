package main 

import( 
	"fmt"
	"net/http"
)

func WebPageTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing Web Page")
}

func main() {
	http.HandleFunc("/", WebPageTest)
	http.ListenAndServe("", nil)
}