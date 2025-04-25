package main

import (
	"fmt"
)

type Packet struct{
	Source string
	Destination string
	Data string
}

type routingTable map[string]string

func(p *Packet) String() string{
	return fmt.Sprintf("Packet:{Source: %s, Destination: %s, Data %s}", p.Source, p.Destination, p.Data)
}

func main(){
	router1RoutingTable:= routingTable{
		"192.168.1.1" : "Router2",
	}

	router2RoutingTable:= routingTable{
		"192.168.1.2": "Final Destination",
	}

	packet:= Packet{
		Source: "192.168.0.0",
		Destination: "192.168.1.2",
		Data: "Hello, This is a Packet",
	}

	fmt.Println("Router table 1")
	ForwardPacket(packet, router1RoutingTable)
	fmt.Println("Router table 2")
	ForwardPacket(packet, router2RoutingTable)
}