package main

import (
	"fmt"
	"time"

	"./../shared"
)

type Git int

func (t *Git) GetVersion(_, reply *shared.Version) error {
	time := time.Now()
	reply.Name = fmt.Sprintf("%s%d", "1.1.", time.Second())
	return nil
}
