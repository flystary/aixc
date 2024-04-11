package arp

import (
	"errors"
	"os"
	"strings"
)

// Entry define the list available in /proc/net/arp
type Entry struct {
	IPAddress string
	HWType    string
	Flags     string
	HWAddress string
	Mask      string
	Device    string
}

type Entrys []Entry

// GetMACFromAddr
func (entries Entrys) GetMACFromAddr(ADDR string) (string, error) {
	for _, entry := range entries {
		addr := strings.Trim(entry.IPAddress, " ")
		if addr == ADDR {
			return entry.HWAddress, nil
		}
	}
	return "", errors.New("ADDR Not Fond")
}

// GetEntryFromMAC get an entry by searching with MAC address
func (entries Entrys) GetEntryFromMAC(mac string) (Entry, error) {
	for _, entry := range entries {
		if entry.HWAddress == mac {
			return entry, nil
		}
	}

	return Entry{}, errors.New("MAC address not found")
}

// GetEntries list ARP entries in /proc/net/arp
func GetEntries() (Entrys, error) {
	fileDatas, err := os.ReadFile("/proc/net/arp")
	if err != nil {
		return nil, err
	}

	entries := Entrys{}
	datas := strings.Split(string(fileDatas), "\n")
	for i, data := range datas {
		if i == 0 || data == "" {
			continue
		}
		parsedData := removeWhiteSpace(strings.Split(data, " "))
		entries = append(entries, Entry{
			IPAddress: parsedData[0],
			HWType:    parsedData[1],
			Flags:     parsedData[2],
			HWAddress: parsedData[3],
			Mask:      parsedData[4],
			Device:    parsedData[5],
		})
	}

	return entries, nil
}

func removeWhiteSpace(tab []string) []string {
	var newTab []string
	for _, element := range tab {
		if element == "" {
			continue
		}
		newTab = append(newTab, element)
	}

	return newTab
}
