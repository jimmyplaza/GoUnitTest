package main

func main() {
	SmtpServer := "smtp.gmail.com"
	Port := "465"
	From := "g2.service@nexusguard.com"
	To := make([]string, 1, 1)
	To[0] = "jimmy.ko@nexusguard.com"
	Title := "Go Birdy!"
	Body := "test!!!"

	MorningMail(SmtpServer, Port, From, To, Title, Body)
}
