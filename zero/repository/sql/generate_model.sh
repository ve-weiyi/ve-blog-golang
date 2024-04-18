rm ../account_model.go
rm ../account_model_gen.go
rm ../model
goctl model mysql ddl -src  test.sql -dir ../model --style go_zero -c --home ../template
