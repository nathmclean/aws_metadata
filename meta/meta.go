package meta

import (
	"log"
)

type Metadata struct {
	amiID string
	amiLaunchIndex string
	amiManifestPath string
	blockDeviceMapping struct {
		ami []string
	}
	hostname string
	instanceAction string
	instanceId string
	instanceType string
	localHostname string
	localIPv4 string
	mac string
	metrics struct{
		vhostmd string
	}
	network struct{
		interfaces struct{
			macs []struct{

			}
		}
	}
}

func Load()  {
	log.Println("Logging test")
}