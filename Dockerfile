# Build Stage
FROM golang:1.21.1-alpine3.18 AS builder 
WORKDIR /app
# First dot means everything from simplebank project directory. Second dot means current working directory inside the image where the files are being copied to. As stated in the previous command WORKDIR /app, the current working directory is /app.
COPY . .
# Build project to single binary file named main
RUN go build -o main main.go
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage - this stage is not required, but shrinks the image size into only containing the binary file and necessary files to run the binary file
FROM alpine:3.18
WORKDIR /app
# stage comes from AS command above 
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate ./migrate
COPY app.env . 
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

# This Expose command is only for documentation for dev reading this file. Doesn't actually do anything 
EXPOSE 8080
# final command that will run the above created binary file
CMD [ "/app/main" ]
# originally used this to run db migrations, but now are running them within go code. Keeping for reference for how to use this. 
ENTRYPOINT [ "/app/start.sh" ]