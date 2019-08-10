package main

import (
	"fmt"

	nordsearch "github.com/tcarrio/nordsearch/pkg"
)

func germanWireguardFilter(s nordsearch.Server) bool {
	return s.Features.WireguardUDP && s.Flag == "DE"
}

func main() {
	fmt.Println("Contacting NordVPN...")

	servers, err := nordsearch.GetServers()
	if err != nil {
		fmt.Printf("Failed to get servers: %s\n", err.Error())
	}

	fmt.Printf("Found %d servers\n", len(servers))

	sorted := servers.Filter(germanWireguardFilter).Sort(nordsearch.LoadSorter)[:10]

	for i, server := range sorted {
		fmt.Printf("#%d Wireguard server: %s (Load: %d)\n", i, server.Name, server.Load)
	}
}
