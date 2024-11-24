package activity_log

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionActivityLog = "activity_log"
)

type ActivityLog struct {
	LogID        primitive.ObjectID `json:"_id" bson:"_id"`
	ClientIP     string             `json:"client_ip" bson:"client_ip"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	Method       string             `json:"method" bson:"method"`
	StatusCode   int                `json:"status_code" bson:"status_code"`
	BodySize     int                `json:"body_size" bson:"body_size"`
	Path         string             `json:"path" bson:"path"`
	Latency      string             `json:"latency" bson:"latency"`
	Error        string             `json:"error" bson:"error"`
	ActivityTime time.Time          `json:"activity_time" bson:"activity_time"`
	ExpireAt     time.Time          `json:"expire_at" bson:"expire_at"`
}
