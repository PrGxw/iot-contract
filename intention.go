package contract

// IntentionKind names something the backend wants a gateway to do. In step 1 an
// intention travels as a command action whose Type is the kind.
type IntentionKind string

const (
	IntentMeshRefresh   IntentionKind = "mesh_refresh"
	IntentMeshAdd       IntentionKind = "mesh_add"
	IntentMeshRemove    IntentionKind = "mesh_remove"
	IntentMeshReset     IntentionKind = "mesh_reset"
	IntentMeshIdleNodes IntentionKind = "mesh_idle_nodes"
	IntentQueryType     IntentionKind = "query_type"
)

// MeshNodesParams are the parameters of mesh_add, mesh_remove and mesh_reset.
type MeshNodesParams struct {
	MACs []string `json:"macs"`
}

// QueryTypeParams are the parameters of query_type.
type QueryTypeParams struct {
	MAC string `json:"mac"`
}
