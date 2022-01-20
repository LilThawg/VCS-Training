# Các phương thức Authentication

## Basic Authentication

Basic Auth là cơ chế xác thực đơn giản nhất của một ứng dụng web và cũng là phương thức xác thực ít được khuyến khích bởi tình bảo mật của nó không an toàn, việc ứng dụng nó rất dễ dàng và đã được nhiều phần mềm / máy chủ tích hợp và xử lý một cách tự động.

Với phương thức này có thể không cần yêu cầu về cookies, session ID... Client cần thêm header Authorization vào tất cả các request. Username và password không được encrypt mà được cấu trúc như sau :

```
Authorization: Basic base64encode(username:password)
```

Ví dụ với username = "lethang" và password = "S3cr3T" :

```
Authorization: Basic bGV0aGFuZzpTM2NyM1Q=
```

```uml
// https://app.gleek.io/
// Sample UML Sequence Diagram
// ** Check Right-Click Menu for more features! **

Client -Get / HTTP/1.1-> Server
Server -HTTP/1.1 401 Unauthorized-> Client
Client -Ask User-> Client
Client -Get / HTTP/1.1<br>Authorization: Basic bGV0aGFuZzpTM2NyM1Q=-> Server
Server -Check credentials-> Server
Server -HTTP/1.1 200 OK<br>or<br>HTTP/1.1 401 Unauthorized-> Client
```

![image](https://user-images.githubusercontent.com/72289126/148673451-e284892b-b471-4e6a-9001-0b65ce11e8df.png)

### Ưu điểm

* Đơn giản, do đó được hầu hết các trình duyệt, webserver (nginx, apache,...) hỗ trợ. Bạn có thể dễ dàng config cho webserver sử dụng Basic Auth với 1 vài dòng config.
* Dễ dàng kết hợp với các phương pháp khác. Do đã được xử lý mặc định trên trình duyệt và webserver thông qua truyền tải http header, các bạn có thể dễ dàng kết hợp phương pháp này với các phương pháp sử dụng cookie, session, token,...

### Nhược điểm

* username và password được gửi trong mọi request, nên có nguy cơ bị lộ tài khoản, kể cả khi sử dụng kết nối bảo mật.
* Liên kết với SSL/TLS, nếu mà website sử dụng mã hoá yếu thì sẽ bị lộ tài khoản ngay lập tức.
* Không có cách nào để đăng xuất tài khoản cả.
* Phiên đăng nhập sẽ không bao giờ bị hết hạn, trừ khi người dùng đổi password.

## Bearer Authentication

Phương thức này hay còn được gọi là token authentication có thể hiểu đơn giản là "cấp quyền truy cấp cho người mang (bearer) token này". Bearer token sẽ cho phép truy cập đến một số tài nguyên hoặc url nhất định và thường là một chuỗi string được mã hóa, được sinh ra bởi server trong response để thực hiện request login.

Khi thực hiện bằng phương thức này thì client phải gửi bearer token này trong header để thực hiện request

```
Authorization: Bearer <token>
```
Cách thức thực hiện với bearer authentication như sau:

Khi user gửi request đến server để lấy một token bằng username, password thông qua SSL, server sẽ trả về một chuỗi access token. 

Access token này chính là bearer token mà client cần phải gửi vào header nếu muốn thực hiện các request khác để server xác thực user đó là đúng. 

Token này có thể là một chuỗi mã hóa với các thuộc tính của user, vai trò của user đó. Khi server nhận được token này, sẽ giải mã sau đó sẽ thực hiện validate request đó trong ứng dụng xem user có được quyền thực hiện request đó hay không.

## API Keys

![image](https://blog.restcase.com/content/images/2019/07/nonref-docs-preso_apikey.png)

Phương thức này được tạo ra với mục đích khắc phục những vấn đề của basic authentication. Trong phương thức này, một giá trị key duy nhất sinh (unique key) ra và gán cho user trong lần đầu tiên, biểu thị rằng user đó được xác định. Khi user quay trở lại hệ thống, unique key (có thể được sinh ra từ việc kết hợp giữa thông tin phần cứng, địa chỉ IP và thời gian của server) được sử dụng để chứng minh rằng user là giống với lần đầu tiên.

API key có thể đặt trong query string như sau : 

```
GET /something?api_key=abcdef12345
```

![image](https://user-images.githubusercontent.com/72289126/148680436-c4c8e113-cdb1-44e2-ab78-2874bccc5516.png)

Hoặc đặt trong request header : 

```
GET /something HTTP/1.1
X-API-Key: abcdef12345
```

Hay dùng như 1 cookie : 

```
GET /something HTTP/1.1
Cookie: X-API-KEY=abcdef12345
```

API Key được coi là một bí mật mà chỉ máy khách và máy chủ biết. Giống như  Basic authentication, xác thực dựa trên khóa API chỉ được coi là an toàn nếu được sử dụng cùng với các cơ chế bảo mật khác như HTTPS / SSL.

## JWT

JWT hay Json Web Token là một loại token được chấp nhận và sử dụng rộng rãi như một tiêu chuẩn của các nền tảng web hiện đại (RFC 7519) bởi nó thỏa mãn được tính chất self-contained, được hỗ trợ bởi nhiều ngôn ngữ và nền tảng và hơn hết là cấu trúc JSON đơn giản, nhỏ gọn hơn rất nhiều so với các loại token khác như Simple Web Tokens (SWT) and Security Assertion Markup Language Tokens (SAML).

![image](https://topdev.vn/blog/wp-content/uploads/2017/12/jwt-la-gi.jpeg)

Cấu trúc của JSON Web Token gồm 3 phần : Header, Payload, Signature. Những thành phần này được ngăn cách với nhau rất rõ ràng bởi ký tự ".". Do đó, chúng ta có thể hình dung ra cấu trúc của nó sẽ theo format sau:

```
<base64-encoded header>.<base64-encoded payload>.<base64-encoded signature>
```


### Header

Phần header sẽ chứa kiểu dữ liệu , và thuật toán sử dụng để mã hóa ra chuỗi JWT

```json
{
    "typ": "JWT",
    "alg": "HS256"
}
```

* typ (type) : chỉ ra rằng đối tượng ở đây là một JWT
* alg (algorithm) : chỉ ra rằng thuật toán dùng để mã hóa cho chuỗi đang được sử dụng ở đây là HS256(Hmac + SHA256)

Và chúng ta có đầu ra của header là : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`

### Payload

Phần tiếp theo của Token chính là Payload. Nó đóng một vai trò rất quan trọng trong JWT, đây là nơi chứa các nội dung của thông tin (claim) mà người sử dụng muốn truyền đi ở bên trong chuỗi.Các thông tin này góp phần mô tả thực thể một cách đơn giản và nhanh chóng hoặc cũng có thể là các thông tin bổ sung thêm cho phần Header. 

Có 3 loại claims thường gặp trong Payload: reserved, public và private claims.

* Reserved: đây là những thông tin đã được quy định trong IANA JSON Web Token Claims registry. Tuy nhiên, những thông tin này không mang tính bắt buộc. Bạn có thể tùy vào từng ứng dụng khác nhau mà bạn có thể để đặt ra những điều kiện ràng buộc đối với những thông tin cần thiết nhất. Ví dụ như:

  * iss (issuer): tổ chức phát hành của Token
  * sub (subject): chủ đề Token
  * aud (audience): đối tượng sử dụng Token
  * exp (expired time): thời điểm token sẽ hết hạn
  * nbf (not before time): token chưa hợp lệ trước thời điểm này
  * iat (issued at): thời điểm token sẽ được phát hành, tính theo UNIX time
  * jti: ID của JWT

* Public: Được định nghĩa tùy theo ý muốn của người sử dụng JWT. Tuy nhiên để tránh tình trạng trùng lặp xảy ra thì nên được quy định ở trong IANA JSON Web Token Registry hoặc là 1 URL có chứa không gian tên không bị trùng lặp.

* Private: Đây là phần thông tin thêm được dùng để truyền tải qua lại giữa các máy khách với nhau. 

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239021
}
```

Đầu ra của payload sẽ là : `eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIxfQ`

### Signature

Phần chữ ký này sẽ được tạo ra bằng cách mã hóa phần header , payload kèm theo một secret key : 

```js
data = base64Encode( header ) + "." + base64Encode( payload )
signature = Hash( data, secret )
```

Mã hoá với secret key = "secret" ta được signature : `L6yL8agHqFnFChgwv2bn6vdEOGTRF-6TMfOwzk8JRyY`

Tổng kết lại ta được 1 chuỗi jwt hoàn chỉnh : 

```jwt
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIxfQ.L6yL8agHqFnFChgwv2bn6vdEOGTRF-6TMfOwzk8JRyY
```

![image](https://user-images.githubusercontent.com/72289126/148696798-8e83dab5-01cc-4a4f-b2f5-403c9a5e78a7.png)

### Khi nào nên dùng JSON Web Token?


Authentication: Đây là trường hợp phổ biến nhất thường sử dụng JWT. Khi người dùng đã đăng nhập vào hệ thống thì những request tiếp theo từ phía người dùng sẽ chứa thêm mã JWT. Điều này cho phép người dùng được cấp quyền truy cập vào các url, service, và resource mà mã Token đó cho phép. Phương pháp này không bị ảnh hưởng bởi Cross-Origin Resource Sharing (CORS) do nó không sử dụng cookie.

Quá trình xác thực khi sử dụng jwt : 

![image](https://res.cloudinary.com/dxxbnrmtx/image/upload/v1581521944/jwt_vlecak.jpg) 

Trao đổi thông tin: JSON Web Token là 1 cách thức khá hay để truyền thông tin an toàn giữa các thành viên với nhau, nhờ vào phần signature của nó. Phía người nhận có thể biết được người gửi là ai thông qua phần signature. Và chữ ký được tạo ra bằng việc kết hợp cả phần header, payload lại nên thông qua đó ta có thể xác nhận được chữ ký có bị giả mạo hay không.

## OAuth2

OAuth là viết tắt của Open với Authentication hoặc Authorization. OAuth là một tiêu chuẩn mở ủy quyền truy cập (Authorization) giúp cho người dùng cho phép các trang web hoặc các ứng dụng có quyền truy cập vào thông tin của họ mà không cần chia sẻ thông tin username và password.

OAuth2 là bản nâng cấp của OAuth1.0, là một giao thức chứng thực cho phép các ứng dụng chia sẻ một phần tài nguyên với nhau mà không cần xác thực qua username và password như cách truyền thống từ đó giúp hạn chế được những phiền toái khi phải nhập username, password ở quá nhiều nơi hoặc đăng ký quá nhiều tài khoản mật khẩu mà chúng ta chẳng thể nào nhớ hết.

Nếu bạn thấy một website (hay phần mềm) nào đó cho phép đăng nhập bằng tài khoản Facebook hay Google mà không cần tạo nick mới thì đó chính là OAuth.

### OAuth2 hoạt động như thế nào?

Khi bạn đăng nhập bằng Facebook hay Gmail, website sẽ dẫn bạn đến trang (hoặc phần mềm) Facebook và liệt kê những quyền mà nó cần phải có để cho phép bạn đăng nhập và sử dụng dịch vụ.

Nếu bạn đồng ý thì lúc này Facebook sẽ phát cho website một cái token. Token này chứa một số quyền hạn nhất định giúp cho website có thể xác minh bạn là ai cũng như giúp cho website có thể hoạt động được.

Nếu website này bị hacker tấn công thì nó chỉ lấy được thông tin hay hoạt động của bạn trên website đó mà không ảnh hưởng đến những website khác mà bạn đang sử dụng.

Do đó cách đăng nhập bằng phương thức OAuth này rất an toàn cho người dùng cuối như chúng ta.

### Các tác nhân (đối tượng) trong Oauth2

OAuth2 làm việc với 4 đối tượng mang những vai trò riêng:

* Resource Owner (User): Là những người dùng ủy quyền cho ứng dụng cho phép truy cập tài khoản của họ. Sau đó ứng dụng được phép truy cập vào những dữ liệu người dùng nhưng bị giới hạn bởi những phạm vi (scope) được cấp phép. (VD: chỉ đọc hay được quyền ghi dữ liệu) => chính là bạn.
* Client (Application): Là những ứng dụng mong muốn truy cập vào dữ liệu người dùng. Trước khi được phép tương tác với dữ liệu thì ứng dụng này phải qua bước ủy quyền của User, và phải được kiểm tra xác nhận thông qua API. => Có thể hiểu là các ứng dụng sử dụng Facebook, Twitter, Google API chẳng hạn.
* Resource Server (API): Nơi lưu trữ thông tin tài khoản của User và được bảo mật.
* Authorization Server (API): àm nhiệm vụ kiểm tra thông tin user (VD: ID), sau đó cấp quyền truy cập cho Application thông qua việc phát sinh "access token".

Resource Server và Authorization Server chính là điểm khác biệt cơ bản giữa OAuth2 và OAuth1 khi tách biệt được hai thao tác: chứng thực (Authorization) và cung cấp thông tin người dùng (Resource) thành 2 Server.

### Sơ đồ luồng hoạt động của OAuth2

Để có quyền truy cập vào các tài nguyên được bảo vệ, `OAuth2.0` đã sử dụng `access token`, nó là một chuỗi đại diện cho các quyền truy cập. `Access Token` được sinh ra cho việc Authorization ở định dạng `JWT`.

Quyền truy cập được đại diện bởi access token, trong OAuth 2.0 gọi là `scope`.

![image](https://images.viblo.asia/4b4ca6c8-c55d-4c45-9301-8b7f312fbee5.png)

![image](https://topdev.vn/blog/wp-content/uploads/2019/03/mo-hinh-oauth.png)

1. Ứng dụng (website hoặc mobile app) yêu cầu ủy quyền để truy cập vào Resource Server (Gmail,Facebook, Twitter hay Github…) thông qua User
2. Nếu User ủy quyền cho yêu cầu trên, Ứng dụng sẽ nhận được ủy quyền từ phía User (dưới dạng một token string)
3. Ứng dụng gửi thông tin định danh (ID) của mình kèm theo ủy quyền của User tới Authorization Server
4. Nếu thông tin định danh được xác thực và ủy quyền hợp lệ, Authorization Server sẽ trả về cho Ứng dụng access_token. Đến đây quá trình ủy quyền hoàn tất.
5. Để truy cập vào tài nguyên (resource) từ Resource Server và lấy thông tin, Ứng dụng sẽ phải đưa ra access_token để xác thực.
6. Nếu access_token hợp lệ, Resource Server sẽ trả về dữ liệu của tài nguyên đã được yêu cầu cho Ứng dụng.

## OpenID Connect

Trước khi đi vào OpenID Connect ta cần làm rõ sự khác nhau giữa OpenID và OAuth2.

https://qastack.vn/programming/1087031/whats-the-difference-between-openid-and-oauth

OpenID là về xác thực (ví dụ: chứng minh bạn là ai), OAuth là về ủy quyền (nghĩa là cấp quyền truy cập vào chức năng / dữ liệu / vv .. mà không phải xử lý xác thực ban đầu).

OpenID Connect (OIDC) Kết hợp các tính năng của OpenID và OAuth tức là cả Xác thực và Ủy quyền.

![image](https://i.stack.imgur.com/LbKkm.png)

OpenID Connect là một lớp nhận dạng được phát triển trên giao thức OAuth 2.0. Nó cho phép khách hàng xác nhận danh tính và nhận thông tin hồ sơ cơ bản về người dùng cuối dựa trên xác thực được thực hiện bằng cách sử dụng máy chủ ủy quyền. Các nhà phát triển ứng dụng và web sử dụng nó để xác thực người dùng mà không chịu trách nhiệm lưu trữ và quản lý mật khẩu. Nó thậm chí còn cung cấp chức năng Đăng nhập một lần (SSO).

OpenID Connect giúp các client kiểm tra user thông qua các authorization server, đồng thời với đó là lấy các thông tin cơ bản của user đó bằng cách gọi đến REST API. OpenID Connect dùng JSON Web Tokens (JWT) trong đó việc authentiation dựa vào các token string trao đổi qua lại giữa client và authentication server. Thông tin mà server gửi về bao gồm một Acess Token, và có thể có (nếu bạn yêu cầu) một ID token.

Lấy ví dụ khi ta đăng nhập dùng tài khoản Google:

1. Khi bạn chọn đăng nhập dùng tài khoản Google, Authen server gửi một Authorization Request tới Google để xin được cấp quyền truy cập.
2. Google xác thực thông tin của bạn hoặc yêu cầu bạn login nếu bạn chưa đăng nhập, và sau đó hỏi bạn việc cấp quyền, tức là các quyền truy cập như avatar, email… của bạn.
3. Khi bạn đã chấp thuận việc đăng nhập, Google sẽ gửi một Access Token, và có thể có ID Token cho Authen server.
4. Authen server có thể nhận thông tin người dùng từ ID Token hoặc sử dụng Access Token để truy cập tới Google API.

Trong đó:

Access Token là một mã bí mật dùng cho ứng dụng để có thể truy cập API. Access Token có thể là một string, JWT token, hay một token không phải là JWT. Nó báo cho API biết được là ứng dụng đã được cấp quyền cho một số action nhất định dựa trên scope mà ứng dụng đã được cấp.

ID Token là một JWT chứa các thông tin nhận dạng. Nó được ứng dụng sử dụng để lấy các thông tin của user như user name, email… và hiển thị lên giao diện. ID Token tuân theo chuẩn công nghiệp (IETF RFC 7519)  và gồm có 3 phần: header, body và chữ ký.

JWT Token chứa các claim là các thông tin về định danh (như là name, email address) của một thực thể (thông thường là user) và các thông tin metadata khác.

![image](https://19yw4b240vb03ws8qm25h366-wpengine.netdna-ssl.com/wp-content/uploads/What-is-OpenID-Connect-flow-graphic-2.jpg)