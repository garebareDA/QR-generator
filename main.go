package main

import(
	"github.com/boombuler/barcode"
  "github.com/boombuler/barcode/qr"
	"image/png"
	"net/http"
	"html/template"
	"fmt"
)

func main(){
	http.HandleFunc("/qr", qrHandler)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
	fmt.Println("listen port 8000")
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
	message, err := r.URL.Query()["url"]
	if(!err || len(message[0])< 1){
		fmt.Println(w, "Missing")
		return
	}

	qrCode, _ := qr.Encode(message[0], qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, nil)
}