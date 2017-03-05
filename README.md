# goblink

A test blink program for use with GPIO on a Raspberry Pi.  Used in conjunction with Resin.io.


## Resin.io Configuration

Dockerfile.template:

```
FROM resin/%%RESIN_MACHINE_NAME%%-golang:latest

# enable systemd
ENV INITSYSTEM on

RUN go get github.com/benr/goblink

# copy start.sh script to root dir
COPY start.sh /

# run start.sh on start
CMD ["bash", "/start.sh"]
```

start.sh:

```
#!/bin/bash

./bin/goblink
```

## References

* https://github.com/stianeikeland/go-rpio/
