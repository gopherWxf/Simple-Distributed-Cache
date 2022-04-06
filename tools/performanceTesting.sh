#----------------基于HTTP----------------#
# 测试Set的性能，value长1000B，总请求100000
../cache-benchmark/cache-benchmark.exe -type http -n 100000 -r 100000 -t set
# 测试Get的性能
../cache-benchmark/cache-benchmark.exe -type http -n 100000 -r 100000 -t get
#----------------基于TCP----------------#
# 测试Set的性能，value长1000B，总请求100000
../cache-benchmark/cache-benchmark.exe -type tcp -n 100000 -r 100000 -t set
# 测试Get的性能
../cache-benchmark/cache-benchmark.exe -type tcp -n 100000 -r 100000 -t get

../cache-benchmark/cache-benchmark.exe -type tcp -n 100000 -d 1 --h 10.29.1.1