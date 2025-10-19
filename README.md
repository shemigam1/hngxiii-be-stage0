# HNGx Stage 0 Profile Service: A Modern Go Gin API 🚀

## Overview

This project delivers a lightweight and efficient backend service built with **Go** and the **Gin** web framework. It serves dynamic profile information and integrates with an external API to fetch real-time data, showcasing robust API development practices and modern Go concurrency patterns.

## Features

- **`github.com/gin-gonic/gin`**: Provides a high-performance HTTP web framework for building RESTful APIs.
- **`net/http`**: Utilized for making external API calls to `catfact.ninja`, demonstrating external service integration.
- **`encoding/json`**: Handles efficient JSON serialization and deserialization for API requests and responses.
- **`time`**: Manages timestamp generation to provide accurate timing information in API responses.
- **API Health Check**: Includes a basic `/ping` endpoint for quick service status verification.
- **Dynamic Profile Data**: Presents a structured `/me` endpoint that combines static user details with dynamically fetched content.
- **Robust Error Handling**: Implements graceful fallback mechanisms for external API failures, ensuring service stability.

## Getting Started

### Environment Variables

This project does not require any environment variables to run, as all necessary configurations are either hardcoded or dynamically fetched.

## Usage

To get this service up and running locally, follow these simple steps:

1.  **Clone the Repository**:

    ```bash
    git clone https://github.com/yourusername/hngxiii-be-stage0.git
    cd hngxiii-be-stage0
    ```

2.  **Run the Application**:
    Ensure you have Go installed (version 1.25.1 or later). Then, you can run the application directly:
    ```bash
    go run main.go
    ```
    The server will start on port `8080`. You should see a message like: `[GIN-debug] Listening and serving HTTP on :8080`.

### API Endpoints

Once the server is running, you can interact with the API using tools like `curl` or Postman.

#### Health Check

Verify the service is running correctly:

```bash
curl http://localhost:8080/ping
```

**Expected Response:**

```
pong
```

#### Get Profile Information

Retrieve the dynamic profile and a fresh cat fact:

```bash
curl http://localhost:8080/me
```

**Expected Response:**

```json
{
  "status": "success",
  "user": {
    "email": "semiloreomotade@gmail.com",
    "name": "Oluwasemilore Omotade-Michaels",
    "stack": "Go/Gin"
  },
  "timestamp": "2024-07-30T10:00:00.000000000Z",
  "fact": "Cats can jump up to six times their height."
}
```

_(Note: The `timestamp` and `fact` will vary with each request.)_

## Technologies Used

| Technology      | Description                                                              |
| :-------------- | :----------------------------------------------------------------------- |
| Go              | Primary programming language, known for its performance and concurrency. |
| Gin             | High-performance HTTP web framework for Go.                              |
| `net/http`      | Standard Go library for HTTP client functionality.                       |
| `encoding/json` | Standard Go library for JSON encoding and decoding.                      |
| `time`          | Standard Go library for handling time-related operations.                |

## API Documentation

### Base URL

`http://localhost:8080`

### Endpoints

#### GET /ping

**Overview**: A simple health check endpoint to verify the server is operational.

**Request**:
No request body or parameters are required.

**Response**:
A plain text string indicating the service status.

```
pong
```

**Errors**:

- **200 OK**: The service is running and responsive.

#### GET /me

**Overview**: Retrieves the current user's profile details along with a dynamically fetched cat fact from an external API.

**Request**:
No request body or parameters are required.

**Response**:
A JSON object containing the user's status, profile information, a timestamp, and a random cat fact.

```json
{
  "status": "success",
  "user": {
    "email": "semiloreomotade@gmail.com",
    "name": "Oluwasemilore Omotade-Michaels",
    "stack": "Go/Gin"
  },
  "timestamp": "2024-07-30T10:00:00.000000000Z",
  "fact": "Cats can jump up to six times their height."
}
```

**Errors**:

- **200 OK**: The request was successful. In case of an external API failure for the cat fact, a default fact is provided, ensuring a successful response.

## License

This project is open-source and available under the MIT License.

## Author Info

**Oluwasemilore Omotade-Michaels**

- LinkedIn: [Semilore Omotade-Michaels](https://www.linkedin.com/in/semiloreomotade)
- Twitter: [@shemigam1](https://twitter.com/shemigam1)

---

[![Go](https://img.shields.io/badge/Go-1.25.1-00ADD8?logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin%20Framework-v1.11.0-008080?logo=gin&logoColor=white)](https://gin-gonic.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
