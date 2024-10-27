// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// type API struct {
// 	Message string "json:message"
// }
// type User struct {
// 	ID        int    "json:id"
// 	FirstName string "json:firstname"
// 	LastName  string "json:lastname"
// 	Age       int    "json: age"
// }

// func main() {
// 	apiRoot := "/api"
// 	// .../api
// 	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
// 		message := API{"API Home"}
// 		output, err := json.Marshal(message)
// 		checkError(err)
// 		// w.Header().Set("Content -Type", "application/json")
// 		fmt.Fprintf(w, string(output))
// 	})

// 	// .../api/users
// 	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
// 		users := []User{
// 			User{ID: 1, FirstName: "Bartu", LastName: "kurnaz", Age: 21},
// 			User{ID: 2, FirstName: "Hakkı", LastName: "bayram", Age: 24},
// 			User{ID: 3, FirstName: "baran", LastName: "bulut", Age: 27},
// 		}
// 		message := users
// 		output, err := json.Marshal(message)
// 		checkError(err)
// 		fmt.Fprintf(w, string(output))

// 	})

// 	// .../api/me

// 	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
// 		user := User{3, "Bartu", "KURNAZ", 21}
// 		message := user
// 		output, err := json.Marshal(message)
// 		checkError(err)
// 		fmt.Fprintf(w, string(output))

// 	})
// 	//...

// 	http.ListenAndServe(":8080", nil)
// }

//	func checkError(err error) {
//		if err != nil {
//			fmt.Println("Fatal Error : ", err.Error())
//			os.Exit(1) // uygulamayı kapat
//		}
//	}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	apiRoot := "/api"

	// .../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)

		w.Header().Set("Content-Type", "application/json") // Başlık ayarlama
		w.Write(output)                                    // Yanıt gönderme
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "Bartu", LastName: "Kurnaz", Age: 21},
			{ID: 2, FirstName: "Hakkı", LastName: "Bayram", Age: 24},
			{ID: 3, FirstName: "Baran", LastName: "Bulut", Age: 27},
		}
		output, err := json.Marshal(users)
		checkError(err)

		w.Header().Set("Content-Type", "application/json") // Başlık ayarlama
		w.Write(output)                                    // Yanıt gönderme
	})

	// .../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 3, FirstName: "Bartu", LastName: "Kurnaz", Age: 21}
		output, err := json.Marshal(user)
		checkError(err)

		w.Header().Set("Content-Type", "application/json") // Başlık ayarlama
		w.Write(output)                                    // Yanıt gönderme
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1) // Uygulamayı kapat
	}
}
