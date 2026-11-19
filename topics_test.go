package contract

import "testing"

func TestInventoryTopic(t *testing.T) {
	got := InventoryTopic(3, 7, "AA:BB:CC:DD:EE:FF")
	if got != "tenants/3/buildings/7/gateways/AA:BB:CC:DD:EE:FF/inventory" {
		t.Fatal(got)
	}
}

func TestInventoryTopicFilter(t *testing.T) {
	if InventoryTopicFilter != "tenants/+/buildings/+/gateways/+/inventory" {
		t.Fatal(InventoryTopicFilter)
	}
}
