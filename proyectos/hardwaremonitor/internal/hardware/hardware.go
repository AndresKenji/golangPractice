package hardware

import (
	"fmt"
	"math"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemInfo struct {
	Hostname     string
	MemoryMB     int
	FreeMemoryMB int
	Os           string
	Platform     string
	Uptime       uint64
}

type NetInfo struct {
	Interfaces []net.InterfaceStat

}

func GetNetSection() (*NetInfo, error) {
	ni := &NetInfo{}
	itf, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ni.Interfaces = itf

	return ni, nil
}

func GetSystemSection() (*SystemInfo, error) {
	si := &SystemInfo{}
	runTimeOs := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}

	si.Hostname = hostStat.Hostname
	si.MemoryMB = int(vmStat.Total / (1024 * 1024))
	si.FreeMemoryMB = int(vmStat.Free / (1024 * 1024))
	si.Os = runTimeOs
	si.Platform = hostStat.Platform
	si.Uptime = hostStat.Uptime

	return si, nil
}

type CpuInfo struct {
	CPU          string
	Cores        int
	LogicalCores int
	Percent      float64
}

func GetCpuSection() (*CpuInfo, error) {
	ci := &CpuInfo{}
	cpuStat, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	logicalcores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	cores, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}

	ci.CPU = cpuStat[0].ModelName
	ci.Cores = cores
	ci.LogicalCores = logicalcores
	ci.Percent = math.Round(percent[0])

	return ci, nil
}

type Disk struct {
	Path string
	TotalSpaceGB float64
	UsedSpaceGB  float64
	FreeSpaceGB  float64
}

type DiskInfo struct {
	Disks []*Disk
}

func GetDiskSection() (*DiskInfo, error) {
	di := &DiskInfo{
		Disks: []*Disk{},
	}
	
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	
	for _, dsk := range partitions {
		diskStat, err := disk.Usage(dsk.Mountpoint)
		if err != nil {
			return nil, err
		}

				
		diskInfo := &Disk{
			Path: diskStat.Path,
			TotalSpaceGB: math.Round(float64(diskStat.Total) / (1024 * 1024 * 1024)),
			FreeSpaceGB:  math.Round(float64(diskStat.Free) / (1024 * 1024 * 1024)),
			UsedSpaceGB:  math.Round(float64(diskStat.Used) / (1024 * 1024 * 1024)),
		}

		di.Disks = append(di.Disks, diskInfo)
	}

	return di, nil
}

func (di *DiskInfo) GetHtml() string {
	html := ""
	for _, dsk := range di.Disks{
		html += fmt.Sprintf(`
		<div class="p-2"  style="max-height: 400px; max-width: 400px;">
			<h4>%s</h4>
			<table class="table table-dark table-striped">
				<tr><td>Total</td><td>%v GB</td></tr>
				<tr><td>Used</td><td>%v GB</td></tr>
				<tr><td>Free</td><td>%v GB</td></tr>
			</table>
		</div>`,dsk.Path, dsk.TotalSpaceGB, dsk.UsedSpaceGB, dsk.FreeSpaceGB)
	}

	return html
}

func (ci *CpuInfo) GetHtml() string {
	html := fmt.Sprintf(`
	<table class="table table-dark table-striped">
		<tr><td>CPU</td><td>%s</td></tr>
		<tr><td>Cores</td><td>%d</td></tr>
		<tr><td>Logical Cores</td><td>%d</td></tr>
		<tr><td>Use Percent</td><td>%v%%</td></tr>
	</table>
	`,ci.CPU, ci.Cores, ci.LogicalCores, ci.Percent)

	return html
}

func (si *SystemInfo) GetHtml() string {
	html := fmt.Sprintf(`
	<table class="table table-dark table-striped">
		<tr><td>Host Name</td><td>%s</td></tr>
		<tr><td>Memory</td><td>%d MB</td></tr>
		<tr><td>Free Memory</td><td>%d MB</td></tr>
		<tr><td>OS</td><td>%s</td></tr>
		<tr><td>Platform</td><td>%s</td></tr>
		<tr><td>Uptime</td><td>%d seg</td></tr>
	</table>
	`,si.Hostname, si.MemoryMB, si.FreeMemoryMB, si.Os, si.Platform, si.Uptime)

	return html

}

func (ni *NetInfo) GetHtml() string {
	html := ""
	for _, itf := range ni.Interfaces{
		html += fmt.Sprintf(`
		<div class="p-2"  style="max-height: 400px; max-width: 400px;">
			<h4>%s</h4>
			<table class="table table-dark table-striped">
				<tr><td>IP</td><td>%s</td></tr>
				<tr><td>MAC</td><td>%s</td></tr>
				<tr><td>MTU</td><td>%d</td></tr>
			</table>
		</div>`,itf.Name, itf.Addrs[0].Addr, itf.HardwareAddr,itf.MTU) 
	}

	return html
}