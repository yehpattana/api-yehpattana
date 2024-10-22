-- Create Customer table
CREATE TABLE [dbo].[Customers] (
  [user_id] nvarchar(255) NOT NULL, 
  [customer_id] INT IDENTITY(1,1) PRIMARY KEY,
  [contact_name] nvarchar(255) NOT NULL,
  -- company_name must be company_id only
  [company_name] nvarchar(255) NULL, 
  [vat_number] nvarchar(255) NULL,
  [phone_number] nvarchar(255) NULL,
  [address] nvarchar(255) NULL,
  [cap] nvarchar(255) NULL,
  [city] nvarchar(255) NULL,
  [province] nvarchar(255) NULL,
  [country] nvarchar(255) NULL,
  [message] nvarchar(max) NULL,
  FOREIGN KEY ([user_id]) REFERENCES [dbo].[Users]([Id])
);
