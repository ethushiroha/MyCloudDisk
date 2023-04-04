package SystemUtils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

type HostProfile struct {
	Name string
	OS   string
}

func GetHostProfile() (*HostProfile, error) {
	info, err := host.Info()
	if err != nil {
		return nil, err
	}
	profile := &HostProfile{
		Name: info.Hostname,
		OS:   info.OS,
	}
	return profile, nil
}

type CPUProfile struct {
	Nums    int
	Percent int
}

func GetCPUProfile() (*CPUProfile, error) {
	percent, err := cpu.Percent(time.Millisecond*300, false)
	if err != nil {
		return nil, err
	}
	profile := &CPUProfile{
		Nums:    runtime.NumCPU(),
		Percent: int((percent[0] + 0.005) * 100),
	}
	return profile, nil
}

type MemoryProfile struct {
	Total   uint64
	Used    uint64
	Percent int
}

func GetMemoryProfile() (*MemoryProfile, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	profile := &MemoryProfile{
		Total:   memory.Total,
		Used:    memory.Used,
		Percent: int((memory.UsedPercent + 0.005) * 100),
	}
	return profile, nil
}
