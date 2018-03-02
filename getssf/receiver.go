package getssf

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

func worker(tweet chan<- interface{}, stream *anaconda.Stream) {

	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			default:
				tweet <- status
			}
		}
	}

}

func followReceiver(streams *Streams, api *anaconda.TwitterApi, follows []string, tweet chan<- interface{}) {

	v := url.Values{}
	if len(follows) > STATUSFILTERFOLLOW {
		panic("ups!!! follows count should be under STATUSFILTERFOLLOW")
	}
	follow := follows[0]
	for i := 1; i < len(follows); i++ {
		follow += "," + follows[i]
	}
	log.Printf("follow is %s\n", follow)
	v.Set("follow", follow)
	stream := api.PublicStreamFilter(v)
	streams.Add(stream)

	go worker(tweet, stream)

}

func trackReceiver(streams *Streams, api *anaconda.TwitterApi, tracks []string, tweet chan<- interface{}) {

	v := url.Values{}
	if len(tracks) > STATUSFILTERTRACK {
		panic("ups!!! tracks count should be under STATUSFILTERTRACK")
	}
	track := tracks[0]
	for i := 1; i < len(tracks); i++ {
		track += "," + tracks[i]
	}
	log.Printf("tracks is %s\n", track)
	v.Set("track", track)
	stream := api.PublicStreamFilter(v)
	streams.Add(stream)

	go worker(tweet, stream)

}

func CreateReceivers(streams *Streams, tweet chan<- interface{}, vKeys *viper.Viper, vAccounts *viper.Viper) {

	log.Printf("called createWotkers")

	tList := vAccounts.GetStringSlice("track")
	log.Printf("tList %s", tList)
	for i := 0; i < len(tList); i += STATUSFILTERTRACK {

		api := anaconda.NewTwitterApi(vKeys.GetString("TWITTER_OAUTH_TOKEN"), vKeys.GetString("TWITTER_OAUTH_TOKEN_SECRET"))
		api.SetLogger(anaconda.BasicLogger)

		end := i + STATUSFILTERTRACK
		if end > len(tList) {
			end = len(tList)
		}
		log.Printf("tList i %d\n", i)
		log.Printf("tList end %d\n", end)
		trackReceiver(streams, api, tList[i:end], tweet)

	}

	fList := vAccounts.GetStringSlice("follow")
	for i := 0; i < len(fList); i += STATUSFILTERFOLLOW {

		api := anaconda.NewTwitterApi(vKeys.GetString("TWITTER_OAUTH_TOKEN"), vKeys.GetString("TWITTER_OAUTH_TOKEN_SECRET"))
		api.SetLogger(anaconda.BasicLogger)

		end := i + STATUSFILTERFOLLOW
		if end > len(fList) {
			end = len(fList)
		}
		log.Printf("fList i %d\n", i)
		log.Printf("fList end %d\n", end)
		followReceiver(streams, api, fList[i:end], tweet)

	}

}
