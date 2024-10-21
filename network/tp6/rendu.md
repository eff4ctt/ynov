# TP6 : Des bo services dans des bo LANs

### 0. Prérequis

*Rien à rendre*

### I. Le setup

*Rien à rendre*

### 1. Tableau d'adressage

*Rien à rendre*

### 2. Marche à suivre

☀️ Prouvez que...

Ping www.ynov.com depuis le serveur dhcp :

```
> ping www.ynov.com
PING www.ynov.com (172.67.74.226) 56(84) bytes of data.
64 bytes from 172.67.74.226 (172.67.74.226): icmp_seq=1 ttl=61 time=206 ms
64 bytes from 172.67.74.226 (172.67.74.226): icmp_seq=2 ttl=61 time=20.3 ms
64 bytes from 172.67.74.226 (172.67.74.226): icmp_seq=3 ttl=61 time=35.8 ms
64 bytes from 172.67.74.226 (172.67.74.226): icmp_seq=4 ttl=61 time=72.2 ms
```


Ping www.google.com depuis le serveur web :

```
> ping www.google.com
PING www.google.com (142.250.75.228) 56(84) bytes of data.
64 bytes from par10s41-in-f4.1e100.net (142.250.75.228): icmp_seq=1 ttl=61 time=100 ms
64 bytes from par10s41-in-f4.1e100.net (142.250.75.228): icmp_seq=2 ttl=61 time=18.9 ms
64 bytes from par10s41-in-f4.1e100.net (142.250.75.228): icmp_seq=3 ttl=61 time=34.5 ms
64 bytes from par10s41-in-f4.1e100.net (142.250.75.228): icmp_seq=4 ttl=61 time=22.9 ms
```

Ping le serveur DNS depuis le serveur DHCP :

```
>   ping 10.6.2.12
PING 10.6.2.12 (10.6.2.12) 56(84) bytes of data.
64 bytes from 10.6.2.12: icmp_seq=1 ttl=63 time=3.12 ms
64 bytes from 10.6.2.12: icmp_seq=2 ttl=63 time=2.81 ms
64 bytes from 10.6.2.12: icmp_seq=3 ttl=63 time=4.07 ms
64 bytes from 10.6.2.12: icmp_seq=4 ttl=63 time=2.39 ms
```

# II. LAN clients

### 1. Serveur DHCP

*Rien à rendre*

### 2. Client

☀️ Prouvez que...

```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:d1:69:d2 brd ff:ff:ff:ff:ff:ff
    inet 10.6.1.137/24 brd 10.6.1.255 scope global dynamic noprefixroute enp0s3
       valid_lft 43068sec preferred_lft 43068sec
    inet6 fe80::a00:27ff:fed1:69d2/64 scope link 
       valid_lft forever preferred_lft forever
```

```
> resolvectl
Global
         Protocols: -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
  resolv.conf mode: stub

Link 2 (enp0s3)
    Current Scopes: DNS
         Protocols: +DefaultRoute -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
Current DNS Server: 1.1.1.1
       DNS Servers: 1.1.1.1

```

```
> ip route show 
default via 10.6.1.254 dev enp0s3 proto dhcp src 10.6.1.137 metric 20100 
10.6.1.0/24 dev enp0s3 proto kernel scope link src 10.6.1.137 metric 100
```

```
> ping ww.ynov.com
PING ww.ynov.com (104.26.10.233) 56(84) bytes of data.
64 bytes from 104.26.10.233: icmp_seq=1 ttl=61 time=79.5 ms
64 bytes from 104.26.10.233: icmp_seq=2 ttl=61 time=115 ms
64 bytes from 104.26.10.233: icmp_seq=3 ttl=61 time=38.0 ms
64 bytes from 104.26.10.233: icmp_seq=4 ttl=61 time=21.7 ms
```

# III. LAN serveurzzzz

### 1. Intro serveur Web

*Rien à rendre*

### 2. Install this shiet

*Rien à rendre*

### 3. Analyse et test

☀️ Déterminer sur quel port écoute le serveur NGINX

```
> sudo ss -lnpt | grep 80
LISTEN 0      511          0.0.0.0:80        0.0.0.0:*    users:(("nginx",pid=1583,fd=6),("nginx",pid=1582,fd=6),("nginx",pid=1581,fd=6))
LISTEN 0      511             [::]:80           [::]:*    users:(("nginx",pid=1583,fd=7),("nginx",pid=1582,fd=7),("nginx",pid=1581,fd=7))
```

☀️ Ouvrir ce port dans le firewall

```
> sudo firewall-cmd --permanent --add-port=80/tcp
success
```

```
> sudo firewall-cmd --reload
success
```
☀️ Visitez le site web !

```
> curl http://10.6.2.11
<!doctype html>
<html>
  <head>
    <meta charset='utf-8'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <title>HTTP Server Test Page powered by: Rocky Linux</title>
    <style type="text/css">
      /*<![CDATA[*/
      
      html {
        height: 100%;
        width: 100%;
      }  

```

# 2. Serveur DNS

 ### 1. Présentation serveur DNS

 *Rien à rendre*

 ### 2. Dans notre TP

 *Rien à rendre*

 ### 3. Zé bardi

 *Rien à rendre*

 ### 4. Analyse du service

 ☀️ Déterminer sur quel(s) port(s) écoute le service BIND9

 ```
> sudo ss -lnpt | grep 53
LISTEN 0      4096       127.0.0.1:953       0.0.0.0:*    users:(("named",pid=2115,fd=31))
LISTEN 0      10         127.0.0.1:53        0.0.0.0:*    users:(("named",pid=2115,fd=34))
LISTEN 0      10         127.0.0.1:53        0.0.0.0:*    users:(("named",pid=2115,fd=35))
LISTEN 0      10             [::1]:53           [::]:*    users:(("named",pid=2115,fd=41))
LISTEN 0      10             [::1]:53           [::]:*    users:(("named",pid=2115,fd=40))
LISTEN 0      4096           [::1]:953          [::]:*    users:(("named",pid=2115,fd=42))
```

 ☀️ Ouvrir ce(s) port(s) dans le firewall

```
 sudo firewall-cmd --permanent --add-port=53/tcp
success
```

```
> sudo firewall-cmd --reload
success
```

### 5. Tests manuels

☀️ Effectuez des requêtes DNS manuellement depuis le serveur DNS lui-même dans un premier temps

```
> dig web.tp6.b1 @10.6.2.12
;; ANSWER SECTION:
web.tp6.b1.		86400	IN	A	10.6.2.11
```

```
> dig dns.tp6.b1 @10.6.2.12
;; ANSWER SECTION:
dns.tp6.b1.		86400	IN	A	10.6.2.12
```

```
> dig ynov.com @10.6.2.12
;; ANSWER SECTION:
ynov.com.		292	IN	A	172.67.74.226
ynov.com.		292	IN	A	104.26.10.233
ynov.com.		292	IN	A	104.26.11.233
```

```
> dig -x 10.6.2.11 @10.6.2.12
;; ANSWER SECTION:
11.2.6.10.in-addr.arpa.	86400	IN	PTR	web.tp6.b1.
```

```
> dig -x 10.6.2.12 @10.6.2.12
;; ANSWER SECTION:
12.2.6.10.in-addr.arpa.	86400	IN	PTR	dns.tp6.b1.
```

☀️ Effectuez une requête DNS manuellement depuis client1.tp6.b1

```
> dig web.tp6.b1 @10.6.2.12
;; ANSWER SECTION:
web.tp6.b1.		86400	IN	A	10.6.2.11
```

☀️ Capturez une requête DNS et la réponse de votre serveur

voir -> `dns.pcapng`

# 3. Serveur DHCP

☀️ Créez un nouveau client `client2.tp6.b1` vitefé

```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:d1:61:fe brd ff:ff:ff:ff:ff:ff
    inet 10.6.1.139/24 brd 10.6.1.255 scope global dynamic noprefixroute enp0s3
       valid_lft 43102sec preferred_lft 43102sec
    inet6 fe80::a00:27ff:fed1:61fe/64 scope link 
       valid_lft forever preferred_lft forever
```

```
> resolvectl
Global
         Protocols: -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
  resolv.conf mode: stub

Link 2 (enp0s3)
    Current Scopes: DNS
         Protocols: +DefaultRoute -LLMNR -mDNS -DNSOverTLS DNSSEC=no/unsupported
Current DNS Server: 10.6.2.12
       DNS Servers: 10.6.2.12
```

```
> curl http://web.tp6.b1
<!doctype html>
<html>
  <head>
    <meta charset='utf-8'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <title>HTTP Server Test Page powered by: Rocky Linux</title>
    <style type="text/css">
      /*<![CDATA[*/
      
      html {
        height: 100%;
        width: 100%;
      }  
```

# GGWP ! 