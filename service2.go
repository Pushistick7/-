package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var dialogueState int = 0

type Message struct {
	Message string `json:"message"`
}

func resetDialogue(w http.ResponseWriter, r *http.Request) {
	dialogueState = 0
	fmt.Println("[DEBUG] Запрос на сброс состояния получен")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Состояние диалога сброшено."})
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var incoming Message
	json.Unmarshal(body, &incoming)
	incomingMessage := strings.TrimSpace(strings.ToLower(incoming.Message))
	fmt.Printf("[DEBUG] Сервис 2 получил сообщение: %s\n", incomingMessage)

	var response string

	if dialogueState == 0 && incomingMessage == "добрый день, как вас зовут?" {
		response = "Меня зовут Ксения Николаевна, а как зовут вас?"
		dialogueState = 1
		fmt.Println("[DEBUG] Состояние диалога обновлено: 1")
	} else if dialogueState == 1 && incomingMessage == "приятно познакомиться, ксения николаевна, мое имя евгений александрович. давайте выпьем чаю?" {
		response = "С удовольствием, Евгений Александрович, я знаю одну хорошую пекарню."
		dialogueState = 2
		fmt.Println("[DEBUG] Состояние диалога обновлено: 2")
	} else if dialogueState == 2 && strings.Contains(incomingMessage, "Отлично, пекарня звучит замечательно!") {
		response = "Договорились! До встречи!"
		dialogueState = 3
		fmt.Println("[DEBUG] Состояние диалога обновлено: 3")
	} else {
		response = "Неизвестное сообщение."
		fmt.Println("[DEBUG] Неизвестное сообщение получено.")
	}

	fmt.Printf("[DEBUG] Сервис 2 отправляет сообщение: %s\n", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: response})
}

func main() {
	http.HandleFunc("/reset", resetDialogue)
	http.HandleFunc("/message", handleMessage)
	fmt.Println("Сервис 2 запущен на порту 5002...")
	log.Fatal(http.ListenAndServe(":5002", nil))
}
