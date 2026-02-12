FROM baseimage #TODO add image


WORKDIR /app

COPY *.go .


RUN ["go run main.go"]
