package main

import (
	"os"
	"net"
	"github.com/trasgum/syslog_event_parser/mesos_agent"
)
//
//type Executors struct {
//	id string
//	name string
//	directory string
//}
//
//type Framework struct {
//	id string
//	name string
//	tasks []Executors
//}
//
//type Agent struct {
//	id string
//	IPAddr net.TCPAddr
//	frameworks []Framework
//}

// Get ip of this machine to connect to mesos
func GetOutboundIP() (string, error) {
	conn, err := net.Dial("tcp", "leader.mesos:5050")
    	if err != nil {
        	return "", err
    	}
    	defer conn.Close()

    	localAddr := conn.LocalAddr().(*net.TCPAddr)
    	return localAddr.IP.String(), nil
}


func onInit(me *mesos_agent.MesosAgent) error{
	/*
	Do everything that is needed to initialize processing (e.g.
	open files, create handles, connect to systems...)
	*/
	if AgentIPAddr, err := GetOutboundIP(); err == nil {
		err :=  me.Poll(AgentIPAddr)
		if err == nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func onReceive(message []byte, len_message int) error{
	out := os.Stdout
	defer out.Close()
	if _, err := out.Write(message[0 : len_message+1]); err != nil {
		return err
	}
	out.Sync()
	return nil

}

func onExit() {
	/*
	Do everything that is needed to finish processing (e.g.
	close files, handles, disconnect from systems...). This is
	being called immediately before exiting.
	*/
}

func main() {
	const BUF_SIZE  = 256
	in := os.Stdin
	//out := os.Stdout
	bufOut := make([]byte, BUF_SIZE)
	n:= 0

	me := mesos_agent.MesosAgent{}
	if err := onInit(&me); err != nil {
		os.Exit(1)
	}

	for {
		if _, err := in.Read(bufOut[n : n+1]); err != nil {
			break
		}
		if bufOut[n] == 0x0a || n == BUF_SIZE {
			//if _, err := out.Write(bufOut[0 : n+1]); err != nil {
			if err := onReceive(bufOut[0 : n+1], n); err != nil {
				break
			}
			n = 0
		} else {
			n++
		}
	}
	//out.Close()
	in.Close()
}
