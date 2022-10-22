package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/tidwall/geodesic"
)

type capital struct {
	country, city string
	lat, long     float64
}

type capital_distance struct {
	cap_a, cap_b capital
	distance     float64
}

func main() {
	f, err := os.Open("../country-capitals.csv")
	if err != nil {
		log.Fatal("Unable to read input file ")
		fmt.Println("this is a test")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	var capitals []capital
	for i, record := range records {
		if i == 0 {
			continue
		}
		lat, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatalln(err)
		}
		long, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatalln(err)
		}

		capitals = append(capitals, capital{record[0], record[1], lat, long})
	}

	var distances []capital_distance
	for i, cap_a := range capitals {
		for _, cap_b := range capitals[i+1:] {
			var dist float64
			geodesic.WGS84.Inverse(cap_a.lat, cap_a.long, cap_b.lat, cap_b.long, &dist, nil, nil)
			distances = append(distances, capital_distance{cap_a, cap_b, dist})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	for _, d := range distances[:10] {
		fmt.Printf("%20s, %20s -> %20s, %20s: %20f\n", d.cap_a.city, d.cap_a.country, d.cap_b.city, d.cap_b.country, d.distance)
	}
}
