package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加一个商品 API请求
taobao.item.add

此接口用于新增一个淘宝商品  
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费 
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemAddRequest struct {
    model.Params
    // 此参数暂时不起作用
    skuSpecIds   string
    // 此参数暂时不起作用
    skuDeliveryTimes   string
    // 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30
    skuHdLength   string
    // 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30
    skuHdHeight   string
    // 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
    skuHdLampQuantity   string
    // 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
    inputStr   string
    // 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
    inputPids   string
    // 更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
    skuProperties   string
    // Sku的数量串，结构如：num1,num2,num3 如：2,3
    skuQuantities   string
    // Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
    skuPrices   string
    // Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
    skuOuterIds   string
    // sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
    skuBarcode   string
    // 设置是否使用发货时间，商品级别，sku级别
    deliveryTimeNeedDeliveryTime   string
    // 发货时间类型：绝对发货时间或者相对发货时间
    deliveryTimeDeliveryTimeType   string
    // 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
    deliveryTimeDeliveryTime   string
    // 参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    msPaymentReferencePrice   string
    // 尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    msPaymentVoucherPrice   string
    // 订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
    msPaymentPrice   string
    // 预约门店是否支持门店自提,1:是
    localityLifeObs   string
    // 新版电子凭证字段
    localityLifeVersion   string
    // 新版电子凭证包 id
    localityLifePackageid   string
    // 生产许可证号
    foodSecurityPrdLicenseNo   string
    // 产品标准号
    foodSecurityDesignCode   string
    // 厂名
    foodSecurityFactory   string
    // 厂址
    foodSecurityFactorySite   string
    // 厂家联系方式
    foodSecurityContact   string
    // 配料表
    foodSecurityMix   string
    // 储藏方法
    foodSecurityPlanStorage   string
    // 保质期，默认有单位，传入数字
    foodSecurityPeriod   string
    // 食品添加剂
    foodSecurityFoodAdditive   string
    // 供货商
    foodSecuritySupplier   string
    // 生产开始日期，格式必须为yyyy-MM-dd
    foodSecurityProductDateStart   string
    // 生产结束日期,格式必须为yyyy-MM-dd
    foodSecurityProductDateEnd   string
    // 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
    foodSecurityStockDateStart   string
    // 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
    foodSecurityStockDateEnd   string
    // 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
    foodSecurityHealthProductNo   string
    // 所在地省份。如浙江
    locationState   string
    // 所在地城市。如杭州 。
    locationCity   string
    // 商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
    num   int64
    // 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。
    price   float64
    // 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。
    type   string
    // 新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性
    stuffStatus   string
    // 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；
    title   string
    // 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制
    desc   string
    // 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale
    approveStatus   string
    // 叶子类目id
    cid   int64
    // 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值
    props   string
    // 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee）
    freightPayer   string
    // 有效期。可选值:7,14;单位:天;默认值:14
    validThru   int64
    // 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)
    hasInvoice   bool
    // 是否有保修。可选值:true,false;默认值:false(不保修)
    hasWarranty   bool
    // 橱窗推荐。可选值:true,false;默认值:false(不推荐)
    hasShowcase   bool
    // 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
    sellerCids   string
    // 支持会员打折。可选值:true,false;默认值:false(不打折)
    hasDiscount   bool
    // 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写
    postFee   float64
    // 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
    expressFee   float64
    // ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
    emsFee   float64
    // 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)
    listTime   string
    // 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
    increment   float64
    // 商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）
    image   []*model.File
    // 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）
    postageId   int64
    // 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
    auctionPoint   int64
    // 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
    propertyAlias   string
    // 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体
    lang   string
    // 商品外部编码，该字段的最大长度是64个字节
    outerId   string
    // 商品所属的产品ID(B商家发布商品需要用)
    productId   int64
    // （推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
    picPath   string
    // 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
    autoFill   string
    // 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
    isTaobao   bool
    // 是否在外店显示
    isEx   bool
    // 是否是3D
    is3D   bool
    // 是否承诺退换货服务!虚拟商品无须设置此项!
    sellPromise   bool
    // 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃
    codPostageId   int64
    // 实物闪电发货
    isLightningConsignment   bool
    // 商品的重量(商超卖家专用字段)
    weight   int64
    // 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。
    isXinpin   bool
    // 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存
    subStock   int64
    // 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
    scenicTicketPayWay   int64
    // 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
    scenicTicketBookCost   string
    // 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）
    itemSize   string
    // 商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。
    itemWeight   string
    // 商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。
    sellPoint   string
    // 商品条形码
    barcode   string
    // 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；
    newprepay   string
    // 商品资质信息
    qualification   string
    // 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
    o2oBindService   bool
    // 宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
    features   string
    // 忽略警告提示.
    ignorewarning   string
    // 售后说明模板id
    afterSaleId   int64
    // 基础色数据，淘宝不使用
    changeProp   string
    // 已废弃
    descModules   string
    // 是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
    isOffline   string
    // 无线的宝贝描述
    wirelessDesc   string
    // 手机类目spu 优化，信息确认字段
    spuConfirm   bool
    // 主图视频id
    videoId   int64
    // 主图视频互动信息id，必须填写主图视频id才能有互动信息id
    interactiveId   int64
    // 租赁扩展信息
    leaseExtendsInfo   string
    // 仅淘小铺卖家需要。佣金比例(15.3对应的佣金比例为15.3%).只支持小数点后1位。多余的位数四舍五入(15.32会保存为15.3%
    brokerage   string
    // 业务身份编码。淘小铺编码为"taobao-taoxiaopu"
    bizCode   string
    // 此字段已经废弃，不再使用
    imageUrls   []string
    // 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
    localityLifeChooseLogis   string
    // 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
    localityLifeExpirydate   string
    // 网点ID
    localityLifeNetworkId   string
    // 码商信息，格式为 码商id:nick
    localityLifeMerchant   string
    // 核销打款 1代表核销打款 0代表非核销打款
    localityLifeVerification   string
    // 退款比例，百分比%前的数字,1-100的正整数值
    localityLifeRefundRatio   int64
    // 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
    localityLifeOnsaleAutoRefundRatio   int64
    // 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
    localityLifeRefundmafee   string
    // 电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
    localityLifeEticket   string
    // 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
    paimaiInfoMode   int64
    // 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。
    paimaiInfoDeposit   int64
    // 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
    paimaiInfoInterval   int64
    // 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
    paimaiInfoReserve   float64
    // 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    paimaiInfoValidHour   int64
    // 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    paimaiInfoValidMinute   int64
    // 全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
    globalStockType   string
    // 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值请填写法定的国家名称，类如（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾），不要使用其他
    globalStockCountry   string
    // 是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改
    supportCustomMade   bool
    // 定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1
    customMadeTypeId   string
    // 全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”
    globalStockDeliveryPlace   string
    // 全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约
    globalStockTaxFreePromise   bool
    // 针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上
    inputCustomCpv   string
    // 针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
    cpvMemo   string
}

// 初始化TaobaoItemAddRequest对象
func NewTaobaoItemAddRequest() *TaobaoItemAddRequest{
    return &TaobaoItemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemAddRequest) GetApiMethodName() string {
    return "taobao.item.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuSpecIds Setter
// 此参数暂时不起作用
func (r *TaobaoItemAddRequest) SetSkuSpecIds(skuSpecIds string) error {
    r.skuSpecIds = skuSpecIds
    r.Set("sku_spec_ids", skuSpecIds)
    return nil
}

// SkuSpecIds Getter
func (r TaobaoItemAddRequest) GetSkuSpecIds() string {
    return r.skuSpecIds
}
// SkuDeliveryTimes Setter
// 此参数暂时不起作用
func (r *TaobaoItemAddRequest) SetSkuDeliveryTimes(skuDeliveryTimes string) error {
    r.skuDeliveryTimes = skuDeliveryTimes
    r.Set("sku_delivery_times", skuDeliveryTimes)
    return nil
}

// SkuDeliveryTimes Getter
func (r TaobaoItemAddRequest) GetSkuDeliveryTimes() string {
    return r.skuDeliveryTimes
}
// SkuHdLength Setter
// 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30
func (r *TaobaoItemAddRequest) SetSkuHdLength(skuHdLength string) error {
    r.skuHdLength = skuHdLength
    r.Set("sku_hd_length", skuHdLength)
    return nil
}

// SkuHdLength Getter
func (r TaobaoItemAddRequest) GetSkuHdLength() string {
    return r.skuHdLength
}
// SkuHdHeight Setter
// 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30
func (r *TaobaoItemAddRequest) SetSkuHdHeight(skuHdHeight string) error {
    r.skuHdHeight = skuHdHeight
    r.Set("sku_hd_height", skuHdHeight)
    return nil
}

// SkuHdHeight Getter
func (r TaobaoItemAddRequest) GetSkuHdHeight() string {
    return r.skuHdHeight
}
// SkuHdLampQuantity Setter
// 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
func (r *TaobaoItemAddRequest) SetSkuHdLampQuantity(skuHdLampQuantity string) error {
    r.skuHdLampQuantity = skuHdLampQuantity
    r.Set("sku_hd_lamp_quantity", skuHdLampQuantity)
    return nil
}

// SkuHdLampQuantity Getter
func (r TaobaoItemAddRequest) GetSkuHdLampQuantity() string {
    return r.skuHdLampQuantity
}
// InputStr Setter
// 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
func (r *TaobaoItemAddRequest) SetInputStr(inputStr string) error {
    r.inputStr = inputStr
    r.Set("input_str", inputStr)
    return nil
}

// InputStr Getter
func (r TaobaoItemAddRequest) GetInputStr() string {
    return r.inputStr
}
// InputPids Setter
// 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
func (r *TaobaoItemAddRequest) SetInputPids(inputPids string) error {
    r.inputPids = inputPids
    r.Set("input_pids", inputPids)
    return nil
}

// InputPids Getter
func (r TaobaoItemAddRequest) GetInputPids() string {
    return r.inputPids
}
// SkuProperties Setter
// 更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
func (r *TaobaoItemAddRequest) SetSkuProperties(skuProperties string) error {
    r.skuProperties = skuProperties
    r.Set("sku_properties", skuProperties)
    return nil
}

// SkuProperties Getter
func (r TaobaoItemAddRequest) GetSkuProperties() string {
    return r.skuProperties
}
// SkuQuantities Setter
// Sku的数量串，结构如：num1,num2,num3 如：2,3
func (r *TaobaoItemAddRequest) SetSkuQuantities(skuQuantities string) error {
    r.skuQuantities = skuQuantities
    r.Set("sku_quantities", skuQuantities)
    return nil
}

// SkuQuantities Getter
func (r TaobaoItemAddRequest) GetSkuQuantities() string {
    return r.skuQuantities
}
// SkuPrices Setter
// Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemAddRequest) SetSkuPrices(skuPrices string) error {
    r.skuPrices = skuPrices
    r.Set("sku_prices", skuPrices)
    return nil
}

// SkuPrices Getter
func (r TaobaoItemAddRequest) GetSkuPrices() string {
    return r.skuPrices
}
// SkuOuterIds Setter
// Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
func (r *TaobaoItemAddRequest) SetSkuOuterIds(skuOuterIds string) error {
    r.skuOuterIds = skuOuterIds
    r.Set("sku_outer_ids", skuOuterIds)
    return nil
}

// SkuOuterIds Getter
func (r TaobaoItemAddRequest) GetSkuOuterIds() string {
    return r.skuOuterIds
}
// SkuBarcode Setter
// sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
func (r *TaobaoItemAddRequest) SetSkuBarcode(skuBarcode string) error {
    r.skuBarcode = skuBarcode
    r.Set("sku_barcode", skuBarcode)
    return nil
}

// SkuBarcode Getter
func (r TaobaoItemAddRequest) GetSkuBarcode() string {
    return r.skuBarcode
}
// DeliveryTimeNeedDeliveryTime Setter
// 设置是否使用发货时间，商品级别，sku级别
func (r *TaobaoItemAddRequest) SetDeliveryTimeNeedDeliveryTime(deliveryTimeNeedDeliveryTime string) error {
    r.deliveryTimeNeedDeliveryTime = deliveryTimeNeedDeliveryTime
    r.Set("delivery_time.need_delivery_time", deliveryTimeNeedDeliveryTime)
    return nil
}

// DeliveryTimeNeedDeliveryTime Getter
func (r TaobaoItemAddRequest) GetDeliveryTimeNeedDeliveryTime() string {
    return r.deliveryTimeNeedDeliveryTime
}
// DeliveryTimeDeliveryTimeType Setter
// 发货时间类型：绝对发货时间或者相对发货时间
func (r *TaobaoItemAddRequest) SetDeliveryTimeDeliveryTimeType(deliveryTimeDeliveryTimeType string) error {
    r.deliveryTimeDeliveryTimeType = deliveryTimeDeliveryTimeType
    r.Set("delivery_time.delivery_time_type", deliveryTimeDeliveryTimeType)
    return nil
}

// DeliveryTimeDeliveryTimeType Getter
func (r TaobaoItemAddRequest) GetDeliveryTimeDeliveryTimeType() string {
    return r.deliveryTimeDeliveryTimeType
}
// DeliveryTimeDeliveryTime Setter
// 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
func (r *TaobaoItemAddRequest) SetDeliveryTimeDeliveryTime(deliveryTimeDeliveryTime string) error {
    r.deliveryTimeDeliveryTime = deliveryTimeDeliveryTime
    r.Set("delivery_time.delivery_time", deliveryTimeDeliveryTime)
    return nil
}

// DeliveryTimeDeliveryTime Getter
func (r TaobaoItemAddRequest) GetDeliveryTimeDeliveryTime() string {
    return r.deliveryTimeDeliveryTime
}
// MsPaymentReferencePrice Setter
// 参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddRequest) SetMsPaymentReferencePrice(msPaymentReferencePrice string) error {
    r.msPaymentReferencePrice = msPaymentReferencePrice
    r.Set("ms_payment.reference_price", msPaymentReferencePrice)
    return nil
}

// MsPaymentReferencePrice Getter
func (r TaobaoItemAddRequest) GetMsPaymentReferencePrice() string {
    return r.msPaymentReferencePrice
}
// MsPaymentVoucherPrice Setter
// 尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddRequest) SetMsPaymentVoucherPrice(msPaymentVoucherPrice string) error {
    r.msPaymentVoucherPrice = msPaymentVoucherPrice
    r.Set("ms_payment.voucher_price", msPaymentVoucherPrice)
    return nil
}

// MsPaymentVoucherPrice Getter
func (r TaobaoItemAddRequest) GetMsPaymentVoucherPrice() string {
    return r.msPaymentVoucherPrice
}
// MsPaymentPrice Setter
// 订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddRequest) SetMsPaymentPrice(msPaymentPrice string) error {
    r.msPaymentPrice = msPaymentPrice
    r.Set("ms_payment.price", msPaymentPrice)
    return nil
}

// MsPaymentPrice Getter
func (r TaobaoItemAddRequest) GetMsPaymentPrice() string {
    return r.msPaymentPrice
}
// LocalityLifeObs Setter
// 预约门店是否支持门店自提,1:是
func (r *TaobaoItemAddRequest) SetLocalityLifeObs(localityLifeObs string) error {
    r.localityLifeObs = localityLifeObs
    r.Set("locality_life.obs", localityLifeObs)
    return nil
}

// LocalityLifeObs Getter
func (r TaobaoItemAddRequest) GetLocalityLifeObs() string {
    return r.localityLifeObs
}
// LocalityLifeVersion Setter
// 新版电子凭证字段
func (r *TaobaoItemAddRequest) SetLocalityLifeVersion(localityLifeVersion string) error {
    r.localityLifeVersion = localityLifeVersion
    r.Set("locality_life.version", localityLifeVersion)
    return nil
}

// LocalityLifeVersion Getter
func (r TaobaoItemAddRequest) GetLocalityLifeVersion() string {
    return r.localityLifeVersion
}
// LocalityLifePackageid Setter
// 新版电子凭证包 id
func (r *TaobaoItemAddRequest) SetLocalityLifePackageid(localityLifePackageid string) error {
    r.localityLifePackageid = localityLifePackageid
    r.Set("locality_life.packageid", localityLifePackageid)
    return nil
}

// LocalityLifePackageid Getter
func (r TaobaoItemAddRequest) GetLocalityLifePackageid() string {
    return r.localityLifePackageid
}
// FoodSecurityPrdLicenseNo Setter
// 生产许可证号
func (r *TaobaoItemAddRequest) SetFoodSecurityPrdLicenseNo(foodSecurityPrdLicenseNo string) error {
    r.foodSecurityPrdLicenseNo = foodSecurityPrdLicenseNo
    r.Set("food_security.prd_license_no", foodSecurityPrdLicenseNo)
    return nil
}

// FoodSecurityPrdLicenseNo Getter
func (r TaobaoItemAddRequest) GetFoodSecurityPrdLicenseNo() string {
    return r.foodSecurityPrdLicenseNo
}
// FoodSecurityDesignCode Setter
// 产品标准号
func (r *TaobaoItemAddRequest) SetFoodSecurityDesignCode(foodSecurityDesignCode string) error {
    r.foodSecurityDesignCode = foodSecurityDesignCode
    r.Set("food_security.design_code", foodSecurityDesignCode)
    return nil
}

// FoodSecurityDesignCode Getter
func (r TaobaoItemAddRequest) GetFoodSecurityDesignCode() string {
    return r.foodSecurityDesignCode
}
// FoodSecurityFactory Setter
// 厂名
func (r *TaobaoItemAddRequest) SetFoodSecurityFactory(foodSecurityFactory string) error {
    r.foodSecurityFactory = foodSecurityFactory
    r.Set("food_security.factory", foodSecurityFactory)
    return nil
}

// FoodSecurityFactory Getter
func (r TaobaoItemAddRequest) GetFoodSecurityFactory() string {
    return r.foodSecurityFactory
}
// FoodSecurityFactorySite Setter
// 厂址
func (r *TaobaoItemAddRequest) SetFoodSecurityFactorySite(foodSecurityFactorySite string) error {
    r.foodSecurityFactorySite = foodSecurityFactorySite
    r.Set("food_security.factory_site", foodSecurityFactorySite)
    return nil
}

// FoodSecurityFactorySite Getter
func (r TaobaoItemAddRequest) GetFoodSecurityFactorySite() string {
    return r.foodSecurityFactorySite
}
// FoodSecurityContact Setter
// 厂家联系方式
func (r *TaobaoItemAddRequest) SetFoodSecurityContact(foodSecurityContact string) error {
    r.foodSecurityContact = foodSecurityContact
    r.Set("food_security.contact", foodSecurityContact)
    return nil
}

// FoodSecurityContact Getter
func (r TaobaoItemAddRequest) GetFoodSecurityContact() string {
    return r.foodSecurityContact
}
// FoodSecurityMix Setter
// 配料表
func (r *TaobaoItemAddRequest) SetFoodSecurityMix(foodSecurityMix string) error {
    r.foodSecurityMix = foodSecurityMix
    r.Set("food_security.mix", foodSecurityMix)
    return nil
}

// FoodSecurityMix Getter
func (r TaobaoItemAddRequest) GetFoodSecurityMix() string {
    return r.foodSecurityMix
}
// FoodSecurityPlanStorage Setter
// 储藏方法
func (r *TaobaoItemAddRequest) SetFoodSecurityPlanStorage(foodSecurityPlanStorage string) error {
    r.foodSecurityPlanStorage = foodSecurityPlanStorage
    r.Set("food_security.plan_storage", foodSecurityPlanStorage)
    return nil
}

// FoodSecurityPlanStorage Getter
func (r TaobaoItemAddRequest) GetFoodSecurityPlanStorage() string {
    return r.foodSecurityPlanStorage
}
// FoodSecurityPeriod Setter
// 保质期，默认有单位，传入数字
func (r *TaobaoItemAddRequest) SetFoodSecurityPeriod(foodSecurityPeriod string) error {
    r.foodSecurityPeriod = foodSecurityPeriod
    r.Set("food_security.period", foodSecurityPeriod)
    return nil
}

// FoodSecurityPeriod Getter
func (r TaobaoItemAddRequest) GetFoodSecurityPeriod() string {
    return r.foodSecurityPeriod
}
// FoodSecurityFoodAdditive Setter
// 食品添加剂
func (r *TaobaoItemAddRequest) SetFoodSecurityFoodAdditive(foodSecurityFoodAdditive string) error {
    r.foodSecurityFoodAdditive = foodSecurityFoodAdditive
    r.Set("food_security.food_additive", foodSecurityFoodAdditive)
    return nil
}

// FoodSecurityFoodAdditive Getter
func (r TaobaoItemAddRequest) GetFoodSecurityFoodAdditive() string {
    return r.foodSecurityFoodAdditive
}
// FoodSecuritySupplier Setter
// 供货商
func (r *TaobaoItemAddRequest) SetFoodSecuritySupplier(foodSecuritySupplier string) error {
    r.foodSecuritySupplier = foodSecuritySupplier
    r.Set("food_security.supplier", foodSecuritySupplier)
    return nil
}

// FoodSecuritySupplier Getter
func (r TaobaoItemAddRequest) GetFoodSecuritySupplier() string {
    return r.foodSecuritySupplier
}
// FoodSecurityProductDateStart Setter
// 生产开始日期，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddRequest) SetFoodSecurityProductDateStart(foodSecurityProductDateStart string) error {
    r.foodSecurityProductDateStart = foodSecurityProductDateStart
    r.Set("food_security.product_date_start", foodSecurityProductDateStart)
    return nil
}

// FoodSecurityProductDateStart Getter
func (r TaobaoItemAddRequest) GetFoodSecurityProductDateStart() string {
    return r.foodSecurityProductDateStart
}
// FoodSecurityProductDateEnd Setter
// 生产结束日期,格式必须为yyyy-MM-dd
func (r *TaobaoItemAddRequest) SetFoodSecurityProductDateEnd(foodSecurityProductDateEnd string) error {
    r.foodSecurityProductDateEnd = foodSecurityProductDateEnd
    r.Set("food_security.product_date_end", foodSecurityProductDateEnd)
    return nil
}

// FoodSecurityProductDateEnd Getter
func (r TaobaoItemAddRequest) GetFoodSecurityProductDateEnd() string {
    return r.foodSecurityProductDateEnd
}
// FoodSecurityStockDateStart Setter
// 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddRequest) SetFoodSecurityStockDateStart(foodSecurityStockDateStart string) error {
    r.foodSecurityStockDateStart = foodSecurityStockDateStart
    r.Set("food_security.stock_date_start", foodSecurityStockDateStart)
    return nil
}

// FoodSecurityStockDateStart Getter
func (r TaobaoItemAddRequest) GetFoodSecurityStockDateStart() string {
    return r.foodSecurityStockDateStart
}
// FoodSecurityStockDateEnd Setter
// 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddRequest) SetFoodSecurityStockDateEnd(foodSecurityStockDateEnd string) error {
    r.foodSecurityStockDateEnd = foodSecurityStockDateEnd
    r.Set("food_security.stock_date_end", foodSecurityStockDateEnd)
    return nil
}

// FoodSecurityStockDateEnd Getter
func (r TaobaoItemAddRequest) GetFoodSecurityStockDateEnd() string {
    return r.foodSecurityStockDateEnd
}
// FoodSecurityHealthProductNo Setter
// 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
func (r *TaobaoItemAddRequest) SetFoodSecurityHealthProductNo(foodSecurityHealthProductNo string) error {
    r.foodSecurityHealthProductNo = foodSecurityHealthProductNo
    r.Set("food_security.health_product_no", foodSecurityHealthProductNo)
    return nil
}

// FoodSecurityHealthProductNo Getter
func (r TaobaoItemAddRequest) GetFoodSecurityHealthProductNo() string {
    return r.foodSecurityHealthProductNo
}
// LocationState Setter
// 所在地省份。如浙江
func (r *TaobaoItemAddRequest) SetLocationState(locationState string) error {
    r.locationState = locationState
    r.Set("location.state", locationState)
    return nil
}

// LocationState Getter
func (r TaobaoItemAddRequest) GetLocationState() string {
    return r.locationState
}
// LocationCity Setter
// 所在地城市。如杭州 。
func (r *TaobaoItemAddRequest) SetLocationCity(locationCity string) error {
    r.locationCity = locationCity
    r.Set("location.city", locationCity)
    return nil
}

// LocationCity Getter
func (r TaobaoItemAddRequest) GetLocationCity() string {
    return r.locationCity
}
// Num Setter
// 商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
func (r *TaobaoItemAddRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r TaobaoItemAddRequest) GetNum() int64 {
    return r.num
}
// Price Setter
// 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。
func (r *TaobaoItemAddRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoItemAddRequest) GetPrice() float64 {
    return r.price
}
// Type Setter
// 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。
func (r *TaobaoItemAddRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoItemAddRequest) GetType() string {
    return r.type
}
// StuffStatus Setter
// 新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性
func (r *TaobaoItemAddRequest) SetStuffStatus(stuffStatus string) error {
    r.stuffStatus = stuffStatus
    r.Set("stuff_status", stuffStatus)
    return nil
}

// StuffStatus Getter
func (r TaobaoItemAddRequest) GetStuffStatus() string {
    return r.stuffStatus
}
// Title Setter
// 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；
func (r *TaobaoItemAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoItemAddRequest) GetTitle() string {
    return r.title
}
// Desc Setter
// 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制
func (r *TaobaoItemAddRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoItemAddRequest) GetDesc() string {
    return r.desc
}
// ApproveStatus Setter
// 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale
func (r *TaobaoItemAddRequest) SetApproveStatus(approveStatus string) error {
    r.approveStatus = approveStatus
    r.Set("approve_status", approveStatus)
    return nil
}

// ApproveStatus Getter
func (r TaobaoItemAddRequest) GetApproveStatus() string {
    return r.approveStatus
}
// Cid Setter
// 叶子类目id
func (r *TaobaoItemAddRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoItemAddRequest) GetCid() int64 {
    return r.cid
}
// Props Setter
// 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值
func (r *TaobaoItemAddRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

// Props Getter
func (r TaobaoItemAddRequest) GetProps() string {
    return r.props
}
// FreightPayer Setter
// 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee）
func (r *TaobaoItemAddRequest) SetFreightPayer(freightPayer string) error {
    r.freightPayer = freightPayer
    r.Set("freight_payer", freightPayer)
    return nil
}

// FreightPayer Getter
func (r TaobaoItemAddRequest) GetFreightPayer() string {
    return r.freightPayer
}
// ValidThru Setter
// 有效期。可选值:7,14;单位:天;默认值:14
func (r *TaobaoItemAddRequest) SetValidThru(validThru int64) error {
    r.validThru = validThru
    r.Set("valid_thru", validThru)
    return nil
}

// ValidThru Getter
func (r TaobaoItemAddRequest) GetValidThru() int64 {
    return r.validThru
}
// HasInvoice Setter
// 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)
func (r *TaobaoItemAddRequest) SetHasInvoice(hasInvoice bool) error {
    r.hasInvoice = hasInvoice
    r.Set("has_invoice", hasInvoice)
    return nil
}

// HasInvoice Getter
func (r TaobaoItemAddRequest) GetHasInvoice() bool {
    return r.hasInvoice
}
// HasWarranty Setter
// 是否有保修。可选值:true,false;默认值:false(不保修)
func (r *TaobaoItemAddRequest) SetHasWarranty(hasWarranty bool) error {
    r.hasWarranty = hasWarranty
    r.Set("has_warranty", hasWarranty)
    return nil
}

// HasWarranty Getter
func (r TaobaoItemAddRequest) GetHasWarranty() bool {
    return r.hasWarranty
}
// HasShowcase Setter
// 橱窗推荐。可选值:true,false;默认值:false(不推荐)
func (r *TaobaoItemAddRequest) SetHasShowcase(hasShowcase bool) error {
    r.hasShowcase = hasShowcase
    r.Set("has_showcase", hasShowcase)
    return nil
}

// HasShowcase Getter
func (r TaobaoItemAddRequest) GetHasShowcase() bool {
    return r.hasShowcase
}
// SellerCids Setter
// 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
func (r *TaobaoItemAddRequest) SetSellerCids(sellerCids string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

// SellerCids Getter
func (r TaobaoItemAddRequest) GetSellerCids() string {
    return r.sellerCids
}
// HasDiscount Setter
// 支持会员打折。可选值:true,false;默认值:false(不打折)
func (r *TaobaoItemAddRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

// HasDiscount Getter
func (r TaobaoItemAddRequest) GetHasDiscount() bool {
    return r.hasDiscount
}
// PostFee Setter
// 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写
func (r *TaobaoItemAddRequest) SetPostFee(postFee float64) error {
    r.postFee = postFee
    r.Set("post_fee", postFee)
    return nil
}

// PostFee Getter
func (r TaobaoItemAddRequest) GetPostFee() float64 {
    return r.postFee
}
// ExpressFee Setter
// 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
func (r *TaobaoItemAddRequest) SetExpressFee(expressFee float64) error {
    r.expressFee = expressFee
    r.Set("express_fee", expressFee)
    return nil
}

// ExpressFee Getter
func (r TaobaoItemAddRequest) GetExpressFee() float64 {
    return r.expressFee
}
// EmsFee Setter
// ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
func (r *TaobaoItemAddRequest) SetEmsFee(emsFee float64) error {
    r.emsFee = emsFee
    r.Set("ems_fee", emsFee)
    return nil
}

// EmsFee Getter
func (r TaobaoItemAddRequest) GetEmsFee() float64 {
    return r.emsFee
}
// ListTime Setter
// 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)
func (r *TaobaoItemAddRequest) SetListTime(listTime string) error {
    r.listTime = listTime
    r.Set("list_time", listTime)
    return nil
}

// ListTime Getter
func (r TaobaoItemAddRequest) GetListTime() string {
    return r.listTime
}
// Increment Setter
// 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
func (r *TaobaoItemAddRequest) SetIncrement(increment float64) error {
    r.increment = increment
    r.Set("increment", increment)
    return nil
}

// Increment Getter
func (r TaobaoItemAddRequest) GetIncrement() float64 {
    return r.increment
}
// Image Setter
// 商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）
func (r *TaobaoItemAddRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoItemAddRequest) GetImage() []*model.File {
    return r.image
}
// PostageId Setter
// 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）
func (r *TaobaoItemAddRequest) SetPostageId(postageId int64) error {
    r.postageId = postageId
    r.Set("postage_id", postageId)
    return nil
}

// PostageId Getter
func (r TaobaoItemAddRequest) GetPostageId() int64 {
    return r.postageId
}
// AuctionPoint Setter
// 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
func (r *TaobaoItemAddRequest) SetAuctionPoint(auctionPoint int64) error {
    r.auctionPoint = auctionPoint
    r.Set("auction_point", auctionPoint)
    return nil
}

// AuctionPoint Getter
func (r TaobaoItemAddRequest) GetAuctionPoint() int64 {
    return r.auctionPoint
}
// PropertyAlias Setter
// 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
func (r *TaobaoItemAddRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

// PropertyAlias Getter
func (r TaobaoItemAddRequest) GetPropertyAlias() string {
    return r.propertyAlias
}
// Lang Setter
// 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体
func (r *TaobaoItemAddRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

// Lang Getter
func (r TaobaoItemAddRequest) GetLang() string {
    return r.lang
}
// OuterId Setter
// 商品外部编码，该字段的最大长度是64个字节
func (r *TaobaoItemAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemAddRequest) GetOuterId() string {
    return r.outerId
}
// ProductId Setter
// 商品所属的产品ID(B商家发布商品需要用)
func (r *TaobaoItemAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoItemAddRequest) GetProductId() int64 {
    return r.productId
}
// PicPath Setter
// （推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
func (r *TaobaoItemAddRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

// PicPath Getter
func (r TaobaoItemAddRequest) GetPicPath() string {
    return r.picPath
}
// AutoFill Setter
// 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
func (r *TaobaoItemAddRequest) SetAutoFill(autoFill string) error {
    r.autoFill = autoFill
    r.Set("auto_fill", autoFill)
    return nil
}

// AutoFill Getter
func (r TaobaoItemAddRequest) GetAutoFill() string {
    return r.autoFill
}
// IsTaobao Setter
// 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
func (r *TaobaoItemAddRequest) SetIsTaobao(isTaobao bool) error {
    r.isTaobao = isTaobao
    r.Set("is_taobao", isTaobao)
    return nil
}

// IsTaobao Getter
func (r TaobaoItemAddRequest) GetIsTaobao() bool {
    return r.isTaobao
}
// IsEx Setter
// 是否在外店显示
func (r *TaobaoItemAddRequest) SetIsEx(isEx bool) error {
    r.isEx = isEx
    r.Set("is_ex", isEx)
    return nil
}

// IsEx Getter
func (r TaobaoItemAddRequest) GetIsEx() bool {
    return r.isEx
}
// Is3D Setter
// 是否是3D
func (r *TaobaoItemAddRequest) SetIs3D(is3D bool) error {
    r.is3D = is3D
    r.Set("is_3D", is3D)
    return nil
}

// Is3D Getter
func (r TaobaoItemAddRequest) GetIs3D() bool {
    return r.is3D
}
// SellPromise Setter
// 是否承诺退换货服务!虚拟商品无须设置此项!
func (r *TaobaoItemAddRequest) SetSellPromise(sellPromise bool) error {
    r.sellPromise = sellPromise
    r.Set("sell_promise", sellPromise)
    return nil
}

// SellPromise Getter
func (r TaobaoItemAddRequest) GetSellPromise() bool {
    return r.sellPromise
}
// CodPostageId Setter
// 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃
func (r *TaobaoItemAddRequest) SetCodPostageId(codPostageId int64) error {
    r.codPostageId = codPostageId
    r.Set("cod_postage_id", codPostageId)
    return nil
}

// CodPostageId Getter
func (r TaobaoItemAddRequest) GetCodPostageId() int64 {
    return r.codPostageId
}
// IsLightningConsignment Setter
// 实物闪电发货
func (r *TaobaoItemAddRequest) SetIsLightningConsignment(isLightningConsignment bool) error {
    r.isLightningConsignment = isLightningConsignment
    r.Set("is_lightning_consignment", isLightningConsignment)
    return nil
}

// IsLightningConsignment Getter
func (r TaobaoItemAddRequest) GetIsLightningConsignment() bool {
    return r.isLightningConsignment
}
// Weight Setter
// 商品的重量(商超卖家专用字段)
func (r *TaobaoItemAddRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r TaobaoItemAddRequest) GetWeight() int64 {
    return r.weight
}
// IsXinpin Setter
// 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。
func (r *TaobaoItemAddRequest) SetIsXinpin(isXinpin bool) error {
    r.isXinpin = isXinpin
    r.Set("is_xinpin", isXinpin)
    return nil
}

// IsXinpin Getter
func (r TaobaoItemAddRequest) GetIsXinpin() bool {
    return r.isXinpin
}
// SubStock Setter
// 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存
func (r *TaobaoItemAddRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

// SubStock Getter
func (r TaobaoItemAddRequest) GetSubStock() int64 {
    return r.subStock
}
// ScenicTicketPayWay Setter
// 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
func (r *TaobaoItemAddRequest) SetScenicTicketPayWay(scenicTicketPayWay int64) error {
    r.scenicTicketPayWay = scenicTicketPayWay
    r.Set("scenic_ticket_pay_way", scenicTicketPayWay)
    return nil
}

// ScenicTicketPayWay Getter
func (r TaobaoItemAddRequest) GetScenicTicketPayWay() int64 {
    return r.scenicTicketPayWay
}
// ScenicTicketBookCost Setter
// 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
func (r *TaobaoItemAddRequest) SetScenicTicketBookCost(scenicTicketBookCost string) error {
    r.scenicTicketBookCost = scenicTicketBookCost
    r.Set("scenic_ticket_book_cost", scenicTicketBookCost)
    return nil
}

// ScenicTicketBookCost Getter
func (r TaobaoItemAddRequest) GetScenicTicketBookCost() string {
    return r.scenicTicketBookCost
}
// ItemSize Setter
// 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）
func (r *TaobaoItemAddRequest) SetItemSize(itemSize string) error {
    r.itemSize = itemSize
    r.Set("item_size", itemSize)
    return nil
}

// ItemSize Getter
func (r TaobaoItemAddRequest) GetItemSize() string {
    return r.itemSize
}
// ItemWeight Setter
// 商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。
func (r *TaobaoItemAddRequest) SetItemWeight(itemWeight string) error {
    r.itemWeight = itemWeight
    r.Set("item_weight", itemWeight)
    return nil
}

// ItemWeight Getter
func (r TaobaoItemAddRequest) GetItemWeight() string {
    return r.itemWeight
}
// SellPoint Setter
// 商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。
func (r *TaobaoItemAddRequest) SetSellPoint(sellPoint string) error {
    r.sellPoint = sellPoint
    r.Set("sell_point", sellPoint)
    return nil
}

// SellPoint Getter
func (r TaobaoItemAddRequest) GetSellPoint() string {
    return r.sellPoint
}
// Barcode Setter
// 商品条形码
func (r *TaobaoItemAddRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r TaobaoItemAddRequest) GetBarcode() string {
    return r.barcode
}
// Newprepay Setter
// 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；
func (r *TaobaoItemAddRequest) SetNewprepay(newprepay string) error {
    r.newprepay = newprepay
    r.Set("newprepay", newprepay)
    return nil
}

// Newprepay Getter
func (r TaobaoItemAddRequest) GetNewprepay() string {
    return r.newprepay
}
// Qualification Setter
// 商品资质信息
func (r *TaobaoItemAddRequest) SetQualification(qualification string) error {
    r.qualification = qualification
    r.Set("qualification", qualification)
    return nil
}

// Qualification Getter
func (r TaobaoItemAddRequest) GetQualification() string {
    return r.qualification
}
// O2oBindService Setter
// 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
func (r *TaobaoItemAddRequest) SetO2oBindService(o2oBindService bool) error {
    r.o2oBindService = o2oBindService
    r.Set("o2o_bind_service", o2oBindService)
    return nil
}

// O2oBindService Getter
func (r TaobaoItemAddRequest) GetO2oBindService() bool {
    return r.o2oBindService
}
// Features Setter
// 宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
func (r *TaobaoItemAddRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

// Features Getter
func (r TaobaoItemAddRequest) GetFeatures() string {
    return r.features
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemAddRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemAddRequest) GetIgnorewarning() string {
    return r.ignorewarning
}
// AfterSaleId Setter
// 售后说明模板id
func (r *TaobaoItemAddRequest) SetAfterSaleId(afterSaleId int64) error {
    r.afterSaleId = afterSaleId
    r.Set("after_sale_id", afterSaleId)
    return nil
}

// AfterSaleId Getter
func (r TaobaoItemAddRequest) GetAfterSaleId() int64 {
    return r.afterSaleId
}
// ChangeProp Setter
// 基础色数据，淘宝不使用
func (r *TaobaoItemAddRequest) SetChangeProp(changeProp string) error {
    r.changeProp = changeProp
    r.Set("change_prop", changeProp)
    return nil
}

// ChangeProp Getter
func (r TaobaoItemAddRequest) GetChangeProp() string {
    return r.changeProp
}
// DescModules Setter
// 已废弃
func (r *TaobaoItemAddRequest) SetDescModules(descModules string) error {
    r.descModules = descModules
    r.Set("desc_modules", descModules)
    return nil
}

// DescModules Getter
func (r TaobaoItemAddRequest) GetDescModules() string {
    return r.descModules
}
// IsOffline Setter
// 是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
func (r *TaobaoItemAddRequest) SetIsOffline(isOffline string) error {
    r.isOffline = isOffline
    r.Set("is_offline", isOffline)
    return nil
}

// IsOffline Getter
func (r TaobaoItemAddRequest) GetIsOffline() string {
    return r.isOffline
}
// WirelessDesc Setter
// 无线的宝贝描述
func (r *TaobaoItemAddRequest) SetWirelessDesc(wirelessDesc string) error {
    r.wirelessDesc = wirelessDesc
    r.Set("wireless_desc", wirelessDesc)
    return nil
}

// WirelessDesc Getter
func (r TaobaoItemAddRequest) GetWirelessDesc() string {
    return r.wirelessDesc
}
// SpuConfirm Setter
// 手机类目spu 优化，信息确认字段
func (r *TaobaoItemAddRequest) SetSpuConfirm(spuConfirm bool) error {
    r.spuConfirm = spuConfirm
    r.Set("spu_confirm", spuConfirm)
    return nil
}

// SpuConfirm Getter
func (r TaobaoItemAddRequest) GetSpuConfirm() bool {
    return r.spuConfirm
}
// VideoId Setter
// 主图视频id
func (r *TaobaoItemAddRequest) SetVideoId(videoId int64) error {
    r.videoId = videoId
    r.Set("video_id", videoId)
    return nil
}

// VideoId Getter
func (r TaobaoItemAddRequest) GetVideoId() int64 {
    return r.videoId
}
// InteractiveId Setter
// 主图视频互动信息id，必须填写主图视频id才能有互动信息id
func (r *TaobaoItemAddRequest) SetInteractiveId(interactiveId int64) error {
    r.interactiveId = interactiveId
    r.Set("interactive_id", interactiveId)
    return nil
}

// InteractiveId Getter
func (r TaobaoItemAddRequest) GetInteractiveId() int64 {
    return r.interactiveId
}
// LeaseExtendsInfo Setter
// 租赁扩展信息
func (r *TaobaoItemAddRequest) SetLeaseExtendsInfo(leaseExtendsInfo string) error {
    r.leaseExtendsInfo = leaseExtendsInfo
    r.Set("lease_extends_info", leaseExtendsInfo)
    return nil
}

// LeaseExtendsInfo Getter
func (r TaobaoItemAddRequest) GetLeaseExtendsInfo() string {
    return r.leaseExtendsInfo
}
// Brokerage Setter
// 仅淘小铺卖家需要。佣金比例(15.3对应的佣金比例为15.3%).只支持小数点后1位。多余的位数四舍五入(15.32会保存为15.3%
func (r *TaobaoItemAddRequest) SetBrokerage(brokerage string) error {
    r.brokerage = brokerage
    r.Set("brokerage", brokerage)
    return nil
}

// Brokerage Getter
func (r TaobaoItemAddRequest) GetBrokerage() string {
    return r.brokerage
}
// BizCode Setter
// 业务身份编码。淘小铺编码为"taobao-taoxiaopu"
func (r *TaobaoItemAddRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoItemAddRequest) GetBizCode() string {
    return r.bizCode
}
// ImageUrls Setter
// 此字段已经废弃，不再使用
func (r *TaobaoItemAddRequest) SetImageUrls(imageUrls []string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

// ImageUrls Getter
func (r TaobaoItemAddRequest) GetImageUrls() []string {
    return r.imageUrls
}
// LocalityLifeChooseLogis Setter
// 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
func (r *TaobaoItemAddRequest) SetLocalityLifeChooseLogis(localityLifeChooseLogis string) error {
    r.localityLifeChooseLogis = localityLifeChooseLogis
    r.Set("locality_life.choose_logis", localityLifeChooseLogis)
    return nil
}

// LocalityLifeChooseLogis Getter
func (r TaobaoItemAddRequest) GetLocalityLifeChooseLogis() string {
    return r.localityLifeChooseLogis
}
// LocalityLifeExpirydate Setter
// 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
func (r *TaobaoItemAddRequest) SetLocalityLifeExpirydate(localityLifeExpirydate string) error {
    r.localityLifeExpirydate = localityLifeExpirydate
    r.Set("locality_life.expirydate", localityLifeExpirydate)
    return nil
}

// LocalityLifeExpirydate Getter
func (r TaobaoItemAddRequest) GetLocalityLifeExpirydate() string {
    return r.localityLifeExpirydate
}
// LocalityLifeNetworkId Setter
// 网点ID
func (r *TaobaoItemAddRequest) SetLocalityLifeNetworkId(localityLifeNetworkId string) error {
    r.localityLifeNetworkId = localityLifeNetworkId
    r.Set("locality_life.network_id", localityLifeNetworkId)
    return nil
}

// LocalityLifeNetworkId Getter
func (r TaobaoItemAddRequest) GetLocalityLifeNetworkId() string {
    return r.localityLifeNetworkId
}
// LocalityLifeMerchant Setter
// 码商信息，格式为 码商id:nick
func (r *TaobaoItemAddRequest) SetLocalityLifeMerchant(localityLifeMerchant string) error {
    r.localityLifeMerchant = localityLifeMerchant
    r.Set("locality_life.merchant", localityLifeMerchant)
    return nil
}

// LocalityLifeMerchant Getter
func (r TaobaoItemAddRequest) GetLocalityLifeMerchant() string {
    return r.localityLifeMerchant
}
// LocalityLifeVerification Setter
// 核销打款 1代表核销打款 0代表非核销打款
func (r *TaobaoItemAddRequest) SetLocalityLifeVerification(localityLifeVerification string) error {
    r.localityLifeVerification = localityLifeVerification
    r.Set("locality_life.verification", localityLifeVerification)
    return nil
}

// LocalityLifeVerification Getter
func (r TaobaoItemAddRequest) GetLocalityLifeVerification() string {
    return r.localityLifeVerification
}
// LocalityLifeRefundRatio Setter
// 退款比例，百分比%前的数字,1-100的正整数值
func (r *TaobaoItemAddRequest) SetLocalityLifeRefundRatio(localityLifeRefundRatio int64) error {
    r.localityLifeRefundRatio = localityLifeRefundRatio
    r.Set("locality_life.refund_ratio", localityLifeRefundRatio)
    return nil
}

// LocalityLifeRefundRatio Getter
func (r TaobaoItemAddRequest) GetLocalityLifeRefundRatio() int64 {
    return r.localityLifeRefundRatio
}
// LocalityLifeOnsaleAutoRefundRatio Setter
// 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
func (r *TaobaoItemAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(localityLifeOnsaleAutoRefundRatio int64) error {
    r.localityLifeOnsaleAutoRefundRatio = localityLifeOnsaleAutoRefundRatio
    r.Set("locality_life.onsale_auto_refund_ratio", localityLifeOnsaleAutoRefundRatio)
    return nil
}

// LocalityLifeOnsaleAutoRefundRatio Getter
func (r TaobaoItemAddRequest) GetLocalityLifeOnsaleAutoRefundRatio() int64 {
    return r.localityLifeOnsaleAutoRefundRatio
}
// LocalityLifeRefundmafee Setter
// 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
func (r *TaobaoItemAddRequest) SetLocalityLifeRefundmafee(localityLifeRefundmafee string) error {
    r.localityLifeRefundmafee = localityLifeRefundmafee
    r.Set("locality_life.refundmafee", localityLifeRefundmafee)
    return nil
}

// LocalityLifeRefundmafee Getter
func (r TaobaoItemAddRequest) GetLocalityLifeRefundmafee() string {
    return r.localityLifeRefundmafee
}
// LocalityLifeEticket Setter
// 电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
func (r *TaobaoItemAddRequest) SetLocalityLifeEticket(localityLifeEticket string) error {
    r.localityLifeEticket = localityLifeEticket
    r.Set("locality_life.eticket", localityLifeEticket)
    return nil
}

// LocalityLifeEticket Getter
func (r TaobaoItemAddRequest) GetLocalityLifeEticket() string {
    return r.localityLifeEticket
}
// PaimaiInfoMode Setter
// 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
func (r *TaobaoItemAddRequest) SetPaimaiInfoMode(paimaiInfoMode int64) error {
    r.paimaiInfoMode = paimaiInfoMode
    r.Set("paimai_info.mode", paimaiInfoMode)
    return nil
}

// PaimaiInfoMode Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoMode() int64 {
    return r.paimaiInfoMode
}
// PaimaiInfoDeposit Setter
// 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。
func (r *TaobaoItemAddRequest) SetPaimaiInfoDeposit(paimaiInfoDeposit int64) error {
    r.paimaiInfoDeposit = paimaiInfoDeposit
    r.Set("paimai_info.deposit", paimaiInfoDeposit)
    return nil
}

// PaimaiInfoDeposit Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoDeposit() int64 {
    return r.paimaiInfoDeposit
}
// PaimaiInfoInterval Setter
// 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
func (r *TaobaoItemAddRequest) SetPaimaiInfoInterval(paimaiInfoInterval int64) error {
    r.paimaiInfoInterval = paimaiInfoInterval
    r.Set("paimai_info.interval", paimaiInfoInterval)
    return nil
}

// PaimaiInfoInterval Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoInterval() int64 {
    return r.paimaiInfoInterval
}
// PaimaiInfoReserve Setter
// 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
func (r *TaobaoItemAddRequest) SetPaimaiInfoReserve(paimaiInfoReserve float64) error {
    r.paimaiInfoReserve = paimaiInfoReserve
    r.Set("paimai_info.reserve", paimaiInfoReserve)
    return nil
}

// PaimaiInfoReserve Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoReserve() float64 {
    return r.paimaiInfoReserve
}
// PaimaiInfoValidHour Setter
// 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
func (r *TaobaoItemAddRequest) SetPaimaiInfoValidHour(paimaiInfoValidHour int64) error {
    r.paimaiInfoValidHour = paimaiInfoValidHour
    r.Set("paimai_info.valid_hour", paimaiInfoValidHour)
    return nil
}

// PaimaiInfoValidHour Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoValidHour() int64 {
    return r.paimaiInfoValidHour
}
// PaimaiInfoValidMinute Setter
// 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
func (r *TaobaoItemAddRequest) SetPaimaiInfoValidMinute(paimaiInfoValidMinute int64) error {
    r.paimaiInfoValidMinute = paimaiInfoValidMinute
    r.Set("paimai_info.valid_minute", paimaiInfoValidMinute)
    return nil
}

// PaimaiInfoValidMinute Getter
func (r TaobaoItemAddRequest) GetPaimaiInfoValidMinute() int64 {
    return r.paimaiInfoValidMinute
}
// GlobalStockType Setter
// 全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
func (r *TaobaoItemAddRequest) SetGlobalStockType(globalStockType string) error {
    r.globalStockType = globalStockType
    r.Set("global_stock_type", globalStockType)
    return nil
}

// GlobalStockType Getter
func (r TaobaoItemAddRequest) GetGlobalStockType() string {
    return r.globalStockType
}
// GlobalStockCountry Setter
// 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值请填写法定的国家名称，类如（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾），不要使用其他
func (r *TaobaoItemAddRequest) SetGlobalStockCountry(globalStockCountry string) error {
    r.globalStockCountry = globalStockCountry
    r.Set("global_stock_country", globalStockCountry)
    return nil
}

// GlobalStockCountry Getter
func (r TaobaoItemAddRequest) GetGlobalStockCountry() string {
    return r.globalStockCountry
}
// SupportCustomMade Setter
// 是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改
func (r *TaobaoItemAddRequest) SetSupportCustomMade(supportCustomMade bool) error {
    r.supportCustomMade = supportCustomMade
    r.Set("support_custom_made", supportCustomMade)
    return nil
}

// SupportCustomMade Getter
func (r TaobaoItemAddRequest) GetSupportCustomMade() bool {
    return r.supportCustomMade
}
// CustomMadeTypeId Setter
// 定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1
func (r *TaobaoItemAddRequest) SetCustomMadeTypeId(customMadeTypeId string) error {
    r.customMadeTypeId = customMadeTypeId
    r.Set("custom_made_type_id", customMadeTypeId)
    return nil
}

// CustomMadeTypeId Getter
func (r TaobaoItemAddRequest) GetCustomMadeTypeId() string {
    return r.customMadeTypeId
}
// GlobalStockDeliveryPlace Setter
// 全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”
func (r *TaobaoItemAddRequest) SetGlobalStockDeliveryPlace(globalStockDeliveryPlace string) error {
    r.globalStockDeliveryPlace = globalStockDeliveryPlace
    r.Set("global_stock_delivery_place", globalStockDeliveryPlace)
    return nil
}

// GlobalStockDeliveryPlace Getter
func (r TaobaoItemAddRequest) GetGlobalStockDeliveryPlace() string {
    return r.globalStockDeliveryPlace
}
// GlobalStockTaxFreePromise Setter
// 全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约
func (r *TaobaoItemAddRequest) SetGlobalStockTaxFreePromise(globalStockTaxFreePromise bool) error {
    r.globalStockTaxFreePromise = globalStockTaxFreePromise
    r.Set("global_stock_tax_free_promise", globalStockTaxFreePromise)
    return nil
}

// GlobalStockTaxFreePromise Getter
func (r TaobaoItemAddRequest) GetGlobalStockTaxFreePromise() bool {
    return r.globalStockTaxFreePromise
}
// InputCustomCpv Setter
// 针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上
func (r *TaobaoItemAddRequest) SetInputCustomCpv(inputCustomCpv string) error {
    r.inputCustomCpv = inputCustomCpv
    r.Set("input_custom_cpv", inputCustomCpv)
    return nil
}

// InputCustomCpv Getter
func (r TaobaoItemAddRequest) GetInputCustomCpv() string {
    return r.inputCustomCpv
}
// CpvMemo Setter
// 针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
func (r *TaobaoItemAddRequest) SetCpvMemo(cpvMemo string) error {
    r.cpvMemo = cpvMemo
    r.Set("cpv_memo", cpvMemo)
    return nil
}

// CpvMemo Getter
func (r TaobaoItemAddRequest) GetCpvMemo() string {
    return r.cpvMemo
}
