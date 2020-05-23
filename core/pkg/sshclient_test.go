package pkg

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"testing"
)

func TestWithKey(*testing.T) {
	user := "root"
	remoteAddr := "47.94.152.92:22"
	keyPath := "/Users/yz/Documents/tron-work/ssh/vol_0821.pem"

	client, err := GetSSHClinet(remoteAddr, user, keyPath)
	if nil != err {
		fmt.Printf("getSSHClient failed:%v\n", err)
		return
	}
	defer client.Close()

	//dialFunc := func(addr string) (net.Conn, error) {
	//	conn, err := client.Dial("tcp", addr)
	//	fmt.Printf("sshClient dial to:%v over:[%v], err:[%v], conn:[%v]\n", addr, "tcp", err, conn)
	//	return conn, err
	//}
	//
	//mysql.RegisterDial("tcpOverSSH", dialFunc)

	driver := "mysql"
	dbUser := "vlm"
	dbPass := "1MzY%5Nj&AxNzU4z"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "cloud"
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcpOverSSH(%s)/%s", dbUser, dbPass, dbHost, dbName))
	db, err := GetRemoteDBOverSSH(driver, dbUser, dbPass, dbHost, dbPort, dbName, client)
	if nil != err {
		fmt.Printf("connect to mysql over tcpOverSSH failed:%v\n", err)
		return
	}

	fmt.Println(db.Ping())

	sqlStr := "select id, itiger_uid, parent_id from ft_user where id = 30"
	if rows, err := db.Query(sqlStr); err == nil {
		for rows.Next() {
			var id int64
			var itigerUID, parentID string
			rows.Scan(&id, &itigerUID, &parentID)
			fmt.Printf("%20v\t%20v\t%20v\t\n", id, itigerUID, parentID)
		}
		rows.Close()
	} else {
		fmt.Printf("query: %s", err.Error())
	}
	db.Close()
}

func TestSSHAgent(*testing.T) {
	var agentClient agent.Agent
	var sshUser string = "leoly"
	// Establish a connection to the local ssh-agent

	osSSHAuthSock := os.Getenv("SSH_AUTH_SOCK")
	fmt.Printf("SSH_AUTH_SOCK:%v\n", osSSHAuthSock)
	if conn, err := net.Dial("unix", osSSHAuthSock); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
		fmt.Printf("get agnetClient:%#v, conn:%#v\n", agentClient, conn)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{},
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		signers, err := agentClient.Signers()
		if nil != err {
			fmt.Printf("get signer failed:%v\n", err)
			return
		}
		for idx, signer := range signers {
			fmt.Printf("%03v%20v%v\n", idx, signer.PublicKey().Type(), hex.EncodeToString(signer.PublicKey().Marshal()))
		}
		//sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	_ = sshConfig

}
