# Form3


## Pre-Requisites

- install docker-compose [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)
- Environment Variables: 
  - MONGO_URI : is defined in docker-compose so will automatically take it from here if the services are ran from docker, inside the code defaults to localhost if we want to run a mongo server in localhost instead.

## Usage

Start services:

- ```docker-compose up ```

On another terminal get the IP of the form3 service container:
-  ```docker inspect form3_service | grep "IPAddress"``` 
  > IPAddress": "172.19.0.3",

Use services e.g:
curl --request GET --url http://172.19.0.3:5000/payments
            
Note: use localhost instead if the service is not running in the container.

## Desing REST API documentation

https://documenter.getpostman.com/view/235847/form3/RWEiLxxR#8f25a0b5-7c86-bdcf-f776-63a8c34f900a



