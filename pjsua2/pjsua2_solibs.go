// +build shared

package pjsua2

/*
#cgo CPPFLAGS: -DPJ_AUTOCONF=1 -O2 -DPJ_IS_BIG_ENDIAN=0
#cgo CPPFLAGS: -DPJ_IS_LITTLE_ENDIAN=1
#cgo CPPFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CPPFLAGS: -I/usr/local/include -I/usr/include
#cgo CPPFLAGS: -I/usr/include/c++/5 -I/usr/include/x86_64-linux-gnu/c++/5
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -lpjsua2
#cgo LDFLAGS: -lpjsua
#cgo LDFLAGS: -lpjsip-ua
#cgo LDFLAGS: -lpjsip-simple
#cgo LDFLAGS: -lpjsip
#cgo LDFLAGS: -lpjmedia-codec
#cgo LDFLAGS: -lpjmedia-videodev
#cgo LDFLAGS: -lpjmedia-audiodev
#cgo LDFLAGS: -lpjmedia
#cgo LDFLAGS: -lpjnath
#cgo LDFLAGS: -lpjlib-util
#cgo LDFLAGS: -lsrtp
#cgo LDFLAGS: -lresample
#cgo LDFLAGS: -lgsmcodec
#cgo LDFLAGS: -lspeex
#cgo LDFLAGS: -lilbccodec
#cgo LDFLAGS: -lg7221codec
#cgo LDFLAGS: -lyuv
#cgo LDFLAGS: -lwebrtc
#cgo LDFLAGS: -lpj
#cgo LDFLAGS: -lssl -lcrypto -luuid -lm -lrt -lpthread -lasound
*/
import "C"
