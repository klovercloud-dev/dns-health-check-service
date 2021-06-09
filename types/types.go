package types
import "net"

type Zone struct {
	Name      string	`json:"name"`
	Locations map[string]struct{}	`json:"locations"`
}

type Record struct {
	A     *New_A_Record `json:"a,omitempty"`
	AAAA  *New_AAAA_Record `json:"aaaa,omitempty"`
	TXT   *New_TXT_Record `json:"txt,omitempty"`
	CNAME *New_CNAME_Record `json:"cname,omitempty"`
	NS    []NS_Record `json:"ns,omitempty"`
	MX    *New_MX_Record `json:"mx,omitempty"`
	SRV   *New_SRV_Record `json:"srv,omitempty"`
	CAA   *New_CAA_Record `json:"caa,omitempty"`
	SOA   *New_SOA_Record `json:"soa,omitempty"`
}

// A Record
type A_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

type New_A_Record struct {
	Type  string            `json:"type"`
	Value FailOver_A_Record `json:"value"`
}

type Simple_A_Record struct {
	Values []A_Record `json:"values"`
}

type FailOver_A_Record struct {
	Primary FailOver_A_Data `json:"primary"`
	Secondary FailOver_A_Data `json:"secondary"`
}

type FailOver_A_Data struct {
	Data              []A_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}

// AAAA Record
type AAAA_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

type New_AAAA_Record struct {
	Type  string            `json:"type"`
	Value FailOver_A_Record `json:"value"`
}

type Simple_AAAA_Record struct {
	Values []AAAA_Record `json:"values"`
}

type FailOver_AAAA_Record struct {
	Primary   FailOver_AAA_Data `json:"primary"`
	Secondary FailOver_AAA_Data `json:"secondary"`
}

type FailOver_AAA_Data struct {
	Data              []AAAA_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}


//TXT Record
type TXT_Record struct {
	Ttl  uint32 `json:"ttl,omitempty"`
	Text string `json:"text"`
}


type New_TXT_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_TXT_Record `json:"value,omitempty"`
}

type Simple_TXT_Record struct {
	Values []TXT_Record `json:"values"`
}

type FailOver_TXT_Record struct {
	Primary   FailOver_TXT_Data `json:"primary"`
	Secondary FailOver_TXT_Data `json:"secondary"`
}

type FailOver_TXT_Data struct {
	Data              []TXT_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}



// CNAME Record
type CNAME_Record struct {
	Ttl  uint32 `json:"ttl,omitempty"`
	Host string `json:"host"`
}

type New_CNAME_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_CNAME_Record `json:"value,omitempty"`
}

type Simple_CNAME_Record struct {
	Values []CNAME_Record `json:"values"`
}

type FailOver_CNAME_Record struct {
	Primary   FailOver_CNAME_Data `json:"primary"`
	Secondary FailOver_CNAME_Data `json:"secondary"`
}

type FailOver_CNAME_Data struct {
	Data              []CNAME_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}


// NS Record
type NS_Record struct {
	Ttl  uint32 `json:"ttl,omitempty"`
	Host string `json:"host"`
}


// MX Record
type MX_Record struct {
	Ttl        uint32 `json:"ttl,omitempty"`
	Host       string `json:"host"`
	Preference uint16 `json:"preference"`
}

type New_MX_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_MX_Record `json:"value,omitempty"`
}

type Simple_MX_Record struct {
	Values []MX_Record `json:"values"`
}

type FailOver_MX_Record struct {
	Primary   FailOver_MX_Data `json:"primary"`
	Secondary FailOver_MX_Data `json:"secondary"`
}

type FailOver_MX_Data struct {
	Data              []MX_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}


// SRV Record
type SRV_Record struct {
	Ttl      uint32 `json:"ttl,omitempty"`
	Priority uint16 `json:"priority"`
	Weight   uint16 `json:"weight"`
	Port     uint16 `json:"port"`
	Target   string `json:"target"`
}

type New_SRV_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_SRV_Record `json:"value,omitempty"`
}

type Simple_SRV_Record struct {
	Values []SRV_Record `json:"values"`
}

type FailOver_SRV_Record struct {
	Primary   FailOver_SRV_Data `json:"primary"`
	Secondary FailOver_SRV_Data `json:"secondary"`
}

type FailOver_SRV_Data struct {
	Data              []SRV_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}


// SOA Record
type SOA_Record struct {
	Ttl     uint32 `json:"ttl,omitempty"`
	Ns      string `json:"ns"`
	MBox    string `json:"MBox"`
	Refresh uint32 `json:"refresh"`
	Retry   uint32 `json:"retry"`
	Expire  uint32 `json:"expire"`
	MinTtl  uint32 `json:"minttl"`
}


type New_SOA_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_SOA_Record `json:"value,omitempty"`
}

type Simple_SOA_Record struct {
	Values []SOA_Record `json:"values"`
}

type FailOver_SOA_Record struct {
	Primary   FailOver_SOA_Data `json:"primary"`
	Secondary FailOver_SOA_Data `json:"secondary"`
}

type FailOver_SOA_Data struct {
	Data              []SOA_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}



// CAA Record
type CAA_Record struct {
	Flag  uint8 `json:"flag"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type New_CAA_Record struct {
	Type  string            `json:"type,omitempty"`
	Value FailOver_CAA_Record `json:"value,omitempty"`
}

type Simple_CAA_Record struct {
	Values []CAA_Record `json:"values"`
}

type FailOver_CAA_Record struct {
	Primary   FailOver_CAA_Data `json:"primary"`
	Secondary FailOver_CAA_Data `json:"secondary"`
}

type FailOver_CAA_Data struct {
	Data              []CAA_Record `json:"data"`
	IsHealthy         bool `json:"isHealthy"`
	HealthCheckConfig FailOverHealthCheckConfig `json:"healthCheckConfig"`
}


type FailOverHealthCheckConfig struct {
	Type	string `json:"type,omitempty"`
	TargetIPs	[]string `json:"targetIPs,omitempty"`
	Port	string `json:"port,omitempty"`
	TargetUrl 	string `json:"targetUrl,omitempty"`
}