package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	botanDB "github.com/DolenL3/botanDB"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("run app: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	client := botanDB.NewClient(0, time.Second)

	exampleKey := "exampleKey"
	exampleKey2 := "exampleKey2"

	// sets
	err := client.Set(ctx, exampleKey, "exampleValue", 0)
	if err != nil {
		return errors.Wrap(err, "botanDB set value")
	}
	err = client.Set(ctx, exampleKey2, 99, time.Second)
	if err != nil {
		return errors.Wrap(err, "botanDB set value")
	}
	// gets
	res, err := client.Get(ctx, exampleKey)
	if err != nil {
		return errors.Wrap(err, "botanDB get value")
	}
	fmt.Printf("%v, type: %T\n", res, res)

	res, err = client.Get(ctx, exampleKey2)
	if err != nil {
		return errors.Wrap(err, "botanDB get value")
	}
	fmt.Printf("%v, type: %T\n", res, res)

	// deletes
	err = client.Delete(ctx, exampleKey)
	if err != nil {
		return errors.Wrap(err, "botanDB delete value")
	}
	// gets after deletion/expiration
	res, err = client.Get(ctx, exampleKey)
	if err != nil {
		if errors.Is(err, botanDB.ErrKeyNotFound) {
			fmt.Println("caugt key not found, returned value:", res)
		} else {
			return errors.Wrap(err, "botanDB get value")
		}
	}
	time.Sleep(time.Second)
	resInt, err := client.GetInt(ctx, exampleKey2)
	if err != nil {
		if errors.Is(err, botanDB.ErrKeyNotFound) {
			fmt.Println("caugt key not found, returned value:", resInt)
		} else {
			return errors.Wrap(err, "botanDB get value")
		}
	}

	return nil
}
