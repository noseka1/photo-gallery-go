# photo-gallery-go

## Build

You can build this project using:

```
make build
```

## Database

This component requires access to a PostgreSQL database. You can create it using:

```
psql -c 'CREATE DATABASE gallery'
psql -c "CREATE USER gallery WITH ENCRYPTED PASSWORD 'password'"
psql -c 'GRANT ALL PRIVILEGES ON DATABASE gallery TO gallery'
```
## Run

You can run this application using:

```
make start
```

After the service starts up you can test it using curl.

To create some photos:

```
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Odie","category":"animals"}' 127.0.0.1:8080/photos
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Garfield","category":"animals"}' 127.0.0.1:8080/photos
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Empire state building","category":"buildings"}' 127.0.0.1:8080/photos
```

To retrieve all created photos:

```
curl -v 127.0.0.1:8080/photos
```

To add some likes to the photo with ID 2:

```
curl -v -X POST -H 'Content-Type: application/json' --data '{"id":2,"likes":5}' 127.0.0.1:8080/likes
curl -v -X POST -H 'Content-Type: application/json' --data '{"id":2,"likes":2}' 127.0.0.1:8080/likes

```

To retrieve likes received by all photos:

```
curl -v 127.0.0.1:8080/likes
```

To retrieve all photos from a specific category ordered by the number of likes:

```
curl 127.0.0.1:8080/query?category=animals
```

Build Docker image (multi-stage build):
```
docker build -t photo-gallery-go .
```

Run using Docker Compose
```
docker-compose up
```

Tips:

Run PostgreSQL independent (not use docker-compose):
```
docker run --name gallery-postgres -p 5432:5432 -e POSTGRES_USER=gallery -e POSTGRES_PASSWORD=password -e POSTGRES_DB=gallery -d postgres

````

Connect to PostgreSQL instance with psql:
```
psql postgresql://gallery:password@127.0.0.1:5432/gallery
```