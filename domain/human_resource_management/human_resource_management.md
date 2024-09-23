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

**Business Logic**
* Add Employee: Thêm mới thông tin nhân viên vào hệ thống, bao gồm các chi tiết cá nhân và công việc.
* Update Employee: Cập nhật thông tin của nhân viên khi có sự thay đổi.
* Terminate Employee: Đánh dấu nhân viên đã rời công ty bằng cách cập nhật ngày kết thúc công việc.
* Department Assignment: Đảm bảo nhân viên thuộc về một phòng ban cụ thể.
* Employee Transfer: Cho phép chuyển nhân viên giữa các phòng ban.

### 2. Departments (Phòng ban)
Holds details about company departments and the manager assigned to each department.

| Trường            | Loại Dữ Liệu | Mô Tả                                                          |
|-------------------|--------------|----------------------------------------------------------------|
| `id`              | `ObjectID`   | Khóa chính                                                     |
| `department_name` | `string`     | Tên phòng ban                                                  |
| `manager_id`      | `ObjectID`   | Tham chiếu tới bảng Employees (khóa ngoại) - quản lý phòng ban |
| `created_at`      | `time.Time`  | Thời gian tạo                                                  |
| `updated_at`      | `time.Time`  | Thời gian cập nhật                                             |

**Business Logic**
* Add Department: Tạo mới phòng ban và chỉ định quản lý.
* Update Department: Cập nhật thông tin phòng ban hoặc quản lý mới.
* Assign Manager: Đảm bảo mỗi phòng ban đều có quản lý.

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

**Business Logic**
* Calculate Net Salary: Tính lương thực nhận dựa trên lương cơ bản, thưởng, và khấu trừ.
* Generate Pay Slip: Tạo phiếu lương chi tiết cho mỗi kỳ trả lương.
* Update Salary: Thay đổi mức lương của nhân viên khi có điều chỉnh.
* Handle Bonuses/Deductions: Quản lý các khoản thưởng hoặc khấu trừ như bảo hiểm, thuế.
* Salary Payment Notification: Gửi thông báo tới nhân viên khi lương đã được trả.
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

**Business Logic**
* Record Attendance: Ghi nhận thời gian vào/ra làm việc của nhân viên mỗi ngày.
* Track Working Hours: Tính tổng số giờ làm việc dựa trên thời gian check-in và check-out.
* Attendance Status: Đánh dấu trạng thái của ngày làm việc (Đi làm, Nghỉ phép, Nghỉ bệnh).
* Overtime Calculation: Tính toán số giờ làm thêm và ghi nhận chúng.


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

**Business Logic**
* Submit Leave Request: Nhân viên có thể gửi đơn xin nghỉ phép với loại nghỉ cụ thể.
* Approve/Reject Leave: Quản lý phê duyệt hoặc từ chối đơn nghỉ phép.
* Leave Balance: Kiểm tra số ngày nghỉ phép còn lại của nhân viên.
* Track Leave Status: Quản lý trạng thái của các đơn nghỉ phép (Chờ duyệt, Đã duyệt, Từ chối).
* Leave Approval/Rejection Notification: Gửi thông báo khi đơn nghỉ phép được duyệt hoặc từ chối.
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

**Business Logic**
* Create/Update Contract: Tạo mới hoặc cập nhật hợp đồng lao động cho nhân viên.
* Contract Status: Theo dõi trạng thái hợp đồng (Hiệu lực, Hết hạn).
* Renew Contract: Gia hạn hợp đồng khi gần hết hạn.
* Terminate Contract: Kết thúc hợp đồng lao động khi hết hạn hoặc khi nhân viên rời công ty.
* Contract Expiry Reminder: Gửi nhắc nhở khi hợp đồng lao động gần hết hạn.
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

**Business Logic**
* Assign Benefits: Gán các phúc lợi như bảo hiểm, phụ cấp cho nhân viên.
* Update Benefits: Điều chỉnh các phúc lợi khi có thay đổi về chính sách công ty.
* Track Benefit Periods: Theo dõi thời gian bắt đầu và kết thúc của các phúc lợi.

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

**Business Logic**
* Submit Performance Review: Người quản lý thực hiện đánh giá hiệu suất của nhân viên.
* Assign Reviewer: Gán người đánh giá cho mỗi nhân viên.
* Calculate Performance Score: Tính toán và ghi nhận điểm hiệu suất.
* Track Review History: Lưu trữ lịch sử các đánh giá hiệu suất của nhân viên.
* Performance Review Reminder: Nhắc nhở về kỳ đánh giá hiệu suất.
* Role-based Access: Phân quyền truy cập dựa trên vai trò (Nhân viên, Quản lý, Admin).
* Data Privacy: Đảm bảo rằng chỉ người được phép mới có quyền truy cập thông tin nhạy cảm như lương và hợp đồng.

### 9. Candidates (Ứng viên)
Bảng này lưu trữ thông tin các ứng viên tiềm năng trong quá trình tuyển dụng. Sau khi được chấp nhận, ứng viên sẽ được chuyển thành nhân viên chính thức.

| Trường             | Loại Dữ Liệu | Mô Tả                                                                |
|--------------------|--------------|----------------------------------------------------------------------|
| `id`               | `ObjectID`   | Khóa chính                                                           |
| `first_name`       | `string`     | Tên ứng viên                                                         |
| `last_name`        | `string`     | Họ ứng viên                                                          |
| `email`            | `string`     | Địa chỉ email                                                        |
| `phone_number`     | `string`     | Số điện thoại                                                        |
| `position_applied` | `string`     | Vị trí ứng tuyển                                                     |
| `department_id`    | `ObjectID`   | Tham chiếu tới bảng Departments (khóa ngoại)                         |
| `resume`           | `string`     | Đường dẫn tới CV/Resume của ứng viên                                 |
| `interview_date`   | `time.Time`  | Ngày phỏng vấn                                                       |
| `status`           | `string`     | Trạng thái ứng tuyển (Phỏng vấn, Đang xem xét, Trúng tuyển, Từ chối) |
| `feedback`         | `string`     | Nhận xét từ nhà tuyển dụng về ứng viên                               |
| `created_at`       | `time.Time`  | Thời gian tạo                                                        |
| `updated_at`       | `time.Time`  | Thời gian cập nhật                                                   |

### Business Logic for Candidates

1. **Candidate Lifecycle Management**
    - **Add Candidate**: Thêm mới ứng viên với các thông tin cá nhân, vị trí ứng tuyển, và ngày phỏng vấn.
    - **Update Candidate Status**: Cập nhật trạng thái ứng viên trong suốt quá trình tuyển dụng (Phỏng vấn, Đang xem xét, Trúng tuyển, Từ chối).
    - **Attach Resume**: Lưu trữ hoặc cập nhật CV/Resume của ứng viên vào hệ thống.
    - **Interview Scheduling**: Lên lịch phỏng vấn, ghi chú về thời gian và người phỏng vấn.
    - **Feedback Collection**: Lưu trữ nhận xét từ nhà tuyển dụng sau mỗi vòng phỏng vấn.

2. **Candidate Conversion to Employee**
    - **Convert to Employee**: Khi ứng viên được tuyển dụng, chuyển toàn bộ thông tin sang bảng **Employees** và yêu cầu gán phòng ban, chức vụ chính thức.
    - **Automatic Onboarding**: Khởi tạo quy trình chào mừng nhân viên mới sau khi chuyển từ ứng viên sang nhân viên.