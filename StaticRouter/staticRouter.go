package main

import (
	"fmt"
)

type Packet struct{
	Source string
	Destination string
	Data string
}

type RoutingTable map[string]string

func(p *Packet) String() string{
	return fmt.Sprintf("Packet: Source: %s, Destination: %s, Data: %s \n", p.Source, p.Destination, p.Data)
}

func (rt RoutingTable )GetnextHop(destination string) string{
	nextHop, exists:= rt[destination]
	if exists{
		return nextHop
	}
	return "No Route Found"
}


func ForwardPacket(p *Packet, routingTable RoutingTable){
	nextHop:= routingTable.GetnextHop(p.Destination)
	fmt.Printf("The NextHop is %s\n", nextHop)
	fmt.Println(p)
}


func main(){
	router1RoutingTable:= RoutingTable{
		"192.168.1.2" : "Router2",
	}

	router2RoutingTable:= RoutingTable{
		"192.168.1.2": "Final Destination",
	}

	packet:= &Packet{
		Source: "192.168.1.1",
		Destination: "192.168.1.2",
		Data: "Hello, This is a Packet",
	}

	fmt.Println("Router table 1")
	ForwardPacket(packet, router1RoutingTable)
	fmt.Println("Router table 2")
	ForwardPacket(packet, router2RoutingTable)
}