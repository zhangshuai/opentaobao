package tbtrade

import (
	"sync"
)

// Trade 结构体
type Trade struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 优惠详情
	PromotionDetails []PromotionDetail `json:"promotion_details,omitempty" xml:"promotion_details>promotion_detail,omitempty"`
	// 物流标签
	ServiceTags []LogisticsTag `json:"service_tags,omitempty" xml:"service_tags>logistics_tag,omitempty"`
	// 服务子订单列表
	ServiceOrders []ServiceOrder `json:"service_orders,omitempty" xml:"service_orders>service_order,omitempty"`
	// 子单物流发货信息
	LogisticsInfos []LogisticsInfo `json:"logistics_infos,omitempty" xml:"logistics_infos>logistics_info,omitempty"`
	// 分阶段支付详情
	StepPayDetails []StepPayDetail `json:"step_pay_details,omitempty" xml:"step_pay_details>step_pay_detail,omitempty"`
	// 同意退款检查标识字段
	AgreeRefundChecks []AgreeRefundCheck `json:"agree_refund_checks,omitempty" xml:"agree_refund_checks>agree_refund_check,omitempty"`
	// 阶段地址详情字段
	AddressDetails []AddressDetail `json:"address_details,omitempty" xml:"address_details>address_detail,omitempty"`
	// 交付计划
	DeliveryPlan []DeliveryPlan `json:"delivery_plan,omitempty" xml:"delivery_plan>delivery_plan,omitempty"`
	// 子订单组合品物流详情（仅支持拆单发货）
	CombineLogisticsDetails []CombineLogisticsDetail `json:"combine_logistics_details,omitempty" xml:"combine_logistics_details>combine_logistics_detail,omitempty"`
	// 物流时效信息
	LogisticsConsignInfo []LogisticsConsignInfo `json:"logistics_consign_info,omitempty" xml:"logistics_consign_info>logistics_consign_info,omitempty"`
	// 发货时效修改记录
	LogisticsModifyInfo []LogisticsModifyInfo `json:"logistics_modify_info,omitempty" xml:"logistics_modify_info>logistics_modify_info,omitempty"`
	// 补差关联订单号列表
	ReceiptRelIds []string `json:"receipt_rel_ids,omitempty" xml:"receipt_rel_ids>string,omitempty"`
	// 交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 买家OpenUid
	BuyerOpenUid string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
	// 交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据 可选值 fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) auto_delivery(自动发货) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易)o2o_offlinetrade（O2O交易）step (万人团)nopaid(无付款订单)pre_auth_type(预授权0元购机交易)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 交易创建时间。格式:yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易编号 (父订单的交易编号)
	Sid string `json:"sid,omitempty" xml:"sid,omitempty"`
	// Acookie订单唯一ID。
	AcookieId string `json:"acookie_id,omitempty" xml:"acookie_id,omitempty"`
	// 交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)	* PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 建议使用trade.promotion_details查询系统优惠系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样
	AdjustFee string `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分
	PostFee string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 卖家发货时间。格式:yyyy-MM-dd HH:mm:ss
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	ReceivedPayment string `json:"received_payment,omitempty" xml:"received_payment,omitempty"`
	// 交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分
	CommissionFee string `json:"commission_fee,omitempty" xml:"commission_fee,omitempty"`
	// 买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段）
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段）
	SellerMemo string `json:"seller_memo,omitempty" xml:"seller_memo,omitempty"`
	// 买家下单的地区
	BuyerArea string `json:"buyer_area,omitempty" xml:"buyer_area,omitempty"`
	// 支付宝交易号，如：2009112081173831
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 商品图片绝对途径
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。
	CodFee string `json:"cod_fee,omitempty" xml:"cod_fee,omitempty"`
	// 货到付款物流状态。初始状态 NEW_CREATED,接单成功 ACCEPTED_BY_COMPANY,接单失败 REJECTED_BY_COMPANY,接单超时 RECIEVE_TIMEOUT,揽收成功 TAKEN_IN_SUCCESS,揽收失败 TAKEN_IN_FAILED,揽收超时 TAKEN_TIMEOUT,签收成功 SIGN_IN,签收失败 REJECTED_BY_OTHER_SIDE,订单等待发送给物流公司 WAITING_TO_BE_SENT,用户取消物流订单 CANCELED
	CodStatus string `json:"cod_status,omitempty" xml:"cod_status,omitempty"`
	// 买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分
	BuyerCodFee string `json:"buyer_cod_fee,omitempty" xml:"buyer_cod_fee,omitempty"`
	// 卖家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。卖家不承担服务费的订单：未发货的订单获取服务费为0，发货后就能获取到正确值。
	SellerCodFee string `json:"seller_cod_fee,omitempty" xml:"seller_cod_fee,omitempty"`
	// 快递代收款。精确到2位小数;单位:元。如:212.07，表示:212元7分
	ExpressAgencyFee string `json:"express_agency_fee,omitempty" xml:"express_agency_fee,omitempty"`
	// 创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送),seller(商家自配送),instant(即时配送)。
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// 买家支付宝账号
	BuyerAlipayNo string `json:"buyer_alipay_no,omitempty" xml:"buyer_alipay_no,omitempty"`
	// 收货人的姓名（加密）
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人国籍
	ReceiverCountry string `json:"receiver_country,omitempty" xml:"receiver_country,omitempty"`
	// 收货人的所在省份
	ReceiverState string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
	// 收货人的所在城市&lt;br/&gt;注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面&lt;br/&gt;建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 收货人的所在地区&lt;br/&gt;注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面&lt;br/&gt;建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
	ReceiverDistrict string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
	// 收货人街道地址
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
	// 收货人的详细地址（加密）
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收货人的邮编
	ReceiverZip string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
	// 收货人的手机号码（加密）
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收货人的电话号码（加密）
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 买家邮件地址
	BuyerEmail string `json:"buyer_email,omitempty" xml:"buyer_email,omitempty"`
	// 卖家支付宝账号
	SellerAlipayNo string `json:"seller_alipay_no,omitempty" xml:"seller_alipay_no,omitempty"`
	// 卖家手机
	SellerMobile string `json:"seller_mobile,omitempty" xml:"seller_mobile,omitempty"`
	// 卖家电话
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家姓名
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 卖家邮件地址
	SellerEmail string `json:"seller_email,omitempty" xml:"seller_email,omitempty"`
	// 交易中剩余的确认收货金额（这个金额会随着子订单确认收货而不断减少，交易成功后会变为零）。精确到2位小数;单位:元。如:200.07，表示:200 元7分
	AvailableConfirmFee string `json:"available_confirm_fee,omitempty" xml:"available_confirm_fee,omitempty"`
	// 超时到期时间。格式:yyyy-MM-dd HH:mm:ss。业务规则：前提条件：只有在买家已付款，卖家已发货的情况下才有效如果申请了退款，那么超时会落在子订单上；比如说3笔ABC，A申请了，那么返回的是BC的列表, 主定单不存在如果没有申请过退款，那么超时会挂在主定单上；比如ABC，返回主定单，ABC的超时和主定单相同
	TimeoutActionTime string `json:"timeout_action_time,omitempty" xml:"timeout_action_time,omitempty"`
	// 交易快照地址
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 交易备注。
	TradeMemo string `json:"trade_memo,omitempty" xml:"trade_memo,omitempty"`
	// 交易促销详细信息
	Promotion string `json:"promotion,omitempty" xml:"promotion,omitempty"`
	// 使用信用卡支付金额数
	CreditCardFee string `json:"credit_card_fee,omitempty" xml:"credit_card_fee,omitempty"`
	// 分阶段付款的订单状态（例如万人团订单等），目前有四个返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)，FRONT_PAID_FRONT_FORFEITED（预售定金罚没）
	StepTradeStatus string `json:"step_trade_status,omitempty" xml:"step_trade_status,omitempty"`
	// 分阶段付款的已付金额（万人团订单已付金额）
	StepPaidFee string `json:"step_paid_fee,omitempty" xml:"step_paid_fee,omitempty"`
	// 电子凭证的垂直信息
	EticketExt string `json:"eticket_ext,omitempty" xml:"eticket_ext,omitempty"`
	// 订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空
	MarkDesc string `json:"mark_desc,omitempty" xml:"mark_desc,omitempty"`
	// 订单的运费险，单位为元
	YfxFee string `json:"yfx_fee,omitempty" xml:"yfx_fee,omitempty"`
	// 运费险支付号
	YfxId string `json:"yfx_id,omitempty" xml:"yfx_id,omitempty"`
	// 运费险类型
	YfxType string `json:"yfx_type,omitempty" xml:"yfx_type,omitempty"`
	// 订单将在此时间前发出，主要用于预售订单
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 物流到货时效截单时间，格式 HH:mm
	ArriveCutTime string `json:"arrive_cut_time,omitempty" xml:"arrive_cut_time,omitempty"`
	// 导购宝=crm
	O2o string `json:"o2o,omitempty" xml:"o2o,omitempty"`
	// 导购员id
	O2oGuideId string `json:"o2o_guide_id,omitempty" xml:"o2o_guide_id,omitempty"`
	// 导购员名称
	O2oGuideName string `json:"o2o_guide_name,omitempty" xml:"o2o_guide_name,omitempty"`
	// 导购员门店id
	O2oShopId string `json:"o2o_shop_id,omitempty" xml:"o2o_shop_id,omitempty"`
	// 导购门店名称
	O2oShopName string `json:"o2o_shop_name,omitempty" xml:"o2o_shop_name,omitempty"`
	// 导购宝提货方式，inshop：店内提货，online：线上发货
	O2oDelivery string `json:"o2o_delivery,omitempty" xml:"o2o_delivery,omitempty"`
	// 外部订单号
	O2oOutTradeId string `json:"o2o_out_trade_id,omitempty" xml:"o2o_out_trade_id,omitempty"`
	// 线下自提门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 拼音名
	HkEnName string `json:"hk_en_name,omitempty" xml:"hk_en_name,omitempty"`
	// 航班号
	HkFlightNo string `json:"hk_flight_no,omitempty" xml:"hk_flight_no,omitempty"`
	// 中文名
	HkChinaName string `json:"hk_china_name,omitempty" xml:"hk_china_name,omitempty"`
	// 证件号码
	HkCardCode string `json:"hk_card_code,omitempty" xml:"hk_card_code,omitempty"`
	// 证件类型001代表港澳通行证类型、002代表入台证003代表护照
	HkCardType string `json:"hk_card_type,omitempty" xml:"hk_card_type,omitempty"`
	// 航班飞行时间
	HkFlightDate string `json:"hk_flight_date,omitempty" xml:"hk_flight_date,omitempty"`
	// 性别M: 男性，F: 女性
	HkGender string `json:"hk_gender,omitempty" xml:"hk_gender,omitempty"`
	// 出生日期
	HkBirthday string `json:"hk_birthday,omitempty" xml:"hk_birthday,omitempty"`
	// 提货地点
	HkPickup string `json:"hk_pickup,omitempty" xml:"hk_pickup,omitempty"`
	// 提货地点id
	HkPickupId string `json:"hk_pickup_id,omitempty" xml:"hk_pickup_id,omitempty"`
	// 商家的预计发货时间
	EstConTime string `json:"est_con_time,omitempty" xml:"est_con_time,omitempty"`
	// 交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔
	TradeFrom string `json:"trade_from,omitempty" xml:"trade_from,omitempty"`
	// 交易外部来源：ownshop(商家官网)
	TradeSource string `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
	// 天猫国际官网直供主订单关税税费
	OrderTaxFee string `json:"order_tax_fee,omitempty" xml:"order_tax_fee,omitempty"`
	// 天猫汽车服务预约时间
	EtSerTime string `json:"et_ser_time,omitempty" xml:"et_ser_time,omitempty"`
	// 电子凭证预约门店地址
	EtShopName string `json:"et_shop_name,omitempty" xml:"et_shop_name,omitempty"`
	// 电子凭证核销门店地址
	EtVerifiedShopName string `json:"et_verified_shop_name,omitempty" xml:"et_verified_shop_name,omitempty"`
	// 车牌号码
	EtPlateNumber string `json:"et_plate_number,omitempty" xml:"et_plate_number,omitempty"`
	// 抢单状态&lt;br/&gt;0,未处理待分发;1,抢单中;2,已抢单;3,已发货;-1,超时;-2,处理异常;-3,匹配失败;-4,取消抢单;-5,退款取消;-9,逻辑删除
	O2oSnatchStatus string `json:"o2o_snatch_status,omitempty" xml:"o2o_snatch_status,omitempty"`
	// 天猫电子凭证家装
	EticketServiceAddr string `json:"eticket_service_addr,omitempty" xml:"eticket_service_addr,omitempty"`
	// 电子凭证扫码购-扫码支付订单type
	EtType string `json:"et_type,omitempty" xml:"et_type,omitempty"`
	// 垂直市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 门店预约自提订单标
	Obs string `json:"obs,omitempty" xml:"obs,omitempty"`
	// 满返红包的金额；如果没有满返红包，则值为 0.00
	PaidCouponFee string `json:"paid_coupon_fee,omitempty" xml:"paid_coupon_fee,omitempty"`
	// 线下门店自提
	ShopPick string `json:"shop_pick,omitempty" xml:"shop_pick,omitempty"`
	// 处方药审核状态
	RxAuditStatus string `json:"rx_audit_status,omitempty" xml:"rx_audit_status,omitempty"`
	// yyyyMMdd
	EsDate string `json:"es_date,omitempty" xml:"es_date,omitempty"`
	// hh:mm-hh:mm
	EsRange string `json:"es_range,omitempty" xml:"es_range,omitempty"`
	// yyyyMMdd
	OsDate string `json:"os_date,omitempty" xml:"os_date,omitempty"`
	// hh:mm-hh:mm
	OsRange string `json:"os_range,omitempty" xml:"os_range,omitempty"`
	// 主订单扩展属性;payCurrency:HKD(港币)、TWD(台币)
	TradeAttr string `json:"trade_attr,omitempty" xml:"trade_attr,omitempty"`
	// 星盘标识字段
	OmniAttr string `json:"omni_attr,omitempty" xml:"omni_attr,omitempty"`
	// 星盘业务字段
	OmniParam string `json:"omni_param,omitempty" xml:"omni_param,omitempty"`
	// assembly
	Assembly string `json:"assembly,omitempty" xml:"assembly,omitempty"`
	// 采购订单标识
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 组装O2O多阶段尾款订单的明细数据 总阶段数，当前阶数，阶段金额（单位：分），支付状态，例如 3_1_100_paid ; 3_2_2000_nopaid
	O2oStepTradeDetail string `json:"o2o_step_trade_detail,omitempty" xml:"o2o_step_trade_detail,omitempty"`
	// 特权定金订单的尾款订单ID
	O2oStepOrderId string `json:"o2o_step_order_id,omitempty" xml:"o2o_step_order_id,omitempty"`
	// 分阶段订单的特权定金订单ID
	O2oEtOrderId string `json:"o2o_et_order_id,omitempty" xml:"o2o_et_order_id,omitempty"`
	// 分阶段订单的特权定金抵扣金额，单位：分
	O2oVoucherPrice string `json:"o2o_voucher_price,omitempty" xml:"o2o_voucher_price,omitempty"`
	// 天猫国际计税优惠金额
	OrderTaxPromotionFee string `json:"order_tax_promotion_fee,omitempty" xml:"order_tax_promotion_fee,omitempty"`
	// tidStr
	TidStr string `json:"tid_str,omitempty" xml:"tid_str,omitempty"`
	// 包含的交易服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// o2oServiceMobile
	O2oServiceMobile string `json:"o2o_service_mobile,omitempty" xml:"o2o_service_mobile,omitempty"`
	// o2oServiceName
	O2oServiceName string `json:"o2o_service_name,omitempty" xml:"o2o_service_name,omitempty"`
	// o2oServiceState
	O2oServiceState string `json:"o2o_service_state,omitempty" xml:"o2o_service_state,omitempty"`
	// o2oServiceCity
	O2oServiceCity string `json:"o2o_service_city,omitempty" xml:"o2o_service_city,omitempty"`
	// o2oServiceDistrict
	O2oServiceDistrict string `json:"o2o_service_district,omitempty" xml:"o2o_service_district,omitempty"`
	// o2oServiceTown
	O2oServiceTown string `json:"o2o_service_town,omitempty" xml:"o2o_service_town,omitempty"`
	// o2oServiceAddress
	O2oServiceAddress string `json:"o2o_service_address,omitempty" xml:"o2o_service_address,omitempty"`
	// o2oStepTradeDetailNew
	O2oStepTradeDetailNew string `json:"o2o_step_trade_detail_new,omitempty" xml:"o2o_step_trade_detail_new,omitempty"`
	// o2oXiaopiao
	O2oXiaopiao string `json:"o2o_xiaopiao,omitempty" xml:"o2o_xiaopiao,omitempty"`
	// o2oContract
	O2oContract string `json:"o2o_contract,omitempty" xml:"o2o_contract,omitempty"`
	// rechargeFee
	RechargeFee string `json:"recharge_fee,omitempty" xml:"recharge_fee,omitempty"`
	// retailStoreCode
	RetailStoreCode string `json:"retail_store_code,omitempty" xml:"retail_store_code,omitempty"`
	// retailOutOrderId
	RetailOutOrderId string `json:"retail_out_order_id,omitempty" xml:"retail_out_order_id,omitempty"`
	// platformSubsidyFee
	PlatformSubsidyFee string `json:"platform_subsidy_fee,omitempty" xml:"platform_subsidy_fee,omitempty"`
	// nrOffline
	NrOffline string `json:"nr_offline,omitempty" xml:"nr_offline,omitempty"`
	// 网厅订单垂直表信息
	WttParam string `json:"wtt_param,omitempty" xml:"wtt_param,omitempty"`
	// sellerNick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// buyerNick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// nrStoreOrderId
	NrStoreOrderId string `json:"nr_store_order_id,omitempty" xml:"nr_store_order_id,omitempty"`
	// 门店 ID
	NrShopId string `json:"nr_shop_id,omitempty" xml:"nr_shop_id,omitempty"`
	// 门店名称
	NrShopName string `json:"nr_shop_name,omitempty" xml:"nr_shop_name,omitempty"`
	// 导购员ID
	NrShopGuideId string `json:"nr_shop_guide_id,omitempty" xml:"nr_shop_guide_id,omitempty"`
	// 导购员名称
	NrShopGuideName string `json:"nr_shop_guide_name,omitempty" xml:"nr_shop_guide_name,omitempty"`
	// 一小时达不处理订单
	NrNoHandle string `json:"nr_no_handle,omitempty" xml:"nr_no_handle,omitempty"`
	// 为tmall.daogoubao.cloudstore时表示云店链路
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 为1，且bizCode不为tmall.daogoubao.cloudstore时，为旗舰店订单
	CloudStore string `json:"cloud_store,omitempty" xml:"cloud_store,omitempty"`
	// 暂不公开
	DoneeNick string `json:"donee_nick,omitempty" xml:"donee_nick,omitempty"`
	// 暂不公开
	DoneeOpenUid string `json:"donee_open_uid,omitempty" xml:"donee_open_uid,omitempty"`
	// 苏宁自提门店code
	SuningShopCode string `json:"suning_shop_code,omitempty" xml:"suning_shop_code,omitempty"`
	// 允许的appkey，逗号分隔
	AllowAppkeys string `json:"allow_appkeys,omitempty" xml:"allow_appkeys,omitempty"`
	// 天猫未来店线下店铺 ID
	RetailStoreId string `json:"retail_store_id,omitempty" xml:"retail_store_id,omitempty"`
	// 区分istore订单来源
	Ua string `json:"ua,omitempty" xml:"ua,omitempty"`
	// linkedmall透传参数
	LinkedmallExtInfo string `json:"linkedmall_ext_info,omitempty" xml:"linkedmall_ext_info,omitempty"`
	// 支付渠道：0 用户主动支付 1 系统代扣 2 保险赔付
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 新零售全渠道订单：订单类型，自提订单：pickUp，电商发货：tmall，门店发货（配送、骑手）：storeSend
	RtOmniSendType string `json:"rt_omni_send_type,omitempty" xml:"rt_omni_send_type,omitempty"`
	// 新零售全渠道订单：发货门店ID
	RtOmniStoreId string `json:"rt_omni_store_id,omitempty" xml:"rt_omni_store_id,omitempty"`
	// 新零售全渠道订单：商家自有发货门店编码
	RtOmniOuterStoreId string `json:"rt_omni_outer_store_id,omitempty" xml:"rt_omni_outer_store_id,omitempty"`
	// 同城预约配送开始时间
	TcpsStart string `json:"tcps_start,omitempty" xml:"tcps_start,omitempty"`
	// 同城业务类型,com.tmall.dsd:定时送,storeDsd-fn-3-1:淘速达3公里蜂鸟配送
	TcpsCode string `json:"tcps_code,omitempty" xml:"tcps_code,omitempty"`
	// 同城预约配送结束时间
	TcpsEnd string `json:"tcps_end,omitempty" xml:"tcps_end,omitempty"`
	// 主订单商家代缴税费
	MTariffFee string `json:"m_tariff_fee,omitempty" xml:"m_tariff_fee,omitempty"`
	// 时效服务身份，如tmallPromise代表天猫时效承诺
	TimingPromise string `json:"timing_promise,omitempty" xml:"timing_promise,omitempty"`
	// 时效服务字段，服务字段，会有多个服务值，以英文半角逗号&#34;,&#34;切割
	PromiseService string `json:"promise_service,omitempty" xml:"promise_service,omitempty"`
	// 物流截单时间，分钟
	CutoffMinutes string `json:"cutoff_minutes,omitempty" xml:"cutoff_minutes,omitempty"`
	// 物流时效，相对时间，单位是天
	EsTime string `json:"es_time,omitempty" xml:"es_time,omitempty"`
	// 最晚发货时间，日期
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 最晚揽收时间，日期
	CollectTime string `json:"collect_time,omitempty" xml:"collect_time,omitempty"`
	// 最晚派送时间，日期
	DispatchTime string `json:"dispatch_time,omitempty" xml:"dispatch_time,omitempty"`
	// 最晚签收时间，日期
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 外部会员id
	OuterPartnerMemberId string `json:"outer_partner_member_id,omitempty" xml:"outer_partner_member_id,omitempty"`
	// 叶子分类
	RootCat string `json:"root_cat,omitempty" xml:"root_cat,omitempty"`
	// 1-gifting订单
	Gifting string `json:"gifting,omitempty" xml:"gifting,omitempty"`
	// 1-coffee gifting订单
	GiftingTakeout string `json:"gifting_takeout,omitempty" xml:"gifting_takeout,omitempty"`
	// 预约安装时间
	OiDate string `json:"oi_date,omitempty" xml:"oi_date,omitempty"`
	// 预约安装时间段
	OiRange string `json:"oi_range,omitempty" xml:"oi_range,omitempty"`
	// 暂不安装
	HoldInstall string `json:"hold_install,omitempty" xml:"hold_install,omitempty"`
	// 订单来源
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 同城站类型
	EasyHomeCityType string `json:"easy_home_city_type,omitempty" xml:"easy_home_city_type,omitempty"`
	// 同城站关联订单号
	NrDepositOrderId string `json:"nr_deposit_order_id,omitempty" xml:"nr_deposit_order_id,omitempty"`
	// 摊位id
	NrStoreCode string `json:"nr_store_code,omitempty" xml:"nr_store_code,omitempty"`
	// 使用淘金币的数量，以分为单位，和订单标propoint中间那一段一样，没有返回null
	Propoint string `json:"propoint,omitempty" xml:"propoint,omitempty"`
	// 是否周期送订单
	ZqsOrderTag string `json:"zqs_order_tag,omitempty" xml:"zqs_order_tag,omitempty"`
	// 天鲜配冰柜id
	TxpFreezerId string `json:"txp_freezer_id,omitempty" xml:"txp_freezer_id,omitempty"`
	// 天鲜配自提方式
	TxpReceiveMethod string `json:"txp_receive_method,omitempty" xml:"txp_receive_method,omitempty"`
	// 透出的额外信息
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 收货地址有变更，返回&#34;1&#34;
	Lm string `json:"lm,omitempty" xml:"lm,omitempty"`
	// 同城购订单来源
	BrandLightShopSource string `json:"brand_light_shop_source,omitempty" xml:"brand_light_shop_source,omitempty"`
	// 同城购渠道店id
	BrandLightShopStoreId string `json:"brand_light_shop_store_id,omitempty" xml:"brand_light_shop_store_id,omitempty"`
	// 标识完美履约订单
	IsWmly string `json:"is_wmly,omitempty" xml:"is_wmly,omitempty"`
	// 全渠道包裹信息
	OmniPackage string `json:"omni_package,omitempty" xml:"omni_package,omitempty"`
	// 新康众扩展数据
	NczExtAttr string `json:"ncz_ext_attr,omitempty" xml:"ncz_ext_attr,omitempty"`
	// 苹果发票详情
	InvoiceDetailPay string `json:"invoice_detail_pay,omitempty" xml:"invoice_detail_pay,omitempty"`
	// 苹果发票详情
	InvoiceDetailMidRefund string `json:"invoice_detail_mid_refund,omitempty" xml:"invoice_detail_mid_refund,omitempty"`
	// 苹果发票详情
	InvoiceDetailAfterRefund string `json:"invoice_detail_after_refund,omitempty" xml:"invoice_detail_after_refund,omitempty"`
	// 买卡订单本金
	ExpandCardBasicPrice string `json:"expand_card_basic_price,omitempty" xml:"expand_card_basic_price,omitempty"`
	// 买卡订单权益金
	ExpandCardExpandPrice string `json:"expand_card_expand_price,omitempty" xml:"expand_card_expand_price,omitempty"`
	// 用卡订单所用的本金
	ExpandCardBasicPriceUsed string `json:"expand_card_basic_price_used,omitempty" xml:"expand_card_basic_price_used,omitempty"`
	// 用卡订单所用的权益金
	ExpandCardExpandPriceUsed string `json:"expand_card_expand_price_used,omitempty" xml:"expand_card_expand_price_used,omitempty"`
	// 配送cp
	DeliveryCps string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
	// 业务身份
	AsdpBizType string `json:"asdp_biz_type,omitempty" xml:"asdp_biz_type,omitempty"`
	// 关联下单订单
	OrderFollowId string `json:"order_follow_id,omitempty" xml:"order_follow_id,omitempty"`
	// 送货上门标
	AsdpAds string `json:"asdp_ads,omitempty" xml:"asdp_ads,omitempty"`
	// 消费者催发货标识，lg表示消费者做过催发货
	ObTag string `json:"ob_tag,omitempty" xml:"ob_tag,omitempty"`
	// 是否疫情登记的订单。0=未登记，1=已登记
	DrugRegister string `json:"drug_register,omitempty" xml:"drug_register,omitempty"`
	// 阶段收货地址标识字段
	StageAddressType string `json:"stage_address_type,omitempty" xml:"stage_address_type,omitempty"`
	// 订单分组ID
	OgId string `json:"og_id,omitempty" xml:"og_id,omitempty"`
	// 承诺/最晚送达时间
	PromiseSignTime string `json:"promise_sign_time,omitempty" xml:"promise_sign_time,omitempty"`
	// 全渠道订单相关字段
	OmnichannelParam string `json:"omnichannel_param,omitempty" xml:"omnichannel_param,omitempty"`
	// 入参fields字段必须包含receiver_name、receiver_address、created、receiver_mobile、receiver_phone 5个字段，否则无法生成oaid。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 淘鲜达生鲜半日达
	ScenarioGroup string `json:"scenario_group,omitempty" xml:"scenario_group,omitempty"`
	// 拼团玩法垂直标
	PlayType string `json:"play_type,omitempty" xml:"play_type,omitempty"`
	// 优先发货时间
	PriorityConsignTime string `json:"priority_consign_time,omitempty" xml:"priority_consign_time,omitempty"`
	// 天猫国际实名认证状态，1：实名完成 0：实名未完成
	RealNameAuthStatus string `json:"real_name_auth_status,omitempty" xml:"real_name_auth_status,omitempty"`
	// 天猫国际第三方报关标识，1：第三方报关 0：非第三方报关
	ThirdPartyCustomsDeclaration string `json:"third_party_customs_declaration,omitempty" xml:"third_party_customs_declaration,omitempty"`
	// 补差类型，ITEM:商品补差 POSTAGE:邮费补差 OTHER：其他补差
	ReceiptType string `json:"receipt_type,omitempty" xml:"receipt_type,omitempty"`
	// 同步到卖家库的时间，taobao.trades.sold.incrementv.get接口返回此字段
	AsyncModified string `json:"async_modified,omitempty" xml:"async_modified,omitempty"`
	// 卡易售垂直表信息，去除下单ip之后的结果
	NutFeature string `json:"nut_feature,omitempty" xml:"nut_feature,omitempty"`
	// 次日达，三日达等送达类型
	LgAgingType string `json:"lg_aging_type,omitempty" xml:"lg_aging_type,omitempty"`
	// 次日达订单送达时间
	LgAging string `json:"lg_aging,omitempty" xml:"lg_aging,omitempty"`
	// 天猫直送服务
	CnService string `json:"cn_service,omitempty" xml:"cn_service,omitempty"`
	// 地址aid字段
	Aid string `json:"aid,omitempty" xml:"aid,omitempty"`
	// 交易编号 (父订单的交易编号)
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 商品数字编号
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 买家使用积分,下单时生成，且一直不变。格式:100;单位:个.
	PointFee int64 `json:"point_fee,omitempty" xml:"point_fee,omitempty"`
	// 买家实际使用积分（扣除部分退款使用的积分），交易完成后生成（交易成功或关闭），交易未完成时该字段值为0。格式:100;单位:个
	RealPointFee int64 `json:"real_point_fee,omitempty" xml:"real_point_fee,omitempty"`
	// 买家获得积分,返点的积分。格式:100;单位:个。返点的积分要交易成功之后才能获得。
	BuyerObtainPointFee int64 `json:"buyer_obtain_point_fee,omitempty" xml:"buyer_obtain_point_fee,omitempty"`
	// 买家备注旗帜（与淘宝网上订单的买家备注旗帜对应，只有买家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
	BuyerFlag int64 `json:"buyer_flag,omitempty" xml:"buyer_flag,omitempty"`
	// 卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
	SellerFlag int64 `json:"seller_flag,omitempty" xml:"seller_flag,omitempty"`
	// 物流到货时效，单位天
	ArriveInterval int64 `json:"arrive_interval,omitempty" xml:"arrive_interval,omitempty"`
	// 物流发货时效，单位小时
	ConsignInterval int64 `json:"consign_interval,omitempty" xml:"consign_interval,omitempty"`
	// 付款时使用的支付宝积分的额度,单位分，比如返回1，则为1分钱
	AlipayPoint int64 `json:"alipay_point,omitempty" xml:"alipay_point,omitempty"`
	// 天猫点券卡实付款金额,单位分
	PccAf int64 `json:"pcc_af,omitempty" xml:"pcc_af,omitempty"`
	// 交易扩展表信息
	TradeExt *TradeExt `json:"trade_ext,omitempty" xml:"trade_ext,omitempty"`
	// 扫码购关联门店
	EtShopId int64 `json:"et_shop_id,omitempty" xml:"et_shop_id,omitempty"`
	// 100
	CouponFee int64 `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
	// TOP拦截标识，0不拦截，1拦截
	TopHold int64 `json:"top_hold,omitempty" xml:"top_hold,omitempty"`
	// 聚划算一起买字段
	ForbidConsign int64 `json:"forbid_consign,omitempty" xml:"forbid_consign,omitempty"`
	// 天猫拼团拦截标示
	TeamBuyHold int64 `json:"team_buy_hold,omitempty" xml:"team_buy_hold,omitempty"`
	// 分享购拦截
	ShareGroupHold int64 `json:"share_group_hold,omitempty" xml:"share_group_hold,omitempty"`
	// 天猫国际拦截
	OfpHold int64 `json:"ofp_hold,omitempty" xml:"ofp_hold,omitempty"`
	// 聚划算火拼标记
	DelayCreateDelivery int64 `json:"delay_create_delivery,omitempty" xml:"delay_create_delivery,omitempty"`
	// top定义订单类型
	Toptype int64 `json:"toptype,omitempty" xml:"toptype,omitempty"`
	// sortInfo
	SortInfo *SortInfo `json:"sort_info,omitempty" xml:"sort_info,omitempty"`
	// 1已排序 2不排序
	Sorted int64 `json:"sorted,omitempty" xml:"sorted,omitempty"`
	// 苏宁自提门店是否有效
	SuningShopValid int64 `json:"suning_shop_valid,omitempty" xml:"suning_shop_valid,omitempty"`
	// 购物金信息输出
	ExpandcardInfo *ExpandCardInfo `json:"expandcard_info,omitempty" xml:"expandcard_info,omitempty"`
	// 天猫商家使用，订单使用的红包信息
	TmallCouponFee int64 `json:"tmall_coupon_fee,omitempty" xml:"tmall_coupon_fee,omitempty"`
	// 三方鉴定信息
	IdentifyInfo *IdentifyInfo `json:"identify_info,omitempty" xml:"identify_info,omitempty"`
	// 物流服务
	LogisticsAgreement *LogisticsAgreement `json:"logistics_agreement,omitempty" xml:"logistics_agreement,omitempty"`
	// 买家的支付宝id号，在UIC中有记录，买家支付宝的唯一标示，不因为买家更换Email账号而改变。
	AlipayId int64 `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// 卖家是否已评价。可选值:true(已评价),false(未评价)
	SellerRate bool `json:"seller_rate,omitempty" xml:"seller_rate,omitempty"`
	// 买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false
	BuyerRate bool `json:"buyer_rate,omitempty" xml:"buyer_rate,omitempty"`
	// 是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含)
	HasPostFee bool `json:"has_post_fee,omitempty" xml:"has_post_fee,omitempty"`
	// 是否3D交易
	Is3D bool `json:"is_3D,omitempty" xml:"is_3D,omitempty"`
	// 是否保障速递，如果为true，则为保障速递订单，使用线下联系发货接口发货，如果未false，则该订单非保障速递，根据卖家设置的订单流转规则可使用物流宝或者常规物流发货。
	IsLgtype bool `json:"is_lgtype,omitempty" xml:"is_lgtype,omitempty"`
	// 表示是否是品牌特卖（常规特卖，不包括特卖惠和特实惠）订单，如果是返回true，如果不是返回false。当此字段与is_force_wlb均为true时，订单强制物流宝发货。
	IsBrandSale bool `json:"is_brand_sale,omitempty" xml:"is_brand_sale,omitempty"`
	// 订单是否强制使用物流宝发货。当此字段与is_brand_sale均为true时，订单强制物流宝发货。此字段为false时，该订单根据流转规则设置可以使用物流宝或者常规方式发货
	IsForceWlb bool `json:"is_force_wlb,omitempty" xml:"is_force_wlb,omitempty"`
	// 订单中是否包含运费险订单，如果包含运费险订单返回true，不包含运费险订单，返回false
	HasYfx bool `json:"has_yfx,omitempty" xml:"has_yfx,omitempty"`
	// 买家可以通过此字段查询是否当前交易可以评论，can_rate=true可以评价，false则不能评价。
	CanRate bool `json:"can_rate,omitempty" xml:"can_rate,omitempty"`
	// 卖家是否可以对订单进行评价
	SellerCanRate bool `json:"seller_can_rate,omitempty" xml:"seller_can_rate,omitempty"`
	// 是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false
	IsPartConsign bool `json:"is_part_consign,omitempty" xml:"is_part_consign,omitempty"`
	// 表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。
	IsDaixiao bool `json:"is_daixiao,omitempty" xml:"is_daixiao,omitempty"`
	// 表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。
	IsWt bool `json:"is_wt,omitempty" xml:"is_wt,omitempty"`
	// 在返回的trade对象上专门增加一个字段zero_purchase来区分，这个为true的就是0元购机预授权交易
	ZeroPurchase bool `json:"zero_purchase,omitempty" xml:"zero_purchase,omitempty"`
	// 是否屏蔽发货
	IsShShip bool `json:"is_sh_ship,omitempty" xml:"is_sh_ship,omitempty"`
	// 邮关订单
	PostGateDeclare bool `json:"post_gate_declare,omitempty" xml:"post_gate_declare,omitempty"`
	// 跨境订单
	CrossBondedDeclare bool `json:"cross_bonded_declare,omitempty" xml:"cross_bonded_declare,omitempty"`
	// 暂不公开
	IsGift bool `json:"is_gift,omitempty" xml:"is_gift,omitempty"`
	// 是否预售
	NewPresell bool `json:"new_presell,omitempty" xml:"new_presell,omitempty"`
	// 是否优享
	YouXiang bool `json:"you_xiang,omitempty" xml:"you_xiang,omitempty"`
	// 区分istore订单和普通订单
	IsIstore bool `json:"is_istore,omitempty" xml:"is_istore,omitempty"`
	// 是否是Openmall订单
	IsOpenmall bool `json:"is_openmall,omitempty" xml:"is_openmall,omitempty"`
	// 是否是码上收订单
	VLogisticsCreate bool `json:"v_logistics_create,omitempty" xml:"v_logistics_create,omitempty"`
	// 是否是非物流订单
	QRPay bool `json:"q_r_pay,omitempty" xml:"q_r_pay,omitempty"`
	// 通用的是否预售，默认是false，需要传general_new_presell参数才能真正识别是否是预售订单
	GeneralNewPresell bool `json:"general_new_presell,omitempty" xml:"general_new_presell,omitempty"`
	// 是否是周期购订单
	IsCycleBuy bool `json:"is_cycle_buy,omitempty" xml:"is_cycle_buy,omitempty"`
	// ascp会自动流转到菜鸟仓发货
	IsForceDc bool `json:"is_force_dc,omitempty" xml:"is_force_dc,omitempty"`
	// 判断订单是否有买家留言，有买家留言返回true，否则返回false
	HasBuyerMessage bool `json:"has_buyer_message,omitempty" xml:"has_buyer_message,omitempty"`
	// 是否是智慧门店订单，只有true，或者 null 两种情况
	IsO2oPassport bool `json:"is_o2o_passport,omitempty" xml:"is_o2o_passport,omitempty"`
	// tmallDelivery
	TmallDelivery bool `json:"tmall_delivery,omitempty" xml:"tmall_delivery,omitempty"`
	// threeplTiming
	ThreeplTiming bool `json:"threepl_timing,omitempty" xml:"threepl_timing,omitempty"`
	// 无物流信息返回true，平台属性，业务不要依赖
	NoShipping bool `json:"no_shipping,omitempty" xml:"no_shipping,omitempty"`
}

var poolTrade = sync.Pool{
	New: func() any {
		return new(Trade)
	},
}

// GetTrade() 从对象池中获取Trade
func GetTrade() *Trade {
	return poolTrade.Get().(*Trade)
}

// ReleaseTrade 释放Trade
func ReleaseTrade(v *Trade) {
	v.Orders = v.Orders[:0]
	v.PromotionDetails = v.PromotionDetails[:0]
	v.ServiceTags = v.ServiceTags[:0]
	v.ServiceOrders = v.ServiceOrders[:0]
	v.LogisticsInfos = v.LogisticsInfos[:0]
	v.StepPayDetails = v.StepPayDetails[:0]
	v.AgreeRefundChecks = v.AgreeRefundChecks[:0]
	v.AddressDetails = v.AddressDetails[:0]
	v.DeliveryPlan = v.DeliveryPlan[:0]
	v.CombineLogisticsDetails = v.CombineLogisticsDetails[:0]
	v.LogisticsConsignInfo = v.LogisticsConsignInfo[:0]
	v.LogisticsModifyInfo = v.LogisticsModifyInfo[:0]
	v.ReceiptRelIds = v.ReceiptRelIds[:0]
	v.Modified = ""
	v.BuyerOpenUid = ""
	v.Title = ""
	v.Type = ""
	v.Created = ""
	v.Sid = ""
	v.AcookieId = ""
	v.Status = ""
	v.Payment = ""
	v.DiscountFee = ""
	v.AdjustFee = ""
	v.PostFee = ""
	v.TotalFee = ""
	v.PayTime = ""
	v.EndTime = ""
	v.ConsignTime = ""
	v.ReceivedPayment = ""
	v.CommissionFee = ""
	v.BuyerMemo = ""
	v.SellerMemo = ""
	v.BuyerArea = ""
	v.AlipayNo = ""
	v.BuyerMessage = ""
	v.PicPath = ""
	v.Price = ""
	v.CodFee = ""
	v.CodStatus = ""
	v.BuyerCodFee = ""
	v.SellerCodFee = ""
	v.ExpressAgencyFee = ""
	v.ShippingType = ""
	v.BuyerAlipayNo = ""
	v.ReceiverName = ""
	v.ReceiverCountry = ""
	v.ReceiverState = ""
	v.ReceiverCity = ""
	v.ReceiverDistrict = ""
	v.ReceiverTown = ""
	v.ReceiverAddress = ""
	v.ReceiverZip = ""
	v.ReceiverMobile = ""
	v.ReceiverPhone = ""
	v.BuyerEmail = ""
	v.SellerAlipayNo = ""
	v.SellerMobile = ""
	v.SellerPhone = ""
	v.SellerName = ""
	v.SellerEmail = ""
	v.AvailableConfirmFee = ""
	v.TimeoutActionTime = ""
	v.SnapshotUrl = ""
	v.TradeMemo = ""
	v.Promotion = ""
	v.CreditCardFee = ""
	v.StepTradeStatus = ""
	v.StepPaidFee = ""
	v.EticketExt = ""
	v.MarkDesc = ""
	v.YfxFee = ""
	v.YfxId = ""
	v.YfxType = ""
	v.SendTime = ""
	v.ArriveCutTime = ""
	v.O2o = ""
	v.O2oGuideId = ""
	v.O2oGuideName = ""
	v.O2oShopId = ""
	v.O2oShopName = ""
	v.O2oDelivery = ""
	v.O2oOutTradeId = ""
	v.ShopCode = ""
	v.HkEnName = ""
	v.HkFlightNo = ""
	v.HkChinaName = ""
	v.HkCardCode = ""
	v.HkCardType = ""
	v.HkFlightDate = ""
	v.HkGender = ""
	v.HkBirthday = ""
	v.HkPickup = ""
	v.HkPickupId = ""
	v.EstConTime = ""
	v.TradeFrom = ""
	v.TradeSource = ""
	v.OrderTaxFee = ""
	v.EtSerTime = ""
	v.EtShopName = ""
	v.EtVerifiedShopName = ""
	v.EtPlateNumber = ""
	v.O2oSnatchStatus = ""
	v.EticketServiceAddr = ""
	v.EtType = ""
	v.Market = ""
	v.Obs = ""
	v.PaidCouponFee = ""
	v.ShopPick = ""
	v.RxAuditStatus = ""
	v.EsDate = ""
	v.EsRange = ""
	v.OsDate = ""
	v.OsRange = ""
	v.TradeAttr = ""
	v.OmniAttr = ""
	v.OmniParam = ""
	v.Assembly = ""
	v.Identity = ""
	v.O2oStepTradeDetail = ""
	v.O2oStepOrderId = ""
	v.O2oEtOrderId = ""
	v.O2oVoucherPrice = ""
	v.OrderTaxPromotionFee = ""
	v.TidStr = ""
	v.ServiceType = ""
	v.O2oServiceMobile = ""
	v.O2oServiceName = ""
	v.O2oServiceState = ""
	v.O2oServiceCity = ""
	v.O2oServiceDistrict = ""
	v.O2oServiceTown = ""
	v.O2oServiceAddress = ""
	v.O2oStepTradeDetailNew = ""
	v.O2oXiaopiao = ""
	v.O2oContract = ""
	v.RechargeFee = ""
	v.RetailStoreCode = ""
	v.RetailOutOrderId = ""
	v.PlatformSubsidyFee = ""
	v.NrOffline = ""
	v.WttParam = ""
	v.SellerNick = ""
	v.BuyerNick = ""
	v.NrStoreOrderId = ""
	v.NrShopId = ""
	v.NrShopName = ""
	v.NrShopGuideId = ""
	v.NrShopGuideName = ""
	v.NrNoHandle = ""
	v.BizCode = ""
	v.CloudStore = ""
	v.DoneeNick = ""
	v.DoneeOpenUid = ""
	v.SuningShopCode = ""
	v.AllowAppkeys = ""
	v.RetailStoreId = ""
	v.Ua = ""
	v.LinkedmallExtInfo = ""
	v.PayChannel = ""
	v.RtOmniSendType = ""
	v.RtOmniStoreId = ""
	v.RtOmniOuterStoreId = ""
	v.TcpsStart = ""
	v.TcpsCode = ""
	v.TcpsEnd = ""
	v.MTariffFee = ""
	v.TimingPromise = ""
	v.PromiseService = ""
	v.CutoffMinutes = ""
	v.EsTime = ""
	v.DeliveryTime = ""
	v.CollectTime = ""
	v.DispatchTime = ""
	v.SignTime = ""
	v.OuterPartnerMemberId = ""
	v.RootCat = ""
	v.Gifting = ""
	v.GiftingTakeout = ""
	v.OiDate = ""
	v.OiRange = ""
	v.HoldInstall = ""
	v.AppName = ""
	v.EasyHomeCityType = ""
	v.NrDepositOrderId = ""
	v.NrStoreCode = ""
	v.Propoint = ""
	v.ZqsOrderTag = ""
	v.TxpFreezerId = ""
	v.TxpReceiveMethod = ""
	v.ExtendInfo = ""
	v.Lm = ""
	v.BrandLightShopSource = ""
	v.BrandLightShopStoreId = ""
	v.IsWmly = ""
	v.OmniPackage = ""
	v.NczExtAttr = ""
	v.InvoiceDetailPay = ""
	v.InvoiceDetailMidRefund = ""
	v.InvoiceDetailAfterRefund = ""
	v.ExpandCardBasicPrice = ""
	v.ExpandCardExpandPrice = ""
	v.ExpandCardBasicPriceUsed = ""
	v.ExpandCardExpandPriceUsed = ""
	v.DeliveryCps = ""
	v.AsdpBizType = ""
	v.OrderFollowId = ""
	v.AsdpAds = ""
	v.ObTag = ""
	v.DrugRegister = ""
	v.StageAddressType = ""
	v.OgId = ""
	v.PromiseSignTime = ""
	v.OmnichannelParam = ""
	v.Oaid = ""
	v.ScenarioGroup = ""
	v.PlayType = ""
	v.PriorityConsignTime = ""
	v.RealNameAuthStatus = ""
	v.ThirdPartyCustomsDeclaration = ""
	v.ReceiptType = ""
	v.AsyncModified = ""
	v.NutFeature = ""
	v.LgAgingType = ""
	v.LgAging = ""
	v.CnService = ""
	v.Aid = ""
	v.Tid = 0
	v.NumIid = 0
	v.Num = 0
	v.PointFee = 0
	v.RealPointFee = 0
	v.BuyerObtainPointFee = 0
	v.BuyerFlag = 0
	v.SellerFlag = 0
	v.ArriveInterval = 0
	v.ConsignInterval = 0
	v.AlipayPoint = 0
	v.PccAf = 0
	v.TradeExt = nil
	v.EtShopId = 0
	v.CouponFee = 0
	v.TopHold = 0
	v.ForbidConsign = 0
	v.TeamBuyHold = 0
	v.ShareGroupHold = 0
	v.OfpHold = 0
	v.DelayCreateDelivery = 0
	v.Toptype = 0
	v.SortInfo = nil
	v.Sorted = 0
	v.SuningShopValid = 0
	v.ExpandcardInfo = nil
	v.TmallCouponFee = 0
	v.IdentifyInfo = nil
	v.LogisticsAgreement = nil
	v.AlipayId = 0
	v.SellerRate = false
	v.BuyerRate = false
	v.HasPostFee = false
	v.Is3D = false
	v.IsLgtype = false
	v.IsBrandSale = false
	v.IsForceWlb = false
	v.HasYfx = false
	v.CanRate = false
	v.SellerCanRate = false
	v.IsPartConsign = false
	v.IsDaixiao = false
	v.IsWt = false
	v.ZeroPurchase = false
	v.IsShShip = false
	v.PostGateDeclare = false
	v.CrossBondedDeclare = false
	v.IsGift = false
	v.NewPresell = false
	v.YouXiang = false
	v.IsIstore = false
	v.IsOpenmall = false
	v.VLogisticsCreate = false
	v.QRPay = false
	v.GeneralNewPresell = false
	v.IsCycleBuy = false
	v.IsForceDc = false
	v.HasBuyerMessage = false
	v.IsO2oPassport = false
	v.TmallDelivery = false
	v.ThreeplTiming = false
	v.NoShipping = false
	poolTrade.Put(v)
}
