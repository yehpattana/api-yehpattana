CREATE TABLE [dbo].[Companies] (
  [id] NVARCHAR(255) NOT NULL PRIMARY KEY,
  [company_code] NVARCHAR(255) NOT NULL,
  [company_name] NVARCHAR(255) NOT NULL,
  [logo] NVARCHAR(255) NOT NULL,
);