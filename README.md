# MedellinC2
Welcome to Medellin C2! The development of this C2 is being used to enhance my coding knowledge, and learn more about red teaming. 

### Structure
```
server
   | 
    server.go
data
   | 
    data.go
cmd
   | 
    create.go
    init.go
    list.go
    listeners.go
    root.go
    server.go
    start.go
MedellinC2
```
![server](server_design.PNG)

![schema](schema.PNG)

### Commands
- `./MedellinC2 init`: Initializes the databases 
- `./MedellinC2 listeners`: Displays the listeners menu
- `./MedellinC2 listeners create`: Create a new listener
- `./MedellinC2 server`: Displays the server menu
- `./MedellinC2 server start`: Starts the C2 server, which allows listeners to listen for connections from agents