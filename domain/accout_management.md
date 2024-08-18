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

# Quản Lý Bán Hàng và Phân Phối

## Tổng Quan
Hệ thống quản lý bán hàng và phân phối bao gồm các mô hình để quản lý thông tin khách hàng, sản phẩm, danh mục sản phẩm, đơn hàng, chi tiết đơn hàng, vận chuyển, thanh toán, hóa đơn và báo cáo bán hàng.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Customers (Khách hàng)
| Trường          | Loại Dữ Liệu  | Mô Tả                                    |
|-----------------|---------------|------------------------------------------|
| `id`            | `ObjectID`    | Khóa chính                               |
| `first_name`    | `string`      | Tên                                      |
| `last_name`     | `string`      | Họ                                       |
| `email`         | `string`      | Địa chỉ email                            |
| `phone_number`  | `string`      | Số điện thoại                            |
| `address`       | `string`      | Địa chỉ                                  |
| `city`          | `string`      | Thành phố                                |
| `country`       | `string`      | Quốc gia                                 |
| `created_at`    | `time.Time`   | Thời gian tạo                            |
| `updated_at`    | `time.Time`   | Thời gian cập nhật                       |

### 2. Products (Sản phẩm)
| Trường              | Loại Dữ Liệu | Mô Tả                                       |
|---------------------|--------------|---------------------------------------------|
| `id`                | `ObjectID`   | Khóa chính                                  |
| `product_name`      | `string`     | Tên sản phẩm                                |
| `description`       | `string`     | Mô tả sản phẩm                              |
| `price`             | `float64`    | Giá bán                                     |
| `quantity_in_stock` | `int`        | Số lượng tồn kho                            |
| `category_id`       | `ObjectID`   | Tham chiếu tới bảng Categories (khóa ngoại) |
| `created_at`        | `time.Time`  | Thời gian tạo                               |
| `updated_at`        | `time.Time`  | Thời gian cập nhật                          |

### 3. Categories (Danh mục sản phẩm)
| Trường          | Loại Dữ Liệu  | Mô Tả                                    |
|-----------------|---------------|------------------------------------------|
| `id`            | `ObjectID`    | Khóa chính                               |
| `category_name` | `string`      | Tên danh mục                             |
| `description`   | `string`      | Mô tả danh mục                           |
| `created_at`    | `time.Time`   | Thời gian tạo                            |
| `updated_at`    | `time.Time`   | Thời gian cập nhật                       |

### 4. SalesOrders (Đơn hàng)
| Trường             | Loại Dữ Liệu | Mô Tả                                                  |
|--------------------|--------------|--------------------------------------------------------|
| `id`               | `ObjectID`   | Khóa chính                                             |
| `order_number`     | `string`     | Mã số đơn hàng                                         |
| `customer_id`      | `ObjectID`   | Tham chiếu tới bảng Customers (khóa ngoại)             |
| `order_date`       | `time.Time`  | Ngày đặt hàng                                          |
| `shipping_address` | `string`     | Địa chỉ giao hàng                                      |
| `total_amount`     | `float64`    | Tổng giá trị đơn hàng                                  |
| `status`           | `string`     | Trạng thái đơn hàng (Đang xử lý, Đã giao hàng, Đã hủy) |
| `created_at`       | `time.Time`  | Thời gian tạo                                          |
| `updated_at`       | `time.Time`  | Thời gian cập nhật                                     |

### 5. OrderDetails (Chi tiết đơn hàng)
| Trường        | Loại Dữ Liệu | Mô Tả                                        |
|---------------|--------------|----------------------------------------------|
| `id`          | `ObjectID`   | Khóa chính                                   |
| `order_id`    | `ObjectID`   | Tham chiếu tới bảng SalesOrders (khóa ngoại) |
| `product_id`  | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)    |
| `quantity`    | `int`        | Số lượng sản phẩm đặt mua                    |
| `unit_price`  | `float64`    | Giá của từng sản phẩm                        |
| `total_price` | `float64`    | Tổng giá (quantity * unit_price)             |
| `created_at`  | `time.Time`  | Thời gian tạo                                |
| `updated_at`  | `time.Time`  | Thời gian cập nhật                           |

### 6. Shipping (Vận chuyển)
| Trường               | Loại Dữ Liệu | Mô Tả                                                  |
|----------------------|--------------|--------------------------------------------------------|
| `id`                 | `ObjectID`   | Khóa chính                                             |
| `order_id`           | `ObjectID`   | Tham chiếu tới bảng SalesOrders (khóa ngoại)           |
| `shipping_method`    | `string`     | Phương thức vận chuyển (Tiêu chuẩn, Nhanh, Quốc tế)    |
| `shipping_date`      | `time.Time`  | Ngày vận chuyển                                        |
| `estimated_delivery` | `time.Time`  | Ngày dự kiến giao hàng                                 |
| `actual_delivery`    | `time.Time`  | Ngày giao hàng thực tế                                 |
| `tracking_number`    | `string`     | Mã số theo dõi vận chuyển                              |
| `status`             | `string`     | Trạng thái vận chuyển (Đang giao, Đã giao, Đã trả lại) |
| `created_at`         | `time.Time`  | Thời gian tạo                                          |
| `updated_at`         | `time.Time`  | Thời gian cập nhật                                     |

### 7. Payments (Thanh toán)
| Trường           | Loại Dữ Liệu | Mô Tả                                                                         |
|------------------|--------------|-------------------------------------------------------------------------------|
| `id`             | `ObjectID`   | Khóa chính                                                                    |
| `order_id`       | `ObjectID`   | Tham chiếu tới bảng SalesOrders (khóa ngoại)                                  |
| `payment_date`   | `time.Time`  | Ngày thanh toán                                                               |
| `payment_method` | `string`     | Phương thức thanh toán (Thẻ tín dụng, Chuyển khoản, Thanh toán khi nhận hàng) |
| `amount_paid`    | `float64`    | Số tiền đã thanh toán                                                         |
| `status`         | `string`     | Trạng thái thanh toán (Đã thanh toán, Chưa thanh toán)                        |
| `created_at`     | `time.Time`  | Thời gian tạo                                                                 |
| `updated_at`     | `time.Time`  | Thời gian cập nhật                                                            |

### 8. Invoices (Hóa đơn)
| Trường         | Loại Dữ Liệu | Mô Tả                                                        |
|----------------|--------------|--------------------------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                                   |
| `order_id`     | `ObjectID`   | Tham chiếu tới bảng SalesOrders (khóa ngoại)                 |
| `invoice_date` | `time.Time`  | Ngày xuất hóa đơn                                            |
| `due_date`     | `time.Time`  | Ngày đáo hạn thanh toán                                      |
| `amount_due`   | `float64`    | Số tiền cần thanh toán                                       |
| `amount_paid`  | `float64`    | Số tiền đã thanh toán                                        |
| `status`       | `string`     | Trạng thái hóa đơn (Đã thanh toán, Chưa thanh toán, Quá hạn) |
| `created_at`   | `time.Time`  | Thời gian tạo                                                |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                                           |

### 9. SalesReports (Báo cáo bán hàng)
| Trường                 | Loại Dữ Liệu | Mô Tả                                           |
|------------------------|--------------|-------------------------------------------------|
| `id`                   | `ObjectID`   | Khóa chính                                      |
| `report_date`          | `time.Time`  | Ngày tạo báo cáo                                |
| `total_sales`          | `float64`    | Tổng doanh thu trong khoảng thời gian nhất định |
| `top_selling_products` | `string`     | Danh sách sản phẩm bán chạy                     |
| `created_at`           | `time.Time`  | Thời gian tạo                                   |
| `updated_at`           | `time.Time`  | Thời gian cập nhật                              |

