# HTTP

## Giới thiệu

HTTP (HyperText Transfer Protocol – Giao thức truyền tải siêu văn bản) là một trong các giao thức chuẩn về mạng Internet, được dùng để liên hệ thông tin giữa Máy cung cấp dịch vụ (Web server) và Máy sử dụng dịch vụ (Web client), là giao thức Client/Server dùng cho World Wide Web – WWW

Giao thức HTTP là một giao thức tầng ứng dụng của bộ giao thức TCP/IP (các giao thức nền tảng cho Internet).

## Sơ đồ hoạt động của HTTP

Giao thức HTTP hoạt động dựa trên mô hình Client – Server. Thông thường khi các bạn lướt web, các máy tính của người dùng sẽ đóng vai trò làm máy khách (Client). Sau một thao tác nào đó của người dùng, các máy khách sẽ gửi yêu cầu đến máy chủ (Server) và chờ đợi câu trả lời từ những máy chủ này.

![image](https://images.viblo.asia/7da268f1-718b-465c-87df-700e766df185.png)

HTTP là một stateless protocol. Hay nói cách khác, request hiện tại không biết những gì đã hoàn thành trong request trước đó.

HTTP cho phép tạo các yêu cầu gửi và nhận các kiểu dữ liệu, do đó cho phép xây dựng hệ thống độc lập với dữ liệu được chuyển giao.

Ngoài ra, khi các hệ thống trao đổi dữ liệu với nhau, chúng cũng sử dụng giao thức này nhưng 2 bên đều là server.

## Các thành phần chính của HTTP

### HTTP - Requests

Một HTTP client (máy khách) gửi một HTTP request (yêu cầu) lên server (máy chủ) nhờ một thông điệp có định dạng như sau:

```
<method> <request-URL> <http-serverion>
<headers>
<body>
```

VD:

![image](https://khuenguyencreator.com/wp-content/uploads/2021/09/http-request.jpg)

**1. Request Line**

Bắt đầu của HTTP Request sẽ là dòng Request-Line bao gồm 3 thông tin:
* Method
* Request URL
* HTTP version

**Method**

Báo cho Server rằng hành động sẽ phải xử lý với thông tin được gửi từ client lên.

Giao thức HTTP định nghĩa một tập các phương thức request, client có thể sử dụng một trong các phương thức này để tạo request tới HTTP server, dưới đây liệt kê một số phương thức phổ biến.

![image](https://images.viblo.asia/b986dced-c499-4051-8efb-5ea5d9b93c02.png)

**Request URL**

Một URL (Uniform Resource Locator) được sử dụng để xác định duy nhất một tài nguyên trên Web. Một URL có cấu trúc như sau:

`protocol://hostname:port/path-and-file-name`

Trong một URL có 4 thành phần:

* Protocol: giao thức tầng ứng dụng được sử dụng bởi client và server
* Hostname: tên DNS domain
* Port: Cổng TCP để server lắng nghe request từ client
* Path-and-file-name: Tên và vị trí của tài nguyên yêu cầu

**HTTP version**

HTTP version là Phiên bản giao thức HTTP đang sử dụng.

Các phiên bản của HTTP gồm: `0.9`, `1.0`, `1.1` và `2.0`.

**2. Request header**

Tiếp theo dòng Request-Line là các trường Request-header, cho phép client gửi thêm các thông tin bổ sung về thông điệp HTTP request và về chính client. Một số trường thông dụng như:

* Accept loại nội dung có thể nhận được từ thông điệp response. Ví dụ: text/plain, text/html…
* Accept-Encoding: các kiểu nén được chấp nhận. Ví dụ: gzip, deflate, xz, exi…
* Connection: tùy chọn điều khiển cho kết nối hiện thời. Ví dụ: keepalive,Upgrade…
* Cookie: thông tin HTTP Cookie từ server
* User-Agent: thông tin về user agent của người dùng

**3. Request Body**

Body là dữ liệu mà Client sẽ gửi lên Server

Một request body có thể là một đoạn text thuần túy, HTML, XML, JSON, Javascript, hoặc một tập các cặp key-value dạng form-data.

Khi bạn mở một web, trình duyệt sẽ nhận payload dạng HTML, nó chính là giao diện mà chúng ta quan sát được trên trình duyệt đó.

Thông thường khi làm việc với các HTTP APIs chúng ta sẽ gửi và nhận các payload dạng JSON hoặc XML.

Không phải tất cả các message HTTP đều phải có payload: POST và PUT có thể có, còn với GET và DELETE thì có thể không có payload.

Sau khi request server sẽ xử lý và phản hổi (response) tới client theo 1 trong 3 phương pháp:
* Server phân tích request nhận được, maps yêu cầu với tập tin trong tập tài liệu của server, và trả lại tập tin yêu cầu cho client.
* Server phân tích request nhận được, maps yêu cầu vào một chương trình trên server, thực thi chương trình và trả lại kết quả của chương trình đó.
* Request từ client không thể đáp ứng, server trả lại thông báo lỗi.

### HTTP – Responses

HTTP response là bản tin trả về từ server sang client, trong đó sẽ có các trường thông tin mà request yêu cầu.

Định dạng gói tin HTTP response cũng gồm 3 phần chính là: Status line, Header và Body.

![image](https://images.viblo.asia/8414d386-f4e5-4b9c-aded-d3b379dc7c20.png)

**1. Response Status line**

Gồm 3 trường là phiên bản giao thức (HTTP version), mã trạng thái (Status code) và mô tả  trạng thái (Status text):

* Phiên bản giao thức (HTTP version): phiên bản của giao thức HTTP mà server hỗ trợ, thường là HTTP/1.0 hoặc HTTP/1.1
* Mã trạng thái (Status code): mô tả trạng thái kết nối dưới dạng số, mỗi trạng thái sẽ được biểu thị bởi một số nguyên. Ví dụ: 200, 404, 302,…
* Mô tả trạng thái (Status text): mô tả trạng thái kết nối dưới dạng văn bản một cách ngắn gọn, giúp người dùng dễ hiểu hơn so với mã trạng thái. Ví du: 200 OK, 404 Not Found, 403 Forbiden,…

Yếu tố Status-Code(mã trạng thái) là một số nguyên 3 ký tự, trong đó ký tự đầu tiên của mã hóa trạng thái định nghĩa hạng (loại) phản hồi và hai ký tự cuối không có bất cứ vai trò phân loại nào. Có 5 giá trị của ký tự đầu tiên như sau:

**1xx: Information Message (Thông tin)**

Nó nghĩa là yêu cầu đã được nhận và tiến trình đang tiếp tục.

* 100 (Continue) : Máy chủ trả về mã này để chỉ ra rằng nó đã nhận được một phần đầu tiên của một yêu cầu và được chờ đợi cho phần còn lại.
* 101 (Switching protocols) : Bên yêu cầu đã yêu cầu các máy chủ để chuyển đổi và máy chủ được thừa nhận rằng nó sẽ làm như vậy.

**2xx: Successful (Thành công)**

Nó nghĩa là hoạt động đã được nhận, được hiểu, và được chấp nhận một cách thành công.

* 200 (Successful) : Các máy chủ xử lý yêu cầu thành công
* 201 (Created) : Yêu cầu đã thành công và các máy chủ tạo ra một nguồn tài nguyên mới.
* 202 (Accepted) : Máy chủ đã chấp nhận yêu cầu, nhưng vẫn chưa xử lý nó
* 203 (Non-authoritative information) : Máy chủ đã chấp nhận yêu cầu, nhưng vẫn chưa xử lý nó
* 204 (No content) : Các máy chủ xử lý yêu cầu thành công, nhưng không trả lại bất kỳ nội dung nào
* 205 (Reset content) : Các máy chủ proccessed yêu cầu thành công, nhưng không trả lại bất kỳ nội dung. Không giống như một phản ứng 204, phản ứng này đòi hỏi người yêu cầu thiết lập lại xem tài liệu
* 206 (Partial content) : Các máy chủ xử lý thành công một phần của một yêu cầu

**3xx: Redirection (Sự điều hướng lại)**

Server thông báo cho client phải thực hiện thêm thao tác để hoàn tất request. Nó nghĩa là hoạt động phải được thực hiện để hoàn thành yêu cầu.

* 301 (Moved Permanently) : Các trang web yêu cầu đã bị di chuyển vĩnh viễn tới URL mới
* 302 (Moved temporarily) : Trang được yêu cầu đã di chuyển tạm thời tới một URL mới
* 303 (See other) : tài nguyên đã được chuyển tạm thời tới địa chỉ Location trong HTTP response.
* 304 (Not modified) : Các trang yêu cầu đã không được sửa đổi kể từ khi yêu cầu cuối cùng. Khi máy chủ trả về phản hồi này, nó không trả lại các nội dung của trang.

**4xx: Client error (Lỗi Client)**

Nó nghĩa là yêu cầu chứa cú pháp không chính xác hoặc không được thực hiện

* 400 (Bad request) : request không đúng dạng, cú pháp.
* 401 (Not authorized) : Đề nghị yêu cầu xác thực. Máy chủ có thể trả về phản hồi này yêu cầu xác thực đăng nhập tài khoản và mật khẩu (thông thường máy chủ trả về phản hồi này nếu client gửi request một trang đăng nhập)
* 403 (Forbidden) : Máy chủ từ chối yêu cầu.(thông thường nếu đăng nhập không thành công máy chủ sẽ trả về mã lỗi này)
* 404 (Not found) : Máy chủ không thể tìm thấy trang yêu cầu. Ví dụ, máy chủ thường trả về mã này nếu có 1 yêu cầu tới một trang không tồn tại trên máy chủ.
* 405 (Method not allowed) : Phương thức được xác định trong yêu cầu là không được cho phép.
* 406 (Not acceptable) : Server chỉ có thể tạo một phản hồi mà không được chấp nhận bởi Client.
* 407 (Proxy authentication required) : Yêu cầu client phải xác thực sử dụng một proxy. Khi máy chủ trả về phản hồi này, nó cũng chỉ ra proxy mà người yêu cầu phải sử dụng.
* 408 (Request timeout) : Request tốn thời gian dài hơn thời gian Server phản hồi
* 409 (Conflict) : Các máy chủ gặp phải một cuộc xung đột thực hiện yêu cầu. Các máy chủ phải bao gồm thông tin về các cuộc xung đột trong các phản ứng. Máy chủ có thể trả về mã này để đáp ứng với yêu cầu PUT xung đột với yêu cầu trước đó, cùng với một danh sách các sự khác biệt giữa các yêu cầu.
* 410 (Gone) : Các máy chủ trả về phản hồi này khi các nguồn tài nguyên yêu cầu đã bị loại bỏ vĩnh viễn. Nó tương tự như một 404 (Không tìm thấy) mã, nhưng đôi khi được sử dụng ở vị trí của một 404 cho nguồn lực được sử dụng để tồn tại nhưng không còn làm. Nếu tài nguyên đã di chuyển vĩnh viễn, bạn nên sử dụng một 301 để xác định vị trí mới của tài nguyên.
* 411 (Length required) : Content-Length không được xác định rõ. Server sẽ không chấp nhận yêu cầu mà không có nó
* 412 (Precondition failed) : Các máy chủ không đáp ứng một trong các điều kiện tiên quyết mà người yêu cầu đưa vào yêu cầu.
* 413 (Request entity too large) : Máy chủ không thể xử lý yêu cầu bởi vì nó là quá lớn đối với các máy chủ để xử lý.
* 414 (Requested URI is too long) : URI yêu cầu (thường là một URL) là quá dài đối với máy chủ để xử lý.
* 416 (Requested range not satisfiable) : Máy chủ trả về mã trạng thái này nếu yêu cầu cho một phạm vi không có sẵn cho trang.
* 417 (Expectation failed) : Máy chủ không thể đáp ứng yêu cầu của các trường yêu cầu, tiêu đề mong đợi.

**5xx: Server Error (Lỗi Server)**

Nó nghĩa là Server thất bại với việc thực hiện một yêu cầu nhìn như có vẻ khả thi.

* 500 (Internal server error) : Các máy chủ gặp lỗi và không thể thực hiện yêu cầu.
* 501 (Not implemented) : Các máy chủ không có các chức năng để thực hiện yêu cầu. Ví dụ, máy chủ có thể trả về mã này khi nó không nhận ra phương thức yêu cầu.
* 502 (Bad gateway) : Các máy chủ đã hoạt động như một gateway hoặc proxy và nhận được một phản ứng không hợp lệ từ máy chủ ngược.
* 503 (Service unavailable) : Máy chủ hiện không có sẵn (vì nó bị quá tải hoặc xuống để bảo trì). Nói chung, đây là một trạng thái tạm thời.
* 504 (Gateway timeout) : Các máy chủ đã hoạt động như một gateway hoặc proxy và đã không nhận được yêu cầu kịp thời từ máy chủ ngược.
* 505 (HTTP version not supported) : Các máy chủ không hỗ trợ phiên bản giao thức HTTP được sử dụng trong yêu cầu.

**2. Response Header**

Header của gói tin response có chức năng tương tựn như gói tin request, giúp server có thể truyền thêm các thông tin bổ sung đến client dưới dạng các cặp `{"Name":Value}`.

**3. Response Body**

Là nơi đóng gói dữ liệu để trả về cho client, thông thường trong duyệt web thì dữ liệu trả về sẽ ở dưới dạng một trang HTML để trình duyệt có thể thông dịch được và hiển thị ra cho người dùng, hoặc trả về dạng JSON, XML khi giao tiếp bằng API.


# RESTful API

## RESTful API là gì?

Để hiểu rõ RESTful API là gì, ta cần làm rõ hai khái niệm thành phần của nó. Bao gồm:

* `API` là viết tắt của **A**pplication **P**rograming **I**nterface (Giao diện lập trình ứng dụng). Thật ra cái giao diện này không phải cho người dùng cuối mà dành cho các nhà phát triển (developer). Thuật ngữ này dùng để chỉ phương thức kết nối từ ứng dụng đến các thư viện khác nhau.
* `REST` là từ viết tắt của **RE**presentational **S**tate **T**ransfer (Truyền trạng thái đại diện). Đây là một tiêu chuẩn dùng để xây dựng API, giúp API trở nên thân thiện hơn với người dùng. Khái niệm này được Roy Fielding giới thiệu lần đầu vào năm 2000 trong luận văn tiến sĩ nổi tiếng của ông.

Do đó, RESTful API (hay còn được gọi là REST API) là thuật ngữ dùng để chỉ những API được xây dựng theo cấu trúc REST. Mặc dù REST có thể được sử dụng với mọi giao thức, nhưng nó thường dùng HTTP cho Web API.

## Những ràng buộc kiến trúc cơ bản của RESTful API

Thực tế cho thấy, có rất nhiều API tự nhận mình là RESTful, nhưng chúng lại không đáp ứng được những ràng buộc mà tiến sĩ Fielding đặt ra. Một RESTful API phải thỏa mãn những ràng buộc (hay điều kiện) sau đây:

**1. Client–server (Máy khách – máy chủ)**

Ràng buộc này hoạt động dựa trên ý tưởng rằng máy khách và máy chủ phải hoàn toàn tách biệt và được phép phát triển riêng lẻ, độc lập. Phương thức hoạt động chính của REST là tách biệt giao diện người dùng ra khỏi dữ liệu lưu trữ.

Với cách thức này, người dùng có thể thực hiện thay đổi với các ứng dụng di động của mình một cách độc lập. Việc này không làm ảnh hưởng đến cấu trúc dữ liệu hoặc thiết kế cơ sở dữ liệu của máy chủ. Ngược lại, việc điều chỉnh cơ sở dữ liệu hoặc thay đổi ứng dụng của máy chủ cũng không ảnh hưởng đến ứng dụng của máy khách.

**2. Stateless (Phi trạng thái)**

Bất kỳ một RESTful API nào cũng phải phi trạng thái. Nghĩa là mỗi yêu cầu (request) từ máy khách đến máy chủ có thể được thực hiện độc lập. Đồng thời, mỗi yêu cầu phải chứa mọi thông tin cần thiết để máy chủ hiểu và hoàn thành nó.

Ngoài ra, yêu cầu của máy khách không thể lạm dụng bất kỳ thông tin nào trên máy chủ. Đó cũng chính là lý do khiến trạng thái phiên (session state) được giữ hoàn toàn trên máy khách. Điều này sẽ giúp tăng độ tin cậy cho API, hạn chế lỗi và giảm tài nguyên sử dụng.

**3. Cacheable (Lưu được vào bộ nhớ cache)**

API phi trạng thái có thể tăng số lượng request, nhất là khi có nhiều cuộc gọi đến và đi. Do đó, RESTful API được thiết kế để lưu trữ dữ liệu vào cache để tăng tính tái sử dụng.

Cụ thể, ràng buộc này yêu cầu mỗi phản hồi phải đánh dấu dữ liệu bên trong nó là lưu được hoặc không lưu được vào cache. Nếu phản hồi lưu được vào cache, thì máy khách có thể sử dụng lại dữ liệu của phản hồi đó cho các yêu cầu tương tự sau này.

**4. Uniform interface (Giao diện thống nhất)**

REST áp dụng những nguyên tắc chung của kỹ thuật phần mềm cho giao diện thành phần. Chính vì lý do đó, tổng thể kiến trúc hệ thống được đơn giản hóa. Khả năng hiển thị của các tương tác cũng được cải thiện đáng kể.

Để có giao diện thống nhất, REST cần nhiều ràng buộc kiến trúc cho các thành phần bên trong. Chẳng hạn, việc biểu diễn tài nguyên trên toàn hệ thống phải tuân theo các nguyên tắc cụ thể. Các nguyên tắc đó bao gồm: quy ước đặt tên, định dạng liên kết, định dạng dữ liệu (XML hoặc JSON).

Ngoài ra, tất cả tài nguyên phải được truy cập thông qua một cách chung như HTTP GET. Việc điều chỉnh, sửa đổi tài nguyên cũng phải sử dụng một cách tiếp cận nhất quán.

**5. Layered system (Hệ thống phân lớp)**

Kiểu hệ thống phân lớp cho phép một kiến trúc chứa nhiều lớp phân cấp. Mỗi lớp sẽ có một chức năng và trách nhiệm cụ thể. Cách thức của REST là hạn chế hành vi của các thành phần trong một lớp. Mỗi thành phần hoàn toàn không thể thấy được những gì ở bên ngoài lớp mà chúng đang tương tác.

**6. Code on demand (Mã theo yêu cầu)**

Ràng buộc này là tùy chọn, không bắt buộc. Do đó, đây cũng là ràng buộc ít được biết đến của REST.

Ràng buộc này cho phép người dùng mở rộng chức năng của máy khách. Họ có thể tải xuống và thực thi mã dưới dạng các ứng dụng nhỏ (applet) hoặc tập lệnh. Điều này sẽ đơn giản hóa công việc cho máy khách, bằng cách giảm các tính năng bắt buộc triển khai từ trước.
