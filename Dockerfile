# FROM alpine
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# RUN update-ca-certificates

# WORKDIR /app/
# ADD ./app /app/
# # ADD ./zoneinfo.zip /usr/lsocal/go/lib/time/
# ADD ./demo.html /app/
# ENTRYPOINT ["./app"]

# Sử dụng image Golang làm base
# FROM golang:latest

# # Đặt thư mục làm việc
# WORKDIR /app

# # Copy mã nguồn vào thư mục làm việc trong container
# COPY . .

# # Biên dịch ứng dụng Golang
# RUN go build -o app

# # Chạy ứng dụng khi container được khởi động
# CMD ["./app"]


# # Sử dụng image base là Windows
# FROM golang:latest

# # Set working directory trong container
# WORKDIR /app

# # Sao chép các file cần thiết từ local vào container
# COPY . .

# # Thực hiện build ứng dụng (ví dụ với ứng dụng Go)
# # RUN go build -o app.exe

# # Thiết lập entrypoint
# CMD ["app.exe"]

# Định nghĩa lớp cơ sở
FROM golang:1.17 AS build

# Set thư mục làm việc
WORKDIR /app

# Sao chép mã nguồn vào container
COPY . .

# Build ứng dụng
RUN go build -o app

# Sử dụng lớp cơ sở alpine
FROM alpine:latest

# Cài đặt các gói cần thiết
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

# Set thư mục làm việc trong container
WORKDIR /app

# Sao chép file đã build từ lớp build sang lớp alpine
COPY --from=build /app/app .

# Sao chép các file tĩnh hoặc tài nguyên khác nếu cần
# ADD ./demo.html /app/

# Chạy ứng dụng khi container được khởi chạy
ENTRYPOINT ["./app"]
