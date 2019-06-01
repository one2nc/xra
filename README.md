# xra

## Usage

```
(srehubenv) meson10@xps:$xra$ ./xra --help
Usage of ./xra:
  -client
    	Run in client mode?
  -config string
    	Path to Json (default "/tmp/goss.json")
  -debug
    	Run in debug mode?
  -host string
    	Hostname or IP (default "xps")
  -nmap
    	Nmap scan? [Required sudo]
```

## Server

```
./xra -config=/home/meson10/workspace/trusting_social/srehub/goss_validator/sample.json -host=10.54.0.12
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5001
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5005
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5008
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5009
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5010
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5000
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5002
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5003
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5004
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5006
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5007
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 5011
2019/05/30 13:16:08 server.go:78: Trying to listen on tcp Port: 22
2019/05/30 13:16:08 server.go:78: Trying to listen on udp Port: 6000
2019/05/30 13:16:08 server.go:78: Trying to listen on udp Port: 6001
2019/05/30 13:16:08 server.go:78: Trying to listen on udp Port: 6011
2019/05/30 13:16:08 server.go:32: listen tcp 0.0.0.0:22: bind: permission denied
^C2019/05/30 13:17:15 server.go:94: interrupt
2019/05/30 13:17:15 server.go:98: Exiting
```

## Client

```
./xra -client -config=/home/meson10/workspace/trusting_social/srehub/goss_validator/sample.json -host=10.54.0.12
fail, tcp, 10.54.0.12 -> cluster-example-new.uvweui.0001.aps1.cache.amazonaws.com:6380
fail, tcp, 10.54.0.12 -> cluster-example-new.uvweui.0001.aps1.cache.amazonaws.com:6379
pass, tcp, 10.54.0.12 -> 8.8.8.8:53
fail, tcp, 10.54.0.12 -> terraform-20190110072303326400000001.cttzy0nxd8dp.ap-south-1.rds.amazonaws.com:3306
fail, tcp, 10.54.0.12 -> terraform-20190110072303326400000001.cttzy0nxd8dp.ap-south-1.rds.amazonaws.com:3307
fail, tcp, 10.54.0.12 -> 10.54.0.26:5008
fail, tcp, 10.54.0.12 -> 10.54.0.26:5001
fail, tcp, 10.54.0.12 -> 10.54.0.26:5006
fail, tcp, 10.54.0.12 -> 10.54.0.26:5002
fail, tcp, 10.54.0.12 -> 10.54.0.26:5010
fail, tcp, 10.54.0.12 -> 10.54.0.26:5004
fail, tcp, 10.54.0.12 -> 13.233.206.75:22
fail, tcp, 10.54.0.12 -> 10.54.0.26:5003
fail, tcp, 10.54.0.12 -> 10.54.0.26:5005
fail, tcp, 10.54.0.12 -> 10.54.0.26:5007
fail, tcp, 10.54.0.12 -> 10.54.0.26:5011
fail, udp, 10.54.0.12 -> 10.54.0.26:6011
fail, tcp, 10.54.0.12 -> 10.54.0.26:5009
fail, udp, 10.54.0.12 -> 10.54.0.26:6001
fail, udp, 10.54.0.12 -> 10.54.0.26:6000
fail, tcp, 10.54.0.12 -> 10.54.0.26:5000

```
