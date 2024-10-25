// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// type Human struct {
// 	FName string
// 	LName string
// 	Age   int
// }

// func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	hum.FName = "Bartu"
// 	hum.LName = "Kurnaz"
// 	hum.Age = 21

// 	// formu parse etmek için
// 	r.ParseForm()

// 	// SUnucudan form bigisini almak için
// 	fmt.Println(r.Form)

// 	//URL'nin path bilgisini almak için
// 	fmt.Println("path", r.URL.Path)

// 	// fmt.Fprint(w, "<table><tr><td><b><Ad></b><Soyad></b></td><td><b><Yas></b></td></tr><tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr><tr></tr><tr></tr><tr><td><b>Path</b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")

// 	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yas</b></td></tr>")
// 	fmt.Fprint(w, "<tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr>")
// 	fmt.Fprint(w, "<tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")

// }

// func main() {
// 	var hum Human
// 	err := http.ListenAndServe("localhost :9000", hum)
// 	checkError(err)
// }
// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FName = "Bartu"
	hum.LName = "Kurnaz"
	hum.Age = 21

	// Formu parse etmek için
	r.ParseForm()

	// Sunucudan form bilgisini almak için
	fmt.Println("Form verisi:", r.Form)

	// URL'nin path bilgisini almak için
	fmt.Println("Path:", r.URL.Path)

	// HTML tablosunu yazdırma
	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yas</b></td></tr>")
	fmt.Fprint(w, "<tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr>")
	fmt.Fprint(w, "<tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")
}

func main() {
	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	checkError(err)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
