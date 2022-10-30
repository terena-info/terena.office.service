package configs

import (
	"context"
	"fmt"
	"time"

	"github.com/terena-info/terena.godriver/gomgo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (cfg *configs) ConnectDatabase() {
	connector := gomgo.ConnectionOption{
		Host:     Env.DB_URI,
		Database: Env.DB_NAME,
		ReadRef:  readpref.Primary(),
		Timeout:  time.Second * 10,
		Context:  context.Background(),
	}
	connector.Connect().WithMessage(fmt.Sprintf("Database: %s", Env.DB_URI))

}
