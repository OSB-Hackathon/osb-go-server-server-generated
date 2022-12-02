/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Plan struct {

	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Free bool `json:"free,omitempty"`

	Bindable bool `json:"bindable,omitempty"`

	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	Schemas *SchemasObject `json:"schemas,omitempty"`

	MaximumPollingDuration int32 `json:"maximum_polling_duration,omitempty"`

	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`

	BindingRotatable bool `json:"binding_rotatable,omitempty"`
}
