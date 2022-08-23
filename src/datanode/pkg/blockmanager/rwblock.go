package blockmanager

import "os"

func WriteBlock(path, blockid string, data []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
