# WeatherAPIProject
 
## A brief project overview
In this project, a both GET and POST API is created to get the information regarding the weather in which the city name needs as the input and to get this an external api is used. Moreover the project is deployed in the docker.


## Steps to run the program ##
- Go installed on your machine
- Git installed on your machine

- Clone the repository using command - "git clone https://github.com/DilpreetSinghCoder/WeatherAPIProject.git"
- Open the terminal the to the main folder.
- To run the API server run command - "go run main.go"
- To check the GET request pass the query paramter named `name` in postman and check "http://localhost:8012/city?name=toronto"
- To check the POST request pass the Content-Type: application/json in the postman and check "http://localhost:8012/city"
	Data example: {"name": "Toronto"}

## The port used in the program is 8012

## Commands to build the docker image ##
## How to build and run docker image 
docker build -t dilpreetsinghhardoi/weatherapiapp:v02 
docker run -d -p 8012:8012 dilpreetsinghhardoi/weatherapiapp:v02

## Push docker image
docker push dilpreetsinghhardoi/weatherapiapp:v02


## Commands to run the test cases ##
To run the tests run command - "go test -v"


