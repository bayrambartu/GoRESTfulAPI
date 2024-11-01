// package main

// import (
// 	"bytes"
// 	"io/ioutil"
// 	"net/http"
// 	"text/template"
// )

// type Page struct {
// 	Title           string
// 	Author          string
// 	Header          string
// 	PageDescription string
// 	Content         string
// 	URI             string
// }

// func loadFile(fileName string) (string, error) {
// 	bytes, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bytes), nil
// }

// func handler(w http.ResponseWriter, r *http.Request) {

// 	// String birleştirme işlemi
// 	var builder bytes.Buffer
// 	builder.WriteString("KodLab yayınevinden çıkardığımız Yazılımcılar İçin İleri Seviye T-SQL kitabımın özellikleri aşağıdaki gibidir;\n")
// 	builder.WriteString("704 Sayfa\n")
// 	builder.WriteString("ISBN: 9.786.055.201.142\n")
// 	builder.WriteString("Fiyat: 37 TL\n")
// 	builder.WriteString("Boyut: 15 x 21\n")
// 	builder.WriteString("2. Baskı\n")

// 	uri := "www.bartukurnaz.com/yazilimcilar-icin-ileri-seviye-t-sql-programlama-kitabi/"

// 	page := Page{
// 		Title:           "Kitap : İleri Seviye T-SQL Programlama",
// 		Author:          "Bayram Bartu Kurnaz ",
// 		Header:          "İleri Seviye T-SQL Programlama",
// 		PageDescription: "İleri Seviye T-SQL Programlama kitap tanıtım sayfası",
// 		Content:         builder.String(),
// 		URI:             "http://" + uri}
// 	t, _ := template.ParseFiles("page.html")
// 	t.Execute(w, page)
// }
// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":9000", nil)

// }
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// String birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("KodLab yayınevinden çıkardığımız Yazılımcılar İçin İleri Seviye T-SQL kitabımın özellikleri aşağıdaki gibidir;\n")
	builder.WriteString("704 Sayfa\n")
	builder.WriteString("ISBN: 9.786.055.201.142\n")
	builder.WriteString("Fiyat: 37 TL\n")
	builder.WriteString("Boyut: 15 x 21\n")
	builder.WriteString("2. Baskı\n")

	uri := "www.bartukurnaz.com/yazilimcilar-icin-ileri-seviye-t-sql-programlama-kitabi/"

	page := Page{
		Title:           "Kitap : İleri Seviye T-SQL Programlama",
		Author:          "Bayram Bartu Kurnaz ",
		Header:          "İleri Seviye T-SQL Programlama",
		PageDescription: "İleri Seviye T-SQL Programlama kitap tanıtım sayfası",
		Content:         builder.String(),
		URI:             "http://" + uri,
	}

	// Şablonu yükleme ve hata kontrolü
	t, err := template.ParseFiles("page.html")
	if err != nil {
		http.Error(w, "Şablon dosyası yüklenemedi", http.StatusInternalServerError)
		fmt.Println("Şablon hatası:", err)
		return
	}

	// Şablonu çalıştırma ve hata kontrolü
	err = t.Execute(w, page)
	if err != nil {
		http.Error(w, "Şablon çalıştırılamadı", http.StatusInternalServerError)
		fmt.Println("Şablon çalıştırma hatası:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Sunucu başlatılıyor: http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
