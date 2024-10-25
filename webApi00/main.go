package main

import (
	"fmt"
	"net/http"
)

// http suncuusnu olslutumak ve istekleri işlemek için gerekli araç

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Web Server başlatılıyor...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
	}
}
