package contract

// DeviceType is a device type in the platform's vocabulary.
type DeviceType string

const (
	TypeLprIl          DeviceType = "lpr_il"
	TypePSIL           DeviceType = "psil"
	TypeDisplay        DeviceType = "display"
	TypeLight          DeviceType = "light"
	TypePSLS           DeviceType = "psls"
	TypeRelayNode      DeviceType = "relay_node"
	TypeBarrier        DeviceType = "barrier"
	TypeParkingBarrier DeviceType = "parking_barrier"
)

// Role is a contracted device role.
type Role string

const (
	RoleRadio          Role = "radio"
	RoleCamera         Role = "camera"
	RolePSIL           Role = "psil"
	RoleDisplay        Role = "display"
	RoleLighting       Role = "lighting"
	RolePSLS           Role = "psls"
	RoleRelay          Role = "relay"
	RoleBarrier        Role = "barrier"
	RoleParkingBarrier Role = "parking_barrier"
)

var typeRoles = []struct {
	t     DeviceType
	roles []Role
}{
	{TypeLprIl, []Role{RoleCamera, RolePSIL}},
	{TypePSIL, []Role{RolePSIL}},
	{TypeDisplay, []Role{RoleDisplay}},
	{TypeLight, []Role{RoleLighting}},
	{TypePSLS, []Role{RolePSLS}},
	{TypeRelayNode, []Role{RoleRelay}},
	{TypeBarrier, []Role{RoleBarrier}},
	{TypeParkingBarrier, []Role{RoleParkingBarrier}},
}

// DeviceTypes returns every contracted device type in a stable order.
func DeviceTypes() []DeviceType {
	out := make([]DeviceType, len(typeRoles))
	for i, tr := range typeRoles {
		out[i] = tr.t
	}
	return out
}

// RolesOf returns the roles a device type is contracted to have. Radio is
// implicit on every device and is never listed. The result is a copy.
func RolesOf(t DeviceType) ([]Role, bool) {
	for _, tr := range typeRoles {
		if tr.t == t {
			return append([]Role(nil), tr.roles...), true
		}
	}
	return nil, false
}
