package utils

import (
	"net"
	"testing"
)

func TestNtoI(t *testing.T) {
	var ipInt uint32
	ipInt = NtoI(net.ParseIP("0.0.0.0"))
	if ipInt != uint32(0) {
		t.Errorf("Invalid value %d for ip 0.0.0.0, Expected 0", ipInt)
	}
	ipInt = NtoI(net.ParseIP("0.0.0.1"))
	if ipInt != uint32(1) {
		t.Errorf("Invalid value %d for ip 0.0.0.1, Expected 1", ipInt)
	}
	ipInt = NtoI(net.ParseIP("128.0.0.0"))
	if ipInt != uint32(2147483648) {
		t.Errorf("Invalid value %d for ip 128.0.0.0, Expected 2147483648", ipInt)
	}
	ipInt = NtoI(net.ParseIP("255.255.255.255"))
	if ipInt != uint32(4294967295) {
		t.Errorf("Invalid value %d for ip 255.255.255.255, Expected 4294967295", ipInt)
	}
}

func TestItoN(t *testing.T) {
	var ip net.IP
	ip = ItoN(uint32(0))
	if ip.String() != "0.0.0.0" {
		t.Errorf("Invalid ip %s for uint32(0). Expected 0.0.0.0", ip.String())
	}
	ip = ItoN(uint32(1))
	if ip.String() != "0.0.0.1" {
		t.Errorf("Invalid ip %s for uint32(1). Expected 0.0.0.1", ip.String())
	}
	ip = ItoN(uint32(2147483648))
	if ip.String() != "128.0.0.0" {
		t.Errorf("Invalid ip %s for uint32(2147483648). Expected 128.0.0.0", ip.String())
	}
	ip = ItoN(uint32(4294967295))
	if ip.String() != "255.255.255.255" {
		t.Errorf("Invalid ip %s for uint32(4294967295). Expected 255.255.255.255", ip.String())
	}
}
