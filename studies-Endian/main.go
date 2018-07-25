package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	Endian = binary.BigEndian // Which endian?
)

func main() {
	db, err := bolt.Open("temp.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer time.Sleep(time.Second)

	tempBucket := []byte("TMP")

	db.Update(func(tx *bolt.Tx) error {
		data := Row{
			Key: Key{
				Head:    0x1A1A1A,
				Mark:    0x1010,
				Counter: 0x01,
			},
			At:      time.Now().UTC().Unix(),
			Payload: 10,
		}

		keyBytes, err := marshal(&data.Key)
		if err != nil {
			return erp(err)
		}

		dataBytes, err := marshal(&data)
		if err != nil {
			return erp(err)
		}

		b, err := tx.CreateBucketIfNotExists(tempBucket)
		if err != nil {
			return erp(err)
		}

		err = b.Put(keyBytes, dataBytes)
		if err != nil {
			return erp(err)
		}

		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tempBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			k, v := k, v
			fmt.Println(k, v)

			dt := Row{}
			err := unmarshal(&dt, v)
			if err != nil {
				return erp(err)
			}
			fmt.Printf("key=%X, value=%v\n", k, dt)
		}

		return nil
	})
}

func erp(e error) error {
	log.Println(e)
	return e
}

func marshal(pointerToData interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, Endian, pointerToData)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unmarshal(pointerToData interface{}, bs []byte) error {
	buffer := bytes.NewBuffer(bs)
	err := binary.Read(buffer, Endian, pointerToData)
	if err != nil {
		return err
	}
	return nil
}

type Row struct {
	Key
	At      int64
	Payload int64
}

type Key struct {
	Head    uint32
	Mark    uint16
	Counter uint16
}
