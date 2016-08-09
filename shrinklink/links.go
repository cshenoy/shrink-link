package shrinklink

import (
	"errors"
	"github.com/mediocregopher/radix.v2/pool"
	// "log"
)

var db *pool.Pool

var ErrNoAlbum = errors.New("models: no album found")

// func init() {
// 	var err error
// 	db, err = pool.New("tcp", "localhost:6379", 10)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

// func Lookup(hash string) error {
// 	reply, err := db.Cmd("HGETALL", "album:"+id).Map()
// 	if err != nil {
// 		return nil, err
// 	} else if len(reply) == 0 {
// 		return nil, ErrNoAlbum
// 	}

// 	return populateAlbum(reply)
// }

// func IncrementLikes(id string) error {
// 	conn, err := db.Get()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Put(conn)

// 	exists, err := conn.Cmd("EXISTS", "album:"+id).Int()
// 	if err != nil {
// 		return err
// 	} else if exists == 0 {
// 		return ErrNoAlbum
// 	}

// 	err = conn.Cmd("MULTI").Err
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.Cmd("HINCRBY", "album:"+id, "likes", 1).Err
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.Cmd("ZINCRBY", "likes", 1, id).Err
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.Cmd("EXEC").Err
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
