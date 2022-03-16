# Transaction là gì?

Transaction là một chủ đề được đề cập đến rất nhiều trong các ứng dụng doanh nghiệp, để đảm bảo tính toàn vẹn của dữ liệu ngay cả trong các hệ thống lớn phát sinh một lớn các thay đổi lên database và đồng thời. Transaction là một tập hợp các hoạt động đọc/ghi xuống database hoặc là chúng đều thực thi thành công hết hoặc không có hoạt động nào được thực thi xuống database.

Để dễ hình dung, mình xin lấy một ví dụ như sau: chúng ta đang làm việc với một ứng dụng bên ngân hàng, ứng dụng này có một tính năng đó là chuyển tiền giữa các tài khoản ngân hàng với nhau, ví dụ chuyển tiền cho tài khoản A sang tài khoản B. Đây là một tính năng mà quá trình xử lý sẽ gồm 2 bước:

- Thứ nhất là trừ tiền trong tài khoản A
- Và bước thứ hai là cộng tiền vào tài khoản B.

Rõ ràng, như các bạn thấy, đây là một quá trình nếu gọi là thành công thì bắt buộc 2 bước trên phải thành công, hoặc nếu thất bại thì bắt buộc 2 bước trên cũng phải thất bại luôn. Nếu chỉ một trong 2 bước trên thành công hoặc thất bại thì chúng ta sẽ gặp vấn đề ngay. Và chúng ta khi xây dựng ứng dụng cũng phải đảm bảo 2 điều kiện trên.

Đây chính là ý tưởng chính của khái niệm transaction trong database. Trong database, một transaction là một tập hợp các thao tác mà ở đó hoặc là các thao tác đó được thực hiện thành công, hoặc là không một thao tác nào được thực hiện thành công cả. Khi các bạn suy nghĩ về transaction trong database, các bạn hãy hình dung nó là “hoặc là tất cả hoặc là không gì hết”.

Có thể lấy thêm ví dụ về 1 Transaction đơn giản nhất là tiến trình cài đặt phần mềm hoặc gỡ bỏ phần mềm. Việc cài đặt được chia thành các bước, thực hiện tuần tự từ đầu đến cuối, nếu toàn bộ các bước thực thi thành công đồng nghĩa với việc tiến trình cài đặt hoặc gỡ bỏ phần mềm thành công và ngược lại, một phép thất bại thì tiến trình phải rollback lại tức sẽ không có bất kỳ thay đổi nào trên máy tính.

Ở đây, chúng ta có một chữ viết tắt ACID (Atomicity, Consistency, Isolation, Durability) định nghĩa các tính chất của một transaction trong database.

## Atomicity (tính đơn vị)

Tính chất này thể hiện khái niệm “hoặc là tất cả hoặc là không gì hết”. Toàn bộ các bước được thực hiện trong một transaction nếu thành công thì phải thành công tất cả, nếu thất bại thì tất cả cũng phải thất bại. Nếu một transaction thành công thì tất cả những thay đổi phải được lưu vào database. Nếu thất bại thì tất cả những thay đổi trong transaction đó phải được rollback về trạng thái ban đầu.

Trong ví dụ của mình, rõ ràng khi đã trừ tiền trong tài khoản A thì bắt buộc việc cộng tiền vào tài khoản B phải xảy ra. Còn không thì không có gì xảy ra hết, đã không trừ tiền trong tài khoản A thì chắc chắn rồi, việc cộng tiền vào tài khoản B cũng không xảy ra.

## Consistency (tính nhất quán)

Consistency nghĩa là tất cả các ràng buộc toàn vẹn dữ liệu(constraints, key, data types, Trigger, Check) phải được thực thi thành công cho mọi transaction phát sinh xuống database, nhầm đảm bảo tính đúng đắn của dữ liệu.

Trong ví dụ của mình thì, nếu đã trừ 10$ trong tài khoản A thì trong tài khoản B phải được cộng 10$.

## Isolation (tính độc lập)

Nếu chúng ta có nhiều transaction cùng một lúc thì phải đảm bảo là các transaction này xảy ra độc lập, không tác động lẫn nhau trên dữ liệu.

Trong ví dụ của mình, giả sử cùng môt lúc xảy ra 2 transaction như trên thì rõ ràng chúng ta không xác định được trạng thái của tài khoản A và tài khoản B sau khi thực hiện cùng 1 lúc. Mà điều này thì rất nguy hiểm.

## Durability (tính bền vững)

Dữ liệu của transaction sau khi thực thi xong được cố định, chính thức và bền vững. Nghĩa là những thay đổi đã được cố định, không có chuyện có thể chuyển lại trạng thái dữ liệu lúc trước khi thực hiện transaction.

Dữ liệu của transaction sau khi thực thi xong được cố định, chính thức và bền vững. Nghĩa là những thay đổi đã được cố định, không có chuyện có thể chuyển lại trạng thái dữ liệu lúc trước khi thực hiện transaction.
