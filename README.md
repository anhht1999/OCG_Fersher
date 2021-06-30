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
![image](https://user-images.githubusercontent.com/48313937/123949267-9289eb00-d9cc-11eb-81cb-f9f627b4cb01.png)
![image](https://user-images.githubusercontent.com/48313937/123949339-a6cde800-d9cc-11eb-94ce-d50e620a17da.png)
![image](https://user-images.githubusercontent.com/48313937/123949387-b4836d80-d9cc-11eb-8bfc-a9d62fc1bc7c.png)
![image](https://user-images.githubusercontent.com/48313937/123949431-c238f300-d9cc-11eb-9d16-62e8b7fb78f4.png)
![image](https://user-images.githubusercontent.com/48313937/123949492-d41a9600-d9cc-11eb-83ae-a613b04acafc.png)
![image](https://user-images.githubusercontent.com/48313937/123949603-f1e7fb00-d9cc-11eb-81ad-5ef310a12499.png)
![image](https://user-images.githubusercontent.com/48313937/123949652-ff04ea00-d9cc-11eb-80c9-9de2e729149a.png)






