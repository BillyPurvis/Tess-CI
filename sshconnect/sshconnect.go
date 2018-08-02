package sshconnect

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// KeyPair Get Key Pair
func KeyPair(keyFile string) (ssh.AuthMethod, error) {

	pem, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(pem)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

// SSHAgent get the SSH agent
func SSHAgent() (ssh.AuthMethod, error) {
	agentSock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeysCallback(agent.NewClient(agentSock).Signers), nil
}

// MakeConnection Makes a Connection to SSH
func MakeConnection(host string, user string, methods ...ssh.AuthMethod) (*ssh.Client, error) {
	fmt.Printf("%v", host)
	// Make Channel

	// TODO: User a credential manager
	config := ssh.ClientConfig{
		User: user,
		Auth: methods,
	}

	// FIXME: We shouldn't need to do this at all
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	return ssh.Dial("tcp", host, &config)

}
