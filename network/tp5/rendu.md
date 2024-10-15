# TP5 : Un ptit LAN à nous

# I. Setup

☀️ Uniquement avec des commandes, prouvez-que :

### Les IP :

client 1 :
```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:a3:16:15 brd ff:ff:ff:ff:ff:ff
    inet 10.5.1.11/24 brd 10.5.1.255 scope global noprefixroute enp0s3
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fea3:1615/64 scope link 
       valid_lft forever preferred_lft forever
```

client 2:
```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:d8:6f:27 brd ff:ff:ff:ff:ff:ff
    inet 10.5.1.12/24 brd 10.5.1.255 scope global noprefixroute enp0s3
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fed8:6f27/64 scope link 
       valid_lft forever preferred_lft forever
```

routeur :
```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:27:d9:9b brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic noprefixroute enp0s3
       valid_lft 559sec preferred_lft 559sec
    inet6 fe80::a00:27ff:fe27:d99b/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:d5:b4:0b brd ff:ff:ff:ff:ff:ff
    inet 10.5.1.254/24 brd 10.5.1.255 scope global noprefixroute enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fed5:b40b/64 scope link 
       valid_lft forever preferred_lft forever
```

### Les hostname :

client 1 :
```
> hostnamectl
 Static hostname: client1
       Icon name: computer-vm
         Chassis: vm 🖴
      Machine ID: 120dcb95d3534d31baca3c21d27c90c9
         Boot ID: 464f692a207d4b74b0050b2a02963c07
  Virtualization: oracle
Operating System: Ubuntu 24.04.1 LTS              
          Kernel: Linux 6.8.0-45-generic
    Architecture: x86-64
 Hardware Vendor: innotek GmbH
  Hardware Model: VirtualBox
Firmware Version: VirtualBox
   Firmware Date: Fri 2006-12-01
    Firmware Age: 17y 10month 2w 
```

client2 :

```
> hostnamectl
 Static hostname: client2
       Icon name: computer-vm
         Chassis: vm 🖴
      Machine ID: 120dcb95d3534d31baca3c21d27c90c9
         Boot ID: 5731e9722dbc48d8aee3b64a9cdc8adc
  Virtualization: oracle
Operating System: Ubuntu 24.04.1 LTS              
          Kernel: Linux 6.8.0-45-generic
    Architecture: x86-64
 Hardware Vendor: innotek GmbH
  Hardware Model: VirtualBox
Firmware Version: VirtualBox
   Firmware Date: Fri 2006-12-01
    Firmware Age: 17y 10month 2w
```

router :
```
> hostnamectl
 Static hostname: router
       Icon name: computer-vm
         Chassis: vm 🖴
      Machine ID: 69be8b6b9c74446aa937c114634349bd
         Boot ID: ff3ffee76470477ab58bd3973cfe7a0f
  Virtualization: oracle
Operating System: Rocky Linux 9.4 (Blue Onyx)       
     CPE OS Name: cpe:/o:rocky:rocky:9::baseos
          Kernel: Linux 5.14.0-427.13.1.el9_4.x86_64
    Architecture: x86-64
 Hardware Vendor: innotek GmbH
  Hardware Model: VirtualBox
Firmware Version: VirtualBox
```

# II. Accès internet pour tous

## 1. Accès internet routeur

☀️ Déjà, prouvez que le routeur a un accès internet

```
> ping www.ynov.com
PING www.ynov.com (104.26.11.233) 56(84) bytes of data.
64 bytes from 104.26.11.233 (104.26.11.233): icmp_seq=1 ttl=51 time=99.1 ms
64 bytes from 104.26.11.233 (104.26.11.233): icmp_seq=2 ttl=51 time=278 ms
64 bytes from 104.26.11.233 (104.26.11.233): icmp_seq=3 ttl=51 time=198 ms
64 bytes from 104.26.11.233 (104.26.11.233): icmp_seq=4 ttl=51 time=72.3 ms
```

☀️ Activez le routage

```
> sudo firewall-cmd --add-masquerade --permanent
success
```

```
sudo firewall-cmd --reload
success
```

## 2. Accès internet clients

Client 1 :
```
> ip route show
default via 10.5.1.254 dev enp0s3 proto static metric 20100 
10.5.1.0/24 dev enp0s3 proto kernel scope link src 10.5.1.11 metric 100
```

Client 2 :

```
> ip route show
default via 10.5.1.254 dev enp0s3 proto static metric 20100 
10.5.1.0/24 dev enp0s3 proto kernel scope link src 10.5.1.12 metric 100
```

CLient 1 : 
```
> ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=114 time=82.5 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=114 time=105 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=114 time=118 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=114 time=232 ms
```

Client 2 :
```
> ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=114 time=155 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=114 time=172 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=114 time=189 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=114 time=305 ms
```

☀️ Prouvez que les clients ont un accès internet

Client 1 :
```
> ping www.google.com
PING www.google.com (142.250.201.4) 56(84) bytes of data.
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=1 ttl=111 time=18.2 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=2 ttl=111 time=22.4 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=3 ttl=111 time=18.1 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=4 ttl=111 time=22.8 ms
```


Client 2 : 

```
> ping www.google.com
PING www.google.com (142.250.201.4) 56(84) bytes of data.
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=1 ttl=111 time=22.7 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=2 ttl=111 time=17.6 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=3 ttl=111 time=22.9 ms
64 bytes from mrs08s19-in-f4.1e100.net (142.250.201.4): icmp_seq=4 ttl=111 time=22.3 ms
```

☀️ Montrez-moi le contenu final du fichier de configuration de l'interface réseau

```
sudo cat /etc/netplan/50-cloud-init.yaml 
[sudo] password for matt: 

# This file is generated from information provided by the datasource.  Changes
# to it will not persist across an instance reboot.  To disable cloud-init's
# network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
    ethernets:
        enp0s3:
            dhcp4: true
            addresses: [10.5.1.12/24]
            routes:
              - to: 0.0.0.0/0
                via: 10.5.1.254
            nameservers:
              addresses: [1.1.1.1]
    version: 2
```

# III. Serveur SSH

☀️ Sur routeur.tp5.b1, déterminer sur quel port écoute le serveur SSH

```
> sudo ss -lnpt | grep 22
LISTEN 0      128          0.0.0.0:22        0.0.0.0:*    users:(("sshd",pid=724,fd=3))
LISTEN 0      128             [::]:22           [::]:*    users:(("sshd",pid=724,fd=4))
```

☀️ Sur routeur.tp5.b1, vérifier que ce port est bien ouvert

```
sudo firewall-cmd --list-all
public (active)
  target: default
  icmp-block-inversion: no
  interfaces: enp0s3 enp0s8
  sources: 
  services: cockpit dhcpv6-client ssh
  ports: 22/tcp
  protocols: 
  forward: yes
  masquerade: yes
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules: 
```

# IV. Serveur DHCP

## 1. Le but

*Rien à rendre*

## 2. 2. Comment le faire

*Rien à rendre*

## 3. Rendu attendu

## A. Installation et configuration du serveur DHCP

Commandes utilisés :

- `sudo dnf install dhcp-server` (pour l'installation)

- `sudo nano /etc/dhcp/dhcpd.conf` (pour la configuration)

- `sudo systemctl start dhcpd` (démarrer le service)

- `sudo systemctl enable dhcpd` (activer le service au démarrage)

- `sudo systemctl status dhcpd` (vérifier son status)

Fichier de configuration : 

```
> sudo cat /etc/dhcp/dhcpd.conf
# Configuration du serveur DHCP

subnet 10.5.1.0 netmask 255.255.255.0 {
  range 10.5.1.137 10.5.1.237;        # Plage d'adresses IP attribuées
  option routers 10.5.1.254;        # Passerelle
  option subnet-mask 255.255.255.0; # Masque de sous-réseau
  option domain-name-servers 1.1.1.1; # Serveurs DNS
}
```

## B. Test avec un nouveau client

☀️ Créez une nouvelle machine client client3.tp5.b1

```
> hostnamectl
 Static hostname: client3
       Icon name: computer-vm
         Chassis: vm 🖴
      Machine ID: 120dcb95d3534d31baca3c21d27c90c9
         Boot ID: 1a9eec6d9cf3483ba830d9f1c27526d0
  Virtualization: oracle
Operating System: Ubuntu 24.04.1 LTS              
          Kernel: Linux 6.8.0-45-generic
    Architecture: x86-64
 Hardware Vendor: innotek GmbH
  Hardware Model: VirtualBox
Firmware Version: VirtualBox
   Firmware Date: Fri 2006-12-01
    Firmware Age: 17y 10month 2w
```

```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:98:77:ff brd ff:ff:ff:ff:ff:ff
    inet 10.5.1.137/24 brd 10.5.1.255 scope global dynamic noprefixroute enp0s3
       valid_lft 43132sec preferred_lft 43132sec
    inet6 fe80::a00:27ff:fe98:77ff/64 scope link 
       valid_lft forever preferred_lft forever
```

```
> ping www.google.com
PING www.google.com (172.217.20.196) 56(84) bytes of data.
64 bytes from waw02s08-in-f4.1e100.net (172.217.20.196): icmp_seq=1 ttl=114 time=14.9 ms
64 bytes from waw02s08-in-f4.1e100.net (172.217.20.196): icmp_seq=2 ttl=114 time=19.8 ms
64 bytes from waw02s08-in-f4.1e100.net (172.217.20.196): icmp_seq=3 ttl=114 time=15.1 ms
64 bytes from waw02s08-in-f4.1e100.net (172.217.20.196): icmp_seq=4 ttl=114 time=14.9 ms
```

## C. Consulter le bail DHCP

☀️ Consultez le bail DHCP qui a été créé pour notre client

```
> cat /var/lib/dhcpd/dhcpd.leases
# The format of this file is documented in the dhcpd.leases(5) manual page.
# This lease file was written by isc-dhcp-4.4.2b1

# authoring-byte-order entry is generated, DO NOT DELETE
authoring-byte-order little-endian;

server-duid "\000\001\000\001.\240\024\331\010\000'\325\264\013";

lease 10.5.1.137 {
  starts 1 2024/10/14 20:36:20;
  ends 2 2024/10/15 08:36:20;
  cltt 1 2024/10/14 20:36:20;
  binding state active;
  next binding state free;
  rewind binding state free;
  hardware ethernet 08:00:27:98:77:ff;
  uid "\001\010\000'\230w\377";
  client-hostname "client3";
}

```

☀️ Confirmez qu'il s'agit bien de la bonne adresse MAC

```
> ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:98:77:ff brd ff:ff:ff:ff:ff:ff
    inet 10.5.1.137/24 brd 10.5.1.255 scope global dynamic noprefixroute enp0s3
       valid_lft 42680sec preferred_lft 42680sec
    inet6 fe80::a00:27ff:fe98:77ff/64 scope link 
       valid_lft forever preferred_lft forever
```