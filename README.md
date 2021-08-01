# GoFnaf
a recreation of five night at freddy's as a text game in GO 

## What works
cameras, battery, enemies, doors

## what i still have to fix/implement
1)
Golden freddy, 

2)
someting that most of the times makes you notice that an animatronic is at your door

3)
a versus mode where one player controls the animatronic and the other controls the night guard

## usage
now that the game has an actual game client (if a bit rough) l'll explain how to run the server and the client

### server
you can decide wether to run the server via a docker container or via the .go file without docker,

### docker
to run this image with docker you have to run 


docker run --rm --name gofnafserver -p 8080:8080 -t -d rickytrevor/gofnafdocker

i'll update this image as soon as i publish a new version of the server

### "standard" way

to run the server in the "standard" you have to build it by going inside the serverside directory and running go build and then go run gofnaf to actually start the server

### client

the procedure is mostly the same, go inside the clientside directory and run go build . and then run the client executable file

### if you want to play the game without hosting an actual client at home

so if you really want to play this "game" i'm hosting a game server on my personal server, you can access it by starting the client and putting as the ip:port

207.180.204.188:8080 (it'll probalby change in the future)
