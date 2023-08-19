# PI RESTful API assignment

An experimental REST API written in Golang that uses a MySQL database to store data and is containerized by Docker for simpler development and installation. All API endpoints provide built-in testers to confirm that business requirements are valid. It also uses third-party libraries or popular frameworks for common tasks to speed up development. The project structure was designed for expansion with further development.

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
| /users        | POST   | Create a new user with the given name and email address |
| /users/{id}   | GET    | Retrieve user information by ID                         |
| /users/search | GET    | Search user by name                                     |
| /users        | PUT    | Update user information by ID                           |

<br />

## Usage Examples
### Create a new user
   ```sh
   curl http://localhost:8080/users --include --header "Content-Type: application/json" --request "POST" --data "{\"FirstName\":\"John\",\"LastName\":\"Doe\",\"Email\":\"john.doe@example.com\"}"
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sat, 19 Aug 2023 02:51:54 GMT\
   Content-Length: 75\
   \
   {"ID":4,"FirstName":"John","LastName":"Doe","Email":"john.doe@example.com"}

<br />

### Retrieve user information by ID  
   ```sh
   curl http://localhost:8080/users/4
   ```
   Response
   > {"ID":4,"FirstName":"John","LastName":"Doe","Email":"john.doe@example.com"}

<br />

### Update user information by ID
   ```sh
   curl http://localhost:8080/users --include --header "Content-Type: application/json" --request "PUT" --data "{\"ID\":4,\"FirstName\":\"_John\",\"LastName\":\"_Doe\",\"Email\":\"j.doe@example.com\"}"
   ```
   Response
   > HTTP/1.1 200 OK\
   Content-Type: application/json; charset=utf-8\
   Date: Sat, 19 Aug 2023 03:17:27 GMT\
   Content-Length: 74\
   \
   {"ID":4,"FirstName":"_John","LastName":"_Doe","Email":"j.doe@example.com"}

<br />

### Search user by name
   ```sh
   curl http://localhost:8080/users/search?name=wee
   ```
   Response
   > [{"ID":1,"FirstName":"Weerachai","LastName":"Ruengrangsan","Email":"wee.ru@gmail.com"},{"ID":2,"FirstName":"Paweena","LastName":"Suksawad","Email":"paw.suk@gmail.com"}]

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
