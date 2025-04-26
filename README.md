# 📨 Routing Simulator in Go

This is a simple Go project that simulates static routing using routing tables. It demonstrates basic networking concepts such as packets, destinations, routing tables, and forwarding logic — all implemented in Go.

---


### Requirements

- Go 1.18 or later

### Running the Program

```bash
go run staticRouter.go
```

## 📦 Features

- Defines a `Packet` structure with source, destination, and data fields.
- Implements a `RoutingTable` as a map from destination IPs to next hop.
- Supports a method to find the next hop using a routing table.
- Simulates packet forwarding through two routers.

---

