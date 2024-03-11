package function

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	//"handler/function/static/ausf/pkg/service"
	"handler/function/static/ausf/ausfin/ausflogger"

	"handler/function/static/ausf/ausfin/sbi/ueauthentication"
	"handler/function/static/ausf/pkg/factory"
	"handler/function/static/ausf/pkg/service"
	"net/http"

	"github.com/urfave/cli"

	logger_util "github.com/free5gc/util/logger"
	"github.com/free5gc/util/version"
)

var AUSF *service.AusfApp

func Handle(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	fmt.Println("Hello World!")
	path := r.URL.EscapedPath()
	fmt.Println(path)

	if strings.HasPrefix(path, "/main") {
		main()
		fmt.Println("main EXECUTED")
	} else if path == "/" {
		fmt.Println(w)
		c, _ := gin.CreateTestContext(w)
		fmt.Println(c)
		//fmt.Println(c.GetRawData())
		c.Request = r
		ueauthentication.HTTPUeAuthenticationsPost(c)
		fmt.Println("HTTPUeAuthenticationsPost EXECUTED")
	} else if strings.HasPrefix(path, "/suci") {
		fmt.Println("path")
		fmt.Println(path)
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		ueauthentication.HTTPUeAuthenticationsAuthCtxID5gAkaConfirmationPut(c)
		fmt.Println("HTTPUeAuthenticationsPost EXECUTED")

	} else {
		http.NotFound(w, r)
	}

	//ueauthentication.HTTPUeAuthenticationsAuthCtxID5gAkaConfirmationPut(c)

	/*
		if c.Writer.Status() == http.StatusCreated { // http.StatusCreated == 201

			ueauthentication.HTTPUeAuthenticationsAuthCtxID5gAkaConfirmationPut(c2)

		} else {
			fmt.Println(w, "La llamada a HTTPUeAuthenticationsPost no fue exitosa: Código de Estado %d", c.Writer.Status())
		}
		// Execute the main function of AUSF main()*/
}

func main() {
	defer func() {
		if p := recover(); p != nil {

			// Print stack for panic to log. Fatalf() will let program exit.
			ausflogger.MainLog.Fatalf("panic: %v\n%s", p, string(debug.Stack()))
		}
	}()

	app := cli.NewApp()
	fmt.Println("newapp executed")
	app.Name = "ausf"
	app.Usage = "5G Authentication Server Function (AUSF)"
	app.Action = action
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
		cli.StringSliceFlag{
			Name:  "log, l",
			Usage: "Output NF log to `FILE`",
		},
	}
	fmt.Println("conf set")

	if err := app.Run(os.Args); err != nil {
		fmt.Println("error set")
		ausflogger.MainLog.Errorf("AUSF Run error: %v\n", err)
	}
}

func action(cliCtx *cli.Context) error {
	fmt.Println("executing action")
	tlsKeyLogPath, err := initLogFile(cliCtx.StringSlice("log"))
	if err != nil {
		return err
	}

	ausflogger.MainLog.Infoln("AUSF version: ", version.GetVersion())

	cfg, err := factory.ReadConfig(cliCtx.String("config"))
	fmt.Println("config read")

	if err != nil {
		fmt.Println("error1")

		return err
	}
	factory.AusfConfig = cfg

	ausf, err := service.NewApp(cfg)
	fmt.Println("new app called")

	if err != nil {
		fmt.Println("error2")
		return err
	}
	AUSF = ausf
	fmt.Println("ausf instance")

	ausf.Start(tlsKeyLogPath)
	fmt.Println("start called")

	return nil
}

func initLogFile(logNfPath []string) (string, error) {
	fmt.Println("init log file executed")
	logTlsKeyPath := ""

	for _, path := range logNfPath {
		if err := logger_util.LogFileHook(ausflogger.Log, path); err != nil {
			return "", err
		}

		if logTlsKeyPath != "" {
			continue
		}

		nfDir, _ := filepath.Split(path)
		tmpDir := filepath.Join(nfDir, "key")
		if err := os.MkdirAll(tmpDir, 0o775); err != nil {
			ausflogger.InitLog.Errorf("Make directory %s failed: %+v", tmpDir, err)
			return "", err
		}
		_, name := filepath.Split(factory.AusfDefaultTLSKeyLogPath)
		logTlsKeyPath = filepath.Join(tmpDir, name)
	}

	return logTlsKeyPath, nil
}

/*
func Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ue-authentications" {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Hello World!")

		service.Start()
		Main_function()

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Invalid request type")
	}
}

func Main_function() {
	fmt.Println("Hola2")

}
*/
