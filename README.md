# exposer
Expose localhost to internet using ngrok and send the ngrok address to your Discord server.

The purpose of this project is to expose a server which does not have a public IP or is in a protected network.
There are other BETTER solutions like [Cloudflare Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/)
and [frp](https://github.com/fatedier/frp) but they need you to have a domain name or public IP server. 

## How to use?
- Make sure you have `Golang` and `Ngrok` installed
- Create `.env` file using `.env.example`
- Run `go build main.go` to build the executable OR directly run using `go run main.go`

To make this run every time after reboot
- Edit `run.sh` and add your correct location 
- Run `chmod +x run.sh` to make it executable
- Set a crobtab for it
```
$ crontab -e
@reboot  /home/user/startup.sh
```
