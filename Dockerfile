FROM golang:alpine3.18 AS builder

RUN apk --no-cache add ca-certificates git
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY cli ./cli
RUN cd cli \
    && CGO_ENABLED=0 go build -o musicbinary

FROM alpine:latest AS final
RUN apk --no-cache upgrade
RUN mkdir /app \
    && adduser -D music --shell /usr/sbin/nologin \
    && chown -R music:music /app
WORKDIR /app
COPY --from=builder --chown=music:music /build/cli/musicbinary .
USER music

CMD ["/app/musicbinary"]