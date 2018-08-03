package main

import (
	"flag"
	"fmt"
	. "fogflow/common/ngsi"
	"os"
	"os/signal"
	"syscall"
)

func HandleNotifyContext(notifyCtxReq *NotifyContextRequest) {
	fmt.Println("===========RECEIVE NOTIFY CONTEXT=========")
	fmt.Printf("<< %+v >>\r\n", notifyCtxReq)

	return
}

func startAgent(port int) {
	agent := NGSIAgent{Port: port}
	agent.Start()
	agent.SetContextNotifyHandler(HandleNotifyContext)
}

func main() {
	myPort := flag.Int("p", 8090, "the port of this agent")
	flag.Parse()

	startAgent(*myPort)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	<-c

}
