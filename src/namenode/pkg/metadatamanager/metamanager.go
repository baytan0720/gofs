package metadatamanager

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

var root int64 = 0
var entryidincrease int64
var mu sync.Mutex

func Start(metadatapath string) {
	startDb(metadatapath)
	initEntryId()
}

func NewEntryId() int64 {
	mu.Lock()
	entryid := entryidincrease
	entryidincrease++
	mu.Unlock()
	return entryid
}

func formatKey(parentid int64, filename string) string {
	return strconv.FormatInt(parentid, 10) + "_" + filename
}

func formatDirVal(entryid int64, modtime string) string {
	return strconv.FormatInt(entryid, 10) + "_0_" + modtime
}

func formatFileVal(entryid, size int64, modtime string, blocks []string) string {
	val := strconv.FormatInt(entryid, 10) + "_1_" + modtime + "_" + strconv.FormatInt(size, 10) + "_" + strconv.Itoa(len(blocks))
	for _, v := range blocks {
		val += "_" + v
	}
	return val
}

func parseKey(key string) (int64, string) {
	ana := strings.Split(key, "_")
	parentid, _ := strconv.ParseInt(ana[0], 10, 64)
	return parentid, ana[1]
}

func parseVal(val string) (entryid, size int64, filetype int, modtime string, blocks []string) {
	ana := strings.Split(val, "_")
	entryid, _ = strconv.ParseInt(ana[0], 10, 64)
	modtime = ana[2]
	if ana[1] == "0" {
		return entryid, 0, 0, modtime, nil
	} else {
		size, _ = strconv.ParseInt(ana[3], 10, 64)
		blocknum, _ := strconv.Atoi(ana[4])
		blocks = make([]string, 0, blocknum)
		for i := 5; i < len(ana); i++ {
			blocks = append(blocks, ana[i])
		}
		return entryid, size, 1, modtime, blocks
	}
}

func getParentId(key string) int64 {
	ana := strings.Split(key, "_")
	parentid, _ := strconv.ParseInt(ana[0], 10, 64)
	return parentid
}

func getFileName(key string) string {
	ana := strings.Split(key, "_")
	return ana[1]
}

func getFileType(val string) int {
	ana := strings.Split(val, "_")
	if ana[1] == "0" {
		return 0
	} else {
		return 1
	}
}

func getEntryId(val string) int64 {
	ana := strings.Split(val, "_")
	entryid, _ := strconv.ParseInt(ana[0], 10, 64)
	return entryid
}

func getSize(val string) int64 {
	ana := strings.Split(val, "_")
	entryid, _ := strconv.ParseInt(ana[3], 10, 64)
	return entryid
}

func getBlocks(val string) []string {
	ana := strings.Split(val, "_")
	blocknum, _ := strconv.Atoi(ana[4])
	blocks := make([]string, 0, blocknum)
	for i := 5; i < len(ana); i++ {
		blocks = append(blocks, ana[i])
	}
	return blocks
}

func updateModtime(val string) string {
	entryid, size, filetype, _, blocks := parseVal(val)
	if filetype == 0 {
		return formatDirVal(entryid, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		return formatFileVal(entryid, size, time.Now().Format("2006-01-02 15:04:05"), blocks)
	}
}
