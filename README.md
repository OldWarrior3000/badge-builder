Tiny service to return images tags from AWS ecr in order to build badges

docker build -t bb .

***
docker build -t bb .
docker run --name bb1 -p 10000:10000 -e AWS_REGION=eu-central-1 -e REGISTRY_ID=XXXXXX -e AWS_ACCESS_KEY_ID=XXXXX -e AWS_SECRET_ACCESS_KEY=XXXXX bb
