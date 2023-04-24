# Use an official Golang runtime as a parent image
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any dependencies required by the project
RUN go mod download

# Build the Go binary
ENV RUN_ENV PROD
RUN go build -o main .

# Expose the port on which the server will run
EXPOSE 8080

# Define the command to run the server when the container starts
CMD ["./main"]
