#!/bin/bash

swig -c++ -go -cgo -intgosize 64 \
	-outcurrentdir \
	-I/usr/local/include \
	~/projects/c/pjsip/pjproject-2.8/pjsip-apps/src/swig/pjsua2.i

# modify path above to match your installation.

