# github地址
https://github.com/esnet/iperf/tags

# EXAMPLE

## 服务端开启
./iperf3 -s

## 客户端丽连接
./iperf3 -c 192.168.12.153 -p 5201 -t 40 -i 2
ip是服务器的ip，-p是服务器启动后选择的端口5201，-t是测试40s，-i是2s测试一次。