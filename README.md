# GoToken API
- [x] Cashiers API 
- [x] Login API
- [ ] Product API
- [x] Category API
- [ ] Payment API
- [ ] Order API
- [ ] Report API


## Current Mappings

* GET `/cashiers` --> List Cashiers
* GET `/cashiers/:cashierId/passcode` --> Provides Passcode 
* POST `/cashiers/` --> Creates Cashier
* PUT `/cashiers/:cashierId` --> Update Cashier
* DELETE `/cashiers/:cashierId` --> Deletes Cashier
* POST `/cashiers/login` --> Login Cashier
* POST `/cashiers/logout` --> Logout Cashier

* GET `/categories` --> List Categories
* POST `/categories` --> Create Categories

## Points To Remember
* Database with tables cashiers and categories should already be existing on you mysql server
