FROM golang:1.25.3 AS build
COPY . ./
RUN go install github.com/a-h/templ/cmd/templ@v0.3.943
RUN templ generate
RUN CGO_ENABLED=1 go build -o "/bin/scnorion-console" .

FROM debian:latest
COPY --from=build /bin/scnorion-console /bin/scnorion-console
COPY ./assets /bin/assets
RUN apt-get update
RUN apt install -y ca-certificates
EXPOSE 1323
EXPOSE 1324
WORKDIR /bin
ENTRYPOINT ["/bin/scnorion-console"]