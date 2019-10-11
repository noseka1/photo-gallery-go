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
