package builder_pattern

import (
	"fmt"
	"testing"
)

func TestFancyLapTopDirector_DirectLapTop(t *testing.T) {

	t.Run("general test", func(t *testing.T) {
		director := &FancyLapTopDirector{
			builder: FancyLapTopBuilder{},
		}
		fmt.Println(director.DirectLapTop())
	})

}
