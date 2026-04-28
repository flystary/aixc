package arp

import (
	"bufio"
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

type Entries []Entry

func GetEntries() (Entries, error) {
	f, err := os.Open("/proc/net/arp")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries Entries
	scanner := bufio.NewScanner(f)

	line := 0
	for scanner.Scan() {
		line++
		if line == 1 {
			continue // header
		}

		fields := strings.Fields(scanner.Text())
		if len(fields) < 6 {
			continue // skip invalid line safely
		}

		entries = append(entries, Entry{
			IPAddress: fields[0],
			HWType:    fields[1],
			Flags:     fields[2],
			HWAddress: fields[3],
			Mask:      fields[4],
			Device:    fields[5],
		})
	}

	return entries, scanner.Err()
}

func (e Entries) GetMACFromAddr(addr string) (string, error) {
	for _, entry := range e {
		if entry.IPAddress == addr {
			return entry.HWAddress, nil
		}
	}
	return "", errors.New("address not found")
}

func (e Entries) GetEntryFromMAC(mac string) (Entry, error) {
	for _, entry := range e {
		if entry.HWAddress == mac {
			return entry, nil
		}
	}
	return Entry{}, errors.New("mac not found")
}
