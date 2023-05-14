# Docker network

Simple run 2 container will be end up with different IP each run
Create a network and put it together will be easy to manage

#### Inspect container
> ```docker container inspect <container name>```

#### Inspect network
> ```docker network inspect <network name>```


# Step to create and add container to network

1. Create a network
> ```docker network connect <network name>```

2. Connect container to network
> ```docker network connect <network name> <container 1 name>

3. Run container in network
> ```docker run --network <network name> <container 2 name>

Now that the container 2 can connect to container 1 by using container 1 name cause the ip address has been added to domain name automatically by docker

```aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 154335758310.dkr.ecr.us-east-1.amazonaws.com```