#!/bin/bash

server_bin_name="gowebbenchmark"

. ./libs.sh

length=${#web_frameworks[@]}

test_result=()

test_web_framework()
{
  echo "testing web framework: $2"
  ./servers/bench_${2}/$server_bin_name -s ${3}ms > alloc.log 2>&1 &
  sleep 2

  throughput=$(wrk -t$cpu_cores -c$4 -d${test_duration} http://127.0.0.1:7030 -s pipeline.lua --latency -- /hello 16| grep Requests/sec | awk '{print $2}')
  echo "throughput: $throughput requests/second"
  test_result[$1]=$throughput

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

echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime-pipeline.csv
test_all 0 1000
echo "0 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime-pipeline.csv
test_all 10 1000
echo "10 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime-pipeline.csv
test_all 100 1000
echo "100 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime-pipeline.csv
test_all 500 1000
echo "500 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime-pipeline.csv



echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency-pipeline.csv
test_all 30 100
echo "100,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency-pipeline.csv
test_all 30 1000
echo "1000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency-pipeline.csv
test_all 30 5000
echo "5000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency-pipeline.csv

mv -f processtime-pipeline.csv ./testresults
mv -f concurrency-pipeline.csv ./testresults
./testresults/plot.sh
