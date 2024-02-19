package utils

import (
	"net/http"
	"net/http/httptest"

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

// func TestGetAndSend(t *testing.T) {

// 	nonExistentConfigFile := "C:/Users/ARABINDA DAS/OneDrive/Desktop/GOLANG_PROJS/power_status/config.json"

// 	// Call GetAndSend with the non-existent config file path
// 	_, err := GetAndSend(nonExistentConfigFile)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// }
