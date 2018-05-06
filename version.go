package baikalver

import (
	"fmt"
	"strconv"
)

// VersionFromFWV returns Firmware version from 'FWV' value
func VersionFromFWV(s string) (string, error) {
	fwv, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}
	fwvn := uint8(fwv)
	return fmt.Sprintf("%d.%d", fwvn>>4&255, fwv&15), nil
}
