package builder

import (
	"github.com/cristian0193/golang-service-template/internal"
	"go.uber.org/zap"
)

func NewDatabase(log *zap.SugaredLogger) (*internal.Database, error) {
	db := internal.New(log)
	err := db.Open()
	return db, err
}
