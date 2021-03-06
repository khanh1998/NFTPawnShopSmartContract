docker network create pawningshop\
docker volume create ganache_volume\
docker run -dp 80:80 --name client --network pawningshop client\
docker run -dp 4000:4000 --name api --network pawningshop api\
docker run -dp 7789:7789 --name notify --network pawningshop notify\
docker run -d --name event_listener --network pawningshop --env-file prod.env event_listener\
docker run -dp 8545:8545 --network pawningshop -v ganache_volume:/app/ganache-data --name ganache trufflesuite/ganache-cli --mnemonic "jealous expect hundred young unlock disagree major siren surge acoustic machine catalog" --networkId 5777 --db /app/ganache-data\

## using secrets with docker compose
by default, docker-compose doen't support secret, but we do have a workaround:
docker-compose up -d
docker-compose exec mongo cat /run/secrets/db_username

## Setup MongoDB replica set for dev
```javascript
rs.initiate({
	_id : 'rsmongo',
	members: [
		{ _id : 0, host : "127.0.0.1:27017" },
		{ _id : 1, host : "127.0.0.1:27018" },
	]
});
```