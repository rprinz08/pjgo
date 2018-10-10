# pjgo - using pjsip with go

This repository describes how to use the great, C/C++ based SIP library [pjsip](http://www.pjsip.org/) with [Go](https://golang.org/).

Pjsip provides a full featured library with almost everything to build Sip based communication software like for example softphones or Sip proxy servers. But this feature-richness comes with a price: projects might become complex quite fast. That's why pjsip provides the C++ based Object Oriented pjsua2 library which combines the lower level pjsip library functions into easy to use classes and components. This makes building Sip software much simpler.

Go makes it relatively easy to access native code by using [goc](https://golang.org/cmd/cgo/) and some special comments in your go source files. This works quite well for C but not so well for C++. So how could we access C++ based pjsua2?

That's where [SWIG](http://www.swig.org/) comes into play - the [Simplified Wrapper and Interface Generator](http://www.swig.org/). SWIG reads some metadata from interface description files **.i** or **.swig** and creates wrappers where higher level languages (e.g. Python) can call lower level pjsip functions. pjsip ships with predefined wrappers for C#, Python and Java.

Fortunately the pjsip source tree contains a SWIG definition for pjsua2 ([see github](https://github.com/pjsip/pjproject/tree/master/pjsip-apps/src/swig)). With this, SWIG can create a Go wrapper for pjsua2.

```sh
swig -c++ -go -cgo -intgosize 64 \
    -outcurrentdir \
    -I/usr/local/include \
    /path/to/pjsip/pjproject-2.8/pjsip-apps/src/swig/pjsua2.i
```

This generates files **pjsua2_wrap.cxx**, **pjsua2_wrap.h** and **pjsua2.go**. These files are not meant to be used by a programmer directly (almost unreadable generated Go code). Instead you import it into your Go project and can start using pjsip. As a last step you must decide if you want to link the static pjsip libraries (resulting in a larger binary) or against pjsip's shared libraries. This can be done by creating a go file (e.g. pjsua2_libs.go) in the same folder as the SWIG generated files which describes which pjsip libraries to link:

for static libraries:
```go
// +build !shared

package pjsua2

/*
#cgo CPPFLAGS: -DPJ_AUTOCONF=1 -O2 -DPJ_IS_BIG_ENDIAN=0
#cgo CPPFLAGS: -DPJ_IS_LITTLE_ENDIAN=1
#cgo CPPFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CPPFLAGS: -I/usr/local/include -I/usr/include
#cgo CPPFLAGS: -I/usr/include/c++/5 -I/usr/include/x86_64-linux-gnu/c++/5
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -lpjsua2-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsua-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-ua-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-simple-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-codec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-videodev-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-audiodev-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjnath-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjlib-util-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lsrtp-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lresample-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lgsmcodec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lspeex-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lilbccodec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lg7221codec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lyuv-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lwebrtc-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpj-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lssl -lcrypto -luuid -lm -lrt -lpthread -lasound
*/
import "C"
```

and for shared libraries:
```go
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
```

The pjsip documentation contains examples for C#, Python and Java. On in this repository you will find a go implementation of these examples.

It can be built either to use the static pjsip libraries with
```sh
go build
```
or to use the shared libraries with
```sh
go build -tags shared
```

See also my original blog post at https://www.min.at/prinz/?x=entry:entry180924-185225
