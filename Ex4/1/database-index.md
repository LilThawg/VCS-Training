# Index trong database

## Khái niệm Index trong database

Index là một cấu trúc dữ liệu được dùng để định vị và truy cập nhanh nhất vào dữ liệu trong các bảng database

Index là một cách tối ưu hiệu suất truy vấn database bằng việc giảm lượng truy cập vào bộ nhớ khi thực hiện truy vấn

## Index database dùng để làm gì?

Giả sử ta có một bảng User lưu thông tin của người dùng, ta muốn lấy ra thông tin của người dùng có trường tên (Name) là `LeThang` . Ta có truy vấn SQL sau: `SELECT * FROM User WHERE Name = 'LeThang';`

Nếu không có Index cho cột Name, truy vấn sẽ phải chạy qua tất cả các Row của bảng User để so sánh và lấy ra những Row thỏa mãn. Vì vậy, khi số lượng bản ghi lớn, chuyện gì sẽ xảy ra ?? Index được sinh ra để giải quyết vấn đề này.

Index trỏ tới địa chỉ dữ liệu trong một bảng, nó giống như mục lục của quyển sách bạn đọc, nó giúp truy vấn trở nên nhanh chóng như việc bạn xem mục lục.

Index có thể được tạo cho một hoặc nhiều cột trong database. Index thường được tạo mặc định cho primary key, foreign key. Ngoài ra, ta cũng có thể tạo thêm index cho các cột nếu cần.

## Cấu trúc của Index

Index gồm:

1. Cột Search Key: chứa bản sao các giá trị của cột được tạo Index
2. Cột Data Reference: chứa con trỏ trỏ đến địa chỉ của bản ghi có giá trị cột index tương ứng

![image](https://images.viblo.asia/c44b421d-0e4d-4bf9-966e-ddc03bc75993.png)

## Ưu và nhược điểm của Index

Ưu điểm :

- Ưu điểm của index là tăng tốc độ tìm kiếm records theo câu lệnh WHERE.
- Không chỉ giới hạn trong câu lệnh SELECT mà với cả xử lý UPDATE hay DELETE có điều kiện WHERE.

Nhược điểm :

- Khi sử dụng index thì tốc độ của những xử lý ghi dữ liệu (Insert, Update, Delete) sẽ bị chậm đi.
- Vì ngoài việc thêm hay update thông tin data thì CSDL cũng cần update lại thông tin index của bảng tương ứng.
- Tốc độ xử lý bị chậm đi cũng tỷ lệ thuận với số lượng index được sử dụng trong bảng.
- Do vậy với những table hay có xử lý insert, update hoặc delete và cần tốc độ xử lý nhanh thì không nên được đánh index.
- Ngoài ra việc đánh index cũng sẽ tốn resource của server như thêm dung lượng cho CSDL.

## Các kiểu Index

### 1. Clustered Index

Clustered index là loại index theo đó các bản ghi trong bảng được sắp thứ tự theo trường index. Khi bảng được tạo clustered index thì bản thân nó trở thành một cây index, với các node lá chứa khóa là các trường được index và cũng đồng thời chứa tất cả các trường còn lại của bảng.

Các đặc điểm sau đây cần lưu ý khi sử dụng Clustered Index:

- Mỗi bảng chỉ được phép có một Clustered Index.
- Khóa chính (Primary Key - PK) chính là một Clustered Index.
- Clustered Index chỉ được tạo trên cột hoặc tập cột có chứa những giá trị hoặc tập giá trị duy nhất.

Ta đi vào ví dụ sau để hiểu rõ :

Tạo bảng có 1 triệu bản ghi :

```sql
Create Table Employees
(
	Id int primary key identity,
	[Name] nvarchar(50),
	Email nvarchar(50),
	Department nvarchar(50)
)
Go

SET NOCOUNT ON
Declare @counter int = 1

While(@counter <= 1000000)
Begin
	Declare @Name nvarchar(50) = 'ABC ' + RTRIM(@counter)
	Declare @Email nvarchar(50) = 'abc' + RTRIM(@counter) + '@pragimtech.com'
	Declare @Dept nvarchar(10) = 'Dept ' + RTRIM(@counter)

	Insert into Employees values (@Name, @Email, @Dept)

	Set @counter = @counter +1

	If(@Counter%100000 = 0)
		Print RTRIM(@Counter) + ' rows inserted'
End
```

#### Clustered Index Seek

![image](https://user-images.githubusercontent.com/72289126/157830692-72f1db56-b359-469e-8cac-8d5fc427d9da.png)

Ở SQL Server Management Studio click on `Include Actual Execution Plan` icon và execute query :

```sql
Select * from Employees where Id = 932000
```

hover vào ta có các thông tin :

![image](https://user-images.githubusercontent.com/72289126/157830988-4de34b05-c9bf-4281-8b03-f887a3280fc5.png)

Chú ý `Clustered Index Seek` có nghĩa là database engine đang sử dụng `clustered index` trên cột `Id` để tìm thông tin employee với Id = 932000 (bởi vì khi khai báo `Id` là `primary key` thì bảng sẽ tự động gán `clustered index` cho cột `Id`).

![image](https://www.pragimtech.com/blog/contribute/article_images/1220210330203211/clustered-index-seek.png)

- Number of rows read = 1
- Actual number of rows for all executions = 1

Number of rows read là số bản ghi SQL cần đọc để đưa ra kết quả và Actual number of rows for all executions là số bản ghi thực tế sau khi truy vấn.

Chúng ta nhận được kết quả ngay lập tức. Đó là sức mạnh của Index.

```sql
DECLARE @EndTime datetime
DECLARE @StartTime datetime
SELECT @StartTime=GETDATE()

-- Write Your Query
Select * from Employees where Id = 992000

SELECT @EndTime=GETDATE()

--This will return execution time of your query
SELECT DATEDIFF(ms,@StartTime,@EndTime) AS [Duration in millisecs]
```

![image](https://user-images.githubusercontent.com/72289126/157839984-5454487f-b872-4d3d-b359-86d7255ad4cf.png)

Thử tạo 1 `clustered index` :

```sql
Create clustered index index_name
on Employees(Name)
```

Ta nhận được kết quả :

```
Msg 1902, Level 16, State 3, Line 16
Cannot create more than one clustered index on table 'Employees'. Drop the existing clustered index 'PK__Employee__3214EC07E4DAC301' before creating another.
```

Bởi vì 1 bảng chỉ có 1 `clustered index` thôi, muốn tạo ta phải Drop the existing clustered index 'PK**Employee**3214EC07E4DAC301' đã tự khởi tạo khi đặt `primary key` đó.

#### Clustered Index Scan

Bây giờ chúng ta hãy thử tìm theo tên :

```sql
Select * from Employees Where Name = 'ABC 932000'
```

![image](https://www.pragimtech.com/blog/contribute/article_images/1220210330203211/clustered-index-scan.png)

Chú ý đây là `Clustered Index Scan` Vì không có Index thích hợp để trợ giúp truy vấn này, công cụ cơ sở dữ liệu không có lựa chọn nào khác ngoài việc đọc mọi bản ghi trong bảng. Đây chính là lý do tại sao Số hàng được đọc là 1 triệu, tức là mọi hàng trong bảng.

- Number of rows read = 1000000
- Actual number of rows for all executions = 1

Chúng ta cần phải đọc 1 triệu bản ghi chỉ để nhận lại 1 kết quả, như vậy không đáng chút nào, giờ cùng xem nó chạy hết thời gian bao nhiêu nhé.

```sql
DECLARE @EndTime datetime
DECLARE @StartTime datetime
SELECT @StartTime=GETDATE()

-- Write Your Query
Select * from Employees where Name = 'ABC 992000'

SELECT @EndTime=GETDATE()

--This will return execution time of your query
SELECT DATEDIFF(ms,@StartTime,@EndTime) AS [Duration in millisecs]
```

![image](https://user-images.githubusercontent.com/72289126/157840695-8637de5e-c7a6-46e9-891d-e4b9aec2ff5d.png)

Rõ ràng là chậm hơn so với khi có index rồi, đây là ví dụ của 1 triệu bản ghi, trên thực tế ở những công ty lớn họ có hàng trăm triệu hoặc cả tỷ bản ghi cơ, thử tưởng tượng nếu ta không đánh index thì mỗi lần request phải đợi bao nhiêu lâu mới được response ?

### 2. Non-Clustered Index

Nonclustered Index được định nghĩa trên bảng trong đó dữ liệu có thể có cấu trúc phân cụm (clustered structure) hoặc ở dạng vun đống (heap).

Mỗi hàng index trong Nonclustered Index sẽ chứa một giá trị khóa nonclustered và một bộ định vị hàng.

Nonclustered Index thường áp dụng cho bảng chứa một lượng bản ghi nhỏ; một bảng có thể có nhiều Nonclustered Index.

Ta tạo 1 `Nonclustered Index` trên cột `Name` :

```sql
CREATE NONCLUSTERED INDEX IX_Employees_Name
ON [dbo].[Employees] ([Name])
```

![image](https://www.pragimtech.com/blog/contribute/article_images/1220210330203211/non-clustered-index-structure-in-sql-server.png)

Bây giờ thử với câu truy vấn tìm theo tên ở trên xem :

```sql
DECLARE @EndTime datetime
DECLARE @StartTime datetime
SELECT @StartTime=GETDATE()

-- Write Your Query
Select * from Employees where Name = 'ABC 992000'

SELECT @EndTime=GETDATE()

--This will return execution time of your query
SELECT DATEDIFF(ms,@StartTime,@EndTime) AS [Duration in millisecs]
```

![image](https://user-images.githubusercontent.com/72289126/157843218-7d9e8a68-770e-43bd-a16e-52b46cfc24ac.png)

Chú ý `Estimated Subtree Cost` khi có và không có `Estimated Subtree Cost` trên cột `Name`

![image](https://www.pragimtech.com/blog/contribute/article_images/1220210330203211/how-nonclustered-index-works.png)

#### So sánh Clustered Index và Non-clustered Index

![image](https://www.pragimtech.com/blog/contribute/article_images/1220210330203211/how-does-a-non-clustered-index-point-to-the-data.png)

- Ở `Clustered Index` thì các Index như `EmployeeId` thì được sắp xếp, còn `Non-clustered Index` thì không được sắp xếp.
- Ở `Clustered Index` thì số chỉ mục trên 1 bảng chỉ có 1 còn `Non-clustered Index` có thể có nhiều.
- Về tốc độ, thì việc tìm kiểm khi có `Clustered Index` chắc chắn sẽ nhanh hơn vì sắp xếp rồi mà, nhưng khi Insert,Update,Delete,... thì `Non-clustered Index` nhanh hơn do không phải quan tâm đến việc sắp xếp thứ tự lại.

### 3. B+ tree Index

B + Tree là một phương pháp tiên tiến của ISAM tổ chức tệp. Nó sử dụng cùng một khái niệm về khóa-chỉ mục, nhưng trong một cấu trúc giống như cây. Cây B + tương tự như cây tìm kiếm nhị phân, nhưng nó có thể có nhiều hơn hai nút lá. Nó chỉ lưu trữ tất cả các bản ghi ở nút lá. Các nút trung gian sẽ có các con trỏ đến các nút lá. Chúng không chứa bất kỳ dữ liệu / bản ghi nào.

Hãy xem xét một bảng sinh viên dưới đây. Khoá chính ở đây là STUDENT_ID. Và mỗi bản ghi chứa thông tin chi tiết của từng sinh viên cùng với giá trị khóa của nó và chỉ mục / con trỏ đến giá trị tiếp theo.

![image](https://www.tutorialcup.com/images/dbms/35/1.png?ezimgfmt=rs:373x249/rscb41/ng:webp/ngcb41)

Trong cây B +, nó có thể được biểu diễn như dưới đây.

![image](https://www.tutorialcup.com/images/dbms/35/2.png?ezimgfmt=rs:783x305/rscb41/ng:webp/ngcb41)

Xin lưu ý rằng nút lá 100 có nghĩa là, nó có tên và địa chỉ của sinh viên với ID 100, như chúng ta đã thấy ở trên.

Từ cấu trúc cây B + ở trên, rõ ràng là :

- Có một nút chính được gọi là gốc của cây - ở đây là 105 là gốc.
- Có một lớp trung gian với các nút. Họ không có hồ sơ thực tế được lưu trữ. Chúng đều là con trỏ tới nút lá. Chỉ nút lá chứa dữ liệu theo thứ tự đã được sắp xếp.
- Các nút bên trái của các nút gốc có giá trị trước của gốc và các nút bên phải có các giá trị tiếp theo của gốc. bên trái 105 là 102 và bên phải 105 là 108.
- Có một nút cuối cùng, được gọi là nút lá, chỉ có các giá trị : 100, 101, 103, 104, 106 và 107
- Tất cả các nút lá đều cân bằng - tất cả các nút lá ở cùng một khoảng cách từ nút gốc. Do đó tìm kiếm bất kỳ bản ghi nào dễ dàng hơn.
- Tìm kiếm bất kỳ bản ghi nào là tuyến tính trong trường hợp này. Bất kỳ bản ghi nào cũng có thể được duyệt qua một con đường duy nhất và được truy cập một cách dễ dàng.
- Vì các nút trung gian chỉ có con trỏ đến nút lá nên cấu trúc cây có chiều cao ngắn hơn. Chiều cao ngắn hơn, tốc độ truyền tải nhanh hơn và do đó truy xuất các bản ghi nhanh hơn.

Ưu điểm của B + Trees :

- Vì tất cả các bản ghi chỉ được lưu trữ trong nút lá và được sắp xếp theo danh sách liên kết tuần tự nên việc tìm kiếm trở nên rất dễ dàng.
- Sử dụng B +, chúng ta có thể truy xuất phạm vi hoặc truy xuất một phần. Đi ngang qua cấu trúc cây giúp việc này trở nên dễ dàng và nhanh chóng hơn.
- Khi số lượng bản ghi tăng / giảm, cấu trúc cây B + phát triển / thu hẹp. Không có giới hạn về kích thước cây B +, giống như chúng tôi có trong ISAM.
- Vì nó là một cấu trúc cây cân bằng, bất kỳ chèn / xóa / cập nhật nào không ảnh hưởng đến hiệu suất.
- Vì chúng ta có tất cả dữ liệu được lưu trữ trong các nút lá và việc phân nhánh nhiều hơn các nút bên trong làm cho chiều cao của cây ngắn hơn. Điều này làm giảm I / O đĩa. Do đó nó hoạt động tốt trong các thiết bị lưu trữ thứ cấp.

Nhược điểm :

- Nếu có bất kỳ sự sắp xếp lại nào của các node trong khi chèn hoặc xóa, thì sẽ tốn chút ít công sức, thời gian và không gian. Nhưng điều đó không đáng so với tốc độ nó đem lại.

Cú pháp:

`CREATE INDEX id_index ON table_name (column_name[, column_name…]) USING BTREE;`

Những đặc điểm của BTREE index:

- Dữ liệu index được tổ chức và lưu trữ theo dạng tree, tức là có root, branch, leaf. Giá trị của các node được tổ chức tăng dần từ trái qua phải.
- Btree index được sử dụng trong các biểu thức so sánh dạng: =, >, >=, <, <=, BETWEEN, LIKE.
- Khi truy vấn dữ liệu thì việc tìm kiếm trong BTREE là 1 quá trình đệ quy, bắt đầu từ root node và tìm kiếm tới branch và leaf, đến khi tìm được tất cả dữ liệu thỏa mãn với điều kiện truy vấn thì mới dùng lại.
- Btree index được sử dụng cho những column trong bảng khi muốn tìm kiếm 1 giá trị nằm trong khoảng nào đó. Ví dụ: tìm kiếm những sinh viên có điểm Toán từ 5-9

### 4. Hash Index

Hash index dựa trên giải thuật Hash Function (hàm băm). Tương ứng với mỗi khối dữ liệu (index) sẽ sinh ra một bucket key (giá trị băm) để phân biệt.

![image](https://images.viblo.asia/8ffa848f-00a3-4b2f-9184-fbe63c8ce58c.png)

Cú pháp:

`CREATE INDEX id_index ON table_name (column_name[, column_name…]) USING HASH;`

Kỹ thuật Hashing được chia thành 2 loại là: Static Hashing và Dynamic Hashing

#### Static Hashing

Trong static hashing, với mỗi key được cung cấp, hash function luôn tính toán và trả về một giá trị địa chỉ giống nhau. Sẽ không có bất kỳ thay đổi nào với địa chỉ bucket này.

Các hoạt động trong static hashing:

- Thêm mới: Khi một record được thêm mới, hash function sẽ generate một bucket address mới
- Tìm kiếm: Khi muốn tìm kiếm record với key đã được đánh index, hash function sẽ sử dụng key đó là params để tính toán và trả về địa chỉ của bucket tương ứng, trong bucket này sẽ chứa value là tất cả con trỏ tới record
- Xóa: Khi muốn xóa một record, hash function sẽ tính toán và tìm bản ghi cần xóa và sau đó xóa nó khỏi bộ nhớ

Bây giờ bạn muốn thêm mới một record nhưng lúc này địa chỉ bucket được hash function tính toán đã được sử dụng. Lúc này sẽ xử lý như thế nào?

Chúng ta có 2 phương pháp để giải quyết:

##### 1.Open Hashing

Với Open hashing, data block phù hợp kế tiếp sẽ được sử dụng để thêm mới record thay vì ghi đè lên address cũ.

![image](https://images.viblo.asia/2a6510f7-eda6-4743-bdeb-94ee5c53af7f.png)

Ví dụ, D3 là record muốn thêm vào, nhưng hash function lại tính toán đến địa chỉ 105 đã được sử dụng. Lúc này, hệ thống sẽ tìm kiếm để bucket khả dụng tiếp theo là 123, và assign D3 cho nó.

##### 2.Closed hashing

Với Closed hashing, một cái data bucket mới sẽ được phân bố cùng địa chỉ và được liên kết sau bucket cũ.

![image](https://images.viblo.asia/2df31df4-e5ed-4191-8b47-ff028f64c86f.png)

Ví dụ, chúng ta muốn thêm D3 vào table. Nhưng hash function tính toán địa chỉ cho D3 là 105 và bucket này đã được sử dụng. Lúc này, một butket mới sẽ được thêm vào và liên kết với butket 105, D3 sẽ được assign cho bucket này.

#### Dynamic Hashing

Hạn chế của Static Hashing là không mở rộng hoặc thu hẹp buckets một cách linh hoạt khi kích thước dữ liệu của database thay đổi.

Với Dynamic Hashing, các data buckets sẽ được tăng hoặc giảm linh hoạt dựa theo sự theo đổi kích thước của database.

Trong Dynamic hashing, hash function sẽ tính toán và tạo ra một tập rất nhiều các giá trị.

Ví dụ, chúng ta có 3 key D1, D2, D3, hash function sẽ generate ra 3 address tương ứng 1001, 0101, 1010. Với phương pháp lưu trữ này, chỉ xét một phần của địa chỉ nêu trên, lúc đầu buckets chỉ sử dụng một bit để lưu trữ data. Vì vậy nó sẽ cố gắng load 3 giá trị kia vào địa chỉ 0 hoặc 1

![image](https://images.viblo.asia/9b34737d-418a-41a9-9037-12892b4a6628.png)

Vậy thì D3 sẽ lưu ở đâu?

Lúc này buckets sẽ được tự động tăng thêm 1 bit và nó sẽ sử dụng 2 bit để lưu trữ, update những giá trị đã tồn tại vào buckets 2 bit và lưu lại giá trị mới thêm vào

![image](https://images.viblo.asia/45ce8ead-0d54-4a55-923c-a409d986a4f0.png)

Những đặc điểm của HASH index:

- Hash index có một vài đặc điểm khác biệt so với Btree index.
- Dữ liệu index được tổ chức theo dạng Key - Value được liên kết với nhau.
- Khác với BTREE, thì HASH index chỉ nên sử dụng trong các biểu thức toán tử là = và <>. Không sử dụng cho toán từ tìm kiếm 1 khoảng giá trị như > hay < .
- Không thể tối ưu hóa toán tử ORDER BY bằng việc sử dụng HASH index bởi vì nó không thể tìm kiếm được phần từ tiếp theo trong Order.
- Toàn bộ nội dung của Key được sử dụng để tìm kiếm giá trị records, khác với BTREE một phần của node cũng có thể được sử dụng để tìm kiếm.
- Hash có tốc độ nhanh hơn kiểu Btree.
