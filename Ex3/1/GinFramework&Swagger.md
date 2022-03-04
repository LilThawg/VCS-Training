# Tìm hiểu về Gin web framework

Tài liệu tham khảo : https://github.com/gin-gonic/gin

Gin là một web framework được viết bằng Go (Golang). Nó có một API giống như martini với hiệu suất nhanh hơn tới 40 lần nhờ có httprouter.

Các tính năng hỗ trợ của GIN :

- Hỗ trợ các restful methods như GET,POST,PUT,DELETE,OPTION, etc.
- Lấy các tham số trên path (ví dụ : /user/john/)
- Lấy các tham số Querystring (cũng trên path) (ví dụ : ?name=john)
- Upload file(Single file or Multiple files)
- Nhóm các routes lại
- Hỗ trợ Middleware (so với Mux thì Gin làm dễ hơn rất nhiều)
- Ghi log file
- Model binding và validation (ràng buộc các thuộc tính theo mỗi model và xác nhận lại)
- Render ProtoBuf XML, JSON, YAML
- Serving data static files, from file, from reader
- Multitemplate
- Redirects(chuyển hướng)
- Custom Middleware
- Goroutines trong middleware
- Run multiple service
- Run multiple service
- ...


# Hướng dẫn sử dụng Gin-Swagger (step by step) (Bài này em lấy ví dụ ở challenge 2)

Hiểu đơn giản swagger nó là tài liệu hướng dẫn sử dụng API. Do vậy bạn cứ làm xong hết việc rồi mới viết hướng dẫn nhé.

Tài liệu : https://github.com/swaggo/gin-swagger

## Khởi tạo swagger

1. `go get -u github.com/swaggo/swag/cmd/swag`
2. `swag init` 

Lúc này nó sẽ gen ra folder `docs`.

Sau khi `swag init` để sinh ra Swagger 2.0 docs ta import 2 thêm các packages sau (Nhấn `quick fix` nó tự get mấy package này về nhé) : 

```
"github.com/swaggo/gin-swagger" // gin-swagger middleware
"github.com/swaggo/files" // swagger embed files
```

Và thêm 1 dòng code : 

```
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

![image](https://user-images.githubusercontent.com/72289126/156089422-923c0bf8-2100-4a8a-907b-e72531747e71.png)

Sau đó ta truy cập : http://localhost:8080/swagger/index.html thì nhận được lỗi 

![image](https://user-images.githubusercontent.com/72289126/156089558-7d07c972-70df-404d-a4ee-f8a5d49ee07a.png)


OK bây giờ ta thêm comment ở trước hàm main như sau (phần không có @ là comment bình thường, phần có @ là phần swagger sẽ đọc và lưu vào file docs, đây có thể hiểu là phần khởi tạo): 

```
// Khai báo Swagger

// @title           Swagger Example API
// @version         1.0
// @description     Đây là swagger api của challenge 2
// @host localhost:8080
```

Sau khi đặt comment xong, ta `swag init` để lưu thông tin lại trong file docs, sau đó thử run lại thì vẫn cái lỗi `Fetch errorInternal Server Error doc.json` đó. Sau một hồi tìm hiểu thì ta phải import thêm file docs vào nữa =="

![image](https://user-images.githubusercontent.com/72289126/156093498-5f640448-6313-4af5-8ebc-6eb50f9ebae7.png)

![image](https://user-images.githubusercontent.com/72289126/156093636-f69ac238-6c2c-4762-9e48-15fd7157d5c4.png)

Ok vậy là hoàn thành xong bước đầu giờ ta đi viết hướng dẫn cho các API

## Viết hướng dẫn sử dụng API

Đọc trước : https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html

### Handler Auth

Bắt đầu ví dụ với Handler Auth nhé. Ta thêm cmt vào trước func SignUp :

```
// @Summary 	Sign Up
// @Description Đăng ký
// @Tags 		Handler Auth
// @Accept		json
// @Produce     json
// @Param 		user body models.UserSignUp true "User Sign Up" 
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /SignUp [post]
```

![image](https://user-images.githubusercontent.com/72289126/156129495-da94a259-bb3e-4d44-adae-98a08eecd2d1.png)

Giải thích sơ qua 1 tý :
- Param : tên param = user, loại param : body (có nghĩa là post object ở trong body đó, còn 1 số loại param như path, query string), có bắt buộc không ? True, "User Sign Up" là comment
- Success và Failure trả về object, với kiểu models.ExampleSuccess hoặc models.ExampleFailure (cái này mình khai báo thêm nha)
![image](https://user-images.githubusercontent.com/72289126/156136148-0aade64d-1979-47df-bed6-9dc63a335593.png)

- Router thì bắt buộc phải có rồi, nếu không nó không hiển thị đâu.

`swag init` và vào xem thành quả thôi 

![image](https://user-images.githubusercontent.com/72289126/156139197-20ba36c1-d8b3-432a-9e69-2194a1e13668.png)

![image](https://user-images.githubusercontent.com/72289126/156139559-a5e4b3cc-4d9c-4ab8-8aee-2402c62194f1.png)

Nhấn `Try it out` và nhập thông tin vào và nhấn `Excute`

![image](https://user-images.githubusercontent.com/72289126/156140352-a5d42f59-e9e1-4cec-8879-1d0040e4d8b4.png)

![image](https://user-images.githubusercontent.com/72289126/156140865-0ac918b7-dfab-441b-a253-c883918fd9a7.png)

Thử trùng email xem sao : 

![image](https://user-images.githubusercontent.com/72289126/156141128-1ca5ddf7-865d-4a59-a6b5-ddd74856e285.png)


Ok vậy là thành công rồi đó, mấy cái sau làm tương tự thôi.

### Handler User

Phần này có jwt token, sau 1 hồi search thì hiện tại em sử dụng : 

```go
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
```
Thêm vào trước func AddUser

```go
// AddUser godoc
// @Summary 	Add a user
// @Description	Thêm 1 User
// @Tags		Handler User (Only Admin)
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		user body models.UserAdd true "Add User"
// @Success		200 {object} models.User
// @Failure     400 {object} models.ExampleFailure
// @Router      /admin/addUser [post]
```

Vì trả về model.User (có gorm) nên ta dùng `swag init --parseDependency --parseInternal` mới update được, run và xem thành quả thôi.

![image](https://user-images.githubusercontent.com/72289126/156278789-0cb4685e-bb20-43e1-83d5-c53cefa11744.png)

Tạo user mới với tài khoản admin

![image](https://user-images.githubusercontent.com/72289126/156279132-ec9dbaa5-c2e5-495f-bfcf-1c514b280274.png)

Click phát nữa (tất nhiên là trùng email rồi)

![image](https://user-images.githubusercontent.com/72289126/156279251-3b93bc3f-e38b-493d-b862-019ae4162bfe.png)

Thử dùng token của user

![image](https://user-images.githubusercontent.com/72289126/156279435-08c1fbac-f003-41c5-8a8f-7e3dec092683.png)

### Validate

import "github.com/go-playground/validator"

Thêm vào struct User 

```go
type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,checkstring"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,checkstring"`
	Role     string `json:"role"`
}
```

![image](https://user-images.githubusercontent.com/72289126/156591268-84f9bfa5-23ad-41b5-ab16-d3b73c5e7bb5.png)

![image](https://user-images.githubusercontent.com/72289126/156591870-23e482e6-732e-43f8-81e5-1acd8ebec1b5.png)

Chạy thử thôi

![image](https://user-images.githubusercontent.com/72289126/156592114-3197a9e4-b49b-4d0d-8f30-b08ca4be2719.png)

![image](https://user-images.githubusercontent.com/72289126/156592289-8fe00c90-df72-4e3c-b03e-63cf9272f033.png)
