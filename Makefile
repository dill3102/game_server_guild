
# schema.sql を generatorにコピる...

GENERATOR_PATH  = cmd/generator
# scheamからmodelを生成する
# go run $(GENERATOR_PATH)/sql_to_models/main.go
# go run cmd/generator/sql_to_models/main.go 


# modelからrepositoryの生成を行う
# go run $(GENERATOR_PATH)/repository/main.go $(GENERATOR_PATH)/sql_to_models/models $(GENERATOR_PATH)/templates/repository_template.go $(GENERATOR_PATH)/dist2
# go run cmd/generator/repository/main.go cmd/generator/sql_to_models/models cmd/generator/templates/repository_template.go cmd/generator/repository/dist2
