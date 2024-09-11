# Quản Lý Bán Hàng và Phân Phối

## Tổng Quan
Hệ thống quản lý bán hàng và phân phối bao gồm các mô hình để quản lý thông tin khách hàng, sản phẩm, danh mục sản phẩm, đơn hàng, chi tiết đơn hàng, vận chuyển, thanh toán, hóa đơn và báo cáo bán hàng.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Customers (Khách hàng)
Stores information about customers, including their contact details (name, email, phone, address) and basic information for future reference and order management.

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
Contains details of the products available for sale, such as product name, description, price, and stock quantity. This helps track inventory and product offerings.

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
Organizes products into different categories for better product classification and easier management. It allows for filtering and categorizing similar products.

| Trường          | Loại Dữ Liệu  | Mô Tả                                    |
|-----------------|---------------|------------------------------------------|
| `id`            | `ObjectID`    | Khóa chính                               |
| `category_name` | `string`      | Tên danh mục                             |
| `description`   | `string`      | Mô tả danh mục                           |
| `created_at`    | `time.Time`   | Thời gian tạo                            |
| `updated_at`    | `time.Time`   | Thời gian cập nhật                       |

### 4. SalesOrders (Đơn hàng)
Tracks customer orders, including customer details, shipping information, total amount, and the status of the order (processing, shipped, or canceled).

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
Contains information about individual items in each order. It tracks the products, quantities, and unit prices associated with a specific order to provide more detailed order history.

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
Records shipment details, including the shipping method, dates, tracking information, and the status of the delivery to monitor the progress of the shipping process.

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
Manages payment details for orders, including the payment method, status (paid or unpaid), and amount paid, ensuring accurate financial records.

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
Tracks the invoices generated for each order, including payment deadlines, amounts due, and the current status (paid or unpaid). It ensures that invoicing and payments are properly managed.

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
Generates reports on sales performance, including total sales, top-selling products, and trends over a specific period. This helps with business analysis and decision-making.

| Trường                 | Loại Dữ Liệu | Mô Tả                                           |
|------------------------|--------------|-------------------------------------------------|
| `id`                   | `ObjectID`   | Khóa chính                                      |
| `report_date`          | `time.Time`  | Ngày tạo báo cáo                                |
| `total_sales`          | `float64`    | Tổng doanh thu trong khoảng thời gian nhất định |
| `top_selling_products` | `string`     | Danh sách sản phẩm bán chạy                     |
| `created_at`           | `time.Time`  | Thời gian tạo                                   |
| `updated_at`           | `time.Time`  | Thời gian cập nhật                              |
