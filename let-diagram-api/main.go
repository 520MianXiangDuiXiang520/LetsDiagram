package main

import (
    "flag"
    `fmt`
    "github.com/gin-gonic/gin"
    `io`
    `log`
    "net/http"
    _ "net/http/pprof"
    `os`
    `time`
)

func init() {
	initialization()
}

func main() {
	debugFlag := flag.Bool("debug", true, "run with debug mode")
	flag.Parse()
	if *debugFlag {
		go func() {
			http.ListenAndServe("0.0.0.0:8899", nil)
		}()
		gin.SetMode(gin.DebugMode)
	} else {
        gin.SetMode(gin.ReleaseMode)
        n := time.Now()
        logFileName := fmt.Sprintf("./logs/let_diagram_api_%d_%d_%d.log", n.Year(), n.Month(), n.Day())
        logF, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            logF = os.Stdout
            log.Printf("can not open %s, err: %v", logFileName, err)
        }
        log.SetOutput(logF)
        gin.DefaultWriter = io.MultiWriter(logF)
	}
	engine := gin.Default()
	rootRoutes(engine)
	defer engine.Run(":8888")
}
