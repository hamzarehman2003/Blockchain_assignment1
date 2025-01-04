# # Use the official Golang image as the base
# FROM golang:1.23

# # Set the working directory
# WORKDIR /app

# # Copy the module files
# COPY go.mod go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy the application source code
# COPY . .

# # Build the application
# # RUN go build -o blockchain main.
# # RUN ls -l /app
# # RUN chmod +x /app/blockchain

# # Expose the application port
# EXPOSE 8001 8002 8003 8004

# # Run the application
# # CMD ["./blockchain"]
# CMD ["go" ,"run" ,"main.go"]

FROM golang:1.23.4

WORKDIR /app

COPY . .


RUN go mod tidy

RUN go build -o main .

CMD ["./main"]

