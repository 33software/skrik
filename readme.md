# Skrik



Skrik is a platform for group chats, made for educational purposes and built with Go.



## Features


- **Authentication**: User registration and login with JWT-based authentication.

- **User Management**: Fetch user profiles and delete users.

- **Chat**: Create chat rooms, save and retrieve messages, and handle real-time messaging via WebSocket.



## Requirements



- Go 1.20 or higher

- Docker and Docker Compose

- PostgreSQL



## Running the Project



### Local Setup

1. Clone the repository:

    ```bash

    git clone <repository-url>

    cd skrik

    ```

2. Configure environment variables in `dev.env` file.

3. Run the project:

    ```bash

    go run cmd/main.go

    ```



### Using Docker

1. Build and start the services:

    ```bash

    docker-compose up --build

    ```

2. The app will be available at `http://localhost:8080`.



## TODO



- [ ] Add Swagger documentation.

- [ ] Implement refresh tokens for JWT authentication.

- [ ] Implement signaling server for voice calling.

## API Routes


### Authentication

- `POST /login`: User login.

- `POST /register`: User registration.



### User Management

- `GET /me`: Fetch current user's profile.

- `DELETE /deleteuser`: Delete a user.



### Chat

- `GET /ws`: Connect to a chat room via WebSocket.

