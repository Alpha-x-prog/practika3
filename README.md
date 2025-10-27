# Minimal Go + Gin microservices (Compose)

- service-a: GET / -> {"message":"Hi from service A!"}
- service-b: GET / -> calls service-a and returns combined message

## Run
docker compose up -d --build

## Test
curl http://localhost:5001/
curl http://localhost:5002/

