docker network create pawningshop
docker volume create ganache_volume
docker run -dp 80:80 --name client --network pawningshop client
docker run -dp 4000:4000 --name api --network pawningshop api
docker run -dp 7789:7789 --name notify --network pawningshop notify
docker run -d --name event_listener --network pawningshop --env-file prod.env event_listener
docker run -dp 8545:8545 --network pawningshop -v ganache_volume:/app/ganache-data --name ganache trufflesuite/ganache-cli --mnemonic "jealous expect hundred young unlock disagree major siren surge acoustic machine catalog" --networkId 5777 --db /app/ganache-data 