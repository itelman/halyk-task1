# Proxy Server

## Routers

### Health Check
- **Endpoint:** `GET /health
  - **Response:** `OK`

### Make Request
- **Endpoint:** `POST /`
  - **Body:**
    ```json
    {
      "method": "GET",
      "url": "http://google.com",
      "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",

      }
    }
    ```

### Get Response
  - **Response:**
    ```json
    {
      "id": "response_id",
      "status": "HTTP-статус ответа стороннего сервиса",
      "headers": { "массив заголовков из ответа стороннего сервиса" },
      "length": "длина содержимого ответа"
    }
    ```

## Models Structure

```sql
Request {
  method  string,
  url     string,
  headers map[string][]string
}

Response {
  id      string,
  status  string,
  headers map[string][]string,
  length  int
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/itelman/hl-task1
   cd hl-task1
   ```
2. **Run the program:**
   ```bash
   go run main.go
   ```
3. **Check the server:**
   Open your browser and go to http://localhost:8080 to ensure the server is running properly.
   
**LINK: https://proxy-server-y7iz.onrender.com/**
