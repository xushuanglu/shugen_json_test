### Docker学习笔记_安装和使用Zookeeper


二、安装

       1、搜索镜像   sudo docker search zookeeper
       2、下载镜像   sudo docker pull zookeeper
       3、查看已下载镜像   sudo docker images zookeeper
       4、使用镜像    sudo docker run -p 2181:2181 --name myZookeeper --restart always -d zookeeper

       容器名称：myZookeeper
       5、查看已运行容器，查到已运行Zookeeper容器的ID   sudo docker ps -a

       6、进入容器，并修改其配置   sudo docker exec -it 94e4e83c4c5a /bin/bash

    说明：
           1)、94e4e83c4c5a是容器ID
           2)、出现bash-4.4#,表明进了容器