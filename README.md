# Orders Application

The **Orders Application** is a Go-based web service for managing and processing orders. This README provides instructions for running the application locally, running tests, and building a Docker container for the application.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your machine.
- Git installed for version control.
- Docker installed (if you plan to use Docker).

## Getting Started

To run the **Orders Application** locally, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/orders-app.git


2. Change to the project directory:

   ```bash
   cd orders-app

3. Install project dependencies:

    ```bash
    go get -d ./...

4. Build and run the application:

    ```bash
    go run main.go

   The application should now be running locally at http://localhost:8080. You can access it using your web browser or API testing tool

5. Running Tests

   To run tests for the Orders Application, execute the following command from the project root directory

    ```bash
    go test ./...


6. Dockerization 

    You can also run the Orders Application in a Docker container. Follow these steps to build and run a Docker container for the application.

    Ensure you have Docker installed on your machine.

    In the project directory, create a Docker image using the provided Dockerfile:

     ```bash
     docker build -t orders-app.

   This command builds a Docker image named "orders-app" based on the Dockerfile in the project directory

   After building the Docker image, you can run the Orders Application in a Docker container


    ```bash
    docker run -p 8080:8080 orders-app



This command maps port 8080 from the Docker container to port 8080 on your host machine, assuming the application listens on port 8080. Adjust the port mapping if needed.

The application should now be accessible at http://localhost:8080, just like when running it locally.