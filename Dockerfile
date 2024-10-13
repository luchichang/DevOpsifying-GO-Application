# stage 1 Importing the Base Image
FROM golang:1.23-alpine as Base

# setting up th workdir
WORKDIR /app

#copying the go.mod dependency file
COPY go.mod .

#installing the Dependency
RUN go mod download

#Copying the source code from host to the container
COPY . . 

# building the image
RUN go build -o main .

###########STAGE 2#############################

# Distroless base image for statically typed web app
FROM gcr.io/distroless/base

#copying the build and static file from 1st Stage
COPY --from=Base /app/main .

COPY --from=Base /app/templates ./templates/

#Exposing the application in port 8080
EXPOSE 8080

#Running the build file in go language
CMD ["./main"]



