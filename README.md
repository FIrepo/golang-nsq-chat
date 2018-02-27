A toy chat app
==============

This project takes ideas from [NSQ-Centric Architecture (GoCon Autumn 2014)][1]
slide.

## Overview

The app contains these parts, but you can add as many as you want:

- Chat server: Responsible for receiving & sending messages between clients.

- Archive daemon: Listen for Archive channel & save messages to Mongodb.

- Bot daemon: Listen for Bot channel, do some analysis, then reply.

```
 _____________
|   "Chat"    |                  Pub to "Chat" topic
|    Topic    |<-------------------------------------------------------.
|_____________|                                                        |
       |                                                               |
       |            _________________              .----> Client 1 --->|
       |           |     Server1     | Go channels |                   |
       |---------->|   NSQ channel   |------------>|----> Client 2 --->|
       |           |_________________|             |                   |
       |                                           '----> Client 3 --->|
       |            _________________                                  |
       |           |     Server2     | Go channels                     |
       |---------->|   NSQ channel   |------------> ...                |
       |           |_________________|                                 |
       |                                                               |
       |                                                               |
       ...                                                             |
       |            _________________                                  |
       |           |     Archive     |             .---------.         |
       |---------->|   NSQ channel   |------------>| Mongodb |         |
       |           |_________________|             |_________|         |
       |            _________________                                  |
       |           |       Bot       |             .---------------.   |
       |---------->|   NSQ channel   |------------>| NLP, commands |-->|
                   |_________________|             |_______________|

```

## Get started

1. Start Mongodb

```sh
mongod
```

2. Start nsq

```sh
docker-compose up -d

export NSQLOOKUPD_HTTP_ADDRESS=0.0.0.0:4161
export NSQD_HTTP_ADDRESS=0.0.0.0:4151
```

3. Get & Build

```sh
go get github.com/manhtai/golang-nsq-chat/cmd/chat
go get github.com/manhtai/golang-nsq-chat/cmd/archive
go get github.com/manhtai/golang-nsq-chat/cmd/bot
```

4. Run

- Chat server

```sh
./chat
```

- Archive daemon

```sh
./archive
```

- Bot daemon

```sh
./bot
```

## Generate cert.pem & key.pem

```sh
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

## Live reload for Chat server

```
go get https://github.com/Unknwon/bra
Bra run
```


[1]: https://www.slideshare.net/guregu/nsqcentric-architecture-gocon-autumn-2014
