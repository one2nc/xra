package main

import (
	"encoding/json"
	"strings"
)

func spitServerConfig(c *MasterConfig, ip string) []byte {
	// figure out the zone
	var zones []string

	for z, ips := range c.Zones {
		for _, _ip := range ips {
			if ip == _ip {
				zones = append(zones, z)
			}
		}
	}

	portMap := make(map[string]map[int64]bool)

	for _, sn := range c.Network {
		for n, rules := range sn {
			for _, z := range zones {
				if n != z {
					continue
				}

				for r, ports := range rules {
					parts := strings.Split(r, "_")
					for _, p := range ports {
						if portMap[parts[len(parts)-1]] == nil {
							portMap[parts[len(parts)-1]] = make(map[int64]bool)
						}
						portMap[parts[len(parts)-1]][p] = true
					}
				}
			}
		}
	}

	finalPortMap := make(map[string][]int64)

	for proto, pm := range portMap {
		for p := range pm {
			finalPortMap[proto] = append(finalPortMap[proto], p)
		}
	}

	fpmBytes, err := json.MarshalIndent(&finalPortMap, "", "    ")
	if err != nil {
		panic(err)
	}

	return fpmBytes
}

func spitClientConfig(c *MasterConfig, ip string) []byte {
	var zones []string

	for z, ips := range c.Zones {
		for _, _ip := range ips {
			if ip == _ip {
				zones = append(zones, z)
			}
		}
	}

	portMap := make(map[string]map[string]map[int64]bool)

	for n, sn := range c.Network {
		for _, z := range zones {
			if n != z {
				continue
			}

			for tz, rules := range sn {
				for _, _ip := range c.Zones[tz] {
					if _ip == ip {
						continue
					}
					if portMap[_ip] == nil {
						portMap[_ip] = make(map[string]map[int64]bool)
					}

					for r, ports := range rules {
						parts := strings.Split(r, "_")
						for _, p := range ports {
							if portMap[_ip][parts[len(parts)-1]] == nil {
								portMap[_ip][parts[len(parts)-1]] = make(map[int64]bool)
							}
							portMap[_ip][parts[len(parts)-1]][p] = true
						}
					}
				}
			}
		}
	}

	finalPortMap := make(map[string]map[string][]int64)

	for tIp, rules := range portMap {
		for r, ports := range rules {
			for p := range ports {
				if finalPortMap[tIp] == nil {
					finalPortMap[tIp] = make(map[string][]int64)
				}

				finalPortMap[tIp][r] = append(finalPortMap[tIp][r], p)
			}
		}
	}

	fpmBytes, err := json.MarshalIndent(&finalPortMap, "", "    ")
	if err != nil {
		panic(err)
	}

	return fpmBytes
}
