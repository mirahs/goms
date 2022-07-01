#!/bin/bash
DIR_ROOT=$(dirname $(readlink -f $0))/

# 发布目录
DIR_REL=/d/cy.worker/release/erlms/


fun_start_server()
{
    cd ${DIR_ROOT}bin/

    go build gomsserver

    ./gomsserver -t server -n 127.0.0.1
}

fun_start_client()
{
    cd ${DIR_ROOT}bin/

    go build gomsserver

    ./gomsserver -t client -n 127.0.0.1
}


fun_sync()
{
    cd ${DIR_REL}
    git pull

    cd ${DIR_ROOT}
    fun_rel

    for dirDep in $(ls deps)
    do
        mkdir ${DIR_REL}deps/${dirDep} -p
        \cp deps/${dirDep}/ebin ${DIR_REL}deps/${dirDep}/ -r
    done
    \cp ebin ${DIR_REL} -r
    \cp priv ${DIR_REL} -r
    \cp .gitignore_rel ${DIR_REL}.gitignore
    \cp ctl.sh ${DIR_REL}
    \cp elog.config ${DIR_REL}

    cd ${DIR_REL}
    git add .
    git commit -m 同步版本
    git push
}

fun_tar()
{
    cd ${DIR_ROOT}
    fun_rel

    DIR_TAR='./erlms/'
    mkdir ${DIR_TAR} -p
    rm -f erlms.tar.gz

    for dirDep in $(ls deps)
    do
        mkdir ${DIR_TAR}deps/${dirDep} -p
        \cp deps/${dirDep}/ebin ${DIR_TAR}deps/${dirDep}/ -r
    done
    \cp ebin ${DIR_TAR} -r
    \cp priv ${DIR_TAR} -r
    \cp ctl.sh ${DIR_TAR}
    \cp elog.config ${DIR_TAR}

    tar zcf erlms.tar.gz ${DIR_TAR}
    rm -rf ${DIR_TAR}
}


fun_help()
{
    echo "start_server          启动服务器"
    echo "start_client          启动客户端"

    echo "sync                  发布版本"
    echo "tar                   打包版本"

    exit 1
}



if [ $# -eq 0 ]
then
	fun_help
else
	fun_$1 $*
fi

exit 0
