CREATE TABLE Log (
    [id] INT IDENTITY(1,1) PRIMARY KEY,
	[end_point] VARCHAR(255),
	[description] VARCHAR(MAX),
    [updated_by] VARCHAR(255),
	[created_at] DATETIME2(0) NOT NULL DEFAULT(GETUTCDATE()),
);
