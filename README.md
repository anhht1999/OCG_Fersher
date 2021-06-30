# OCG_Fersher

# BACK-END

PORT: 3000

CSDL ở thư mục sql
ERD
![image](https://user-images.githubusercontent.com/48313937/123947494-a6345200-d9ca-11eb-8024-81dbcd431614.png)
Chạy file app.go (lệnh go run app.go)
#### Các API hoàn thành 
<br>//API Product 
func routeProduct(router *fiber.Router) {
 <br>   //Liệt kê Prodcuts, phân trang ,tìm kiếm bằng tên sản phầm, lọc giá từ cao đến thấp và ngược lại 
<br>	(*router).Get("/", controller.GetProducts) 
 <br>   //lấy product theo ID để đưa vào chi tiết sản phẩm
<br>	(*router).Get("/:id", controller.GetProductById) 
}
<br>//API Category
<br>func routeCategory(router *fiber.Router) {
<br>    ////Liệt kê các Category
<br>	(*router).Get("/", controller.GetCategories)
}

# FRONT-END

PORT:8080
<br>Cài đặt: npm install or yarn install
<br>Development: yarn server



