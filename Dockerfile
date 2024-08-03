FROM golang:latest AS build
WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /src/server *.go

FROM scratch AS webecho
COPY --from=build /src/server /bin/server
EXPOSE 8080
CMD ["/bin/server"]
