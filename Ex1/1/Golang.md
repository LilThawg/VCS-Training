# GOLANG

## Giới thiệu Golang

Go, thường được gọi là Golang, là một ngôn ngữ lập trình mã nguồn mở của Google và được phát hành ổn định đầu tiên vào năm 2011. Robert Griesemer, Rob Pike và Ken Thompson lần đầu tiên khởi xướng thiết kế cho Go vào năm 2007 và nó có mã nguồn mở vào năm 2009 .

Golang giống như C++ hay Java, nó cũng là một ngôn ngữ dùng để lập trình. Tuy nhiên điểm khiến ngôn ngữ Go trở nên khác biệt nằm ở sự đơn giản của nó. Nó sở hữu những cú pháp khá tinh gọn. Điều này có thể khiến những người mới học thấy khó khăn. Dù vậy cú pháp của Go có độ tương đồng rất lớn với C++. Vì thế nếu bạn đã quen thuộc với C++ thì việc học Golang không còn là điều khó khăn.

Tương tự như vậy, Go cũng có rất nhiều điểm chung với Java. Nó cũng được dùng để xây dựng và phát triển các ứng dụng cross-platform. Điều này khiến nó dễ dàng hòa nhập vào cộng đồng lập trình dù được sinh sau đẻ muộn.

Nhìn chung, ngôn ngữ lập trình Go được đánh giá như một công cụ nhỏ gọn và sắc bén. Nó giúp vận hành chương trình một cách nhanh chóng tại nhiều điều kiện khác nhau. Golang sở hữu độ tương thích rất cao, giống như cách mà Google phủ sóng thanh tìm kiếm của thế giới vậy.

Việc khai thác tối đa sức mạnh của các bộ xử lý đa lõi và phần cứng thế hệ mới đối với các ngôn ngữ hiện có được xem như là việc không thể làm được. Bởi những giới hạn vốn có của các ngôn ngữ lập trình trên máy tính như C, C++, Java,... Bấy lâu nay, các vấn đề xử lý đa lõi vẫn là chuyện của hệ điều hành.

Google đưa ra ngôn ngữ Go như là một cách tiếp cận khác về vấn đề xử lý đa lõi. Thay vì chỉ có hệ điều hành được phép cung cấp tài nguyên và xử lý, thì các phần mềm cũng có thể tương tác trực tiếp với nền tảng đa lõi giúp cho việc xử lý nhanh hơn rất nhiều.

### Cú pháp 

Về mặt cú pháp thì Go rất giống ngôn ngữ C, tuy nhiên nó có nhiều thay đổi trong thiết kế để an toàn hơn và có cú pháp súc tích và dễ đọc. Go cho phép lập trình viên vừa khai báo và khởi tạo biến cùng một lúc mà không cần phải chỉ định kiểu dữ liệu `i:=3` hoặc `name:="Hello, world!"`, điều này trái ngược với cú pháp của ngôn ngữ C `int i = 3;` và `const char *s = "Hello, world!"`.Ở Cuối mỗi dòng lệnh cũng không cần kết thúc bằng dấu chấm phẩy và mỗi hàm có thể trả về nhiều hơn một giá trị.

Ví dụ : 

In ra "Hello World!"

```go
 package main
 
 import "fmt"
 
 func main() {
   fmt.Println("Hello World!")
 }
```

Ví dụ về trả về nhiều hơn một giá trị

```go
package main

import "fmt"

// Phần `(int, int)` trong chữ ký hàm thể hiện rằng
// hàm này trả về 2 giá trị kiểu int
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Ở đây chúng ta sử dụng hai biến a và b để đón dữ liệu trả về
    // từ hàm vals()
    a, b:= vals()
    fmt.Println(a)
    fmt.Println(b)

    // Ta cũng có thể chỉ nhận về một tập con của giá trị trả về
    // bằng cách sử dụng ký hiệu  `_`.
    _, c:= vals()
    fmt.Println(c)
}
```

Ví dụ về xử lý song song (concurrency)

```go
package main

import (
	"fmt"
)

var (
	naturalChan = make(chan int)
	squaredChan = make(chan int)
	items       = make([]map[int]int, 10)
)

func natural() {
	for i:= range items {
		naturalChan <- i
	}
	close(naturalChan)
}

func square() {
	for _ = range items {
		x:= <-naturalChan
		squaredChan <- x * x
	}
	close(squaredChan)
}

func main() {
	go natural()
	go square()

	for _ = range items {
		select {
		case squared:= <-squaredChan:
			fmt.Printf("Squared %d\n", squared)
		}
	}
}
```

## Ưu điểm của Golang
* Dễ học và sử dụng 

Mặc dù Go có thể không phổ biến như JavaScript hoặc Python,C,C++,... nhưng nó có những điểm tương đồng với những ngôn ngữ này, nên việc học nó khi đã biết một vài ngôn ngữ là rất nhanh.

* Tốc độ làm việc


Run nhị phân thường chậm hơn so với C  tuy nhiên sự khác biệt này được đánh giá là không quá đáng kể với các ứng dụng. Đa số, hiệu suất làm việc của Go tốt ngang trong những phần công việc lớn và nhanh hơn so với các ngôn ngữ lập trình có độ nổi tiếng cao như Java, Python hay Ruby về tốc độ. 

* Khả năng tương thích cao 

Go có thể cung cấp những lợi ích trên cho lập trình viên mà không lo bị mất quyền truy cập vào các hệ thống dưới. Ngoài tra, phương mềm của Go có khả năng liên kết với thư viện C ở bên ngoài hoặc thực hiện các hệ thống native.

* Tính tiện lợi cao 

Không phải tự nhiên là Golang lại được so sánh với nhiều loại ngôn ngữ lập trình đến thế. Nó thậm chí còn không thua kém những loại ngôn ngữ kịch bản như Python nhờ khả năng đáp ứng vô vàn những nhu cầu lập trình thường thấy.


Ngôn ngữ Golang sở hữu một tính năng nổi bật và độc quyền mang tên goroutines. Goroutines tồn tại như một công cụ tích cực giúp giải quyết rất nhiều vấn đề còn tồn tại. Nó có thời gian khởi động nhanh hơn threads thông thường. Công cụ này cũng sở hữu đa kênh và có khả năng cho phép sự giao tiếp giữa các kênh này. Ngoài ra, goroutines còn có mutex locking, một tính năng cho phép khóa lại các cấu trúc dữ liệu để việc đọc và ghi nhớ không xảy ra xung đột

* Đa hỗ trợ

Toolchain Go đã có sẵn dưới dạng binary của Linux, MacOS hoặc Windows được xem như là một container trong Docker. Go đã được đặt mặc định với ở trong các bản phát hành phổ biến của Linux, giúp cho quá trình triển khai Go source sẽ dễ dàng hơn đối với những nền tảng trên. Ngoài ra, Go còn hỗ trợ cho nhiều development environment của các bên thứ ba.

* Linh hoạt: 

Hầu hết, các file executable đều được tạo ra bằng toolchain của Go nên có thể hoạt động độc lập hơn mà không cần phải dùng hệ điều hành. Ngoài ra, nó còn được sử dụng để thực hiện biên dịch mọi chương trình nhị phân qua nhiều nền tảng. 

## Nhược điểm của Golang

* Đôi lúc nó quá đơn giản

Một trong những ưu điểm chính của Go, cũng là một trong những điểm yếu lớn nhất của nó. Go có thể là một ngôn ngữ dễ dàng để chọn, nhưng điều đó mang lại sự thiếu linh hoạt.

* Go vẫn còn là một ngôn ngữ trẻ

Có ngôn ngữ rất nhiều tiềm năng hứa hẹn, tuy nhiên nó vẫn là một ngôn ngữ trẻ và đương nhiên sẽ gặp khó khăn khi phải cạnh trạnh với những ngôn ngữ đã được ra đời trước nó. Thư viện chuẩn của Go được thiết kế thông minh và hiệu quả, nhưng nó phải cạnh tranh với các ngôn ngữ như Java được hỗ trợ bởi một bộ mã khổng lồ và có một cộng đồng hùng hậu. Mặc dù Go cuối cùng có thể bắt kịp các ngôn ngữ khác và Go cũng có những điểm mạnh riêng của nó nhưng Go vẫn còn một chặng đường dài trong việc xây dựng thư viện hỗ trợ.

* Không hỗ trợ công cụ Generics

Golang không có khả năng hỗ trợ thuốc generic bởi vì nó cho phép bạn tạo một mã rất rõ ràng. Generics là tốt ở một mức độ nào đó, tuy nhiên việc sử dụng sai cũng thường xuyên xảy ra nên bạn cần phải lưu ý. 

* Không hỗ trợ thừa kế
* Không hỗ trợ quá tải toán tử hoặc ghi đè phương thức
* Không hỗ trợ thao tác trên con trỏ (vì lý do bảo mật)

## Concurrency

Concurrency (tính đồng thời) là khả năng xử lí nhiều tác vụ cùng 1 lúc. Ví dụ, khi bạn đang lướt web, có thể bạn đang download file trong khi đang nghe nhạc, đồng thời đang scroll trang. Nếu trình duyệt không thể thực hiện chúng cùng 1 lúc, bạn sẽ phải đợi đến khi mọi file download xong, mới có thể nghe nhạc, rồi mới có thể scroll. Điều này nghe rất khó chịu phải không nào?

Vậy làm thế nào để 1 máy tính có CPU 1 nhân có thể xử lí nhiều tác vụ cùng lúc, trong khi 1 nhân CPU chỉ có thể xử lí 1 việc tại 1 thời điểm? Concurrency tức là xử lí 1 tác vụ tại 1 thời điểm, nhưng CPU không xử lí hết 1 tác vụ rồi mới đến tác vụ khác, mà sẽ dành 1 lúc cho tác vụ này, 1 lúc cho tác vụ kia. Do vậy, chúng ta có cảm giác máy tính thực hiện nhiều tác vụ cùng 1 lúc, nhưng thực chất chỉ có 1 tác vụ được xử lí tại 1 thời điểm.

Cùng xem biểu đồ dưới để rõ hơn về cách CPU phân bố khi chúng ta sử dựng web ở ví dụ trên.

![image](https://user-images.githubusercontent.com/72289126/146537842-89621576-91a9-4f84-a286-1d72a14a4bea.png)

Từ biểu đồ trên, chúng ta có thể thấy rằng, CPU 1 nhân phân chia thời gian làm việc dựa trên độ ưu tiên của cùng tác vụ.

Ví dụ, khi đang scroll trang, việc nghe nhạc sẽ có độ ưu tiên thấp hơn, nên có thể nhạc của bạn sẽ bị dừng do đường truyền kém, nhưng bạn vẫn có thể kéo trang lên xuống.

## Parallelism

Ở phần trên, chúng ta chỉ nói về CPU 1 nhân, nhưng nếu máy tính có CPU nhiều nhân thì sao? 

Với ví dụ lượt web, CPU 1 nhân cần chia thời gian sử dụng cho các tác vụ. Còn với CPU nhiều nhân, chúng ta có thể xử lí nhiều tác vụ cùng lúc trên các nhân khác nhau.

![image](https://user-images.githubusercontent.com/72289126/146538090-fe8be525-060a-4251-b9bd-8c31c8060bcb.png)

Cơ chế xử lí nhiều tác vụ song song với nhau được gọi là parallelism. Khi CPU có nhiều nhân, chúng ta có thể sử dụng nhiều nhân để xử lí nhiều thứ 1 lúc.

## Concurrency vs Parallelism

Go khuyến khích việc sử dụng goroutines chỉ trên 1 nhân, tuy nhiên, chúng ta cũng có thể chỉnh sửa để goroutines chạy trên các nhân khác nhau. Có thể coi goroutines là 1 function trong go.

Có vài điểm khác biệt giữa concurrency và parallelism. Concurrency xử lí nhiều 1 tác vụ 1 lúc, còn parallelism là thực hiện nhiều tác vụ cùng 1 lúc. Parrallism không phải lúc nào cũng tốt hơn concurrency. Chúng ta sẽ nghiên cứu sâu hơn về vấn đề này trong những bài viết tới.

Để hiểu rõ hơn về concurrency trong go và cách sử dụng chúng trong code, trước hết chúng ta cần hiểu hơn về 1 tiến trình xử lí (process) của máy tính.

## Gorountine

Goroutines là các hàm hoặc phương thức chạy đồng thời với các hàm/ phương thức khác. Goroutines có thể được coi là những luồng gọn nhẹ. Chi phí tạo một Goroutine tương đối thấp so với một luồng. Do vậy, những ứng dụng Go có hàng ngàn Goroutines chạy đồng thời là điều hết sức bình thường.

Đây là một chương trình bình thường với hàm `hello` in ra tin nhắn `Hello world goroutine` và sau đó trong hàm main in ra tiếp tin nhắn  `main function`

```go
package main

import (  
    "fmt"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    hello()
    fmt.Println("main function")
}
```

Output : 

```
Hello world goroutine
main function
```

Thế nhưng khi khi ta đặt hàm `hello` làm 1 goroutine (Nhập từ khóa go vào trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời) 

```go
package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
}

```

thì ta lại chỉ có output : 

```
main function
```

Tại sao lại có chuyện này xảy ra ? 

Không giống hàm, hệ thống không chờ Goroutine chạy xong, mà trả lại ngay lập tức dòng code tiếp theo ngay sau khi Goroutine gọi và bất kỳ giá trị trả lại nào từ Goroutine sẽ bị bỏ qua.

Lưu ý là ở đây có 2 goroutine, đầu tiên là goroutine của hàm `main`, thứ hai là goroutine của hàm `hello`.

Goroutine chính cần chạy để các Goroutine khác chạy được. Nếu Goroutine chính dừng lại thì chương trình cũng dừng và không Goroutine nào chạy nữa.

Sau khi gọi go hello() tại dòng 12, hệ thống trả về ngay lập tức dòng tiếp theo của code mà không đợi việc in ra `Hello world goroutine` kết thúc và in ra main function. Sau đó Goroutine chính dừng khi không còn đoạn code nào để thực thi và do đó, hello Goroutine không còn cơ hội để chạy.

Chúng ta sửa bằng cách cho chương trình sleep 1s để chờ hàm `hello` chạy như sau:

```go
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(time.Second)
	fmt.Println("main function")
}
```

Phương pháp sử dụng sleep trong Gorountine khá là `...` :D nghĩ xem giờ 1 hàm ta cho nghỉ 1s vậy vòng for 10000 lần = 10000s à ? 

Thay vào đó chúng ta hãy sử dụng `sync`

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello world goroutine ", i)
		time.Sleep(time.Second)
	}
}
func goodbye() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Goodbye world goroutine ", i)
		time.Sleep(time.Second)
	}
}
func main() {
	//B1: tạo 1 biến kiểu sync.WaitGroup: wg
	var wg = sync.WaitGroup{} 
	//B2: wg.Add(số lượng goroutine phải đợi)
	wg.Add(2)
	//B3: ở mỗi goroutine gọi method wg.Done() trước khi return
	go func() {
		hello()
		wg.Done()
	}()

	go func() {
		goodbye()
		wg.Done()
	}()
	//B4: gọi method wg.Wait()
	wg.Wait()
	fmt.Println("Done")
}
```

(time.Sleep ở đây chỉ để chỉ rõ cho mọi người thấy rằng 2 gorountine đang chạy đồng thời)

Output : 

```
Hello world goroutine  1
Goodbye world goroutine  1
Hello world goroutine  2
Goodbye world goroutine  2
Goodbye world goroutine  3
Hello world goroutine  3
Hello world goroutine  4
Goodbye world goroutine  4
Goodbye world goroutine  5
Hello world goroutine  5
Done
```

## Channel

Channels có thể được coi là các đường ống sử dụng mà Goroutines giao tiếp. Tương tự như cách nước chảy từ đầu này sang đầu kia trong đường ống, dữ liệu có thể được gửi từ một đầu và nhận từ đầu kia bằng channels.

Channel được khởi tạo với cú pháp như sau : 

```go
ch := make(chan int) //khởi tạo channel ch kiểu int
```

Gửi và nhận từ channel :

```go
data := <- ch // đọc dữ liệu từ kênh ch và gán vào biến data  
ch <- data // gửi data tới kênh ch
```

**Unbuffered channel**

Unbuffered channel không có khoảng trống để chứa dữ liệu, yêu cầu cả 2 goroutines gửi và nhận đều sẵn sàng cùng lúc. Khi 1 goroutine gửi dữ liệu vào channel, luồng xử lý sẽ bị block lại cho tới khi có 1 goroutine đọc từ channel này ra.

Khai báo unbuffered channel:

```go
ch := make(chan int)
```

![image](https://i.stack.imgur.com/Lq7ZG.png)

Ví dụ về Unbuffered channel : 

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(2)
	//Nhận dữ liệu
	go func() {
		data := <-ch
		fmt.Println(data)
		wg.Done()
	}()
	//Gửi dữ liệu
	go func() {
		data := "LilThawg"
		ch <- data
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Doneee")
}
```

**Buffrered channel**

Buffrered channel có từ 1 khoảng trống trở lên để chứa dữ liệu, không yêu cầu cả 2 goroutines gửi và nhận cùng phải sẵn sàng cùng lúc. Luồng xử lý chỉ bị block nếu: goroutines gửi bị block nếu không còn khoảng trống trong channel; goroutines nhận bị block nếu trong channel không có dữ liệu.

Khai báo buffered channel : 

```go
ch := make(chan int, 10)
```

![image](https://i.stack.imgur.com/kSx6w.png)

Ví dụ về Buffrered channel : 

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(2)
	//Nhận dữ liệu
	go func(ch <-chan string) {
		for i := range ch {
			fmt.Println(i)
			time.Sleep(time.Second)
		}

		wg.Done()
	}(ch)
	//Gửi dữ liệu
	go func(ch chan<- string) {
		for s := 'A'; s <= 'Z'; s++ {
			data := string(s)
			ch <- data
		}
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println("Doneee")
}
```

