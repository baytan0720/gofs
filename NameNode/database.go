package main

import (
	"log"

	"github.com/xujiajun/nutsdb"
)

func opendb() *nutsdb.DB {
	dbop := nutsdb.DefaultOptions
	dbop.Dir = "./metadata"
	db, err := nutsdb.Open(dbop)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func set(key, val []byte, bucket string) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}

func get(key []byte, bucket string) []byte {
	var res []byte
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(bucket, key); err != nil {
				return err
			} else {
				res = e.Value
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
	return res
}

func del(key []byte, bucket string) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(bucket, key); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}

func getAll(bucket string) []*nutsdb.Entry {
	entries := []*nutsdb.Entry{}
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			var err error
			entries, err = tx.GetAll(bucket)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
	return entries
}
