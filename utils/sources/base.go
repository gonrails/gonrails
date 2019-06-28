package sources

import (
	"github.com/one-hole/gonrails/utils/sources/rabbits"
)

func init() {
	go rabbits.RunRabbit()
}
