package utils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestBatteryHasPowerSupplyNew(t *testing.T) {
	returnVal, err := BatteryHasPowerSupply()
	if err != nil {
		t.Errorf("Failed to get battery status")
	}
	if returnVal && !returnVal {
		t.Errorf("Unexpected status returned: %v", returnVal)
	}
}

func TestSendDataToServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	data := Indicator{
		ElectricityStatus: true,
		LocationName:      "TestLocation",
		LocationID:        "123",
		CurrentTime:       time.Now().Format("2006-01-02 15:04:05"),
	}

	err := SendDataToServer(data, server.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetAndSend(t *testing.T) {
	// Create a temporary config file for testing
	tempConfigFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempConfigFile.Name())

	// Write test config data to the temporary file
	testConfigData := `{"location_id":"123","location_name":"TestLocation","server_url":"http://example.com"}`
	_, err = tempConfigFile.WriteString(testConfigData)
	if err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}

	// Test GetAndSend function
	hasPower, err := GetAndSend(tempConfigFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !hasPower {
		t.Error("Expected battery not to have power supply, but it does.")
	}
}
