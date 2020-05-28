
### git新建项目流程

>
*  git init    初始化git项目库

*  git add .   添加文件

*  git commit  提交文件

*  git remote add origin https://github.com/xxxxxxxxxx.git  添加文件到git地址

*  git push -f origin master     //如果是第一次提交的话，直接提交
>

### 教你如何修改github上的项目语言类型
>
什么是项目语言类型
使用github的朋友都知道，上传一个项目时都会显示项目的语言类型

如何修改语言类型
在你的项目根目录下创建一个.gitattributes文件

文件里加一行代码

*.js linguist-language=Vue
然后提交到github后，你就会发现javascript语言类型的项目变成了vue语言类型项目啦！

这样别人在搜索你的项目时，就可以快速定位查看你的项目了。
>


git之将主分支merge到自己分支

在进行git merge master时 git提示Already up to date.但其实并没有将最近代码拉下来。
百度后，找得到原因是master上的代码不是最新的
1 .切换到主干

$ git checkout master

更新主干代码
$ git pull origin master

切换回分支
$ git checkout dev

执行合并操作
$ git merge master

取消合并
$ git merge --abort

从远端拉取分支
$ git remote prune origin