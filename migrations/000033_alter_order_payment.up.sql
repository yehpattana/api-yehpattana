ALTER TABLE Orders
    ADD payment_status NVARCHAR (255) NOT NULL DEFAULT 'pending';
        
ALTER TABLE Orders
    ADD payment_id NVARCHAR (255);