FROM golang:latest AS build

WORKDIR /app

COPY . ./

RUN go mod vendor

RUN go build -o /kyc cmd/citiesd/main.go


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /kyc /kyc

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/kyc"]