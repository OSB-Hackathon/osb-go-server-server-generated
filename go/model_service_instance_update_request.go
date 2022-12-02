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

type ServiceInstanceUpdateRequest struct {

	Context *Context `json:"context,omitempty"`

	ServiceId string `json:"service_id"`

	PlanId string `json:"plan_id,omitempty"`

	Parameters interface{} `json:"parameters,omitempty"`

	PreviousValues *ServiceInstancePreviousValues `json:"previous_values,omitempty"`

	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`
}
