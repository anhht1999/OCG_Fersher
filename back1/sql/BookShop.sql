CREATE TABLE `products` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `author` varchar(50) NOT NULL,
  `publisher` varchar(50),
  `supplier` varchar(50),
  `year` int,
  `page` int,
  `price` int,
  `category_id` int
);

CREATE TABLE `product_prices` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `price` int,
  `discount` int,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `product_medias` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `uri` varchar(255) NOT NULL,
  `media_type` ENUM ('photo', 'vIDeo', 'PDF')
);

CREATE TABLE `categories` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `parent_id` int
);

CREATE TABLE `relate_products` (
  `product_id` int,
  `relate_id` int,
  `relation` ENUM ('oldversion', 'newversion', 'similar', 'recommend') NOT NULL
);

CREATE TABLE `users` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `mobile` varchar(255) UNIQUE,
  `password` varchar(255) NOT NULL
);

CREATE TABLE `customers` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int
);

CREATE TABLE `addresses` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int
);

CREATE TABLE `cities` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`ID`);

ALTER TABLE `product_prices` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `product_medias` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `relate_products` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `relate_products` ADD FOREIGN KEY (`relate_id`) REFERENCES `products` (`ID`);

ALTER TABLE `customers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`ID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`ID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`city_id`) REFERENCES `cities` (`ID`);
