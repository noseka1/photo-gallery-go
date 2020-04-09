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

## Testing the Application

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
## Deploying to OpenShift

Create a new project if it doesn't exist:

```
oc new-project photo-gallery-go
```

Deploy a PostgreSQL database:

```
oc new-app \
--template postgresql-persistent \
--param DATABASE_SERVICE_NAME=postgresql \
--param POSTGRESQL_USER=gallery \
--param POSTGRESQL_PASSWORD=password \
--param POSTGRESQL_DATABASE=gallery
```

Define a binary build (this will reuse the Go artifacts that were you built at the beginning):

```
oc new-build \
--name photo-gallery \
--binary \
--strategy docker
```

Start the binary build:

```
oc start-build \
photo-gallery \
--from-dir . \
--follow
```

Deploy the application:

```
oc new-app \
--image-stream photo-gallery \
--name photo-gallery \
--env GALLERY_DB_HOST=postgresql
```

Expose the application to the outside world:

```
oc expose svc photo-gallery
```

## OpenShift Pipelines

Install the OpenShift Client Task from the TektonCD catalog:

```
oc create -f https://raw.githubusercontent.com/tektoncd/catalog/5d22dcb133d83b5cd94aee64084c329d39e15239/openshift-client/openshift-client-task.yaml
```

Create the pipeline Kubernetes objects:

```
oc create -f pipelines
```

Start the pipeline execution to build an application image from source:

```
tkn pipeline start photo-gallery-pipeline \
--resource git=photo-gallery-git \
--resource image=photo-gallery-image \
--serviceaccount pipeline
```

## Deploying as a Knative Function

```
oc apply -f deploy/knative/service.yml
```
