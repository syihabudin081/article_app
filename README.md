# OrdentTest Application (Article Online)

This application is built with Go Fiber for web handling, GORM for database interactions, JWT for authentication, and Docker for containerization. It provides functionality to manage users, articles, and comments.

## Features

- **User Authentication**: Secure user authentication using JWT.
- **Article Management**: CRUD operations for articles with category and status options.
- **Comment Management**: Users can create, update, and delete comments on articles.
- **Database Relations**: Articles and comments are linked to users with foreign keys for relational data.

## Technologies Used

- **Go Fiber**: Web framework for handling HTTP requests.
- **GORM**: ORM library for Go.
- **PostgreSQL**: Database.
- **JWT**: JSON Web Tokens for secure user authentication.
- **Docker**: Containerization for easy deployment.

## Project Structure

```plaintext
OrdentTest/
├── config/               # Configuration files
├── controllers/          # Controllers for handling HTTP requests
├── models/               # Database models
├── repositories/         # Data access layer
├── services/             # Business logic layer
├── utils/                # Utility functions (e.g., JWT handling)
├── main.go               # Application entry point
├── Dockerfile            # Dockerfile for building the app
└── docker-compose.yml    # Docker Compose file
```

## Getting Started

These instructions will help you set up and run the application on your local machine.

### Prerequisites

- **Go** (v1.20+)
- **Docker** and **Docker Compose**
- **PostgreSQL** (if not using Docker for the database)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/syihabudin081/article_app.git
   cd OrdentTest
   ```

2. **Set Up Environment Variables**:

   Create a `.env` file in the root directory with the following variables:
   ```env
   DATABASE_URL=host=localhost port=5432 user=postgres password=123 dbname=article_db sslmode=disable
   ```

3. **Build and Start the Application with Docker Compose**:

   Run the following command to set up the app and database containers with the necessary configurations:
   ```bash
   docker-compose up --build
   ```
### ERD Diagram
 open folder `erd` and open `ERD.png` or ERD.md file

### App Architecture
 open folder `app architecture` and open `app architecture.png`

## API Documentation

You can access the API documentation at:
- Local: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Online: [SwaggerHub Documentation](https://app.swaggerhub.com/apis/SYIHABUPNYK/article_online_ordent/1.0.0)

--- 

This `README.md` now reflects your use of Go Fiber and provides setup instructions with Docker Compose.