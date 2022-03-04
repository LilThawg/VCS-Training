# Tìm hiểu về OpenAPI và Swagger

## OpenAPI là gì ?

OpenAPI còn được gọi là Public API, nghe tên gọi ta có thể lờ mờ hình dung đây là API mở. Cụ thể hơn nó là một giao diện lập trình ứng dụng được cung cấp công khai cho các dev.Các API mở được xuất bản trên internet và được chia sẻ miễn phí, cho phép chủ sở hữu của dịch vụ có thể truy cập mạng cấp quyền truy cập phổ biến cho người tiêu dùng.

Open API được biết đến với tên đầy đủ là Open API Specification. Đây được biết đến là một loại định dạng dùng để mô tả API cho một Rest APIs hiện nay. Với duy nhất một file Open API, bạn sẽ có thể mô tả được toàn bộ API. Điều này có nghĩa là việc mô tả sẽ bao gồm cả những nội dung sau đây:

- Tạo điều kiện hoạt động các các endpoints hay là các users cũng như cách thức hoạt động của các endpoints đó như get/users, post/users, ...
- Hiển thị rõ được các tham số về đầu vào và đầu ra của mỗi hoạt động.
- Thể hiện các phương thức xác thực được sử dụng.
- Thể hiện các thông tin liên lạc, các chứng chỉ, những điều khoản liên quan đến việc sử dụng và các thông tin liên quan khác.

Thực tế thì API Specifications hiện có thể được viết bằng các định dạng như JSOL hay YAML. Đây được biết đến là hai định dạng có lợi cho cả người dùng và hệ thống máy tính khi rất dễ đọc, dễ hiểu để sử dụng.

## Swagger là gì ?

Nếu như Open API có ý nghĩa quan trọng như trên thì điều gì đã tạo nên được bộ mô tả này? Và đó chính là Swagger. Swagger được hiểu là một công cụ có mã nguồn mở và dùng để xây dựng nên Open API Specifications. Công cụ này sẽ giúp bạn trong việc thiết kế, tạo dựng các tài liệu cũng như việc sử dụng Rest APIs. Hiểu đơn giản nó giống với POSTMAN nhưng có hướng dẫn sử dụng từng API cụ thể, chi tiết hơn :) 

Với các nhà phát triển, khi sử dụng Swagger sẽ được cung cấp 3 tool chính như sau: 

- Tool Swagger Editor: Được sử dụng để thiết kế, xây dựng nên các APIs một cách mới hoàn toàn hoặc là có thể edit lại từ những APIs có sẵn với việc tận dụng một file config.
- Tool Swagger Codegen: Có tác dụng trong việc generate ra code thông qua sử dụng các file config sẵn có trước đó.
- Tool Swagger UI: Ứng dụng trong việc generate các file ra HTML, CSS,...xuất phát từ một file config.

Với những tool được liệt kê ở trên của Swagger thì Swagger UI được biết đến là một tool có sự thông dụng phổ biến nhất hiện nay. Với tool Swagger UI, tool này có ứng dụng rất lớn trong việc xây dựng giao diện cho các tài liệu bắt nguồn từ file config áp dụng dưới chuẩn của Ipen API. Giao diện được tạo ra bởi tool này thường có tính tường minh và khá rõ ràng, hiện ra một cách cụ thể nhất cho các nhà phát triển. Điều này sẽ giúp ích rất lớn cho cả người dùng lẫn các lập trình viên trong việc đọc hiểu và sử dụng. Thêm vào đó, đây cũng là dẫn chứng có việc các developers sử dụng file config nhưng lại có sự tách biệt một cách hoàn toàn giữa các tác vụ với nhau.

Mỗi API được sử dụng trong quá trình này sẽ cho chúng ta biết một cách chính xác nhất về nguồn vào và nguồn ra một cách chi tiết. Thêm vào đó chính là việc những trường cần phải được gửi lên hệ thống cũng như các trạng thái kết quả có thể được trả về. Điều đặc biệt nhất có lẽ chính là việc ta có thể đưa các dữ liệu vào trong để test thử các kết quả liệu có thực sự chính xác và đảm bảo tính đúng đắn hay không. 

## Cấu trúc cơ bản của API

### 1. Metadata hay Info

Hầu hết, mỗi Open API Specifications được sử dụng đều sẽ bắt đầu với từ khóa “Open API” nhằm mục đích cho việc khai báo tên của phiên bản đó. Với loại phiên bản được sử dụng sẽ có ý nghĩa trong việc định nghĩa lại toàn bộ những cấu trúc ở trong API. Phân info sẽ có các thông tin cơ bản về API như title, description và các version. Cụ thể thì:

- `Title` sẽ là tên mà bạn đặt cho API của mình.
- `Description` chính là các thông tin về API của bạn được đưa ra một cách chi tiết avf xuất hiện ở nhiều mặt cụ thể khác nhau. Với việc mô tả này thì bạn hoàn toàn có thể viết thành nhiều dòng nếu như quá dài và sử dụng cú pháp hỗ trợ như markdown.
- `Version` là phiên bản được sử dụng trong quá trình tạo dựng của bạn với API.

Metadata hay Info sẽ có chức năng trong việc hỗ trợ đưa ra các từ khóa về các thông tin liên quan như thông tin liên lạc, thông tin về chứng chỉ, các điều khoản trong việc sử dụng,...

### 2.Servers

Để có thể test được các API thì bạn cần có một đường dẫn liên quan đến servers. Và đây chính là phần tạo ra một đường dẫn cụ thể của servers được sử dụng để thực hiện chức năng trên. Trong phần này, việc định nghĩa một hay là nhiều các servers khác nhau sẽ hoàn toàn tùy thuộc vào quyết định của bạn.

### 3.Paths

Đây là phần trọng tâm của API. Với phần này, nhiệm vụ của bạn và những nhà lập trình viên khác chính là định nghĩa những paths xuất hiện trong API hay các phương thức, những tham số cụ thể tồn tại trong API đó.

- Phần này sẽ bắt đầu bằng từ khóa paths
- Sau đó là đến những path trong API (/user/{userId})
- Tiếp đến là phương thức của API (GET, POST, DELETE, PUT ...)
- summary là phần mô tả tóm tắt của API
- parameters: sẽ là những tham số truyền vào API. Bạn có thể set tham số required hay không, mô tả nó (description) hoặc validate. Đặc biệt trong phần này. bạn có thể chỉ định 1 schema (hiểu nôm na là 1 Model) để có thể định nghĩa cho phần tham số thông qua schema & $ref
- response là phần trả về của server. Bạn có thể định nghĩa những HTTP code: 200, 404, 500 ... với những mô tả cho từng trường hợp.

Thêm vào đó, ở phần này, các parameters sẽ sở hữu khá nhiều loại khai báo khác nhau đứng sau từ khóa “in” mà chúng ta đều cần phải lưu ý:

- In body: Giúp cho người dùng có một không gian để tạo ra một dữ liệu đầu vào. Ở không gian này, người dùng hoàn toàn có thể nhập  các dữ liệu về những yêu cầu cơ bản vào trong đó.
- In formData: Giúp cho người dùng tạ được input đã được định trước để có thể thực hiện nhập các dữ liệu yêu cầu theo từng miền đã được định sẵn.
- In path: Sử dụng trong việc tạo lập một input vào trong giá trị để khai báo trong các routers, thông thường được gọi là ID. 
- In query: Được sử dụng trong việc tạo input nhập vào các giá trị thống kê theo các miền định sẵn để dùng trong việc gửi các query request.
- In header: Thực hiện việc dùng để khai báo các giá trị có trong header của yêu cầu mà bạn cần thực hiện truyền tải lên. 

### 4.Schema

Nói một cách dễ hiểu và đơn giản nhất thì Schema có thể được định nghĩa như một model. Schema sẽ được thực hiện phần khai báo thông qua việc sử dụng từ khóa `component` và `schemas`(Lưu ý: những chỗ gọi đến schema này phải chỉ định chính xác đường dẫn VD: `$ref: "#/components/schemas/User"`). 

- Tham số đầu tiên là tên của Model (User)
- Tiếp đó sẽ là phần kiểu định dạng (object)
- Sau đó là phần thuộc tính của Model này

