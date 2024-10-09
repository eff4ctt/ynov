# TP2 : Hey yo tell a neighbor tell a friend

## I. Simplest LAN

### 1. Quelques pings

🌞 Prouvez que votre configuration est effective

```
> ip addr show vboxnet0
3: vboxnet0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 0a:00:27:00:00:00 brd ff:ff:ff:ff:ff:ff
    inet 192.168.56.1/24 brd 192.168.56.255 scope global vboxnet0
       valid_lft forever preferred_lft forever
    inet6 fe80::800:27ff:fe00:0/64 scope link 
       valid_lft forever preferred_lft forever
```

🌞 Tester que votre LAN + votre adressage IP est fonctionnel

ping depuis la VM (en ssh) : 

```
> ping 192.168.56.1
PING 192.168.56.1 (192.168.56.1) 56(84) bytes of data.
64 bytes from 192.168.56.1: icmp_seq=1 ttl=64 time=0.409 ms
64 bytes from 192.168.56.1: icmp_seq=2 ttl=64 time=0.338 ms
64 bytes from 192.168.56.1: icmp_seq=3 ttl=64 time=0.389 ms
64 bytes from 192.168.56.1: icmp_seq=4 ttl=64 time=0.492 ms
```

ping depuis ma machine : 

```
> ping 192.168.56.104
PING 192.168.56.104 (192.168.56.104) 56(84) bytes of data.
64 bytes from 192.168.56.104: icmp_seq=1 ttl=64 time=0.677 ms
64 bytes from 192.168.56.104: icmp_seq=2 ttl=64 time=0.665 ms
64 bytes from 192.168.56.104: icmp_seq=3 ttl=64 time=0.763 ms
64 bytes from 192.168.56.104: icmp_seq=4 ttl=64 time=0.668 ms
```

🌞 Capture de ping

La capture -> `ping.pcapng`

## II. Utilisation des ports


🌞 Sur le PC serveur (ma machine)

`nc -l 6969`

🌞 Sur le PC serveur toujours

```
> sudo ss -lntp
LISTEN    0         1                  0.0.0.0:6969             0.0.0.0:*
```

🌞 Sur le PC client (ma VM)

`nc 192.168.56.1 6969`

🌞 Echangez-vous des messages

```
> nc 192.168.56.1 6969

ça marche ? (serveur)
oui c'est super (client)
```

🌞 Utilisez une commande qui permet de voir la connexion en cours

```
> sudo ss -lnpt
LISTEN    0         1                  0.0.0.0:6969             0.0.0.0:*        users:(("nc",pid=18721,fd=3))
```

🌞 Faites une capture Wireshark complète d'un échange

La capture -> `netcat1.pcapng`

🌞 Inversez les rôles

Sur le pc serveur (ma VM)

```
> sudo ss -lnpt
LISTEN    0         1                  0.0.0.0:2323             0.0.0.0:*        users:(("nc",pid=8182,fd=3))
```
netcat2 : 

La capture -> `netcat2.pcapng`


## III. Analyse de vos applications usuelles

### 1. Serveur web

```
> nc repubblica.it 80

GET / HTTP/1.1
Host: repubblica.it

HTTP/1.1 301 Moved Permanently
Content-Length: 0
Connection: keep-alive
Cache-Control: private, no-cache, no-store, must-revalidate, pre-check=0, post-check=0
Date: Mon, 07 Oct 2024 20:37:13 GMT
Expires: Tue, 31 Mar 1981 05:00:00 GMT
Location: https://www.repubblica.it/
Server: Varnish
X-Varnish: 613047358
X-Cache: Miss from cloudfront
Via: 1.1 56f08e51c16f365de3e0991809e86e7c.cloudfront.net (CloudFront)
X-Amz-Cf-Pop: CDG52-P5
X-Amz-Cf-Id: T71jTQEb04KNpp0je2jTIlPRsAUwUTHI12d3C4s7fUPSU4c1BTUQIQ==
```

🌞 Utilisez Wireshark pour capturer du trafic HTTP

La capture -> `http.pcapng`

🌞 Pour les 5 applications

Discord :

```
> ss -tnp | grep -i "discord"
ESTAB 0      0       192.168.1.42:52224 162.159.130.234:443  users:(("Discord",pid=30828,fd=25))
```

La capture -> `discord.pcapng`

Firefox :

```
> ss -tnp | grep -i "firefox"
ESTAB 0      0                                192.168.1.42:48996            140.82.113.25:443  users:(("firefox",pid=4012,fd=50)) 
ESTAB 0      0                                192.168.1.42:48602            34.107.243.93:443  users:(("firefox",pid=4012,fd=53))
```

*La capture -> `firefox.pcapng`

Steam :

```
> ss -tnp | grep -i "steam"
ESTAB 0      0                                192.168.1.42:55964               23.33.233.98:443   users:(("steamwebhelper",pid=52011,fd=40))
ESTAB 0      0                                   127.0.0.1:37997                  127.0.0.1:57230 users:(("steam",pid=51594,fd=98))         
ESTAB 0      0                                   127.0.0.1:39222                  127.0.0.1:27060 users:(("steamwebhelper",pid=52011,fd=51))
ESTAB 0      0                                192.168.1.42:55968               23.33.233.98:443   users:(("steamwebhelper",pid=52011,fd=42))
ESTAB 0      0                                192.168.1.42:60146               2.18.131.137:443   users:(("steamwebhelper",pid=52011,fd=46))
ESTAB 0      0                                192.168.1.42:55994               23.33.233.98:443   users:(("steamwebhelper",pid=52011,fd=62))
ESTAB 0      0                                192.168.1.42:55982               23.33.233.98:443   users:(("steamwebhelper",pid=52011,fd=43))
ESTAB 0      0                                   127.0.0.1:44448                  127.0.0.1:36467 users:(("steamwebhelper",pid=52011,fd=29))
ESTAB 0      0                                   127.0.0.1:36467                  127.0.0.1:44448 users:(("steam",pid=51594,fd=99))         
ESTAB 0      0                                192.168.1.42:48783             162.254.196.79:27019 users:(("steam",pid=51594,fd=104))        
ESTAB 0      0                                   127.0.0.1:57230                  127.0.0.1:37997 users:(("steamwebhelper",pid=52011,fd=28))
ESTAB 0      0                                   127.0.0.1:27060                  127.0.0.1:39222 users:(("steam",pid=51594,fd=149))        
ESTAB 0      0                                192.168.1.42:58804             172.64.145.151:443   users:(("steamwebhelper",pid=52011,fd=48))
ESTAB 0      0                                192.168.1.42:58782             172.64.145.151:443   users:(("steamwebhelper",pid=52011,fd=49))
ESTAB 0      0                                192.168.1.42:36702               23.218.19.90:443   users:(("steamwebhelper",pid=52011,fd=34))
ESTAB 0      0                                192.168.1.42:52868               96.16.122.82:443   users:(("steamwebhelper",pid=52011,fd=39))
```

La capture -> `steam.pcapng`

Telegram Desktop :

```
> ss -tnp | grep -i "telegram"
ESTAB 0      0                                192.168.1.42:59324              149.154.167.92:443  users:(("telegram-deskto",pid=56155,fd=63))
```

La capture -> `telegram.pcapng`

Github Desktop :

```
> ss -tnp | grep -i "github"
ESTAB 0      0                                192.168.1.42:44886         140.82.112.26:443  users:(("github-desktop",pid=60838,fd=27))
ESTAB 0      0                                192.168.1.42:53814         140.82.114.21:443  users:(("github-desktop",pid=60838,fd=28))
ESTAB 0      0                                192.168.1.42:37906          140.82.121.6:443  users:(("github-desktop",pid=60838,fd=26))
```

La capture -> `github.pcapng`

## *Ecrit par eff4ctt*