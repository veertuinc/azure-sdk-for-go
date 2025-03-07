//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

// ARecord - An A record.
type ARecord struct {
	// The IPv4 address of this A record.
	IPv4Address *string `json:"ipv4Address,omitempty"`
}

// AaaaRecord - An AAAA record.
type AaaaRecord struct {
	// The IPv6 address of this AAAA record.
	IPv6Address *string `json:"ipv6Address,omitempty"`
}

// CloudError - An error response from the service.
type CloudError struct {
	// Cloud error body.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response from the service.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// CnameRecord - A CNAME record.
type CnameRecord struct {
	// The canonical name for this CNAME record.
	Cname *string `json:"cname,omitempty"`
}

// MxRecord - An MX record.
type MxRecord struct {
	// The domain name of the mail host for this MX record.
	Exchange *string `json:"exchange,omitempty"`

	// The preference value for this MX record.
	Preference *int32 `json:"preference,omitempty"`
}

// PrivateZone - Describes a Private DNS zone.
type PrivateZone struct {
	// The ETag of the zone.
	Etag *string `json:"etag,omitempty"`

	// The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the Private DNS zone.
	Properties *PrivateZoneProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateZoneListResult - The response to a Private DNS zone list operation.
type PrivateZoneListResult struct {
	// Information about the Private DNS zones.
	Value []*PrivateZone `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// PrivateZoneProperties - Represents the properties of the Private DNS zone.
type PrivateZoneProperties struct {
	// READ-ONLY; Private zone internal Id
	InternalID *string `json:"internalId,omitempty" azure:"ro"`

	// READ-ONLY; The maximum number of record sets that can be created in this Private DNS zone. This is a read-only property
	// and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty" azure:"ro"`

	// READ-ONLY; The maximum number of virtual networks that can be linked to this Private DNS zone. This is a read-only property
	// and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinks *int64 `json:"maxNumberOfVirtualNetworkLinks,omitempty" azure:"ro"`

	// READ-ONLY; The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
	// This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinksWithRegistration *int64 `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" azure:"ro"`

	// READ-ONLY; The current number of record sets in this Private DNS zone. This is a read-only property and any attempt to
	// set this value will be ignored.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty" azure:"ro"`

	// READ-ONLY; The current number of virtual networks that are linked to this Private DNS zone. This is a read-only property
	// and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinks *int64 `json:"numberOfVirtualNetworkLinks,omitempty" azure:"ro"`

	// READ-ONLY; The current number of virtual networks that are linked to this Private DNS zone with registration enabled. This
	// is a read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinksWithRegistration *int64 `json:"numberOfVirtualNetworkLinksWithRegistration,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the resource. This is a read-only property and any attempt to set this value will
	// be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateZonesClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateZonesClient.BeginCreateOrUpdate
// method.
type PrivateZonesClientBeginCreateOrUpdateOptions struct {
	// The ETag of the Private DNS zone. Omit this value to always overwrite the current zone. Specify the last-seen ETag value
	// to prevent accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new Private DNS zone to be created, but to prevent updating an existing zone. Other values will be
	// ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateZonesClientBeginDeleteOptions contains the optional parameters for the PrivateZonesClient.BeginDelete method.
type PrivateZonesClientBeginDeleteOptions struct {
	// The ETag of the Private DNS zone. Omit this value to always delete the current zone. Specify the last-seen ETag value to
	// prevent accidentally deleting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateZonesClientBeginUpdateOptions contains the optional parameters for the PrivateZonesClient.BeginUpdate method.
type PrivateZonesClientBeginUpdateOptions struct {
	// The ETag of the Private DNS zone. Omit this value to always overwrite the current zone. Specify the last-seen ETag value
	// to prevent accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateZonesClientGetOptions contains the optional parameters for the PrivateZonesClient.Get method.
type PrivateZonesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateZonesClientListByResourceGroupOptions contains the optional parameters for the PrivateZonesClient.ListByResourceGroup
// method.
type PrivateZonesClientListByResourceGroupOptions struct {
	// The maximum number of record sets to return. If not specified, returns up to 100 record sets.
	Top *int32
}

// PrivateZonesClientListOptions contains the optional parameters for the PrivateZonesClient.List method.
type PrivateZonesClientListOptions struct {
	// The maximum number of Private DNS zones to return. If not specified, returns up to 100 zones.
	Top *int32
}

// ProxyResource - The resource model definition for an ARM proxy resource.
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PtrRecord - A PTR record.
type PtrRecord struct {
	// The PTR target domain name for this PTR record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// RecordSet - Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
type RecordSet struct {
	// The ETag of the record set.
	Etag *string `json:"etag,omitempty"`

	// The properties of the record set.
	Properties *RecordSetProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RecordSetListResult - The response to a record set list operation.
type RecordSetListResult struct {
	// Information about the record sets in the response.
	Value []*RecordSet `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// RecordSetProperties - Represents the properties of the records in the record set.
type RecordSetProperties struct {
	// The list of A records in the record set.
	ARecords []*ARecord `json:"aRecords,omitempty"`

	// The list of AAAA records in the record set.
	AaaaRecords []*AaaaRecord `json:"aaaaRecords,omitempty"`

	// The CNAME record in the record set.
	CnameRecord *CnameRecord `json:"cnameRecord,omitempty"`

	// The metadata attached to the record set.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// The list of MX records in the record set.
	MxRecords []*MxRecord `json:"mxRecords,omitempty"`

	// The list of PTR records in the record set.
	PtrRecords []*PtrRecord `json:"ptrRecords,omitempty"`

	// The SOA record in the record set.
	SoaRecord *SoaRecord `json:"soaRecord,omitempty"`

	// The list of SRV records in the record set.
	SrvRecords []*SrvRecord `json:"srvRecords,omitempty"`

	// The TTL (time-to-live) of the records in the record set.
	TTL *int64 `json:"ttl,omitempty"`

	// The list of TXT records in the record set.
	TxtRecords []*TxtRecord `json:"txtRecords,omitempty"`

	// READ-ONLY; Fully qualified domain name of the record set.
	Fqdn *string `json:"fqdn,omitempty" azure:"ro"`

	// READ-ONLY; Is the record set auto-registered in the Private DNS zone through a virtual network link?
	IsAutoRegistered *bool `json:"isAutoRegistered,omitempty" azure:"ro"`
}

// RecordSetsClientCreateOrUpdateOptions contains the optional parameters for the RecordSetsClient.CreateOrUpdate method.
type RecordSetsClientCreateOrUpdateOptions struct {
	// The ETag of the record set. Omit this value to always overwrite the current record set. Specify the last-seen ETag value
	// to prevent accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new record set to be created, but to prevent updating an existing record set. Other values will be
	// ignored.
	IfNoneMatch *string
}

// RecordSetsClientDeleteOptions contains the optional parameters for the RecordSetsClient.Delete method.
type RecordSetsClientDeleteOptions struct {
	// The ETag of the record set. Omit this value to always delete the current record set. Specify the last-seen ETag value to
	// prevent accidentally deleting any concurrent changes.
	IfMatch *string
}

// RecordSetsClientGetOptions contains the optional parameters for the RecordSetsClient.Get method.
type RecordSetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RecordSetsClientListByTypeOptions contains the optional parameters for the RecordSetsClient.ListByType method.
type RecordSetsClientListByTypeOptions struct {
	// The suffix label of the record set name to be used to filter the record set enumeration. If this parameter is specified,
	// the returned enumeration will only contain records that end with ".".
	Recordsetnamesuffix *string
	// The maximum number of record sets to return. If not specified, returns up to 100 record sets.
	Top *int32
}

// RecordSetsClientListOptions contains the optional parameters for the RecordSetsClient.List method.
type RecordSetsClientListOptions struct {
	// The suffix label of the record set name to be used to filter the record set enumeration. If this parameter is specified,
	// the returned enumeration will only contain records that end with ".".
	Recordsetnamesuffix *string
	// The maximum number of record sets to return. If not specified, returns up to 100 record sets.
	Top *int32
}

// RecordSetsClientUpdateOptions contains the optional parameters for the RecordSetsClient.Update method.
type RecordSetsClientUpdateOptions struct {
	// The ETag of the record set. Omit this value to always overwrite the current record set. Specify the last-seen ETag value
	// to prevent accidentally overwriting concurrent changes.
	IfMatch *string
}

// Resource - The core properties of ARM resources
type Resource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SoaRecord - An SOA record.
type SoaRecord struct {
	// The email contact for this SOA record.
	Email *string `json:"email,omitempty"`

	// The expire time for this SOA record.
	ExpireTime *int64 `json:"expireTime,omitempty"`

	// The domain name of the authoritative name server for this SOA record.
	Host *string `json:"host,omitempty"`

	// The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTTL *int64 `json:"minimumTtl,omitempty"`

	// The refresh value for this SOA record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`

	// The retry time for this SOA record.
	RetryTime *int64 `json:"retryTime,omitempty"`

	// The serial number for this SOA record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
}

// SrvRecord - An SRV record.
type SrvRecord struct {
	// The port value for this SRV record.
	Port *int32 `json:"port,omitempty"`

	// The priority value for this SRV record.
	Priority *int32 `json:"priority,omitempty"`

	// The target domain name for this SRV record.
	Target *string `json:"target,omitempty"`

	// The weight value for this SRV record.
	Weight *int32 `json:"weight,omitempty"`
}

// SubResource - Reference to another subresource.
type SubResource struct {
	// Resource ID.
	ID *string `json:"id,omitempty"`
}

// TrackedResource - The resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TxtRecord - A TXT record.
type TxtRecord struct {
	// The text value of this TXT record.
	Value []*string `json:"value,omitempty"`
}

// VirtualNetworkLink - Describes a link to virtual network for a Private DNS zone.
type VirtualNetworkLink struct {
	// The ETag of the virtual network link.
	Etag *string `json:"etag,omitempty"`

	// The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the virtual network link to the Private DNS zone.
	Properties *VirtualNetworkLinkProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// VirtualNetworkLinkListResult - The response to a list virtual network link to Private DNS zone operation.
type VirtualNetworkLinkListResult struct {
	// Information about the virtual network links to the Private DNS zones.
	Value []*VirtualNetworkLink `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// VirtualNetworkLinkProperties - Represents the properties of the Private DNS zone.
type VirtualNetworkLinkProperties struct {
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled?
	RegistrationEnabled *bool `json:"registrationEnabled,omitempty"`

	// The reference of the virtual network.
	VirtualNetwork *SubResource `json:"virtualNetwork,omitempty"`

	// READ-ONLY; The provisioning state of the resource. This is a read-only property and any attempt to set this value will
	// be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The status of the virtual network link to the Private DNS zone. Possible values are 'InProgress' and 'Done'.
	// This is a read-only property and any attempt to set this value will be ignored.
	VirtualNetworkLinkState *VirtualNetworkLinkState `json:"virtualNetworkLinkState,omitempty" azure:"ro"`
}

// VirtualNetworkLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginCreateOrUpdate
// method.
type VirtualNetworkLinksClientBeginCreateOrUpdateOptions struct {
	// The ETag of the virtual network link to the Private DNS zone. Omit this value to always overwrite the current virtual network
	// link. Specify the last-seen ETag value to prevent accidentally overwriting
	// any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new virtual network link to the Private DNS zone to be created, but to prevent updating an existing
	// link. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginDelete
// method.
type VirtualNetworkLinksClientBeginDeleteOptions struct {
	// The ETag of the virtual network link to the Private DNS zone. Omit this value to always delete the current zone. Specify
	// the last-seen ETag value to prevent accidentally deleting any concurrent
	// changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientBeginUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginUpdate
// method.
type VirtualNetworkLinksClientBeginUpdateOptions struct {
	// The ETag of the virtual network link to the Private DNS zone. Omit this value to always overwrite the current virtual network
	// link. Specify the last-seen ETag value to prevent accidentally overwriting
	// any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientGetOptions contains the optional parameters for the VirtualNetworkLinksClient.Get method.
type VirtualNetworkLinksClientGetOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkLinksClientListOptions contains the optional parameters for the VirtualNetworkLinksClient.List method.
type VirtualNetworkLinksClientListOptions struct {
	// The maximum number of virtual network links to return. If not specified, returns up to 100 virtual network links.
	Top *int32
}
