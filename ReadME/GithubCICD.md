##  Deploy The project using  Continuous Delivery to GKE using Github Actions.

1. We deploy our app to GCP so first we have to make sure this API are enabled.
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

```Shell
    # get gatewayAPI IP
    kubectl get gateways
```

8. configure DNS records for domain.

![screenshot](/assets/CloudDNS.JPG)

### Quick Run - Github Actions !
1. CI workflow

![screenshot](/assets/CI.JPG)

2. Note after our CI has run successfully our CD is triggerd but i have configured the **Environment** Production to require review before accessing so that makes a trigger for us to enable manually for the CD because we maybe does not want to go production with each commit.

![screenshot](/assets/ProductionReview.JPG)

3. Lets approve that review, this will trigger the CD pipeline

![screenshot](/assets/ProductionReview2.JPG)

4. CD workflow

![screenshot](/assets/CD.JPG)

5. you should wait a little then open the domain and the website should be running.