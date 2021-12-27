2. Xây dựng 1 service API sử dụng Golang và Restful API
- Dựng 1 service API sử dụng Golang và viết API tương tác với database postgres
theo chuẩn Restful bao gồm các API (4đ)
    - Lấy danh sách user, có thể tìm kiếm theo tên
    - Thêm mới một user
    - Sửa thông tin một user
    - Xoá user khỏi database
- Thông tin user bao gồm: tên, ngày/tháng/năm sinh, giới tính, email

**Chuẩn bị**

```
Postgresql - storing data
GORM - communicate with database
Gorrila/mux - serve api
Postman - test api
```


1. Tạo go.mod để định nghĩa đường dẫn của các module

```cmd
go mod init github.com/VCS-trainning/Ex1/2/crud
```

go.sum: Để lưu vết version của thư viện mà project đang sử dụng.

2. Tải các thư viện

Gorrila/mux - Server API

```cmd
go get -u github.com/gorilla/mux
```

GORM - communicate with database postgres

```cmd
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

**Kết quả**

Tạo user với method post

![image](https://user-images.githubusercontent.com/72289126/147465855-62e7d931-8338-44e1-9d2c-6a9d4b5e5d59.png)

Kiểm tra trong database

![image](https://user-images.githubusercontent.com/72289126/147466046-ab6786f4-1720-46b4-82f0-b508d6d0335f.png)

Post thêm 2 user nữa, sau đó sử dụng method get để lấy toàn bộ thông tin user 

![image](https://user-images.githubusercontent.com/72289126/147467539-c254830e-741e-4f3b-8149-9618a66867d0.png)

Tìm kiếm user theo tên Alice

![image](https://user-images.githubusercontent.com/72289126/147467662-5b114fd7-f316-4b78-bc34-4c5563b11b13.png)

Chỉnh sửa thông tin user có id = 2 và kiểm tra lại

![image](https://user-images.githubusercontent.com/72289126/147468014-ee529e49-3ef1-4ebe-bf47-f91ddd10a0ff.png)

![image](https://user-images.githubusercontent.com/72289126/147468067-4ef7c301-0b96-47e1-be87-45044f599b2f.png)

Delete user có id = 2 và kiểm tra lại

![image](https://user-images.githubusercontent.com/72289126/147468209-a1dc8be2-1a3c-4ba7-ac54-492ddcba7e26.png)

![image](https://user-images.githubusercontent.com/72289126/147468272-e6a46f8b-02a0-4d2c-b8a9-d24ea84ee8d7.png)
