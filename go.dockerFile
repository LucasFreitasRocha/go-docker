FROM golang:1.20.4 as build

WORKDIR  /app

COPY ./src ./

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/server"]