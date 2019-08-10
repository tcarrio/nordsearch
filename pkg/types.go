package nordsearch

// Server is the data type returned from the NordVPN /api/servers endpoint
type Server struct {
	ID             uint64        `json:"id"`
	IPAddress      string        `json:"ip_address"`
	SearchKeywords []interface{} `json:"search_keywords"`
	Categories     []Category    `json:"categories"`
	Name           string        `json:"name"`
	Domain         string        `json:"domain"`
	Price          uint64        `json:"price"`
	Flag           string        `json:"flag"`
	Country        string        `json:"country"`
	Location       Location      `json:"location"`
	Load           uint64        `json:"load"`
	Features       Features      `json:"features"`
}

// Category is a metadata field for a Server
type Category struct {
	Name string `json:"name"`
}

// Features describes the features offered by a server
type Features struct {
	Ikev2              bool `json:"ikev2"`
	OpenvpnUDP         bool `json:"openvpn_udp"`
	OpenvpnTCP         bool `json:"openvpn_tcp"`
	Socks              bool `json:"socks"`
	Proxy              bool `json:"proxy"`
	PPTP               bool `json:"pptp"`
	L2TP               bool `json:"l2tp"`
	OpenvpnXorUDP      bool `json:"openvpn_xor_udp"`
	OpenvpnXorTCP      bool `json:"openvpn_xor_tcp"`
	ProxyCybersex      bool `json:"proxy_cybersec"`
	ProxySSL           bool `json:"proxy_ssl"`
	ProxySSLCybersec   bool `json:"proxy_ssl_cybersec"`
	IkeV2v6            bool `json:"ikev2_v6"`
	OpenvpnUDPv6       bool `json:"openvpn_udp_v6"`
	OpenvpnTCPv6       bool `json:"openvpn_tcp_v6"`
	WireguardUDP       bool `json:"wireguard_udp"`
	OpenvpnUDPTLSCrypt bool `json:"openvpn_udp_tls_crypt"`
	OpenvpnTCPTLSCrypt bool `json:"openvpn_tcp_tls_crypt"`
}
