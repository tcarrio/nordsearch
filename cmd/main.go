package main

import (
	"fmt"

	nordsearch "github.com/tcarrio/nordsearch/pkg"
)

func main() {
	fmt.Println("Contacting NordVPN...")

	servers, err := nordsearch.GetServers()
	if err != nil {
		fmt.Printf("Failed to get servers: %s\n", err.Error())
	}

	fmt.Printf("Found %d servers\n", len(servers))

	wireguards := servers.Filter(func(s nordsearch.Server) bool {
		return s.Features.WireguardUDP && s.Flag == "DE"
	})

	// for _, server := range wireguards {
	// 	fmt.Printf("Found Wireguard server: %s\n", server.Name)
	// }

	sorted := wireguards.Sort(nordsearch.LoadSorter)

	for i, server := range sorted[:10] {
		fmt.Printf("#%d Wireguard server: %s (Load: %d)\n", i, server.Name, server.Load)
	}
}
