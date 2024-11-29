# Start from the official golang image with newer version
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Copy the source code into the container
COPY . .

# Install dependencies (if any)
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]