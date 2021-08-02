# ImenAria interview Golang task

## How to Run?
This project has 2+1 microservices, so you should run all of them in parallel.
1. User microservice
2. Bank Card microservice
3. Gateway microservice

**Hint: You can use shortcuts in Makefile to run project correctly.**

**Notice: You should create a .env file in root directory to configure the project(use .env.template)**

## How to Test?
You can use PostMan and get project collection with link below.
```
https://www.getpostman.com/collections/e95dc12db0711705b27c
```

## What more?
In this project, you will see a gateway microservice which is the restful part of the project.
This part will talk with two other microservice with grpc protocol and respond to client in json format.

Uploading avatar works separately, first of all you should use upload avatar api to upload the avatar then pass the avatar path that you'd gotten to create user api with the name and gender and birth year.

