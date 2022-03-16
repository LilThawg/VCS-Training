# Mô hình ERD

## 1 - Mô hình erd là gì?

Mô hình erd được viết tắt bởi từ Entity Relationship Diagram được hiểu là mô hình thực thể kết hợp hay còn gọi là thực thể liên kết. Mô hình này còn được biết tới với các gọi khác là er (viết tắt của từ Entity Relationship model). Vậy mô hình er là gì? Mô hình erd hay er bao gồm các thực thể, những mối kết hợp và đặc biệt là danh sách thuộc tính.

## 2 - Vai trò của mô hình E-R trong quá trình thiết kế cơ sở dữ liệu:

Mục tiêu của mô hình E-R trong quá trình thiết kế cơ sở dữ liệu đó là phân tích dữ liệu, xác định các đơn vị thông tin cơ bản cần thiết của tổ chức, mô tả cấu trúc và mối liên hệ giữa chúng

![image](https://images.viblo.asia/6d0787fc-73e6-4779-b4dd-14db89e44eee.png)

E-R là mô hình trung gian để chuyển những yêu cầu quản lý dữ liệu trong thế giới thực thành mô hình cơ sở dữ liệu quan hệ

## 3 - Cách vẽ sơ đồ erd (theo P.Chen)

![image](https://user-images.githubusercontent.com/72289126/157619213-90303b39-94cc-4f62-a9c5-7e109d041fc5.png)

![image](https://user-images.githubusercontent.com/72289126/157619530-0e6f2088-4de0-4f56-bae5-b56a47da681c.png)

Sau khi nắm được mô hình erd là gì, chúng ta sẽ tìm hiểu cách vẽ sơ đồ erd. Để vẽ được sơ đồ erd, ta cần lưu ý một số ký hiệu sau:

- Hình chữ nhật: biểu diễn thực thể
- Hình elip: biểu diễn thuộc tính, trong hình elip có ghi tên thuộc tính
- Hình thoi: biểu diễn quan hệ

Các bước vẽ sơ đồ erd:

- Thông qua việc liệt kê và lựa chọn thông tin dựa trên giấy tờ, hồ sơ
- Xác định mối quan hệ giữa thực thể và thuộc tính của nó
- Xác định mối quan hệ có thể có giữa các thực thể và mối kết hợp
- Vẽ mô hình erd bằng các ký hiệu sau đó chuẩn hóa và thu gọn sơ đồ

![image](https://i1.wp.com/dinhnghia.vn/wp-content/uploads/2018/08/mo-hinh-erd-la-gi-2.png?resize=600%2C334&ssl=1)

## 4 - Các thành phần cơ bản của mô hình E-R

Các thành phần cơ bản của mô hình E-R là :

- Thực thể và tập thực thể
- Thuộc tính
- Mối quan hệ giữa các tập thực thể

### a) Thực thể và tập thực thể

Thực thể là một đối tượng trong thế giới thực.

Một nhóm bao gồm các thực thể tương tự nhau tạo thành một tập thực thể

Việc lựa chọn các tập thực thể là một bước vô cùng quan trọng trong việc xây dựng sơ đồ về mối quan hệ thực thể

Thực thể là 1 thứ có thể xác định được (như một người, khái niệm, đối tượng hoặc sự kiện,...) có thể được lưu trữ các dữ liệu về chính nó. Các thực thể hãy được coi là các danh từ như: Sinh viên, khách hàng, sản phẩm, ô tô,... Thực thể có ký hiệu là hình chữ nhật.

![image](https://d3hi6wehcrq5by.cloudfront.net/itnavi-blog/Entity%20Relationship%20Diagram%20l%C3%A0%20g%C3%AC%206.jpg)

Là một thứ có thể xác định được như một người, đối tượng, khái niệm hoặc sự kiện,… Thực thể có thể là các danh từ như: sinh viên, khách hàng, ô tô hoặc sản phẩm.

### b) Thuộc tính

Thực thế sẽ có các thuộc tính liên quan tới nó. Ví dụ con người thì sẽ có họ tên, giới tính, ngày sinh, số CMND, địa chỉ,... Sinh viên thì có mã số sinh viên, trường, lớp,... Thuộc tính trong mô hình ER được ký hiệu là hình bầu dục hoặc hình tròn, tên thuộc tính bên trong. Thuộc tính thuộc thực thể nào thì được nối đến thực thể đó.

Thuộc tính bao gồm các loại như sau:

- Thuộc tính đơn – không thể tách nhỏ ra được
- Thuộc tính phức hợp – có thể tách ra thành các thành phần nhỏ hơn

Các loại giá trị của thuộc tính:

- Đơn trị: các thuộc tính có giá trị duy nhất cho một thực thể (VD: số CMND, …)
- Đa trị: các thuộc tính có một tập giá trị cho cùng một thực thể (VD: bằng cấp, …)
- Suy diễn được (năm sinh <----> tuổi)

Mỗi thực thể đều được phân biệt bởi thuộc tính khóa

Ví dụ :

![image](https://images.viblo.asia/37256d0c-6712-49b8-b73a-96fd5379c680.png)

![image](https://images.viblo.asia/9c49dce9-f46d-4293-9b5f-11b3e78a83e7.png)

### c) Mối quan hệ giữa các tập thực thể

Thành phần mối quan hệ thể hiện quan hệ giữa các thực thể trong mô hình. Nó cũng là thành phần đóng vai trò vô cùng quan trọng. Quan hệ này được ký hiệu bằng hình thoi, bên trong là tên mối quan hệ, nối đến các thực thể có quan hệ với nhau.

Mối quan hệ giữa các thực thể sẽ có các kiểu sau:

- quan hệ 1 - 1
- quan hệ 1 - n (hoặc n - 1)
- quan hệ n - n

Mối quan hệ giữa thực thể còn được đánh bảng số thể hiện số chiều của mối quan hệ.

![image](https://images.viblo.asia/a6800064-5364-4a44-9727-ca635c36fc24.png)

Ví dụ :

- Một phòng ban có nhiều nhân viên

![image](https://images.viblo.asia/f6cf6681-b353-479c-9b44-eb056663cae5.png)

- Một nhân viên chỉ thuộc 1 phòng ban

![image](https://images.viblo.asia/cca3520d-fb00-46d5-a263-970c488fb572.png)

- Một nhân viên có thể được phân công vào nhiều dự án hoặc không được phân công vào dự án nào

![image](https://images.viblo.asia/54f73971-9d60-470f-a3ec-724a7b4096c1.png)

- Một nhân viên có thể là trưởng phòng của 1 phòng ban nào đó

![image](https://images.viblo.asia/7e7e4af7-8e16-4465-ae72-684fb47e8701.png)

- Một loại thực thể có thể tham gia nhiều lần vào một quan hệ với nhiều vai trò khác nhau

![image](https://images.viblo.asia/0fa04207-52cf-4371-b390-a7e1f12522b4.png)
