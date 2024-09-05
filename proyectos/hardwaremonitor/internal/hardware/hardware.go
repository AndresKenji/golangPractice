package hardware

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Hostname     string
	MemoryMB     int
	FreeMemoryMB int
	Os           string
	Platform     string
	Uptime       uint64
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
	ci.Percent = percent[0] // O considera promediar el valor si tienes múltiples CPUs

	return ci, nil
}

type DiskInfo struct {
	TotalSpaceGB float64
	UsedSpaceGB  float64
	FreeSpaceGB  float64
}

func GetDiskSection() (*DiskInfo, error) {
	di := &DiskInfo{}
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	di.TotalSpaceGB = float64(diskStat.Total) / (1024 * 1024 * 1024)
	di.FreeSpaceGB = float64(diskStat.Free) / (1024 * 1024 * 1024)
	di.UsedSpaceGB = float64(diskStat.Used) / (1024 * 1024 * 1024)
	return di, nil
}

func (di *DiskInfo) GetHtml() string {
	html := fmt.Sprintf(`
	<table class="table table-dark table-striped">
		<tr>
			<td>Total</td>
			<td>%v</td>
		</tr>
		<tr>
			<td>Used</td>
			<td">%v</td>
		</tr>
		<tr>
			<td>Free</td>
			<td">%v</td>
		</tr>
    </table>
	`, di.TotalSpaceGB, di.UsedSpaceGB, di.FreeSpaceGB)

	return html
}

func (ci *CpuInfo) GetHtml() string {
	html := fmt.Sprintf(`
	<table class="table table-dark table-striped">
		<tr>
			<td>CPU</td>
			<td>%s</td>
		</tr>
		<tr>
			<td>Cores</td>
			<td>%d</td>
		</tr>
		<tr>
			<td>Logical Cores</td>
			<td>%d</td>
		</tr>
		<tr>
			<td>Use Percent</td>
			<td>%v</td>
		</tr>
	</table>
	`,ci.CPU, ci.Cores, ci.LogicalCores, ci.Percent)

	return html
}

func (si *SystemInfo) GetHtml() string {
	html := fmt.Sprintf(`
	<table class="table table-dark table-striped">
		<tr>
			<td>Host Name</td>
			<td>%s</td>
		</tr>
		<tr>
			<td>Memory MB</td>
			<td>%d</td>
		</tr>
		<tr>
			<td>Free Memory MB</td>
			<td>%d</td>
		</tr>
		<tr>
			<td>OS</td>
			<td>%s</td>
		</tr>
		<tr>
			<td>Platform</td>
			<td>%s</td>
		</tr>
		<tr>
			<td>Uptime</td>
			<td>%d</td>
		</tr>
	</table>
	`,si.Hostname, si.MemoryMB, si.FreeMemoryMB, si.Os, si.Platform, si.Uptime)

	return html

}