rm -r ../model
goctl model mysql ddl -src  test.sql -dir ../model --style go_zero -c --home ../../../../script/template
