CREATE TABLE [dbo].[Products] (
  -- Essential product information ------------
  [id] NVARCHAR(255) NOT NULL PRIMARY KEY,
  [name] NVARCHAR(255) NOT NULL,
  [product_code] NVARCHAR(255) NOT NULL,
  [master_code] NVARCHAR(255),
  [color_code] NVARCHAR(255),
  [product_status] NVARCHAR(255) NOT NULL CHECK ([product_status] IN ('available', 'hidden', 'out_of_stock')),
  [cover_image] NVARCHAR(255),
  [front_image] NVARCHAR(255),
  [back_image] NVARCHAR(255),
  [price] DECIMAL(12, 4) NOT NULL,
  -- Specific field for service only (don't have in YPT biz table) ---------------
  [use_as_primary_data] BIT DEFAULT 0,
  -- Filter Product Information ---------------
  [product_group] NVARCHAR(255),
  [season] NVARCHAR(255),
  [gender] NVARCHAR(255) NOT NULL CHECK ([gender] IN ('male', 'female', 'unisex', 'kids')),
  [product_class] NVARCHAR(255),
  [collection] NVARCHAR(255),
  [category] NVARCHAR(255),
  [brand] NVARCHAR(255),
  [is_club] BIT,
  [club_name] NVARCHAR(255),
	-- Extra product information ----------------
  [remark] NVARCHAR(MAX),
  [launch_date] DATETIME2(0),
  [size_chart] NVARCHAR(255),
  [pack_size] NVARCHAR(255),
  [current_supplier] NVARCHAR(255),
  [description] NVARCHAR(MAX),
  [fabric_content] NVARCHAR(255),
  [fabric_type] NVARCHAR(255),
  [weight] DECIMAL(18, 2),
  -- product log -------------------------
  [created_by_company] NVARCHAR(255) NOT NULL,
  [created_by] NVARCHAR(255) NOT NULL,
  [edited_by] NVARCHAR(255) NOT NULL,
  [created_at] DATETIME2(0) NOT NULL DEFAULT(GETUTCDATE()),
  [updated_at] DATETIME2(0) NOT NULL DEFAULT(GETUTCDATE())
);

