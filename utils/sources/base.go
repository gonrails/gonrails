package sources

import (
	"github.com/gonrails/gonrails/utils/sources/rabbits"
)

func init() {
	go rabbits.RunRabbit()
}
