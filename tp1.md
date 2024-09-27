# TP1 : Les premiers pas de bébé B1

### I. Récolte d'informations

🌞 Adresses IP de ta machine

```
> ip a
2: wlp3s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether b8:1e:a4:6c:56:97 brd ff:ff:ff:ff:ff:ff
    inet 10.33.77.206/20 brd 10.33.79.255 scope global dynamic noprefixroute wlp3s0
       valid_lft 81792sec preferred_lft 81792sec
    inet6 fe80::d7d6:514d:c48a:982f/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
```

🌞 Si t'as un accès internet normal, d'autres infos sont forcément dispos...

```
> ip route show 
default via 10.33.79.254 dev wlp3s0 proto dhcp src 10.33.77.206 metric 600 
10.33.64.0/20 dev wlp3s0 proto kernel scope link src 10.33.77.206 metric 600 
```

```
> resolvectl
Global
         Protocols: -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
  resolv.conf mode: stub

Link 2 (wlp3s0)
    Current Scopes: DNS
         Protocols: +DefaultRoute -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
Current DNS Server: 8.8.8.8
       DNS Servers: 8.8.8.8 1.1.1.1
```

```
> nmcli connection show 7c167f37-6ab6-4fbd-82ae-48daaa681582 | grep -i 'dhcp'
DHCP4.OPTION[3]:                        dhcp_server_identifier = 10.33.79.254
```

### II. Utiliser le réseau

🌞 Envoie un ping vers...

``` 
> ping 10.33.77.206
PING 10.33.77.206 (10.33.77.206) 56(84) bytes of data.
64 bytes from 10.33.77.206: icmp_seq=1 ttl=64 time=0.057 ms
64 bytes from 10.33.77.206: icmp_seq=2 ttl=64 time=0.047 ms
64 bytes from 10.33.77.206: icmp_seq=3 ttl=64 time=0.105 ms
64 bytes from 10.33.77.206: icmp_seq=4 ttl=64 time=0.048 ms
```

```
> ping 127.0.0.1
PING 127.0.0.1 (127.0.0.1) 56(84) bytes of data.
64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.034 ms
64 bytes from 127.0.0.1: icmp_seq=2 ttl=64 time=0.055 ms
64 bytes from 127.0.0.1: icmp_seq=3 ttl=64 time=0.082 ms
64 bytes from 127.0.0.1: icmp_seq=4 ttl=64 time=0.080 ms
```

🌞 On continue avec ping. Envoie un ping vers...

```
> ping 10.33.79.254
PING 10.33.79.254 (10.33.79.254) 56(84) bytes of data.
```

```
> ping 10.33.66.78 
PING 10.33.66.78 (10.33.66.78) 56(84) bytes of data.
64 bytes from 10.33.66.78: icmp_seq=1 ttl=64 time=159 ms
64 bytes from 10.33.66.78: icmp_seq=2 ttl=64 time=180 ms
64 bytes from 10.33.66.78: icmp_seq=3 ttl=64 time=202 ms
64 bytes from 10.33.66.78: icmp_seq=4 ttl=64 time=123 ms
```

```
> ping www.thinkerview.com
PING www.thinkerview.com (172.67.189.166) 56(84) bytes of data.
64 bytes from 172.67.189.166: icmp_seq=1 ttl=55 time=15.0 ms
64 bytes from 172.67.189.166: icmp_seq=2 ttl=55 time=19.6 ms
64 bytes from 172.67.189.166: icmp_seq=3 ttl=55 time=14.9 ms
64 bytes from 172.67.189.166: icmp_seq=4 ttl=55 time=18.5 ms
```

🌞 Faire une requête DNS à la main

```
> dig www.thinkerview.com
; <<>> DiG 9.18.28-0ubuntu0.24.04.1-Ubuntu <<>> www.thinkerview.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 49566
;; flags: qr rd ra; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;www.thinkerview.com.		IN	A

;; ANSWER SECTION:
www.thinkerview.com.	300	IN	A	188.114.96.6
www.thinkerview.com.	300	IN	A	188.114.97.6

;; Query time: 198 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Fri Sep 27 16:20:07 CEST 2024
;; MSG SIZE  rcvd: 80
```

```
> dig www.wikileaks.com
; <<>> DiG 9.18.28-0ubuntu0.24.04.1-Ubuntu <<>> www.wikileaks.org
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 29455
;; flags: qr rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;www.wikileaks.org.		IN	A

;; ANSWER SECTION:
www.wikileaks.org.	1800	IN	CNAME	wikileaks.org.
wikileaks.org.		1800	IN	A	80.81.248.21
wikileaks.org.		1800	IN	A	51.159.197.136

;; Query time: 199 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Fri Sep 27 16:21:08 CEST 2024
;; MSG SIZE  rcvd: 92
```

```
> dig www.torproject.com
; <<>> DiG 9.18.28-0ubuntu0.24.04.1-Ubuntu <<>> www.torproject.org
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 46452
;; flags: qr rd ra; QUERY: 1, ANSWER: 5, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;www.torproject.org.		IN	A

;; ANSWER SECTION:
www.torproject.org.	150	IN	A	95.216.163.36
www.torproject.org.	150	IN	A	116.202.120.165
www.torproject.org.	150	IN	A	116.202.120.166
www.torproject.org.	150	IN	A	204.8.99.144
www.torproject.org.	150	IN	A	204.8.99.146

;; Query time: 309 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Fri Sep 27 16:24:26 CEST 2024
;; MSG SIZE  rcvd: 127
```

### III. Sniffer le réseau

🌞 J'attends dans le dépôt git de rendu un fichier ping.pcap

*Fichier disponible sur mon git*

🌞 Livrez un deuxième fichier : dns.pcap

*Fichier disponible sur mon git*

### IV. Network scanning et adresses IP

🌞 Effectue un scan du réseau auquel tu es connecté

```
>  nmap -sn -PR 10.33.64.0/20
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-09-27 17:29 CEST
Nmap scan report for 10.33.66.78
Host is up (0.015s latency).
Nmap scan report for 10.33.70.82
Host is up (0.018s latency).
Nmap scan report for 10.33.71.156
Nmap scan report for 10.33.79.0
Host is up (0.047s latency).
Nmap scan report for 10.33.79.2
Host is up (0.22s latency).
Nmap scan report for 10.33.79.115
Host is up (0.017s latency).
Nmap done: 4096 IP addresses (63 hosts up) scanned in 101.87 seconds
```

🌞 Changer d'adresse IP

```
> sudo ip addr add 10.33.71.157/20 dev wlp3s0

> ip addr show wlp3s0
2: wlp3s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether b8:1e:a4:6c:56:97 brd ff:ff:ff:ff:ff:ff
    inet 10.33.77.206/20 brd 10.33.79.255 scope global dynamic noprefixroute wlp3s0
       valid_lft 73403sec preferred_lft 73403sec
    inet 10.33.71.157/20 scope global secondary wlp3s0
       valid_lft forever preferred_lft forever
    inet6 fe80::d7d6:514d:c48a:982f/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever

> sudo ip addr del 10.33.77.206/20 dev wlp3s0
ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: wlp3s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether b8:1e:a4:6c:56:97 brd ff:ff:ff:ff:ff:ff
    inet 10.33.71.157/20 scope global wlp3s0
       valid_lft forever preferred_lft forever
    inet6 fe80::d7d6:514d:c48a:982f/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: vboxnet0: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 0a:00:27:00:00:00 brd ff:ff:ff:ff:ff:ff

> sudo ip route add default via 10.33.79.254

> ping google.fr
PING google.fr (142.250.75.227) 56(84) bytes of data.
64 bytes from par10s41-in-f3.1e100.net (142.250.75.227): icmp_seq=1 ttl=117 time=102 ms
64 bytes from par10s41-in-f3.1e100.net (142.250.75.227): icmp_seq=2 ttl=117 time=15.3 ms
64 bytes from par10s41-in-f3.1e100.net (142.250.75.227): icmp_seq=3 ttl=117 time=19.0 ms
```