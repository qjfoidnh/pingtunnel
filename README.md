# Pingtunnel-v6

修改自[https://github.com/cxjava/pingtunnel][pingtunnel]

pingtunnel是把tcp/udp/sock5流量伪装成icmp流量进行转发的工具。用于突破网络封锁，或是绕过WIFI网络的登陆验证，或是在某些网络加快网络传输速度。
pingtunnel-v6为原版pingtunnel的ipv6协议版本，不支持ipv4


# 功能
* 某些服务器的tcp、udp流量被禁止，可以通过pingtunnel绕过。
* 某些场合如学校、咖啡厅、机场，需要登录跳转验证，可以通过pingtunnel绕过。
* 某些网络，tcp、udp传输很慢，可以通过pingtunnel加速网络。

# 使用
### 编译
```
git clone https://github.com/qjfoidnh/pingtunnel.git
cd pingtunnel &&  go build -o pingtunnel cmd/pingtunnel/main.go
```

### 安装服务端
* 准备好一个具有公网ipv6的服务器，如AWS上的EC2，假定域名或者公网ip是www.yourserver.com。将编译出的的可执行文件传到服务器上
```
sudo ./pingtunnel -type server
```
* (可选)关闭系统默认的v6 ping，低版本内核可能不支持
```
echo 1 >/proc/sys/net/ipv6/icmp/echo_ignore_all
```

### 安装客户端(高玩推荐)
* 将编译出的可执行文件传到客户机
* 然后用**管理员权限**运行，不同的转发功能所对应的命令如下
* 如果看到有ping pong的log，说明连接正常
##### 转发sock5
```
# 如果使用ip地址，地址需要用[]括起来
pingtunnel.exe -type client -l :4455 -s www.yourserver.com -sock5 1
```
##### 转发tcp
```
pingtunnel.exe -type client -l :4455 -s www.yourserver.com -t www.yourserver.com:4455 -tcp 1
```
##### 转发udp
```
pingtunnel.exe -type client -l :4455 -s www.yourserver.com -t www.yourserver.com:4455
```
     
# Usage
    通过伪造ping，把tcp/udp/sock5流量通过远程服务器转发到目的服务器上。用于突破某些运营商封锁TCP/UDP流量。
    
    Usage:

    // server
    pingtunnel -type server

    // client, Forward udp
    pingtunnel -type client -l LOCAL_IP:4455 -s SERVER_IP -t SERVER_IP:4455

    // client, Forward tcp
    pingtunnel -type client -l LOCAL_IP:4455 -s SERVER_IP -t SERVER_IP:4455 -tcp 1

    // client, Forward sock5, implicitly open tcp, so no target server is needed
    pingtunnel -type client -l LOCAL_IP:4455 -s SERVER_IP -sock5 1

    -type     服务器或者客户端

    服务器参数:

    -key      设置的密码，默认0

    -nolog    不写日志文件，只打印标准输出，默认0

    -noprint  不打印屏幕输出，默认0

    -loglevel 日志文件等级，默认info

    -maxconn  最大连接数，默认0，不受限制

    -maxprt   server最大处理线程数，默认100

    -maxprb   server最大处理线程buffer数，默认1000

    -conntt   server发起连接到目标地址的超时时间，默认1000ms

    客户端参数:

    -l        本地的地址，发到这个端口的流量将转发到服务器

    -s        服务器的地址，流量将通过隧道转发到这个服务器

    -t        远端服务器转发的目的地址，流量将转发到这个地址

    -timeout  本地记录连接超时的时间，单位是秒，默认60s

    -key      设置的密码，默认0

    -tcp      设置是否转发tcp，默认0

    -tcp_bs   tcp的发送接收缓冲区大小，默认1MB

    -tcp_mw   tcp的最大窗口，默认10000

    -tcp_rst  tcp的超时发送时间，默认400ms

    -tcp_gz   当数据包超过这个大小，tcp将压缩数据，0表示不压缩，默认0

    -tcp_stat 打印tcp的监控，默认0

    -nolog    不写日志文件，只打印标准输出，默认0

    -noprint  不打印屏幕输出，默认0

    -loglevel 日志文件等级，默认info

    -sock5    开启sock5转发，默认0

    -profile  在指定端口开启性能检测，默认0不开启

    -s5filter sock5模式设置转发过滤，默认全转发，设置CN代表CN地区的直连不转发

    -s5ftfile sock5模式转发过滤的数据文件，默认读取当前目录的GeoLite2-Country.mmdb
