# Phân biệt giữa Authentication và Authorization

![image](https://res.cloudinary.com/practicaldev/image/fetch/s--VYXihGsl--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/ras8no1uj4ih1ogzy89h.png)

Trước khi tìm hiểu chi tiết, ta sẽ nói ngắn gọn về 2 thuật ngữ này.

Authentication (xác thực) là quá trình kiểm tra danh tính hay nói đơn giản hơn nó là quá trình xác minh bạn là ai. Nó đi tìm câu trả lời cho câu hỏi "Bạn là ai ?"

Authorization (uỷ quyền) là quá trình cấp quyền truy cập vào hệ thống hay nói đơn giản hơn nó là quá trình xác minh xem bạn có những quyền hạn nào. Nó đi tìm câu trả lời cho câu hỏi "Bạn được phép làm gì ?"

Ví dụ khi đi làm, để vào công ty bạn cần có thẻ nhân viên thì mới được vào làm - đó là Authentication. Sau khi vượt qua bước xác thực, khi vào công ty có rất nhiều phòng ban, tuỳ thuộc chức vụ của bạn là gì thì mới được phép vào các phòng tương ứng - đó là Authorization.

## Authentication (Xác thực)

Authentication là về việc xác thực thông tin đăng nhập của bạn như tên người dùng hoặc định danh người dùng và mật khẩu để xác minh danh tính của bạn. Hệ thống xác thực danh tính người dùng thông qua mật khẩu đăng nhập. Authentication thường được thực hiện bởi tên người dùng và mật khẩu, đôi khi kết hợp với các yếu tố xác thực, trong đó đề cập đến các cách khác nhau để được xác thực.

Các nhân tố xác thực (Authentication Factor) mà hệ thống sử dụng để xác minh một danh tính trước khi cấp cho họ quyền truy cập vào thực hiện bất cứ điều gì trên ứng dụng của bạn.

Dựa trên cấp độ bảo mật, Authentication Factor có thể thay đổi theo một trong các cách sau:

- Single-Factor Authentication (Xác thực một lớp): Nó là phương thức xác thực đơn giản nhất thường dựa vào mật khẩu đơn giản để cấp cho người dùng quyền truy cập vào một hệ thống cụ thể là một website hoặc network. Người này có thể yêu cầu quyền truy cập vào hệ thống chỉ bằng một trong các thông tin đăng nhập để xác minh danh tính của mình. Ví dụ phổ biến nhất về xác thực một yếu tố sẽ là thông tin đăng nhập chỉ yêu cầu mật khẩu đối với tên người dùng hoặc địa chỉ email.
![image](https://doubleoctopus.com/wp-content/uploads/2021/08/One-factor-authentication-1.png)

- Two-Factor Authentication (Xác thực hai lớp): Như tên của nó, nó có một quy trình xác minh gồm hai bước, không chỉ yêu cầu tên người dùng và mật khẩu, mà còn một thứ mà chỉ người dùng biết, để đảm bảo mức độ bảo mật bổ sung, chẳng hạn như pin ATM, chỉ người dùng mới biết. Sử dụng tên người dùng và mật khẩu cùng với một thông tin bí mật bổ sung khiến cho những kẻ lừa đảo hầu như không thể đánh cắp dữ liệu có giá trị.
![image](https://doubleoctopus.com/wp-content/uploads/2021/08/2FA-1-1-1-1.png)


- Multi-Factor Authentication (Xác thực nhiều lớp): Nó có một phương thức xác thực tiên tiến nhất sử dụng hai hoặc nhiều mức bảo mật từ các loại xác thực độc lập để cấp quyền truy cập cho người dùng vào hệ thống. Tất cả các yếu tố phải độc lập với nhau để loại bỏ bất kỳ lỗ hổng nào trong hệ thống.
![image](https://doubleoctopus.com/wp-content/uploads/2021/08/Multi-factor-authentication-1-1.png)

Authentication có thể kết hợp một hoặc nhiều các yếu tố sau: 

+ Password hay Pin: Là mật khẩu hay mã pin. Mật khẩu là một phương pháp xác thực vô cùng đơn giản, dễ triển khai nên đang được sử dụng rất rộng rãi và phổ biến. 
Mỗi khi người dùng thực hiện truy cập vào thì mỗi hệ thống đều sẽ lưu lại mật khẩu dưới dạng mã hóa một chiều (các loại mã hóa có thể là md5, sha1, hoặc tự viết cơ chế mã hóa).
Đây là tính năng sẽ đảm bảo cho mật khẩu dù bị hack nhưng cũng không thể khôi phục được thành chuỗi gốc.  Đây là phương pháp có nhiều biến thể khác nhau như: thiết kế dưới dạng biến thể Swipe Pattern PIN hoặc mật khẩu chỉ bằng cách dùng trong một lần (nó chuyên sử dụng cho những chức năng quan trọng).

+ Security token: Các mã bảo mật. Đây là phương pháp dựa vào thuật toán mã hóa khóa công cộng và khóa cá nhân để xác thực.
Muốn đăng nhập vào hệ thống, thì bạn chỉ cần có khóa cá nhân ở trên máy rồi thực hiện đăng nhập vào hệ thống mà không cần phải nhớ đến những thông tin đăng nhập như việc sử dụng mật khẩu. Thường thì các hệ thống quản trị server sẽ thường xuyên áp dụng biện pháp này.

+ Biometric verification: Sẽ xác minh sinh trắc học, sử dụng tròng mắt, dấu vân tay hoặc khuôn mặt là một trong những phương pháp xác thực dựa trên các yếu tố đặc trưng của một người. Phương pháp này mang lại ưu điểm là Id và mật khẩu sẽ luôn được đi cùng nhau nên bạn hoàn toàn không cần phải lo lắng bị quên hay lạc mất nó.
Mỗi khi bạn muốn đăng nhập lại thì chỉ cần chủ động sử dụng các yếu tố xác thực này một cách dễ dàng, mà không gặp bất kỳ khó khăn nào. 

Mặc dù có nhiều phương pháp để có thể xác thực cho một tài khoản, tuy nhiên thì bạn sẽ không thể tránh khỏi những rủi ro như: mất đi mật khẩu, vân tay bị đánh cắp, mất mã khóa cá nhân

## Authorization (Uỷ quyền)

Authorization là quá trình để xác định xem người dùng được xác thực có quyền truy cập vào các tài nguyên cụ thể hay không. Nó xác minh quyền của bạn để cấp cho bạn quyền truy cập vào các tài nguyên như thông tin, cơ sở dữ liệu, file… Authorization thường được đưa ra sau khi xác thực xác nhận các đặc quyền của bạn để thực hiện. Nói một cách đơn giản hơn, nó giống như cho phép ai đó chính thức làm điều gì đó hoặc bất cứ điều gì.

Authorization xảy ra sau khi hệ thống của bạn được authentication (xác thực) thành công, cuối cùng cho phép bạn toàn quyền truy cập các tài nguyên như thông tin, file, cơ sở dữ liệu, quỹ, địa điểm, hầu hết mọi thứ. Nói một cách đơn giản, authorization xác định khả năng của bạn để truy cập hệ thống và ở mức độ nào. Khi danh tính của bạn được hệ thống xác minh sau khi xác thực thành công, bạn sẽ được phép truy cập tài nguyên của hệ thống.

Ví dụ quy trình xác minh và xác nhận mã nhân viên và mật khẩu trong một tổ chức được gọi là Authentication, nhưng xác định nhân viên nào có quyền truy cập vào tầng nào được gọi là Authorization.

Các hình thức phân quyền thường gặp là:

+ Role-based authorization: Phân quyền dựa trên vai trò của người dùng. Ví dụ trong WordPress có các role như là  Subscriber, Contributor, Author, Editor, Administrator và mỗi một role sẽ có những quyền khác nhau và mỗi người dùng sẽ được phân role có quyền tương ứng. Đối với những hệ thống có nhiều người dùng thì role-based là cách tiếp cận tốt nhất để tiết kiệm thời gian trong việc phân quyền.

+ Object-based authorization: Phân quyền theo đối tượng. Cách này sẽ phân quyền cho từng đối tượng cụ thể. Ví dụ những đối tượng trong nhóm A, B được phân quyền chỉnh sửa các bài viết trong danh mục. Nhưng đối tượng trong nhóm A chỉ chỉnh sửa được bài viết trong danh mục C, đối trượng trong nhóm B chỉ chỉnh sửa bài viết trong danh mục D.

## Quá trình Authentication và Authorization diễn ra như nào?

![image](https://blog.ntechdevelopers.com/wp-content/uploads/2021/12/security.png)

Bước 1: Khi người dùng muốn truy cập vào một tài nguyên nào đó, bước đầu tiên họ sẽ phải gửi một request truy cập đến nguồn chứa tài nguyên đó. Để xác định người dùng đó là ai trong hệ thống thì nguồn chứa tài nguyên phải sinh ra một dấu hiệu trước đó. Đây chính là bước đăng ký mà các bạn vẫn hay làm. Hoặc có thể là do phía Admin cấp cho bạn tài khoản định danh. Dấu hiệu này sẽ có thể là mã nhân viên, tên định danh, tên đăng nhập…

Bước 2: Người dùng sẽ vào ứng dụng của bạn và điền các thông tin của dấu hiệu kể trên và yêu cầu quyền truy cập vào ứng dụng. Đây chính là quá trình đăng nhập của bạn.

Bước 3: Nguồn tài nguyên sẽ kiểm tra yêu cầu với thông tin truy cập trên và xác định được người đó là ai, họ có thể là khách, là người mua hàng, người quản lý, hay cộng tác viên. Khi này tài nguyên cần muốn biết phân loại được người truy cập vừa rồi có được quyền sử dụng, chỉnh sửa hay chỉ xem tài nguyên thì hệ thống sẽ tiếp tục xác thực quyền.

Bước 4, 5: Có nhiều cách xác thực quyền truy cập dựa trên các yêu cầu sau khi xác thực định danh. Đến đây khi hệ thống biết bạn có những quyền gì nó sẽ chỉ cấp cho bạn đúng quyền hạn đó mà thôi.

Bước 6,7,8: Khi nhận được quyền hạn truy cập các request tiếp theo bạn có thể truy cập được tài nguyên với quyền hạn đi kèm đã cấp.

