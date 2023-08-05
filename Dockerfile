FROM golang:1.20-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

EXPOSE 4000

CMD ["go", "run", "src/main/main.go"]
