Build the image:

```sh
PROJECT=$(basename `pwd`)
```
```sh
docker image build -t $PROJECT-image . --build-arg user_id=`id -u` --build-arg group_id=`id -g`
```

Run docker containers:

```sh
docker container run \
  -it \
  --rm \
  --init \
  --mount type=bind,src=`pwd`,dst=/app \
  --name $PROJECT-container \
  $PROJECT-image \
  /bin/zsh
```

## Compile your app inside the Docker container

To compile, but not run your app inside the Docker instance, you can write something like:

```sh
docker container run \
  --rm \
  --mount type=bind,src=`pwd`,dst=/usr/src/myapp \
  -w /usr/src/myapp \
  --name $PROJECT-container \
  $PROJECT-image \
  go build -v hello.go
```
