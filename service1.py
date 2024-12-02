from flask import Flask, request, jsonify

app = Flask(__name__) #экземпляр приложения

# Переменная для лучшего понимания, как сервисам общаться
dialogue_state = 0

# Имена участников диалога
service1_name = "Евгений Александрович"
service2_name = "Ксения Николаевна"

@app.route('/reset', methods=['POST'])
def reset_dialogue():
    global dialogue_state
    print("Запрос на сброс состояния получен")  # Отладочное сообщение, чтобы заново запускать беседу
    dialogue_state = 0  
    response = jsonify({"message": "Состояние диалога сброшено."})
    response.headers["Content-Type"] = "application/json; charset=utf-8"
    return response

@app.route('/message', methods=['POST'])
def receive_message():
    global dialogue_state
    print(f"[DEBUG] Текущее состояние диалога: {dialogue_state}")
    incoming_message = request.json.get('message', '').strip()
    print(f"[DEBUG] Сервис 1 получил сообщение: {incoming_message}")

    response = "Неизвестное сообщение."  # Ответ по умолчанию, если никто друг друга не понял

#логика диалога
    if dialogue_state == 0:
        # Начало диалога
        response = "Добрый день, как вас зовут?"
        dialogue_state += 1
    elif dialogue_state == 1 and "как зовут вас" in incoming_message.lower():
        # Ответ на вопрос, который задал сервер 2
        response = f"Приятно познакомиться, {service2_name}, мое имя {service1_name}. Давайте выпьем чаю?"
        dialogue_state += 1
    elif dialogue_state == 2 and "я знаю одну хорошую пекарню" in incoming_message.lower():
        # Завершение диалога
        response = "Отлично, пекарня звучит замечательно!"
        dialogue_state += 1

    print(f"[DEBUG] Сервис 1 отправляет сообщение: {response}")
    return jsonify({"message": response})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001) #запуск сервиса