package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"encoding/json"

	"github.com/gregdel/pushover"
	log "github.com/sirupsen/logrus"
)

var (
	configfilename string
	pocreds        pushovercredentials
)

type pushovercredentials struct {
	Token     string
	Recipient string
}

func init() {
	// read commandline options
	flag.StringVar(&configfilename, "config", "startup-pushover.json", "Filename of config file")
	flag.Parse()

	// parse config file
	file, err := os.Open(configfilename)
	if err != nil {
		log.Fatalf("Unable to open config file: %s", err)
	}
	defer file.Close()

	jdec := json.NewDecoder(file)
	err = jdec.Decode(&pocreds)
	if err != nil {
		log.Fatalf("Unable to parse config file: %s", err)
	}

}

func main() {
	var (
		subject, messagebody bytes.Buffer
		interfaces           []net.Interface
		hostname             string
		err                  error
	)

	// get hostname
	hostname, err = os.Hostname()
	if err != nil {
		log.Println(err)
		hostname = "unknown"
	}
	fmt.Fprintf(&subject, "Host %s is online", hostname)

	// enumerate interfaces
	interfaces, err = net.Interfaces()
	if err != nil {
		log.Println(err)
		interfaces = []net.Interface{}
	}
	for _, i := range interfaces {
		addrs, _ := i.Addrs()
		if i.Flags&net.FlagLoopback == 0 && i.Flags&net.FlagUp != 0 && !strings.HasPrefix(i.Name, "docker") && !strings.HasPrefix(i.Name, "virbr") && len(addrs) > 0 {
			fmt.Fprintf(&messagebody, "Interface <b>%s</b>: ", i.Name)
			for _, a := range addrs {
				stripped := strings.Split(a.String(), "/")[0]
				fmt.Fprintf(&messagebody, "<a href=\"ssh://%s\">%s</a>\n", stripped, stripped)
			}
		}
	}

	// prepare pushover message
	message := &pushover.Message{
		Message:   messagebody.String(),
		Title:     subject.String(),
		Priority:  pushover.PriorityNormal,
		Timestamp: time.Now().Unix(),
		Retry:     60 * time.Second,
		Expire:    time.Hour,
		Sound:     pushover.SoundMagic,
		HTML:      true,
	}

	// create pushover sender
	app := pushover.New(pocreds.Token)

	// set pushover recipient
	recipient := pushover.NewRecipient(pocreds.Recipient)

	// send pushover message
	resp, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Println(err)
	}
	if resp.Status != 1 {
		log.Println("Error sending notification: ", resp)
	} else {
		log.Println("Successfully sent notification")
	}
}
