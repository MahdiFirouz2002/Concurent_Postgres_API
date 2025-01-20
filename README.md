# Concurrent PostgreSQL API

This project is a concurrent Go application that reads user data from a JSON file and inserts it into a PostgreSQL database. It uses 10 Goroutines for concurrent insertion and employs a `sync.WaitGroup` and channels to manage concurrency. The project also uses the Gin web framework to serve an API that allows fetching user details by ID.

## Features

- **Concurrent Data Insertion**: Inserts 10,000 user records into PostgreSQL concurrently using 10 Goroutines.
- **PostgreSQL**: All user data is stored in a PostgreSQL database.
- **Gin Framework**: Provides an API to query user information by user ID.
- **WaitGroup & Channels**: Utilizes `sync.WaitGroup` and channels for managing concurrent Goroutines and ensuring safe data handling.
  
## Prerequisites

Before running the project, make sure you have the following installed:

- **Go (v1.18 or higher)**
- **PostgreSQL** with a running database
- **Gin Framework** (The project uses `github.com/gin-gonic/gin` for API routing)
- **JSON File**: A JSON file containing the user data to be inserted into the database.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/MahdiFirouz2002/Concurent_Postgres_API.git
   cd concurrent-postgres-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your PostgreSQL database:

   - Create a PostgreSQL database and table for users and addresses.

     Example SQL:

     ```sql
        CREATE TABLE users(
            id VARCHAR(200) PRIMARY KEY,
            name VARCHAR(200),
            email VARCHAR(200),
            phone_number VARCHAR(200)
        );

        CREATE TABLE addresses(
            id SERIAL PRIMARY KEY,
            street VARCHAR(200),
            city VARCHAR(200),
            state VARCHAR(200),
            zip_code VARCHAR(200),
            country VARCHAR(200),
            userId VARCHAR(200) REFERENCES users ON DELETE CASCADE
        );
     ```

4. Update the database connection string in the `ConnectToDB` function with your PostgreSQL credentials and database details.

   Example:

   ```go
   const (
       DbUser     = "postgres"
       DbPassword = "your_password"
       DbName     = "your_database"
       DbHost     = "localhost"
       DbSSLMode  = "disable"
   )
   ```

5. Create a `users.json` file in the project directory with 10,000 user records. Example structure:

   ```json
   [
       {
           "id": "7f6d09fa-ccc2-498e-bd71-77a45c12853a",
           "name": "John Doe",
           "email": "john.doe@example.com",
           "phone_number": "+1234567890"
       },
       ...
   ]
   ```

## Usage

1. **Start the API server**:

   To start the server, run:

   ```bash
   go run main.go
   ```

   This will start the Gin server at `http://localhost:8080`.

2. **Insert Data into PostgreSQL**:

   The program will begin inserting the 10,000 user records from `users.json` into the PostgreSQL database using 10 concurrent Goroutines. The insertion process will complete once all records are successfully added.

3. **API to Fetch User by ID**:

   Once the data is inserted, you can query a user by their ID using the following API endpoint:

   - **GET** `/user/:id`
   
     Example:

     ```bash
     curl http://localhost:8080/user/7f6d09fa-ccc2-498e-bd71-77a45c12853a
     ```

   This will return the user data as a JSON object:

   ```json
    {
	"data": {
		"id": "645df623-c037-4935-9c9e-0b8e2ee78911",
		"name": "Federico Ryan",
		"email": "rustyraynor@mcglynn.name",
		"phone_number": "6434290414",
		"addresses": [
			{
				"street": "8077 Placeberg",
				"city": "Albuquerque",
				"state": "North Dakota",
				"zip_code": "46656",
				"country": "Iran (Islamic Republic of)"
			},
			{
				"street": "2646 Prairiemouth",
				"city": "Baton Rouge",
				"state": "Texas",
				"zip_code": "68113",
				"country": "Japan"
			},
			{
				"street": "47677 Port Mountainschester",
				"city": "San Jose",
				"state": "Missouri",
				"zip_code": "51903",
				"country": "Palestine, State of"
			},
			{
				"street": "5804 Wellsport",
				"city": "Mesa",
				"state": "New Jersey",
				"zip_code": "93089",
				"country": "Somalia"
			}
		]
	}
}
   ```