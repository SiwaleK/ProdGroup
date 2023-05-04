CREATE TABLE prodgroup (
  prodgroupid INT NOT NULL,
  th_name VARCHAR(255),
  en_name VARCHAR(255)
);

CREATE TABLE promotion (
  PromotionID VARCHAR(36),
  PromotionType INT NOT NULL,
  StartDate timestamptz NOT NULL,
  EndDate timestamptz NOT NULL,
  Description VARCHAR(1024),
  Conditions JSON NOT NULL
);

CREATE TABLE promotion_applied_items_id (
  PromotionDetailID VARCHAR(36),
  PromotionID VARCHAR(36),
  SKUID VARCHAR(36)
);

CREATE TABLE payment_method (
  PaymentMethodID INT NOT NULL,
  PaymentName VARCHAR(255)
);

ALTER TABLE promotion_applied_items_id ADD FOREIGN KEY (PromotionDetailID) REFERENCES promotion (PromotionID);
