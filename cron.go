package main

import (
	"time"

	"github.com/cuijxin/go-gin-example/models"
	"github.com/cuijxin/go-gin-example/pkg/logging"
	"github.com/robfig/cron"
)

func main() {
	logging.Info("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		logging.Info("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("* * * * * *", func() {
		logging.Info("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
