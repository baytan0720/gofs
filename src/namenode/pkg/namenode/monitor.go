package namenode

import (
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type Monitor struct {
	Status       int    `json:"status"`
	Starttime    string `json:"starttime"`
	Cpunum       int    `json:"cpunum"`
	Totalmem     int    `json:"totalmem"`
	Totaldisk    int    `json:"totaldisk"`
	Replicanum   int    `json:"replicanum"`
	Blocksize    int    `json:"blocksize"`
	datanodelist []*datanode
}

type Load struct {
	Cpudata  []int    `json:"cpudata"`
	Memdata  []int    `json:"memdata"`
	Timedata []string `json:"timedata"`
	Diakused int      `json:"diskused"`
	Diskfree int      `json:"diskfree"`
}

type DataNode struct {
	Id        int    `json:"id"`
	Starttime string `json:"starttime"`
	Address   string `json:"address"`
	Disk      int    `json:"disk"`
	Used      int    `json:"used"`
	Status    string `json:"status"`
}

var Loa *Load
var Moni *Monitor

func (nn *NameNode) monitorServer() {
	cpu, _ := cpu.Counts(true)
	mem, _ := mem.VirtualMemory()
	disk, _ := disk.Usage("/")
	Moni = &Monitor{
		Status:       0,
		Starttime:    nn.Starttime,
		Cpunum:       cpu,
		Totalmem:     int(mem.Total) >> 30,
		Totaldisk:    int(disk.Total) >> 30,
		Replicanum:   nn.ReplicaNum,
		Blocksize:    int(nn.BlockSize) >> 20,
		datanodelist: nn.DataNodeList,
	}
	Loa = &Load{
		Cpudata:  make([]int, 6),
		Memdata:  make([]int, 6),
		Timedata: make([]string, 6),
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.File("../static/index.html")
	})
	e.GET("/api/init", initargs)
	e.GET("/api/shudown", shutdown)
	e.GET("/api/load", getload)
	e.GET("/api/datanodes", getdatanodes)
	e.GET("/api/block", getblock)
	e.GET("/api/log", getLog)
	go updateload()
	go e.Start(":8090")
}

func initargs(c echo.Context) error {
	return c.JSON(200, Moni)
}

func shutdown(c echo.Context) error {
	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
	return c.String(200, "")
}

func getload(c echo.Context) error {
	return c.JSON(200, Loa)
}

func updateload() {
	for {
		Loa.Timedata = append(Loa.Timedata[1:], time.Now().Format("15:04"))
		percpu, _ := cpu.Percent(0, true)
		Loa.Cpudata = append(Loa.Cpudata[1:], int(percpu[0]))
		stat, _ := mem.VirtualMemory()
		Loa.Memdata = append(Loa.Memdata[1:], int(stat.UsedPercent))
		disk, _ := disk.Usage("/")
		Loa.Diakused = int(disk.UsedPercent)
		Loa.Diskfree = 100 - Loa.Diakused
		time.Sleep(time.Minute)
	}
}

func getdatanodes(c echo.Context) error {
	datanodes := make([]*DataNode, 0, len(Moni.datanodelist))
	for i, v := range Moni.datanodelist {
		if v == nil {
			continue
		}
		datanodes = append(datanodes, &DataNode{
			Id:        i,
			Starttime: v.info.StartTime,
			Address:   v.info.Addr,
			Disk:      int(v.load.TotalDisk) >> 30,
			Used:      int(v.load.UsedDisk) >> 30,
			Status:    "Active",
		})
	}
	return c.JSON(200, datanodes)
}

func getblock(c echo.Context) error {
	index, _ := strconv.Atoi(c.Request().FormValue("id"))
	dn := Moni.datanodelist[index]
	if dn != nil {
		c.JSON(200, dn.info.Blocks)
	}
	return c.String(400, "")
}

func getLog(c echo.Context) error {
	return c.File("../../../logs/" + Moni.Starttime + ".log")
}
