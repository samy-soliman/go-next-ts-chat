[![Continuous Integration](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CI.yml/badge.svg?branch=main&event=push)](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CI.yml)
[![Continuous Deployment](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CD.yml/badge.svg?branch=main&event=push)](https://github.com/samy-soliman/go-next-ts-chat/actions/workflows/CD.yml)

# Realtime chat built with Go, Next, and Typescript

![screenshot](/assets/appScreanShot2.JPG)


## In this project, we will:
1. Develop a Realtime web socket chat using Go, Next, Tailwind CSS, Typescript, and PostgreSQL.
2. Deploy it using Docker Compose.
3. Deploy it using Continuous Delivery to GKE using Github Actions.

# Development

- In this section, we look at the development perspective of the project, getting to know its architecture and how it works.
- In simple, it's a three-tier app:
    - The frontend is written in Next.js, tailwind css.
    - The backend with Golang.
    - The database with PostgreSQL.
- To know more about the development process of this app -> [Development](https://github.com/samy-soliman/go-next-ts-chat/blob/main/ReadME/Development.md)

# Deployment - How To Get the Project Working

## 1- Docker Compose
In This part we:
1. Use docker to containerize our **Nextjs** frontend, **Golang** backend and **Postgresql** database.
2. Make a docker compose file to spin up the project at the click of a button.
3. To see a run in action and To Learn how to get that working >> [Docker Compose](https://github.com/samy-soliman/go-next-ts-chat/blob/main/ReadME/Docker-Compose.md)

## 2- GitHub Actions

<div align="center">
  
![](https://github.com/samy-soliman/go-next-ts-chat/blob/main/assets/GitHubCICD.gif)


</div>

In This part we:
1. Configure a domain and DNS server records for our app.
2. Write Kubernetes manifets to our containerized app.
3. Use of Kubenetes gatwayAPI to serve as a http loadbalancer to our app.
4. Write Two Github Actions Pipelines
    - CI for building, testing, linting and artifact saving.
    - CD for deploying our app to GKE.
5. Implement continuous Delivery methodology, beacause maybe we do not want all commits to go for production.
    - We implemented it using github actions **Environment** review option which gives us two features:
        - a trigger for our production deployment.
        - a review for the changes.
6. To see a run in action and To Learn how to get that working >> [Github Actions CICD](https://github.com/samy-soliman/go-next-ts-chat/blob/main/ReadME/GithubCICD.md)
