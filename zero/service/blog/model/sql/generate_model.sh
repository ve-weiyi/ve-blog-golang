#-f 选项用于强制删除文件而不提示确认，-r 选项用于递归删除目录中的文件。
rm -f ../*
goctl model mysql ddl -src  test.sql -dir ../ --style go_zero -c --home ../../../../resource/template
