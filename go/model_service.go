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

type Service struct {

	Name string `json:"name"`

	Id string `json:"id"`

	Description string `json:"description"`

	Tags []string `json:"tags,omitempty"`

	Requires []string `json:"requires,omitempty"`

	Bindable bool `json:"bindable"`

	InstancesRetrievable bool `json:"instances_retrievable,omitempty"`

	BindingsRetrievable bool `json:"bindings_retrievable,omitempty"`

	AllowContextUpdates bool `json:"allow_context_updates,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`

	BindingRotatable bool `json:"binding_rotatable,omitempty"`

	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	Plans []Plan `json:"plans"`
}
