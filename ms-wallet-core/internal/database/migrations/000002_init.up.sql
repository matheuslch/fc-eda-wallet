INSERT INTO clients (id,name,email,created_at) VALUES
	 ('59c929ee-0b5b-4dbb-a4c5-5f32b91e84aa','John Doe','John@j.com','2024-03-12 20:49:33.000'),
	 ('7aaae6b9-b15b-4217-aa1f-fcfdc4265ff3','Jane Doe','jane@j.com','2024-03-12 20:48:58.000');

INSERT INTO accounts (id,client_id,balance,created_at) VALUES
	 ('494f2113-a136-4eb6-a682-9bc5cb30ba80','59c929ee-0b5b-4dbb-a4c5-5f32b91e84aa',1000,'2024-03-12 20:49:42.000'),
	 ('96ded354-3410-40b3-8ca2-a3dce2deb269','7aaae6b9-b15b-4217-aa1f-fcfdc4265ff3',100,'2024-03-12 20:49:15.000');

INSERT INTO transactions (id,account_id_from,account_id_to,amount,created_at) VALUES
	 ('655f1989-125c-4efb-8cdd-af8071fc91f2','96ded354-3410-40b3-8ca2-a3dce2deb269','494f2113-a136-4eb6-a682-9bc5cb30ba80',10.0,'2024-03-12 20:51:53');