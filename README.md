# SHOP ECOMMERCE - VN
## Tìm hiểu về mô hình ERP
### ERP là gì ?
Mô hình ERP (Enterprise Resource Planning - Hoạch định nguồn lực doanh nghiệp) là một 
hệ thống phần mềm quản lý tích hợp các quy trình kinh doanh cốt lõi của một tổ chức. 
ERP kết nối và quản lý các phòng ban và bộ phận khác nhau trong một doanh nghiệp, 
giúp tối ưu hóa hoạt động, cải thiện hiệu quả và cung cấp dữ liệu thời gian thực để hỗ trợ ra quyết định.

### Các thành phần chính của ERP
1. **Quản lý tài chính (Financial Management)**: Quản lý kế toán, ngân sách, thuế và báo cáo tài chính.
2. **Quản lý nhân sự (Human Resource Management - HRM)**: Quản lý tiền lương, chấm công, tuyển dụng và đào tạo.
3. **Quản lý sản xuất (Manufacturing)**: Hoạch định sản xuất, quản lý chuỗi cung ứng và kiểm soát tồn kho.
4. **Quản lý bán hàng và phân phối (Sales and Distribution)**: Theo dõi đơn hàng, quản lý kho và giao hàng.
5. **Quản lý quan hệ khách hàng (Customer Relationship Management - CRM)**: Quản lý thông tin khách hàng, chăm sóc và duy trì mối quan hệ với khách hàng.

### Lợi ích của hệ thống ERP
* **Tăng cường hiệu suất**: Tích hợp dữ liệu từ các bộ phận giúp giảm thiểu lỗi, tăng hiệu quả công việc.
* **Quyết định chính xác**: Dữ liệu thời gian thực giúp lãnh đạo đưa ra các quyết định chiến lược đúng đắn.
* **Tự động hóa**: Tự động hóa các quy trình kinh doanh giúp giảm bớt công việc thủ công và nâng cao năng suất.****

## Về dự án


## Run Programming
How to run this project?
We can run this Go Backend Clean Architecture project with or without Docker. Here, I am providing both ways to run this project.

#### Clone this project
    cd your-workspace
- Move to your workspace


#### Clone this project into your workspace
- git clone https://github.com/KuroNgo/FEIT.git

#### Move to the project root directory
    cd FEIT

#### Run without Docker
- Create a file .env similar to .env.example at the root directory with your configuration.
- Install go if not installed on your machine.
- Install MongoDB if not installed on your machine.
- Important: Change the DB_HOST to localhost (DB_HOST=localhost) in .env configuration file. DB_HOST=mongodb is needed only when you run with Docker.
- Run go run cmd/main.go.
- Access API using http://localhost:8080
#### Run with Docker
- Create a file .env similar to .env.example at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run docker-compose up -d.
- Access API using http://localhost:8080
#### How to run the test?
#### Run all tests
    go test ./...
#### How to generate the mock code?
- In this project, to test, we need to generate mock code for the use-case, repository, and database.

#### Generate mock code for the usecase and repository
    mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

#### Generate mock code for the database
    mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
- Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the corresponding command to regenerate the mock code for testing.


## FIX User API