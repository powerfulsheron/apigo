## Getting started:
```ssh
$ sudo docker-compose up

or

$ sudo make up
```

Then you can open the React frontend at localhost:3000 and the RESTful GoLang API at localhost:5000

Changing any frontend (VUE) code locally will cause a hot-reload in the browser with updates and changing any backend (GoLang) code locally will also automatically update any changes.

## Connect to Postgres:

```ssh

make connect

or

osbkone@osbkone-XPS-15-7590:~/dev/go/apigo$ docker exec -it d9f bash
root@d9f273587f3e:/# psql -U docker
docker-# \c postgres
```

## Testing:

```ssh
make test

or

docker exec -it 834 bash
root@8344ae56c9b8:/go/src/apigo/back# cd tests
root@8344ae56c9b8:/go/src/apigo/back/tests# go get github.com/bxcodec/faker
root@8344ae56c9b8:/go/src/apigo/back/tests# go test -v
```

## Endpoints

### Auth

#### Create user
```javascript
POST
http://localhost:5000/users
{
    "email":"lolo@gmail.com",
    "pass":"secret",
    "first_name": "lolo",
    "last_name": "canava",
    "birth_date": "0001-01-01T00:00:00Z"
}
```

#### Login and get JWT token
```javascript
POST
http://localhost:5000/login
{
    "email":"lolo@gmail.com",
    "pass":"secret"
}
```

### User

#### Modify user
```javascript
PUT
http://localhost:5000/users/:uuid
{
    "email":"lolo@gmail.com",
    "pass":"secret",
    "first_name": "lolo",
    "last_name": "canava"
}
```

#### Delete user
```javascript
DELETE
http://localhost:5000/users/:uuid
{

}
```

#### Get all votes

```
GET
http://localhost:5000/votes
{

}
```

#### Get vote

```
GET
http://localhost:5000/votes/:uuid
{

}
```

#### Create vote

```
POST
http://localhost:5000/votes
{
    "title":"brexit"
    "description":"Vote du brexit"
}
```

#### Update vote

```
PUT
http://localhost:5000/votes/:uuid
{
    "title":"brexit 2"
    "description":"Vote du brexit 2"
}
```

#### Delete vote

```
DELETE
http://localhost:5000/votes/:uuid
{

}
```

#### Vote

```
POST
http://localhost:5000/votes/:uuid
{

}
```