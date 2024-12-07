package application

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	testSuccessRequestBody := []struct {
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
	for i := 0; i < len(testSuccessRequestBody); i++ {
		expr := testSuccessRequestBody[i].expression
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
		if gotRes[1] != fmt.Sprintf("%f", testSuccessRequestBody[i].expectedResult) {
			t.Errorf("expected %f, got %v", testSuccessRequestBody[i].expectedResult, string(body))
		}
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr string
	}{
		{
			name:        "simple",
			expression:  "1+1*",
			expectedErr: "invalid expression\n",
		},
		{
			name:        "wrong expression",
			expression:  "((2+2-*(2",
			expectedErr: "invalid expression\n",
		},
		{
			name:        "empty",
			expression:  "",
			expectedErr: "empty expression\n",
		},
	}

	for i := 0; i < len(testCasesFail); i++ {
		expr := testCasesFail[i].expression
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

		gotRes := strings.Split(string(body), ": ")
		if gotRes[1] != fmt.Sprintf("%v", testCasesFail[i].expectedErr) {
			t.Errorf("expression \"%v\" is invalid, but result \"%v\" was obtained", testCasesFail[i].expression, gotRes[1])
		}
	}
}
