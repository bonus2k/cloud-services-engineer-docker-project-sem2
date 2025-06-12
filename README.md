# Cloud Services Engineer - Docker Project (Semester 2)

## 📋 Описание проекта

В рамках лабораторной работы реализован multi-container проект:

- Frontend (Vue.js)
- Backend (Go)
- Traefik (reverse proxy + маршрутизация)
- Healthcheck endpoints
- Разделение конфигураций dev/prod по разным Compose файлам
- Сборка, проверка и публикация образов автоматизированы с помощью GitHub Actions (CI/CD)

---

## 🚀 Запуск проекта

### 1️⃣ Профиль разработки (dev)

Использует отдельный Compose файл `docker-compose.dev.yml`:

```bash
docker compose -f docker-compose.dev.yml up --build
```

- Frontend → build из `./frontend`
- Backend → build из `./backend`
- Traefik → общий
- Переменные API_PREFIX подставляются из `.env`

### 2️⃣ Профиль продакшена (prod)

Использует отдельный Compose файл `docker-compose.yml`:

```bash
docker compose up -d
```

- Frontend → image `${DOCKER_USER}/docker-project-frontend:latest`
- Backend → image `${DOCKER_USER}/docker-project-backend:latest`
- Traefik → общий

---

## ⚙️ Конфигурируемость

Все основные параметры вынесены в `.env` файл:

```env
# API prefix для dev профиля
API_PREFIX=/api

# API prefix для prod профиля
API_PREFIX_PROD=/api/v1

# Имя пользователя DockerHub
DOCKER_USER=yourdockerhubusername
```

Файл `.env.example` добавлен для шаблона.

### Использование переменных:

- Frontend → build-arg `VUE_APP_API_URL`
- Backend → Traefik router rule + middleware prefix

---

## 🔒 Безопасность контейнеров

Контейнеры запускаются с ограничениями:

- `cap_drop: ALL`
- `security_opt: no-new-privileges:true`
- `read_only: true`
- non-root user `1000:1000`
- tmpfs для временных каталогов

---

## ❤️ Healthchecks

Настроены healthcheck-и для всех контейнеров:

- Backend → `/healthcheck`
- Frontend → доступность `index.html`
- Traefik → `/ping`

---

## ⚙️ CI/CD Pipeline

GitHub Actions workflow:

- Сборка образов Backend и Frontend
- Проверка образов с помощью Trivy на уязвимости (`HIGH`, `CRITICAL`)
- Публикация образов в DockerHub
- Запуск `docker-compose.yml` (через фиксированную джобу)

---

## 🏆 Соответствие требованиям

| Требование | Выполнено |
|------------|-----------|
| Использование Compose | ✅ |
| Разделение dev/prod конфигураций | ✅ |
| Конфигурируемость контейнеров | ✅ |
| Безопасность (cap_drop, no-new-privileges, read_only) | ✅ |
| Healthchecks | ✅ |
| CI/CD Pipeline (Trivy + push) | ✅ |
| Документация (README.md) | ✅ |

---

## Запуск проекта

### Dev профиль

```bash
docker compose -f docker-compose.dev.yml up --build
```

### Prod профиль

```bash
docker compose up -d
```

---

## 📜 Дополнительно

- Frontend и Backend используют минимизированные и оптимизированные образы:
  - Backend → `scratch`
  - Frontend → `nginx:1.25-alpine-slim`

- Используются non-root пользователи, ограниченные права, tmpfs.

---

## Контакты

Автор: Спартак Катвицкий

---