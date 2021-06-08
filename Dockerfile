FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/fuegoDeQuasar .
EXPOSE 8081
CMD /out/fuegoDeQuasar

