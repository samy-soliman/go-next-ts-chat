##  Deploy The project using Docker Compose.

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
</br>