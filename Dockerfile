FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# TODO: Add dockerignore
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /matchma

EXPOSE 8080
ENTRYPOINT ["/matchma"]
CMD [ "--testing", "--player-count=200" ] # If you want more players for testing, change this number
