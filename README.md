# Sample Cars API Overview

The Cars API is designed to manage and retrieve information about cars. It provides capabilities to add new cars, retrieve specific cars by their ID, list all cars, and update existing car records.

## Endpoints:
* GET /cars: List all the cars in the database.
* GET /car?id={id}: Retrieve details of a specific car by its ID.
* POST /car: Add a new car to the database.
* PUT /car: Update details of an existing car.

```mermaid
sequenceDiagram
    participant Client as Client
    participant Server as Server
    Client->>Server: GET /cars
    Server-->>Client: Returns list of all cars
    Client->>Server: GET /car?id={id}
    Server-->>Client: Returns car by ID or error message
    Client->>Server: POST /car (with car details in body)
    Server-->>Client: Responds with success or error message
    Client->>Server: PUT /car (with car details in body)
    Server-->>Client: Responds with update confirmation or error message
```

## Data Model:
```mermaid
classDiagram
    class CarRecord {
        +ID : string
        +Make : string
        +Model : string
        +Category : string
        +Package : string
        +Color : string
        +Year : int
        +Mileage : int
        +Price : int
    }
```

API will not save data past its lifetime.

## Development:

To run:

``` bash
go run ./app
```

To test:

``` bash
go test ./...
```

To run integration tests:

```bash
go run ./app &
go test -tags=integration
```
