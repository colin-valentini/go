package textproto

// This package is a simple demonstration of how to embed and decode textprotos
// for use in a Go program.

import (
	"embed"
	"log"
	"sync"

	pb "github.com/colin-valentini/go/internal/proto/example"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	//go:embed *.textproto
	embedded_textprotos embed.FS
	lang_textprotos     = map[pb.Language]string{
		pb.Language_LANGUAGE_ENGLISH: "english_greeting.textproto",
		pb.Language_LANGUAGE_SPANISH: "spanish_greeting.textproto",
		pb.Language_LANGUAGE_ITALIAN: "italian_greeting.textproto",
	}
)

var (
	parseOnce      sync.Once
	parsedRespones = map[pb.Language]*pb.GreetingResponse{}
)

func parseTextProtos(responses map[pb.Language]*pb.GreetingResponse) {
	for lang, file_name := range lang_textprotos {
		textpb, err := embedded_textprotos.ReadFile(file_name)
		if err != nil {
			log.Fatalf("failed to read embedded textproto for %s: %v", lang.String(), err)
		}
		res := &pb.GreetingResponse{}
		if err := prototext.Unmarshal(textpb, res); err != nil {
			log.Fatalf("failed to unmarshall textproto for %s: %v", lang.String(), err)
		}
		responses[lang] = res
	}
}

func GetCannedResponses() map[pb.Language]*pb.GreetingResponse {
	parseOnce.Do(func() {
		parseTextProtos(parsedRespones)
	})
	return parsedRespones
}
