# Backend завдання

## Реалізації Backend

### Python Flask API
- [ ] Налаштування базової структури Flask додатку
- [ ] Реалізація REST endpoints згідно OpenAPI специфікації
- [ ] Middleware для обробки запитів та відповідей
- [ ] Error handling та логування

### Go API (Alternative Implementation)
- [x] Базова структура Go додатку з Gin framework
- [x] REST endpoints для всіх основних функцій
- [x] Middleware для CORS, логування та обробки помилок
- [x] Інтеграція з OpenAI API
- [x] Система AI персонажів (Ci, ПоДія, Настрій, Маля, Казкар)
- [x] Docker конфігурація для Go backend
- [ ] Реалізація повної функціональності інтеграцій

### Ci-асистент
- [ ] Інтеграція з OpenAI API
- [ ] Система промптів для різних персонажів  
- [ ] Механізм трансформації між персонажами
- [ ] Збереження контексту розмов

### API інтеграції
- [ ] Google Calendar API налаштування
- [ ] Dropbox API для медіа файлів
- [ ] OpenWeatherMap API інтеграція
- [ ] Health API підключення
- [ ] FreeAstrologyAPI налаштування

### Telegram Bot
- [ ] Створення та налаштування бота
- [ ] Команди для управління подіями
- [ ] Завантаження файлів через бота
- [ ] Синхронізація з веб-інтерфейсом

## Технічні деталі

### Структура проекту

#### Python Flask
`
/app
  /api
    /endpoints
    /models
    /services
  /core
  /integrations
  /utils
`

#### Go Backend
`
/Backend_Go
  /cmd                 # Application entry point
  /internal
    /api              # HTTP handlers and routes
      /handlers       # Request handlers
      /middleware     # Middleware components
    /core             # Configuration and services
    /integrations     # External API clients  
    /models           # Data models
    /services         # Business logic
  /pkg                # Shared utilities
`

### Конфігурація
- Environment variables (.env)
- API ключі та секрети
- Docker налаштування
