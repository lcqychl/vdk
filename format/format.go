package format

import (
	"github.com/lcqychl/vdk/av/avutil"
	"github.com/lcqychl/vdk/format/aac"
	"github.com/lcqychl/vdk/format/flv"
	"github.com/lcqychl/vdk/format/mp4"
	"github.com/lcqychl/vdk/format/rtmp"
	"github.com/lcqychl/vdk/format/rtsp"
	"github.com/lcqychl/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
