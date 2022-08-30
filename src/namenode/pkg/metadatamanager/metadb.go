package metadatamanager

import (
	"log"
	"strconv"

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
// 	}
// }

func dbPrefixScan(prefix string, limit int) [][]string {
	list := make([][]string, 0, limit)
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if entries, _, err := tx.PrefixScan(bucket, []byte(prefix), 0, limit); err != nil {
				return err
			} else {
				for _, v := range entries {
					list = append(list, []string{string(v.Key), string(v.Value)})
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
