# Use the official Go image from the Docker Hub
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Build the Go application
RUN go build -o orders main.go

# Expose a port for the Go application to listen on (adjust as needed)
EXPOSE 8080

# Define the command to run the Go application
CMD ["./orders"]
