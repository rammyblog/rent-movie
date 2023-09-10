# Rent movie

## Introduction

This is a simple movie rental application API built in GO. The objective was to try and learn GO concepts. Apart from the api, it contains a worker that periodically send emails to defaulter every day.

It was built with lot of struggles and love ❤️

## Requirements

- [Go](https://golang.org) - v1.20 above
- [Postgres](https://mysql.com) - v14 above

## Installation

> **NOTE**</br>
>
> - Before the fourth step, you should have created a database for this application
> - The supported database is Postgres

- Clone this repo

  ```bash
  git clone https://github.com/rammyblog/rent-movie
  ```

- Change directory to project directory

  ```bash
  cd rent-movie
  ```

- Copy `.env` template

  ```bash
  cp .env.example .env
  ```

- Add correct database credentials to the `.env` file, credentials include:

  - `PORT`: This is the port the application will be served on
  - `DB_HOST`: This is your database host name/IP address
  - `DB_NAME`: This is the name of the database created for the application
  - `DB_USER`: This is your database user
  - `DB_PORT`: This is your database port
  - `DB_PASS`: This is your database password if any, it should be left blank if no password is configured (localhost)
  - `JWT_SECRET` This is for auth.. You can put random string there

- Seed Movies Data

  ```bash
  make seed
  ```

- Run application

  ```bash
  make run
  ```

  This will serve this application on the port you specified in the `.env` file.

## Usage
The API is documented here 👇👇</br>
https://documenter.getpostman.com/view/15213147/2s9YC1WuHX
## Contributors

| Contributor Name                                       | Role      | Tool                                    | Language(s) |
| ------------------------------------------------------ | --------- | --------------------------------------- | ----------- |
| [Babatunde Onasanya](https://twitter.com/simply_rammy) | Developer | [VSCode](https://code.visualstudio.com) | Go          |
