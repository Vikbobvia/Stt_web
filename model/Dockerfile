
# Use the official Golang image as the base
FROM golang:1.23.1

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application listens on (adjust if needed)
EXPOSE 8080

# Define the command to run when the container starts
CMD ["go run main.go"]
