# API Документація Cimeika

## Базовий URL
`
http://localhost:8000/api
`

## Аутентифікація
[Визначити схему аутентифікації]

## Endpoints

### Ci Assistant

#### POST /chat/completion
Отримання відповіді від Ci асистента

**Request:**
`json
{
    'message': 'Привіт, Ci!',
    'persona': 'ci',  // опціонально
    'context': {}     // опціонально
}
`

**Response:**
`json
{
    'response': 'Привіт! Як справи?',
    'persona': 'ci',
    'suggestions': ['Створити подію', 'Перевірити настрій']
}
`

### ПоДія (Події)

#### POST /podia/events
Створення нової події

**Request:**
`json
{
    'title': 'Зустріч з друзями',
    'date': '2024-07-17',
    'time': '18:00',
    'location': 'Кафе центр',
    'notes': 'Не забути подарунок'
}
`

#### GET /podia/events
Отримання списку подій

**Query params:**
- date_from (опціонально)
- date_to (опціонально)
- limit (опціонально)

### Настрій

#### POST /nastriy/mood
Фіксація настрою

**Request:**
`json
{
    'mood_level': 7,
    'description': 'Гарний день, багато енергії',
    'activities': ['прогулянка', 'робота']
}
`

#### GET /nastriy/advice
Отримання ШІ-поради

**Request:**
`json
{
    'current_mood': 'стрес',
    'context': 'багато роботи'
}
`

### Маля (Творчість)

#### GET /mala/tasks
Отримання творчих завдань

#### POST /mala/upload
Завантаження творчої роботи

**Form data:**
- file: зображення
- title: назва роботи
- category: категорія творчості

### Дані

#### GET /data/weather
Поточна погода

#### GET /data/astrology
Астрологічний прогноз

## Status Codes

- 200: Успішно
- 400: Помилка запиту
- 401: Неавторизований доступ
- 500: Серверна помилка

## Rate Limiting

- 100 запитів на хвилину для chat endpoints
- 1000 запитів на годину для інших endpoints
