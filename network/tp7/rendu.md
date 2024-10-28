# TP7 - On dit chiffrer pas crypter

## II. Serveur Web

### 1. HTTP

#### B. Configuration

🌞 Lister les ports en écoute sur la machine

```
[root@web ~]# ss -lnp | grep "nginx"
tcp   LISTEN 0      511                                       0.0.0.0:80               0.0.0.0:*    users:(("nginx",pid=11279,fd=6),("nginx",pid=11278,fd=6))                                    
tcp   LISTEN 0      511                                          [::]:80                  [::]:*    users:(("nginx",pid=11279,fd=7),("nginx",pid=11278,fd=7)
```

🌞 Ouvrir le port dans le firewall de la machine

```
sudo firewall-cmd --permanent --add-port=8888/tcp
sudo firewall-cmd --reload

[root@web ~]# sudo firewall-cmd --list-all
public (active)
  target: default
  icmp-block-inversion: no
  interfaces: enp0s3
  sources: 
  services: cockpit dhcpv6-client ssh
  ports: 80/tcp
  protocols: 
  forward: yes
  masquerade: no
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules:
```

#### C. Tests client

🌞 Vérifier que ça a pris effet

```
vincent@vinent-vbox:~$ ping sitedefou.tp7.b1
PING sitedefou.tp7.b1 (10.7.1.11) 56(84) bytes of data.
64 bytes from sitedefou.tp7.b1 (10.7.1.11): icmp_seq=1 ttl=64 time=0.599 ms
^C
--- sitedefou.tp7.b1 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 0.599/0.599/0.599/0.000 ms

vincent@vinent-vbox:~$ curl sitedefou.tp7.b1
rami malek nous contrôle tous

(j'ai pas écrit meow dans la page désolé)

```

#### D. Analyze

🌞 Capture tcp_http.pcap

-> `tcp_http.pcap`

🌞 Voir la connexion établie

```
vincent@vinent-vbox:~$ sudo ss -tnp | grep "10.7.1.11"
ESTAB 0      0               10.7.1.101:53546         10.7.1.11:80    users:(("firefox",pid=5033,fd=115))
```

### 2. On rajoute un S

#### A. Config

🌞 Lister les ports en écoute sur la machine

```
[root@web ~]# ss -tlnp | grep "nginx"
LISTEN 0      511        10.7.1.11:443       0.0.0.0:*    users:(("nginx",pid=11397,fd=6),("nginx",pid=11396,fd=6))
```

🌞 Gérer le firewall

```
[root@web ~]# sudo firewall-cmd --permanent --add-port=443/tcp
success
[root@web ~]# sudo firewall-cmd --permanent --remove-port=80/tcp
success
[root@web ~]# sudo firewall-cmd --reload
success
```

#### B. Test test test analyyyze

```
vincent@vinent-vbox:~$ curl -k https://sitedefou.tp7.b1
rami malek nous contrôle tous
```

🌞 Capture tcp_https.pcap
-> `tcp_https.pcap`

## III. Serveur VPN

### 1. Install et conf Wireguard

🌞 Prouvez que vous avez bien une nouvelle carte réseau wg0

```
ip a
4: wg0: <POINTOPOINT,NOARP,UP,LOWER_UP> mtu 1420 qdisc noqueue state UNKNOWN group default qlen 1000
    link/none 
    inet 10.7.200.1/24 scope global wg0
       valid_lft forever preferred_lft forever
```

🌞 Déterminer sur quel port écoute Wireguard

```
 [root@vpn ~]# ss -lnp | grep "51820"
udp   UNCONN 0      0                                         0.0.0.0:51820            0.0.0.0:*                                                                                                  
udp   UNCONN 0      0                                            [::]:51820               [::]:*
```

🌞 Ouvrez ce port dans le firewall

```
[root@vpn ~]# sudo firewall-cmd --permanent --add-port=51820/udp
sudo firewall-cmd --reload
```

### 3. Proofs

🌞 Ping ping ping !

```
vincent@vinent-vbox:~$ ping 10.7.200.1
PING 10.7.200.1 (10.7.200.1) 56(84) bytes of data.
64 bytes from 10.7.200.1: icmp_seq=1 ttl=64 time=0.983 ms
^C
--- 10.7.200.1 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 0.983/0.983/0.983/0.000 ms
```

🌞 Capture ping1_vpn.pcap

-> `ping1_vpn`

🌞 Capture ping2_vpn.pcap

-> `ping2_vpn`

🌞 Prouvez que vous avez toujours un accès internet

```
vincent@vinent-vbox:~$ traceroute 1.1.1.1
traceroute to 1.1.1.1 (1.1.1.1), 30 hops max, 60 byte packets
 1  _gateway (10.7.200.1)  1.709 ms  1.556 ms  1.715 ms
 2  10.7.1.254 (10.7.1.254)  2.445 ms  2.428 ms  2.415 ms
 3  10.0.2.2 (10.0.2.2)  2.589 ms  2.687 ms  2.675 ms
 4  * * *
 5  * * *
 6  * * *
 7  196.224.65.86.rev.sfr.net (86.65.224.196)  6.169 ms  5.932 ms  5.687 ms
 8  164.147.6.194.rev.sfr.net (194.6.147.164)  14.297 ms  14.063 ms  13.924 ms
 9  162.158.20.24 (162.158.20.24)  19.005 ms * *
10  162.158.20.31 (162.158.20.31)  19.129 ms  19.081 ms 162.158.20.48 (162.158.20.48)  19.043 ms
11  one.one.one.one (1.1.1.1)  18.978 ms  18.937 ms  18.992 ms
```

### 4. Private service

🌞 Visitez le service Web à travers le VPN

```
vincent@vinent-vbox:~$ curl -k https://sitedefou.tp7.b1
rami malek nous contrôle tous

vincent@vinent-vbox:~$ cat /etc/hosts
127.0.0.1 localhost
127.0.1.1 vinent-vbox
10.7.200.37 sitedefou.tp7.b1
```