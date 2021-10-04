[read this tutorial first](https://docs.mongodb.com/manual/tutorial/deploy-replica-set-with-keyfile-access-control/#std-label-deploy-repl-set-with-auth)

1. docker compose \
```yml
version: '3.1'
services:
  mongo:
    image: mongo
    ports:
      - 27018:27018
    volumes:
      - ./mongodb/file.key:/auth/file.key
    networks:
      - pawningshopdev
    command: /usr/bin/mongod --replSet rsmongo --keyFile /auth/file.key --bind_ip_all --port 27018

  mongo1:
    image: mongo
    ports:
      - 27019:27019
    volumes:
      - ./mongodb/file.key:/auth/file.key
    networks:
      - pawningshopdev
    command: /usr/bin/mongod --replSet rsmongo --keyFile /auth/file.key --bind_ip_all --port 27019

  mongo2:
    image: mongo
    ports:
      - 27020:27020
    volumes:
      - ./mongodb/file.key:/auth/file.key
    networks:
      - pawningshopdev
    command: /usr/bin/mongod --replSet rsmongo --keyFile /auth/file.key --bind_ip_all --port 27020
```
`sudo docker-compose -f docker-compose-dev.yml up  -d`
2. establish replica set \
`docker-compose exec mongo mongosh --port 27018`
```js
var cfg = {
    "_id": "rsmongo",
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "mongo:27018",
            "priority": 2
        },
        {
            "_id": 1,
            "host": "mongo1:27019",
            "priority": 1
        },
        {
            "_id": 2,
            "host": "mongo2:27020",
            "priority": 0
        },
    ]
};
rs.initiate(cfg, { force: true });
rs.secondaryOk();
db.getMongo().setReadPref('primary');
rs.status();
```
3. Create root user \
wait a few moment so the replica set can pick a node to the primary, that is usually node 1 `mongo` which is running on port 27018, which is also the instance that you are connected using `mongosh`. \
run bellow on on primary node to create root user: \
```js
admin = db.getSiblingDB("admin");
admin.createUser(
	{
		user: "khanh",
		pwd: "handsome", // or cleartext password
		roles: [ "root" ]
	}
);
```

in replica set, you cannot use `MONGO_INITDB_ROOT_USERNAME` or `MONGO_INITDB_ROOT_PASSWORD` to setup password. because instance in replica set not work until the replica set configuration is done.

4. Now you can login into the replica set using connection string bellow:
`mongodb://khanh:handsome@localhost:27018,localhost:27019,localhost:27020/?authSource=admin&replicaSet=rsmongo&readPreference=secondary`

5. to login any db instance in replica set using:
sudo docker-compose exec mongo1 mongosh --port 27019 -u khanh -p handsome --authenticationDatabase admin