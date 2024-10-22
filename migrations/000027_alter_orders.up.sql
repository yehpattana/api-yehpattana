CREATE SEQUENCE dbo.OrderSeq
    AS INT
    START WITH 1
    INCREMENT BY 1;
    
ALTER TABLE Orders
    ADD order_no NVARCHAR(255);

ALTER TABLE Orders
    ADD shipping_address NVARCHAR(MAX);

ALTER TABLE Orders
    ADD tracking_no NVARCHAR(255);