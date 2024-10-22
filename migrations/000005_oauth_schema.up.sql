CREATE TABLE [dbo].[Oauth] (
  [Id] nvarchar(255) NOT NULL PRIMARY KEY,
  [user_id] nvarchar(255) NOT NULL,
  [access_token] nvarchar(max) NOT NULL,
  [refresh_token] nvarchar(max) NOT NULL,
  [created_at] datetime2(0) NOT NULL DEFAULT(getutcdate()),
  [updated_at] datetime2(0) NOT NULL DEFAULT(getutcdate()),
  FOREIGN KEY ([user_id]) REFERENCES [dbo].[Users]([Id])
);
