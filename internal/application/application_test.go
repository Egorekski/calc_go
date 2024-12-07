package application

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO: write README.md file
// TODO: Make tests for `application.go`
func TestCalcHandler(t *testing.T) {
	testRequestBody := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}
	for i := 0; i < len(testRequestBody); i++ {
		expr := testRequestBody[i].expression
		requestBody := strings.NewReader(`{"expression":"` + expr + `"}`)
		req, err := http.NewRequest("POST", "/calc", requestBody)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		CalcHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		gotRes := strings.Split(string(body), " ")
		if gotRes[1] != fmt.Sprintf("%f", testRequestBody[i].expectedResult) {
			t.Errorf("expected %f, got %v", testRequestBody[i].expectedResult, string(body))
		}
	}
}
