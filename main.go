package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thermal-printer/lib"
	"thermal-printer/lib/configs/epson"
	"time"
)

type QueryRequest struct {
	Query [][]string `json:"query"`
}

func main() {
	v := lib.GetDevice(0x04b8, 0x0e27)

	MainPrinter := epson.CreateEpsonPrinter(v.Out)

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {

		// Handle preflight OPTIONS requests, requisited by CORS
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			return
		}

		var decoded QueryRequest
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Fatal("Decode error\n", err.Error())
		}

		var mode string = "left"

		fmt.Println(r.Host, len(decoded.Query), "queries")

		//init, clear buffers
		MainPrinter.Clear()
		// set character set to "TURKISH"
		MainPrinter.Write(epson.WPC1254_Turkish)

		for _, v := range decoded.Query {
			switch v[0] {
			case "println":

				if mode == "left" {
					inputStr := v[1] + "\n"
					inputASCII := lib.String2ExtASCII(inputStr)
					MainPrinter.Write(inputASCII)
				}
				if mode == "center" {
					inputStr := lib.CenterString(v[1], 49) + "\n"
					inputASCII := lib.String2ExtASCII(inputStr)
					MainPrinter.Write(inputASCII)
				}

			case "cut":
				MainPrinter.FeedLines(5)
				time.Sleep(10000)
				MainPrinter.FullCut()

			case "center":
				mode = "center"

			case "left":
				mode = "left"
			}
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	http.HandleFunc("/get/width", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"width\":48}"))
	})

	fmt.Println("Listening 8888")
	http.ListenAndServe(":8888", nil)

}

func convertToExtendedASCII(text string) string {
	extendedASCII := ""
	for _, char := range text {
		if char >= 128 && char <= 255 {
			extendedASCII += string(char)
		}
	}
	return extendedASCII
}
