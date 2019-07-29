# Introduction

Scope of this software is just to play around with Kubernetes (Minikube). </br>
This software is exposing a set of rest-api to manage a collection of ToDo operations and is accessing postgres to store and read Todos. </br>
We will see how this app can be deployed locally, using just Docker containers to link one container to another or using kubernetes, minikube in this case. </br>
In the past I also used this app using mysql: </br>
https://github.com/DanielePalaia/web-service-kubernetes</br>

## Datastore and rest api

The todos operations are saved in a mysql datastore defined in datastore.sql

```
CREATE TABLE ToDo (
	    ID serial,
	    Topic varchar(255),
	    Completed int,
	    Due varchar(255) DEFAULT '',
	    PRIMARY KEY (ID)
);
```

The software exposes these rest api (GET, POST, PUT, DELETE) which can be tested with curl.</br>
Use GET to get all todo items inside the collection </br>
Use PUT to create a new item for the Todo Collection </br>
Use POST for update and DELETE for deletion </br>

```
curl http://localhost:8080/todos
curl -H "Content-Type: application/json" -d '{"Topic":"New TodoElem", "Completed":0}' -X PUT http://localhost:8080/todos
curl http://localhost:8080/todos/1
curl -H "Content-Type: application/json" -d '{"Id":0,"name":"New TodoElem Updated"}' -X POST http://localhost:8080/todos
curl -X DELETE http://localhost/todos/1
curl -X DELETE http://localhost/todos
```

A Swagger documentation that allows you to test this interface is also provided:

![Screenshot](./pics/pic2.png)
![Screenshot](./pics/pic3.png)

## Running the app locally
In conf file please specify the right database connection information,
then you can simply run the binary provided in ./bin and run the binary </br>

**./kubernetes-postgres**</br></br>
Then go to the swagger interface</br>
http://localhost:8080/docs/index.html</br>
![Screenshot](./pics/pic4.png) </br>
Initially you will receive an empty list of todos. Fill the todo with the other rest api with curl or the swagger doc.
</br>

## Running on kubernetes/minikube

### 1. Install minikube
On mac is enough to:</br>
**brew cask install virtualbox**</br>
**brew cask install kubectl**</br>
**brew cask install docker** </br>

### 2. Start minikube and dashboard
Start minikube: </br>
**minikube start** </br>
Run dashboard </br>
**minikube dashboard**</br>

![Screenshot](./pics/minikube.png) </br>

### 3. Install Postgresql on minikube 
Follow this guideline to create a volume a pod and a service: </br>
https://severalnines.com/blog/using-kubernetes-deploy-postgresql
