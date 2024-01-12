# ![Realtime chat built with Go, Next, and Typescript](image.png)

![Architecture](/assets/Architecture.png)

In this project we will build a **(Realtime)** chat built with **(Go)**, **(Next)**, **(tailwindcss)** ,  **(Typescript)** and **(PostgresSql)**.

## Hub Architecture

![Initial Hub Architecture](/assets/hub_initial.jpg)

First, we have the hub running on a separate goroutine which is the central place that manages different channels and contains a map of rooms. The hub has a Register and an Unregister channel to register/unregister clients, and a Broadcast channel that receives a message and broadcasts it out to all the other clients in the same room.

![Client joins room](/assets/join_room.jpg)

A room is initially empty. Only when a client hits the `/ws/joinRoom` endpoint, that will create a new client object in the room and it will be registered through the hub's Register channel.

![Hub Architecture](/assets/hub_architecture.jpg)

Each client has a `writeMessage` and a `readMessage` method. `readMessage` reads the message through the client's websocket connection and send the message to the Broadcast channel in the hub, which will then broadcast the message out to every client in the same room. The `writeMessage` method in each of those clients will write the message to its websocket connection, which will be handled on the frontend side to display the messages accordingly.

## How To Get the Project Working
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

![Architecture](/assets/appScreanShot3.JPG)

    - Test by Postman

![Architecture](/assets/appScreanShot3.JPG)

## Quick Run,  YAY!
1. I have steps 1 to 6 covered to i will jump with deploying the app.

![Architecture](/Images/1.PNG)

2. Confirm our infrastructure is created inside gcp console

3. Now i am going to ssh into my private vm.

![Architecture](/Images/2.PNG)

4. Now to clone the mongo and nodejs files, i have them in a diffrent repo so you can use it but do not forget to alter the images tags.

```Shell
    git clone https://github.com/samy-soliman/nodejs-k8s.git
```
Notes:
- Only the **Management VM (private)** will have access to internet through the **NAT**.
- The **GKE cluster (private)** will NOT have access to the internet.
- The **Management VM** will be used to manage the **GKE cluster** and **build/push** images to the **Artifact Registry**.
- All deployed images must be stored in Artifact Registry.