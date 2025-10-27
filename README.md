# Minimal Go + Gin microservices (Compose + Swarm)

- service-a: GET / -> {"message":"Hi from service A!"}
- service-b: GET / -> calls service-a and returns combined message

## Practical 3.1 (Compose)

### Build Local Images

```bash
# service-a
docker build -t practika3-service-a:latest services/service-a

# service-b
docker build -t practika3-service-b:latest services/service-b
```

### Run with Compose

```bash
docker compose up -d
```

### Test
```bash
curl http://localhost/a/
curl http://localhost/b/
```

---

## Practical 3.2 (Docker Swarm)

### Deploy with Swarm (one node)

```bash
docker swarm init
docker stack deploy -c stack.yaml app
docker stack services app
```

### Test
```bash
curl http://localhost/a/
curl http://localhost/b/
```

### Optional: Push to Registry

If you want to use remote images:

```bash
# Build and push to GHCR (replace <GITHUB_USER> and <REPO>)
docker build -t ghcr.io/<GITHUB_USER>/<REPO>-service-a:latest services/service-a
docker push ghcr.io/<GITHUB_USER>/<REPO>-service-a:latest

docker build -t ghcr.io/<GITHUB_USER>/<REPO>-service-b:latest services/service-b
docker push ghcr.io/<GITHUB_USER>/<REPO>-service-b:latest

# Then update stack.yaml to use ghcr.io images
```

