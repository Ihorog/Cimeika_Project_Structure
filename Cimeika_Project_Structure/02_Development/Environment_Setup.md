# Налаштування середовища розробки

## Docker конфігурація

### Dockerfile
`dockerfile
FROM python:3.9
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["gunicorn", "--bind", "0.0.0.0:8000", "app:app"]
`

### DevContainer
- Python 3.9 + JDK 21
- Налаштування для VS Code
- Готове середовище для розробки

## Environment Variables (.env)

`env
# OpenAI API
OPENAI_API_KEY=your_openai_key

# Google Calendar
GOOGLE_CALENDAR_API_KEY=your_google_key

# Dropbox
DROPBOX_ACCESS_TOKEN=your_dropbox_token

# Weather
OPENWEATHER_API_KEY=your_weather_key

# Astrology  
ASTROLOGY_API_KEY=your_astrology_key

# Health
HEALTH_API_KEY=your_health_key

# Telegram
TELEGRAM_BOT_TOKEN=your_telegram_token
`

## Залежності

### Python (requirements.txt)
`
Flask==2.3.0
Gunicorn==20.1.0
openai==0.27.0
google-api-python-client==2.86.0
dropbox==11.36.0
requests==2.31.0
python-telegram-bot==20.3.0
`

### Запуск локально
`ash
# Встановлення залежностей
pip install -r requirements.txt

# Запуск Flask додатку
python app.py

# Або через Gunicorn
gunicorn --bind 0.0.0.0:8000 app:app
`
"@

    # =================== AI Personas ===================
    "03_AI_Personas\Ci_Core.md" = @"
# Ci - Центральний асистент

## Концепція

Ci є єдиним інтерфейсом взаємодії з усією платформою Cimeika. Він не просто виконує команди, а розуміє контекст, навчається та адаптується до потреб користувача.

## Можливості

### Комунікація
- **Текстовий чат**: природне спілкування українською мовою
- **Голосовий інтерфейс**: розпізнавання та синтез мови
- **Файловий обмін**: робота з документами, зображеннями
- **Генерація контенту**: створення текстів, порад, рекомендацій

### Трансформації
Ci може перетворюватися на спеціалізованих персонажів:
- **ПоДія**: при роботі з календарем та подіями
- **Настрій**: при обговоренні емоцій та самопочуття  
- **Маля**: при творчих завданнях з дітьми
- **Казкар**: [потребує специфікації]

### Адаптивність
- Аналіз історії взаємодій
- Врахування емоційного стану
- Персоналізація рекомендацій
- Контекстна підказка інструментів

## Технічна реалізація

### OpenAI інтеграція
`python
# Базовий промпт для Ci
base_prompt = \"\"\"
Ти - Ci, центральний асистент платформи Cimeika.
Твоя роль: розумний помічник, що допомагає користувачу 
організувати життя через планування, настрій, творчість.
\"\"\"
`

### Система трансформацій
`python
personas = {
    'ci': {'prompt': base_prompt, 'tools': ['general']},
    'podiya': {'prompt': event_prompt, 'tools': ['calendar', 'planning']},
    'nastriy': {'prompt': mood_prompt, 'tools': ['wellness', 'meditation']},
    'malya': {'prompt': creative_prompt, 'tools': ['games', 'art']},
    'kazkar': {'prompt': story_prompt, 'tools': ['storytelling']}
}
`

### Контекстна пам'ять
- Збереження історії розмов
- Профіль користувача та вподобання
- Аналіз патернів використання
- Персоналізація відповідей

## Візуальна ідентичність

### Основний образ Ci
- Логотип та фірмовий стиль
- Кольорова схема (основна палітра)
- Анімації переходів між персонажами

### Персонажні трансформації
- Зміна кольорової схеми інтерфейсу
- Адаптація стилю спілкування
- Спеціалізовані іконки та візуальні елементи
