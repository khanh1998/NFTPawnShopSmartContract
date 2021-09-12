docker build -t api .
docker run -dp 4000:4000 --name api api