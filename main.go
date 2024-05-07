package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"thermal-printer/lib"
	"thermal-printer/lib/configs/epson"
	"time"

	"github.com/joho/godotenv"
)

type QueryRequest struct {
	Query [][]string `json:"query"`
}

func main() {

	// LOAD .env
	if err := godotenv.Load(); err != nil {
		// log.Fatalf("Error loading .env: %v", err)
	}

	printerName := lib.GetEnvOrDefault("PRINTER_NAME", "TM-T20X")
	printerVendor := lib.GetEnvOrDefault("PRINTER_VENDOR", "EPSON")
	printerWidth, err := strconv.Atoi(lib.GetEnvOrDefault("PRINTER_WIDTH", "48"))

	if err != nil {
		log.Fatalf("Bad format in PRINTER_WIDTH (%s):\n%v", lib.GetEnvOrDefault("PRINTER_WIDTH", "48"), err)
	}

	printerCharsetName := lib.GetEnvOrDefault("PRINTER_CHARSET", "default")
	printerCharset := epson.CharacterSet[printerCharsetName]
	if printerCharset == nil {
		fmt.Printf("Printer CharSet '%s' not exists. Using 'default'.\n", printerCharsetName)
		printerCharset = epson.CharacterSet["default"]
	}

	v := lib.GetDeviceByName(printerVendor, printerName)

	fmt.Printf("Using %s %s (%s:%s)\n", v.VendorName, v.ProductName, v.VID, v.PID)

	MainPrinter := epson.CreateEpsonPrinter(v.Out, printerCharset)

	var printing = false

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {

		// Handle preflight OPTIONS requests, requisited by CORS
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			return
		}

		if printing {
			for i := 0; i < 60; i++ {
				fmt.Printf("\rWaiting %d seconds", i+1)
				time.Sleep(1 * time.Second)
				if !printing {
					fmt.Printf("\r" + strings.Repeat(" ", 20))
					fmt.Printf("\r")
					break
				}
			}

		}

		if printing {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"error\":true,\"The printer are busy.\"}"))
			return
		}

		// Locks new requests
		printing = true

		var decoded QueryRequest
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if err != nil {
			log.Fatal("Decode error:\n", err.Error())
		}

		if len(decoded.Query) > 100 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"error\":true,\"The num of queries exceed the max of queries per request.\"}"))
			return
		}

		fmt.Println(r.Method, r.Host, r.URL.Path)

		var mode string = "left"

		//init, clear buffers
		MainPrinter.Clear()

		for _, v := range decoded.Query {

			switch v[0] {
			case "println":

				if mode == "left" {
					inputStr := v[1] + "\n"
					inputASCII := lib.String2ExtASCII(inputStr)
					MainPrinter.Write(inputASCII)
				}
				if mode == "center" {
					inputStr := lib.CenterString(v[1], printerWidth) + "\n"
					inputASCII := lib.String2ExtASCII(inputStr)
					MainPrinter.Write(inputASCII)
				}

			case "cut":
				MainPrinter.FeedLines(5)
				MainPrinter.FullCut()

			case "center":
				mode = "center"

			case "left":
				mode = "left"
			}
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"message\":\"Query created with success.\"}"))

		// unlocking for new requests
		printing = false
	})

	http.HandleFunc("/get/width", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.Host, r.URL.Path)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != "GET" {
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"width\":" + strconv.Itoa(printerWidth) + "}"))

	})

	fmt.Println("Listening 8888")
	http.ListenAndServe(":8888", nil)

}
