# Use the official Go image as a base image
FROM golang:latest
RUN apt-get update && apt-get install -y iputils-ping

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]