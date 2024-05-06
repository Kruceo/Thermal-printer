package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thermal-printer/lib"
)

type QueryRequest struct {
	Query [][]string `json:"query"`
}

func main() {
	// Initialize a new Context.
	v := lib.GetDevice(0x04b8, 0x0e27)

	epOut := v.Out

	// inputs := [][]byte{

	// 	{0x1B, 0x52, 12}, //select latin america(12) character set
	// 	{0x1B, 0x40},     //init
	// 	// {0x10, 0x14},       //clear
	// 	[]byte(strings.Repeat("-", 47) + "\n"),
	// 	[]byte(lib.CenterString("VASCO", 47)),
	// 	{0x1b, 0x64, 5},    //feed 5 lines
	// 	{0x1D, 0x56, 0x00}, //cut
	// }

	// for _, v := range inputs {
	// 	bytesWritten, err := epOut.Write(v)
	// 	if err != nil {
	// 		log.Fatalf("Erro ao escrever dados: %v", err)
	// 	}
	// 	fmt.Printf("%d bytes enviados para o dispositivo\n", bytesWritten)
	// }

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		var decoded QueryRequest
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Fatal("Erro ao decodificar\n", err.Error())
		}

		var mode string = "left"

		for _, v := range decoded.Query {
			switch v[0] {
			case "println":
				if mode == "left" {
					epOut.Write([]byte(v[1] + "\n"))
				}
				if mode == "center" {
					epOut.Write([]byte(lib.CenterString(v[1], 49) + "\n"))
				}

			case "cut":
				epOut.Write([]byte{0x1b, 0x64, 5})
				epOut.Write([]byte{0x1D, 0x56, 0x00})

			case "center":
				mode = "center"

			case "left":
				mode = "left"
			}
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/get/width", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("49"))
	})

	fmt.Println("Listening 8888")
	http.ListenAndServe(":8888", nil)

}
