package main

import (
	"fmt"

	"rsc.io/quote" // run "go mod tidy" -> this command will add the module as a requirement
)

func main()  {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(greeting("Andrew"))
	
}

func greeting(name string) string {
	message:= fmt.Sprintf("Hi %v. Welcome!", name)
	return message
}

// import (
//     "fmt"
//     "net/http"
// )

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//     })

//     http.ListenAndServe(":80", nil)
// }