//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type IDNameModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListNetworksBody struct {
	Networks     []GetSpecificNetworkBody `json:"networks"`
	NetworkCount int                      `json:"networkCount"`
}

type GetSpecificNetworkBody struct {
	ID                      int           `json:"id" tf:"id"`
	Name                    string        `json:"name" tf:"name"`
	Zone                    IDNameModel   `json:"zone"`
	DisplayName             string        `json:"displayName" tf:"display_name,computed"`
	Type                    IDModel       `json:"type"`
	TypeID                  int           `json:"-" tf:"type_id,computed"`
	Owner                   IDNameModel   `json:"owner"`
	Code                    string        `json:"code" tf:"code,computed"`
	Category                string        `json:"category"`
	ExternalID              string        `json:"externalId" tf:"external_id,computed"`
	InternalID              string        `json:"internalId" tf:"internal_id,computed"`
	UniqueID                string        `json:"uniqueId" tf:"unique_id,computed"`
	ExternalType            string        `json:"externalType"`
	RefType                 string        `json:"refType"`
	RefID                   int           `json:"refId"`
	DhcpServer              bool          `json:"dhcpServer" tf:"dhcp_server"`
	Status                  string        `json:"status" tf:"status,computed"`
	Visibility              string        `json:"visibility" tf:"visibility"`
	EnableAdmin             bool          `json:"enableAdmin"`
	ScanNetwork             bool          `json:"scanNetwork" tf:"scan_network"`
	Active                  bool          `json:"active" tf:"active"`
	DefaultNetwork          bool          `json:"defaultNetwork"`
	AssignPublicIP          bool          `json:"assignPublicIp"`
	ApplianceURLProxyBypass bool          `json:"applianceUrlProxyBypass" tf:"appliance_url_proxy_bypass"`
	ZonePool                IDNameModel   `json:"zonePool"`
	AllowStaticOverride     bool          `json:"allowStaticOverride"`
	Tenants                 []IDNameModel `json:"tenants"`
}

type CreateNetworkResourcePermission struct {
	All   bool      `json:"all"`
	Sites []IDModel `json:"sites"`
}

type CreateNetworkRequest struct {
	Network             CreateNetwork                   `json:"network"`
	ResourcePermissions CreateNetworkResourcePermission `json:"resourcePermissions"`
}

type CreateNetwork struct {
	Name                    string              `json:"name" tf:"name"`
	Description             string              `json:"description" tf:"description"`
	CloudID                 int                 `json:"-" tf:"cloud_id"`
	Zone                    IDModel             `json:"zone" tf:"zone"`
	Type                    IDModel             `json:"type" tf:"type"`
	Cidr                    string              `json:"cidr" tf:"cidr"`
	Gateway                 string              `json:"gateway" tf:"gateway"`
	DNSPrimary              string              `json:"dnsPrimary" tf:"primary_dns"`
	DNSSecondary            string              `json:"dnsSecondary" tf:"secondary_dns"`
	Config                  CreateNetworkConfig `json:"config" tf:"config,sub"`
	DhcpServer              string              `json:"dhcpServer" tf:"dhcp_server"`
	Pool                    int                 `json:"pool" tf:"pool"`
	ScanNetwork             string              `json:"scanNetwork" tf:"scan_network"`
	ApplianceURLProxyBypass string              `json:"applianceUrlProxyBypass" tf:"appliance_url_proxy_bypass"`
	NoProxy                 string              `json:"noProxy" tf:"no_proxy"`
	ScopeID                 string              `json:"scopeId" tf:"scode_id"`
}

type CreateNetworkConfig struct {
	ConnectedGateway string `json:"connectedGateway" tf:"connected_gateway"`
	VlanIDs          string `json:"vlanIDs" tf:"vlan_ids"`
}
