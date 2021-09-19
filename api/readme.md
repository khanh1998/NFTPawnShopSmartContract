docker build -t api . \
docker run -dp 4000:4000 -e ENV=PROD --name api api \

When run in localhost: `ENV=DEV go run .` \
When run build version in localhost: `ENV=PROD ./NFTPawningShopBackend` \
When run build version inside docker container: `./NFTPawningShopBackend` \
Don't need to specify env when run program inside docker container because the `--env-file` is specified in `docker-compose.yml`. \
If you run this container with simple `docker run`, you have to add `ENV=PROD` or `ENV=DEV` environment variable to the command.

When ENV=PROD, the application read env variables from `prod.env` file \
When ENV=DEV, the application read env variables from `dev.env` file \
When env is not specified, it read directly from system environment using `os.Readenv`