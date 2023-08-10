# Net-cat

## Table-of-Contents
1. [Description](#description)
2. [Authors](#authors)
3. [Usage:](#usage)
4. [Implementation details: algorithm](#implementation-details-algorithm)
### Description:
***
Hi *Talent*!
This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

#### Go
***
- **GO**, also called Golang or Go language, is an open source programming language that Google developed.

### Authors:
***
+ Seynabou Niang - https://learn.zone01dakar.sn/git/sniang
* Masseck Thiaw - https://learn.zone01dakar.sn/git/mthiaw
- Vincent Félix Ndour - https://learn.zone01dakar.sn/git/vindour

### Usage:
***
A little intro about how to install:
```
$ git clone https://learn.zone01dakar.sn/git/vindour/net-cat.git
$ cd net-cat
$ Must open at least to terminal (one that handle the server and a other for client(s)) the chatroom can't exceed 10 clients 
```

[def]: #usage-how-to-run
Open terminals :
+ In the first terminal type : ***go run main.go [&& specified port if you don't want to run it on a default port]***
- In the other terminals type : ***nc localhost port [port that server run]***

## Implementation-details-algorithm:
***

Our program works in a similar way to net-cat by implementing a discussion group to be able to communicate between two or more clients through messages to do this we have created a HandleClient function which makes it possible to manage interactions between clients by: 
- recording the name
+ allowing to send messages 
* saving the history in a ".txt" file
- notifying of the exit of the client from the server... 

This function acts using a goroutine which itself is included in another goroutine processing incoming connections using two channels.
The first channel for saving entering connections and the last for shutdown the server.