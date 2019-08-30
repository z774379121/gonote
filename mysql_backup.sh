
#!/bin/sh
# mysql_backup.sh: 备份mysql数据信息，并且只保留最新的10份.
    #数据库用户名
    db_user="root"
    #数据库密码
    db_passwd="passwd"
    #数据库IP
    db_host="localhost"
    #数据库名
    db_name="db"
 
    #数据库备份信息保存位置.
 
    backup_dir="/home/jj/backup/mysql"
 
    #文件保存日期格式 (dd-mm-yyyy)
 
    time="$(date +"%d-%m-%Y")"
 
    # mysql, mysqldump所在目录，不同的安装会有不同目录
    #如果只对数据进行备份，可以不用填写
    MYSQL="/usr/local/mysql/bin/mysql"
 
    MYSQLDUMP="/usr/bin/mydumper"
 
    MKDIR="/bin/mkdir"
 
    RM="/bin/rm"
 
    MV="/bin/mv"
 
    GZIP="/bin/gzip"
 
    #检查备份目录 不存在进行存储主目录创建
 
    test ! -d $backup_dir && $MKDIR "$backup_dir"
 
    # 检查备份目录 不存在进行存储副目录创建
 
    test ! -d "$backup_dir/backup.0/" && $MKDIR "$backup_dir/backup.0/"
 
    # 获取所有的数据库信息
 
    #all_db="$($MYSQL -u $db_user -h $db_host -p$db_passwd -Bse 'show databases')"
 
    #for db in $all_db
 
    #do
 
    # $MYSQLDUMP -u $db_user -h $db_host -p$db_passwd $db_name | $GZIP -9 > "$backup_dir/backup.0/$time.$db_name.gz"

    $MYSQLDUMP -u $db_user -h $db_host -p $db_passwd -B $db_name -o "$backup_dir/backup.0/$time.$db_name"
    # 压缩
    # /usr/bin/tar -czvf "$backup_dir/backup.0/$time.$db_name.tar.gz" "$backup_dir/backup.0/$time.$db_name"
    # $RM -rf "$backup_dir/backup.0/$time.$db_name"

    # | $GZIP -9 > "$backup_dir/backup.0/$time.$db_name.gz"
 
    #done
 
    # 删除旧的备份信息
 
    test -d "$backup_dir/backup.10/" && $RM -rf "$backup_dir/backup.10"
 
    # rotate backup directory
 
    for int in 9 8 7 6 5 4 3 2 1 0
 
    do
 
    if(test -d "$backup_dir"/backup."$int")
 
    then
 
    next_int=`expr $int + 1`
 
    $MV "$backup_dir"/backup."$int" "$backup_dir"/backup."$next_int"
 
    fi
 
    done
 
    exit 0;