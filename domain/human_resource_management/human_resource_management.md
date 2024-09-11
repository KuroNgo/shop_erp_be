# Quản Lý Nhân Sự

## Tổng Quan
Hệ thống quản lý nhân sự bao gồm các mô hình để quản lý thông tin nhân viên, phòng ban, lương, chấm công, đơn xin nghỉ phép, hợp đồng lao động, phúc lợi, và đánh giá hiệu suất.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Employees (Nhân viên)
Stores personal and work-related information of employees, including their department and role within the company.

| Trường          | Loại Dữ Liệu | Mô Tả                                        |
|-----------------|--------------|----------------------------------------------|
| `id`            | `ObjectID`   | Khóa chính                                   |
| `first_name`    | `string`     | Tên                                          |
| `last_name`     | `string`     | Họ                                           |
| `email`         | `string`     | Địa chỉ email                                |
| `phone_number`  | `string`     | Số điện thoại                                |
| `date_of_birth` | `time.Time`  | Ngày sinh                                    |
| `address`       | `string`     | Địa chỉ                                      |
| `position`      | `string`     | Chức vụ (ví dụ: Nhân viên, Trưởng phòng)     |
| `department_id` | `ObjectID`   | Tham chiếu tới bảng Departments (khóa ngoại) |
| `start_date`    | `time.Time`  | Ngày bắt đầu làm việc                        |
| `end_date`      | `time.Time`  | Ngày kết thúc làm việc (nếu có)              |
| `created_at`    | `time.Time`  | Thời gian tạo                                |
| `updated_at`    | `time.Time`  | Thời gian cập nhật                           |

### 2. Departments (Phòng ban)
Holds details about company departments and the manager assigned to each department.

| Trường            | Loại Dữ Liệu | Mô Tả                                                          |
|-------------------|--------------|----------------------------------------------------------------|
| `id`              | `ObjectID`   | Khóa chính                                                     |
| `department_name` | `string`     | Tên phòng ban                                                  |
| `manager_id`      | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại) - quản lý phòng ban |
| `created_at`      | `time.Time`  | Thời gian tạo                                                  |
| `updated_at`      | `time.Time`  | Thời gian cập nhật                                             |

### 3. Salaries (Lương)
Tracks employee salaries, including bonuses, deductions, and net pay.

| Trường        | Loại Dữ Liệu | Mô Tả                                              |
|---------------|--------------|----------------------------------------------------|
| `id`          | `ObjectID`   | Khóa chính                                         |
| `employee_id` | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại)         |
| `base_salary` | `float64`    | Lương cơ bản                                       |
| `bonus`       | `float64`    | Thưởng                                             |
| `deductions`  | `float64`    | Các khoản khấu trừ (bảo hiểm, thuế)                |
| `net_salary`  | `float64`    | Lương thực nhận (base_salary + bonus - deductions) |
| `pay_date`    | `time.Time`  | Ngày trả lương                                     |
| `created_at`  | `time.Time`  | Thời gian tạo                                      |
| `updated_at`  | `time.Time`  | Thời gian cập nhật                                 |

### 4. Attendance (Chấm công)
Records daily attendance, including check-in/out times and hours worked.

| Trường           | Loại Dữ Liệu | Mô Tả                                      |
|------------------|--------------|--------------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                                 |
| `employee_id`    | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại) |
| `date`           | `time.Time`  | Ngày chấm công                             |
| `check_in_time`  | `time.Time`  | Thời gian vào làm                          |
| `check_out_time` | `time.Time`  | Thời gian tan làm                          |
| `hours_worked`   | `float64`    | Số giờ làm việc                            |
| `status`         | `string`     | Trạng thái (Đi làm, Nghỉ phép, Nghỉ bệnh)  |
| `created_at`     | `time.Time`  | Thời gian tạo                              |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                         |

### 5. LeaveRequests (Đơn xin nghỉ phép)
Manages employee leave requests, including type of leave and approval status.

| Trường        | Loại Dữ Liệu | Mô Tả                                                       |
|---------------|--------------|-------------------------------------------------------------|
| `id`          | `ObjectID`   | Khóa chính                                                  |
| `employee_id` | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại)                  |
| `leave_type`  | `string`     | Loại nghỉ phép (Nghỉ bệnh, Nghỉ phép năm, Nghỉ không lương) |
| `start_date`  | `time.Time`  | Ngày bắt đầu nghỉ                                           |
| `end_date`    | `time.Time`  | Ngày kết thúc nghỉ                                          |
| `status`      | `string`     | Trạng thái đơn (Đã duyệt, Chờ duyệt, Từ chối)               |
| `created_at`  | `time.Time`  | Thời gian tạo                                               |
| `updated_at`  | `time.Time`  | Thời gian cập nhật                                          |

### 6. Contracts (Hợp đồng lao động)
Stores details of employee contracts, including type, salary, and contract duration.

| Trường          | Loại Dữ Liệu | Mô Tả                                      |
|-----------------|--------------|--------------------------------------------|
| `id`            | `ObjectID`   | Khóa chính                                 |
| `employee_id`   | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại) |
| `contract_type` | `string`     | Loại hợp đồng (Thời vụ, Dài hạn, Thử việc) |
| `start_date`    | `time.Time`  | Ngày bắt đầu hợp đồng                      |
| `end_date`      | `time.Time`  | Ngày kết thúc hợp đồng                     |
| `salary`        | `float64`    | Mức lương trong hợp đồng                   |
| `status`        | `string`     | Trạng thái hợp đồng (Hiệu lực, Hết hạn)    |
| `created_at`    | `time.Time`  | Thời gian tạo                              |
| `updated_at`    | `time.Time`  | Thời gian cập nhật                         |

### 7. Benefits (Phúc lợi)
Tracks employee benefits like insurance and allowances.

| Trường         | Loại Dữ Liệu | Mô Tả                                                           |
|----------------|--------------|-----------------------------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                                      |
| `employee_id`  | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại)                      |
| `benefit_type` | `string`     | Loại phúc lợi (Bảo hiểm y tế, Bảo hiểm xã hội, Phụ cấp ăn trưa) |
| `amount`       | `float64`    | Giá trị phúc lợi                                                |
| `start_date`   | `time.Time`  | Ngày bắt đầu phúc lợi                                           |
| `end_date`     | `time.Time`  | Ngày kết thúc phúc lợi                                          |
| `created_at`   | `time.Time`  | Thời gian tạo                                                   |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                                              |

### 8. PerformanceReviews (Đánh giá hiệu suất)
Logs employee performance reviews, scores, and comments from reviewers.

| Trường              | Loại Dữ Liệu | Mô Tả                                          |
|---------------------|--------------|------------------------------------------------|
| `id`                | `ObjectID`   | Khóa chính                                     |
| `employee_id`       | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại)     |
| `review_date`       | `time.Time`  | Ngày đánh giá                                  |
| `reviewer_id`       | `ObjectID`   | Tham chiếu tới bảng Employees (người đánh giá) |
| `performance_score` | `float64`    | Điểm hiệu suất                                 |
| `comments`          | `string`     | Nhận xét                                       |
| `created_at`        | `time.Time`  | Thời gian tạo                                  |
| `updated_at`        | `time.Time`  | Thời gian cập nhật                             |
