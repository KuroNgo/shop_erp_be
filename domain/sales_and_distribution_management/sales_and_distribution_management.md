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

**Business Logic:**

* Quản lý thông tin liên lạc và thông tin cá nhân của khách hàng.
* Hỗ trợ theo dõi lịch sử mua hàng và cung cấp dịch vụ cá nhân hóa dựa trên hồ sơ khách hàng.
* Giúp doanh nghiệp quản lý chăm sóc khách hàng tốt hơn, từ đó nâng cao trải nghiệm khách hàng.

**Use Case:**

* Theo dõi các thông tin khách hàng để cá nhân hóa các chương trình khuyến mãi và dịch vụ.
* Quản lý các thông tin liên hệ của khách hàng cho mục đích bán hàng và marketing.

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

**Business Logic:**

* Quản lý chi tiết về sản phẩm, bao gồm mô tả, giá cả, và số lượng tồn kho.
* Hỗ trợ theo dõi tình trạng tồn kho để tránh việc thiếu hàng hoặc tồn đọng hàng hóa không cần thiết.
* Cung cấp dữ liệu sản phẩm để lập hóa đơn và theo dõi bán hàng.

**Use Case:**

* Theo dõi số lượng tồn kho để quản lý cung ứng sản phẩm hiệu quả.
* Quản lý thông tin giá bán sản phẩm cho các đơn hàng.

### 3. Categories (Danh mục sản phẩm)
Organizes products into different categories for better product classification and easier management. It allows for filtering and categorizing similar products.

| Trường          | Loại Dữ Liệu  | Mô Tả                                    |
|-----------------|---------------|------------------------------------------|
| `id`            | `ObjectID`    | Khóa chính                               |
| `category_name` | `string`      | Tên danh mục                             |
| `description`   | `string`      | Mô tả danh mục                           |
| `created_at`    | `time.Time`   | Thời gian tạo                            |
| `updated_at`    | `time.Time`   | Thời gian cập nhật                       |

**Business Logic:**

* Phân loại sản phẩm thành các danh mục để dễ dàng quản lý và tìm kiếm sản phẩm.
* Cung cấp khả năng lọc sản phẩm theo danh mục trong quá trình bán hàng và quản lý kho.

**Use Case:**

* Tạo danh mục sản phẩm mới khi công ty mở rộng dòng sản phẩm.
* Lọc sản phẩm theo danh mục để tìm kiếm và quản lý sản phẩm dễ dàng hơn.

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

**Business Logic:**

* Theo dõi các đơn hàng từ khách hàng, bao gồm chi tiết về khách hàng, địa chỉ giao hàng, và tổng giá trị đơn hàng.
* Hỗ trợ quản lý trạng thái đơn hàng để theo dõi quá trình từ khi đặt hàng đến khi giao hàng thành công.
* Cung cấp cơ sở dữ liệu để lập hóa đơn và thực hiện thanh toán.

**Use Case:**

* Ghi nhận đơn hàng mới và theo dõi trạng thái giao hàng.
* Tạo các đơn hàng trực tiếp từ thông tin khách hàng và sản phẩm đã có sẵn.

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

**Business Logic:**

* Ghi nhận chi tiết từng sản phẩm trong mỗi đơn hàng, bao gồm số lượng và giá đơn vị.
* Tính toán tổng giá trị của từng mục hàng trong đơn hàng để hỗ trợ việc thanh toán và lập hóa đơn.
* Cung cấp dữ liệu chi tiết để phân tích bán hàng và lập báo cáo.

**Use Case:**

* Theo dõi số lượng và giá của từng sản phẩm trong một đơn hàng cụ thể.
* Lập báo cáo chi tiết về các sản phẩm bán chạy dựa trên dữ liệu chi tiết đơn hàng.

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

**Business Logic:**

* Quản lý thông tin vận chuyển bao gồm phương thức, ngày giao hàng dự kiến, và trạng thái giao hàng.
* Giúp doanh nghiệp theo dõi quá trình vận chuyển, đảm bảo hàng hóa được giao đúng thời gian và đúng địa chỉ.
* Cung cấp khả năng kiểm tra trạng thái vận chuyển của đơn hàng theo thời gian thực.

**Use Case:**

* Cập nhật trạng thái vận chuyển khi đơn hàng được giao.
* Theo dõi đơn hàng bị trả lại hoặc thất lạc trong quá trình vận chuyển.

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

**Business Logic:**

* Ghi nhận thông tin thanh toán, bao gồm phương thức và trạng thái thanh toán.
* Hỗ trợ quản lý các khoản thanh toán chưa hoàn thành và theo dõi hạn mức thanh toán của khách hàng.
* Liên kết với hóa đơn và đơn hàng để đảm bảo thanh toán đúng số tiền và thời hạn.

**Use Case:**

* Ghi nhận thanh toán từ khách hàng và theo dõi trạng thái thanh toán của đơn hàng.
* Quản lý các đơn hàng chưa được thanh toán và nhắc nhở khách hàng về hạn thanh toán.

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

**Business Logic:**

* Tạo hóa đơn cho các đơn hàng đã hoàn thành, bao gồm thông tin về số tiền cần thanh toán và hạn thanh toán.
* Hỗ trợ theo dõi các hóa đơn chưa thanh toán, hóa đơn quá hạn và gửi nhắc nhở thanh toán đến khách hàng.
* Tích hợp với hệ thống thanh toán để đánh dấu hóa đơn là đã thanh toán khi nhận được tiền.

**Use Case:**

* Phát hành hóa đơn cho khách hàng sau khi giao hàng hoặc hoàn tất dịch vụ.
* Theo dõi các hóa đơn chưa thanh toán và gửi thông báo nhắc nhở cho khách hàng.

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

**Business Logic:**

* Tạo các báo cáo phân tích bán hàng bao gồm doanh thu, sản phẩm bán chạy, và xu hướng bán hàng theo thời gian.
* Cung cấp dữ liệu cho việc đánh giá hiệu suất kinh doanh và hỗ trợ ra quyết định chiến lược.
* Hỗ trợ xuất báo cáo theo nhiều định dạng để phục vụ việc kiểm toán và phân tích.

**Use Case:**

* Tạo báo cáo doanh thu hàng tháng để đánh giá tình hình kinh doanh.
* Theo dõi các sản phẩm bán chạy nhất trong quý để lập kế hoạch cung ứng phù hợp.