package logmanager

import (
	"io"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func Start(logpath, starttime string) {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.Stamp,
		NoColors:        true,
		FieldsOrder:     []string{"o"},
	})
	f, err := os.OpenFile(logpath+"/"+starttime+".log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("logmanager start fail:", err)
	}
	log.SetOutput(io.MultiWriter(f))
}
