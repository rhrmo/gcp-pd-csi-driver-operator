// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IPv4GatewayConfigApplyConfiguration represents a declarative configuration of the IPv4GatewayConfig type for use
// with apply.
type IPv4GatewayConfigApplyConfiguration struct {
	InternalMasqueradeSubnet *string `json:"internalMasqueradeSubnet,omitempty"`
}

// IPv4GatewayConfigApplyConfiguration constructs a declarative configuration of the IPv4GatewayConfig type for use with
// apply.
func IPv4GatewayConfig() *IPv4GatewayConfigApplyConfiguration {
	return &IPv4GatewayConfigApplyConfiguration{}
}

// WithInternalMasqueradeSubnet sets the InternalMasqueradeSubnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalMasqueradeSubnet field is set to the value of the last call.
func (b *IPv4GatewayConfigApplyConfiguration) WithInternalMasqueradeSubnet(value string) *IPv4GatewayConfigApplyConfiguration {
	b.InternalMasqueradeSubnet = &value
	return b
}
