# Quản Lý Tài Chính

## Tổng Quan
Hệ thống quản lý tài chính bao gồm các mô hình để quản lý tài khoản tài chính, giao dịch, hóa đơn, ngân sách, thanh toán, báo cáo tài chính và thuế.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Accounts (Tài khoản tài chính)
Stores information about financial accounts like bank accounts or cash wallets, tracking balance and account type.

| Trường           | Loại Dữ Liệu | Mô Tả                                |
|------------------|--------------|--------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                           |
| `account_name`   | `string`     | Tên tài khoản (ví dụ: Ngân hàng A)   |
| `account_number` | `string`     | Số tài khoản                         |
| `balance`        | `float64`    | Số dư hiện tại                       |
| `account_type`   | `string`     | Loại tài khoản (tiền mặt, ngân hàng) |
| `created_at`     | `time.Time`  | Thời gian tạo                        |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                   |

**Business Logic:**
* Quản lý các tài khoản tài chính của công ty bao gồm tài khoản ngân hàng, ví điện tử, tài khoản tiền mặt, v.v.
* Cung cấp thông tin về số dư hiện tại của từng tài khoản, giúp doanh nghiệp theo dõi được lượng tiền mặt có sẵn.
* Cho phép ghi nhận và quản lý các giao dịch tài chính dựa trên tài khoản để kiểm soát nguồn tiền vào/ra.
* Hỗ trợ kiểm tra số dư tài khoản trước khi thực hiện thanh toán hoặc giao dịch lớn.

**Use Case:**

* Kiểm tra số dư tài khoản ngân hàng trước khi duyệt thanh toán lương cho nhân viên.
* Ghi nhận giao dịch thu nhập vào tài khoản sau khi khách hàng thanh toán hóa đơn.

### 2. Transactions (Giao dịch tài chính)
Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.

| Trường             | Loại Dữ Liệu | Mô Tả                                                  |
|--------------------|--------------|--------------------------------------------------------|
| `id`               | `ObjectID`   | Khóa chính                                             |
| `transaction_date` | `time.Time`  | Ngày giao dịch                                         |
| `account_id`       | `ObjectID`   | Tham chiếu tới bảng Accounts (khóa ngoại)              |
| `amount`           | `float64`    | Số tiền giao dịch                                      |
| `transaction_type` | `string`     | Loại giao dịch (thu nhập, chi tiêu)                    |
| `description`      | `string`     | Mô tả ngắn về giao dịch                                |
| `category_id`      | `ObjectID`   | Tham chiếu tới bảng TransactionCategories (khóa ngoại) |
| `created_at`       | `time.Time`  | Thời gian tạo                                          |
| `updated_at`       | `time.Time`  | Thời gian cập nhật                                     |

**Business Logic:**

* Theo dõi chi tiết từng giao dịch tài chính (thu nhập, chi tiêu) của công ty để nắm rõ nguồn tiền vào/ra.
* Cung cấp khả năng phân loại giao dịch để dễ dàng quản lý các loại thu nhập và chi tiêu.
* Hỗ trợ truy xuất lịch sử giao dịch để phục vụ cho việc kiểm toán và báo cáo tài chính.

**Use Case:**

* Ghi nhận giao dịch chi tiêu khi mua hàng hóa hoặc dịch vụ từ nhà cung cấp.
* Theo dõi giao dịch thu nhập khi nhận tiền từ khách hàng.

### 3. TransactionCategories (Danh mục giao dịch)
Categorizes transactions into types like income or expenses for better organization.

| Trường           | Loại Dữ Liệu  | Mô Tả                                   |
|------------------|---------------|-----------------------------------------|
| `id`             | `ObjectID`    | Khóa chính                              |
| `category_name`  | `string`      | Tên danh mục (lương, chi phí)           |
| `category_type`  | `string`      | Loại danh mục (thu nhập, chi tiêu)      |

**Business Logic:**

* Phân loại các giao dịch tài chính thành các nhóm như thu nhập và chi tiêu để phục vụ cho việc phân tích và báo cáo.
* Giúp doanh nghiệp có cái nhìn chi tiết hơn về các nguồn thu và các loại chi phí, từ đó dễ dàng điều chỉnh ngân sách và đưa ra quyết định kinh doanh.
* Hỗ trợ việc tạo các báo cáo theo danh mục để theo dõi tình hình tài chính theo từng loại giao dịch.

**Use Case:**

* Phân loại giao dịch thu nhập từ bán hàng vào danh mục "Doanh thu bán hàng."
* Phân loại chi phí văn phòng vào danh mục "Chi phí văn phòng."

### 4. Invoices (Hóa đơn)
Tracks issued invoices, including amounts, due dates, and payment status.

| Trường           | Loại Dữ Liệu | Mô Tả                                                |
|------------------|--------------|------------------------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                                           |
| `invoice_number` | `string`     | Số hóa đơn                                           |
| `invoice_date`   | `time.Time`  | Ngày phát hành hóa đơn                               |
| `customer_id`    | `ObjectID`   | Tham chiếu tới bảng Customers (khóa ngoại)           |
| `total_amount`   | `float64`    | Tổng số tiền hóa đơn                                 |
| `status`         | `string`     | Trạng thái (đã thanh toán, chưa thanh toán, quá hạn) |
| `due_date`       | `time.Time`  | Ngày đến hạn thanh toán                              |
| `created_at`     | `time.Time`  | Thời gian tạo                                        |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                                   |

**Business Logic:**

* Quản lý hóa đơn phát hành cho khách hàng, theo dõi số tiền cần thu và tình trạng thanh toán.
* Giúp doanh nghiệp kiểm soát dòng tiền từ khách hàng bằng cách theo dõi các hóa đơn đã phát hành, các hóa đơn đến hạn và hóa đơn quá hạn.
* Tích hợp với hệ thống thanh toán để đánh dấu hóa đơn là đã thanh toán khi nhận được tiền từ khách hàng.

**Use Case:**

* Phát hành hóa đơn cho khách hàng sau khi cung cấp dịch vụ.
* Theo dõi hóa đơn quá hạn để nhắc nhở khách hàng thanh toán.

### 5. Budgets (Ngân sách)
Manages budgets by setting limits for specific categories over a time period.

| Trường        | Loại Dữ Liệu | Mô Tả                                                  |
|---------------|--------------|--------------------------------------------------------|
| `id`          | `ObjectID`   | Khóa chính                                             |
| `budget_name` | `string`     | Tên ngân sách                                          |
| `amount`      | `float64`    | Số tiền ngân sách                                      |
| `start_date`  | `time.Time`  | Ngày bắt đầu ngân sách                                 |
| `end_date`    | `time.Time`  | Ngày kết thúc ngân sách                                |
| `category_id` | `ObjectID`   | Tham chiếu tới bảng TransactionCategories (khóa ngoại) |
| `created_at`  | `time.Time`  | Thời gian tạo                                          |
| `updated_at`  | `time.Time`  | Thời gian cập nhật                                     |

**Business Logic:**

* Quản lý ngân sách cho từng dự án, phòng ban, hoặc loại giao dịch tài chính để đảm bảo chi tiêu không vượt quá mức giới hạn.
* Giúp doanh nghiệp lập kế hoạch tài chính và kiểm soát chi tiêu trong các giai đoạn cụ thể.
* Cung cấp báo cáo so sánh giữa số tiền đã chi tiêu và ngân sách đã lập để đưa ra điều chỉnh kịp thời.

**Use Case:**

* Thiết lập ngân sách cho chiến dịch marketing trong quý 4 và theo dõi chi tiêu để đảm bảo không vượt ngân sách.
* Tạo ngân sách hàng năm cho từng phòng ban và theo dõi hiệu suất tài chính.

### 6. Payments (Thanh toán)
Records payments made towards invoices, linking them to accounts and payment methods.

| Trường           | Loại Dữ Liệu | Mô Tả                                           |
|------------------|--------------|-------------------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                                      |
| `payment_date`   | `time.Time`  | Ngày thanh toán                                 |
| `amount`         | `float64`    | Số tiền thanh toán                              |
| `payment_method` | `string`     | Phương thức thanh toán (tiền mặt, chuyển khoản) |
| `invoice_id`     | `ObjectID`   | Tham chiếu tới bảng Invoices (khóa ngoại)       |
| `account_id`     | `ObjectID`   | Tham chiếu tới bảng Accounts (khóa ngoại)       |
| `created_at`     | `time.Time`  | Thời gian tạo                                   |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                              |

**Business Logic:**

* Quản lý việc thanh toán hóa đơn, giúp theo dõi chính xác số tiền đã thanh toán, phương thức thanh toán và tài khoản thanh toán.
* Hỗ trợ đối chiếu các khoản thanh toán với hóa đơn để đảm bảo tính chính xác và tránh thanh toán trùng lặp.
* Cung cấp khả năng theo dõi các khoản thanh toán cho nhà cung cấp hoặc nhân viên.

**Use Case:**

* Thanh toán hóa đơn cho nhà cung cấp thông qua tài khoản ngân hàng.
* Ghi nhận thanh toán từ khách hàng và đối chiếu với hóa đơn đã phát hành.

### 7. FinancialReports (Báo cáo tài chính)
Generates financial reports over a specified date range for analysis and tracking.

| Trường         | Loại Dữ Liệu  | Mô Tả                                           |
|----------------|---------------|-------------------------------------------------|
| `id`           | `ObjectID`    | Khóa chính                                      |
| `report_name`  | `string`      | Tên báo cáo                                     |
| `start_date`   | `time.Time`   | Ngày bắt đầu kỳ báo cáo                         |
| `end_date`     | `time.Time`   | Ngày kết thúc kỳ báo cáo                        |
| `generated_at` | `time.Time`   | Ngày báo cáo được tạo                           |
| `data`         | `interface{}` | Dữ liệu báo cáo (có thể lưu trữ dưới dạng JSON) |

**Business Logic:**

* Tạo các báo cáo tài chính tổng hợp dựa trên dữ liệu giao dịch, ngân sách và các khoản thanh toán để phân tích hiệu quả tài chính của doanh nghiệp.
* Giúp ban lãnh đạo nắm bắt được tình hình tài chính qua các báo cáo về doanh thu, chi phí, lợi nhuận, dòng tiền, và nợ phải thu.
* Hỗ trợ việc đưa ra các quyết định chiến lược dựa trên báo cáo tài chính chi tiết và rõ ràng.

**Use Case:**

* Tạo báo cáo thu nhập ròng hàng tháng để đánh giá lợi nhuận của công ty.
* Xuất báo cáo tài chính hàng năm để phục vụ kiểm toán.

### 8. Taxes (Thuế)
Defines tax types and rates applicable to invoices, transactions, or other financial elements.

| Trường          | Loại Dữ Liệu | Mô Tả                                           |
|-----------------|--------------|-------------------------------------------------|
| `id`            | `ObjectID`   | Khóa chính                                      |
| `tax_name`      | `string`     | Tên loại thuế (VAT, Thuế thu nhập doanh nghiệp) |
| `rate`          | `float64`    | Tỷ lệ thuế                                      |
| `applicable_to` | `string`     | Phạm vi áp dụng (hóa đơn, giao dịch, lương)     |
| `created_at`    | `time.Time`  | Thời gian tạo                                   |
| `updated_at`    | `time.Time`  | Thời gian cập nhật                              |

**Business Logic:**

* Quản lý các loại thuế liên quan đến giao dịch, hóa đơn hoặc các khoản thu nhập khác của doanh nghiệp.
* Tự động tính toán thuế dựa trên các quy định pháp luật hiện hành và áp dụng cho từng hóa đơn hoặc giao dịch.
* Hỗ trợ doanh nghiệp trong việc lập báo cáo thuế và nộp thuế đúng hạn.

**Use Case:**

* Áp dụng thuế VAT cho hóa đơn bán hàng.
* Tính toán thuế thu nhập doanh nghiệp hàng năm.