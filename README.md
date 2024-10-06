
# Phone Case CRUD API

This is a simple RESTful API built with Go to manage a collection of **Phone Cases**. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on phone cases. 

The API uses the **net/http** built-in Go package along with **gorilla/mux** for routing.

## Project Structure

```
go-phonecase-crud/
├── controllers/
│   └── phonecase_controller.go       # Handles CRUD operations logic
├── models/
│   └── phonecase.go                  # Defines the PhoneCase struct
├── routes/
│   └── routes.go                     # Defines routes for the API
├── go.mod                            # Module dependencies
├── go.sum                            # Checksums for dependencies
└── main.go                           # Entry point of the application
```

## Features

- **Create a new phone case** (POST `/phonecases`)
- **Get all phone cases** (GET `/phonecases`)
- **Get a phone case by ID** (GET `/phonecases/{id}`)
- **Update a phone case by ID** (PUT `/phonecases/{id}`)
- **Delete a phone case by ID** (DELETE `/phonecases/{id}`)

## Prerequisites

- [Go](https://golang.org/dl/) (v1.20 or later)
- [Git](https://git-scm.com/)

## Setup

1. Clone the repository:

```bash
git clone https://github.com/your-username/phone-case-crud.git
cd phone-case-crud
```

2. Initialize the project and install dependencies:

```bash
go mod tidy
```

This will ensure all necessary dependencies (like **gorilla/mux**) are installed.

## Run the Application

To start the API server, run:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Endpoints

### 1. **Create a Phone Case**
   - **Method**: `POST`
   - **Endpoint**: `/phonecases`
   - **Payload**:
     ```json
     {
       "title": "iPhone 13 Case",
       "description": "A sleek case for iPhone 13"
     }
     ```

### 2. **Get All Phone Cases**
   - **Method**: `GET`
   - **Endpoint**: `/phonecases`

### 3. **Get a Phone Case by ID**
   - **Method**: `GET`
   - **Endpoint**: `/phonecases/{id}`

### 4. **Update a Phone Case**
   - **Method**: `PUT`
   - **Endpoint**: `/phonecases/{id}`
   - **Payload**:
     ```json
     {
       "title": "Updated Case",
       "description": "Updated description",
       "status": "sold out"
     }
     ```

### 5. **Delete a Phone Case**
   - **Method**: `DELETE`
   - **Endpoint**: `/phonecases/{id}`

## Testing the API

You can use tools like [Postman](https://www.postman.com/) or `curl` to test the API endpoints.

- **Create a Phone Case**:

```bash
curl -X POST http://localhost:8080/phonecases -d '{"title": "iPhone 13 Case", "description": "A sleek case for iPhone 13"}' -H "Content-Type: application/json"
```

- **Get All Phone Cases**:

```bash
curl http://localhost:8080/phonecases
```

- **Update a Phone Case**:

```bash
curl -X PUT http://localhost:8080/phonecases/1 -d '{"title": "Updated Case", "description": "Updated description", "status": "sold out"}' -H "Content-Type: application/json"
```

- **Delete a Phone Case**:

```bash
curl -X DELETE http://localhost:8080/phonecases/1
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gorilla Mux](https://github.com/gorilla/mux) for providing the router.
