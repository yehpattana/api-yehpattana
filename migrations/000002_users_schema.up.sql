-- Create User table
CREATE TABLE [dbo].[Users] (
  [Id] nvarchar(255) NOT NULL PRIMARY KEY,
  [created_at] datetime2(0) NOT NULL DEFAULT(getutcdate()),
  [updated_at] datetime2(0) NOT NULL DEFAULT(getutcdate()),
  [email] nvarchar(255) NOT NULL UNIQUE,
  [password] nvarchar(255) NOT NULL,
  [is_actived] BIT NOT NULL,
  [role] nvarchar(255) NOT NULL,
);
