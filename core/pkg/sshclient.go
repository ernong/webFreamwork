package pkg

import (
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/ssh"
)

// getSSHClinet ...
// https://github.com/helloyi/go-sshclient/blob/master/sshclient.go
func GetSSHClinet(remoteAddr string, user, privateKeyPath string) (sshClient *ssh.Client, retErr error) {
	//privateKeyPath := "/Users/leoly/.ssh/volume/vlm_cn.key"

	key, err := ioutil.ReadFile(privateKeyPath)
	if nil != err {
		retErr = fmt.Errorf("read key:%v failed:%v\n", privateKeyPath, err)
		return
	}
	signer, err := ssh.ParsePrivateKey(key)
	if nil != err {
		retErr = fmt.Errorf("parse private key failed:%v\n", err)
		return
	}

	config := ssh.ClientConfig{
		//Config:            ssh.Config{},
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remoteAddr net.Addr, pk ssh.PublicKey) error {
			fmt.Printf("HostKeyCallback: Hostname:[%v], Address:[%v], PK:[%v]\n",
				hostname, remoteAddr.String(), hex.EncodeToString(pk.Marshal()))
			return nil
		}),
		//BannerCallback:    nil,
		//ClientVersion:     "",
		//HostKeyAlgorithms: nil,
		//Timeout:           0,
	}

	//vlm1 := "120.78.70.77:22"
	sshClient, err = ssh.Dial("tcp", remoteAddr, &config)
	if nil != err {
		retErr = fmt.Errorf("dial to remote failed:%v\n", err)
		return
	}

	fmt.Printf("clinet server version:%s\n", sshClient.ServerVersion())
	return
}

// getRemoteDBOverSSH
func GetRemoteDBOverSSH(driver, dbUser, dbPasswd, dbHost, dbPort, dbName string, sshClient *ssh.Client) (*sqlx.DB, error) {
	dialFunc := func(addr string) (net.Conn, error) {
		return sshClient.Dial("tcp", addr)
	}

	connName := "tcpOverSSH"
	// connName = "tcp"
	mysql.RegisterDial(connName, dialFunc)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		dbUser, dbPasswd, connName, dbHost, dbPort, dbName)
	fmt.Printf("dsn:[%v]\n", dsn)
	driverName := "mysql"
	//dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	return sqlx.Connect(driverName, dsn)
	//return sqlx.Connect(driver, dsn)
	//return sql.Open(driver, dsn)
}

// GetRemoteGormDBOverSSH
func GetRemoteGormDBOverSSH(driver, dbUser, dbPasswd, dbHost, dbPort, dbName string, sshClient *ssh.Client) (*gorm.DB, error) {
	dialFunc := func(addr string) (net.Conn, error) {
		return sshClient.Dial("tcp", addr)
	}

	connName := "tcpOverSSH"
	// connName = "tcp"
	mysql.RegisterDial(connName, dialFunc)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		dbUser, dbPasswd, connName, dbHost, dbPort, dbName)
	fmt.Printf("dsn:[%v]\n", dsn)
	driverName := "mysql"
	//dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	//return sqlx.Connect(driverName, dsn)
	return gorm.Open(driverName, dsn)
	//return sqlx.Connect(driver, dsn)
	//return sql.Open(driver, dsn)
}
