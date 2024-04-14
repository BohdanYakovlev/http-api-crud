# HTTP API for User Management
This Go code provides an HTTP API for managing users stored in a PostgreSQL database. It allows creating, reading, updating, and deleting user records via HTTP requests.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd http-api-crud
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```
5. Make sure you have a PostgreSQL database set up. You can use the provided db.sql file to create the required schema and tables.

6. Update the dbOpenConfig constant in the main.go file with your PostgreSQL connection details.

## Endpoints
  - GET /users: Retrieves a list of all users.
  - GET /users/{id}: Retrieves a user by ID.
  - POST /users/{username}/{phone}: Creates a new user with the specified username and phone number.
  - DELETE /users/{id}: Deletes a user by ID.
  - PUT /users/{id}/{username}/{phone}: Updates a user's information by ID.

## Conclusion

This HTTP API provides basic CRUD functionality for managing users in a PostgreSQL database. Feel free to reach out if you have any questions or encounter any issues while running the code.

Feel free to reach out if you have any questions or encounter any issues while running the code.
