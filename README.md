## HOW TO RUN

### SETUP
1. Clone this repository.
2. In the `handlers/init.go` file, edit the variables `DbPort`, `DbUsername`, `DbPassword` according to the local Postgres server configuration.
3. From the root directory, run `go run main.go`
4. Make requests to the API endpoints listed below.

### API
- POST : **localhost:3000/create**
  - Create new user
  - JSON request body format: `{ "username": "userone", "email": "userone@gmail.com" }`

- GET : **localhost:3000/list**
  - Get all users in pagination format

- GET : **localhost:3000/list?page=0**
  - Get users by page number
  
- PATCH : **localhost:3000/update/1**
  - Update user info and returns user
  - JSON request body format: `{ "username": "userone" }` or `{ "email": "userone@gmail.com" }`
  
- DELETE : **localhost:3000/delete/1**
  - Delete user
