package contract

import (
	"encoding/json"
	"testing"
)

func TestEnvelope_NewerThan(t *testing.T) {
	e := Envelope{BootID: "b2", Seq: 1}
	if !e.NewerThan("b1", 900) {
		t.Error("a new boot must win regardless of seq")
	}
	if e.NewerThan("b2", 1) {
		t.Error("same seq is not newer")
	}
	if e.NewerThan("b2", 5) {
		t.Error("lower seq is not newer")
	}
	if !(Envelope{BootID: "b2", Seq: 6}).NewerThan("b2", 5) {
		t.Error("higher seq on the same boot is newer")
	}
}

func TestWrap_RoundTrip(t *testing.T) {
	inv := Inventory{
		Devices: []InventoryDevice{{MAC: "AA:BB:CC:DD:EE:01", Type: TypeLprIl}, {MAC: "AA:BB:CC:DD:EE:02"}},
		Links:   []MeshLink{{MAC: "AA:BB:CC:DD:EE:01", Parent: ParentGateway}},
	}
	env, err := Wrap("AA:BB:CC:00:00:01", "boot-1", 3, KindInventory, inv)
	if err != nil {
		t.Fatal(err)
	}
	if env.Kind != KindInventory || env.Seq != 3 || env.BootID != "boot-1" || env.GatewayID != "AA:BB:CC:00:00:01" {
		t.Fatalf("envelope fields: %+v", env)
	}
	raw, err := json.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	var back Envelope
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatal(err)
	}
	var got Inventory
	if err := json.Unmarshal(back.Body, &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Devices) != 2 || got.Devices[0].Type != TypeLprIl || got.Devices[1].Type != "" || got.Links[0].Parent != ParentGateway {
		t.Fatalf("round trip lost data: %+v", got)
	}
}

func TestInventoryDevice_JSONOmitsUnknownType(t *testing.T) {
	raw, _ := json.Marshal(InventoryDevice{MAC: "AA:BB:CC:DD:EE:02"})
	if string(raw) != `{"mac":"AA:BB:CC:DD:EE:02"}` {
		t.Fatalf("got %s", raw)
	}
}
