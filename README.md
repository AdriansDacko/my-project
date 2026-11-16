# flask-demo-12-2026



## Kā darbināt ssh tuneli:

```sudo apt update```

```sudo apt install sshpass openssh-client -y```


Izveidot failu atbilstoši sava datora numuram:
1. dators - 9001
2. dators - 9002
3. utt.
piemēram 1. datoram
```sudo nano /etc/systemd/system/flask-tunnel-9001.service```

ievietot failā šo (pievērst uzmanību porta numuriem - 9*** jāsakrīt ar datora numuru):


``` 
[Unit]
Description=Reverse SSH tunnel to Raspberry Pi
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=ubuntuuser

ExecStart=/usr/bin/sshpass -p 'rpipassword' /usr/bin/ssh -N \
    -R 127.0.0.1:9001:127.0.0.1:5000 \
    -o StrictHostKeyChecking=accept-new \
    -o ServerAliveInterval=30 \
    -o ServerAliveCountMax=3 \
    -o ExitOnForwardFailure=yes \
    raspi@192.168.8.188

Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target


```
nomainīt paroles, piekļuves datus, IP un palaist:
```
sudo systemctl daemon-reload
sudo systemctl enable flask-tunnel-9001
sudo systemctl start flask-tunnel-9001

# pārbaudīt vai status ir aktīvs
sudo systemctl status flask-tunnel-9001

# apskatīt žurnālfailus
journalctl -u flask-tunnel-9001 -f
```


## Kā darbināt flask serveri:

1. pārliecināties, ka eksistē virtual environment un to uzstādīt, ja nepieciešams:

    izveidot jaunu
    ``` python -m venv myfirstproject ```
    
    aktivizēt
    ```source myfirstproject/bin/activate```

    rezultātā komandrindā jāredz kas līdzīgs šim:
    (myfirstproject) ... $
    

2. uz virtual environment uzlikt flask:

``` pip install flask ```

3. iedarbināt flask

``` flask --app app.py run ```

