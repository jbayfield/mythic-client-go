package mythic

// VPS -
type VPS struct {
	Name        string             `json:"name"`
	Identifier  string             `json:"identifier"`
	Product     string             `json:"product"`
	Dormant     bool               `json:"dormant"`
}
