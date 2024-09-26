// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func IploadbalancingUdpFarmServerResourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"address": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Required:   true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Description:         "IPv4 address (e.g., 192.0.2.0)",
			MarkdownDescription: "IPv4 address (e.g., 192.0.2.0)",
		},
		"backend_id": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Computed:            true,
			Description:         "Synonym for farm_id",
			MarkdownDescription: "Synonym for farm_id",
		},
		"display_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Optional:            true,
			Description:         "Human readable name for your server, this field is for you",
			MarkdownDescription: "Human readable name for your server, this field is for you",
		},
		"farm_id": schema.Int64Attribute{
			CustomType: ovhtypes.TfInt64Type{},
			Required:   true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplace(),
			},
			Description:         "Id of your farm",
			MarkdownDescription: "Id of your farm",
		},
		"port": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Optional:            true,
			Description:         "Port attached to your server ([1..49151]). Inherited from farm if null",
			MarkdownDescription: "Port attached to your server ([1..49151]). Inherited from farm if null",
		},
		"server_id": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Computed:            true,
			Description:         "Id of your server",
			MarkdownDescription: "Id of your server",
		},
		"service_name": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Required:   true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Description:         "The internal name of your IP load balancing",
			MarkdownDescription: "The internal name of your IP load balancing",
		},
		"status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Possible values for server status",
			MarkdownDescription: "Possible values for server status",
			Validators: []validator.String{
				stringvalidator.OneOf(
					"active",
					"inactive",
				),
			},
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type IploadbalancingUdpFarmServerModel struct {
	Address     ovhtypes.TfStringValue `tfsdk:"address" json:"address"`
	BackendId   ovhtypes.TfInt64Value  `tfsdk:"backend_id" json:"backendId"`
	DisplayName ovhtypes.TfStringValue `tfsdk:"display_name" json:"displayName"`
	FarmId      ovhtypes.TfInt64Value  `tfsdk:"farm_id" json:"farmId"`
	Port        ovhtypes.TfInt64Value  `tfsdk:"port" json:"port"`
	ServerId    ovhtypes.TfInt64Value  `tfsdk:"server_id" json:"serverId"`
	ServiceName ovhtypes.TfStringValue `tfsdk:"service_name" json:"serviceName"`
	Status      ovhtypes.TfStringValue `tfsdk:"status" json:"status"`
}

func (v *IploadbalancingUdpFarmServerModel) MergeWith(other *IploadbalancingUdpFarmServerModel) {

	if (v.Address.IsUnknown() || v.Address.IsNull()) && !other.Address.IsUnknown() {
		v.Address = other.Address
	}

	if (v.BackendId.IsUnknown() || v.BackendId.IsNull()) && !other.BackendId.IsUnknown() {
		v.BackendId = other.BackendId
	}

	if (v.DisplayName.IsUnknown() || v.DisplayName.IsNull()) && !other.DisplayName.IsUnknown() {
		v.DisplayName = other.DisplayName
	}

	if (v.FarmId.IsUnknown() || v.FarmId.IsNull()) && !other.FarmId.IsUnknown() {
		v.FarmId = other.FarmId
	}

	if (v.Port.IsUnknown() || v.Port.IsNull()) && !other.Port.IsUnknown() {
		v.Port = other.Port
	}

	if (v.ServerId.IsUnknown() || v.ServerId.IsNull()) && !other.ServerId.IsUnknown() {
		v.ServerId = other.ServerId
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

}

type IploadbalancingUdpFarmServerWritableModel struct {
	Address     *ovhtypes.TfStringValue `tfsdk:"address" json:"address,omitempty"`
	DisplayName *ovhtypes.TfStringValue `tfsdk:"display_name" json:"displayName,omitempty"`
	Port        *ovhtypes.TfInt64Value  `tfsdk:"port" json:"port,omitempty"`
	Status      *ovhtypes.TfStringValue `tfsdk:"status" json:"status,omitempty"`
}

func (v IploadbalancingUdpFarmServerModel) ToCreate() *IploadbalancingUdpFarmServerWritableModel {
	res := &IploadbalancingUdpFarmServerWritableModel{}

	if !v.Address.IsUnknown() {
		res.Address = &v.Address
	}

	if !v.DisplayName.IsUnknown() {
		res.DisplayName = &v.DisplayName
	}

	if !v.Port.IsUnknown() {
		res.Port = &v.Port
	}

	if !v.Status.IsUnknown() {
		res.Status = &v.Status
	}

	return res
}

func (v IploadbalancingUdpFarmServerModel) ToUpdate() *IploadbalancingUdpFarmServerWritableModel {
	res := &IploadbalancingUdpFarmServerWritableModel{}

	if !v.DisplayName.IsUnknown() {
		res.DisplayName = &v.DisplayName
	}

	if !v.Port.IsUnknown() {
		res.Port = &v.Port
	}

	if !v.Status.IsUnknown() {
		res.Status = &v.Status
	}

	return res
}
