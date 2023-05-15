-- name: GetPromotion :many
SELECT * FROM promotion;


-- name: GetPromotionByID :one
SELECT * FROM promotion 
WHERE Promotionid =$1 ;

-- name: GetPromotionAppliedItemID :many
SELECT * FROM promotion_applied_items_id; 


-- name: PostPromotion :exec
WITH promotion AS (
  INSERT INTO promotion (Promotionid,Promotiontitle, PromotionType, Startdate, Enddate, Description, Condition)
  VALUES ($1, $2, $3, $4, $5, $6,$7)
  RETURNING *
),
promotion_applied_items_id AS (
  INSERT INTO promotion_applied_items_id (Promotiondetail_id, Promotionid, skuid)
  VALUES ($8, (SELECT Promotionid FROM promotion), $9)
  RETURNING *
)
SELECT promotion.*, promotion_applied_items_id.*
FROM promotion, promotion_applied_items_id;

-- name: PostPromotionTable :exec
INSERT INTO promotion (Promotionid, Promotiontitle, PromotionType, Startdate, Enddate, Description, Condition)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: PostPromotionApplied :exec
INSERT INTO promotion_applied_items_id (Promotiondetail_id, Promotionid, skuid)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetProdgroup :many
SELECT * FROM prodgroup;

-- name: GetProdgroupByID :many 
SELECT * FROM prodgroup
WHERE prodgroupid =$1;


-- name: GetPaymentMethod :many 
SELECT * FROM payment_method;

-- name: GetPaymentConfig :many 
SELECT
    pc.is_cash,
    pc.is_qrcode,
    pc.is_paotang,
    pc.is_tongfah,
    pc.is_coupon,
    pc.printer_type,
    br.account_name,
    br.account_code

FROM
    posclient pc,
    branch br
   
;

-- Add unique constraint to posclient table
ALTER TABLE posclient
ADD CONSTRAINT uq_posclient_unique_columns UNIQUE (is_cash, is_qrcode, is_paotang, is_tongfah, is_coupon, printer_type);


-- name: UpsertPaymentConfig :exec
WITH upsert_data AS (
    INSERT INTO posclient (is_cash, is_qrcode, is_paotang, is_tongfah, is_coupon, printer_type)
    VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT (pos_client_id)
    DO UPDATE SET
        is_cash = EXCLUDED.is_cash,
        is_qrcode = EXCLUDED.is_qrcode,
        is_paotang = EXCLUDED.is_paotang,
        is_tongfah = EXCLUDED.is_tongfah,
        is_coupon = EXCLUDED.is_coupon,
        printer_type = EXCLUDED.printer_type
    RETURNING pos_client_id
)
SELECT *
FROM upsert_data;
