//+build darwin build

package hyperkit

import "fmt"

const (
	MachineDriverCommand = "crc-driver-hyperkit"
	MachineDriverVersion = "0.12.6"
	HyperkitCommand      = "hyperkit"
)

var (
	HyperkitDownloadURL      = fmt.Sprintf("https://github.com/code-ready/machine-driver-hyperkit/releases/download/v%s/hyperkit", MachineDriverVersion)
	MachineDriverDownloadURL = fmt.Sprintf("https://github.com/code-ready/machine-driver-hyperkit/releases/download/v%s/crc-driver-hyperkit", MachineDriverVersion)
)
