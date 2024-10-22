ALTER TABLE [dbo].[Colors]
ALTER COLUMN [code] nvarchar(255);

ALTER TABLE [dbo].[Colors]
ADD CONSTRAINT UQ_Colors_Code UNIQUE ([code]);

ALTER TABLE [dbo].[Products]
	ADD CONSTRAINT FK_Products_Colors FOREIGN KEY (color_code) REFERENCES [dbo].[Colors] (code);