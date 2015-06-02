package bridge

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/docker/libnetwork/driverapi"
	"github.com/docker/libnetwork/ipallocator"
	"github.com/docker/libnetwork/netlabel"
	"github.com/docker/libnetwork/netutils"
	"github.com/docker/libnetwork/options"
	"github.com/docker/libnetwork/portmapper"
	"github.com/docker/libnetwork/sandbox"
	"github.com/docker/libnetwork/types"
	"github.com/vishvananda/netlink"
)

const (
	networkType = "bridge"
	vethPrefix = "veth"
	vethLen = 7
	containerVethPrefix = "eth"
	maxAllocatePortAttempts = 10
	ifaceID = 1
)

var (
	ipAllocator *ipallocator.IPAllocator
	portMapper  *portmapper.PortMapper
)

type configuration struct{
	EnableIPForwarding bool
}

type networkConfiguration struct {
	BridgeName string
	AddressIPv4 net.IPNet
	FixedCIDR net.IPNet
	EnableIPTables        bool
	EnableIPMasquerade    bool
	EnableICC             bool
	Mtu                   int
	DefaultGatewayIPv4    net.IP	\
	DefaultBindingIP net.IP
	AllowDefaultBridge bool
	EnableUserLandProxy bool 
}

type endpointConfiguration struct {
		
}