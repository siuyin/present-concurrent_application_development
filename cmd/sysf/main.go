package main

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/siuyin/dflt"
)

//10 OMIT
func main() {
	fmt.Println("sysf main program")

	sc := stanConnect() // get a messaging server connection
	defer sc.Close()

	userStart()
	orderStart()
	delvStart()
	webStart(sc)

	select {} // wait forever
}

//20 OMIT
func stanConnect() stan.Conn {
	clusterID := dflt.EnvString("NATS_CLUSTER_ID", "test-cluster")
	clientID := dflt.EnvString("NATS_CLIENT_ID", "sysf")
	natsURL := dflt.EnvString("NATS_URL", "nats://127.0.0.1:4222/")
	sc, err := stan.Connect(clusterID, clientID,
		stan.NatsURL(natsURL),
		stan.SetConnectionLostHandler(
			func(_ stan.Conn, err error) {
				log.Fatalf("FATAL: stanConnect: %v", err)
			}),
	)
	if err != nil {
		log.Fatalf("FATAL: stanConnect: %v", err)
	}
	return sc
}

//30 OMIT
