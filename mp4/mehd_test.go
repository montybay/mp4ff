package mp4_test

import (
	"testing"

	"github.com/montybay/mp4ff/mp4"
)

func TestMehd(t *testing.T) {
	mehd := &mp4.MehdBox{FragmentDuration: 1234}
	boxDiffAfterEncodeAndDecode(t, mehd)
}
