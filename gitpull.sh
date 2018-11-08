#! /bin/bash
# 切换到项目目录
cd /data/soft/src/web_render_service/
branch=$1
if [  -z "$branch" ];then
     echo "Please input branch!"
     exit
fi
# 执行git命令
git pull origin $branch
