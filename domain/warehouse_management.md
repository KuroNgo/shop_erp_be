# Quản Lý Kho

## Tổng Quan
Hệ thống quản lý kho bao gồm các mô hình để quản lý thông tin về kho lưu trữ, sản phẩm, tồn kho, di chuyển kho, nhà cung cấp, đơn mua hàng, chi tiết đơn mua hàng và điều chỉnh tồn kho.

## Các Bảng và Mô Hình Dữ Liệu

### 1. Warehouses (Kho)
| Trường           | Loại Dữ Liệu | Mô Tả              |
|------------------|--------------|--------------------|
| `id`             | `ObjectID`   | Khóa chính         |
| `warehouse_name` | `string`     | Tên kho            |
| `location`       | `string`     | Vị trí của kho     |
| `capacity`       | `int`        | Công suất của kho  |
| `created_at`     | `time.Time`  | Thời gian tạo      |
| `updated_at`     | `time.Time`  | Thời gian cập nhật |

### 2. Products (Sản phẩm)
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
| Trường         | Loại Dữ Liệu | Mô Tả                                       |
|----------------|--------------|---------------------------------------------|
| `id`           | `ObjectID`   | Khóa chính                                  |
| `product_id`   | `ObjectID`   | Tham chiếu tới bảng Products (khóa ngoại)   |
| `warehouse_id` | `ObjectID`   | Tham chiếu tới bảng Warehouses (khóa ngoại) |
| `quantity`     | `int`        | Số lượng tồn kho                            |
| `created_at`   | `time.Time`  | Thời gian tạo                               |
| `updated_at`   | `time.Time`  | Thời gian cập nhật                          |

### 4. StockMovements (Di chuyển kho)
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
