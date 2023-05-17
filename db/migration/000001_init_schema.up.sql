CREATE TABLE prodgroup (
  prodgroupid INT NOT NULL,
  th_name VARCHAR(255),
  en_name VARCHAR(255)
);

CREATE TABLE promotion (
  Promotionid VARCHAR(36),
  Promotiontitle VARCHAR(36),
  PromotionType INT NOT NULL,
  Startdate timestamptz NOT NULL,
  Enddate timestamptz NOT NULL,
  Description VARCHAR(1024),
  Condition JSON NOT NULL

);

CREATE TABLE promotion_applied_items_id (
  Promotiondetail_id VARCHAR(36),
  Promotionid VARCHAR(36),
  skuid VARCHAR(36)
);

CREATE TABLE payment_method (
  PaymentMethodID INT NOT NULL,
  PaymentName VARCHAR(255)
);

CREATE TABLE posclient (
	pos_client_id VARCHAR(36) NULL,
	branch_id VARCHAR(36) NULL,
	merchant_id VARCHAR(36) NULL,
	rd_number VARCHAR(36) NULL,
	is_drawer SMALLINT NULL,
	is_barcode SMALLINT NULL,
	is_cash SMALLINT NULL,
	is_qrcode SMALLINT NULL,
	is_paotang SMALLINT NULL,
	is_tongfah SMALLINT NULL,
	is_coupon SMALLINT NULL,
	session_type SMALLINT NULL,
	barcode_reader_type SMALLINT NULL,
	printer_type SMALLINT NULL,
	is_active SMALLINT NOT NULL,
	pos_running VARCHAR(5) NULL,
	fr_pos_running VARCHAR(5) NULL,
	payment_mode SMALLINT NULL
);

CREATE TABLE branch (
	branch_id VARCHAR(36) NULL,
	merchant_id VARCHAR(36) NULL,
	branch_no VARCHAR(10) NULL,
	branch_name VARCHAR(50) NULL,
	branch_address VARCHAR(255) NULL,
	branch_email VARCHAR(50) NULL,
	account_name VARCHAR(255) NULL,
	account_code VARCHAR(255) NULL,
	is_active TINYINT NOT NULL,
	branch_address2 VARCHAR(255) NULL,
	branch_subdistrict VARCHAR(30) NULL,
	branch_district VARCHAR(30) NULL,
	branch_province VARCHAR(30) NULL,
	branch_zipcode VARCHAR(10) NULL,
	is_inventory TINYINT NULL,
	is_alert_inventory TINYINT NULL
);

ALTER TABLE promotion_applied_items_id ADD FOREIGN KEY (PromotionDetailID) REFERENCES promotion (PromotionID);
