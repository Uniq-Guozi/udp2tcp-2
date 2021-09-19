# udp2tcp
Change UDP to TCP
Fix Code From https://github.com/TongxiJi/udp2tcp

## Build
See cmd/build.sh

## Clinet
udp2tcp -c :11111 -a  (udp2tcp-server's ip):17002

## Server
udp2tcp -s :17002 -z  (app-server's ip):17002
