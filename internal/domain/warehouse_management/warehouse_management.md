# Quản Lý Kho

## Tổng Quan
Hệ thống quản lý kho bao gồm các mô hình để quản lý thông tin về kho lưu trữ, sản phẩm, tồn kho, di chuyển kho, nhà cung cấp, đơn mua hàng, chi tiết đơn mua hàng và điều chỉnh tồn kho.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Warehouses (Kho)
This table stores information about different warehouses, including their location, capacity, and basic details. It helps manage the distribution and storage of products across different facilities.

| Trường           | Loại Dữ Liệu | Mô Tả              |
|------------------|--------------|--------------------|
| `id`             | `ObjectID`   | Khóa chính         |
| `warehouse_name` | `string`     | Tên kho            |
| `location`       | `string`     | Vị trí của kho     |
| `capacity`       | `int`        | Công suất của kho  |
| `created_at`     | `time.Time`  | Thời gian tạo      |
| `updated_at`     | `time.Time`  | Thời gian cập nhật |

### 2. Products (Sản phẩm)
Contains details of products such as name, description, price, and associated categories. It is central to managing what items are available for sale and for inventory control.

| Trường         | Loại Dữ Liệu | Mô Tả                                       |
|----------------|--------------|---------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                  |
| `product_name` | `string`     | Tên sản phẩm                                |
| `description`  | `string`     | Mô tả sản phẩm                              |
| `price`        | `float64`    | Giá bán                                     |
| `category_id`  | `ObjectID`   | Tham chiếu tới bảng Categories (khóa ngoại) |
| `created_at`   | `time.Time`  | Thời gian tạo                               |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                          |

### 3. Inventory (Tồn kho)
Tracks the stock levels of products in different warehouses. It helps monitor how much stock is available and where it is located.

| Trường         | Loại Dữ Liệu | Mô Tả                                       |
|----------------|--------------|---------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                  |
| `product_id`   | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)   |
| `warehouse_id` | `ObjectID`   | Tham chiếu tới bảng Warehouses (khóa ngoại) |
| `quantity`     | `int`        | Số lượng tồn kho                            |
| `created_at`   | `time.Time`  | Thời gian tạo                               |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                          |

### 4. StockMovements (Di chuyển kho)
Records the movements of products in and out of warehouses, such as when items are received (Nhập kho) or dispatched (Xuất kho). This helps track changes in inventory levels.

| Trường          | Loại Dữ Liệu | Mô Tả                                           |
|-----------------|--------------|-------------------------------------------------|
| `id`            | `ObjectID`   | Khóa chính                                      |
| `product_id`    | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)       |
| `warehouse_id`  | `ObjectID`   | Tham chiếu tới bảng Warehouses (khóa ngoại)     |
| `movement_type` | `string`     | Loại di chuyển (Nhập kho, Xuất kho)             |
| `quantity`      | `int`        | Số lượng di chuyển                              |
| `movement_date` | `time.Time`  | Ngày di chuyển                                  |
| `reference`     | `string`     | Tham chiếu đến đơn hàng hoặc tài liệu liên quan |
| `created_at`    | `time.Time`  | Thời gian tạo                                   |
| `updated_at`    | `time.Time`  | Thời gian cập nhật                              |

### 5. Suppliers (Nhà cung cấp)
Stores supplier information, including contact details. This table is crucial for managing relationships with vendors and tracking where products are sourced.

| Trường           | Loại Dữ Liệu | Mô Tả              |
|------------------|--------------|--------------------|
| `id`             | `ObjectID`   | Khóa chính         |
| `supplier_name`  | `string`     | Tên nhà cung cấp   |
| `contact_person` | `string`     | Người liên hệ      |
| `phone_number`   | `string`     | Số điện thoại      |
| `email`          | `string`     | Địa chỉ email      |
| `address`        | `string`     | Địa chỉ            |
| `created_at`     | `time.Time`  | Thời gian tạo      |
| `updated_at`     | `time.Time`  | Thời gian cập nhật |

### 6. PurchaseOrders (Đơn mua hàng)
Manages purchase orders sent to suppliers. It includes order numbers, supplier references, order dates, and status, ensuring that product orders are properly tracked.

| Trường         | Loại Dữ Liệu | Mô Tả                                             |
|----------------|--------------|---------------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                        |
| `order_number` | `string`     | Mã số đơn hàng                                    |
| `supplier_id`  | `ObjectID`   | Tham chiếu tới bảng Suppliers (khóa ngoại)        |
| `order_date`   | `time.Time`  | Ngày đặt hàng                                     |
| `total_amount` | `float64`    | Tổng giá trị đơn hàng                             |
| `status`       | `string`     | Trạng thái đơn hàng (Đang xử lý, Đã nhận, Đã hủy) |
| `created_at`   | `time.Time`  | Thời gian tạo                                     |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                                |

### 7. PurchaseOrderDetails (Chi tiết đơn mua hàng)
Provides detailed information about the individual products in each purchase order, including quantities and prices. It links the products ordered from suppliers to their respective purchase orders.

| Trường              | Loại Dữ Liệu | Mô Tả                                           |
|---------------------|--------------|-------------------------------------------------|
| `id`                | `ObjectID`   | Khóa chính                                      |
| `purchase_order_id` | `ObjectID`   | Tham chiếu tới bảng PurchaseOrders (khóa ngoại) |
| `product_id`        | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)       |
| `quantity`          | `int`        | Số lượng sản phẩm                               |
| `unit_price`        | `float64`    | Giá của từng sản phẩm                           |
| `total_price`       | `float64`    | Tổng giá (quantity * unit_price)                |
| `created_at`        | `time.Time`  | Thời gian tạo                                   |
| `updated_at`        | `time.Time`  | Thời gian cập nhật                              |

### 8. StockAdjustments (Điều chỉnh tồn kho)
Keeps records of stock adjustments, such as increases or decreases in inventory due to reasons like damage, returns, or reallocation. This ensures accurate stock levels are maintained.

| Trường            | Loại Dữ Liệu | Mô Tả                                       |
|-------------------|--------------|---------------------------------------------|
| `id`              | `ObjectID`   | Khóa chính                                  |
| `product_id`      | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)   |
| `warehouse_id`    | `ObjectID`   | Tham chiếu tới bảng Warehouses (khóa ngoại) |
| `adjustment_type` | `string`     | Loại điều chỉnh (Tăng, Giảm)                |
| `quantity`        | `int`        | Số lượng điều chỉnh                         |
| `reason`          | `string`     | Lý do điều chỉnh                            |
| `adjustment_date` | `time.Time`  | Ngày điều chỉnh                             |
| `created_at`      | `time.Time`  | Thời gian tạo                               |
| `updated_at`      | `time.Time`  | Thời gian cập nhật                          |
