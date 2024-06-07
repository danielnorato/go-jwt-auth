# Go JWT Auth

A simple authentication service in Go using the Chi router, GORM for ORM, and PostgreSQL as the database. This service provides endpoints for user registration, login, and JWT-based authentication.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project demonstrates a simple implementation of JWT-based authentication in Go. It uses the Chi router for routing, GORM for interacting with the PostgreSQL database, and provides basic functionalities for user registration and login.

## Features

- User Registration
- User Login
- JWT Authentication
- Secure Password Storage

## Prerequisites

- Go 1.16 or later
- PostgreSQL 12 or later
- Git

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/danielnorato/go-jwt-auth.git
    cd go-jwt-auth
    ```

2. Install the dependencies:

    ```bash
    go mod download
    ```

## Configuration

Create a `.env` file in the root directory and add the following configuration variables:

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
