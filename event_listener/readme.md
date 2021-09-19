docker build -t event_listener . \
docker run -dp 7789:7789 -e ENV=PROD --name event_listener event_listener \

When run in localhost: `ENV=DEV go run .` \
When run build version in server or localhost: `ENV=PROD ./khanh` \
When run inside docker container using docker compose: `./khanh` \
Because the `--env-file` of event_listener is specified in docker-compose.yml, so you don't have to specify env by yourself. \
If you run this container with simple `docker run`, you have to add `ENV=PROD` or `ENV=DEV` environment variable to the command.

When ENV=PROD, the application read env variables from `prod.env` file \
When ENV=DEV, the application read env variables from `dev.env` file \
When env is not specified, it read directly from system environment using `os.Readenv`