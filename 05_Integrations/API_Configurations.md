# Конфігурації API

## OpenAI API

### Налаштування
`python
import openai

openai.api_key = os.getenv('OPENAI_API_KEY')

# Базові параметри для Ci
ci_config = {
    'model': 'gpt-4',
    'temperature': 0.7,
    'max_tokens': 1000,
    'presence_penalty': 0.1
}
`

### Промпти персонажів
`python
personas_prompts = {
    'ci': 'Ти Ci - центральний асистент платформи Cimeika...',
    'podiya': 'Ти ПоДія - експерт планування та організації...',
    'nastriy': 'Ти Настрій - турботливий помічник ментального здоров\'я...',
    'malya': 'Ти Маля - веселий помічник дитячої творчості...',
    'kazkar': 'Ти Казкар - [потребує специфікації]...'
}
`

## Google Calendar API

### Аутентифікація
`python
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build

# OAuth 2.0 flow
credentials = Credentials.from_authorized_user_info(token_data)
service = build('calendar', 'v3', credentials=credentials)
`

### Основні операції
`python
# Створення події
event = {
    'summary': 'Назва події',
    'start': {'dateTime': '2024-07-17T10:00:00'},
    'end': {'dateTime': '2024-07-17T11:00:00'},
}
service.events().insert(calendarId='primary', body=event).execute()
`

## Dropbox API

### Налаштування
`python
import dropbox

dbx = dropbox.Dropbox(os.getenv('DROPBOX_ACCESS_TOKEN'))

# Завантаження файлу
def upload_file(file_data, filename):
    path = f'/Cimeika/Creativity/{filename}'
    return dbx.files_upload(file_data, path)
`

## Telegram Bot API

### Налаштування
`python
from telegram import Bot, Update
from telegram.ext import Application, CommandHandler

bot_token = os.getenv('TELEGRAM_BOT_TOKEN')
application = Application.builder().token(bot_token).build()

# Команди
application.add_handler(CommandHandler('create_event', create_event_command))
application.add_handler(CommandHandler('view_events', view_events_command))
`

## Інші API

### OpenWeatherMap
`python
weather_api_key = os.getenv('OPENWEATHER_API_KEY')
weather_url = f'http://api.openweathermap.org/data/2.5/weather'
`

### Health API
`python
health_api_key = os.getenv('HEALTH_API_KEY')
# Специфічні налаштування залежно від провайдера
`

## Error Handling

### Загальна стратегія
`python
def api_call_with_retry(api_func, max_retries=3):
    for attempt in range(max_retries):
        try:
            return api_func()
        except Exception as e:
            if attempt == max_retries - 1:
                raise e
            time.sleep(2 ** attempt)
`

### Rate Limiting
- OpenAI: моніторинг використання токенів
- Google Calendar: дотримання квот API
- Telegram: обмеження повідомлень на секунду
