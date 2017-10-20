# this will run a concurrent benchmark on the server
# create a file called results (with the results)
# and then assume unchanged the file from git to not pollute git history
git update-index --assume-unchanged .results

ab \
  -n 100000 \
  -c 50 \
  -k -v 1 \
  -H "Accept-Encoding: gzip, deflate" \
  -T "application/json" \
  -p ./scripts/bench.txt http://localhost:8080/ > .results \

# to re-track the .results file for updated benchmarks
# $ git update-index --no-assume-unchanged .results
# coomit the change and then
# git update-index --assume-unchanged .results
# prior to a new push
