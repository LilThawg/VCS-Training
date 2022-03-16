# Chuẩn hoá dữ liệu là gì ?

Chuẩn hóa là một kỹ thuật thiết kế cơ sở dữ liệu tổ chức các bảng theo cách làm giảm sự dư thừa và phụ thuộc của dữ liệu hay có thể hiểu đơn giản là thiết kế database theo 1 quy chuẩn nhất định nào đó.

Chuẩn hóa cơ sở dữ liệu nhằm 3 mục đích:

1. Giảm thiểu dư thừa dữ liệu
2. Đảm bảo tính nhất quán của dữ liệu
3. Hỗ trợ dễ dàng cập nhật, truy xuất dữ liệu

# Các kiểu chuẩn hóa cơ sở dữ liệu

Có 4 dạng chuẩn hóa cơ bản đó là:

- First Normal Form (1NF): dạng chuẩn 1NF
- Second Normal Form (2NF): dạng chuẩn 2NF
- Third Nomal Form (3NF): dạng chuẩn 3NF
- Boyce-Codd Normal Form (BCNF): dạng chuẩn Boyce-Codd

Các dạng chuẩn được sắp xếp theo thứ tự từ thấp đến cao. Để chuẩn hóa 2NF thì cơ sở dữ liệu của bạn phải đạt chuẩn 1NF, tương tự nếu đạt chuẩn 3NF thì phải đạt chuẩn 1NF và 2NF. Và chuẩn Boyce-Codd sẽ bao gồm 3 loại chuẩn 1NF, 2NF và 3NF.

## Dạng chuẩn 1NF

Quy tắc 1NF :

- Các trường thuộc tính phải là nguyên tố, không được chứa giá trị phức.
- Không chứa các thuộc tính gây lặp.
- Không chứa các thuộc tính có thể tính toán từ các thuộc tính khác.

Ví dụ bảng chưa chuẩn hoá 1NF :

![image](https://images.viblo.asia/db98ce98-08e2-4d39-bb29-59bab55acb77.png)

Bảng có 3 khóa chính là `customer_id`, `order_id` và `product_id`.

Bảng dữ liệu này vi phạm tất cả quy tắc của chuẩn 1NF vì:

- `address`, `phone`, `customer_name` chứa các giá trị trùng lặp.
- giá trị `address` trong từng hàng không phải là đơn trị (chỉ có 1 giá trị).
- `total_amount` hoàn toàn có thể tính toán được bằng cách `quantity` \* `unit_price`, không nhất thiết phải đưa vào bảng, gây ra dư thừa dữ liệu.

Chuẩn hoá lại bảng như sau :

- Tách các thuộc tính lặp trong bảng như: `customer_name`, `phone` ra thành một bảng mới là `customers` có khóa là `customer_id`.
- Tách `address` thành một bảng riêng có khóa là `customer_id` để biết địa chỉ đó thuộc về `customer` nào.
- Loại bỏ thuộc tính `total_amount`

![image](https://images.viblo.asia/2134df5f-2b7f-4639-9976-66b6879482d7.png)

## Dạng chuẩn 2NF

Quy tắc 2NF :

- Phải đạt dạng chuẩn 1
- Các trường thuộc tính không phải là khoá chính, phải phụ thuộc hoàn toàn vào khoá chính. Không được phép phụ thuộc vào 1 phần của khoá chính.

Quy tắc chuẩn hóa từ chuẩn 1NF thành 2NF:

- Bước 1: Loại bỏ các thuộc tính không khóa phụ thuộc vào một bộ phận khóa chính và tách ra thành một bảng riêng, khóa chính của bảng là bộ phận của khóa mà chúng phụ thuộc vào.
- Bước 2: Các thuộc tính còn lại lập thành một quan hệ, khóa chính của nó là khóa chính ban đầu.

Bảng dữ liệu bên trên vẫn chưa đạt chuẩn 2NF vì :

- một số thuộc tính như description , unit_price phụ thuộc vào 1 phần của khóa là product_id chứ không cần phụ thuộc cả vào tập khóa (customer_id, order_id, product_id), hay thuộc tính customer_name và phone cũng chỉ phụ thuộc vào customer_id, thuộc tính order_date phụ thuộc vào customer_id và order_id, thuộc tính quantity phụ thuộc vào order_id và product_id.

Vậy nên để đạt chuẩn 2NF thì ta sẽ thiết kế tiếp bảng dữ liệu chuẩn 1NF như sau:

- Tách các thuộc tính (product_id, description, unit_price) thành một bảng riêng là products.
- Các thuộc tính (customer_id, order_id, order_date) làm thành một bảng, mình đặt tên là orders.
- Còn lại các thuộc tính (order_id, product_id, quantity) làm thành một bảng trung gian giữa products và orders, mình đặt là order_products.

Chỉ cần tuân thủ 2 chuẩn mà ta đã được cơ sở dữ liệu chuẩn hóa như sau:

![image](https://images.viblo.asia/3a692db9-598a-41e9-928d-841404a6c83f.png)

## Dạng chuẩn 3NF

Quy tắc 3NF :

- Phải đạt chuẩn 2
- Các trường thuộc tính không phải khoá chính, phải phụ thuộc trực tiếp vào khoá chính. Không được phép phụ thuộc bắc cầu thông qua thuộc tính khác. (nghĩa là tất cả các thuộc tính không khóa phải được suy ra trực tiếp từ thuộc tính khóa)

Quy tắc chuẩn hóa từ 2NF thành 3NF:

- Bước 1: Loại bỏ các thuộc tính phụ thuộc bắc cầu ra khỏi quan hệ và tách chúng thành quan hệ riêng có khóa chính là thuộc tính bắc cầu.
- Bước 2: Các thuộc tính còn lại lập thành một quan hệ có khóa chính là khóa ban đầu.

Bảng trên đạt chuẩn 2NF và cũng đạt chuẩn luôn 3NF rồi nên ta lấy ví dụ khác.

Bảng sau vi phạm 3NF:

![image](https://user-images.githubusercontent.com/72289126/157527404-e5c2c1be-554a-417a-9503-1c03b7bdc7b8.png)

Ta thấy rằng `Tên lớp` và `Sĩ Số` không phụ thuộc trực tiếp vào khoá chính `Mã SV` mà phụ thuộc vào `Mã lớp`. Do vậy 2 thuộc tính đó phụ thuộc bắc cầu thông qua `Mã lớp`.

Bảng đạt chuẩn 3NF :

![image](https://user-images.githubusercontent.com/72289126/157528911-f2b7ba02-d0fe-408d-9a98-d29c9cf0f9eb.png)

Còn 1 bảng mới với khoá chính là `Mã lớp` với 2 trường `Tên lớp` và `Sĩ số` liên kết với bảng trên thì mọi người tự vẽ nhé :D

## Dạng chuẩn Boyce-Codd

Quy tắc BCNF :

- Phải đạt chuẩn 3NF
- Không có thuộc tính khóa nào phụ thuộc vào thuộc tính không khóa

Quy tắc chuẩn hóa 3NF thành Boyce-Codd:

- Bước 1: Loại bỏ các thuộc tính khóa phụ thuộc hàm vào thuộc tính không khóa ra khỏi quan hệ
- Bước 2: Tách thuộc tính vừa loại bỏ thành một quan hệ riêng có khoá chính là thuộc tính không khóa gây ra phụ thuộc.

Ví dụ :

![image](https://user-images.githubusercontent.com/72289126/157530240-42d10fce-cb76-4a6b-aa17-42d983393000.png)

# Summary

![image](https://images.viblo.asia/f67dc66b-2714-46bc-88f8-ed3ef41a4864.png)

