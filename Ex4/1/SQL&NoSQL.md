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

## So sánh SQL và NoSQL

| Tham số                   | SQL                                                                                                                                                     | NoSQL                                                                                                                                                                                          |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Định nghĩa                | Cơ sở dữ liệu SQL chủ yếu được gọi là RDBMS hoặc Cơ sở dữ liệu quan hệ                                                                                  | Cơ sở dữ liệu NoSQL chủ yếu được gọi là cơ sở dữ liệu không liên quan hoặc phân tán                                                                                                            |
| Design for                | RDBMS truyền thống sử dụng cú pháp và truy vấn SQL để phân tích và lấy dữ liệu để có thêm thông tin chi tiết. Chúng được sử dụng cho các hệ thống OLAP. | Hệ thống cơ sở dữ liệu NoSQL bao gồm nhiều loại công nghệ cơ sở dữ liệu khác nhau. Các cơ sở dữ liệu này được phát triển để đáp ứng nhu cầu trình bày cho sự phát triển của ứng dụng hiện đại. |
| Ngôn ngữ Query            | Structured query language (SQL)                                                                                                                         | Không có ngôn ngữ query                                                                                                                                                                        |
| Type                      | SQL databases là cơ sở dữ liệu dựa trên bảng                                                                                                            | NoSQL databases có thể dựa trên tài liệu, cặp khóa-giá trị, cơ sở dữ liệu biểu đồ                                                                                                              |
| Schema                    | SQL databases có lược đồ được xác định trước                                                                                                            | NoSQL databases sử dụng lược đồ động cho dữ liệu phi cấu trúc.                                                                                                                                 |
| Khả năng mở rộng          | SQL databases có thể mở rộng theo chiều dọc                                                                                                             | NoSQL databases có thể mở rộng theo chiều ngang                                                                                                                                                |
| Ví dụ                     | Oracle, Postgres, and MS-SQL.                                                                                                                           | MongoDB, Redis, , Neo4j, Cassandra, Hbase.                                                                                                                                                     |
| Phù hợp cho               | Đây là 1 lựa chọn lý tưởng cho môi trường truy vấn phức tạp                                                                                             | Không phù hợp với truy vấn phức tạp                                                                                                                                                            |
| Lưu trữ dữ liệu phân cấp  | SQL databases không thích hợp cho việc lưu trữ dữ liệu phân cấp.                                                                                        | Phù hợp hơn cho kho lưu trữ dữ liệu phân cấp vì nó hỗ trợ phương thức cặp khóa-giá trị.                                                                                                        |
| Variations                | Một loại có biến thể nhỏ                                                                                                                                | Nhiều loại khác nhau bao gồm các kho khóa-giá trị, cơ sở dữ liệu tài liệu và cơ sở dữ liệu đồ thị.                                                                                             |
| Năm phát triển            | Nó được phát triển vào những năm 1970 để giải quyết các vấn đề với lưu trữ tệp phẳng                                                                    | Được phát triển vào cuối những năm 2000 để khắc phục các vấn đề và hạn chế của SQL databases.                                                                                                  |
| Open-source               | Một sự kết hợp của mã nguồn mở như Postgres & MySQL, và thương mại như Oracle Database.                                                                 | Open-source                                                                                                                                                                                    |
| Tính nhất quán            | Nó phải được cấu hình cho sự nhất quán chặt chẽ.                                                                                                        | Nó phụ thuộc vào DBMS như một số cung cấp tính nhất quán mạnh mẽ như MongoDB, trong khi những người khác cung cấp chỉ cung cấp sự nhất quán cuối cùng, như Cassandra.                          |
| Được sử dụng tốt nhất cho | RDBMS database là tùy chọn thích hợp để giải quyết các vấn đề về ACID.                                                                                  | NoSQL được sử dụng tốt nhất để giải quyết các vấn đề về tính khả dụng của dữ liệu                                                                                                              |
| Tầm quan trọng            | Nó nên được sử dụng khi hiệu lực dữ liệu là siêu quan trọng                                                                                             | Sử dụng khi nó quan trọng hơn để có dữ liệu nhanh hơn dữ liệu chính xác                                                                                                                        |
| Lựa chọn tốt nhất         | Khi bạn cần hỗ trợ truy vấn động                                                                                                                        | Sử dụng khi bạn cần mở rộng quy mô dựa trên yêu cầu thay đổi                                                                                                                                   |
| Hardware                  | Specialized DB hardware (Oracle Exadata, etc.)                                                                                                          | Commodity hardware                                                                                                                                                                             |
| Network                   | Highly available network (Infiniband, Fabric Path, etc.)                                                                                                | Commodity network (Ethernet, etc.)                                                                                                                                                             |
| Loại lưu trữ              | Highly Available Storage (SAN, RAID, etc.)                                                                                                              | Commodity drives storage (standard HDDs, JBOD)                                                                                                                                                 |
| Tính năng tốt nhất        | Hỗ trợ đa nền tảng, Bảo mật và miễn phí                                                                                                                 | Dễ sử dụng, hiệu suất cao và công cụ linh hoạt.                                                                                                                                                |
| Mô hình ACID và BASE      | ACID (Atomicity, nhất quán, cách ly và độ bền) là một chuẩn cho RDBMS                                                                                   | Cơ bản (Về cơ bản có sẵn, trạng thái mềm, phù hợp cuối cùng) là một mô hình của nhiều hệ thống NoSQL                                                                                           |
| Performance               | SQL hoạt động tốt và nhanh thì việc desgin tốt là cực kì quan trọng và ngược lại.                                                                       | Nhanh hơn SQL NoSQL thì denormalized cho phép bạn lấy được tất cả thông tin về một item cụ thể với các codition mà không cần JOIN liên quan hoặc truy vấn SQL phức tạp.                        |
| Kết luận                  | Dự án đã có yêu cầu dữ liệu rõ ràng xác định quan hệ logic có thể được xác định trước.                                                                  | Phù hợp với những dự án yêu cầu dữ liệu không liên quan, khó xác định, đơn giản mềm dẻo khi đang phát triển                                                                                    |

## Các loại NoSQL DB

### 1.Document

Document store được gọi là các cơ sở dữ liệu hướng tài liệu, một thiết kế riêng biệt cho việc lưu trữ tài liệu dạng văn kiện JSON, BSON hoặc XML. Vì là cấu trúc dữ liệu không ràng buộc khác với SQL, các CSDL này không đòi hỏi người dùng tự tạo bảng nhập liệu trước khi nhập dữ liệu vào. Các tài liệu có thể chứa bất kì dữ liệu nào. CSDL dạng này có các cặp khoá – giá trị nhưng cũng có đính kèm các trị số siêu dữ liệu (*metadata) giúp việc truy vấn (*query) dễ dàng hơn.

![image](https://webassets.mongodb.com/_com_assets/cms/Relational_vs_DocumentDB-imgngssl17.png)

Trong sơ đồ bên trái bạn có thể thấy các hàng và cột, và ở bên phải, có một cơ sở dữ liệu tài liệu có cấu trúc tương tự như JSON. Đối với cơ sở dữ liệu quan hệ, bạn phải biết bạn có những cột nào, v.v. Tuy nhiên, đối với NoSQL document database, bạn có kho dữ liệu như đối tượng JSON. Bạn không cần phải xác định cái nào làm cho nó linh hoạt. Loại document này chủ yếu được sử dụng cho các hệ thống CMS, nền tảng blog, phân tích thời gian thực và các ứng dụng thương mại điện tử. Document database không nên sử dụng cho các giao dịch phức tạp yêu cầu nhiều hoạt động hoặc truy vấn dựa trên các cấu trúc tổng hợp khác nhau.

#### Điểm mạnh

CSDL hướng tài liệu rất linh hoạt, có thể xử lí dữ liệu nửa cấu trúc và không cấu trúc rất tốt. Người dùng không cần quan tâm tới dạng dữ liệu khi setup, điều này tốt trong trường hợp bạn không lường trước được dạng dữ liệu nào bạn sẽ cần lưu trữ.

Người dùng có thể thiết kế một cấu trúc cho một tài liệu cụ thể mà không ảnh hưởng tới các tài liệu khác. Schema cho CSDL cũng có thể được tuỳ chỉnh mà không gây ra thời gian downtime, giúp đem đến high availability (tính sẵn sàng cao). Thời gian ghi dữ liệu cũng rất nhanh.

Ngoài tính linh hoạt, các lập trình viên còn ưa chuộng document store bởi tính dễ dàng mở rộng theo chiều ngang của chúng. Quá trình chia sẻ cũng dễ hiểu và dễ thao tác hơn so với CSDL quan hệ, nên document store có thể mở rộng nhanh và dễ dàng.

#### Điểm yếu

CSDL dạng lưu trữ tài liệu hy sinh các yếu tố ACID để đổi lấy sự linh hoạt. Ngoài ra, việc truỵ vấn chỉ có thể được thực hiện trong từng tài liệu, không thể truy vấn dữ liệu trên nhiều tài liệu khác nhau.

Nên sử dụng CSDL dạng lưu trữ tài liệu trong các trường hợp nào?

Dữ liệu phi cấu trúc hoặc không có cấu trúc
Quản lý nội dung
Phân tích dữ liệu chuyên sâu
Tạo mẫu nhanh

Các hệ thống DBMS Document database NoSQL tiêu biểu:

- Amazon SimpleDB
- CouchDB
- MongoDB
- Riak
- Lotus Notes

### 2.Key-Value

Với Key Value, Dữ liệu được lưu trữ trong các cặp khóa / giá trị (Key/Value Pair). Nó được thiết kế theo cách để xử lý nhiều dữ liệu và tải nặng. Cơ sở dữ liệu lưu trữ cặp khóa-giá trị lưu trữ dữ liệu dưới dạng bảng băm trong đó mỗi khóa là duy nhất và giá trị có thể là JSON, BLOB (Binary Large Objects), chuỗi, v.v.

![image](https://upload.wikimedia.org/wikipedia/commons/5/5b/KeyValue.PNG)

Đây là một trong những ví dụ cơ sở dữ liệu NoSQL cơ bản nhất. Loại cơ sở dữ liệu NoSQL này được sử dụng như một bộ sưu tập, từ điển, mảng kết hợp, v.v. Key value stores giúp developer lưu trữ dữ liệu không có schema.

#### Điểm mạnh

Dạng CSDL này có rất nhiều lợi thế. Nó rất linh hoạt và có thể xử lí nhiều loại dữ liệu một cách nhanh chóng. Các chìa khoá được dùng để truy xuất thẳng tới các giá trị tìm kiếm, mà không cần thông qua quá trình index (quá trình tìm kiếm dữ liệu và đánh giá độ chính xác của dữ liệu đó của hệ thống CSDL), giúp quá trình tìm kiếm diễn ra nhanh chóng. Tính linh động cũng là một điểm mạnh của CSDL dạng này: lưu trữ key – value có thể được chuyển từ hệ thống này sang hệ thống khác mà không cần code lại. Cuối cùng, CSDL key – value có thể mở rộng theo chiều ngang dễ dàng và chi phí vận hành thấp.

#### Điểm yếu

Tính linh hoạt của CSDL dạng key – value bị đánh đổi bởi tính chính xác. Hầu như rất khó để truy xuất giá trị chính xác từ CSDL dạng này vì dữ liệu được lưu trữ theo blob, nên kết quả trả về hầu như đều theo blob. Điều này gây ra khó khan khi báo cáo số liệu hoặc cần chỉnh sửa một phần của các giá trị. Cuối cùng, không phải objects nào cũng có thể được cấu hình thành cặp chìa khoá – giá trị được.

### 3.Wide-column (column family)

![image](https://thehonestcoder.com/wp-content/uploads/2020/11/choosing-the-right-database-wide-column-3-1-1024x623.png)

Mô hình wide – column là một dạng lưu CSDL phi quan hệ lưu trữ theo dạng cột. Mô hình này có vài điểm tương đồng với mô hình key – value nhưng cũng có vài tính chất của dạng CSDL quan hệ.

Mô hình wide – column dựa trên khái niệm keyspace thay vì schema. Một keyspace bao gồm nhiều cụm column (tương tự như table nhưng linh hoạt hơn về cấu trúc), mỗi cụm bao gồm nhiều hàng và nhiều cột riêng biệt. Mỗi hàng không cần phải có số lượng hoặc loại cột. Một timestamp quyết định phiên bản gần nhất của data.

Với Wide-column dữ liệu được lưu trữ trong database dưới dạng các cột. Mỗi cột được xử lý riêng biệt. Giá trị của cơ sở dữ liệu cột đơn được lưu trữ liền kề.

#### Điểm mạnh

Loại CSDL này có cả lợi ích của CSDL quan hệ và phi quan hệ, có thể xử lí dữ liệu cấu trúc và phi cấu trúc, đồng thời cũng dễ dàng nâng cấp. So với CSDL quan hệ, khả năng mở rộng theo chiều ngang cũng dễ dàng và nhanh chóng hơn.

CSDL dạng cột có khả năng nén tốt hơn CSDL dạng dòng. Đồng thời, data set lớn có thể dễ dàng duyệt hơn. Mô hình wide – column có khả năng xử lí tốt các yêu cầu truy xuất tập trung.

#### Điểm yếu

CSDL dạng cột dễ dàng update theo cụm, bù lại việc upload và update số liệu cá nhân rất khó. Cộng thêm thực tế là mô hình wide – column chậm hơn so với CSDL quan hệ khi xử lí các giao dịch.

### 4.Graph

![image](https://www.profium.com/wp-content/uploads/2020/01/graph-database.jpg)

Cơ sở dữ liệu kiểu đồ thị (Graph Based) lưu trữ các thực thể cũng như các mối quan hệ giữa các thực thể đó. Thực thể được lưu trữ dưới dạng một node với mối quan hệ là các cạnh. Một cạnh cho biết mối quan hệ giữa các node. Mỗi node và cạnh có một mã định danh duy nhất.

#### Điểm mạnh

So với cơ sở dữ liệu quan hệ trong đó các bảng được kết nối với nhau một cách lỏng lẻo, cơ sở dữ liệu Đồ thị có bản chất là đa quan hệ. Mối quan hệ truyền tải nhanh chóng vì chúng đã được ghi lại vào DB và không cần phải tính toán chúng. Cơ sở dữ liệu đồ thị chủ yếu được sử dụng cho mạng xã hội, hậu cần, dữ liệu không gian.

#### Điểm yếu

Thiếu tính đồng thời hiệu suất cao (high performance concurrency): Trong nhiều trường hợp, graph database cung cấp các kiểu đọc và kiểu ghi đơn, điều này cản trở sự đồng thời và hiệu suất, do đó phần nào hạn chế tính song song phân luồng (threaded parallelism).

Thiếu ngôn ngữ chuẩn: Việc thiếu sự thiết lập và một ngôn ngữ khai báo chuẩn là một vấn của NoSQL graph database.

Thiếu tính song song (parallelism): việc phân vùng một biểu đồ là một vấn đề. Hầu hết các graph database không cung cấp các truy vấn song song trên các biểu đồ lớn.
