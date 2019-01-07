if [ "$1" == "-d" ]
then
  APP_HOST=0.0.0.0 
else
  APP_HOST=localhost
fi

if [ "$2" == "" ]
then
  CONCURRENT_CALLS=300
  RESULTS_FILE=.con300.results.log
else
  CONCURRENT_CALLS=$2
  RESULTS_FILE=.con$2.results.log
fi

ab \
  -n 50000 \
  -c $CONCURRENT_CALLS \
  -k -v 1 \
  -H "Accept-Encoding: gzip, deflate" \
  -T "application/json" \
  -p ./scripts/bench.txt "http://$APP_HOST:8081/" > $RESULTS_FILE \
