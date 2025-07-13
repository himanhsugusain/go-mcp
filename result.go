package server

const protocolVersion = "2025-06-18"

type InitResponse struct {
	ProtocolVersion string `json:"protocolVersion"`
	Capabilities Capabilities `json:"capabilities"`
	ServerInfo ServerInfo `json:"serverInfo"`
	Instructions string `json:"instructions"`
}

type ServerInfo struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Version string  `json:"version"`
}
