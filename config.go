package main

import "fmt"

func getZone(c *MasterConfig, ip string) string {
	for z, ips := range c.Zones {
		for _, _ip := range ips {
			if ip == _ip {
				return z
			}
		}
	}
	return ""
}

func spitServerConfig(c *MasterConfig, ip string) *PortMap {
	portMap := PortMap{"tcp": []int{}, "udp": []int{}}

	zone := getZone(c, ip)

	for _, destinations := range c.Network {
		for dest, conf := range destinations {
			if dest != zone {
				continue
			}

			for _, p := range conf.AllowTCP {
				portMap["tcp"] = append(portMap["tcp"], p)
			}

			for _, p := range conf.AllowUDP {
				portMap["udp"] = append(portMap["udp"], p)
			}
			for _, p := range conf.ThroughputTest {
				portMap["throughputTest"] = append(portMap["throughputTest"], p)
			}
		}
	}

	return &portMap
}

func spitClientConfig(c *MasterConfig, ip string) map[string]PortMap {
	zone := getZone(c, ip)

	portMap := map[string]PortMap{}

	fmt.Printf("\nZone is %v", zone)
	for target, conf := range c.Network[zone] {
		pm := PortMap{"tcp": conf.AllowTCP, "udp": conf.AllowUDP, "throughputTest": conf.ThroughputTest}

		for _, _ip := range c.Zones[target] {
			if _ip == ip {
				continue
			}

			portMap[_ip] = pm
		}
	}

	return portMap

}
