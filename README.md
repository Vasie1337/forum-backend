## Project Overview
This project is a backend server for a forum application. It is built using Go and the Gin web framework, with GORM for interacting with the database.

## Features
- **User Authentication**: Login and token generation for users and admins.
- **Key Redemption**: Users can redeem keys.
- **Admin Management**: Admins can manage users and keys.

## How to Run
1. Build the project:
```
go build -o build/server.exe cmd/server/main.go
```
2. Run the server:
```
./build/server.exe
```
The server will start on http://localhost:8080.

## Endpoints
- **User Login**: POST /user/login
- **Admin Login**: POST /admin/login
- **Get All Users**: GET /admin/users
- **Redeem Key**: GET /user/redeem-key

## Dependencies
- Gin
- GORM
- JWT

## License
This project is licensed under the MIT License. See the LICENSE file for details.