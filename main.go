package main

import "fmt"
import "html/template"
import "net/http"

func isian(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "input.html")
		return
	}

	j := r.FormValue("isian")

	berhasil(w, r, j)

}

func berhasil(w http.ResponseWriter, r *http.Request, g string) {

	var data = map[string]string{"nama": g}

	var t, err = template.ParseFiles("tampil.html")

	if err != nil {
		fmt.Println("no name")
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", isian)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
