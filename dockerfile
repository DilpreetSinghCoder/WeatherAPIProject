# Setting from a Golang base image with the specific version
FROM golang:1.22 


# Set the Current Working Directory inside the container
WORKDIR /weatherAPI


#Coping all the data to the main app directory 
COPY . .

#Now download the dependencies of the project
RUN go mod download

#This will build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weatherAPI .

#This will expose the port to listen at the runtime
EXPOSE 8012

CMD [ "./weatherAPI" ]