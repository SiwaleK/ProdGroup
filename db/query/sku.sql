-- name: GetPromotion :many
SELECT * FROM promotion;


-- name: GetPromotionByID :one
SELECT * FROM promotion 
WHERE Promotionid =$1 ;

-- name: GetPromotionAppliedItemID :many
SELECT * FROM promotion_applied_items_id; 

-- name: GetProdgroup :many
SELECT * FROM prodgroup;


-- name: GetPaymentMethod :many 
SELECT * FROM payment_method;

-- name: GetPaymentConfig :many 
SELECT
    pc.is_cash,
    pc.is_qrcode,
    pc.is_paotang,
    pc.is_tongfah,
    pc.is_coupon,
    br.account_name,
    br.account_code

FROM
    posclient pc,
    branch br
   
;

--name: UpdateBranch :exec
UPDATE 
     branch
SET 
   
    merchant_id = $2,
    branch_no = $3, 
    branch_name = $4, 
    branch_address = $5, 
    branch_email = $6, 
    account_name = $7, 
    account_code = $8, 
    is_active = $9, 
    branch_address2 =$10, 
    branch_subdistrict = $11, 
    branch_district = $12, 
    branch_province = $13, 
    branch_zipcode = $14,
    is_inventory = $15, 
    is_alert_inventory = $16

WHERE
    branch_id = $1;
    
