# 🚀 Автоматизоване розгортання Cimeika Central Command System

## Швидкий старт

### Для Ubuntu Server (simei@Cihub:~$)

```bash
# 1. Завантажити скрипт
wget https://raw.githubusercontent.com/your-repo/cimeika/main/auto_deploy_cimeika.sh
# або
curl -O https://raw.githubusercontent.com/your-repo/cimeika/main/auto_deploy_cimeika.sh

# 2. Надати права виконання
chmod +x auto_deploy_cimeika.sh

# 3. Запустити автоматизоване розгортання
./auto_deploy_cimeika.sh

# Додаткові опції:
./auto_deploy_cimeika.sh --quick          # Швидкий режим
./auto_deploy_cimeika.sh --skip-docker    # Без Docker
./auto_deploy_cimeika.sh --skip-update    # Без оновлення системи
```

### Для Windows (F:\Cimeika_Project_Structure)

```cmd
REM 1. Завантажити batch файл
REM Зберегти auto_deploy_cimeika.bat у будь-яку директорію

REM 2. Запустити від імені адміністратора
auto_deploy_cimeika.bat

REM Або подвійний клік по файлу в провіднику
```

### Альтернативний спосіб для Windows через WSL

```bash
# В WSL терміналі
cd /mnt/f
./auto_deploy_cimeika.sh --quick
```

## Детальний процес

### Що робить автоматизований скрипт:

#### 1. Підготовка середовища
- ✅ Детекція операційної системи
- ✅ Перевірка та встановлення залежностей
- ✅ Оновлення системних пакетів
- ✅ Встановлення Python 3.9+, pip, git
- ✅ Встановлення Docker (опціонально)

#### 2. Створення структури проекту
```
Cimeika_Project_Structure/
├── 01_Project_Core/        # Основний код системи
│   ├── modules/
│   ├── utils/
│   ├── config/
│   └── main.py            # Головний модуль
├── 02_Development/         # Розробка та тести
│   ├── tests/
│   ├── scripts/
│   └── quick_test.py      # Швидкий тест
├── 03_AI_Personas/         # AI персонажі
├── 05_Integrations/        # Інтеграції
├── 08_Deployment/          # Розгортання
│   ├── scripts/           # Скрипти управління
│   ├── configs/           # Конфігурації
│   └── docker/            # Docker файли
├── logs/                   # Логи системи
├── data/                   # Дані та кеш
└── certificates/           # SSL сертифікати
```

#### 3. Конфігурація системи
- ✅ Створення `.env` файлу з базовими налаштуваннями
- ✅ Генерація безпечних ключів шифрування
- ✅ Налаштування віртуального Python середовища
- ✅ Встановлення всіх необхідних залежностей
- ✅ Створення Docker Compose конфігурації

#### 4. Створення основних модулів
- ✅ Flask веб-сервер з API endpoints
- ✅ Health check система
- ✅ Redis інтеграція
- ✅ PostgreSQL підтримка
- ✅ Prometheus моніторинг

#### 5. Скрипти управління
- ✅ `start.sh`/`start.bat` - Запуск системи
- ✅ `stop.sh` - Зупинка системи (Linux)
- ✅ `restart.sh` - Перезапуск системи (Linux)
- ✅ `status.sh`/`status.bat` - Статус системи

## Після розгортання

### 1. Перевірка роботи

```bash
# Linux
./08_Deployment/scripts/status.sh

# Windows
08_Deployment\scripts\status.bat

# Або через браузер
curl http://localhost:8000/health
```

### 2. Налаштування

Відредагуйте файл `.env`:

```bash
# Безпека (ОБОВ'ЯЗКОВО для production!)
CIMEIKA_MASTER_KEY=your-generated-key-here
JWT_SECRET=your-jwt-secret-here

# Зовнішні сервіси
OPENAI_API_KEY=your-openai-api-key
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
GOOGLE_CALENDAR_CREDENTIALS=certificates/google_credentials.json

# База даних (для production)
DATABASE_URL=postgresql://user:password@localhost:5432/cimeika

# Email сповіщення
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
ALERT_EMAILS=admin@example.com
```

### 3. Запуск системи

#### Режим розробки (Linux):
```bash
# Активація віртуального середовища
source venv/bin/activate

# Запуск
python 01_Project_Core/main.py

# Або через скрипт
./08_Deployment/scripts/start.sh
```

#### Режим розробки (Windows):
```cmd
REM Активація віртуального середовища
venv\Scripts\activate.bat

REM Запуск
python 01_Project_Core\main.py

REM Або через скрипт
08_Deployment\scripts\start.bat
```

#### Production режим (Docker):
```bash
# Запуск всіх сервісів
docker-compose up -d

# Перегляд логів
docker-compose logs -f cimeika-central

# Зупинка
docker-compose down
```

### 4. Тестування

```bash
# Швидкий тест
python 02_Development/quick_test.py

# Повні тести
python -m pytest 02_Development/tests/

# Windows
02_Development\quick_test.bat
```

## API Endpoints

Після запуску доступні наступні endpoints:

- **Health Check**: `GET http://localhost:8000/health`
- **System Info**: `GET http://localhost:8000/system/info`
- **Component Registration**: `POST http://localhost:8000/system/register`
- **Persona Transformation**: `POST http://localhost:8000/ci/transform`

### Приклади використання:

```bash
# Перевірка здоров'я системи
curl http://localhost:8000/health

# Отримання інформації про систему
curl http://localhost:8000/system/info

# Реєстрація нового компонента
curl -X POST http://localhost:8000/system/register \
  -H "Content-Type: application/json" \
  -d '{
    "component_name": "My Component",
    "component_type": "chat_interface",
    "access_level": "user"
  }'

# Трансформація персонажа
curl -X POST http://localhost:8000/ci/transform \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "message": "Потрібно запланувати зустріч",
    "target_persona": "podia"
  }'
```

## Моніторинг

- **API Health**: http://localhost:8000/health
- **Prometheus Metrics**: http://localhost:9090
- **Логи**: `logs/cimeika.log`
- **System Status**: `./08_Deployment/scripts/status.sh`

## Troubleshooting

### Поширені проблеми:

#### 1. Python не знайдено
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install python3 python3-pip python3-venv

# CentOS/RHEL
sudo yum install python3 python3-pip

# Windows
# Завантажте з https://python.org/downloads/
# Або через Microsoft Store
```

#### 2. Порт 8000 зайнятий
```bash
# Знайти процес що використовує порт
sudo netstat -tulpn | grep :8000
# або
sudo lsof -i :8000

# Змінити порт в .env файлі
API_PORT=8001
```

#### 3. Redis недоступний
```bash
# Встановити Redis
sudo apt install redis-server

# Або використати Docker
docker run -d -p 6379:6379 redis:alpine

# Або відключити Redis в .env
REDIS_URL=""
```

#### 4. Проблеми з правами доступу (Linux)
```bash
# Додати користувача до групи docker
sudo usermod -aG docker $USER

# Змінити власника директорії
sudo chown -R $USER:$USER /home/simei/Cimeika_Project_Structure

# Перелогінитися
logout
```

## Розширена конфігурація

### Systemd сервіс (Linux)
```bash
# Встановити як системний сервіс
sudo cp 08_Deployment/configs/cimeika.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable cimeika
sudo systemctl start cimeika

# Статус сервісу
sudo systemctl status cimeika
```

### Nginx reverse proxy
```bash
# Встановити Nginx
sudo apt install nginx

# Копіювати конфігурацію
sudo cp 08_Deployment/configs/nginx.conf /etc/nginx/sites-available/cimeika
sudo ln -s /etc/nginx/sites-available/cimeika /etc/nginx/sites-enabled/
sudo systemctl reload nginx
```

### SSL сертифікати
```bash
# Let's Encrypt (безкоштовні SSL)
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com

# Або розмістити сертифікати в certificates/
```

## Підтримка

- **Документація**: `06_Documentation/`
- **Логи розгортання**: `logs/auto_deploy_*.log`
- **GitHub Issues**: [посилання на репозиторій]
- **Telegram**: @cimeika_support

## Ліцензія

Створено автоматичним скриптом розгортання Cimeika v2.0

---

**⚠️ Важливо**: Перед використанням в production обов'язково:
1. Оновіть всі ключі безпеки в `.env`
2. Налаштуйте SSL сертифікати
3. Обмежте доступ до системи
4. Налаштуйте бекапи бази даних
5. Протестуйте всі функції системи