# Concurrent Application Design with Go functions, channels and context

Getting started with concurrent application design.

Pointers on how an application can be structured
based on sharing memory (state) by communicating.

Communcation is by Go channels and timing contexts.

```
docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
go get golang.org/x/tools/cmd/present
setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH

present -http 0.0.0.0:3999 -orighost {docker_host_IP}
eg1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
eg2. present -http 0.0.0.0:3999 -orighost 127.0.0.1
```
