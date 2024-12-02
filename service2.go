package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// глобальная переменная для общения серверов
var dialogueState int = 0

// структура сообщения Go, происходит сериализация(преобразуем значение Go в кусочек байтов(строку))
type Message struct {
	Message string `json:"message"`
}

// отладочное сообщение, чтобы заново запускать беседу
func resetDialogue(w http.ResponseWriter, r *http.Request) {
	dialogueState = 0
	fmt.Println("[DEBUG] Запрос на сброс состояния получен")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Состояние диалога сброшено."})
}

// эта функция обрабатывает входящие сообщения, читает тело запроса и сохраняет в переменную body
func handleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// десериализация JSON (преобразуем кусочек байтов(строку) в значение Golang)
	var incoming Message
	json.Unmarshal(body, &incoming) //
	incomingMessage := strings.TrimSpace(strings.ToLower(incoming.Message))
	fmt.Printf("[DEBUG] Сервис 2 получил сообщение: %s\n", incomingMessage)

	//вся логика диалога
	var response string

	if dialogueState == 0 && incomingMessage == "добрый день, как вас зовут?" {
		response = "Меня зовут Ксения Николаевна, а как зовут вас?"
		dialogueState = 1
		fmt.Println("[DEBUG] Состояние диалога обновлено: 1")
	} else if dialogueState == 1 && incomingMessage == "приятно познакомиться, ксения николаевна, мое имя евгений александрович. давайте выпьем чаю?" {
		response = "С удовольствием, Евгений Александрович, я знаю одну хорошую пекарню."
		dialogueState = 2
		fmt.Println("[DEBUG] Состояние диалога обновлено: 2")
	} else if dialogueState == 2 && strings.Contains(incomingMessage, strings.ToLower("Отлично, пекарня звучит замечательно!")) {
		response = "Договорились! До встречи!"
		dialogueState = 3
		fmt.Println("[DEBUG] Состояние диалога обновлено: 3")
	} else {
		response = "Неизвестное сообщение."
		fmt.Println("[DEBUG] Неизвестное сообщение получено.")
	}

	//формируем ответ, решила добавить это отладочное сообщение, чтобы было видно, как происходит диалог
	fmt.Printf("[DEBUG] Сервис 2 отправляет сообщение: %s\n", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: response})
}

// функция запускает сервер на порту 5002 и привязывает функции обработки к маршрутам /reset и /message
func main() {
	http.HandleFunc("/reset", resetDialogue)
	http.HandleFunc("/message", handleMessage)
	fmt.Println("Сервис 2 запущен на порту 5002...")
	log.Fatal(http.ListenAndServe(":5002", nil))
}
