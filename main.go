package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello my fellas, hello Argo my friend</h1>"))
	})
	http.ListenAndServe(":8081", nil)
}
