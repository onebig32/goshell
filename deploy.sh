#! /bin/bash
# 切换到项目目录
WORK_DIR="/data/www/go/src/$1/"
if [ ! -d $WORK_DIR ]; then
       echo "Service not found!"
       exit
fi
cd $WORK_DIR
branch=$2
if [  -z "$branch" ];then
     echo "Please input branch!"
     exit
fi


# 1.判断拉取是否当前分支
nowBranch=`git symbolic-ref --short -q HEAD`
if [ "$nowBranch" != "$branch" ];then
  echo "Git拉取分支[$branch]不是当前分支[$nowBranch]!"
  exit
fi

# 2.执行deploy脚本
deployFile="${WORK_DIR}deploy.sh"
if [ ! -f $deployFile ]; then
       echo "deploy.sh not found!"
       exit
fi
/bin/sh deploy.sh

