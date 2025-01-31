# Public API for Retrieving Basic Information

## 📌 Overview
This is a simple **public API** built with **Go (Golang)** that returns:  
- Your registered email address  
- The current date and time in **ISO 8601** format  
- The GitHub repository URL  

The API is publicly accessible and includes **CORS handling** for cross-origin requests.

---

## 🚀 Features
- **Lightweight** and fast  
- **CORS enabled** for cross-origin requests  
- **Dynamic timestamp** in **UTC (ISO 8601)** format  
- **JSON response** for easy integration  

---

## 📡 API Endpoint

**Base URL:**  


### **📥 Request**
- **Method:** `GET`  
- **Headers:**  
  - `Content-Type: application/json`

### **📤 Response (200 OK)**
```json
{
  "email": "uviekugbere.theo@gmail.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/Theo-uviekugbere/hng-internship-go"
}
```
---
## ⚙️ How to Run Locally

### 1️⃣ Clone the Repository

```sh
git clone git@github.com:yourusername/your-repo.git
cd your-repo
```

### 2️⃣ Install Go (if not installed)
Check if Go is installed:

```sh
go version
```

### 3️⃣ Run the API

```sh
go run main.go
```

### 4️⃣ Test the API
Open your browser or use curl:

```sh
curl -X GET http://localhost:8080/