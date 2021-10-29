// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type GetAllNetworkRouter struct {
	NetworkRouters []GetNetworkRouter `json:"networkRouters"`
}

type GetSpecificRouterResp struct {
	NetworkRouter GetNetworkRouter `json:"networkRouter"`
}

type CreateRouterResp struct {
	Success bool `json:"success"`
	ID      int  `json:"id" tf:"id,computed"`
}

type GetNetworkRouter struct {
	ID              int           `json:"id" tf:"id,computed"`
	Code            string        `json:"code"`
	Name            string        `json:"name"`
	Description     interface{}   `json:"description"`
	Category        string        `json:"category"`
	DateCreated     string        `json:"dateCreated"`
	LastUpdated     string        `json:"lastUpdated"`
	RouterType      string        `json:"routerType"`
	Status          string        `json:"status"`
	Enabled         bool          `json:"enabled"`
	ExternalIP      interface{}   `json:"externalIp"`
	ExternalID      string        `json:"externalId"`
	ProviderID      string        `json:"providerId" tf:"provider_id,computed"`
	Type            IDModel       `json:"type"`
	NetworkServer   IDModel       `json:"networkServer"`
	Zone            IDModel       `json:"zone"`
	Instance        interface{}   `json:"instance"`
	ExternalNetwork interface{}   `json:"externalNetwork"`
	Site            interface{}   `json:"site"`
	Interfaces      []interface{} `json:"interfaces"`
}

type CreateRouterRequest struct {
	NetworkRouter CreateRouterRequestRouter `json:"networkRouter"`
}

type CreateRouterRequestRouter struct {
	ID              int                       `json:"-" tf:"id,computed"`
	Name            string                    `json:"name" tf:"name"`
	Type            IDModel                   `json:"type,omitempty"`
	TypeID          int                       `json:"-" tf:"type_id,computed"`
	Enabled         bool                      `json:"enabled" tf:"enable"`
	Site            IDStringModel             `json:"site,omitempty"`
	GroupID         string                    `json:"-" tf:"group_id"`
	NetworkServer   IDModel                   `json:"networkServer,omitempty"`
	NetworkServerID int                       `json:"-" tf:"network_server_id,computed"`
	EnableBGP       bool                      `json:"enableBgp"`
	Config          CreateRouterRequestConfig `json:"config"`

	// for tftags parsing
	TfTier0Config *CreateRouterTier0Config `json:"-" tf:"tier0_config,sub"`
	TfTier1Config *CreateRouterTier1Config `json:"-" tf:"tier1_config,sub"`
}

type CreateRouterRequestConfig struct {
	EdgeCluster   string `json:"edgeCluster,omitempty"`
	HaMode        string `json:"haMode,omitempty"`
	FailOver      string `json:"failOver,omitempty"`
	Tier0Gateways string `json:"tier0Gateway,omitempty"`
	CreateRouterTier0Config
}

type CreateRouterTier0Config struct {
	Bgp
	RouteRedistributionTier0
	RouteRedistributionTier1
	// For tftag parsing
	TfEdgeCluster string                   `json:"-" tf:"edge_cluster"`
	TfHaMode      string                   `json:"-" tf:"ha_mode"`
	TfFailOver    string                   `json:"-" tf:"fail_over"`
	TfBGP         Bgp                      `json:"-" tf:"bgp,sub"`
	TfRRTier0     RouteRedistributionTier0 `json:"-" tf:"route_redistribution_tier0,sub"`
	TfRRTier1     RouteRedistributionTier1 `json:"-" tf:"route_redistribution_tier1,sub"`
}

type CreateRouterTier1Config struct {
	TfEdgeCluster        string             `json:"-" tf:"edge_cluster"`
	TfTier0Gateways      string             `json:"-" tf:"tier0_gateway,omitempty"`
	TfRouteAdvertisement RouteAdvertisement `json:"-" tf:"route_advertisement,sub"`
}

type RouteRedistributionTier0 struct {
	TIER0STATIC            bool `json:"TIER0_STATIC" tf:"tier0_static"`
	TIER0NAT               bool `json:"TIER0_NAT" tf:"tier0_nat"`
	TIER0IPSECLOCALIP      bool `json:"TIER0_IPSEC_LOCAL_IP" tf:"tier0_ipsec_local_ip"`
	TIER0DNSFORWARDERIP    bool `json:"TIER0_DNS_FORWARDER_IP" tf:"tier0_dns_forwarder_ip"`
	TIER0SERVICEINTERFACE  bool `json:"TIER0_SERVICE_INTERFACE" tf:"tier0_service_interface"`
	TIER0EXTERNALINTERFACE bool `json:"TIER0_EXTERNAL_INTERFACE" tf:"tier0_external_interface"`
	TIER0LOOPBACKINTERFACE bool `json:"TIER0_LOOPBACK_INTERFACE" tf:"tier0_loopback_interface"`
	TIER0SEGMENT           bool `json:"TIER0_SEGMENT" tf:"tier0_segment"`
}

type RouteRedistributionTier1 struct {
	TIER1SERVICEINTERFACE bool `json:"TIER1_SERVICE_INTERFACE," tf:"tier1_service_interface"`
	TIER1SEGMENT          bool `json:"TIER1_SEGMENT" tf:"tier1_segment"`
	RouteAdvertisement
	Tier1Connected          bool `json:"-" tf:"tier1_connected"`
	Tier1StaticRoutes       bool `json:"-" tf:"tier1_static_routes"`
	TIER1DNSFORWARDERIP     bool `json:"-" tf:"tier1_dns_forwarder_ip"`
	TIER1STATIC             bool `json:"-" tf:"tier1_static"`
	TIER1LBVIP              bool `json:"-" tf:"tier1_lb_vip"`
	TIER1NAT                bool `json:"-" tf:"tier1_nat"`
	TIER1LBSNAT             bool `json:"-" tf:"tier1_lb_snat"`
	TIER1IPSECLOCALENDPOINT bool `json:"-" tf:"tier1_ipsec_local_endpoint"`
}

type RouteAdvertisement struct {
	Tier1Connected          bool `json:"TIER1_CONNECTED" tf:"tier1_connected"`
	Tier1StaticRoutes       bool `json:"TIER1_STATIC_ROUTES" tf:"tier1_static_routes"`
	TIER1DNSFORWARDERIP     bool `json:"TIER1_DNS_FORWARDER_IP" tf:"tier1_dns_forwarder_ip"`
	TIER1STATIC             bool `json:"TIER1_STATIC" tf:"tier1_static"`
	TIER1LBVIP              bool `json:"TIER1_LB_VIP" tf:"tier1_lb_vip"`
	TIER1NAT                bool `json:"TIER1_NAT" tf:"tier1_nat"`
	TIER1LBSNAT             bool `json:"TIER1_LB_SNAT" tf:"tier1_lb_snat"`
	TIER1IPSECLOCALENDPOINT bool `json:"TIER1_IPSEC_LOCAL_ENDPOINT" tf:"tier1_ipsec_local_endpoint"`
}

type Bgp struct {
	LOCALASNUM     int    `json:"LOCAL_AS_NUM,omitempty" tf:"local_as_num"`
	ECMP           bool   `json:"ECMP" tf:"ecmp"`
	MULTIPATHRELAX bool   `json:"MULTIPATH_RELAX" tf:"multipath_relax"`
	INTERSRIBGP    bool   `json:"INTER_SR_IBGP" tf:"inter_sr_ibgp"`
	RESTARTMODE    string `json:"RESTART_MODE,omitempty" tf:"restart_mode"`
	RESTARTTIME    int    `json:"RESTART_TIME,omitempty" tf:"restart_time"`
	STALEROUTETIME int    `json:"STALE_ROUTE_TIME,omitempty" tf:"stale_route_time"`
	TfEnableBgp    bool   `json:"-" tf:"enable_bgp"`
}

type GetNetworlRouterTypes struct {
	NetworkRouterTypes []NetworkRouterTypes `json:"networkRouterTypes"`
}

type NetworkRouterTypes struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type GetNetworkServicesResp struct {
	NetworkServices []GetNetworkServices `json:"networkServices"`
}

type GetNetworkServices struct {
	ServiceType               string      `json:"serviceType"`
	ServiceTypeName           string      `json:"serviceTypeName"`
	Type                      string      `json:"type"`
	TypeName                  string      `json:"typeName"`
	Name                      string      `json:"name"`
	ID                        int         `json:"id"`
	IntegrationID             int         `json:"integrationId"`
	CanEdit                   bool        `json:"canEdit"`
	CanDelete                 bool        `json:"canDelete"`
	Status                    string      `json:"status"`
	LastUpdated               string      `json:"lastUpdated"`
	BrandingImageName         interface{} `json:"brandingImageName"`
	SupportsTenantPermissions bool        `json:"supportsTenantPermissions"`
}

type CreateRouterNatRequest struct {
	CreateRouterNat CreateRouterNat `json:"networkRouterNAT"`
}

type CreateRouterNat struct {
	ID                 int                   `json:"-" tf:"id,computed"`
	Name               string                `json:"name,omitempty" tf:"name"`
	Description        string                `json:"description,omitempty" tf:"description"`
	Enabled            bool                  `json:"enabled" tf:"enabled"`
	Config             CreateRouterNatConfig `json:"config" tf:"config,sub"`
	SourceNetwork      string                `json:"sourceNetwork,omitempty" tf:"source_network"`
	DestinationNetwork string                `json:"destinationNetwork,omitempty" tf:"destination_network"`
	TranslatedNetwork  string                `json:"translatedNetwork,omitempty" tf:"translated_network"`
	TranslatedPorts    string                `json:"translatedPorts,omitempty" tf:"translated_ports"`
	Priority           int                   `json:"priority" tf:"priority"`
	RouterID           int                   `json:"-" tf:"router_id"`
}

type TfCreateRouterNat struct {
	NatRule []CreateRouterNat `tf:"nat_rule"`
}

type CreateRouterNatConfig struct {
	Action   string `json:"action" tf:"action"`
	Service  string `json:"service" tf:"service"`
	Firewall string `json:"firewall" tf:"firewall"`
	Scope    string `json:"scope" tf:"scope"`
	Logging  bool   `json:"logging" tf:"logging"`
}

type CreateRouterNatResponse struct {
	IDModel
	SuccessOrErrorMessage
}

type GetSpecificRouterNatResponse struct {
	GetSpecificRouterNat GetSpecificRouterNat `json:"networkRouterNAT"`
}

type GetSpecificRouterNat struct {
	ID                 int    `json:"id" tf:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Enabled            bool   `json:"enabled"`
	SourceNetwork      string `json:"sourceNetwork"`
	DestinationNetwork string `json:"destinationNetwork"`
	TranslatedNetwork  string `json:"translatedNetwork"`
	SourcePorts        string `json:"sourcePorts"`
	DestinationPorts   string `json:"destinationPorts"`
	TranslatedPorts    string `json:"translatedPorts"`
	Priority           int    `json:"priority"`
	SyncSource         string `json:"syncSource"`
	InternalID         string `json:"internalId"`
	ExternalID         string `json:"externalId"`
	ProviderID         string `json:"providerId"`
	DateCreated        string `json:"dateCreated"`
	LastUpdated        string `json:"lastUpdated"`
	IsDeprecated       bool   `json:"-" tf:"is_deprecated,computed"`
}
