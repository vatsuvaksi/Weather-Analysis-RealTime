# Use the official Golang image as a parent image
FROM golang:1.21.0

# Set the working directory to /app
WORKDIR /app

# Copy everything from the current directory to the /app directory in the container
COPY . .

# Build the Go application
RUN go build .

# Expose port 7000 (where your application is running)
EXPOSE 7070

# Run the Go application
CMD ["./real-time-weather-app"]
