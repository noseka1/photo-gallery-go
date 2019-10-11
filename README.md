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
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Odie","category":"animals"}' localhost:8080/photos
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Garfield","category":"animals"}' localhost:8080/photos
curl -v -X POST -H 'Content-Type: application/json' --data '{"name":"Empire state building","category":"buildings"}' localhost:8080/photos
```

To retrieve all created photos:

```
curl -v localhost:8080/photos
```

To add some likes to the photo with ID 2:

```
curl -v -X POST -H 'Content-Type: application/json' --data '{"id":2,"likes":5}' localhost:8080/likes
curl -v -X POST -H 'Content-Type: application/json' --data '{"id":2,"likes":2}' localhost:8080/likes

```

To retrieve likes received by all photos:

```
curl -v localhost:8080/likes
```

To retrieve all photos from a specific category ordered by the number of likes:

```
curl localhost:8080/query?category=animals
```
