package nordsearch

// Servers provides encapsulating type for server list
type Servers []Server

type serverFilter func(Server) bool

type serverSorter func(Server, Server) bool

// Filter provides a way to perform a filter across the full list of servers and return a subset
func (s Servers) Filter(fn serverFilter) (servers Servers) {
	servers = make([]Server, 0)
	for _, server := range s {
		if fn(server) {
			servers = append(servers, server)
		}
	}
	return
}

// Sort will sort the list of servers based on the function provided
func (s Servers) Sort(fn serverSorter) (servers Servers) {
	servers = make(Servers, len(s))
	copy(servers, s)

	for i := 1; i < len(servers); i++ {
		for j := 0; j < len(servers); j++ {
			if fn(servers[i], servers[j]) {
				servers[i], servers[j] = servers[j], servers[i]
			}
		}
	}
	return
}

// LoadSorter implements a serverSorter based on load
func LoadSorter(a Server, b Server) bool {
	return a.Load < b.Load
}
