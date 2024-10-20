# EcommerceManager

EcommerceManager is a simple e-commerce management application built using Go, Gin framework, GORM, and SQLite. It supports basic CRUD operations for four core models: Categories, Products, Variants, and Orders.

## Features

- Manage Categories (with hierarchical relationships)
- Manage Products and their Variants
- Manage Customer Orders with inventory management
- Fully implemented CRUD APIs for each model
- SQLite as the database for persistence

## Technology Stack

- **Language:** Go
- **Framework:** Gin (for HTTP routing)
- **ORM:** GORM (for database interactions)
- **Database:** SQLite

## Models

### 1. Category

Represents product categories with hierarchical relationships (parent-child).

- **Fields:**
  - `ID` (primary key)
  - `Name` (category name)
  - `ParentID` (foreign key to itself for subcategories)
  - `ChildCategories` (subcategories)
  - `ChildProducts` (products)
  - `Variants` (Variants)
  - `CreatedAt`, `UpdatedAt`

### 2. Product

Represents products that belong to categories and can have multiple variants.

- **Fields:**
  - `ID` (primary key)
  - `Name` (product name)
  - `CategoryID` (foreign key to `Category`)
  - `Description` (product details)
  - `Variants` (Variants)
  - `Price` (base price)
  - `CreatedAt`, `UpdatedAt`

### 3. Variant

Represents product variants (e.g., size, color).

- **Fields:**
  - `ID` (primary key)
  - `ProductID` (foreign key to `Product`)
  - `Name` (variant name, e.g., size or color)
  - `Price` (variant-specific price)
  - `quantity` (available quantity)
  - `CreatedAt`, `UpdatedAt`

### 4. Order

Represents customer orders for products.

- **Fields:**
  - `ID` (primary key)
  - `VariantID` (foreign key to `Variant`)
  - `Quantity` (number of items ordered)
  - `TotalPrice` (total cost of the order)
  - `CreatedAt`, `UpdatedAt`

## API Endpoints

### Category

- `GET /api/category` - Retrieve all categories
- `POST /api/category` - Create a new category
- `GET /api/category/:id` - Get category by ID
- `PUT /api/category/:id` - Update a category
- `DELETE /api/category/:id` - Delete a category

### Product

- `GET /api/product` - Retrieve all products
- `POST /api/product` - Create a new product
- `GET /api/product/:id` - Get product by ID
- `PUT /api/product/:id` - Update a product
- `DELETE /api/product/:id` - Delete a product

### Variant

- `GET /api/variant` - Retrieve all variants
- `POST /api/variant` - Create a new variant
- `GET /api/variant/:id` - Get variant by ID
- `PUT /api/variant/:id` - Update a variant
- `DELETE /api/variant/:id` - Delete a variant

### Order

- `GET /api/order` - Retrieve all orders
- `POST /api/order` - Create a new order

## Setup and Installation

### Prerequisites

- Go 1.18 or higher
- SQLite

### Steps to Run the Project

1. Clone the repository:
    ```bash
    git clone https://github.com/AbhinavP003/GO-APPS.git
    ```

2. Navigate to the project directory:
    ```bash
    cd EcommerceManager
    ```

3. Install dependencies:
    ```bash
    go mod tidy
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

5. The API will be available at `http://localhost:8080`.

## Database Schema

GORM will automatically create and manage the SQLite schema for the models upon running the application.

## Future Enhancements

- Add authentication for secure API access
- Implement pagination for large datasets

## Postman collection

-https://www.postman.com/maintenance-geologist-73425626/ecommerce-manager/collection/dbzmvfz/category?action=share&creator=22805703


# Swagger docs
Swagger UI  can be accessed for running application on : 
http://localhost:8080/swagger/index.html
