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
  # echo "Git拉取分支[$branch]不是当前分支[$nowBranch]!"

  # 拉取远程分支
  git fetch origin $branch:$branch
  if [ $? != 0 ];then
     echo "Git拉取分支[$branch]失败"
     exit
  fi

  # 切换远程分支
  git checkout $branch
  if [ $? != 0 ];then
     echo "Git切换远程分支[$branch]失败"
     exit
  fi
fi

# 2.执行git命令
git pull origin $branch
