ALTER TABLE Companies ADD currency NVARCHAR (255) NOT NULL DEFAULT 'THB';

ALTER TABLE Companies ADD minimum_cost_avoid_shipping DECIMAL(12, 2) NOT NULL DEFAULT 0;