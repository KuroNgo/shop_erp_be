# Quản Lý Tài Chính

## Tổng Quan
Hệ thống quản lý tài chính bao gồm các mô hình để quản lý tài khoản tài chính, giao dịch, hóa đơn, ngân sách, thanh toán, báo cáo tài chính và thuế.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Accounts (Tài khoản tài chính)
| Trường           | Loại Dữ Liệu | Mô Tả                                |
|------------------|--------------|--------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                           |
| `account_name`   | `string`     | Tên tài khoản (ví dụ: Ngân hàng A)   |
| `account_number` | `string`     | Số tài khoản                         |
| `balance`        | `float64`    | Số dư hiện tại                       |
| `account_type`   | `string`     | Loại tài khoản (tiền mặt, ngân hàng) |
| `created_at`     | `time.Time`  | Thời gian tạo                        |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                   |

### 2. Transactions (Giao dịch tài chính)
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

### 3. TransactionCategories (Danh mục giao dịch)
| Trường           | Loại Dữ Liệu  | Mô Tả                                   |
|------------------|---------------|-----------------------------------------|
| `id`             | `ObjectID`    | Khóa chính                              |
| `category_name`  | `string`      | Tên danh mục (lương, chi phí)           |
| `category_type`  | `string`      | Loại danh mục (thu nhập, chi tiêu)      |

### 4. Invoices (Hóa đơn)
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

### 5. Budgets (Ngân sách)
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

### 6. Payments (Thanh toán)
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

### 7. FinancialReports (Báo cáo tài chính)
| Trường         | Loại Dữ Liệu  | Mô Tả                                           |
|----------------|---------------|-------------------------------------------------|
| `id`           | `ObjectID`    | Khóa chính                                      |
| `report_name`  | `string`      | Tên báo cáo                                     |
| `start_date`   | `time.Time`   | Ngày bắt đầu kỳ báo cáo                         |
| `end_date`     | `time.Time`   | Ngày kết thúc kỳ báo cáo                        |
| `generated_at` | `time.Time`   | Ngày báo cáo được tạo                           |
| `data`         | `interface{}` | Dữ liệu báo cáo (có thể lưu trữ dưới dạng JSON) |

### 8. Taxes (Thuế)
| Trường          | Loại Dữ Liệu | Mô Tả                                           |
|-----------------|--------------|-------------------------------------------------|
| `id`            | `ObjectID`   | Khóa chính                                      |
| `tax_name`      | `string`     | Tên loại thuế (VAT, Thuế thu nhập doanh nghiệp) |
| `rate`          | `float64`    | Tỷ lệ thuế                                      |
| `applicable_to` | `string`     | Phạm vi áp dụng (hóa đơn, giao dịch, lương)     |
| `created_at`    | `time.Time`  | Thời gian tạo                                   |
| `updated_at`    | `time.Time`  | Thời gian cập nhật                              |
