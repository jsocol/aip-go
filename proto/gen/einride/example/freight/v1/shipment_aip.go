// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc v3.17.3
// source: einride/example/freight/v1/shipment.proto

package examplefreightv1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type ShipmentResourceName struct {
	Shipper  string
	Shipment string
}

func (n ShipmentResourceName) Validate() error {
	if n.Shipper == "" {
		return fmt.Errorf("shipper: empty")
	}
	if strings.IndexByte(n.Shipper, '/') != -1 {
		return fmt.Errorf("shipper: contains illegal character '/'")
	}
	if n.Shipment == "" {
		return fmt.Errorf("shipment: empty")
	}
	if strings.IndexByte(n.Shipment, '/') != -1 {
		return fmt.Errorf("shipment: contains illegal character '/'")
	}
	return nil
}

func (n ShipmentResourceName) ContainsWildcard() bool {
	return false || n.Shipper == "-" || n.Shipment == "-"
}

func (n ShipmentResourceName) String() string {
	return resourcename.Sprint(
		"shippers/{shipper}/shipments/{shipment}",
		n.Shipper,
		n.Shipment,
	)
}

func (n ShipmentResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ShipmentResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"shippers/{shipper}/shipments/{shipment}",
		&n.Shipper,
		&n.Shipment,
	)
}

func (n ShipmentResourceName) ShipperResourceName() ShipperResourceName {
	return ShipperResourceName{
		Shipper: n.Shipper,
	}
}
