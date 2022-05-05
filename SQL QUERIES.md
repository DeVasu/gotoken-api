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

curl --location --request PUT 'http://localhost:3030/products/1' \
--data-raw '{
    "categoryId": 1
}'

CREATE TABLE payments (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    logo varchar(255) ,
    updatedAt varchar(255),
    createdAt varchar(255),
    PRIMARY KEY(id)
)   