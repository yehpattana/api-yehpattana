-- Create Orders Table
CREATE TABLE
    Orders (
        order_id INT PRIMARY KEY IDENTITY (1, 1),
        order_detail VARBINARY(MAX),
        customer_id INT NOT NULL,
        status NVARCHAR (255) NOT NULL DEFAULT 'pending',
        created_at datetime2 (0) NOT NULL DEFAULT (getutcdate ()),
        updated_at datetime2 (0) NOT NULL DEFAULT (getutcdate ()),
        FOREIGN KEY (customer_id) REFERENCES Customers (customer_id),
    );