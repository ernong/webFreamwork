package main

import (
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"oceanEngineService/bus/db"
	"oceanEngineService/bus/route"
	"oceanEngineService/core/pkg"
	"oceanEngineService/core/util"
	"time"
)

var (
	cfg  = pflag.StringP("config", "c", "./conf/config.yaml", "server config file path")
	test = pflag.StringP("test", "t", "yes", "test flag")
)

func main() {

	pflag.Parse()

	pkg.Init(*cfg)
	db.InitDBFromCfg()
	util.InitGo(100)

	gin.SetMode(viper.GetString("run_mode"))

	g := gin.New()
	//store := cookie.NewStore([]byte("universal inc. ill sick blibli"))
	//g.Use(sessions.Sessions("session", store))

	route.Load(
		g,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the route has no response, or it might took too long to start up", err)
		}
		log.Info("the route has been deployed successfully")
	}()
	log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	http.ListenAndServe(viper.GetString("addr"), g)
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(viper.GetString("url") + "/friendcycle/api/check/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the route, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the route error")
}
