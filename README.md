# Automated QR-Code Table Ordering System - Backend API

![Golang](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/fiber-000?style=for-the-badge&logo=go&logoColor=white)
![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)

This is the robust, high-performance Golang backend for the Automated QR-Code Table Ordering System. It uses the Fiber framework for lightning-fast HTTP routing and GORM for database operations against a Supabase PostgreSQL instance.

## 🌟 Key Features
- **Clean Architecture / Domain-Driven Design (DDD)**: Ensures the codebase is highly maintainable, scalable, and easy to test.
- **Blazing Fast**: Built on Go Fiber, one of the fastest web frameworks available.
- **Real-Time Readiness**: Exposes endpoints tailored for a Kitchen Display System (KDS) for real-time order tracking.
- **System Metrics Dashboard**: Built-in Fiber monitor middleware to track latency, memory, and CPU usage. (Access via `/metrics`).
- **Call Waiter Integration**: Endpoint dedicated to allowing customers to ping staff directly from their table.
- **Supabase Integration**: Fully configured for Supabase Connection Pooling (PgBouncer) via GORM.

## 🚀 Setup & Installation

### 1. Prerequisites
- Go 1.22+
- Supabase PostgreSQL Database

### 2. Clone and Install
```bash
git clone https://github.com/rusdifn08/Automated-QR-Code-Table-Ordering-System.git
cd backend
go mod tidy
```

### 3. Environment Variables
Create a `.env` file in the root directory and add your Supabase connection string:
```env
PORT=8080
DATABASE_URL=postgresql://postgres.[YOUR_PROJECT]:[YOUR_PASSWORD]@aws-1-ap-southeast-2.pooler.supabase.com:6543/postgres
```

### 4. Run the Server
```bash
go run cmd/api/main.go
```
*Note: The server will automatically migrate the database schema upon starting.*

## 🛣️ API Endpoints

### Menus
- `GET /api/menus` - Fetch all available menu items

### Orders
- `GET /api/orders` - Fetch all orders (Used by KDS Dashboard)
- `POST /api/orders` - Create a new order with cart items
- `GET /api/orders/:id` - Fetch order status
- `PATCH /api/orders/:id/status` - Update order status (pending, preparing, served, paid)

### Tables
- `GET /api/tables/:number` - Verify table details
- `POST /api/tables/:number/call-waiter` - Flag table as needing assistance
- `POST /api/tables/:number/resolve-assistance` - Resolve assistance request

### System
- `GET /health` - API Health check
- `GET /metrics` - Fiber Monitor Dashboard (Memory, Latency, CPU)
