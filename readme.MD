# GO-Docker

Golang RESTAPI implementation in docker

## Step to run on docker

### build docker

    sudo docker build -t go-docker .

### run docker

    sudo docker run -p 8080:8081 -it go-docker

### Running our Container In the Background

    sudo docker run -p 8080:8081 -d go-docker

## API ENDPOINTS

    http://localhost:8080/v1/health-check

    http://localhost:8080/v1/homepage