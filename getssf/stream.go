package getssf

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
)

type Streams struct {
	stream []*anaconda.Stream
}

func (s *Streams) Add(stream *anaconda.Stream) {
	log.Print("called Streams add")
	s.stream = append(s.stream, stream)
}

func (s *Streams) Stop() {
	log.Print("called Streams stop")
	for _, stream := range s.stream {
		stream.Stop()
	}
}
