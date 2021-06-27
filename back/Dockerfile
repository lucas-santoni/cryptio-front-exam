FROM golang:alpine AS builder
WORKDIR /back
COPY . .
RUN go build -o back .

FROM golang:alpine AS runner
WORKDIR /app
COPY --from=builder /back/back back
CMD /app/back