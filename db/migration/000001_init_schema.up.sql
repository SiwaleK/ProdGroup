



CREATE TABLE IF NOT EXISTS readallprovinceresp (
  ProvinceId   int,
	ProvinceName varchar(20)
);


CREATE TABLE IF NOT EXISTS backendposdatasku (
  skuid varchar(36) PRIMARY KEY,
  barcodepos varchar(20),
  productname varchar(255),
  brandid int,
  productgroupid int,
  productcatid int,
  productsubcatid int,
  productsizeid int,
  productunit int,
  packsize varchar(255),
  unit int,
  banforpracharat int,
  isvat smallint,
  createby varchar(36),
  createdate timestamptz NOT NULL DEFAULT (now()),
  isactive smallint NOT NULL,
  merchantid varchar(36),
  mapsku varchar(36),
  isfixprice smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS backendposdata_sku_branch_price (
  skuid VARCHAR(36),
  merchantid VARCHAR(36),
  branchid VARCHAR(36),
  price NUMERIC(16,4),
  startdate timestamptz NOT NULL DEFAULT (now()),
  enddate timestamptz,
  isactive smallint NOT NULL,
  PRIMARY KEY (skuid, merchantid, branchid)
);

CREATE TABLE IF NOT EXISTS prodsubcat (
	prodsubcatid INT NOT NULL,
	th_name VARCHAR(255) NULL,
	en_name VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS "public"."prodgroup" (
    "prodgroupid" int4 NOT NULL,
    "th_name" varchar(255),
    "en_name" varchar(255),
    PRIMARY KEY ("prodgroupid")
);

CREATE TABLE IF NOT EXISTS promotion (
	promotionid VARCHAR(36) NULL,
  promotiontitle varchar(36) NULL,
	promotiontype INT NOT NULL,
	startdate timestamptz NOT NULL,
	enddate timestamptz NOT NULL,
	Description VARCHAR(1024) NULL,
	conditions JSON NOT NULL
);

CREATE TABLE IF NOT EXISTS promotion_applied_items_id (
	promotiondetailid VARCHAR(36) NULL,
	promotionid VARCHAR(36) NULL,
	skuid VARCHAR(36) NULL
);


CREATE TABLE IF NOT EXISTS brand (
	brandid SERIAL NOT NULL,
	th_brand VARCHAR(255) NULL,
	en_brand VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS prodgroup_prodcat (
	prodgroupid INT NOT NULL,
	prodcatid INT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
	prodcatid INT NOT NULL,
	th_name VARCHAR(255) NULL,
	en_name VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS geographies (
  geographyId serial PRIMARY KEY,
  geographyName varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS provinces (
  provinceId SERIAL PRIMARY KEY,
  provinceCode VARCHAR(2) NOT NULL,
  provinceName VARCHAR(150) NOT NULL,
  provinceNameEN VARCHAR(150) NOT NULL,
  geographyId INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS districts (
  districtId serial NOT NULL,
  districtCode varchar(4)  NOT NULL,
  districtName varchar(150)  NOT NULL,
  districtNameEN varchar(150)  NOT NULL,
  provinceId integer NOT NULL DEFAULT 0,
  PRIMARY KEY (districtId)
);

CREATE TABLE IF NOT EXISTS subdistricts (
  subDistrictId varchar(6) NOT NULL,
  zipCode int NOT NULL,
  subDistrictName varchar(150) NOT NULL,
  subDistrictNameEN varchar(150) NOT NULL,
  districtId int NOT NULL DEFAULT '0'
);

CREATE TABLE IF NOT EXISTS "public"."branch" (
    "branchid" uuid,
    "merchantid" uuid,
    "branchno" varchar(10),
    "branchname" varchar(50),
    "branchaddress" varchar(255),
    "branchemail" varchar(50),
    "accountname" varchar(255),
    "accountcode" varchar(255),
    "isactive" bool NOT NULL,
    "branchaddress2" varchar(255),
    "branchsubdistrict" varchar(30),
    "branchdistrict" varchar(30),
    "branchprovince" varchar(30),
    "branchzipcode" varchar(10),
    "isinventory" bool,
    "isalertinventory" bool
);

CREATE TABLE IF NOT EXISTS "public"."merchant" (
    "merchantid" uuid,
    "merchantname" varchar(255),
    "merchantaddress" varchar(255),
    "merchanttel" varchar(50),
    "taxid" varchar(20),
    "isactive" bool,
    "isvat" bool
);

CREATE TABLE IF NOT EXISTS "public"."posclient" (
    "posclientid" uuid,
    "branchid" uuid,
    "merchantid" uuid,
    "rdnumber" uuid,
    "isdrawer" bool,
    "isbarcode" bool,
    "iscash" bool,
    "isqrcode" bool,
    "ispaotang" bool,
    "istongfah" bool,
    "iscoupon" bool,
    "sessiontype" int2,
    "barcodereadertype" int2,
    "printertype" int2,
    "isactive" bool NOT NULL,
    "posrunning" varchar(5),
    "frposrunning" varchar(5),
    "paymentmode" int2
);

CREATE TABLE IF NOT EXISTS "public"."useraccount" (
    "userid" uuid,
    "login" varchar(255),
    "password" varchar(255),
    "lastlogin" timestamp,
    "authtype" bpchar(1),
    "usertype" bpchar(1),
    "failedcount" int4,
    "createby" uuid,
    "createdate" timestamp NOT NULL,
    "updateby" uuid,
    "updatedate" timestamp NOT NULL,
    "isactive" bool NOT NULL
);

CREATE TABLE IF NOT EXISTS "public"."userlogin" (
    "userid" uuid,
    "posclientid" uuid,
    "lastlogin" timestamp,
    "lastlogout" timestamp,
    "version" varchar(255)
);

CREATE TABLE IF NOT EXISTS saleitem (
    SaleItemID VARCHAR(36) NULL,
    SaleOrderID VARCHAR(36) NULL,
    Seq INT NOT NULL,
    SKUID VARCHAR(36) NULL,
    ProductName VARCHAR(255) NULL,
    Price DECIMAL(16,4) NULL,
    Quantity DECIMAL(16,4) NULL,
    FullPrice DECIMAL(16,4) NULL,
    BeforeVatSale DECIMAL(16,4) NULL,
    AfterVatSale DECIMAL(16,4) NULL,
    POSClientID VARCHAR(36) NULL,
    BranchID VARCHAR(36) NULL,
    CompCode VARCHAR(36) NULL,
    PromotionID INT NULL,
    CreateBy VARCHAR(36) NULL,
    CreateDate timestamptz NOT NULL,
    IsActive SMALLINT NOT NULL,
    VoidType SMALLINT NULL,
    VoidBy VARCHAR(36) NULL,
    VoidDate timestamptz NULL
);
CREATE TABLE IF NOT EXISTS backendposdata_payment (
    PaymentID VARCHAR(36) NULL,
    SaleOrderID VARCHAR(36) NULL,
    Seq INT NOT NULL,
    PaymentMethod SMALLINT NOT NULL,
    Amount DECIMAL(16,4) NULL,
    AmountRecieve DECIMAL(16,4) NULL,
    POSSessionID VARCHAR(255) NULL,
    POSClientID VARCHAR(36) NULL,
    CreateBy VARCHAR(36) NULL,
    CreateDate timestamptz NOT NULL,
    IsActive SMALLINT NOT NULL,
    VoidType SMALLINT NULL,
    VoidBy VARCHAR(36) NULL,
    VoidDate timestamptz NULL
);

CREATE TABLE IF NOT EXISTS saleorder (
    SaleOrderID VARCHAR(36) NULL,
    DocNo VARCHAR(16) NULL,
    POSSessionID VARCHAR(36) NULL,
    BeforeVATSale DECIMAL(16,4) NULL,
    TotalDiscount DECIMAL NOT NULL CHECK (TotalDiscount >= 0),
    VATSale DECIMAL(16,4) NULL,
    TotalSale DECIMAL(16,4) NULL,
    POSClientID VARCHAR(36) NULL,
    BranchID VARCHAR(36) NULL,
    MerchantID VARCHAR(36) NULL,
    MemberID VARCHAR(36) NULL,
    Status SMALLINT NOT NULL,
    CreateBy VARCHAR(36) NULL,
    CreateDate timestamptz NOT NULL,
    IsActive SMALLINT NOT NULL,
    VoidType SMALLINT NULL,
    VoidBy VARCHAR(36) NULL,
    VoidDate timestamptz NULL
);

CREATE TABLE IF NOT EXISTS payment_method (
  PaymentMethodID INT NOT NULL,
  PaymentName VARCHAR(255)
);