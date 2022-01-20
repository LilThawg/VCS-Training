2. Xây dựng 1 website quản lý user

Xây dựng 1 website quản lý user bao gồm các tính năng:
* Màn hình signup/login (1đ)
* Trang chủ (hiển thị danh sách user, và chỉ user đã đăng nhập mới có quyền xem) (1đ)
* API tạo role, tạo permission (5đ)
    ▪ Signup, login, logout (3đ)
    ▪ Tạo role: Admin, basic user (1đ)
    ▪ Tạo permission: Admin (view all users, delete user) (1đ)
* Sử dụng JWT
* Có khả năng mở rộng để thêm mới role, user

Tham khảo : https://www.bacancytechnology.com/blog/golang-jwt

# Mô tả quá trình 

## Màn hình trước khi đăng nhập 

![image](https://user-images.githubusercontent.com/72289126/150287934-5a1f71d7-2156-4905-acf0-3a53f64e7d35.png)

## Đăng ký 2 tài khoản admin và user 

![image](https://user-images.githubusercontent.com/72289126/150288544-dd0d4fcf-484c-43ca-aa80-5e2cdea5161f.png)

![image](https://user-images.githubusercontent.com/72289126/150288646-c4d84ed1-7fe9-4262-91bf-b85052c75139.png)

## Kiểm tra trong database

![image](https://user-images.githubusercontent.com/72289126/150288779-7746413f-aa32-492d-8796-33adee81c5d1.png)

sau đó sửa role của admin thành admin

![image](https://user-images.githubusercontent.com/72289126/150289081-cfe14dde-d834-4c07-8dd1-55b2e5ae2bfa.png)

## Bây giờ login với tài khoản admin trước

![image](https://user-images.githubusercontent.com/72289126/150291147-7d5b2b70-6a6a-4cb3-b606-aaaee6b4094e.png)

### sau khi đăng nhập 

![image](https://user-images.githubusercontent.com/72289126/150291245-7f0c2938-144f-4b23-adfa-005709049f05.png)

### click vào get all user 

![image](https://user-images.githubusercontent.com/72289126/150291367-357e6db9-6a66-4701-9d39-9fb5d00fda97.png)

### click vào Add new user và nhập thông tin user muốn thêm 

![image](https://user-images.githubusercontent.com/72289126/150291591-daacabf0-230d-443f-8e9e-b3c3371b5146.png)

#### Kiểm tra trong database

![image](https://user-images.githubusercontent.com/72289126/150291769-1fe86b8d-f991-4b33-a9ff-23000acde14f.png)

### click vào Delete user và nhập user muốn xoá theo email 

![image](https://user-images.githubusercontent.com/72289126/150292084-566eead6-e2e5-453c-9246-74996e2feb63.png)

#### Kiểm tra trong database

![image](https://user-images.githubusercontent.com/72289126/150292239-a3c37dd0-503b-4875-80b1-1e926acd3be7.png)

### Bấm log out để clear token và quay lại trang chủ

![image](https://user-images.githubusercontent.com/72289126/150292688-65a068fc-9862-43cb-8b65-6357ff0771ad.png)

## Login với tài khoản user

![image](https://user-images.githubusercontent.com/72289126/150292890-8c5ff49f-be6f-4599-b3fa-fcc7ae7cbd5b.png)

### User vẫn truy cập vào được các đường dẫn nhưng server xác thực rằng user thì không có 3 quyền kia, user chỉ có thể vào và log out thôi

![image](https://user-images.githubusercontent.com/72289126/150293895-abe3a50f-ce79-406f-8a30-970ad893ff48.png)

![image](https://user-images.githubusercontent.com/72289126/150294133-ab25c81d-32e7-48cd-8f3a-1efe4fb7f1c1.png)

