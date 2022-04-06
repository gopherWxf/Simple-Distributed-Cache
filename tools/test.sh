#----------------基于HTTP----------------#
# 查看状态
curl 127.0.0.1:12345/status
# 插入一个kv
curl -v  127.0.0.1:12345/cache/testkey -XPUT -d testvalue
# 查看key的val
curl 127.0.0.1:12345/cache/testkey
# 查看状态
curl 127.0.0.1:12345/status
# 删除key
curl 127.0.0.1:12345/cache/testkey -XDELETE
# 查看状态
curl 127.0.0.1:12345/status
#----------------基于TCP----------------#
../client/client.exe -c set -k testkey -v testvalue
../client/client.exe -c get -k testkey
curl 127.0.0.1:12345/status
../client/client.exe -c del -k testkey
curl 127.0.0.1:12345/status

go run main.go -n 10.29.1.1
go run main.go -n 10.29.1.2 -c 10.29.1.1
go run main.go -n 10.29.1.3 -c 10.29.1.2
curl 10.29.1.1:12345/cluster
curl 10.29.1.1:12345/status
curl 10.29.1.2:12345/status
curl 10.29.1.3:12345/status

../cache-benchmark/cache-benchmark.exe -type tcp -n 100000 -d 1 --h 10.29.1.1

../client/client.exe -c set -k keya -v a -h 10.29.1.3
../client/client.exe -c set -k keyb -v b -h 10.29.1.3
../client/client.exe -c set -k keyc -v c -h 10.29.1.3
../client/client.exe -c set -k keyd -v d -h 10.29.1.3
../client/client.exe -c set -k keye -v e -h 10.29.1.3
../client/client.exe -c set -k keyf -v f -h 10.29.1.3
../client/client.exe -c get -k keya -h 10.29.1.1

curl 10.29.1.1:12345/rebalance -XPOST

