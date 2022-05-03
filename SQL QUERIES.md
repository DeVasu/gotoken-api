```CREATE TABLE products (
    id int NOT NULL AUTO_INCREMENT,
    categoryId int NOT NULL,
    name varchar(255) NOT NULL,
    image varchar(255) NOT NULL,
    price int NOT NULL,
    stock int NOT NULL,
    updatedAt datetime,
    createdAt datetime,
    discountQty int,
    discountType varchar(40),
    discountResult int,
    discountExpiredAt varchar(255),
    PRIMARY KEY(id)
)```

ALTER TABLE products
modify column createdAt varchar(255);