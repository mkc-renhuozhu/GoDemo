## this is a git command tutorial

### 1.repository  仓储

* remote repository : origin

#### 2.1创建一个本地仓储然后推送到git上面

* git init localReposiroty--这是一个本地的文件夹或者说是一个已经新建好了的本地项目

* 在git上面创建一个名称为localReposiroty这个的Repository

* git remote add localReposiroty https://github.com/mkc-renhuozhu/localReposiroty.git

* git push --set-upstream localReposiroty master

### 2.branch 分支

#### 2.1 创建一个名称为branch1的分支

* git branch branch1  

#### 2.2 切换到branch1分支

* git checkout branch1

### 3.commit  一次新的提交

### 4.githash   会产生对应的一次唯一的提交标示

#### 相关命令：

* git status

* git diff

* git add

* git commit -a -m

* git push

### 5.代码的回滚:

#### 回滚到指定的commit，显示变更的文件

* git reset  githash   

#### 回滚到指定的commit ,不显示变更的文件

* git reset --hard  githash   

#### 如果需要回滚到指定的commit并且提交的话：

##### 先针对这个分支做备份

* git branch branch1BackUp  

##### 回滚到指定的commit ,不显示变更的文件

* git reset --hard  githash   
 
##### 强制覆盖remote端代码

* git push -f  

### 6.tag

#### 标签，可以为某一个具体的commit打一个标记, 为以后可能会对当前版本做比较或者启用做记录

#### 创建tag

* git tag -a v1.0.1 36e0061849ec7ceeaf55f3a22ee3d49d7ad08aaa  -m 'add tag for current commit'

* git push origin v1.0.1

#### 删除tag

* git tag -d tagname 

### 7.merge  代码的合并

#### 将指定的githash合并到当前的branch

* git merge githash

#### 查看当前branch的log

* git log



HEAD 指的是最新的commit

因此  git reset --hart  HEAD  指的就是回滚到当前最新的代码