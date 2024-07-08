package engine

import (
	"sync"

	"api.brutecore/libs/lib_db"
)

type Database struct {
	id        int64
	data_type int64
	db        *lib_db.DB
	buffer    []ComboListRecord
	mul       sync.Mutex
	sWorkerC  int
}

func NewDatabase(d *lib_db.DB, database_id int64) (*Database, error) {
	res, err := d.QueryRow(lib_db.TxRead, "SELECT DATA_TYPE FROM DATABASE WHERE ID = $1", database_id)
	if err != nil {
		return nil, err
	}

	return &Database{
		id:        database_id,
		db:        d,
		data_type: (*res)[0]["data_type"].(int64),
	}, nil
}

func (d *Database) GetStartBatch(session_id int64) error {
	res, err := d.db.QueryRow(lib_db.TxRead, QGetMaxConId, session_id)
	if err != nil {
		return err
	}
	max_con_id := (*res)[0]["M"].(string)

	res, err = d.db.QueryRow(lib_db.TxRead, QGetComboListBatchStart, d.id, max_con_id, session_id, d.id, max_con_id, d.sWorkerC*15)
	if err != nil {
		return err
	}

	for _, item := range *res {
		d.buffer = append(d.buffer, ComboListRecord{
			id:            item["id"].(int64),
			database_link: item["database_link"].(int64),
			database_id:   d.id,
			data:          item["data"].(string),
			con_id:        item["con_id"].(string),
		})
	}

	return nil
}

func (d *Database) GetComboLine() *ComboListRecord {
	d.mul.Lock()
	defer d.mul.Unlock()

	var cl ComboListRecord
	switch len(d.buffer) {
	case 0:
		cl.data = "-1"
	case 1:
		res, err := d.db.QueryRow(lib_db.TxRead, QGetComboListBatch, d.id, d.buffer[0].con_id, d.sWorkerC*15)
		if err != nil {
			cl.data = "~"
		} else {
			for _, item := range *res {
				d.buffer = append(d.buffer, ComboListRecord{
					id:            item["id"].(int64),
					database_link: item["database_link"].(int64),
					database_id:   d.id,
					data:          item["data"].(string),
					con_id:        item["con_id"].(string),
				})
			}

			cl = d.buffer[0]
			d.buffer = d.buffer[1:]
		}
	default:
		cl = d.buffer[0]
		d.buffer = d.buffer[1:]
	}

	return &cl
}
