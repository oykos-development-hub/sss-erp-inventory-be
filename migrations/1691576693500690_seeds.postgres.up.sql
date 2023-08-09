--itemi koji su povezani sa procurements artiklima
INSERT INTO items (
   id, article_id, type, class_type_id, depreciation_type_id, supplier_id,
   serial_number, inventory_number, title, abbreviation, internal_ownership,
   office_id, target_user_profile_id, unit, amount, net_price, gross_price,
   description, date_of_purchase,
   invoice_number, active, date_of_assessment,
   price_of_assessment, lifetime_of_assessment_in_months, created_at,
   updated_at, invoice_file_id, file_id
)
VALUES (
   1, 1, 'Kauc', 21, 24, 1,
	'SN001', 'INV001', 'Kauc', 'I1', false,
   27, 2, 'Piece', 2, 100, 120,
   'Kauc za kancelariju 1', '2023-01-15',
   'INV123456', true, '2024-01-15', 150, 24,
   NOW(), NOW(), 701, 801
),
(
   2, 2, 'Stolica', 22, 25, 1,
	'SN002', 'INV002', 'Stolica', 'I2', false,
   28, 2, 'Piece', 2, 50, 60,
   'Stolica za kancelariju 2', '2023-01-15',
   'INV1234567',true, '2024-01-15', 150, 24,
   NOW(), NOW(), 701, 801
);

--itemi koji nijesu povezani sa artiklima
INSERT INTO items (
   id, type, class_type_id, depreciation_type_id, supplier_id,
   serial_number, inventory_number, title, abbreviation, internal_ownership,
   office_id, target_user_profile_id, unit, amount, net_price, gross_price,
   description, date_of_purchase,
   invoice_number, active, date_of_assessment,
   price_of_assessment, lifetime_of_assessment_in_months, created_at,
   updated_at, invoice_file_id, file_id
)
VALUES (
   3, 'Laptop', 23, 26, 1,
	'SN003', 'INV003', 'Laptop', 'L1', true,
   27, 3, 'Piece', 2, 300, 420,
   'Laptop za kancelariju 1', '2023-01-15',
   'INV1234567', true, '2024-01-15', 150, 24,
   NOW(), NOW(), 701, 801
),
(
   4, 'Zgrada', 22, 25, 1,
	'SN004', 'INV004', 'Zgrada', 'Z2', false,
   28, 3, 'Piece', 2, 50000, 60000,
   'Zgrada', '2023-01-15',
   'INV12345678',true, '2024-01-15', 50000, 480,
   NOW(), NOW(), 701, 801
),
(
   5, 'Livada', 24, 27, 1,
	'SN005', 'INV005', 'Livada', 'L2', true,
   27, 3, 'Piece', 2, 30000, 42000,
   'Livada', '2023-01-15',
   'INV123567', true, '2024-01-15', 28000, 480,
   NOW(), NOW(), 701, 801
);

--real estates za zgradu i livadu

INSERT INTO real_estates (
    id,item_id, title, type_id, square_area, land_serial_number,
    estate_serial_number, ownership_type, ownership_scope,
    ownership_investment_scope, limitations_description,
    limitation_id, property_document, document, file_id,
    created_at, updated_at
)
VALUES
    (1,4, 'Zgrada', 'Zgrada', 500, 'LSN001', 'ESN001', 'Private', 'Scope 1', 'Investment Scope 1', 'No limitations', 'Limitation ID 1', 'Property Doc 1', 'Document 1', 201, NOW(), NOW()),
    (2,5, 'Livada', 'Livada', 1000, 'LSN002', 'ESN002', 'Public', 'Scope 2', 'Investment Scope 2', 'Limitations exist', 'Limitation ID 2', 'Property Doc 2', 'Document 2', 202, NOW(), NOW());

--dispatches

INSERT INTO dispatches (
	id,
    type,
    source_user_profile_id,
    source_organization_unit_id,
    target_user_profile_id,
    target_organization_unit_id,
    is_accepted,
    serial_number,
    office_id,
    dispatch_description,
    file_id,
    created_at,
    updated_at
)
VALUES
    (1,'Outgoing', 2, 1, 4, 2, true, 'DS001', 27, 'Prenos izmedju sudova Niksic-Berane', 501, NOW(), NOW()),
    (2,'Incoming', 4, 2, 2, 1, true, 'DS002', 27, 'Prenos izmedju sudova Berane-NIksic', 502, NOW(), NOW()),
    (3,'Outgoing', 6, 5, null, 2, false, 'DS003', 28, 'Prenos IT-Berane', 503, NOW(), NOW());

--assessments, dodata je i jedna stara procjena i tri aktivne

INSERT INTO assessments (
	id, inventory_id, active, depreciation_type_id, user_profile_id, gross_price_new, gross_price_difference, date_of_assessment, created_at, updated_at, file_id
)
VALUES
    (1,3, true, 24, 3, 500, 200, '2023-01-15', NOW(), NOW(), 401),
    (2,4, true, 24, 3, 30000, 3000, '2023-02-20', NOW(), NOW(), 402),
    (3,5, true, 25, 4, 8000, 1000, '2023-03-10', NOW(), NOW(), 403),
    (4,5, false, 25,4, 9000, 1000, '2023-03-10', NOW(), NOW(), 403);

--dispatch_items

insert into dispatch_items(id, dispatch_id, inventory_id)
values
(1,1,1),
(2,1,2),
(3,1,3),
(4,2,1),
(5,2,2),
(6,2,3);


