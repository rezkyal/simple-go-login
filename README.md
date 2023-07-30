# Simple go signup-login

A simple signup + login service using JWT for basic authentication

## Prequisite
1. Go >= 1.15
2. Docker
3. Docker-compose

## Tech stack
This service use:
1. Go
2. Postgresql for the database
3. Redis for the cache
4. JWT for authentication
5. Gin as the HTTP serving framework

## Structure

### App
Contains the logic to implement the singleton and initiate the struct needed

### Handler
Contains the interface logics, such as the input/output data transformation, this project have 2 handlers:
1. Signup, the handler related to signup flow
    1. `Signup` function to handle the `/useraccount/signup` endpoint, to handle the signup flow
2. Login, the handler related to login flow
    1. `Login` function to handle the `/useraccount/login` endpoint, to handle login flow
    2. `IsLoggedIn` function to handle the `/useraccount/data/check` endpoint, to handle the logic to check the flow

### Usecase
Contains business flow logics, this project has 1 usecase:
1. User, the usecase that related to user data/interaction
    1. `RegisterNewUser` function to handle new user registration process
    2. `Login` function to handle login process

### Repo
Contains third parties interaction logics, this projects have 1 repo:
1. User, the repo that related to user data
    1. `SaveNewUser` function to save new user data
    2. `GetUserData` function to get specific data based on their email

### Entity
Contains the entities/structs of the service

### Midllewares
Contains the middleware that will be used on the HTTP endpoint, such as JWT validation checking

### Pkg
Contains the package that warp the resource to help it executed as needed, such as adding JSON marshall/unmarshall in redis before executing the command

### Utils
Contains small functions that help data transformation/validation, such as JWT token validation/creation functions

### Config file
A config file (`files/config.json`) is the config file to control the overall configuration of the service

## How to run
### Mac/Linux
if you use mac/linux, you can use the Makefile command that I have created, on the root folder, run:
1. `make init`, this will download the dependencies needed
2. `make deps-up`, this will run the docker-compose for the postgres and redis container
3. open another terminal because the current terminal will be used for the docker-compose
4. `make run`, this will run the golang service using `github.com/cosmtrek/air`
5. the endpoint will be available at `localhost:8080`


### Windows
run these command on cmd/powershell on the root folder:
1. `docker-compose up`, this will run the docker-compose for the postgres and redis container
2. open another cmd/powershell because the current terminal will be used for the docker-compose
1. `go mod vendor`, this will download the dependencies needed
2. `go install github.com/comstrek/air@latest`, this will download comstrek air that will be used to run the service
3. `air`, this will run the golang service using `github.com/cosmtrek/air`
5. the endpoint will be available at `localhost:8080`

### Other notes
to delete the current database data, you can delete the folder `.dev/dbdata` or using `make delete-db-data` if you are using mac/linux

## Endpoint availables
### Signup endpoint
`/useraccount/signup`, a `POST` endpoint to register new users, the request payload example:
```json
{
    "biography": "testing",
    "date_of_birth": "1990-03-28",
    "email": "rezkyal@mail.com",
    "fullname": "Rezky Alamsyah",
    "location": "Balikpapan",
    "password": "password",
    "phone_number": "123456789",
    "profile_photo": "http://img.profile",
    "sex": 1    
}
```

possible responses:
- success
```json
{
    "error": null,
    "success": true
}
```

- error
```json
{
    "error": "[SaveNewUser] error when create new user, err: ERROR: duplicate key value violates unique constraint \"useraccount_email_key\" (SQLSTATE 23505)"
}
```

### Login endpoint
`/useraccount/login`, a `POST` endpoint to login using the registered email and password, the request payload example:
```json
{
    "email": "rezkyal@mail.com",
    "password": "password"
}
```

possible responses:
- success
```json
{
    "error": null,
    "is_password_correct": true,
    "success": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTA3MTMwNzQsInVzZXJfaWQiOjF9.1CkQliHVLpCyo7UbPSP_JmjCuTMyjWU1lkKATl2bFgA"
}
```

- password incorrect
```json
{
    "error": null,
    "is_password_correct": false,
    "success": true,
    "token": ""
}
```

- error
```json
{
    "error": "[Login] error when GetUserData, err: [GetUserData] error when query the user data, err: record not found"
}
```

### Login check endpoint
`/useraccount/data/check`, a `GET` endpoint to check whether the authorization token is valid or not, no request payload, but need to fill the `Authorization` HTTP header with `Bearer <token>`, where the `<token>` can be obtained from the login endpoint

possible responses:
- success
```json
{
    "success": true
}
```

- unauthorized
Will get `Unauthorized` and HTTP 401 response status

## Unit Test
### Mac/Linux
To run the unit test + coverage report, execute the `make testcover`

### Windows
Run this on cmd/powershell
```
go test -v -coverprofile cover.out ./...
go tool cover -func cover.out
```