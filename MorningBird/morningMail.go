package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func MorningMail(SmtpServer, Port, From string, Too []string, Title, BodyMsg string) {
	var To string
	for _, t := range Too {
		To = To + t + " , "
	}
	To = To[:len(To)-3]
	uurl := "http://gcptools.nexusguard.com:445/morningbird"
	var myClient = &http.Client{
		Transport: &http.Transport{
			//Dial: timeoutDialer(time.Duration(10)*time.Second,
			//      time.Duration(10)*time.Second),
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	Title = Title
	BodyMsg = Title + "<br>" + BodyMsg

	v := url.Values{}
	v.Set("to", To)
	v.Set("from", "g2.service@nexusguard.com")
	v.Set("subject", Title)
	v.Set("content", BodyMsg)
	v.Set("publickey", "cba2eb")
	v.Set("privatekey", "c3e12e")

	//out, _ := json.Marshal(m)
	//outReader := bytes.NewReader([]byte(out))
	//res, err := myClient.Post(uurl, "application/x-www-form-urlencoded", outReader)
	//res, err := myClient.Post(url, "application/json", outReader)
	//res, err := myClient.PostForm(uurl, url.Values{ "from" : { "g2.service@nexusguard.com" }, "to" : { "jimmy.ko@nexusguard.com, stickbob@gmail.com"  }, "subject" : {"aaaa"}, "content":{"ttttt"}, "publickey":{"cba2eb"}, "privatekey":{"c3e12e"}  })
	res, err := myClient.PostForm(uurl, v)
	if err != nil {
		fmt.Printf("MorningBird Mail Error:%s\n", err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Printf("MorningBird Mail Error code: %d,url:%s\n", res.StatusCode, uurl)
	}
	//err = json.Unmarshal([]byte(res.Body), &obj)
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Body Error:%s\n", err)
		//errMsg := fmt.Sprintf("%s",err)
		//WriteToSyslog(3,"Monitor-MorningMail",errMsg)
		res.Body.Close()
	}
	var obj interface{}
	err = json.Unmarshal(contents, &obj)
	fmt.Println(obj)
	if err != nil {
		//errMsg := fmt.Sprintf("%s",err)
		fmt.Printf("MorningMail JSON Error:%s => %s\n", uurl, err)
		//WriteToSyslog(3,"Monitor-MorningMail",errMsg)
	}
	res.Body.Close()
}
