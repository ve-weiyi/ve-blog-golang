rm ../account_model.go
rm ../account_model_gen.go
goctl model mysql ddl -src  t_user.sql -dir ../ --style go_zero -c --home ../template
