# 🎫 Ticketing System — Datacenter

![Go](https://img.shields.io/badge/GO-1.25-00ADD8?style=flat-square&logo=go&logoColor=white) ![Gin](https://img.shields.io/badge/GIN-1.12-00ADD8?style=flat-square&logo=go&logoColor=white) ![Flutter](https://img.shields.io/badge/FLUTTER-3.x-02569B?style=flat-square&logo=flutter&logoColor=white) ![Dart](https://img.shields.io/badge/DART-≥3.0-0175C2?style=flat-square&logo=dart&logoColor=white) ![Firebase](https://img.shields.io/badge/FIREBASE-AUTH-FFCA28?style=flat-square&logo=firebase&logoColor=black) ![PostgreSQL](https://img.shields.io/badge/POSTGRESQL-15-4169E1?style=flat-square&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/DOCKER-READY-2496ED?style=flat-square&logo=docker&logoColor=white)

**Ticketing System Datacenter** adalah platform full-stack untuk manajemen tiket kendala di lokasi datacenter, dibangun untuk **PT Dwi Mitra Ekatama Mandiri**. Aplikasi ini menghubungkan **Engineer** di lapangan yang melaporkan kendala, dengan **Admin** yang melakukan review, approval, maupun rejection — seluruhnya secara real-time melalui mobile app dengan tampilan **dark-mode premium** dan backend API yang aman dengan **Firebase Authentication** serta **Role-Based Access Control (RBAC)**.

> ⚠️ **Repository ini bersifat private.** Seluruh credentials, API keys, dan konfigurasi sensitif dikecualikan dari version control melalui `.gitignore`.

---

## 🏗️ Architecture

```text
ticketing-system-datacenter/
│
├── backend/              # Go REST API Server
│   ├── cmd/api/          # Application entrypoint
│   ├── internal/
│   │   ├── config/       # Environment & app configuration
│   │   ├── controllers/  # HTTP request handlers (User, Ticket)
│   │   ├── handlers/     # Route handler utilities
│   │   ├── middleware/    # Firebase Auth token verification
│   │   ├── models/       # Data structures & domain types
│   │   ├── repository/   # Database access layer (PostgreSQL)
│   │   └── service/      # Business logic layer
│   ├── scripts/          # SQL schema & migration scripts
│   ├── Dockerfile        # Multi-stage Docker build
│   └── vercel.json       # Vercel serverless deployment config
│
├── mobile/               # Flutter Mobile Application
│   ├── lib/
│   │   ├── main.dart     # App entrypoint with Firebase init
│   │   ├── models/       # Dart data models (Ticket)
│   │   ├── screens/      # UI Screens (Login, Home, Create, Detail)
│   │   └── services/     # API service layer (HTTP + Auth)
│   ├── assets/images/    # App logo & branding assets
│   └── android/          # Android platform configuration
│
└── scripts/              # DevOps & utility scripts
```

---

## ⚙️ Core Stack

### Backend

| Layer | Technology | Purpose |
| ------- | ----------- | --------- |
| **Language** | Go 1.25 | High-performance compiled backend |
| **HTTP Framework** | Gin v1.12 | Fast, minimalist web framework |
| **Database** | PostgreSQL (pgx v5) | Relational database with native Go driver |
| **Authentication** | Firebase Admin SDK v4 | Token verification & custom claims (RBAC) |
| **Containerization** | Docker (multi-stage) | Lightweight production images |
| **Deployment** | Vercel / Self-hosted | Serverless or containerized deployment |

### Mobile

| Layer | Technology | Purpose |
| ------- | ----------- | --------- |
| **Framework** | Flutter 3.x | Cross-platform mobile UI framework |
| **Language** | Dart ≥3.0 | Modern, type-safe client language |
| **Auth** | Firebase Auth 5.x | Email/password authentication |
| **State** | StreamBuilder + FutureBuilder | Reactive UI state management |
| **Networking** | http package | RESTful API communication |
| **Media** | image_picker | Camera & gallery access for ticket photos |

### Infrastructure

| Component | Technology | Purpose |
| ----------- | ----------- | --------- |
| **Auth Provider** | Firebase Authentication | Centralized identity & token management |
| **Database** | PostgreSQL | Persistent relational data storage |
| **Realtime Sync** | Firebase Custom Claims | Role-based access synced across client & server |
| **Build** | Docker multi-stage | Optimized ~15MB production container |

---

## 🔑 Key Features

- **🔐 Role-Based Access Control (RBAC)**
  - `admin` — Review all tickets, approve/reject with reason
  - `engineer` — Submit new tickets, view own ticket history

- **🎫 Ticket Lifecycle Management**
  - Create → Pending → Approved / Rejected
  - Full audit trail with timestamps and approver info

- **🔥 Firebase-Powered Authentication**
  - Secure email/password sign-in
  - Server-side token verification on every API call
  - Custom claims for seamless role synchronization

- **📱 Premium Dark-Mode UI**
  - Material Design 3 with custom color scheme (`#6C63FF`)
  - Smooth fade-in animations & gradient backgrounds
  - Responsive card-based ticket list with status indicators

- **🌐 RESTful API (v1)**
  - `GET  /health` — Service health check
  - `GET  /api/v1/user/sync` — Sync user profile & role
  - `GET  /api/v1/tickets` — List all tickets
  - `POST /api/v1/tickets` — Create a new ticket
  - `PATCH /api/v1/tickets/:id` — Approve or reject a ticket

---

## 🔒 Security & Privacy

| Measure | Implementation |
| --------- | --------------- |
| **Credentials** | `.env`, `serviceAccountKey.json` excluded via `.gitignore` |
| **Token Auth** | Firebase ID Token verified server-side on every request |
| **RBAC Enforcement** | Admin-only actions validated in controller layer |
| **CORS** | Configurable cross-origin policy in middleware |
| **No Hardcoded Secrets** | All sensitive values loaded from environment variables |
| **Multi-stage Build** | Source code excluded from production Docker image |

---

## 🚀 Getting Started

### Prerequisites

- **Go** ≥ 1.25
- **Flutter SDK** ≥ 3.0
- **PostgreSQL** ≥ 14
- **Firebase Project** with Authentication enabled
- **Docker** (optional, for containerized deployment)

### Backend Setup

```bash
# 1. Navigate to backend
cd backend

# 2. Copy environment template and configure
cp .env.production .env
# Edit .env with your DB_URL and Firebase credential path

# 3. Place your Firebase service account key
# Download from Firebase Console → Project Settings → Service Accounts
cp /path/to/your/key.json ./serviceAccountKey.json

# 4. Initialize database schema
psql -U your_user -d your_db -f scripts/schema.sql

# 5. Run the server
go run cmd/api/main.go
```

### Mobile Setup

```bash
# 1. Navigate to mobile
cd mobile

# 2. Install dependencies
flutter pub get

# 3. Configure API endpoint
# Edit lib/services/api_service.dart → update _baseUrl

# 4. Run on device/emulator
flutter run
```

### Docker Deployment

```bash
cd backend
docker build -t ticketing-datacenter .
docker run -p 8080:8080 --env-file .env ticketing-datacenter
```

---

## 📂 Environment Variables

| Variable | Description | Required |
| ---------- | ------------- | ---------- |
| `PORT` | Server port (default: `8080`) | ✅ |
| `DB_URL` | PostgreSQL connection string | ✅ |
| `GIN_MODE` | `debug` or `release` | ✅ |
| `FIREBASE_CREDENTIAL_PATH` | Path to Firebase service account JSON | ✅ |

> ⚠️ **Never commit `.env` or `serviceAccountKey.json` to version control.**

---

## 🧪 API Quick Reference

All protected endpoints require `Authorization: Bearer <firebase_id_token>` header.

```text
┌─────────┬─────────────────────────┬────────────────────────────┐
│ Method  │ Endpoint                │ Access                     │
├─────────┼─────────────────────────┼────────────────────────────┤
│ GET     │ /health                 │ Public                     │
│ GET     │ /api/v1/user/sync       │ Authenticated              │
│ GET     │ /api/v1/tickets         │ Authenticated              │
│ POST    │ /api/v1/tickets         │ Authenticated (Engineer)   │
│ PATCH   │ /api/v1/tickets/:id     │ Authenticated (Admin)      │
└─────────┴─────────────────────────┴────────────────────────────┘
```

---

## 🗄️ Database Schema

```sql
users     (id, name, email, role)
sites     (id, name, address)
tickets   (id, user_id, site_id, description, status, photo_url,
           created_at, approved_by, approved_at, rejection_reason)
```

**Relationships:** `tickets.user_id → users.id` · `tickets.site_id → sites.id` · `tickets.approved_by → users.id`

---

## 🤝 Contributing

This is a private, internal project for **PT Dwi Mitra Ekatama Mandiri**. Access is restricted to authorized team members only.

---

*Built with ❤️ using **Go**, **Flutter**, **Firebase** & **PostgreSQL** — © 2026 PT Dwi Mitra Ekatama Mandiri. All Rights Reserved.*
