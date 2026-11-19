package contract

// ParentGateway is the parent of every node directly under the mesh root. The
// root itself is never a device in the inventory.
const ParentGateway = "gateway"

// InventoryDevice is one device on a gateway's list. An empty Type means the
// gateway does not know the device's type.
type InventoryDevice struct {
	MAC  string     `json:"mac"`
	Type DeviceType `json:"type,omitempty"`
}

// MeshLink is a device's last known parent: another device's MAC or ParentGateway.
type MeshLink struct {
	MAC    string `json:"mac"`
	Parent string `json:"parent"`
}

// Inventory is a full snapshot of a gateway's active devices. It carries
// existence and structure only, never presence or status.
type Inventory struct {
	Devices []InventoryDevice `json:"devices"`
	Links   []MeshLink        `json:"links"`
}
