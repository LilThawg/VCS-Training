# Phân biệt SQL và NoSQL Database

## SQL là gì ?

SQL là thuật ngữ viết tắt của từ `Structured Query Language`. Là một ngữ truy vấn cấu trúc; nó dùng để xử lý cơ sở dữ liệu quan hệ. SQL có rất nhiều mệnh đề như toán tử, biểu thức, truy vấn và truy vấn con.

Hiện lập trình SQL có thể được sử dụng hiệu quả để chèn, tìm kiếm và cập nhập các sản phẩm hoặc xóa các bản ghi cơ sở dữ liệu. Có thể nói, SQL có thể làm rất nhiều thứ nhưng nó không giới hạn việc tối ưu hóa và duy trì cơ sở dữ liệu.

Các SQL phổ biến là: PostgreSQL, MySQL, Oracle và Microsoft SQL Server,...

## NoSQL là gì ?

NoSQL thuật ngữ viết tắt `None-Relational SQL` là ngược lại với SQL (có một số tài liệu ghi là `Cơ sở dữ liệu NoSQL là viết tắt của "Not only SQL" - "Không chỉ SQL" hoặc "Not SQL" - "Không phải SQL"; 2 cách hiểu đều đúng`), nó được sử dụng với mục đích tương đối giống như SQL, nhưng là đối với các cơ sở dữ liệu không quan hệ và không yêu cầu một lược đồ cố định và mở rộng.

Cơ sở dữ liệu (Database) NoSQL được sử dụng cho các kho dữ liệu phân tán với nhu cầu lưu trữ dữ liệu khổng lồ.

NoSQL được sử dụng lưu trữ dữ liệu lớn (Big Data) và các ứng dụng web với thời gian thực.

Để có được một hệ thống cơ sở dữ liệu NoSQL bao gồm hoạt động các cơ sở dữ liệu công nghệ lưu trữ những cơ sở dữ liệu có cấu trúc, bán cấu trúc, không có cấu trúc hoặc thậm chí là đa hình.

Các NoSQL phổ biến là: Redis, RavenDB Cassandra, MongoDB, BigTable, HBase, Neo4j và CouchDB,...

## Sự khác nhau giữa SQL và NoSQL

![image](https://starship-knowledge.com/wp-content/uploads/2020/12/nosql_vs_sql-1024x724.jpeg)

So sánh SQL và NoSQL
--------------------

| Tham số | SQL | NoSQL |
| --- | --- | --- |
| Định nghĩa | Cơ sở dữ liệu SQL chủ yếu được gọi là RDBMS hoặc Cơ sở dữ liệu quan hệ | Cơ sở dữ liệu NoSQL chủ yếu được gọi là cơ sở dữ liệu không liên quan hoặc phân tán |
| Design for | RDBMS truyền thống sử dụng cú pháp và truy vấn SQL để phân tích và lấy dữ liệu để có thêm thông tin chi tiết. Chúng được sử dụng cho các hệ thống OLAP. | Hệ thống cơ sở dữ liệu NoSQL bao gồm nhiều loại công nghệ cơ sở dữ liệu khác nhau. Các cơ sở dữ liệu này được phát triển để đáp ứng nhu cầu trình bày cho sự phát triển của ứng dụng hiện đại. |
| Ngôn ngữ Query | Structured query language (SQL) | Không có ngôn ngữ query |
| Type | SQL databases là cơ sở dữ liệu dựa trên bảng | NoSQL databases có thể dựa trên tài liệu, cặp khóa-giá trị, cơ sở dữ liệu biểu đồ |
| Schema | SQL databases có lược đồ được xác định trước | NoSQL databases sử dụng lược đồ động cho dữ liệu phi cấu trúc. |
| Khả năng mở rộng | SQL databases có thể mở rộng theo chiều dọc | NoSQL databases có thể mở rộng theo chiều ngang |
| Ví dụ | Oracle, Postgres, and MS-SQL. | MongoDB, Redis, , Neo4j, Cassandra, Hbase. |
| Phù hợp cho | Đây là 1 lựa chọn lý tưởng cho môi trường truy vấn phức tạp | Không phù hợp với truy vấn phức tạp |
| Lưu trữ dữ liệu phân cấp | SQL databases không thích hợp cho việc lưu trữ dữ liệu phân cấp. | Phù hợp hơn cho kho lưu trữ dữ liệu phân cấp vì nó hỗ trợ phương thức cặp khóa-giá trị. |
| Variations | Một loại có biến thể nhỏ | Nhiều loại khác nhau bao gồm các kho khóa-giá trị, cơ sở dữ liệu tài liệu và cơ sở dữ liệu đồ thị. |
| Năm phát triển | Nó được phát triển vào những năm 1970 để giải quyết các vấn đề với lưu trữ tệp phẳng | Được phát triển vào cuối những năm 2000 để khắc phục các vấn đề và hạn chế của SQL databases. |
| Open-source | Một sự kết hợp của mã nguồn mở như Postgres & MySQL, và thương mại như Oracle Database. | Open-source |
| Tính nhất quán | Nó phải được cấu hình cho sự nhất quán chặt chẽ. | Nó phụ thuộc vào DBMS như một số cung cấp tính nhất quán mạnh mẽ như MongoDB, trong khi những người khác cung cấp chỉ cung cấp sự nhất quán cuối cùng, như Cassandra. |
| Được sử dụng tốt nhất cho | RDBMS database là tùy chọn thích hợp để giải quyết các vấn đề về ACID. | NoSQL được sử dụng tốt nhất để giải quyết các vấn đề về tính khả dụng của dữ liệu |
| Tầm quan trọng | Nó nên được sử dụng khi hiệu lực dữ liệu là siêu quan trọng | Sử dụng khi nó quan trọng hơn để có dữ liệu nhanh hơn dữ liệu chính xác |
| Lựa chọn tốt nhất | Khi bạn cần hỗ trợ truy vấn động | Sử dụng khi bạn cần mở rộng quy mô dựa trên yêu cầu thay đổi |
| Hardware | Specialized DB hardware (Oracle Exadata, etc.) | Commodity hardware |
| Network | Highly available network (Infiniband, Fabric Path, etc.) | Commodity network (Ethernet, etc.) |
| Loại lưu trữ | Highly Available Storage (SAN, RAID, etc.) | Commodity drives storage (standard HDDs, JBOD) |
| Tính năng tốt nhất | Hỗ trợ đa nền tảng, Bảo mật và miễn phí | Dễ sử dụng, hiệu suất cao và công cụ linh hoạt. |
| Mô hình ACID và BASE | ACID (Atomicity, nhất quán, cách ly và độ bền) là một chuẩn cho RDBMS | Cơ bản (Về cơ bản có sẵn, trạng thái mềm, phù hợp cuối cùng) là một mô hình của nhiều hệ thống NoSQL |
| Performance | SQL hoạt động tốt và nhanh thì việc desgin tốt là cực kì quan trọng và ngược lại. | Nhanh hơn SQL NoSQL thì denormalized cho phép bạn lấy được tất cả thông tin về một item cụ thể với các codition mà không cần JOIN liên quan hoặc truy vấn SQL phức tạp. |
| Kết luận | Dự án đã có yêu cầu dữ liệu rõ ràng xác định quan hệ logic có thể được xác định trước. | Phù hợp với những dự án yêu cầu dữ liệu không liên quan, khó xác định, đơn giản mềm dẻo khi đang phát triển |

## Một số câu lệnh SQL cơ bản

**1) Trên cơ sở dữ liệu (database)**

create database: tạo cơ sở dữ liệu

```sql
create database HoSoSinhVien
```

drop database: xóa cơ sở dữ liệu 

```sql
drop database HoSoSinhVien
```

alter database: sửa các thông tin của cơ sở dữ liệu
```sql
alter database HoSoSinhVien modify name = HSSV
```

**2) Trên bảng (table)**

Nhớ các kiểu số liệu: text, binary, numberic, money, datetime, bit, variant

Nhớ các các ràng buộc: default, check, unique, foreign, primary

Nhớ các kiểu toàn vẹn: entity, domain, referential, user

Nhớ các thuộc tính bổ trợ: identity, null

create table: tạo bảng

```sql
create table SinhVien (MaSV int,TenSV nchar(50))
```

drop table: xóa bảng

```sql
drop table SinhVien
```

alter table...add: thêm cột

```sql
alter table SinhVien add QueQuan int
```

alter table...drop column: xóa cột

```sql
alter table SinhVien drop column QueQuan
```

alter table...add primary key: thêm khóa chính

```sql
alter table SinhVien alter column MaSV int not null
go
alter table SinhVien
add primary key (MaSV)
```

**3) Trên bản ghi (record)**

insert..values: thêm các bản ghi vào bảng

```sql
insert into SinhVien (MaSV, TenSV, QueQuan) values (1, N'Nguyễn Văn A', 1)
```

delete: xóa các bản ghi từ bảng

```sql
delete from SinhVien where MaSV=1
```

update: sửa các bản ghi trong bảng

```sql
UPDATE SinhVien
SET TenSV = 'thang' 
WHERE MaSV=1;
```

**4) Truy vấn (query)**

Hỗ trợ truy vấn: distinct, top, as, identity

Phép toán tập hợp: in, like, between

Các hàm tổng nhóm: sum, max, min, avg

4.1) Truy vấn đơn giản

select *: Hiện tất cả bảng

```sql
select * from SinhVien
```

select: Hiện một số cột

```sql
select MaSV,TenSV  from SinhVien
```

select..where: Hiện một số dòng / bản ghi

```sql
select TenSV from SinhVien where MaSv = 1
```

select..order by: Hiện và sắp xếp theo điểm rồi theo tên

```sql
select TenSV, DiemTB from SinhVien
order by DiemTB desc, TenSV asc // asc sắp sếp tăng dần, desc là giảm dần
```

select..distinct: Hiện danh sách giá trị không trùng lặp

```sql
select distinct QueQuan from SinhVien
```

select..top: Hiện các dòng đầu tiên trong bảng

```sql
select top 3 TenSV, DiemTB from SinhVien
```

4.2) Truy vấn lồng nhau (nested query)

select..where (select)

Hiện tất cả những người trong bảng nhân viên có lương bằng lương lớn nhất của những người có trong công ty:

```sql
select TenNV, Luong from NhanVien
where Luong = (select max(Luong) from NhanVien)
```

select..where (in)

Hiện tất cả những người trong bảng nhân viên có lương lớn nhất hoặc lớn nhì của những người có trong công ty:

```sql
select TenNV, Luong from NhanVien
where Luong in (select top 2 Luong from NhanVien order by Luong)
```

4.3) Truy vấn tổng nhóm (subtotal query / grouping query)

select..group by: Thống kê theo tiêu chí

Hiện ra số lượng các nhân viên ứng với từng quê

```sql
select QueQuan, count(*) from NhanVien
group by QueQuan
```

select..having: Hiện ra một số nhóm phù hợp

Chỉ đếm số lượng người ở Hải Phòng và số lượng người ở Hà nội

```sql
select QueQuan, count(*) from NhanVien
group by QueQuan
having (QueQuan = ‘HP’, QueQuan = ‘HN’)
```

4.4) Truy vấn liên bảng (cross table query / joining query)

select..inner join: ghép các cặp bản ghi thỏa mãn điều kiện

Ghép bảng nhân viên và hiện ra tên nhân viên và tên địa phương

```sql
select NhanVien.TenNV, DiaPhuong.TenDP
from NhanVien 
inner join DiaPhuong on NhanVien.QueQuan = DiaPhuong.MaDP
``` 

select..left outer join: lấy tất cả phía trái và ghép (nếu có) với phía phải

Lấy tất cả những nhân viên kể cả những nhân viên có quê quán không hợp lệ (nghĩa là mã quê quán không có trong bảng địa phương)

```sql
select NhanVien.TenNV, DiaPhuong.TenDP
from NhanVien 
left outer join DiaPhuong on NhanVien.QueQuan = DiaPhuong.MaDP
```

select..right outer join: lấy tất cả phía phải và ghép (nếu có) với phía trái

Lấy tất cả những địa phương ghép với nhân viên, các địa phương không hợp lệ sẽ được ghép với bộ dữ liệu rỗng. Không hiện ra các nhân viên không có mã quê quán phù hợp

```sql
select NhanVien.TenNV, DiaPhuong.TenDP
from NhanVien 
right outer join DiaPhuong on NhanVien.QueQuan = DiaPhuong.MaDP
```

select..full outer join: lấy từ hai phía và ghép nếu có

Lấy tất cả những nhân viên (nếu không có quê quán phù hợp thì ghép với bộ dữ liệu rỗng) và tất cả những địa phương kể cả không có nhân viên.

```sql
select NhanVien.TenNV, DiaPhuong.TenDP
from NhanVien 
right outer join DiaPhuong on NhanVien.QueQuan = DiaPhuong.MaDP
```

select..cross join: trả về tất cả các cặp có thể ghép

Ghép từng nhân viên với tất cả các địa phương. Như vậy nếu có m nhân viên và có n địa phương thì bảng đích sẽ có m*n dòng. n dòng đầu cho nhân viên thứ nhất ghép với các địa phương. n dòng sau cho nhân viên thứ hai ghép với các địa phương. và tiếp tục như thế tới nhân viên thứ m.

```sql
select NhanVien.TenNV, DiaPhuong.TenDP
from NhanVien 
cross join DiaPhuong
```






