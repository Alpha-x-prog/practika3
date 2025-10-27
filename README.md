# Minimal Go + Gin microservices (Compose)

- service-a: GET / -> {"message":"Hi from service A!"}
- service-b: GET / -> calls service-a and returns combined message

## Run
docker compose up -d --build

## Test
curl http://localhost:5001/
curl http://localhost:5002/

---

## Practical 3.2 (NGINX + Docker Swarm)

This setup adds NGINX as a reverse proxy routing `/a/` to service-a and `/b/` to service-b.

### Build & Push Images

Build and push images to GitHub Container Registry (replace `<GITHUB_USER>` and `<REPO>`):

```bash
cd services/service-a
docker build -t ghcr.io/<GITHUB_USER>/<REPO>-service-a:latest .
docker push ghcr.io/<GITHUB_USER>/<REPO>-service-a:latest

cd ../service-b
docker build -t ghcr.io/<GITHUB_USER>/<REPO>-service-b:latest .
docker push ghcr.io/<GITHUB_USER>/<REPO>-service-b:latest
```

### Deploy with Docker Swarm

```bash
docker swarm init
docker network create --driver overlay app-net
docker stack deploy -c stack.yaml app
docker stack services app
curl http://<NODE_IP>/a/healthz
curl http://<NODE_IP>/b/
```

