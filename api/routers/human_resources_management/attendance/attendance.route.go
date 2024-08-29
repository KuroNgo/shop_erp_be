package attendance_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	attendancecontroller "shop_erp_mono/api/controllers/human_resources_management/attendance"
	"shop_erp_mono/bootstrap"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	attendancerepository "shop_erp_mono/repository/human_resource_management/attendance/repository"
	attendanceusecase "shop_erp_mono/usecase/human_resource_management/attendence/usecase"
	"time"
)

func AttendanceRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	at := attendancerepository.NewAttendanceRepository(db, attendancedomain.CollectionAttendance, employeesdomain.CollectionEmployee)
	attendance := &attendancecontroller.AttendanceController{
		AttendanceUseCase: attendanceusecase.NewAttendanceUseCase(timeout, at),
		Database:          env,
	}

	router := group.Group("/attendances")
	router.GET("/get/_id", attendance.FetchOneAttendanceByID)
	router.GET("/get/email", attendance.FetchOneAttendanceByEmail)
	router.GET("/get/all", attendance.FetchAllAttendance)
	router.POST("/create", attendance.CreateOneAttendance)
	router.PUT("/update", attendance.UpdateOneAttendance)
	router.DELETE("/delete", attendance.DeleteOneAttendance)
}
