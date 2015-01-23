/* Monitor */
package main 

import (
	"fmt"
	"net/http" 
	"os"
	"time"
	"log"
	"code.google.com/p/gcfg"
	"net/smtp"
	"bytes"
	//"strconv"
)

type Config struct {
    Server struct {
		    Url string
    		IntervalSeconds int 
    }
}

func loadConfiguration(cfgFile string) Config {
    var err error
    var cfg Config

    if cfgFile!= "" {
        err = gcfg.ReadFileInto(&cfg, cfgFile)
    }
    if err != nil {
        fmt.Println(err)
        log.Fatalf("Failed to parse gcfg data: %s", err)
        os.Exit(2)
    } 
    return cfg
}

func sendMail(smtpStr, sender, recipient string) {
	//sender = "dev.service.log@nexusguard.com"
	c, err := smtp.Dial(smtpStr)
        if err != nil {
                log.Fatal(err)
        }
        // Set the sender and recipient.
        c.Mail(sender)
        c.Rcpt(recipient)
        // Send the email body.:q

        wc, err := c.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        buf := bytes.NewBufferString("This is the email body.")
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
}

func monitorServer(url string, seconds int){
	for{
		response, err := http.Get(url) 
		if err != nil {
	        //log.Printf("Failed to get response, %s\n", err)
			writeLogToFile("DIE!")
	     	time.Sleep(time.Duration(seconds) * time.Second)
	     	continue
		}
		if response.Status == "200 OK" { 
			//log.Println("Alive!")
			writeLogToFile("Alive!")
		} else {
			log.Println("response status is: %s", response.Status)
		}
     	time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func writeLogToFile (logMsg string){
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(logMsg)
}


func main() {
	//sendMail("smtp.gmail.com:25","stickbob@gmail.com","stickbob@gmail.com")
	os.Exit(0)
	var cfg Config
	var Url string
	var IntervalSeconds int
	cfgFile := "ini.gcfg"

	cfg = loadConfiguration(cfgFile)
	Url = cfg.Server.Url
	IntervalSeconds = cfg.Server.IntervalSeconds
	log.Printf("Start monitoring: %s", Url)
	//IntervalSeconds = strconv.Atoi(IS)
	//os.Exit(0)

	monitorServer(Url, IntervalSeconds)
	/*go monitorServer(Url, IntervalSeconds)
	for (
		time.Sleep(1* time.Mi)
		)
		*/
}
