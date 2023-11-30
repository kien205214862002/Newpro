# Stage build
# Build image dựa trên base image
FROM golang:1.19-alpine as builder

# Copy toàn bộ files trong project vào folder app trong images
COPY ./ /app/

# Set working directory là folder app
WORKDIR /app/

# Cài đặt các depedencies cho project
RUN go mod download

# Build project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-airbnb .

# Stage runner
FROM alpine

WORKDIR /app/
# copy file thực thi go từ stage trước đó
COPY --from=builder /app/go-airbnb .
# copy configs
COPY config/*.yml ./config/
# copy wait-for
COPY wait-for .
# copy migrations
COPY migrations migrations

CMD [ "./go-airbnb" ]









