
# Order Management Service
-----------------------------
### To run this project just use docker compose command

```bash
docker compose up -d --build
```

### What's in this stack
1. This app with alpine golang image (exposed on externally on port **8080**).
2. MySQL for DB (expose internally).
3. Phpmyadmin for DB UI (exposed on externally on port **8083**).
4. Phpmyadmin User root, password root.

**Note: Please use different port if default port is busy on your environment**

### Project setup flow
While booting up the app,
1. Migrate the DB into MYSQL server running on this container
2. Two default user has been created as seeding.

### API Endpoints
* ### Login

```bash
curl --location 'localhost:8080/api/v1/login' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username" : "abc@mailinator.com",
    "password" : "321dsa"
}'
```
### Response
```json
{
    "type": "success",
    "message": "login successful",
    "code": 200,
    "data": {
        "token_type": "Bearer",
        "expires_in": 432000,
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOiIyMDI1LTAyLTA1VDIxOjExOjA5LjM4MTQ0MDQ1KzA2OjAwIiwiUGhvbmUiOiIwMTkwMTkwMTkwMiIsIlVzZXJuYW1lIjoiYWJjQG1haWxpbmF0b3IuY29tIn0.HEpnFbLHwdBKd0wMAxDSHHRHBp3GfY2Mn1ctGDrkIzo",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOiIyMDI1LTAxLTMxVDIxOjQxOjA5LjM4MTQ5NjI1OSswNjowMCIsIlBob25lIjoiMDE5MDE5MDE5MDIiLCJVc2VybmFtZSI6ImFiY0BtYWlsaW5hdG9yLmNvbSJ9.9p9NPkTHGCy1MqFKGGqh4Hn78C-2o7oY_TKKxSes4Ug"
    }
}
```
* ### Create Order

```bash
curl --location 'localhost:8080/api/v1/orders' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOiIyMDI1LTAyLTA1VDE2OjA5OjQ1Ljk0MzU4NDY5NyswNjowMCIsIlBob25lIjoiMDE5MDE5MDE5MDEiLCJVc2VybmFtZSI6IjAxOTAxOTAxOTAxQG1haWxpbmF0b3IuY29tIn0.bOP_IPDPmH-pPxYsidCluK353Ssb978aOefG6fgAC38' \
--data '{
    "store_id": 131172,
    "merchant_order_id": "MID092737s",
    "recipient_name": "Jahir Raihan",
    "recipient_phone": "01789699003",
    "recipient_address": "Banani, Gulshan 2, Dhaka, Bangladesh",
    "recipient_city": 1,
    "recipient_zone": 3,
    "recipient_area": 14,
    "delivery_type": 48,
    "item_type": 2,
    "special_instruction": "",
    "item_quantity": 1,
    "item_weight": 5,
    "amount_to_collect": 1200,
    "item_description": ""
}'
```

### Response

```json
{
    "type": "success",
    "message": "Order successfully placed",
    "code": 200,
    "data": {
        "consignment_id": "97843a42-0575-4d6e-b79b-396b811a40ca",
        "order_created_at": "2025-01-31T16:35:54.702163056Z",
        "merchant_order_id": "MID092737s",
        "delivery_fee": 120,
        "order_status": "Pending"
    }
}
```

* ### Get Order List of logged in user with/without pagination

```bash
curl --location 'localhost:8080/api/v1/orders/all?paginate=true&limit=1&page=1' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOiIyMDI1LTAyLTA1VDIwOjU5OjM2LjY3ODU4NDk2NiswNjowMCIsIlBob25lIjoiMDE5MDE5MDE5MDIiLCJVc2VybmFtZSI6ImFiY0BtYWlsaW5hdG9yLmNvbSJ9.q9BGEPF2FkbRgPLo2E2a7bMQDQs4ei_KXg0ihR_cbco'
```

```json
{
  "type": "success",
  "message": "successfully found",
  "code": 200,
  "data": {
    "orders": [
      {
        "consignment_id": "c2cbaa00-ea5a-4ec1-9893-dc3a8c35bc8e",
        "order_created_at": "2025-01-31T16:26:27.771Z",
        "merchant_order_id": "MID092737s",
        "recipient_name": "Jahir Raihan",
        "recipient_address": "Banani, Gulshan 2, Dhaka, Bangladesh",
        "recipient_phone": "01789699003",
        "order_amount": 1200,
        "total_fee": 120,
        "order_type_id": 2,
        "cod_fee": 12,
        "delivery_fee": 120,
        "order_status": "Pending",
        "order_type": "Delivery",
        "item_type": "Parcel"
      }
    ],
    "paginate": {
      "total": 3,
      "current_page": 1,
      "per_page": 1,
      "total_in_page": 1,
      "last_page": 3
    }
  }
}
```




