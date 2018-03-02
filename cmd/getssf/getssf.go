package main

import (
	"log"

	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"io"
	"os"

	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/stakada7/getssf/getssf"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {

	versionPrinted := flag.Bool("v", false, "getssf version")
	flag.Parse()

	if *versionPrinted {
		getssf.PrintVersion()
		return
	}

	o, err := os.OpenFile("output/filter.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("tweet.json ojson create missing")
	}
	defer o.Close()

	log.SetOutput(io.MultiWriter(o, os.Stdout))
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	runtime.SetBlockProfileRate(1)
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	vKeys := viper.New()
	vKeys.SetConfigType("yaml")
	vKeys.SetDefault("TWITTER_CONSUMER_KEY", "")
	vKeys.SetDefault("TWITTER_CONSUMER_SECRET", "")
	vKeys.SetDefault("TWITTER_OAUTH_TOKEN", "")
	vKeys.SetDefault("TWITTER_OAUTH_TOKEN_SECRET", "")
	vKeys.SetConfigName("keys")
	vKeys.AddConfigPath("./conf")
	err = vKeys.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading keys.yaml")
	}

	vAccounts := viper.New()
	vAccounts.SetConfigType("yaml")
	vAccounts.SetDefault("track", []string{})
	vAccounts.SetDefault("follow", []string{})
	vAccounts.SetConfigName("account")
	vAccounts.AddConfigPath("./conf")
	err = vAccounts.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading account.yaml")
	}

	anaconda.SetConsumerKey(vKeys.GetString("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(vKeys.GetString("TWITTER_CONSUMER_SECRET"))

	streams := &getssf.Streams{}
	tweet := make(chan interface{})

	// アカウントリストの再読込
	vAccounts.WatchConfig()
	vAccounts.OnConfigChange(func(e fsnotify.Event) {
		// 動いているstreamがあれば停止させる
		streams.Stop()
		// streamを稼働させる
		getssf.CreateReceivers(streams, tweet, vKeys, vAccounts)
	})

	getssf.CreateReceivers(streams, tweet, vKeys, vAccounts)

	// for debug counter
	var count uint64 = 1
	for {
		go getssf.Output(<-tweet, count)
		count++
	}

}
