version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn : "root:root@tcp(localhost:3306)/genlearn?charset=utf8mb4&parseTime=true&loc=Local"
  # 选择mysql或者其他引擎，比方sqlserver
  db  : "mysql"
  # 指定要生成的table,流控则全部
  # 指定输出目录
  outPath :  "./dao/query"
  # 输出的代码，默认gen.go
  outFile :  ""
  # 是否生成单元测试
  withUnitTest  : false
  # generated model code's package name
  # 生成的model的代码的包名
  modelPkgName  : ""
  # 使用指针当字段是空的
  fieldNullable : false
  # 生成的字段带有gorm tag
  fieldWithIndexTag : true
  # 生成的字段时候带有gorm type 标签
  fieldWithTypeTag  : true
