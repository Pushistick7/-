import requests

# URL сервисов через OpenResty
service1_url = "http://localhost:8080/service1/message"
service2_url = "http://localhost:8080/service2/message"

# Начало диалога
print("Диалог начался...\n")

# Сервис 1 отправляет первое сообщение
response1 = requests.post(service1_url, json={"message": ""})
print(f"Сервис 1: {response1.json()['message']}")

# Сервис 2 отвечает на первое сообщение
response2 = requests.post(service2_url, json={"message": response1.json()['message']})
print(f"Сервис 2: {response2.json()['message']}")

# Сервис 1 отвечает на второе сообщение
response3 = requests.post(service1_url, json={"message": response2.json()['message']})
print(f"Сервис 1: {response3.json()['message']}")

# Сервис 2 отвечает на третье сообщение
response4 = requests.post(service2_url, json={"message": response3.json()['message']})
print(f"Сервис 2: {response4.json()['message']}")

# Сервис 1 отвечает на четвёртое сообщение (завершение диалога)
response5 = requests.post(service1_url, json={"message": response4.json()['message']})
print(f"Сервис 1: {response5.json()['message']}")
