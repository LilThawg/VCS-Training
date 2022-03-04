2. Xây dựng 1 website quản lý user

    a. Viết lại các API (Challenge 1,2) theo framework (4d)
    
        i. Sử dụng Swagger để định nghĩa các API đã có

        ii. Sử dụng openapi middleware của web framework (go/gin) để validate các trường dữ liệu trong các API.
        (Ví dụ: user name, permission, role name: (type: string, required: true, regex: Chỉ chứ các ký tự [a-Z], -, _, [0-9])

    b. Tính năng quản lý (CRUD) với các bài post. (3d)

        i. User có thể tạo, sửa, xóa một bài post trên trang cá nhân.
        (Sử dụng token của user khác (không phải user tạo post) không có quyền được sửa, xóa bài post đó)

        ii. Admin có thể xem, xóa, sửa được tất cả các bài post của mọi users.

        iii. Phân trang khi số lượng bài post > 10

        iv. Các yêu cầu giống mục a