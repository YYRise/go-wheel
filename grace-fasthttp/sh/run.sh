#!/bin/sh
RUN_LOG_PATH=./run.log
PROJECT_PATH=$(cd "$(dirname "$0")";cd ../; pwd)
BIN_NAME=grace-fasthttp-bin
BIN_PATH=$PROJECT_PATH/$BIN_NAME
GO_CMD=go
$GO_CMD build -o $BIN_PATH

function start()
{
    echo "启动:"
    pid=`ps -ef|grep $BIN_NAME | grep -v grep | awk '{print $2}'`
    if [ "$pid" != "" ]
    then
        echo "进程已存在:"$pid
        return
    fi
    echo "启动ing"
    /usr/bin/nohup $BIN_PATH >> $RUN_LOG_PATH 2>&1 &
    echo "启动成功"
}

function grace()
{
    echo "平滑重启:"
    pid=`ps -ef|grep $BIN_NAME | grep -v grep | awk '{print $2}'`
    if [ "$pid" != "" ]
    then
        echo "进程id:"$pid
        kill -HUP $pid
    else
        echo "进程不存在:"
        start
    fi
}

function restart()
{
    echo "重启:"
    pid=`ps -ef|grep $BIN_NAME | grep -v grep | awk '{print $2}'`
    if [ "$pid" != "" ]
    then
        echo "kill 进程:"$pid
        kill -9 $pid
        start
    else
        echo "进程不存在:"
        start
    fi
}



case "$1" in
    start)
        start
    ;;
    restart)
        restart
    ;;
    grace)
        grace
    ;;
esac
