package md

import (
	"os"
	"testing"
	"time"

	"gopkg.in/logger.v1"
)

var mdUserApi *GoMdAPI

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	time.Sleep(2 * time.Second)
	tearDown()
	os.Exit(retCode)
}
func setUp() {

	log.SetOutputLevel(log.Ldebug)
	log.Debug("start")
	//api.InitStore()
	//as := api.GetStore()
	//go func() {
	//mdUserApi = api.NewTraderAPI("tcp://180.168.146.187:10000", "9999", "099035", "your password")
	mdUserApi = NewMdAPI("tcp://180.168.146.187:10010", "9999", "099037", "123456")
	spi := NewGoMdSimpleSPI()
	spi.InitCMdSPI()
	mdUserApi.RegisterSpi(spi)
	mdUserApi.RegisterFront()
	mdUserApi.Connect()

	go func() {
		mdUserApi.Join()
	}()

}
func tearDown() {
	mdUserApi.Release()
}

func TestGoMdApi_UserLogin(t *testing.T) {

	for {
		result, err := mdUserApi.UserLogin()
		if err != nil {
			//todo
		}
		if result == 0 {
			break
		}
		time.Sleep(2 * time.Second)
		result, err = mdUserApi.UserLogin()

	}
}
