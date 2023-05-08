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
SELECT * FROM paymentmethod;