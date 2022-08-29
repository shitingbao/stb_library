# github地址
https://github.com/esnet/iperf/tags

# EXAMPLE

## 服务端开启
作为一个测速点
./iperf3 -s

## 客户端连接
连接测速点，来测试本地的网速
./iperf3 -c 192.168.12.153 -p 5201 -t 40 -i 2
ip是服务器的ip，-p是服务器启动后选择的端口5201，-t是测试40s，-i是2s测试一次。

## 注意一个服务点主要作为测试的连接点，不同的测试点对应的效率可能会不同