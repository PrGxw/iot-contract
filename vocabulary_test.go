package contract

import (
	"reflect"
	"testing"
)

func TestRolesOf_LprIl(t *testing.T) {
	roles, ok := RolesOf(TypeLprIl)
	if !ok || !reflect.DeepEqual(roles, []Role{RoleCamera, RolePSIL}) {
		t.Fatalf("got %v, %v", roles, ok)
	}
}

func TestRolesOf_UnknownType(t *testing.T) {
	if roles, ok := RolesOf("gateway"); ok || roles != nil {
		t.Fatalf("got %v, %v", roles, ok)
	}
}

func TestRolesOf_NeverListsRadio(t *testing.T) {
	for _, ty := range DeviceTypes() {
		roles, ok := RolesOf(ty)
		if !ok || len(roles) == 0 {
			t.Errorf("%s: no roles", ty)
		}
		for _, r := range roles {
			if r == RoleRadio {
				t.Errorf("%s lists radio, which is implicit", ty)
			}
		}
	}
}

func TestDeviceTypes_EightInStableOrder(t *testing.T) {
	want := []DeviceType{TypeLprIl, TypePSIL, TypeDisplay, TypeLight, TypePSLS, TypeRelayNode, TypeBarrier, TypeParkingBarrier}
	if got := DeviceTypes(); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v", got)
	}
}

func TestRolesOf_ReturnsCopy(t *testing.T) {
	r, _ := RolesOf(TypeLprIl)
	r[0] = "mutated"
	if again, _ := RolesOf(TypeLprIl); again[0] != RoleCamera {
		t.Fatal("RolesOf must not expose internal state")
	}
}
