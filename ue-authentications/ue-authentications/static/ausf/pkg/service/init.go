package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/sirupsen/logrus"

	"handler/function/static/ausf/ausfin/ausflogger"
	ausf_context "handler/function/static/ausf/ausfin/context"
	"handler/function/static/ausf/ausfin/sbi/consumer"
	"handler/function/static/ausf/ausfin/sbi/ueauthentication"
	"handler/function/static/ausf/pkg/factory"

	"github.com/free5gc/util/httpwrapper"
	logger_util "github.com/free5gc/util/logger"
)

type AusfApp struct {
	cfg     *factory.Config
	ausfCtx *ausf_context.AUSFContext
}

func NewApp(cfg *factory.Config) (*AusfApp, error) {
	ausf := &AusfApp{cfg: cfg}
	ausf.SetLogEnable(cfg.GetLogEnable())
	ausf.SetLogLevel(cfg.GetLogLevel())
	ausf.SetReportCaller(cfg.GetLogReportCaller())

	ausf_context.Init()
	ausf.ausfCtx = ausf_context.GetSelf()
	return ausf, nil
}

func (a *AusfApp) SetLogEnable(enable bool) {
	ausflogger.MainLog.Infof("Log enable is set to [%v]", enable)
	if enable && ausflogger.Log.Out == os.Stderr {
		return
	} else if !enable && ausflogger.Log.Out == ioutil.Discard {
		return
	}

	a.cfg.SetLogEnable(enable)
	if enable {
		ausflogger.Log.SetOutput(os.Stderr)
	} else {
		ausflogger.Log.SetOutput(ioutil.Discard)
	}
}

func (a *AusfApp) SetLogLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		ausflogger.MainLog.Warnf("Log level [%s] is invalid", level)
		return
	}

	ausflogger.MainLog.Infof("Log level is set to [%s]", level)
	if lvl == ausflogger.Log.GetLevel() {
		return
	}

	a.cfg.SetLogLevel(level)
	ausflogger.Log.SetLevel(lvl)
}

func (a *AusfApp) SetReportCaller(reportCaller bool) {
	ausflogger.MainLog.Infof("Report Caller is set to [%v]", reportCaller)
	if reportCaller == ausflogger.Log.ReportCaller {
		return
	}

	a.cfg.SetLogReportCaller(reportCaller)
	ausflogger.Log.SetReportCaller(reportCaller)
}

func (a *AusfApp) Start(tlsKeyLogPath string) {
	fmt.Println("start started")
	ausflogger.InitLog.Infoln("Server started")

	router := logger_util.NewGinWithLogrus(ausflogger.GinLog)
	ueauthentication.AddService(router)

	pemPath := factory.AusfDefaultCertPemPath
	keyPath := factory.AusfDefaultPrivateKeyPath
	sbi := factory.AusfConfig.Configuration.Sbi
	if sbi.Tls != nil {
		pemPath = sbi.Tls.Pem
		keyPath = sbi.Tls.Key
	}

	self := a.ausfCtx
	// Register to NRF
	profile, err := consumer.BuildNFInstance(self)
	if err != nil {
		ausflogger.InitLog.Error("Build AUSF Profile Error")
	}
	_, self.NfId, err = consumer.SendRegisterNFInstance(self.NrfUri, self.NfId, profile)
	if err != nil {
		ausflogger.InitLog.Errorf("AUSF register to NRF Error[%s]", err.Error())
	}
	fmt.Println("registered in nrf done")
	//addr := fmt.Sprintf("%s:%d", self.BindingIPv4, self.SBIPort)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	fmt.Println("notify send")
	go func() {
		fmt.Println("signalChannel")
		fmt.Println(signalChannel)

		defer func() {
			if p := recover(); p != nil {
				// Print stack for panic to log. Fatalf() will let program exit.
				ausflogger.InitLog.Fatalf("panic: %v\n%s", p, string(debug.Stack()))
			}
		}()

		<-signalChannel
		a.Terminate()
		os.Exit(0)
	}()
	fmt.Println("NewHttp2Server antes")

	server, err := httpwrapper.NewHttp2Server("0.0.0.0:80", tlsKeyLogPath, router)

	if server == nil {
		ausflogger.InitLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}

	if err != nil {
		ausflogger.InitLog.Warnf("Initialize HTTP server: +%v", err)
	}
	fmt.Println("1")

	serverScheme := factory.AusfConfig.Configuration.Sbi.Scheme
	fmt.Println("2")
	fmt.Println(serverScheme)
	//fmt.Println(server.inShutdown)
	fmt.Println(server.Addr)

	if serverScheme == "http" {
		fmt.Println("http")
		// AQUI ES EL PROBLEMA!!!!!
		//err = server.ListenAndServe()
		//print(err)
	} else if serverScheme == "https" {
		fmt.Println("https")
		err = server.ListenAndServeTLS(pemPath, keyPath)
	}
	fmt.Println("3")

	if err != nil {
		ausflogger.InitLog.Fatalf("HTTP server setup failed: %+v", err)
	}

}

func (a *AusfApp) Terminate() {
	ausflogger.InitLog.Infof("Terminating AUSF...")
	// deregister with NRF
	problemDetails, err := consumer.SendDeregisterNFInstance()
	if problemDetails != nil {
		ausflogger.InitLog.Errorf("Deregister NF instance Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		ausflogger.InitLog.Errorf("Deregister NF instance Error[%+v]", err)
	} else {
		ausflogger.InitLog.Infof("Deregister from NRF successfully")
	}

	ausflogger.InitLog.Infof("AUSF terminated")
}
