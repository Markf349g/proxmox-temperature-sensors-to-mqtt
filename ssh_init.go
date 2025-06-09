package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHStruct struct {
	Form    *SSHForm
	Conn    *ssh.Client
	Session *ssh.Session
}

func SSHInit(form *SSHForm) *SSHStruct {
	sshPtr := SSHStruct{}
	sshPtr.Form = form

	conf := &ssh.ClientConfig{
		User: sshPtr.Form.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPtr.Form.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	var err error
	for {
		sshPtr.Conn, err = ssh.Dial("tcp", sshPtr.Form.Host, conf)
		if err != nil {
			log.Println(err.Error())
			time.Sleep(time.Duration(sshPtr.Form.Delay) * time.Millisecond)
		} else {
			break
		}
	}

	return &sshPtr
}

func (sshPtr *SSHStruct) Close() {
	sshPtr.Conn.Close()
	sshPtr.Session.Close()
}

func SSHRequest(sshPtr *SSHStruct, request string) (string, error) {
	var err error
	var out []byte

	sshPtr.Session, err = sshPtr.Conn.NewSession()
	if err != nil {
		log.Fatalln("invalid new session ssh request")
	}

	defer sshPtr.Session.Close()
	out, err = sshPtr.Session.CombinedOutput(request)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), err
}

func WaitTheInternet(delayTime time.Duration) {
	for {
		resp, err := http.Get("http://example.com")
		if err != nil {
			log.Println("invalid http get request")
			time.Sleep(delayTime)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				break
			} else {
				log.Println("invalid http status code")
				time.Sleep(delayTime)
			}
		}
	}
}
