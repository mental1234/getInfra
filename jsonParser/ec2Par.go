package jsonParser

import "encoding/json"

func UnmarshalWel(data []byte) (Wel, error) {
	var r Wel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Wel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Wel struct {
	Reservations []Reservation `json:"Reservations"`
}

type Reservation struct {
	Instances     []Instance `json:"Instances"`    
	OwnerID       string     `json:"OwnerId"`      
	ReservationID string     `json:"ReservationId"`
}

type Instance struct {
	AmiLaunchIndex                   int64                            `json:"AmiLaunchIndex"`                  
	Architecture                     string                           `json:"Architecture"`                    
	BlockDeviceMappings              []BlockDeviceMapping             `json:"BlockDeviceMappings"`             
	CapacityReservationSpecification CapacityReservationSpecification `json:"CapacityReservationSpecification"`
	ClientToken                      string                           `json:"ClientToken"`                     
	CPUOptions                       CPUOptions                       `json:"CpuOptions"`                      
	EbsOptimized                     bool                             `json:"EbsOptimized"`                    
	EnaSupport                       bool                             `json:"EnaSupport"`                      
	HibernationOptions               HibernationOptions               `json:"HibernationOptions"`              
	Hypervisor                       string                           `json:"Hypervisor"`                      
	ImageID                          string                           `json:"ImageId"`                         
	InstanceID                       string                           `json:"InstanceId"`                      
	InstanceType                     string                           `json:"InstanceType"`                    
	KeyName                          string                           `json:"KeyName"`                         
	LaunchTime                       string                           `json:"LaunchTime"`                      
	Monitoring                       Monitoring                       `json:"Monitoring"`                      
	NetworkInterfaces                []NetworkInterface               `json:"NetworkInterfaces"`               
	Placement                        Placement                        `json:"Placement"`                       
	PrivateDNSName                   string                           `json:"PrivateDnsName"`                  
	PrivateIPAddress                 string                           `json:"PrivateIpAddress"`                
	PublicDNSName                    string                           `json:"PublicDnsName"`                   
	RootDeviceName                   string                           `json:"RootDeviceName"`                  
	RootDeviceType                   string                           `json:"RootDeviceType"`                  
	SecurityGroups                   []Group                          `json:"SecurityGroups"`                  
	SourceDestCheck                  bool                             `json:"SourceDestCheck"`                 
	State                            State                            `json:"State"`                           
	StateReason                      StateReason                      `json:"StateReason"`                     
	StateTransitionReason            string                           `json:"StateTransitionReason"`           
	SubnetID                         string                           `json:"SubnetId"`                        
	Tags                             []Tag                            `json:"Tags"`                            
	VirtualizationType               string                           `json:"VirtualizationType"`              
	VpcID                            string                           `json:"VpcId"`                           
}

type BlockDeviceMapping struct {
	DeviceName string `json:"DeviceName"`
	Ebs        Ebs    `json:"Ebs"`       
}

type Ebs struct {
	AttachTime          string `json:"AttachTime"`         
	DeleteOnTermination bool   `json:"DeleteOnTermination"`
	Status              string `json:"Status"`             
	VolumeID            string `json:"VolumeId"`           
}

type CPUOptions struct {
	CoreCount      int64 `json:"CoreCount"`     
	ThreadsPerCore int64 `json:"ThreadsPerCore"`
}

type CapacityReservationSpecification struct {
	CapacityReservationPreference string `json:"CapacityReservationPreference"`
}

type HibernationOptions struct {
	Configured bool `json:"Configured"`
}

type Monitoring struct {
	State string `json:"State"`
}

type NetworkInterface struct {
	Attachment         Attachment         `json:"Attachment"`        
	Description        string             `json:"Description"`       
	Groups             []Group            `json:"Groups"`            
	MACAddress         string             `json:"MacAddress"`        
	NetworkInterfaceID string             `json:"NetworkInterfaceId"`
	OwnerID            string             `json:"OwnerId"`           
	PrivateDNSName     string             `json:"PrivateDnsName"`    
	PrivateIPAddress   string             `json:"PrivateIpAddress"`  
	PrivateIPAddresses []PrivateIPAddress `json:"PrivateIpAddresses"`
	SourceDestCheck    bool               `json:"SourceDestCheck"`   
	Status             string             `json:"Status"`            
	SubnetID           string             `json:"SubnetId"`          
	VpcID              string             `json:"VpcId"`             
}

type Attachment struct {
	AttachTime          string `json:"AttachTime"`         
	AttachmentID        string `json:"AttachmentId"`       
	DeleteOnTermination bool   `json:"DeleteOnTermination"`
	DeviceIndex         int64  `json:"DeviceIndex"`        
	Status              string `json:"Status"`             
}

type Group struct {
	GroupID   string `json:"GroupId"`  
	GroupName string `json:"GroupName"`
}

type PrivateIPAddress struct {
	Primary          bool   `json:"Primary"`         
	PrivateDNSName   string `json:"PrivateDnsName"`  
	PrivateIPAddress string `json:"PrivateIpAddress"`
}

type Placement struct {
	AvailabilityZone string `json:"AvailabilityZone"`
	GroupName        string `json:"GroupName"`       
	Tenancy          string `json:"Tenancy"`         
}

type State struct {
	Code int64  `json:"Code"`
	Name string `json:"Name"`
}

type StateReason struct {
	Code    string `json:"Code"`   
	Message string `json:"Message"`
}

type Tag struct {
	Key   string `json:"Key"`  
	Value string `json:"Value"`
}

