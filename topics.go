package contract

import "fmt"

// InventoryTopicFilter matches every gateway's inventory topic.
const InventoryTopicFilter = "tenants/+/buildings/+/gateways/+/inventory"

// InventoryTopic is the topic a gateway publishes its inventory on.
func InventoryTopic(tenantID, buildingID uint, gatewayMAC string) string {
	return fmt.Sprintf("tenants/%d/buildings/%d/gateways/%s/inventory", tenantID, buildingID, gatewayMAC)
}

// ConfigTopicFilter matches every gateway's config topic.
const ConfigTopicFilter = "tenants/+/buildings/+/gateways/+/config"

// ConfigTopic is the topic the backend publishes a gateway's config documents on.
func ConfigTopic(tenantID, buildingID uint, gatewayMAC string) string {
	return fmt.Sprintf("tenants/%d/buildings/%d/gateways/%s/config", tenantID, buildingID, gatewayMAC)
}

// ConfigReportTopicFilter matches every gateway's config report topic.
const ConfigReportTopicFilter = "tenants/+/buildings/+/gateways/+/config_report"

// ConfigReportTopic is the topic a gateway publishes its config report on.
func ConfigReportTopic(tenantID, buildingID uint, gatewayMAC string) string {
	return fmt.Sprintf("tenants/%d/buildings/%d/gateways/%s/config_report", tenantID, buildingID, gatewayMAC)
}
