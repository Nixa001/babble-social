# Social Network Project

A Facebook-like social network application built with a modern JS framework and Go backend, featuring real-time interactions through WebSockets.

## Features

- User Registration and Authentication
- User Profiles (Public/Private)
- Posts and Comments
- Followers System
- Groups and Events
- Real-Time Private Messaging
- Notifications
- Docker Containerization

## Tech Stack

- Frontend: [Next.js](https://nextjs.org/) [React](https://fr.react.dev/) [React Query](https://react-query.tanstack.com/) [TailwindCSS](https://tailwindcss.com/)
- Backend: [Go](https://go.dev/doc/)
- Database: [SQLite](https://www.sqlite.org/docs.html)
- Real-Time Communication: [Gorilla (Go)](https://pkg.go.dev/github.com/gorilla/websocket)
- Containerization: [Docker](https://www.docker.com/)

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/social-network.git
   cd social-network
   ```

2. Build and run the Docker containers:
   ```
   docker-compose up --build
   ```

3. Open your browser and navigate to `http://localhost:3000` (or the port specified for your frontend)

## Project Structure

```
.
├── frontend/
│   ├── Dockerfile
│   └── ... (frontend code)
├── backend/
│   ├── Dockerfile
│   ├── pkg/
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   └── sqlite/
│   │   │   │       ├── 000001_create_users_table.up.sql
│   │   │   │       ├── 000001_create_users_table.down.sql
│   │   │   │       └── ... (other migration files)
│   │   │   └── sqlite/
│   │   │       └── sqlite.go
│   │   └── ... (other packages)
│   └── server.go
└── docker-compose.yml
```

## Features in Detail

### User Management

- Registration with fields: Email, Password, First Name, Last Name, Date of Birth, Avatar (optional), Nickname (optional), About Me (optional)
- Login and session management using cookies
- Public and private profile options

### Posts and Comments

- Create posts with privacy settings (public, private, almost private)
- Comment on posts
- Include images or GIFs in posts and comments

### Followers System

- Send and accept/decline follow requests
- Automatic following for public profiles

### Groups

- Create and join groups
- Invite users to groups
- Request to join groups
- Create and respond to events within groups

### Private Messaging

- Real-time chat functionality
- Send messages to followers or users you're following
- Group chat rooms
- Emoji support

### Notifications

- Follow requests
- Group invitations
- Group join requests
- New events in groups

## Database Migrations

We use SQL migrations to manage the database schema. Migrations are located in `backend/pkg/db/migrations/sqlite/`.

To apply migrations, the application uses the `golang-migrate` package.

## Security

- Password hashing using bcrypt
- Session management with cookies

## Docker

The project uses two Docker containers:

1. Frontend container
2. Backend container

These containers are defined in their respective Dockerfiles and orchestrated using Docker Compose.


## Acknowledgments

- Gorilla WebSocket library
- SQLite database
- golang-migrate for database migrations
- bcrypt for password hashing
