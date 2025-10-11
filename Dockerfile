FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /linky . 


FROM scratch


COPY --from=builder /linky /linky


COPY web /web


ENTRYPOINT ["/linky"]