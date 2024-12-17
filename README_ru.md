# Калькулятор

**Калькулятор** — это простой калькулятор, написанный на Go, реализованный как HTTP-сервер и CLI-приложение. Он выполняет базовые арифметические операции (сложение, вычитание, умножение, деление) и возвращает результаты в формате JSON для HTTP или напрямую в консоли для CLI.

---

## Перевод

- en [English](README.md)
- ru [Русский](README_ru)

---

## Возможности

- Разработан на **Go**, что обеспечивает высокую производительность и простоту.
- Поддерживает два режима работы:
    - **Режим HTTP-сервера** для API-интеграций.
    - **Режим CLI** для быстрых вычислений через терминал.
- Поддержка базовых операций:
    - **Сложение**
    - **Вычитание**
    - **Умножение**
    - **Деление**
- Работает с числами в формате `float`.
- Простые и интуитивно понятные интерфейсы как для HTTP, так и для CLI.
- Обработка ошибок, включая деление на ноль.

---

## Начало работы

### Требования

- **Go** (минимальная версия: `1.18`)  
  [Скачать и установить Go](https://golang.org/dl/).

---

### Установка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/your-username/http-calculator.git
  

---

### Использование
#### Режим HTTP-сервера

1. Запустите сервер:

   Раскомментируйте `app.RunServer()` и закомментируйте `app.Run()` в файле `main.go`.

   Сервер запустится по адресу http://localhost:8080/api/v1/calculate. Вы также можете указать свой порт, задав переменную окружения.  
   Пример:
    ```bash
    export PORT='ваш-порт' && go run ./cmd/main.go 
    ```

2. Используйте HTTP-клиент (например, curl или Postman) для взаимодействия с API.

#### Формат запроса

Отправьте POST-запрос с JSON-данными, содержащими строку с математическим выражением:
```json
{
  "expression": "2+2"    
}
```

#### Пример запроса

```bash
curl -X POST http://localhost:8080 \
-H "Content-Type: application/json" \
-d '{"expression": "2+2"}'
```

#### Ответ

```json
{
    "result": 4
}
```

---

### Режим CLI

1. Запустите приложение в режиме CLI:

   Раскомментируйте `app.Run()` и закомментируйте `app.RunServer()` в файле `main.go`.  
   Затем выполните команду для запуска приложения:
    ```bash
    go run ./cmd/main.go
    ```

2. Следуйте интерактивным подсказкам для выполнения операций. Например:
```bash
2024/12/07 11:48:33 Welcome to the CLI Calculator!
2024/12/07 11:48:33 input expression
1+1
2024/12/07 11:48:37 1+1 = 2
2024/12/07 11:48:37 input expression
exit
2024/12/07 11:48:43 aplication was successfully closed
```
3. Введите `exit`, чтобы завершить работу программы.

---

### Структура проекта

```graphql
calc_go/
├── cmd/ 
     └── main.go                          # Точка входа в приложение
├── internal/                             # Содержит реализацию HTTP-сервера
        └── application/
                ├── application.go        # Логика приложения
                └── application_test.go   # Тесты для приложения
├── pkg/           
     └── calculation/
              ├── calculation.go          # Утилиты для выполнения вычислений
              ├── calculation_test.go     # Тесты для вычислений
              └── errors.go               # Обработка ошибок
└── README.md                             # Документация проекта

```

---

### План развития проекта
- Добавить поддержку расширенных математических операций (например, возведение в степень, квадратный корень)
- Реализовать логирование запросов и результатов.
- Расширить режим CLI дополнительными интерактивными функциями.

---

### Вклад

Мы рады вашему участию! Чтобы внести вклад:

1. Форкните репозиторий.
2. Создайте новую ветку для своей функции: `git checkout -b feature-name`. 
3. Зафиксируйте изменения: `git commit -m "Добавлена новая функция"`. 
4. Опубликуйте ветку: `git push origin feature-name`. 
5. Откройте `Pull Request`.

---
### Лицензия
Проект распространяется по лицензии MIT. Подробности можно найти в файле LICENSE.