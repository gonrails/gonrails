package sources

import (
	"kalista/utils/sources/rabbits"
)

func init() {
	go rabbits.RunRabbit()
}
