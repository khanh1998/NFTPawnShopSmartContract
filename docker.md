docker network create pawningshop
docker run -dp 80:80 --name client --network pawningshop client
docker run -dp 4000:4000 --name api --network pawningshop api
docker run -dp 7789:7789 --name notify --network pawningshop notify
