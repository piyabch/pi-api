# PI RESTful API assignment

An experimental REST API written in Golang that uses a MySQL database to store data and is containerized by Docker for simpler development and installation. The authorization is included and implemented using JWT. All API endpoints provide built-in testers to confirm that business requirements are valid. It also uses third-party libraries or popular frameworks for common tasks to speed up development. The project structure was designed for expansion with further development.

<br />

## Prerequisites

These tools must be installed on your local machine in order to run the example.

* [Docker](https://docs.docker.com/get-docker/)
* [Go](https://go.dev/doc/install) (only required if you wish to do automated testing)

<br />

## Installation

1. Clone the repository.
   ```sh
   git clone https://github.com/piyabch/pi-api.git
   ```
2. Go to the project directory.
   ```sh
   cd pi-api
   ```
3. Configure the running ports of Web and MySQL at these files if the defaults aren't available.

    * /config/app-config.env
   ```js
    MYSQL_ADDRESS=mysql:3306
    WEB_ADDRESS=:8080
   ```   
   
    * docker-compose.yml
   ```js
    app:
        ports:
        - "8080:8080"
    mysql:
        environment:
            MYSQL_TCP_PORT: 3306
        ports:
        - "3306:3306"
   ``` 
4. Builds, creates, starts, and attaches to containers for a service.
   ```sh
   docker compose up -d
   ```
5. Test whether the installation is finished.
   ```sh
   curl http://localhost:8080/

   ```

<br />

## Summary of API Endpoints

| URI           | Method | Description                                             |
| ------------- | ------ | ------------------------------------------------------- |
| /auth         | POST   | Authorize to get the token for calling other APIs       |
| /users        | POST   | Create a new user with the given name and email address |
| /users/{id}   | GET    | Retrieve user information by ID                         |
| /users/search | GET    | Search user by name                                     |
| /users        | PUT    | Update user information by ID                           |

<br />

## Authorization
This project also demonstrates the authorization process using JSON Web Tokens (JWT). All API requests are required to set the authorization token in the HTTP header; otherwise, the request will be rejected. Following are some examples of making an authorization request:
### Request authorization
   The email and password that were hard-coded are as below. If you wish to make further development, please change it later.
   * Email: admin@example.com
   * Password: defaultpassword
   ```sh
   curl -i -H "Content-Type: application/json" -X "POST" -d "{\"email\":\"admin@example.com\",\"password\":\"defaultpassword\"}"  http://localhost:8080/auth
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sun, 20 Aug 2023 08:15:36 GMT\
   Content-Length: 191\
   \
   {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW5AZXhhbXBsZS5jb20iLCJleHAiOjE2OTI2MDU3MzYsImlhdCI6MTY5MjUxOTMzNiwiaXNzIjoicGkifQ.JygtNovibat6nfQIbFbF1zfl4-Nm9uxq-TMypATE6VI"}
### The authorized token
If the input email and password are correct, you will get the authorized token as in the above example. It will be valid for 24 hours. Please keep it for use in calling APIs.

<br />

## Usage Examples
### Create a new user
   ```sh
   curl -i -H "Content-Type: application/json" -H "Authorization: YOUR_TOKEN" -X "POST" -d "{\"firstName\":\"John\",\"lastName\":\"Doe\",\"email\":\"john.doe@example.com\"}" http://localhost:8080/users
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sat, 19 Aug 2023 02:51:54 GMT\
   Content-Length: 75\
   \
   {"id":4,"firstname":"John","lastname":"Doe","email":"john.doe@example.com"}

<br />

### Retrieve user information by ID  
   ```sh
   curl -i -H "Authorization: YOUR_TOKEN" http://localhost:8080/users/1
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sun, 20 Aug 2023 04:07:31 GMT\
   Content-Length: 75\
   \
   {"id":1,"firstname":"Weerachai","lastname":"Ruengrangsan","email":"wee.ru@gmail.com"}

<br />

### Update user information by ID
   ```sh
   curl -i -H "Content-Type: application/json" -H "Authorization: YOUR_TOKEN" -X "PUT" -d "{\"id\":4,\"firstName\":\"_John\",\"lastName\":\"_Doe\",\"email\":\"j.doe@example.com\"}" http://localhost:8080/users
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sat, 19 Aug 2023 03:17:27 GMT\
   Content-Length: 74\
   \
   {"id":4,"firstname":"_John","lastname":"_Doe","email":"j.doe@example.com"}

<br />

### Search user by name
   ```sh
   curl -i -H "Authorization: YOUR_TOKEN" http://localhost:8080/users/search?name=wee
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sun, 20 Aug 2023 04:10:11 GMT\
   Content-Length: 168\
   \
   [{"id":1,"firstname":"Weerachai","lastname":"Ruengrangsan","email":"wee.ru@gmail.com"},{"id":2,"firstname":"Paweena","lastname":"Suksawad","email":"paw.suk@gmail.com"}]

<br />

## Automated testing

The included test suit provides support for automated testing of all API endpoints. It is intended to be used in concert with the "go test" command, which automates execution of making requests and validate the results. Here are command to begin testing.

   ```sh
   cd pi-api
   go test
   ```
   Response
   > PASS\
   ok      github.com/piyabch/pi-api 0.259s
