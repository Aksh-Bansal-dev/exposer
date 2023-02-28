package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
    port := getEnv("PORT","")
    if len(port)==0{
        log.Fatal("Cannot find PORT env")
    }
    fmt.Println("Tunnel port:",port)
	cmd := exec.Command("ngrok", "tcp", port, "--log=stdout")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 4 && s[:3] == "url" {
            fmt.Println("Sending discord message...")
            err=sendDiscordMsg(s[4:])
            if err!=nil{
                log.Fatal(err)
            }
		}
	}
}

type DiscordMsg struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func sendDiscordMsg(content string) error{
	msg := DiscordMsg{Content: content, Username: getEnv("SERVER_NAME", "default")}
	payload, err := json.Marshal(msg)
	if err != nil {
		return errors.New("Error marshalling content")
	}
    discordWebhookUrl := getEnv("DISCORD_WEBHOOK_URL","")
    if len(discordWebhookUrl)==0{
        return errors.New("Cannot find DISCORD_WEBHOOK_URL env")
    }
	req, err := http.NewRequest("POST", discordWebhookUrl, bytes.NewBuffer(payload))
	if err != nil {
        return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
    return nil
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
