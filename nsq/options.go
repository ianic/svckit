/* Postavljanje opcija pattern ukraden od: http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
Primjer:
	consumer := nsq.MustNewConsumer("listici.novi", handler,
		nsq.MaxInFlight(200),
		nsq.Channel("my-app#ephemeral"),
	)
*/
package nsq

import (
	"github.com/minus5/svckit/log"
	"strings"

	gonsq "github.com/nsqio/go-nsq"
)

type nsqLogger struct {
}

// Ovdje ulaze logovi iz go-nsq liba.
// Da ne idu na stderr postavim SetLogger na producer i consumer.
// Pa onda mogu nesto korisni napraviti s njima.
// Ref: https://github.com/nsqio/go-nsq/blob/0b80d6f05e15ca1930e0c5e1d540ed627e299980/delegates.go#L6
func (n *nsqLogger) Output(calldepth int, s string) error {
	a := log.NewAgregator(nil, calldepth)
	a.S("lib", "svckit.nsq.gonsq")
	if strings.HasPrefix(s, "INF") {
		a.Info(s)
		return nil
	}
	if strings.HasPrefix(s, "WRN") {
		a.Info(s)
		return nil
	}
	if strings.HasPrefix(s, "ERR") {
		if !strings.Contains(s, "TOPIC_NOT_FOUND") {
			a.ErrorS(s)
		}
		return nil
	}
	a.Debug(s)
	return nil
}

type options struct {
	maxInFlight      int
	channel          string
	lookupdHTTPAddrs []string
	nsqdTCPAddr      string
	logger           *nsqLogger
	logLevel         gonsq.LogLevel
}

func (c *options) apply(opts ...func(*options)) *options {
	for _, fn := range opts {
		fn(c)
	}
	return c
}

func MaxInFlight(m int) func(*options) {
	return func(o *options) {
		o.maxInFlight = m
	}
}

func Channel(c string) func(*options) {
	return func(o *options) {
		o.channel = c
	}
}