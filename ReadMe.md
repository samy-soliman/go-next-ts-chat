[![Continuous Integration](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CI.yml/badge.svg?branch=main&event=push)](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CI.yml)
[![Continuous Deployment](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CD.yml/badge.svg?branch=main&event=push)](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CD.yml)

# Realtime chat built with Go, Next, and Typescript

![screenshot](/assets/appScreanShot2.JPG)

## In this project, we will build a Realtime web socket chat using Go, Next, Tailwind CSS, Typescript, and PostgreSQL. We then deploy it using different methods:
1. Docker Compose file
2. Automated Delivery to GKE using Github Actions CICD pipeline for building, testing, artifact saving, and deployment of the project. I also configured a public domain using Kubernetes gatewayAPI. 

# Development

In this section, we look at the development perspective of the project, getting to know its architecture and how it works. In simple words, it's a three-tier app: the frontend is written in Next.js, the backend with Golang, and the database with PostgreSQL. It is a chat system where users can create different chat rooms and join them for chatting.

## Hub Architecture

![Initial Hub Architecture](/assets/hub_initial.jpg)

First, we have the hub running on a separate goroutine, which is the central place that manages different channels and contains a map of rooms. The hub has a Register and an Unregister channel to register/unregister clients, and a Broadcast channel that receives a message and broadcasts it out to all the other clients in the same room.

![Client joins room](/assets/join_room.jpg)

A room is initially empty. Only when a client hits the `/ws/joinRoom` endpoint, that will create a new client object in the room and it will be registered through the hub's Register channel.

![Hub Architecture](/assets/hub_architecture.jpg)

Each client has a `writeMessage` and a `readMessage` method. `readMessage` reads the message through the client's websocket connection and sends the message to the Broadcast channel in the hub, which will then broadcast the message out to every client in the same room. The `writeMessage` method in each of those clients will write the message to its websocket connection, which will be handled on the frontend side to display the messages accordingly.

# How To Get the Project Working

## Docker Compose
1. Clone The Repo.

```Shell
    git clone https://github.com/samy-soliman/go-next-ts-chat.git
```

2. I made a **Docker Compose** file to spin up the project, so make sure you have docker installed on your system.

3. Run Docker Compose file

```Shell
    docker compose up
```

4. now to prepare the PostgresSql database with our migrations, we will use **golang-migrate** CLI tool to run our migrations. here is its link in the time of writting this ReadME > [golang-migrate](https://github.com/golang-migrate/migrate).

5. Now Apply the migrations to the db.

```Shell
    cd server
    migrate -path db/migrations -database "postgresql://root:password@localhost:5432/go-chat?sslmode=disable" -verbose up
```

6. You can test the db is working either by entering inside the db container and access the db engine or by using the go api throw a http client like postman.
    - Test by container

    ```Shell
        # Enter to db container
        docker exec -it db psql go-chat
        # enter this to list db tables
        \d
        # make sure you see users table created
    ```

    ![screenshot](/assets/appScreanShot3.JPG)

    - Test by Postman

    ![screenshot](/assets/appScreanShot4.JPG)

7. Now You are Ready to go 

### Quick Run - Docker Compose!
1. Jump in to your browser and type http://localhost:3000 to enter NextJs app.

![screenshot](/assets/appScreanShot5.JPG)

2. Enter some data to register

3. Open up a seconed private browser windows to test the web socket.

4. Create a Room then join it using the two users.

![screenshot](/assets/appScreanShot6.JPG)

5. Start Chatting :D

![screenshot](/assets/appScreanShot2.JPG)

![screenshot](/assets/appScreanShot1.JPG)

## GitHub Actions
1. We deploy our app to GCP so first wr have to make sure this API are enabled.
    - Enable Cloud Domains API
    - Enable Cloud DNS API
    - Enable Compute Engine API
    - Enable Kubernetes Engine API
2. Make a GCP Service account to authenticate our GITHUB ACTIONS, Download json key and put it in a secrect in Production Environment, also make sure to substitute all the secrets with your own.
3. I am going to use a Domain from GCP cloud Domains to use for our project you may choose not to so you to configure your own k8s files for you specific implementation.
4. First Register domain

![screenshot](/assets/CloudDomains.JPG)

5. Lucky for as at this phase we already wrote our docker images in the first phase of deploying with docker compose.

6. all you need is to clone the project and provide your own secrets in the github actions workflows and you are good to go.

7. a point to consider is that we need to create the k8s for gatewayAPI first then take its generated IP and put it in our DNS records for the domain.


5. Lucky for as at this phase we already wrote our docker images in the first phase of deploying with docker compose.

6. all you need is to clone the project and provide your own secrets in the github actions workflows and you are good to go.

7. a point to consider is that we need to create the k8s for gatewayAPI first then take its generated IP and put it in our DNS records for the domain.

```Shell
    # get gatewayAPI IP
    kubectl get gateways
```

8. configure DNS records for domain.

![screenshot](/assets/CloudDNS.JPG)

### Quick Run - Github Actions !
1. CI workflow

![screenshot](/assets/CI.JPG)

2. CD workflow

![screenshot](/assets/CD.JPG)

3. you should wait a little then open the domain and the website should be running.