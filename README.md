# hithere

hithere helps visualize and verify your clusters load balancing.

## Usage

0. Initialise your docker swarm and join nodes, docker login to ghcr

1. 
```bash
docker service create --with-registry-auth --name hithere --env host={{.Node.Hostname}} --publish published=8083,target=8083 --replicas 20 ghcr.io/joschahenningsen/hithere/hithere
```

2. 
```bash
$ curl -s 127.0.0.1:8083
> Hello From server1

$ curl -s 127.0.0.1:8083
> Hello From server7

$ curl -s 127.0.0.1:8083
> Hello From server42
```

3. Repeat until it gets boring
