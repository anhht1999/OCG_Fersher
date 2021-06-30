# OCG_Fersher

# BACK-END

PORT: 3000

CSDL ở thư mục sql

Chạy file app.go (lệnh go run app.go)
#### Các API hoàn thành 
//API Product 
func routeProduct(router *fiber.Router) {
    //Liệt kê Prodcuts, phân trang ,tìm kiếm bằng tên sản phầm, lọc giá từ cao đến thấp và ngược lại 
	(*router).Get("/", controller.GetProducts) 
    //lấy product theo ID để đưa vào chi tiết sản phẩm
	(*router).Get("/:id", controller.GetProductById) 
}
//API Category
func routeCategory(router *fiber.Router) {
    ////Liệt kê các Category
	(*router).Get("/", controller.GetCategories)
}

# FRONT-END

PORT:8080
Cài đặt: npm install or yarn install
Development: yarn server



