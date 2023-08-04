# Assignment

## Requirements

- Go (Golang) version 1.16 or higher
- MongoDB

## Getting Started

1. Clone the repository to your local machine:

* git clone https://github.com/shyamriti/spotbuzz-app.git
* cd golang-crud-api

## To install the dependancy

go mod tidy

## Run the application

go run main.go

## Folder structer
* Controller Folder Contains all the endpoint functions

* Database folder contains the database(PostgreSQL) connection configurations

* Model folder contains the Database Schemas

* Router folder, all the routes are defined in this folder

* Main.go- It is the main execution/entry point of the project

# RestAPI end points

*  ### (POST)localhost:8080/players - This route will Insert the response given by the end user into the database table 

* ### (PUT)localhost:8080/players/:id - This route will update data corresponds to the player id

* ### (GET)localhost:8080/players - all the data available in the table will be displayed by hitting this route

* ### (DELETE)localhost:8080/players/:id - This route will delete the data of a matched player id

* ### (GET)localhost:8080/players/rank/:val - Get the data of a matched player ranked "val"

* ### (GET)localhost:8080/players/random - Get the data of a random player

## To run the application using Docker, use the following commands:

* docker-compose up --build
