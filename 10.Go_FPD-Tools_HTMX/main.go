package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type AltitudeData struct {
	Altitude          string
	ConversionFactors map[string]float32
}

var altitudeData []AltitudeData

func loadAltitudeData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &altitudeData); err != nil {
		return err
	}

	return nil
}

func sortMapKeys(m map[string]float32) []string {
	// Manually specify the desired order of keys
	order := []string{"ISA-30", "ISA-20", "ISA-10", "ISA", "ISA+10", "ISA+15", "ISA+20", "ISA+30"}

	// Create a map to check for key existence
	exists := make(map[string]bool)
	for _, key := range order {
		exists[key] = true
	}

	// Append other keys not present in the order
	for key := range m {
		if !exists[key] {
			order = append(order, key)
		}
	}

	return order
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").Funcs(template.FuncMap{"sortMapKeys": sortMapKeys}).ParseFiles("index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, altitudeData)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	if err := loadAltitudeData("altitude_data.json"); err != nil {
		fmt.Println("Error loading altitude data from json:", err)
		return
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on :5050...")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
