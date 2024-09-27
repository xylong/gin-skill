package bootstrap

import (
	"fmt"
	"gin-skill/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Migrate() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./app/dao",
		ModelPkgPath: "./app/models",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便
		FieldCoverable: false,

		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,

		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})

	g.UseDB(DB())

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		//"timestamp": func(detailType gorm.ColumnType) (dataType string) { return "int64" },           // 自定义时间
		"decimal": func(detailType gorm.ColumnType) (dataType string) { return "decimal.Decimal" }, // 金额类型全部转换为第三方库,github.com/shopspring/decimal
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("updatedAt", func(tag field.GormTag) field.GormTag {
		return tag.Append("autoUpdateTime")
	})
	autoCreateTimeField := gen.FieldGORMTag("createdAt", func(tag field.GormTag) field.GormTag {
		return tag.Append("autoCreateTime")
	})
	softDeleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField, autoUpdateTimeField, autoCreateTimeField, softDeleteField}
	//allModel := g.GenerateAllTable(fieldOpts...)
	//g.ApplyBasic(allModel...)

	generateModel(g, fieldOpts)
	g.Execute()
}

// DB 获取db
func DB() *gorm.DB {
	once.Do(func() {
		dbConfig := global.App.Config.Database

		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			dbConfig.UserName,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
			dbConfig.Charset,
		)

		connectDb(dsn)
	})

	return db
}

func connectDb(dsn string) *gorm.DB {
	var (
		err error
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger:                                   logger.Default.LogMode(logger.Info),
		Logger: getGormLogger(),
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(global.App.Config.Database.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(global.App.Config.Database.MaxOpenConns)
	mysqlDB.SetConnMaxLifetime(time.Second * 10)
	mysqlDB.SetConnMaxIdleTime(time.Second * 30)

	return db
}

func generateModel(generator *gen.Generator, fieldOpts []gen.ModelOpt) {
	Address := generator.GenerateModel("address")
	Profile := generator.GenerateModel("user_profiles")

	User := generator.GenerateModel("users",
		append(
			fieldOpts,
			// user 一对多 address 关联, 外键`user_id`在 address 表中
			gen.FieldRelate(field.HasMany, "Address", Address, &field.RelateConfig{GORMTag: field.GormTag{
				"foreignKey": []string{"UserID"},
				"references": []string{"ID"},
			}}),
			// user 一对一 user_profiles 关联, 外键`user_id`在 user_profile 表中
			gen.FieldRelate(field.HasOne, "Profile", Profile, &field.RelateConfig{GORMTag: field.GormTag{
				"foreignKey": []string{"UserID"},
				"references": []string{"ID"},
			}}),
		)...,
	)

	Address = generator.GenerateModel("address",
		append(
			fieldOpts,
			gen.FieldRelate(field.BelongsTo, "User", User, &field.RelateConfig{GORMTag: field.GormTag{
				"foreignKey": []string{"UserID"},
				"references": []string{"ID"},
			}}),
		)...,
	)

	generator.ApplyBasic(User, Address)
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.Dir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // 禁用彩色打印
	})
}
