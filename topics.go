package contract

import "fmt"

// InventoryTopicFilter matches every gateway's inventory topic.
const InventoryTopicFilter = "tenants/+/buildings/+/gateways/+/inventory"

// InventoryTopic is the topic a gateway publishes its inventory on.
func InventoryTopic(tenantID, buildingID uint, gatewayMAC string) string {
	return fmt.Sprintf("tenants/%d/buildings/%d/gateways/%s/inventory", tenantID, buildingID, gatewayMAC)
}
