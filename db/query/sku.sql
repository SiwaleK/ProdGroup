-- name: GetPromotion :many
SELECT * FROM Promotion;


-- name: GetPromotionByID :one
SELECT * FROM Promotion 
WHERE Promotionid =$1 ;

-- name: GetPromotionAppliedItemID :many
SELECT * FROM promotion_applied_items_id; 

-- name: GetProdgroup :many
SELECT * FROM prodgroup;


