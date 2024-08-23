#!/bin/bash

server_bin_name="gowebbenchmark"

. ./libs.sh

length=${#web_frameworks[@]}

test_result=()
test_latency_result=()
test_alloc_result=()

###############
cpu_cores=4
# 不要小于10s
test_duration=15s
################

test_web_framework()
{
    index=$1
    framework=$2
    sleepTime=$3
    concurrency=$4
  echo "testing web framework: $2 $3"
  ./$server_bin_name $2 $3 > alloc.log 2>&1 &
  sleep 3
  ## 压测时间不得低于5秒
  wrk -t$cpu_cores -c$4 -d$test_duration http://127.0.0.1:7030/hello > tmp.log
  throughput=$(< "tmp.log" grep Requests/sec|awk '{print $2}')
  latency=$(< "tmp.log" grep Latency | awk '{print $2}')
  # 有可能是微秒
  if [ "${latency%us}" != "${latency}" ];then
    latencynano="${latency%us}"
    latency=$(awk 'BEGIN{printf "%.2f\n",('$latencynano'/'1000')}')
  else
    latency=${latency%ms}
  fi
  alloc=$(< "alloc.log" grep HeapAlloc | awk '{print $2}')
  echo "throughput: $throughput requests/second, latency: $latency ms, alloc: $alloc"
  test_result[$1]=$throughput
  test_latency_result[$1]=$latency
  test_alloc_result[$1]=$alloc

  pkill -9 $server_bin_name
  sleep 2
  echo "finished testing $2"
  echo
}

test_all()
{
  echo "###################################"
  echo "                                   "
  echo "      ProcessingTime  $1ms         "
  echo "      Concurrency     $2           "
  echo "                                   "
  echo "###################################"
  for ((i=0; i<length; i++))
  do
  	test_web_framework "$i" "${web_frameworks[$i]}" "$1" "$2"
  done
}


pkill -9 $server_bin_name

## 耗时测试
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime.csv
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime_latency.csv
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime_alloc.csv
test_all 0 100
echo "0 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
echo "0 ms,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> processtime_latency.csv
echo "0 ms,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> processtime_alloc.csv
test_all 10 100
echo "10 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
echo "10 ms,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> processtime_latency.csv
echo "10 ms,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> processtime_alloc.csv
test_all 100 100
echo "100 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
echo "100 ms,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> processtime_latency.csv
echo "100 ms,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> processtime_alloc.csv
test_all 500 100
echo "500 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
echo "500 ms,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> processtime_latency.csv
echo "500 ms,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> processtime_alloc.csv

## 并发测试
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency.csv
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency_latency.csv
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency_alloc.csv
test_all 0 100
echo "100,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
echo "100,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> concurrency_latency.csv
echo "100,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> concurrency_alloc.csv
test_all 0 1000
echo "1000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
echo "1000,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> concurrency_latency.csv
echo "1000,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> concurrency_alloc.csv
test_all 0 5000
echo "5000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
echo "5000,"$(IFS=$','; echo "${test_latency_result[*]}" ) >> concurrency_latency.csv
echo "5000,"$(IFS=$','; echo "${test_alloc_result[*]}" ) >> concurrency_alloc.csv



mv -f processtime.csv ./testresults
mv -f processtime_alloc.csv ./testresults
mv -f processtime_latency.csv ./testresults

mv -f concurrency.csv ./testresults
mv -f concurrency_latency.csv ./testresults
mv -f concurrency_alloc.csv ./testresults
./testresults/plot_mac.sh
