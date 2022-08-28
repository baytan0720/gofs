package blockmanager

import (
	"os"
)

func WriteBlock(path, blockid string, data []byte) error {
	file, err := os.OpenFile(path+"/"+blockid, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}

func ReadBlock(path, blockid string) ([]byte, error) {
	file, err := os.OpenFile(path+"/"+blockid, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	databuf := make([]byte, info.Size())
	file.Read(databuf)
	return databuf, nil
}

func DeleteBlock(path, blockid string) error {
	err := os.Remove(path + "/" + blockid)
	if err != nil {
		return err
	}
	return nil
}
