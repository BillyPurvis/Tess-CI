package main

import (
	"fmt"
	"os"

	"github.com/BillyPurvis/tess-ci/sshconnect"
)

func main() {
	fmt.Println("Starting Tess CI...")

	// Check there's an arugment
	if len(os.Args) == 5 { // FIXME: Use Flags package
		hostname := os.Args[1]
		port := os.Args[2]
		user := os.Args[3]
		cmd := os.Args[4]

		agent, err := sshconnect.SSHAgent()
		if err != nil {
			panic(err) // FIXME: Better Error handling
		}

		KeyPair, err := sshconnect.KeyPair("/Users/billypurvis/.ssh/id_rsa")
		if err != nil {
			panic(err) // FIXME: Better Error handling
		}

		// Create Connection
		client, err := sshconnect.MakeConnection(
			fmt.Sprintf("%v:%v", hostname, port),
			user,
			agent,
			KeyPair,
		)
		defer client.Close()

		if err != nil {
			panic(err)
		}

		// Start New Client Connection
		session, err := client.NewSession()

		defer session.Close()
		if err != nil {
			fmt.Printf("%v", err)
		}

		// Get Output
		session.Stdout = os.Stdout

		err = session.Run(cmd)

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Println("\n\n ****** Command Run Successfully! 🎉 ****** \n\n ")
	}

}
