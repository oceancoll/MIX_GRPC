# MIX_GRPC
begoo&amp;gomicro demo


#编译proto文件
./build_proto.sh

#使用方式及注意点
1.下载相关包文件

cd MIX_GRPC
go download

2.启动user-srv服务

go run main.go

3.启动sale-srv服务

go run main.go

4.安装begoo服务

go get github.com/astaxie/beego

go get github.com/beego/bee

bee添加环境变量

5.测试user-srv服务

go run /user-srv/client/main.go

6.测试sale-srv服务

go run /sale-srv/client/main.go

7.启动begoo服务

bee run

#docker化服务

1.分别在user-srv,sale-srv,beego-api中新增Dockerfile和Makefile文件，具体见文件

Makefile的build中，如果基础镜像是alpine,则在go build中增加CGO_ENABLED=0，防止动态编译

Makefile的run中，增加-v /etc/localtime:/etc/localtime:ro 可以同步宿主机的时区

2.关于新建user-client和sale-client的原因：

docker 本身有自己的mdns，user-srv run时，client.go会找不到节点，故client服务也需加入docker服务

3.具体启动方式

make build && make run
