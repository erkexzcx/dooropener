# dooropener

My very own door opener, using simple servo (taken from Arduino projects), glued to the domophone. :)

Controlled via Telegram bot command or simple HTTP request to secret endpoint. Under the hood it uses Pi-Blaster to set PWM (for servo angle control).

<img src="https://raw.githubusercontent.com/erkexzcx/dooropener/master/gallery/pic1.jpg" alt="pic1" width="450" height="600"/>

<img src="https://raw.githubusercontent.com/erkexzcx/dooropener/master/gallery/pic2.jpg" alt="pic2" width="450" height="600"/>

<img src="https://raw.githubusercontent.com/erkexzcx/dooropener/master/gallery/pic3.jpg" alt="pic3" width="450" height="600"/>

<img src="https://raw.githubusercontent.com/erkexzcx/dooropener/master/gallery/pic4.jpg" alt="pic4" width="450" height="600"/>

# Usage

1. Create Telegram bot: https://core.telegram.org/bots

Set below commands for your bot (using BotFather):
```
open - Open door
```

2. Install [Golang](https://golang.org/doc/install).

3. Install [Pi-blaster](https://github.com/sarfata/pi-blaster).

4. Install `dooropener` service:
```
cd ~
git clone https://github.com/erkexzcx/dooropener
cd dooropener
go build -ldflags="-s -w" -o dooropener ./cmd/dooropener/dooropener.go

cp dooropener.example.yml dooropener.yml
vim dooropener.yml

sudo cp dooropener.service /etc/systemd/system/
sudo vim /etc/systemd/system/dooropener.service

systemctl start dooropener.service
systemctl enable dooropener.service
```

# Upgrade

Upgrade `dooropener` using below commands:
```
cd ~/dooropener
git pull
go build -ldflags="-s -w" -o dooropener ./cmd/dooropener/dooropener.go
systemctl restart dooropener.service
```
