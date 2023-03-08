package main_test

import (
	"log"
	"os"
	"testing"
	"main/app"
)

var a App

// Testing to check if database is properly set up before test and cleaned after tests are finished.
func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM tickets")
	a.DB.Exec("ALTER SEQUENCE tickets_serial_number_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS tickets
(
	serial_number SERIAL,
	device_type VARCHAR ( 50 ) NOT NULL,
	device_info TEXT NOT NULL,
	CONSTRAINT tickets_pkey PRIMARY KEY (serial_number)
)`


// func TestEmptyTable(t *testing.T) {
// 	clearTable()

// 	req, _ := http.NewRequest("GET", "/tickets", nil)
//     response := executeRequest(req)

//     checkResponseCode(t, http.StatusOK, response.Code)

//     if body := response.Body.String(); body != "[]" {
//         t.Errorf("Expected an empty array. Got %s", body)
//     }
// }

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	a.Router.ServeHTTP(rr, req)

// 	return rr
// }

// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 			t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }