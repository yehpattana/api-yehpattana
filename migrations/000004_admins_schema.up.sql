
-- Create Admin table (corrected)
CREATE TABLE [dbo].[Admins] (
  [user_id] nvarchar(255) NOT NULL UNIQUE, 
  [user_name] nvarchar(255) NOT NULL,
  FOREIGN KEY ([user_id]) REFERENCES [dbo].[Users]([Id])
);
