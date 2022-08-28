package metadatamanager

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xujiajun/nutsdb"
)

var db *nutsdb.DB
var bucket string = "metadata"

func startDb(dbpath string) {
	var err error
	db, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(dbpath),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func dbPut(key, val string) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket, []byte(key), []byte(val), 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}

func dbGet(key string) string {
	var val string
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(bucket, []byte(key)); err != nil {
				return err
			} else {
				val = string(e.Value)
			}
			return nil
		}); err != nil {
		return val
	}
	return val
}

func dbDelete(key string) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(bucket, []byte(key)); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}

// func GetAll() {
// 	if err := db.View(
// 		func(tx *nutsdb.Tx) error {
// 			entries, err := tx.GetAll(bucket)
// 			if err != nil {
// 				return err
// 			}
// 			for _, entry := range entries {
// 				log.Println(string(entry.Key), string(entry.Value))
// 			}
// 			return nil
// 		}); err != nil {
// 		log.Println(err)
// 	}
// }

func dbPrefixScan(prefix string, limit int) []string {
	list := make([]string, 0, limit)
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if entries, _, err := tx.PrefixScan(bucket, []byte(prefix), 0, limit); err != nil {
				return err
			} else {
				for _, v := range entries {
					list = append(list, getFileName(string(v.Key)))
				}
			}
			return nil
		}); err != nil {
		return nil
	}
	return list
}

func initEntryId() {
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get("entryid", []byte("id")); err != nil {
				entryidincrease = 1
				return err
			} else {
				entryidincrease, _ = strconv.ParseInt(string(e.Value), 10, 64)
			}
			return nil
		}); err != nil {
		db.Update(
			func(tx *nutsdb.Tx) error {
				if err := tx.Put("entryid", []byte("id"), []byte("1"), 0); err != nil {
					return err
				}
				return nil
			})
	}
}

func loadBackup(metadatapath, metabackup string) {
	backup, err := os.Open(metabackup + "/0.dat")
	if err != nil {
		return
	}
	defer backup.Close()
	os.Remove(metadatapath + "/0.dat")
	metafile, err := os.Create(metadatapath + "/0.dat")
	if err != nil {
		return
	}
	defer metafile.Close()
	io.Copy(metafile, backup)
}

func dbBackup(path string, t int) {
	for {
		time.Sleep(time.Duration(t) * time.Minute)
		db.Update(
			func(tx *nutsdb.Tx) error {
				if err := tx.Put("entryid", []byte("id"), []byte(strconv.FormatInt(entryidincrease, 10)), 0); err != nil {
					return err
				}
				return nil
			})
		err := db.Backup(path)
		if err != nil {
			logrus.WithField("o", "Backup").Warn("MetaData backup fail")
		} else {
			logrus.WithField("o", "Backup").Info("MetaData backup")
		}
	}
}
