package nordsearch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// // Servers provides encapsulating type for server list
// type Servers []Server

// type serverFilter func(Server) bool

// // Filter provides a way to perform a filter across the full list of servers and return a subset
// func (s Servers) Filter(fn serverFilter) (servers Servers) {
// 	servers = make([]Server, len(s))
// 	for _, server := range s {
// 		if fn(server) {
// 			servers = append(servers, server)
// 		}
// 	}
// 	return
// }

// GetServers returns a list of servers from the NordVPN API
func GetServers() (Servers, error) {
	var servers Servers

	resp, err := http.Get("https://nordvpn.com/api/server")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
