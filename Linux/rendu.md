# TP Avance : "Mission Ultime : Sauvegarde et Securisation"

### Etape 1 : Analyse et nettoyage du serveur 

- Lister les taches cron pour detecter des backdoors :

    `cat /etc/passwd`
    
    `crontab -l -u attacker` 

- Identifier et supprimer les fichiers caches :

    `cat script.sh`, `cat /tmp/.hidden_script`, `rm .hidden_script .hidden_file malicious.sh`

- Analyser les connexions reseau actives :

    `netstat -tulnp`

### Etape 2 : Configuration avancee de LVM

Creer un snapshot de securite pour /mnt/secure_data :

    ```bash
    lvcreate --size 500M --snapshot --name secure_data_snapshot /dev/vg_secure/secure_data
    ```

Tester la restauration du snapshot :

`rm /mnt/secure_data/sensitive1.txt`

`mkdir /mnt/secure_data_snapshot`

`mount /dev/vg_secure/secure_data_snapshot /mnt/secure_data_snapshot`

`cp /mnt/secure_data_snapshot/nom_du_fichier /mnt/secure_data/`

`umount /mnt/secure_data_snapshot`

`lvremove /dev/vg_secure/secure_data_snapshot`

Optimiser l'espace disque :

`lvextend --size +10G /dev/vg_secure/secure_data`

`resize2fs /dev/vg_secure/secure_data`

### Etape 3 : Automatisation avec un script de sauvegarde

Créer un script secure_backup.sh :

`nano /usr/local/bin/secure_backup.sh`

Ajoutez une fonction de rotation des sauvegardes :

Code du script :

```bash
#!/bin/bash

SOURCE_DIR="/mnt/secure_data"
BACKUP_DIR="/backup"
DATE=$(date +%Y%m%d)
BACKUP_FILE="$BACKUP_DIR/secure_data_$DATE.tar.gz"

mkdir -p $BACKUP_DIR

tar --exclude='*/.*/' --exclude='*.tmp' --exclude='*.log' -czf $BACKUP_FILE -C $SOURCE_DIR .

if [ $? -eq 0 ]; then
    echo "La sauvegarde a réussi ! Le fichier est sauvegardé sous $BACKUP_FILE"
else
    echo "La sauvegarde a échoué."
    exit 1
fi

backup_count=$(ls $BACKUP_DIR/secure_data_*.tar.gz | wc -l)

if [ $backup_count -gt 7 ]; then
    ls -t $BACKUP_DIR/secure_data_*.tar.gz | tail -n +8 | xargs rm -f
    echo "Les anciennes sauvegardes ont été supprimées. Seules les 7 dernières sauvegardes sont conservées."
else
    echo "Pas de rotation nécessaire, moins de 7 sauvegardes existantes."
fi
```

Testez le script :

`/usr/local/bin/secure_backup.sh`

Automatisez avec une tâche cron :

`crontab -e`

`0 3 * * * /usr/local/bin/secure_backup.sh`

### Étape 4 : Surveillance avancée avec auditd

Configurer auditd pour surveiller /etc :

`auditctl -w /etc -p wa`

Tester la surveillance :

`touch /etc/test_file`

`echo "test" > /etc/test_file`

`cat /var/log/audit/audit.log | grep /etc`

Analyser les événements :

`ausearch -f /etc`

`ausearch -f /etc > /var/log/audit_etc.log`

### Étape 5 : Sécurisation avec Firewalld

Configurer un pare-feu pour SSH et HTTP/HTTPS uniquement :

`systemctl start firewalld`

`systemctl enable firewalld`

`firewall-cmd --permanent --add-service=ssh`

`firewall-cmd --permanent --add-service=http`

`firewall-cmd --permanent --add-service=https`

`firewall-cmd --permanent --set-default-zone=drop`

`firewall-cmd --reload`

Bloquer des IP suspectes :

`firewall-cmd --permanent --add-rich-rule='rule family="ipv4" source address="192.168.56.0" reject'`

`firewall-cmd --reload`

Restreindre SSH à un sous-réseau spécifique :

`firewall-cmd --permanent --add-rich-rule='rule family="ipv4" service name="ssh" source address="192.168.56.0/24" accept'`

`firewall-cmd --permanent --add-rich-rule='rule family="ipv4" service name="ssh" reject'`

`firewall-cmd --reload`

# PARTIE - DLC (BONUS)

### Étape 1 : Analyse avancée et suppression des traces suspectes

Rechercher des utilisateurs récemment ajoutés :

`sudo grep "new user" /var/log/secure`

Trouver les fichiers récemment modifiés dans des répertoires critiques :

`sudo find /etc /usr/local/bin /var -type f -mtime -7`

Lister les services suspects activés :

`sudo systemctl list-unit-files --state=enabled`

Supprimer une tâche cron suspecte :

`sudo crontab -u attacker -r`

### Étape 2 : Configuration avancée de LVM

Créer un snapshot du volume logique :

`lvcreate --size 500M --snapshot --name secure_data_snapshot /dev/vg_secure/secure_data`

Tester le snapshot :

`mount /dev/vg_secure/secure_data_snapshot /mnt/snapshot`

`ls /mnt/snapshot`

Simuler une restauration :

`rm /mnt/secure_data/test_file`

`cp /mnt/snapshot/test_file /mnt/secure_data/`

### Étape 3 : Renforcement du pare-feu avec des règles dynamiques

Bloquer les attaques par force brute :

```bash
firewall-cmd --permanent --add-rich-rule='rule family="ipv4" service name="ssh" limit value="5/m" reject'
firewall-cmd --reload
```

Restreindre l’accès SSH à une plage IP spécifique :

```bash
firewall-cmd --permanent --add-rich-rule='rule family="ipv4" service name="ssh" source address="192.168.x.0/24" accept'
firewall-cmd --permanent --add-rich-rule='rule family="ipv4" service name="ssh" reject'
firewall-cmd --reload
```

```bash
sudo firewall-cmd --permanent --new-zone=webzone
sudo firewall-cmd --permanent --zone=webzone --add-service=http
sudo firewall-cmd --permanent --zone=webzone --add-service=https
sudo firewall-cmd --permanent --zone=webzone --add-interface=eth0
sudo firewall-cmd --reload
```

### Étape 4 : Création d'un script de surveillance avancé

Ecrivez un script monitor.sh :

```bash
#!/bin/bash

LOG_FILE="/var/log/monitor.log"

echo "$(date) - Connexions actives :" >> $LOG_FILE
ss -tuln >> $LOG_FILE

echo "$(date) - Fichiers modifiés dans /etc :" >> $LOG_FILE
find /etc -type f -mmin -5 >> $LOG_FILE

if [ "$(find /etc -type f -mmin -5 | wc -l)" -gt 0 ]; then
    echo "Alerte : Fichier modifié dans /etc" | mail -s "Alerte - Fichier modifié" admin@example.com
fi
```

Ajoutez une alerte par e-mail :

`apt-get install mailutils `

(Configuration du serveur)

Automatisez le script :

`crontab -e`

`*/5 * * * * /path/to/monitor.sh`

### Étape 5 : Mise en place d’un IDS (Intrusion Detection System)

Installer et configurer AIDE :

`yum install aide`

`aide --init`

`vi /etc/aide/aide.conf`

```bash
/etc           FIPSR
/bin           FIPSR
/sbin          FIPSR
/usr/bin       FIPSR
```

`aide --update`

Tester AIDE :

`aide --check`

# FIN DU TP !

