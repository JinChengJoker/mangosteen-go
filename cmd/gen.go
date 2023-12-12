package cmd

import (
	"mangosteen/database"

	"gorm.io/gen"
)

func genDao() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.UseDB(database.DB)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
