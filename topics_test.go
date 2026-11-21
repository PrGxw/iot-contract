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

func TestConfigTopics(t *testing.T) {
	if got := ConfigTopic(1, 2, "AA:BB"); got != "tenants/1/buildings/2/gateways/AA:BB/config" {
		t.Fatal(got)
	}
	if ConfigTopicFilter != "tenants/+/buildings/+/gateways/+/config" {
		t.Fatal(ConfigTopicFilter)
	}
	if got := ConfigReportTopic(1, 2, "AA:BB"); got != "tenants/1/buildings/2/gateways/AA:BB/config_report" {
		t.Fatal(got)
	}
	if ConfigReportTopicFilter != "tenants/+/buildings/+/gateways/+/config_report" {
		t.Fatal(ConfigReportTopicFilter)
	}
}
