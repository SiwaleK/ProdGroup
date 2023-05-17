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


-- name: GetPosClientMethod :many
SELECT
    pc.is_cash,
    pc.is_paotang,
    pc.is_qrcode,
    pc.is_tongfah,
    pc.is_coupon,
    br.account_name,
    br.account_code
FROM
    posclient pc
JOIN
    branch br ON pc.branch_id = br.branch_id
WHERE
    pc.pos_client_id = $1;




-- name: GetPaymentConfig :many 
SELECT 
  br.account_name,
  br.account_code
FROM
  branch br;



