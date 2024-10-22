CREATE TABLE [dbo].[Stock] (
  [id] NVARCHAR(255) NOT NULL PRIMARY KEY,
  [product_id] NVARCHAR(255) NOT NULL,
  [size] NVARCHAR(50) NOT NULL,
  [size_remark] NVARCHAR(255),
  [quantity] INT NOT NULL,
  [price] DECIMAL(12, 4) NOT NULL,
  [item_status] NVARCHAR(255) NOT NULL CHECK ([item_status] IN ('available', 'out_of_stock')),
  [created_at] DATETIME2(0) NOT NULL DEFAULT(GETUTCDATE()),
  [updated_at] DATETIME2(0) NOT NULL DEFAULT(GETUTCDATE()),
  FOREIGN KEY ([product_id]) REFERENCES [dbo].[Products]([id])
);