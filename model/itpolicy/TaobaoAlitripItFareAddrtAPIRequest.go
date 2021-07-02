package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareAddrtAPIRequest 【国际机票自有政策】单条往返添加 API请求
// taobao.alitrip.it.fare.addrt
//
// 自有政策往返添加接口
type TaobaoAlitripItFareAddrtAPIRequest struct {
	model.Params
	// 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
	_outFileCode string
	// 文件编号
	_fileCode string
	// （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
	_productType string
	// （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
	_stockMode string
	// 是否1/2RT，1、请填写 是或者否；默认为否
	_isRT string
	// （后期字段，预留）,1/2RT类型，当需要多填入多个时，请以","分隔 1、可填写 、旅行有效期、排除旅行有效期、班期 ；表明1/2RT 混舱计算时，取严还是各取各 2、默认值是 全部各取各
	_rtType string
	// 可组文件编号， 当需要多填入多个时，请以","分隔 1、标记可组文件的编号政策信息，可填写空白； 2、如果是否1/2RT 字段为是，则此字段为必输项
	_combinationFilecode string
	// （后期字段，预留）,是否允许缺口，1、为是或否；默认为否
	_isAllowOj string
	// （后期字段，预留）,缺口类型，1、可填单缺、双缺、始发地缺、目的地缺、或为空；默认为空（当允许缺口组合时，此项为必输项）
	_ojType string
	// （后期字段，预留）,可组缺口文件编号,当需要多填入多个时，请以","分隔 1、标记政策信息，可填写空白； 2、如果是否缺口 字段为是，则此字段为必输项
	_combinationOjFilecode string
	// 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
	_ticketingAirline string
	// 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
	_saleAirline string
	// 城市/机场选项，默认为城市1、可以填写：“机场",“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
	_addressOption string
	// 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
	_tripType string
	// 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
	_originLand string
	// 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
	_destination string
	// 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
	_transitLand string
	// 舱位， 用","表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
	_cabin string
	// 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
	_restrictFlightNo string
	// 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
	_excludeFlightNo string
	// 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
	_validDate4Dep string
	// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
	_excludeDateRange4Dep string
	// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripDatePoint4Dep string
	// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripExcludeDatePoint4Dep string
	// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
	_flightDateRestrict4Dep string
	// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_flightDatePoint4Dep string
	// 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
	_validDate4Ret string
	// 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
	_excludeDateRange4Ret string
	// 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripDatePoint4Ret string
	// 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripExcludeDatePoint4Ret string
	// 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
	_flightDateRestrict4Ret string
	// 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_flightDatePoint4Ret string
	// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
	_saleDate string
	// 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
	_minStay string
	// 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
	_maxStay string
	// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
	_adultPassengerIdentity string
	// 最小出行人数,数字1-9
	_minTravelPerson int64
	// 最大出行人数,数字1-9
	_maxTravelPerson int64
	// （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
	_gv2ChildRule string
	// 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_nationality string
	// 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_excludeNationality string
	// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
	_passengerAge string
	// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
	_ticketPrice int64
	// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
	_adultTax int64
	// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
	_childPrice string
	// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
	_childTax int64
	// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
	_returnPoint float64
	// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
	_adjustMoney int64
	// 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
	_rtCommissionFormula string
	// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
	_earlyTicketingTimeLimit int64
	// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
	_lateTicketingTimeLimit int64
	// 大客户编码，文本框
	_vipCode string
	// （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
	_fareSource string
	// （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
	_isCreatePnr string
	// 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
	_bookingOffice string
	// 必填项 赋值范围 境外电子凭证,旅行发票,差额行程单发票,等额行程单
	_receipts string
	// 是否校验票面价,1、可填写 是或者否；默认为否
	_isValidatPrice string
	// （已废除字段）,去程全部未使用可否退票，录入是或否
	_isCanRefund4Dep string
	// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
	_refundPrice4Dep string
	// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
	_refundPartPrice4Dep string
	// （已废除字段）,回程全部未使用可否退票，录入是或否
	_isCanRefund4Ret string
	// （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
	_refundPrice4Ret string
	// （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
	_refundPartPrice4Ret string
	// （已废除字段）,去程全部未使用可否改期，录入是或否
	_isCanReissue4Dep string
	// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
	_reissuePrice4Dep string
	// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
	_reissuePartPrice4Dep string
	// （已废除字段）,回程全部未使用可否改期，录入是或否
	_isCanReissue4Ret string
	// （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
	_reissuePrice4Ret string
	// （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
	_reissuePartPrice4Ret string
	// （已废除字段）,去程NOSHOW规定时限，输入正整数
	_noShowTimeLimit4Dep int64
	// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
	_isNoShowCanRefund4Dep string
	// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
	_isNoShowCanReissue4Dep string
	// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
	_noShowPenalty4Dep int64
	// （已废除字段）,回程NOSHOW规定时限，输入正整数
	_noShowTimeLimit4Ret int64
	// （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
	_isNoShowCanRefund4Ret string
	// （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
	_isNoShowCanReissue4Ret string
	// （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
	_noShowPenalty4Ret int64
	// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
	_luggageRule4Dep string
	// （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
	_luggageRule4Ret string
	// 备注,出票备注文本
	_remark string
	// 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
	_workingHours string
	// （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
	_refundRule string
	// （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
	_reissueRule string
	// （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
	_noshowRule string
	// 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
	_luggageRule string
	// 运价渠道 可选listing宝贝  默认listing
	_applyChannel string
	// 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
	_commodityType string
	// 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
	_codeSharingType string
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
	_buyTicketNotice string
	// 必填项，全部未使用可否退票，可输入:是，否
	_isCanAllRefund string
	// 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
	_refundFeeAllUnused string
	// 全部未使用退票币种，只能录入币种三字码，默认值CNY
	_refundCurrencyAllUnused string
	// 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_refundFeeTypeAllUnused string
	// 必填项，部分未使用可否退票，可输入：是，否
	_isCanPartRefund string
	// 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
	_refundFeePartUnused string
	// 部分未使用退票币种，可录入币种三字码，默认值CNY
	_refundCurrencyPartUnused string
	// 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_refundFeeTypePartUnused string
	// 必填项，去程可否改期，可输入是或否
	_canDepChange string
	// 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
	_depChangeFee string
	// 去程改期币种，可录入币种三字码，默认值CNY
	_depChangeCurrency string
	// 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_depChangeFeeType string
	// 必填项，回程可否改期，可输入是或否
	_canRetChange string
	// 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
	_retChangeFee string
	// 回程改期币种，可输入币种三字码，默认值CN
	_retChangeCurrency string
	// 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_retChangeFeeType string
	// 必填项，NOSHOW是否有限制，可输入是或否
	_noshowRestrict string
	// NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
	_noshowTimeRestrict string
	// NOSHOW时限单位(小时/天, 默认为小时)
	_noshowTimeRestrictUnit string
	// NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可改期,不可改期
	_noshowRuleType string
	// NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
	_noshowFee string
	// NOSHOW币种,可录入币种三字码，默认值CNY
	_noshowCurrency string
	// 运价基础，最大长度8
	_farebasis string
	// 运价类型，最大长度3
	_fareTypeCode string
	// 运价tariff，最大长度3
	_tariff string
	// 运价规则id，最大长度4
	_ruleId string
	// 运价组合适用方向,0(或者字段不存在):不限制/1:仅作用在去程/2:仅作用在回程
	_fareDirectDestrict int64
}

// NewTaobaoAlitripItFareAddrtRequest 初始化TaobaoAlitripItFareAddrtAPIRequest对象
func NewTaobaoAlitripItFareAddrtRequest() *TaobaoAlitripItFareAddrtAPIRequest {
	return &TaobaoAlitripItFareAddrtAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareAddrtAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.addrt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareAddrtAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OutFileCode Setter
// 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOutFileCode(_outFileCode string) error {
	r._outFileCode = _outFileCode
	r.Set("outFileCode", _outFileCode)
	return nil
}

// Get OutFileCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOutFileCode() string {
	return r._outFileCode
}

// Set is FileCode Setter
// 文件编号
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFileCode(_fileCode string) error {
	r._fileCode = _fileCode
	r.Set("fileCode", _fileCode)
	return nil
}

// Get FileCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFileCode() string {
	return r._fileCode
}

// Set is ProductType Setter
// （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetProductType(_productType string) error {
	r._productType = _productType
	r.Set("productType", _productType)
	return nil
}

// Get ProductType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetProductType() string {
	return r._productType
}

// Set is StockMode Setter
// （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetStockMode(_stockMode string) error {
	r._stockMode = _stockMode
	r.Set("stockMode", _stockMode)
	return nil
}

// Get StockMode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetStockMode() string {
	return r._stockMode
}

// Set is IsRT Setter
// 是否1/2RT，1、请填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsRT(_isRT string) error {
	r._isRT = _isRT
	r.Set("isRT", _isRT)
	return nil
}

// Get IsRT Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsRT() string {
	return r._isRT
}

// Set is RtType Setter
// （后期字段，预留）,1/2RT类型，当需要多填入多个时，请以","分隔 1、可填写 、旅行有效期、排除旅行有效期、班期 ；表明1/2RT 混舱计算时，取严还是各取各 2、默认值是 全部各取各
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRtType(_rtType string) error {
	r._rtType = _rtType
	r.Set("rtType", _rtType)
	return nil
}

// Get RtType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRtType() string {
	return r._rtType
}

// Set is CombinationFilecode Setter
// 可组文件编号， 当需要多填入多个时，请以","分隔 1、标记可组文件的编号政策信息，可填写空白； 2、如果是否1/2RT 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCombinationFilecode(_combinationFilecode string) error {
	r._combinationFilecode = _combinationFilecode
	r.Set("combinationFilecode", _combinationFilecode)
	return nil
}

// Get CombinationFilecode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCombinationFilecode() string {
	return r._combinationFilecode
}

// Set is IsAllowOj Setter
// （后期字段，预留）,是否允许缺口，1、为是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsAllowOj(_isAllowOj string) error {
	r._isAllowOj = _isAllowOj
	r.Set("isAllowOj", _isAllowOj)
	return nil
}

// Get IsAllowOj Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsAllowOj() string {
	return r._isAllowOj
}

// Set is OjType Setter
// （后期字段，预留）,缺口类型，1、可填单缺、双缺、始发地缺、目的地缺、或为空；默认为空（当允许缺口组合时，此项为必输项）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOjType(_ojType string) error {
	r._ojType = _ojType
	r.Set("ojType", _ojType)
	return nil
}

// Get OjType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOjType() string {
	return r._ojType
}

// Set is CombinationOjFilecode Setter
// （后期字段，预留）,可组缺口文件编号,当需要多填入多个时，请以","分隔 1、标记政策信息，可填写空白； 2、如果是否缺口 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCombinationOjFilecode(_combinationOjFilecode string) error {
	r._combinationOjFilecode = _combinationOjFilecode
	r.Set("combinationOjFilecode", _combinationOjFilecode)
	return nil
}

// Get CombinationOjFilecode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCombinationOjFilecode() string {
	return r._combinationOjFilecode
}

// Set is TicketingAirline Setter
// 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTicketingAirline(_ticketingAirline string) error {
	r._ticketingAirline = _ticketingAirline
	r.Set("ticketingAirline", _ticketingAirline)
	return nil
}

// Get TicketingAirline Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTicketingAirline() string {
	return r._ticketingAirline
}

// Set is SaleAirline Setter
// 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetSaleAirline(_saleAirline string) error {
	r._saleAirline = _saleAirline
	r.Set("saleAirline", _saleAirline)
	return nil
}

// Get SaleAirline Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetSaleAirline() string {
	return r._saleAirline
}

// Set is AddressOption Setter
// 城市/机场选项，默认为城市1、可以填写：“机场",“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAddressOption(_addressOption string) error {
	r._addressOption = _addressOption
	r.Set("addressOption", _addressOption)
	return nil
}

// Get AddressOption Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAddressOption() string {
	return r._addressOption
}

// Set is TripType Setter
// 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripType(_tripType string) error {
	r._tripType = _tripType
	r.Set("tripType", _tripType)
	return nil
}

// Get TripType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripType() string {
	return r._tripType
}

// Set is OriginLand Setter
// 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOriginLand(_originLand string) error {
	r._originLand = _originLand
	r.Set("originLand", _originLand)
	return nil
}

// Get OriginLand Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOriginLand() string {
	return r._originLand
}

// Set is Destination Setter
// 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDestination(_destination string) error {
	r._destination = _destination
	r.Set("destination", _destination)
	return nil
}

// Get Destination Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDestination() string {
	return r._destination
}

// Set is TransitLand Setter
// 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTransitLand(_transitLand string) error {
	r._transitLand = _transitLand
	r.Set("transitLand", _transitLand)
	return nil
}

// Get TransitLand Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTransitLand() string {
	return r._transitLand
}

// Set is Cabin Setter
// 舱位， 用","表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCabin(_cabin string) error {
	r._cabin = _cabin
	r.Set("cabin", _cabin)
	return nil
}

// Get Cabin Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCabin() string {
	return r._cabin
}

// Set is RestrictFlightNo Setter
// 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRestrictFlightNo(_restrictFlightNo string) error {
	r._restrictFlightNo = _restrictFlightNo
	r.Set("restrictFlightNo", _restrictFlightNo)
	return nil
}

// Get RestrictFlightNo Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRestrictFlightNo() string {
	return r._restrictFlightNo
}

// Set is ExcludeFlightNo Setter
// 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeFlightNo(_excludeFlightNo string) error {
	r._excludeFlightNo = _excludeFlightNo
	r.Set("excludeFlightNo", _excludeFlightNo)
	return nil
}

// Get ExcludeFlightNo Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeFlightNo() string {
	return r._excludeFlightNo
}

// Set is ValidDate4Dep Setter
// 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetValidDate4Dep(_validDate4Dep string) error {
	r._validDate4Dep = _validDate4Dep
	r.Set("validDate4Dep", _validDate4Dep)
	return nil
}

// Get ValidDate4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetValidDate4Dep() string {
	return r._validDate4Dep
}

// Set is ExcludeDateRange4Dep Setter
// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeDateRange4Dep(_excludeDateRange4Dep string) error {
	r._excludeDateRange4Dep = _excludeDateRange4Dep
	r.Set("excludeDateRange4Dep", _excludeDateRange4Dep)
	return nil
}

// Get ExcludeDateRange4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeDateRange4Dep() string {
	return r._excludeDateRange4Dep
}

// Set is TripDatePoint4Dep Setter
// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripDatePoint4Dep(_tripDatePoint4Dep string) error {
	r._tripDatePoint4Dep = _tripDatePoint4Dep
	r.Set("tripDatePoint4Dep", _tripDatePoint4Dep)
	return nil
}

// Get TripDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripDatePoint4Dep() string {
	return r._tripDatePoint4Dep
}

// Set is TripExcludeDatePoint4Dep Setter
// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripExcludeDatePoint4Dep(_tripExcludeDatePoint4Dep string) error {
	r._tripExcludeDatePoint4Dep = _tripExcludeDatePoint4Dep
	r.Set("tripExcludeDatePoint4Dep", _tripExcludeDatePoint4Dep)
	return nil
}

// Get TripExcludeDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripExcludeDatePoint4Dep() string {
	return r._tripExcludeDatePoint4Dep
}

// Set is FlightDateRestrict4Dep Setter
// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDateRestrict4Dep(_flightDateRestrict4Dep string) error {
	r._flightDateRestrict4Dep = _flightDateRestrict4Dep
	r.Set("flightDateRestrict4Dep", _flightDateRestrict4Dep)
	return nil
}

// Get FlightDateRestrict4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDateRestrict4Dep() string {
	return r._flightDateRestrict4Dep
}

// Set is FlightDatePoint4Dep Setter
// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDatePoint4Dep(_flightDatePoint4Dep string) error {
	r._flightDatePoint4Dep = _flightDatePoint4Dep
	r.Set("flightDatePoint4Dep", _flightDatePoint4Dep)
	return nil
}

// Get FlightDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDatePoint4Dep() string {
	return r._flightDatePoint4Dep
}

// Set is ValidDate4Ret Setter
// 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetValidDate4Ret(_validDate4Ret string) error {
	r._validDate4Ret = _validDate4Ret
	r.Set("validDate4Ret", _validDate4Ret)
	return nil
}

// Get ValidDate4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetValidDate4Ret() string {
	return r._validDate4Ret
}

// Set is ExcludeDateRange4Ret Setter
// 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeDateRange4Ret(_excludeDateRange4Ret string) error {
	r._excludeDateRange4Ret = _excludeDateRange4Ret
	r.Set("excludeDateRange4Ret", _excludeDateRange4Ret)
	return nil
}

// Get ExcludeDateRange4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeDateRange4Ret() string {
	return r._excludeDateRange4Ret
}

// Set is TripDatePoint4Ret Setter
// 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripDatePoint4Ret(_tripDatePoint4Ret string) error {
	r._tripDatePoint4Ret = _tripDatePoint4Ret
	r.Set("tripDatePoint4Ret", _tripDatePoint4Ret)
	return nil
}

// Get TripDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripDatePoint4Ret() string {
	return r._tripDatePoint4Ret
}

// Set is TripExcludeDatePoint4Ret Setter
// 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripExcludeDatePoint4Ret(_tripExcludeDatePoint4Ret string) error {
	r._tripExcludeDatePoint4Ret = _tripExcludeDatePoint4Ret
	r.Set("tripExcludeDatePoint4Ret", _tripExcludeDatePoint4Ret)
	return nil
}

// Get TripExcludeDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripExcludeDatePoint4Ret() string {
	return r._tripExcludeDatePoint4Ret
}

// Set is FlightDateRestrict4Ret Setter
// 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDateRestrict4Ret(_flightDateRestrict4Ret string) error {
	r._flightDateRestrict4Ret = _flightDateRestrict4Ret
	r.Set("flightDateRestrict4Ret", _flightDateRestrict4Ret)
	return nil
}

// Get FlightDateRestrict4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDateRestrict4Ret() string {
	return r._flightDateRestrict4Ret
}

// Set is FlightDatePoint4Ret Setter
// 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDatePoint4Ret(_flightDatePoint4Ret string) error {
	r._flightDatePoint4Ret = _flightDatePoint4Ret
	r.Set("flightDatePoint4Ret", _flightDatePoint4Ret)
	return nil
}

// Get FlightDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDatePoint4Ret() string {
	return r._flightDatePoint4Ret
}

// Set is SaleDate Setter
// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetSaleDate(_saleDate string) error {
	r._saleDate = _saleDate
	r.Set("saleDate", _saleDate)
	return nil
}

// Get SaleDate Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetSaleDate() string {
	return r._saleDate
}

// Set is MinStay Setter
// 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMinStay(_minStay string) error {
	r._minStay = _minStay
	r.Set("minStay", _minStay)
	return nil
}

// Get MinStay Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMinStay() string {
	return r._minStay
}

// Set is MaxStay Setter
// 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMaxStay(_maxStay string) error {
	r._maxStay = _maxStay
	r.Set("maxStay", _maxStay)
	return nil
}

// Get MaxStay Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMaxStay() string {
	return r._maxStay
}

// Set is AdultPassengerIdentity Setter
// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdultPassengerIdentity(_adultPassengerIdentity string) error {
	r._adultPassengerIdentity = _adultPassengerIdentity
	r.Set("adultPassengerIdentity", _adultPassengerIdentity)
	return nil
}

// Get AdultPassengerIdentity Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdultPassengerIdentity() string {
	return r._adultPassengerIdentity
}

// Set is MinTravelPerson Setter
// 最小出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMinTravelPerson(_minTravelPerson int64) error {
	r._minTravelPerson = _minTravelPerson
	r.Set("minTravelPerson", _minTravelPerson)
	return nil
}

// Get MinTravelPerson Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMinTravelPerson() int64 {
	return r._minTravelPerson
}

// Set is MaxTravelPerson Setter
// 最大出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMaxTravelPerson(_maxTravelPerson int64) error {
	r._maxTravelPerson = _maxTravelPerson
	r.Set("maxTravelPerson", _maxTravelPerson)
	return nil
}

// Get MaxTravelPerson Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMaxTravelPerson() int64 {
	return r._maxTravelPerson
}

// Set is Gv2ChildRule Setter
// （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetGv2ChildRule(_gv2ChildRule string) error {
	r._gv2ChildRule = _gv2ChildRule
	r.Set("gv2ChildRule", _gv2ChildRule)
	return nil
}

// Get Gv2ChildRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetGv2ChildRule() string {
	return r._gv2ChildRule
}

// Set is Nationality Setter
// 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNationality(_nationality string) error {
	r._nationality = _nationality
	r.Set("nationality", _nationality)
	return nil
}

// Get Nationality Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNationality() string {
	return r._nationality
}

// Set is ExcludeNationality Setter
// 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeNationality(_excludeNationality string) error {
	r._excludeNationality = _excludeNationality
	r.Set("excludeNationality", _excludeNationality)
	return nil
}

// Get ExcludeNationality Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeNationality() string {
	return r._excludeNationality
}

// Set is PassengerAge Setter
// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetPassengerAge(_passengerAge string) error {
	r._passengerAge = _passengerAge
	r.Set("passengerAge", _passengerAge)
	return nil
}

// Get PassengerAge Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetPassengerAge() string {
	return r._passengerAge
}

// Set is TicketPrice Setter
// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTicketPrice(_ticketPrice int64) error {
	r._ticketPrice = _ticketPrice
	r.Set("ticketPrice", _ticketPrice)
	return nil
}

// Get TicketPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTicketPrice() int64 {
	return r._ticketPrice
}

// Set is AdultTax Setter
// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdultTax(_adultTax int64) error {
	r._adultTax = _adultTax
	r.Set("adultTax", _adultTax)
	return nil
}

// Get AdultTax Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdultTax() int64 {
	return r._adultTax
}

// Set is ChildPrice Setter
// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetChildPrice(_childPrice string) error {
	r._childPrice = _childPrice
	r.Set("childPrice", _childPrice)
	return nil
}

// Get ChildPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetChildPrice() string {
	return r._childPrice
}

// Set is ChildTax Setter
// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetChildTax(_childTax int64) error {
	r._childTax = _childTax
	r.Set("childTax", _childTax)
	return nil
}

// Get ChildTax Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetChildTax() int64 {
	return r._childTax
}

// Set is ReturnPoint Setter
// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReturnPoint(_returnPoint float64) error {
	r._returnPoint = _returnPoint
	r.Set("returnPoint", _returnPoint)
	return nil
}

// Get ReturnPoint Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReturnPoint() float64 {
	return r._returnPoint
}

// Set is AdjustMoney Setter
// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdjustMoney(_adjustMoney int64) error {
	r._adjustMoney = _adjustMoney
	r.Set("adjustMoney", _adjustMoney)
	return nil
}

// Get AdjustMoney Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdjustMoney() int64 {
	return r._adjustMoney
}

// Set is RtCommissionFormula Setter
// 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRtCommissionFormula(_rtCommissionFormula string) error {
	r._rtCommissionFormula = _rtCommissionFormula
	r.Set("rtCommissionFormula", _rtCommissionFormula)
	return nil
}

// Get RtCommissionFormula Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRtCommissionFormula() string {
	return r._rtCommissionFormula
}

// Set is EarlyTicketingTimeLimit Setter
// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetEarlyTicketingTimeLimit(_earlyTicketingTimeLimit int64) error {
	r._earlyTicketingTimeLimit = _earlyTicketingTimeLimit
	r.Set("earlyTicketingTimeLimit", _earlyTicketingTimeLimit)
	return nil
}

// Get EarlyTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetEarlyTicketingTimeLimit() int64 {
	return r._earlyTicketingTimeLimit
}

// Set is LateTicketingTimeLimit Setter
// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLateTicketingTimeLimit(_lateTicketingTimeLimit int64) error {
	r._lateTicketingTimeLimit = _lateTicketingTimeLimit
	r.Set("lateTicketingTimeLimit", _lateTicketingTimeLimit)
	return nil
}

// Get LateTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLateTicketingTimeLimit() int64 {
	return r._lateTicketingTimeLimit
}

// Set is VipCode Setter
// 大客户编码，文本框
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetVipCode(_vipCode string) error {
	r._vipCode = _vipCode
	r.Set("vipCode", _vipCode)
	return nil
}

// Get VipCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetVipCode() string {
	return r._vipCode
}

// Set is FareSource Setter
// （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareSource(_fareSource string) error {
	r._fareSource = _fareSource
	r.Set("fareSource", _fareSource)
	return nil
}

// Get FareSource Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareSource() string {
	return r._fareSource
}

// Set is IsCreatePnr Setter
// （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCreatePnr(_isCreatePnr string) error {
	r._isCreatePnr = _isCreatePnr
	r.Set("isCreatePnr", _isCreatePnr)
	return nil
}

// Get IsCreatePnr Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCreatePnr() string {
	return r._isCreatePnr
}

// Set is BookingOffice Setter
// 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetBookingOffice(_bookingOffice string) error {
	r._bookingOffice = _bookingOffice
	r.Set("bookingOffice", _bookingOffice)
	return nil
}

// Get BookingOffice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetBookingOffice() string {
	return r._bookingOffice
}

// Set is Receipts Setter
// 必填项 赋值范围 境外电子凭证,旅行发票,差额行程单发票,等额行程单
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReceipts(_receipts string) error {
	r._receipts = _receipts
	r.Set("receipts", _receipts)
	return nil
}

// Get Receipts Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReceipts() string {
	return r._receipts
}

// Set is IsValidatPrice Setter
// 是否校验票面价,1、可填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsValidatPrice(_isValidatPrice string) error {
	r._isValidatPrice = _isValidatPrice
	r.Set("isValidatPrice", _isValidatPrice)
	return nil
}

// Get IsValidatPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsValidatPrice() string {
	return r._isValidatPrice
}

// Set is IsCanRefund4Dep Setter
// （已废除字段）,去程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanRefund4Dep(_isCanRefund4Dep string) error {
	r._isCanRefund4Dep = _isCanRefund4Dep
	r.Set("isCanRefund4Dep", _isCanRefund4Dep)
	return nil
}

// Get IsCanRefund4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanRefund4Dep() string {
	return r._isCanRefund4Dep
}

// Set is RefundPrice4Dep Setter
// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPrice4Dep(_refundPrice4Dep string) error {
	r._refundPrice4Dep = _refundPrice4Dep
	r.Set("refundPrice4Dep", _refundPrice4Dep)
	return nil
}

// Get RefundPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPrice4Dep() string {
	return r._refundPrice4Dep
}

// Set is RefundPartPrice4Dep Setter
// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPartPrice4Dep(_refundPartPrice4Dep string) error {
	r._refundPartPrice4Dep = _refundPartPrice4Dep
	r.Set("refundPartPrice4Dep", _refundPartPrice4Dep)
	return nil
}

// Get RefundPartPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPartPrice4Dep() string {
	return r._refundPartPrice4Dep
}

// Set is IsCanRefund4Ret Setter
// （已废除字段）,回程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanRefund4Ret(_isCanRefund4Ret string) error {
	r._isCanRefund4Ret = _isCanRefund4Ret
	r.Set("isCanRefund4Ret", _isCanRefund4Ret)
	return nil
}

// Get IsCanRefund4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanRefund4Ret() string {
	return r._isCanRefund4Ret
}

// Set is RefundPrice4Ret Setter
// （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPrice4Ret(_refundPrice4Ret string) error {
	r._refundPrice4Ret = _refundPrice4Ret
	r.Set("refundPrice4Ret", _refundPrice4Ret)
	return nil
}

// Get RefundPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPrice4Ret() string {
	return r._refundPrice4Ret
}

// Set is RefundPartPrice4Ret Setter
// （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPartPrice4Ret(_refundPartPrice4Ret string) error {
	r._refundPartPrice4Ret = _refundPartPrice4Ret
	r.Set("refundPartPrice4Ret", _refundPartPrice4Ret)
	return nil
}

// Get RefundPartPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPartPrice4Ret() string {
	return r._refundPartPrice4Ret
}

// Set is IsCanReissue4Dep Setter
// （已废除字段）,去程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanReissue4Dep(_isCanReissue4Dep string) error {
	r._isCanReissue4Dep = _isCanReissue4Dep
	r.Set("isCanReissue4Dep", _isCanReissue4Dep)
	return nil
}

// Get IsCanReissue4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanReissue4Dep() string {
	return r._isCanReissue4Dep
}

// Set is ReissuePrice4Dep Setter
// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePrice4Dep(_reissuePrice4Dep string) error {
	r._reissuePrice4Dep = _reissuePrice4Dep
	r.Set("reissuePrice4Dep", _reissuePrice4Dep)
	return nil
}

// Get ReissuePrice4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePrice4Dep() string {
	return r._reissuePrice4Dep
}

// Set is ReissuePartPrice4Dep Setter
// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePartPrice4Dep(_reissuePartPrice4Dep string) error {
	r._reissuePartPrice4Dep = _reissuePartPrice4Dep
	r.Set("reissuePartPrice4Dep", _reissuePartPrice4Dep)
	return nil
}

// Get ReissuePartPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePartPrice4Dep() string {
	return r._reissuePartPrice4Dep
}

// Set is IsCanReissue4Ret Setter
// （已废除字段）,回程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanReissue4Ret(_isCanReissue4Ret string) error {
	r._isCanReissue4Ret = _isCanReissue4Ret
	r.Set("isCanReissue4Ret", _isCanReissue4Ret)
	return nil
}

// Get IsCanReissue4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanReissue4Ret() string {
	return r._isCanReissue4Ret
}

// Set is ReissuePrice4Ret Setter
// （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePrice4Ret(_reissuePrice4Ret string) error {
	r._reissuePrice4Ret = _reissuePrice4Ret
	r.Set("reissuePrice4Ret", _reissuePrice4Ret)
	return nil
}

// Get ReissuePrice4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePrice4Ret() string {
	return r._reissuePrice4Ret
}

// Set is ReissuePartPrice4Ret Setter
// （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePartPrice4Ret(_reissuePartPrice4Ret string) error {
	r._reissuePartPrice4Ret = _reissuePartPrice4Ret
	r.Set("reissuePartPrice4Ret", _reissuePartPrice4Ret)
	return nil
}

// Get ReissuePartPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePartPrice4Ret() string {
	return r._reissuePartPrice4Ret
}

// Set is NoShowTimeLimit4Dep Setter
// （已废除字段）,去程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowTimeLimit4Dep(_noShowTimeLimit4Dep int64) error {
	r._noShowTimeLimit4Dep = _noShowTimeLimit4Dep
	r.Set("noShowTimeLimit4Dep", _noShowTimeLimit4Dep)
	return nil
}

// Get NoShowTimeLimit4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowTimeLimit4Dep() int64 {
	return r._noShowTimeLimit4Dep
}

// Set is IsNoShowCanRefund4Dep Setter
// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanRefund4Dep(_isNoShowCanRefund4Dep string) error {
	r._isNoShowCanRefund4Dep = _isNoShowCanRefund4Dep
	r.Set("isNoShowCanRefund4Dep", _isNoShowCanRefund4Dep)
	return nil
}

// Get IsNoShowCanRefund4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanRefund4Dep() string {
	return r._isNoShowCanRefund4Dep
}

// Set is IsNoShowCanReissue4Dep Setter
// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanReissue4Dep(_isNoShowCanReissue4Dep string) error {
	r._isNoShowCanReissue4Dep = _isNoShowCanReissue4Dep
	r.Set("isNoShowCanReissue4Dep", _isNoShowCanReissue4Dep)
	return nil
}

// Get IsNoShowCanReissue4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanReissue4Dep() string {
	return r._isNoShowCanReissue4Dep
}

// Set is NoShowPenalty4Dep Setter
// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowPenalty4Dep(_noShowPenalty4Dep int64) error {
	r._noShowPenalty4Dep = _noShowPenalty4Dep
	r.Set("noShowPenalty4Dep", _noShowPenalty4Dep)
	return nil
}

// Get NoShowPenalty4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowPenalty4Dep() int64 {
	return r._noShowPenalty4Dep
}

// Set is NoShowTimeLimit4Ret Setter
// （已废除字段）,回程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowTimeLimit4Ret(_noShowTimeLimit4Ret int64) error {
	r._noShowTimeLimit4Ret = _noShowTimeLimit4Ret
	r.Set("noShowTimeLimit4Ret", _noShowTimeLimit4Ret)
	return nil
}

// Get NoShowTimeLimit4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowTimeLimit4Ret() int64 {
	return r._noShowTimeLimit4Ret
}

// Set is IsNoShowCanRefund4Ret Setter
// （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanRefund4Ret(_isNoShowCanRefund4Ret string) error {
	r._isNoShowCanRefund4Ret = _isNoShowCanRefund4Ret
	r.Set("isNoShowCanRefund4Ret", _isNoShowCanRefund4Ret)
	return nil
}

// Get IsNoShowCanRefund4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanRefund4Ret() string {
	return r._isNoShowCanRefund4Ret
}

// Set is IsNoShowCanReissue4Ret Setter
// （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanReissue4Ret(_isNoShowCanReissue4Ret string) error {
	r._isNoShowCanReissue4Ret = _isNoShowCanReissue4Ret
	r.Set("isNoShowCanReissue4Ret", _isNoShowCanReissue4Ret)
	return nil
}

// Get IsNoShowCanReissue4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanReissue4Ret() string {
	return r._isNoShowCanReissue4Ret
}

// Set is NoShowPenalty4Ret Setter
// （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowPenalty4Ret(_noShowPenalty4Ret int64) error {
	r._noShowPenalty4Ret = _noShowPenalty4Ret
	r.Set("noShowPenalty4Ret", _noShowPenalty4Ret)
	return nil
}

// Get NoShowPenalty4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowPenalty4Ret() int64 {
	return r._noShowPenalty4Ret
}

// Set is LuggageRule4Dep Setter
// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule4Dep(_luggageRule4Dep string) error {
	r._luggageRule4Dep = _luggageRule4Dep
	r.Set("luggageRule4Dep", _luggageRule4Dep)
	return nil
}

// Get LuggageRule4Dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule4Dep() string {
	return r._luggageRule4Dep
}

// Set is LuggageRule4Ret Setter
// （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule4Ret(_luggageRule4Ret string) error {
	r._luggageRule4Ret = _luggageRule4Ret
	r.Set("luggageRule4Ret", _luggageRule4Ret)
	return nil
}

// Get LuggageRule4Ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule4Ret() string {
	return r._luggageRule4Ret
}

// Set is Remark Setter
// 备注,出票备注文本
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRemark() string {
	return r._remark
}

// Set is WorkingHours Setter
// 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetWorkingHours(_workingHours string) error {
	r._workingHours = _workingHours
	r.Set("workingHours", _workingHours)
	return nil
}

// Get WorkingHours Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetWorkingHours() string {
	return r._workingHours
}

// Set is RefundRule Setter
// （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundRule(_refundRule string) error {
	r._refundRule = _refundRule
	r.Set("refundRule", _refundRule)
	return nil
}

// Get RefundRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundRule() string {
	return r._refundRule
}

// Set is ReissueRule Setter
// （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissueRule(_reissueRule string) error {
	r._reissueRule = _reissueRule
	r.Set("reissueRule", _reissueRule)
	return nil
}

// Get ReissueRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissueRule() string {
	return r._reissueRule
}

// Set is NoshowRule Setter
// （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRule(_noshowRule string) error {
	r._noshowRule = _noshowRule
	r.Set("noshowRule", _noshowRule)
	return nil
}

// Get NoshowRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRule() string {
	return r._noshowRule
}

// Set is LuggageRule Setter
// 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule(_luggageRule string) error {
	r._luggageRule = _luggageRule
	r.Set("luggageRule", _luggageRule)
	return nil
}

// Get LuggageRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule() string {
	return r._luggageRule
}

// Set is ApplyChannel Setter
// 运价渠道 可选listing宝贝  默认listing
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetApplyChannel(_applyChannel string) error {
	r._applyChannel = _applyChannel
	r.Set("applyChannel", _applyChannel)
	return nil
}

// Get ApplyChannel Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetApplyChannel() string {
	return r._applyChannel
}

// Set is CommodityType Setter
// 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCommodityType(_commodityType string) error {
	r._commodityType = _commodityType
	r.Set("commodityType", _commodityType)
	return nil
}

// Get CommodityType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCommodityType() string {
	return r._commodityType
}

// Set is CodeSharingType Setter
// 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCodeSharingType(_codeSharingType string) error {
	r._codeSharingType = _codeSharingType
	r.Set("codeSharingType", _codeSharingType)
	return nil
}

// Get CodeSharingType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCodeSharingType() string {
	return r._codeSharingType
}

// Set is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// Get ExtendAttributes Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// Set is BuyTicketNotice Setter
// 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetBuyTicketNotice(_buyTicketNotice string) error {
	r._buyTicketNotice = _buyTicketNotice
	r.Set("buyTicketNotice", _buyTicketNotice)
	return nil
}

// Get BuyTicketNotice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetBuyTicketNotice() string {
	return r._buyTicketNotice
}

// Set is IsCanAllRefund Setter
// 必填项，全部未使用可否退票，可输入:是，否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanAllRefund(_isCanAllRefund string) error {
	r._isCanAllRefund = _isCanAllRefund
	r.Set("isCanAllRefund", _isCanAllRefund)
	return nil
}

// Get IsCanAllRefund Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanAllRefund() string {
	return r._isCanAllRefund
}

// Set is RefundFeeAllUnused Setter
// 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeAllUnused(_refundFeeAllUnused string) error {
	r._refundFeeAllUnused = _refundFeeAllUnused
	r.Set("refundFeeAllUnused", _refundFeeAllUnused)
	return nil
}

// Get RefundFeeAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeAllUnused() string {
	return r._refundFeeAllUnused
}

// Set is RefundCurrencyAllUnused Setter
// 全部未使用退票币种，只能录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundCurrencyAllUnused(_refundCurrencyAllUnused string) error {
	r._refundCurrencyAllUnused = _refundCurrencyAllUnused
	r.Set("refundCurrencyAllUnused", _refundCurrencyAllUnused)
	return nil
}

// Get RefundCurrencyAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundCurrencyAllUnused() string {
	return r._refundCurrencyAllUnused
}

// Set is RefundFeeTypeAllUnused Setter
// 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeTypeAllUnused(_refundFeeTypeAllUnused string) error {
	r._refundFeeTypeAllUnused = _refundFeeTypeAllUnused
	r.Set("refundFeeTypeAllUnused", _refundFeeTypeAllUnused)
	return nil
}

// Get RefundFeeTypeAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeTypeAllUnused() string {
	return r._refundFeeTypeAllUnused
}

// Set is IsCanPartRefund Setter
// 必填项，部分未使用可否退票，可输入：是，否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanPartRefund(_isCanPartRefund string) error {
	r._isCanPartRefund = _isCanPartRefund
	r.Set("isCanPartRefund", _isCanPartRefund)
	return nil
}

// Get IsCanPartRefund Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanPartRefund() string {
	return r._isCanPartRefund
}

// Set is RefundFeePartUnused Setter
// 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeePartUnused(_refundFeePartUnused string) error {
	r._refundFeePartUnused = _refundFeePartUnused
	r.Set("refundFeePartUnused", _refundFeePartUnused)
	return nil
}

// Get RefundFeePartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeePartUnused() string {
	return r._refundFeePartUnused
}

// Set is RefundCurrencyPartUnused Setter
// 部分未使用退票币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundCurrencyPartUnused(_refundCurrencyPartUnused string) error {
	r._refundCurrencyPartUnused = _refundCurrencyPartUnused
	r.Set("refundCurrencyPartUnused", _refundCurrencyPartUnused)
	return nil
}

// Get RefundCurrencyPartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundCurrencyPartUnused() string {
	return r._refundCurrencyPartUnused
}

// Set is RefundFeeTypePartUnused Setter
// 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeTypePartUnused(_refundFeeTypePartUnused string) error {
	r._refundFeeTypePartUnused = _refundFeeTypePartUnused
	r.Set("refundFeeTypePartUnused", _refundFeeTypePartUnused)
	return nil
}

// Get RefundFeeTypePartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeTypePartUnused() string {
	return r._refundFeeTypePartUnused
}

// Set is CanDepChange Setter
// 必填项，去程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCanDepChange(_canDepChange string) error {
	r._canDepChange = _canDepChange
	r.Set("canDepChange", _canDepChange)
	return nil
}

// Get CanDepChange Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCanDepChange() string {
	return r._canDepChange
}

// Set is DepChangeFee Setter
// 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeFee(_depChangeFee string) error {
	r._depChangeFee = _depChangeFee
	r.Set("depChangeFee", _depChangeFee)
	return nil
}

// Get DepChangeFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeFee() string {
	return r._depChangeFee
}

// Set is DepChangeCurrency Setter
// 去程改期币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeCurrency(_depChangeCurrency string) error {
	r._depChangeCurrency = _depChangeCurrency
	r.Set("depChangeCurrency", _depChangeCurrency)
	return nil
}

// Get DepChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeCurrency() string {
	return r._depChangeCurrency
}

// Set is DepChangeFeeType Setter
// 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeFeeType(_depChangeFeeType string) error {
	r._depChangeFeeType = _depChangeFeeType
	r.Set("depChangeFeeType", _depChangeFeeType)
	return nil
}

// Get DepChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeFeeType() string {
	return r._depChangeFeeType
}

// Set is CanRetChange Setter
// 必填项，回程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCanRetChange(_canRetChange string) error {
	r._canRetChange = _canRetChange
	r.Set("canRetChange", _canRetChange)
	return nil
}

// Get CanRetChange Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCanRetChange() string {
	return r._canRetChange
}

// Set is RetChangeFee Setter
// 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeFee(_retChangeFee string) error {
	r._retChangeFee = _retChangeFee
	r.Set("retChangeFee", _retChangeFee)
	return nil
}

// Get RetChangeFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeFee() string {
	return r._retChangeFee
}

// Set is RetChangeCurrency Setter
// 回程改期币种，可输入币种三字码，默认值CN
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeCurrency(_retChangeCurrency string) error {
	r._retChangeCurrency = _retChangeCurrency
	r.Set("retChangeCurrency", _retChangeCurrency)
	return nil
}

// Get RetChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeCurrency() string {
	return r._retChangeCurrency
}

// Set is RetChangeFeeType Setter
// 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeFeeType(_retChangeFeeType string) error {
	r._retChangeFeeType = _retChangeFeeType
	r.Set("retChangeFeeType", _retChangeFeeType)
	return nil
}

// Get RetChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeFeeType() string {
	return r._retChangeFeeType
}

// Set is NoshowRestrict Setter
// 必填项，NOSHOW是否有限制，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRestrict(_noshowRestrict string) error {
	r._noshowRestrict = _noshowRestrict
	r.Set("noshowRestrict", _noshowRestrict)
	return nil
}

// Get NoshowRestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRestrict() string {
	return r._noshowRestrict
}

// Set is NoshowTimeRestrict Setter
// NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowTimeRestrict(_noshowTimeRestrict string) error {
	r._noshowTimeRestrict = _noshowTimeRestrict
	r.Set("noshowTimeRestrict", _noshowTimeRestrict)
	return nil
}

// Get NoshowTimeRestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowTimeRestrict() string {
	return r._noshowTimeRestrict
}

// Set is NoshowTimeRestrictUnit Setter
// NOSHOW时限单位(小时/天, 默认为小时)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowTimeRestrictUnit(_noshowTimeRestrictUnit string) error {
	r._noshowTimeRestrictUnit = _noshowTimeRestrictUnit
	r.Set("noshowTimeRestrictUnit", _noshowTimeRestrictUnit)
	return nil
}

// Get NoshowTimeRestrictUnit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowTimeRestrictUnit() string {
	return r._noshowTimeRestrictUnit
}

// Set is NoshowRuleType Setter
// NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可改期,不可改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRuleType(_noshowRuleType string) error {
	r._noshowRuleType = _noshowRuleType
	r.Set("noshowRuleType", _noshowRuleType)
	return nil
}

// Get NoshowRuleType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRuleType() string {
	return r._noshowRuleType
}

// Set is NoshowFee Setter
// NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowFee(_noshowFee string) error {
	r._noshowFee = _noshowFee
	r.Set("noshowFee", _noshowFee)
	return nil
}

// Get NoshowFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowFee() string {
	return r._noshowFee
}

// Set is NoshowCurrency Setter
// NOSHOW币种,可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowCurrency(_noshowCurrency string) error {
	r._noshowCurrency = _noshowCurrency
	r.Set("noshowCurrency", _noshowCurrency)
	return nil
}

// Get NoshowCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowCurrency() string {
	return r._noshowCurrency
}

// Set is Farebasis Setter
// 运价基础，最大长度8
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFarebasis(_farebasis string) error {
	r._farebasis = _farebasis
	r.Set("farebasis", _farebasis)
	return nil
}

// Get Farebasis Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFarebasis() string {
	return r._farebasis
}

// Set is FareTypeCode Setter
// 运价类型，最大长度3
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareTypeCode(_fareTypeCode string) error {
	r._fareTypeCode = _fareTypeCode
	r.Set("fareTypeCode", _fareTypeCode)
	return nil
}

// Get FareTypeCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareTypeCode() string {
	return r._fareTypeCode
}

// Set is Tariff Setter
// 运价tariff，最大长度3
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTariff(_tariff string) error {
	r._tariff = _tariff
	r.Set("tariff", _tariff)
	return nil
}

// Get Tariff Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTariff() string {
	return r._tariff
}

// Set is RuleId Setter
// 运价规则id，最大长度4
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRuleId(_ruleId string) error {
	r._ruleId = _ruleId
	r.Set("ruleId", _ruleId)
	return nil
}

// Get RuleId Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRuleId() string {
	return r._ruleId
}

// Set is FareDirectDestrict Setter
// 运价组合适用方向,0(或者字段不存在):不限制/1:仅作用在去程/2:仅作用在回程
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareDirectDestrict(_fareDirectDestrict int64) error {
	r._fareDirectDestrict = _fareDirectDestrict
	r.Set("fareDirectDestrict", _fareDirectDestrict)
	return nil
}

// Get FareDirectDestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareDirectDestrict() int64 {
	return r._fareDirectDestrict
}
