ALTER TABLE Menu
    ADD company_id NVARCHAR(255),
    FOREIGN KEY(company_id) REFERENCES Companies(id);

ALTER TABLE Sub_Menu
    ADD company_id NVARCHAR(255);