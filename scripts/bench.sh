# this will run a concurrent benchmark on the server
# create a file called results (with the results)
# and then checkout the file from git to not pollute git history
cd ab_bench \
  && ./bench.sh \
  && cd .. \
  && git checkout .results
