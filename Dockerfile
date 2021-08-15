FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app

COPY ./ ./
RUN go mod download

RUN export GO111MODULE=on

RUN go build -o /planningpoker

CMD [ "/planningpoker" ]
