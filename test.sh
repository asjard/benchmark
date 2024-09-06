#!/bin/bash

server_bin_name="gowebbenchmark"

. ./libs.sh

length=${#web_frameworks[@]}

test_result=()

test_web_framework()
{
  echo "testing web framework: $2"
  args=""
  if [ $3 -gt 0 ];then
    args="$args -s ${3}ms"
  else
    args="$args -c true"
  fi
  ./servers/bench_${2}/$server_bin_name ${args} > alloc.log 2>&1 &
  sleep 2

  throughput=$(wrk -t$cpu_cores -c$4 -d${test_duration} http://127.0.0.1:7030/hello)
  echo "$throughput"
  test_result[$1]=$(echo "$throughput" | grep Requests/sec | awk '{print $2}')

  pkill -9 $server_bin_name
  sleep 1
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
  	test_web_framework $i ${web_frameworks[$i]} $1 $2
  done
}


pkill -9 $server_bin_name

echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime.csv
test_all 0 100
echo "0 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 10 100
echo "10 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 100 100
echo "100 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 500 100
echo "500 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv


echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency.csv
test_all 30 100
echo "100,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
test_all 30 1000
echo "1000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
test_all 30 5000
echo "5000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv


test_all -1 5000
echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > cpubound.csv
echo "cpu-bound,"$(IFS=$','; echo "${test_result[*]}" ) >> cpubound.csv

echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > cpubound-concurrency.csv
test_all -1 100
echo "100,"$(IFS=$','; echo "${test_result[*]}" ) >> cpubound-concurrency.csv
test_all -1 1000
echo "1000,"$(IFS=$','; echo "${test_result[*]}" ) >> cpubound-concurrency.csv
test_all -1 5000
echo "5000,"$(IFS=$','; echo "${test_result[*]}" ) >> cpubound-concurrency.csv


mv -f processtime.csv ./testresults
mv -f concurrency.csv ./testresults
mv -f cpubound.csv ./testresults
mv -f cpubound-concurrency.csv ./testresults
./testresults/plot.sh
