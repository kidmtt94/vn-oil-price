# Build Stage
# First pull Golang image
FROM golang:1.19-alpine as build-env
 
# Set environment variable
ENV APP_NAME vn-oil-price
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run Stage
FROM alpine:3.14
 
# Set environment variable
ENV APP_NAME vn-oil-price
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
RUN apk --no-cache add tzdata
# Expose application port
EXPOSE 80
 
# Start app
CMD ./$APP_NAME