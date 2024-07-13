## Название проекта: Система управления данными о скважинах

**Проект для автоматизации управления данными о скважинах и объектах в нефтедобывающей отрасли.**

### Описание:

Проект представляет собой серверное приложение, реализованное на языке Go, которое предоставляет REST API для работы с данными о скважинах, объектах, истории скважин за день и планами скважин за день.  

Данные хранятся в реляционной базе данных PostgreSQL. API спроектировано с использованием стандартных принципов REST и документировано с помощью Swagger.

### Основные компоненты:

* **Backend на Go:**  Серверная часть, реализующая логику обработки запросов к базе данных.
* **REST API**:  Стандартизированный HTTP интерфейс для взаимодействия с приложением.
* **PostgreSQL**: База данных для хранения и управления данными о скважинах, объектах, истории и планах.
* **Swagger**:  Инструмент для генерации интерактивной документации API.

### Примеры использования API:

#### **Работа с объектами:**

* **Получение всех объектов:**
  ```bash
  curl -X GET http://localhost:8080/objects
  ```

* **Создание нового объекта:**
  ```bash
  curl -X POST http://localhost:8080/objects -H "Content-Type: application/json" -d "{\"id\":1, \"name\":\"New Object\", \"type\":1}" 
  ```

* **Обновление объекта:**
  ```bash
  curl -X PUT http://localhost:8080/objects -H "Content-Type: application/json" -d "{\"id\":1, \"name\":\"Updated Object\", \"type\":2}"
  ```

* **Удаление объекта:**
  ```bash
  curl -X DELETE "http://localhost:8080/objects?id=1"
  ```

#### **Работа со скважинами:**

* **Получение всех скважин:**
  ```bash
  curl -X GET http://localhost:8080/wells
  ```

* **Создание новой скважины:**
  ```bash
  curl -X POST http://localhost:8080/wells -H "Content-Type: application/json" -d "{\"well\":1, \"ngdu\":1, \"cdng\":1, \"kust\":1, \"mest\":1}"
  ```

* **Обновление скважины:**
  ```bash
  curl -X PUT http://localhost:8080/wells -H "Content-Type: application/json" -d "{\"well\":1, \"ngdu\":2, \"cdng\":2, \"kust\":2, \"mest\":2}"
  ```

* **Удаление скважины:**
  ```bash
  curl -X DELETE "http://localhost:8080/wells?well=1"
  ```

#### **Работа с историей скважин за день:**

* **Получение всех исторических данных:**
  ```bash
  curl -X GET http://localhost:8080/well_day_histories
  ```

* **Создание новой записи истории:**
  ```bash
  curl -X POST http://localhost:8080/well_day_histories -H "Content-Type: application/json" -d "{\"well\":4455, \"date_fact\":\"2024-12-10\", \"debit\":10, \"ee_consume\":50.5, \"expenses\":5.123, \"pump_operating\":10}"
  ```

* **Обновление записи истории:**
  ```bash
  curl -X PUT http://localhost:8080/well_day_histories -H "Content-Type: application/json" -d "{\"well\":4455, \"date_fact\":\"2024-12-10\", \"debit\":15, \"ee_consume\":55.5, \"expenses\":5.678, \"pump_operating\":12}"
  ```

* **Удаление записи истории:**
  ```bash
  curl -X DELETE "http://localhost:8080/well_day_histories?well=4455&date_fact=2024-12-10"
  ```

#### **Работа с планами скважин за день:**

* **Получение всех плановых данных:**
  ```bash
  curl -X GET http://localhost:8080/well_day_plans
  ```

* **Создание нового плана:**
  ```bash
  curl -X POST http://localhost:8080/well_day_plans -H "Content-Type: application/json" -d "{\"well\":4455, \"date_plan\":\"2024-12-10\", \"debit\":20, \"ee_consume\":60.5, \"expenses\":6.789, \"pump_operating\":15}"
  ```

* **Обновление плана:**
  ```bash
  curl -X PUT http://localhost:8080/well_day_plans -H "Content-Type: application/json" -d "{\"well\":4455, \"date_plan\":\"2024-12-10\", \"debit\":25, \"ee_consume\":65.5, \"expenses\":7.891, \"pump_operating\":18}"
  ```

* **Удаление плана:**
  ```bash
  curl -X DELETE "http://localhost:8080/well_day_plans?well=4455&date_plan=2024-12-10"
  ```

### Дополнительная информация:

* Документация API доступна по адресу: `http://localhost:8080/swagger/index.html`.