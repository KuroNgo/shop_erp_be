package attendance_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionAttendance = "attendance"
)

type Attendance struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID   primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	Date         time.Time          `bson:"date" json:"date"`
	CheckInTime  time.Time          `bson:"check_in_time" json:"check_in_time"`
	CheckOutTime time.Time          `bson:"check_out_time" json:"check_out_time"`
	HoursWorked  int8               `bson:"hours_worked" json:"hours_worked"`
	Status       string             `bson:"status" json:"status"` // Example values: "Present", "Leave", "Sick"
	Reason       string             `bson:"reason" bson:"reason"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	EmailEmployee string    `bson:"employee" json:"employee"`
	Date          time.Time `bson:"date" json:"date"`
	CheckInTime   time.Time `bson:"check_in_time" json:"check_in_time"`
	CheckOutTime  time.Time `bson:"check_out_time" json:"check_out_time"`
	Status        string    `bson:"status" json:"status"` // Example values: "Present", "Leave", "Sick"
	Reason        string    `bson:"reason" json:"reason"`
}

type Output struct {
	Attendance Attendance                `bson:"attendance" json:"attendance"`
	Employee   employees_domain.Employee `bson:"employee" json:"employee"`
}
