# 💳 Transaction Processing Engine

A backend service that simulates a **payment switch authorization system** for processing card transactions.

---

## 🚀 Features

* 🔐 Secure PIN validation using SHA-256 hashing
* 💳 Card management with in-memory storage
* 💰 Transaction processing (Withdraw / Top-up)
* 📜 Transaction history logging
* ⚡ Fast in-memory operations using Go maps

---

## 🏗️ Architecture

```
Client (Postman / Curl)
        │
        ▼
   REST API (Gin)
        │
        ▼
   Handlers Layer
        │
        ▼
   Business Logic
        │
        ▼
   In-Memory Storage (HashMap)
```

---

## 📂 Project Structure

```
.
├── main.go
├── handlers/
├── models/
├── storage/
├── utils/
└── README.md
```

---

## ⚙️ Setup & Run

### 1. Clone Repo

```bash
git clone https://github.com/YOUR_USERNAME/transaction-processing-engine.git
cd transaction-processing-engine
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run Server

```bash
go run main.go
```

Server runs on:

```
http://localhost:8080
```

---

## 🔌 API Endpoints

### 1️⃣ Process Transaction

**POST** `/api/transaction`

#### Request

```json
{
  "cardNumber": "4123456789012345",
  "pin": "1234",
  "type": "withdraw",
  "amount": 200
}
```

#### Curl Example

```bash
curl -X POST http://localhost:8080/api/transaction \
-H "Content-Type: application/json" \
-d '{
  "cardNumber": "4123456789012345",
  "pin": "1234",
  "type": "withdraw",
  "amount": 200
}'
```

---

### 2️⃣ Get Balance

**GET** `/api/card/balance/{cardNumber}`

#### Curl Example

```bash
curl http://localhost:8080/api/card/balance/4123456789012345
```

---

### 3️⃣ Transaction History

**GET** `/api/card/transactions/{cardNumber}`

#### Curl Example

```bash
curl http://localhost:8080/api/card/transactions/4123456789012345
```

---

## 📬 Postman Testing

### Step 1:

Open Postman → Create New Request

### Step 2:

Set method to **POST**

### Step 3:

URL:

```
http://localhost:8080/api/transaction
```

### Step 4:

Go to **Body → raw → JSON**

Paste:

```json
{
  "cardNumber": "4123456789012345",
  "pin": "1234",
  "type": "withdraw",
  "amount": 200
}
```

---

## 🔐 Security

* PIN stored using **SHA-256 hashing**
* No plaintext PIN stored
* No PIN logging

---

## 📊 Sample Responses

### ✅ Success

```json
{
  "status": "SUCCESS",
  "respCode": "00",
  "balance": 800
}
```

### ❌ Invalid Card

```json
{
  "status": "FAILED",
  "respCode": "05",
  "message": "Invalid card"
}
```

### ❌ Invalid PIN

```json
{
  "status": "FAILED",
  "respCode": "06",
  "message": "Invalid PIN"
}
```

---

## 🛠️ Tech Stack

* Go (Golang)
* Gin Framework
* In-Memory Storage (Maps)

---

## 🎯 Future Improvements

* Database integration (PostgreSQL / MongoDB)
* JWT Authentication
* Rate limiting
* Docker support

---

## 👨‍💻 Author

**Syed Mohammed Sylani**
