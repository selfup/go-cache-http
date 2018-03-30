if [ "$1" == "-d" ]
then
  APP_HOST=0.0.0.0 
else
  APP_HOST=localhost
fi

ab \
  -n 50000 \
  -c 300 \
  -k -v 1 \
  -H "Accept-Encoding: gzip, deflate" \
  -T "application/json" \
  -p ./scripts/bench.txt "http://$APP_HOST:8081/" > .results.log \

git checkout -- .results.log
