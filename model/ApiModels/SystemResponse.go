package ApiModels

import "MyCloudDisk/utils/SystemUtils"

type systemData struct {
	Host *SystemUtils.HostProfile   `json:"host"`
	Cpu  *SystemUtils.CPUProfile    `json:"cpu"`
	Mem  *SystemUtils.MemoryProfile `json:"mem"`
}

type SystemRespBody struct {
	BasicRespBody
	systemData `json:"data"`
}

func (system *SystemRespBody) SetData(data map[string]interface{}) RespBody {
	if data == nil {
		return system
	}
	if cpu, ok := data["cpu"]; ok {
		system.Cpu = cpu.(*SystemUtils.CPUProfile)
	}
	if host, ok := data["host"]; ok {
		system.Host = host.(*SystemUtils.HostProfile)
	}
	if mem, ok := data["mem"]; ok {
		system.Mem = mem.(*SystemUtils.MemoryProfile)
	}
	return system
}
