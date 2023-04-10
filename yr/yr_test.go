package yr

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

// antall linjer i filen er 16756
func TestFileLineCount(t *testing.T) {
	// lineCount, err := GetFileLineCount("C:\Users/simon/reps/is-105/go/minyr2/kjevik-temp-celsius-20220318-20230318.csv")
	filename := ("kjevik-temp-celsius-20220318-20230318.csv")
	expectedLines := 16756

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to scan file %s: %v", filename, err)
	}

	if lineCount != expectedLines {
		t.Errorf("unexpected line count in file %s: expected %d, got %d", filename, expectedLines, lineCount)
	}
}

// gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42,8"
func TestConversion8(t *testing.T) {
	// Open the converted CSV file
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader to read the converted CSV file
	reader := csv.NewReader(file)

	// Loop through each line of the converted CSV file
	for {
		// Read a single line from the converted CSV file
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// Reached end of file, break loop
				break
			} else {
				t.Errorf("Error reading file: %v", err)
				return
			}
		}

		// Check if the line matches the specified line
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "18.03.2022 01:50" {
			// Check if the temperature value has been correctly converted
			want := "42.8"
			got := line[3]
			if got != want {
				t.Errorf("Conversion error. Got %v, want %v", got, want)
			}
			return
		}
	}
	t.Errorf("Line not found in file.")
}

// gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32"
func TestConversion32(t *testing.T) {
	// Open the converted CSV file
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader to read the converted CSV file
	reader := csv.NewReader(file)

	// Loop through each line of the converted CSV file
	for {
		// Read a single line from the converted CSV file
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// Reached end of file, break loop
				break
			} else {
				t.Errorf("Error reading file: %v", err)
				return
			}
		}

		// Check if the line matches the specified line
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "07.03.2023 18:20" {
			// Check if the temperature value has been correctly converted
			want := "32.0"
			got := line[3]
			if got != want {
				t.Errorf("Conversion error. Got %v, want %v", got, want)
			}
			return
		}
	}

	t.Errorf("Line not found in file.")
}

// gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12.2"
func TestConversion2(t *testing.T) {
	// Open the converted CSV file
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader to read the converted CSV file
	reader := csv.NewReader(file)

	// Loop through each line of the converted CSV file
	for {
		// Read a single line from the converted CSV file
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// Reached end of file, break loop
				break
			} else {
				t.Errorf("Error reading file: %v", err)
				return
			}
		}

		// Check if the line matches the specified line
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "08.03.2023 02:20" {
			// Check if the temperature value has been correctly converted
			want := "12.2"
			got := line[3]
			if got != want {
				t.Errorf("Conversion error. Got %v, want %v", got, want)
			}
			return
		}
	}

	t.Errorf("Line not found in file.")
}

/*
gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want)
"Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av
STUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen
*/
func TestLastLineOfFile(t *testing.T) {
	// Open the file for reading
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Scan through the file line by line, keeping track of the last line encountered
	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	// Check that the last line contains the expected text
	expectedText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Simon Helgen"
	if !strings.Contains(lastLine, expectedText) {
		t.Errorf("last line of file does not contain expected text. got: %q, want substring: %q", lastLine, expectedText)
	}
}

// en test som sjekker at gjennomsnittempraturen er 8.56
func TestAverageTemperature(t *testing.T) {
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var sum float64
	var count int
	for i := 1; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if i == 1 || i == 16756 {
			continue // skip the first and last line
		}
		temp, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			t.Fatalf("failed to parse temperature: %v", err)
		}
		sum += temp
		count++
	}
	avg := sum / float64(count)

	expected := 8.56
	if avg != expected {
		t.Errorf("average temperature is %v, expected %v", avg, expected)
	}
}
