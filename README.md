# Study Coroutine

Study ST(state-threads) and Go(goroutine).

## Usage

First, clone this git and start docker:

```
docker pull ossrs/dev &&
git clone https://github.com/winlinvip/study-coroutine.git && 
cd study-coroutine && docker run -it -v `pwd`:/tmp -w /tmp ossrs/dev bash
```

Second, build all:

```
(cd st-server/st-1.9 && make EXTRA_CFLAGS=-DMALLOC_STACK linux-debug) &&
(cd st-server && g++ -Ist-1.9/obj -g -O0 main.cpp st-1.9/obj/libst.a -o st-server) &&
(cd go-client && go build .) &&
(cd go-server && go build .)
```

## ST(state-threads)

Run st-server:

```
docker run -it -v `pwd`:/tmp -w /tmp ossrs/dev ./st-server/st-server
```

Run go-client with 10k connections:

```
dockerID=`docker ps --format "{{.ID}} {{.Image}}" |grep 'ossrs/dev' |awk '{print $1}'` &&
docker exec -it $dockerID ./go-client/go-client -c 60000 -s 90
```

To stop server:

```
docker kill `docker ps --format "{{.ID}} {{.Image}}" |grep 'ossrs/dev' |awk '{print $1}'`
```

## Go

Run go-server:

```
docker run -it -v `pwd`:/tmp -w /tmp ossrs/dev ./go-server/go-server
```

Run go-client with 10k connections:

```
dockerID=`docker ps --format "{{.ID}} {{.Image}}" |grep 'ossrs/dev' |awk '{print $1}'` &&
docker exec -it $dockerID ./go-client/go-client -c 60000 -s 90
```

To stop server:

```
docker kill `docker ps --format "{{.ID}} {{.Image}}" |grep 'ossrs/dev' |awk '{print $1}'`
```

