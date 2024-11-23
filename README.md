## Về dự án

## Acknowledgments
This project was inspired by the structure and ideas from [Clean-Architecture](https://github.com/amitshekhariitbhu/go-backend-clean-architecture).
All code in this project was written from scratch and does not directly copy any part of the original source.

## About
This project is learn how to build Clean Architecture, ERP system and use Services Layer Design (Software Architecture) isolated business/application logic for scalable and maintainable

## Run Programming
How to run this project?
We can run this Go Backend Clean Architecture project with or without Docker. Here, I am providing both ways to run this project.

#### Run without Docker
- Create a file .env similar to .env.example at the root directory with your configuration.
- Install go if not installed on your machine.
- Install MongoDB if not installed on your machine.
- Important: Change the DB_HOST to localhost (DB_HOST=localhost) in .env configuration file. DB_HOST=mongodb is needed only when you run with Docker.
- Run go run cmd/main.go.
- Access API using http://localhost:8080
#### Run with Docker
- Create a file .env similar to .env.example at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run docker-compose up -d.
- Access API using http://localhost:8080
#### How to run the test?
#### Run all tests
    go test ./...
#### How to generate the mock code?
- In this project, to test, we need to generate mock code for the use-case, repository, and database.

#### Generate mock code for the usecase and repository
    mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

#### Generate mock code for the database
    mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
- Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the corresponding command to regenerate the mock code for testing.
