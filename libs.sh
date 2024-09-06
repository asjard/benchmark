#!/bin/bash

# web_frameworks=("default" "atreugo" "beego" "bone" "chi" "denco" "don" "echo"  "fasthttp" \
# "fasthttp-routing" "fasthttp/router" "fiber" "gear" "gearbox" "gin" "goframe" "goji" "gorestful" \
# "gorilla" "gorouter" "gorouterfasthttp" "go-ozzo" "goyave" "httprouter" "httptreemux" "indigo" "lars" \
# "lion" "muxie" "negroni" "pat" "pulse" "pure" "r2router" "tango" "tiger" "tinyrouter" "violetear" \
# "vulcan" "webgo")

web_frameworks=("default" "beego" "gin" "echo" "fasthttp" "asjard" "go_zero" "go_chassis" "kratos")



ROOTDIR=$(cd $(dirname $0);pwd)
###############
cpu_cores=4
# 不要小于5s
test_duration=15s
################

# export ASJARD_CONF_DIR=${ROOTDIR}/servers/asjard/conf
# export GO_ZERO_CONF_FILE=${ROOTDIR}/servers/go_zero/etc/gozero-api.yaml
# export CHASSIS_HOME=${ROOTDIR}/servers/go_chassis
