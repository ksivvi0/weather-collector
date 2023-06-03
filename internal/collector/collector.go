package collector

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"weather-collector/internal/config"
)

type Collector struct {
	Logger *logrus.Logger
	Config *config.Config
}

func NewCollector() (*Collector, error) {
	l := logrus.New()
	l.Out = os.Stdout

	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	return &Collector{Logger: logrus.New(), Config: cfg}, nil
}

func (c *Collector) Run(sigs chan os.Signal) {
	for {
		select {
		case s := <-sigs:
			c.Logger.Printf("Catch signal: %v", s)
			return
		default:
			time.Sleep(time.Second)
			c.Logger.Println("Hard work")
		}
	}
}
