testport is a command to test a network connection to a server/port 

Usage:
```
testport <server> <port> [<network>] [<timeout>]
Defaults:
network: tcp
timeout: 5
```

Examples:

    # check port 80 on localhost
    testport localhost 80
    # check port 1080 (udp) on localhost
    testport localhost 1080 udp
    # check port 443 (tcp) on localhost (with a 10 second timeout)
    testport localhost 443 tcp 10
