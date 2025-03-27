package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
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
	http.HandleFunc("/calculateTAS", calculateTAS)

	fmt.Println("Server is running on :5050...")
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func calculateTAS(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		ias := r.FormValue("ias")
		alt := r.FormValue("alt")
		isa := r.FormValue("isa")

		// Check if values for table are present
		if len(altitudeData) == 0 {
			fmt.Errorf("altitudeData is empty")
		}

		targetAltitude := convertStringToInt(alt)
		if err != nil {
			http.Error(w, "Error converting altitude to integer", http.StatusInternalServerError)
			return
		}

		_, closestIndex, err := findClosestAltitude(targetAltitude)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		key := fmt.Sprintf("ISA%v", isa)
		calculatedTAS := float64(altitudeData[closestIndex].ConversionFactors[key]) * convertStringToFloat(ias)
		calculatedTASStr := fmt.Sprintf("%.2f", calculatedTAS)

		response := map[string]string{"calculatedTAS": calculatedTASStr}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func convertStringToInt(altitudeStr string) int {
	altitude, _ := strconv.Atoi(altitudeStr)
	return altitude
}

func convertStringToFloat(altitudeStr string) float64 {
	altitude, _ := strconv.ParseFloat(altitudeStr, 64)
	return altitude
}

func findClosestAltitude(targetAltitude int) (*AltitudeData, int, error) {
	if len(altitudeData) == 0 {
		return nil, -1, fmt.Errorf("altitudeData is empty")
	}

	closestIndex := 0
	minDifference := math.Abs(float64(targetAltitude - convertStringToInt(altitudeData[0].Altitude)))

	for i := 1; i < len(altitudeData); i++ {
		currentAltitude := convertStringToInt(altitudeData[i].Altitude)
		currentDifference := math.Abs(float64(targetAltitude - currentAltitude))

		if currentDifference < minDifference {
			minDifference = currentDifference
			closestIndex = i
		}
	}

	return &altitudeData[closestIndex], closestIndex, nil
}
