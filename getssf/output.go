package getssf

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
)

func Output(tweet interface{}, count uint64) {

	switch t := tweet.(type) {
	case anaconda.Tweet:
		log.Printf("%#v\n", t.Text)
		//default:
		//	log.Printf("%#v\n", t)
		//case anaconda.FriendsList:
		//	log.Printf("%#v\n", t)
		//case anaconda.StatusDeletionNotice:
		//	log.Printf("%#v\n", t)
		//case anaconda.KScrubGeo:
		//	log.Printf("%#v\n", t)
		//case anaconda.Limit:
		//	log.Printf("%#v\n", t)
		//case anaconda.StatusWithheld:
		//	log.Printf("%#v\n", t)
		//case anaconda.UserWithheld:
		//	log.Printf("%#v\n", t)
		//case anaconda.Disconnect:
		//	log.Printf("%#v\n", t)
		//case anaconda.:
		//	log.Printf("%#v\n", t)
		//case anaconda.DirectMessage:
		//	log.Printf("%#v\n", t)
	}

}
