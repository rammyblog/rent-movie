# Rent movie


## Introduction

❤️

## Requirements

* [Go](https://golang.org) -  v1.20 above
* [Postgres](https://mysql.com) - v14 above

## Installation

>**NOTE**</br>
> * Before the fourth step, you should have created a database for this application
> * The supported database is MySQL

* Clone this repo

  ```bash
  git clone https://github.com/rammyblog/rent-movie
  ```

* Change directory to project directory

  ```bash
  cd fragrance
  ```

* Copy `.env` template

  ```bash
  cp .env.example .env
  ```

* Add correct database credentials to the `.env` file, credentials include:
  - `PORT`: This is the port the application will be served on
  - `DB_HOST`: This is your database host name/IP address
  - `DB_NAME`: This is the name of the database created for the application
  - `DB_NAME`: This is your database user
  - `DB_PASS`: This is your database password if any, it should be left blank if no password is configured (localhost)

* Run application

  ```bash
  make run
  ```

  This will serve this application on the port you specified in the `.env` file.

## Usage

## Contributors

|   Contributor Name	| Role  	|  Tool 	| Language(s)  	|
|---	|---	|---	|---	|
|   [Babatunde Onasanya](https://twitter.com/simply_rammy)	|  Developer 	|   [VSCode](https://code.visualstudio.com)	|  Go 	|