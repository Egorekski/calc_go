package application

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pashapdev/calc_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// Функция запуска приложения
// тут будем чиать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	log.Println("Welcome to the CLI Calculator!")
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}

		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)

		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// 404 error
	if r.URL.Path != "/api/v1/calculate" {
		http.NotFound(w, r)
		return
	}

	// 405 error
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

		// 422 error
		if err.Error() == ErrUnmarshalBool ||
			err.Error() == ErrUnmarshalNumber ||
			err.Error() == ErrUnmarshalObject {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 418 error
	if request.Expression == "coffee" {
		http.Error(w, http.StatusText(http.StatusTeapot), http.StatusTeapot)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) {
			http.Error(w, fmt.Sprintf("error: %v", err.Error()), http.StatusBadRequest)
		} else if errors.Is(err, calculation.ErrEmptyExpression) {
			http.Error(w, fmt.Sprintf("error: %v", err.Error()), http.StatusBadRequest)
		} else if errors.Is(err, calculation.ErrDivisionByZero) {
			http.Error(w, fmt.Sprintf("error: %v", err.Error()), http.StatusBadRequest)
		} else {
			http.Error(w, fmt.Sprintf("unknown error: %v", err.Error()), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, fmt.Sprintf("result: %f", result))
		if err != nil {
			return
		}
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
