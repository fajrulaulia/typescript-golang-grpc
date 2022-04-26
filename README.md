# Rest API and Service GRPC

communicate between Rest API with Service GRPC using` Typescript` as GRPC Client and `Golang` as GRPC Server

## Image
- GRPC Service: `faawidia/webservice-grpc`
- Rest API: `faawidia/webapi-rest:latest`

## Before spin up
- install on your local  `npm` , `nodejs`,  `go`, `docker`, `makefile` and `protoc`


- (running for production mode)

## Development Mode
run in development if you changes in local machine.
 - ### GRPC Service
    - ensures that the go. mod file matches the source code in the module. 
        ``` bash
        $ go mod tidy
        ```
    - Run in development 
        ```
        $ go run cmd/main.go
 - ### Rest API
    - install all deps
        ``` bash
        $ npm i
        ```
    - set `.env` to add port and WEBSERVICE (just optional)
    - Run in
        ``` bash
        $ npm start
        ```
## Production Mode
see a `Makefile` on two project folder(`webapi` and `webservice`), run `make image-build` to build image, then push to regitry or you can adding in docker-compose, just make it communicate by network, for more detail, you can check on `docker-compose.yaml`


## Example

![Alt Text](https://github.com/fajrulaulia/typescript-golang-grpc/blob/master/screenrecorder/screen-capture.gif?raw=true)




