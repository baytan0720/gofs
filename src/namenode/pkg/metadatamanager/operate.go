package metadatamanager

import (
	"gofs/src/service"
	"strconv"
	"strings"
	"time"
)

func Put(parentid int64, filename string, entryid, size int64, modtime string, blocks []string) {
	dbPut(formatKey(parentid, filename), formatFileVal(entryid, size, modtime, blocks))
}

func Get(path string) (service.FileStatus, []string) {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana)-1; i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return service.FileStatus_fPathNotFound, nil
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return service.FileStatus_fIsFile, nil
		}
		parentid = entryid
	}
	val := dbGet(formatKey(parentid, ana[len(ana)-1]))
	if val == "" {
		return service.FileStatus_fPathNotFound, nil
	}
	return service.FileStatus_fPathFound, getBlocks(val)
}

func Mkdir(path, dirname string) service.FileStatus {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana); i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return service.FileStatus_fPathNotFound
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return service.FileStatus_fIsFile
		}
		parentid = entryid
	}
	if dbGet(formatKey(parentid, dirname)) != "" {
		return service.FileStatus_fExist
	}
	dbPut(formatKey(parentid, dirname), formatDirVal(NewEntryId(), time.Now().Format("2006-01-02 15:04:05")))
	return 0
}

func List(path string) ([]string, service.FileStatus) {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana); i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return nil, service.FileStatus_fPathNotFound
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return nil, service.FileStatus_fIsFile
		}
		parentid = entryid
	}
	list := make([]string, 0, 128)
	for _, v := range dbPrefixScan(strconv.FormatInt(parentid, 10)+"_", 128) {
		list = append(list, getFileName(v[0]))
	}
	return list, 0
}

func Rename(path, newname string) service.FileStatus {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana)-1; i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return service.FileStatus_fPathNotFound
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return service.FileStatus_fIsFile
		}
		parentid = entryid
	}
	val := dbGet(formatKey(parentid, ana[len(ana)-1]))
	if val == "" {
		return service.FileStatus_fPathNotFound
	}
	if dbGet(formatKey(parentid, newname)) != "" {
		return service.FileStatus_fExist
	}
	dbPut(formatKey(parentid, newname), updateModtime(val))
	dbDelete(formatKey(parentid, ana[len(ana)-1]))
	return 0
}

func Delete(path string) (service.FileStatus, []string) {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana)-1; i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return service.FileStatus_fPathNotFound, nil
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return service.FileStatus_fIsFile, nil
		}
		parentid = entryid
	}
	val := dbGet(formatKey(parentid, ana[len(ana)-1]))
	if val == "" {
		return service.FileStatus_fPathNotFound, nil
	}
	if getFileType(val) == 0 {
		delDir(parentid, getEntryId(val), ana[len(ana)-1])
	} else {
		dbDelete(formatKey(parentid, ana[len(ana)-1]))
		return 0, getBlocks(val)
	}
	return 0, nil
}

func Stat(path string) (*service.FileInfo, service.FileStatus) {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana)-1; i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return nil, service.FileStatus_fPathNotFound
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return nil, service.FileStatus_fIsFile
		}
		parentid = entryid
	}
	val := dbGet(formatKey(parentid, ana[len(ana)-1]))
	if val == "" {
		return nil, service.FileStatus_fPathNotFound
	}
	_, size, filetype, modtime, _ := parseVal(val)
	if filetype == 0 {
		return &service.FileInfo{
			FileType: service.FileStatus_fIsDirectory,
			Modtime:  modtime,
		}, 0
	}
	return &service.FileInfo{
		FileType: service.FileStatus_fIsFile,
		Size:     size,
		Modtime:  modtime,
	}, 0
}

func delDir(parentid, entryid int64, filename string) {
	for _, entry := range dbPrefixScan(strconv.FormatInt(entryid, 10)+"_", 128) {
		if getFileType(entry[1]) == 0 {
			delDir(entryid, getEntryId(entry[1]), getFileName(entry[0]))
		} else {
			dbDelete(formatKey(entryid, getFileName(entry[0])))
		}
	}
	dbDelete(formatKey(parentid, filename))
}

func PutCheckPath(path, filename string) (service.FileStatus, int64) {
	ana := strings.Split(path, "/")
	if ana[len(ana)-1] == "" {
		ana = ana[:len(ana)-1]
	}
	parentid := root
	for i := 1; i < len(ana); i++ {
		val := dbGet(formatKey(parentid, ana[i]))
		if val == "" {
			return service.FileStatus_fPathNotFound, -1
		}
		entryid, _, filetype, _, _ := parseVal(val)
		if filetype == 1 {
			return service.FileStatus_fIsFile, -1
		}
		parentid = entryid
	}
	if dbGet(formatKey(parentid, filename)) != "" {
		return service.FileStatus_fExist, -1
	}
	return service.FileStatus_fPathFound, parentid
}
