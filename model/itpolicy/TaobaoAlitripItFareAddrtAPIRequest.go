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
	_isRt string
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
	_validDate4dep string
	// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
	_excludeDateRange4dep string
	// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripDatePoint4dep string
	// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripExcludeDatePoint4dep string
	// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
	_flightDateRestrict4dep string
	// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_flightDatePoint4dep string
	// 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
	_validDate4ret string
	// 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
	_excludeDateRange4ret string
	// 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripDatePoint4ret string
	// 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripExcludeDatePoint4ret string
	// 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
	_flightDateRestrict4ret string
	// 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_flightDatePoint4ret string
	// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
	_saleDate string
	// 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
	_minStay string
	// 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
	_maxStay string
	// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
	_adultPassengerIdentity string
	// （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
	_gv2childRule string
	// 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_nationality string
	// 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_excludeNationality string
	// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
	_passengerAge string
	// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
	_childPrice string
	// 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
	_rtCommissionFormula string
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
	_isCanRefund4dep string
	// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
	_refundPrice4dep string
	// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
	_refundPartPrice4dep string
	// （已废除字段）,回程全部未使用可否退票，录入是或否
	_isCanRefund4ret string
	// （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
	_refundPrice4ret string
	// （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
	_refundPartPrice4ret string
	// （已废除字段）,去程全部未使用可否改期，录入是或否
	_isCanReissue4dep string
	// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
	_reissuePrice4dep string
	// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
	_reissuePartPrice4dep string
	// （已废除字段）,回程全部未使用可否改期，录入是或否
	_isCanReissue4ret string
	// （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
	_reissuePrice4ret string
	// （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
	_reissuePartPrice4ret string
	// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
	_isNoShowCanRefund4dep string
	// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
	_isNoShowCanReissue4dep string
	// （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
	_isNoShowCanRefund4ret string
	// （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
	_isNoShowCanReissue4ret string
	// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
	_luggageRule4dep string
	// （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
	_luggageRule4ret string
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
	// 最小出行人数,数字1-9
	_minTravelPerson int64
	// 最大出行人数,数字1-9
	_maxTravelPerson int64
	// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
	_ticketPrice int64
	// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
	_adultTax int64
	// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
	_childTax int64
	// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
	_returnPoint float64
	// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
	_adjustMoney int64
	// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
	_earlyTicketingTimeLimit int64
	// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
	_lateTicketingTimeLimit int64
	// （已废除字段）,去程NOSHOW规定时限，输入正整数
	_noShowTimeLimit4dep int64
	// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
	_noShowPenalty4dep int64
	// （已废除字段）,回程NOSHOW规定时限，输入正整数
	_noShowTimeLimit4ret int64
	// （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
	_noShowPenalty4ret int64
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
func (r TaobaoAlitripItFareAddrtAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutFileCode is OutFileCode Setter
// 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOutFileCode(_outFileCode string) error {
	r._outFileCode = _outFileCode
	r.Set("outFileCode", _outFileCode)
	return nil
}

// GetOutFileCode OutFileCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOutFileCode() string {
	return r._outFileCode
}

// SetFileCode is FileCode Setter
// 文件编号
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFileCode(_fileCode string) error {
	r._fileCode = _fileCode
	r.Set("fileCode", _fileCode)
	return nil
}

// GetFileCode FileCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFileCode() string {
	return r._fileCode
}

// SetProductType is ProductType Setter
// （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetProductType(_productType string) error {
	r._productType = _productType
	r.Set("productType", _productType)
	return nil
}

// GetProductType ProductType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetProductType() string {
	return r._productType
}

// SetStockMode is StockMode Setter
// （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetStockMode(_stockMode string) error {
	r._stockMode = _stockMode
	r.Set("stockMode", _stockMode)
	return nil
}

// GetStockMode StockMode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetStockMode() string {
	return r._stockMode
}

// SetIsRt is IsRt Setter
// 是否1/2RT，1、请填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsRt(_isRt string) error {
	r._isRt = _isRt
	r.Set("isRT", _isRt)
	return nil
}

// GetIsRt IsRt Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsRt() string {
	return r._isRt
}

// SetRtType is RtType Setter
// （后期字段，预留）,1/2RT类型，当需要多填入多个时，请以&#34;,&#34;分隔 1、可填写 、旅行有效期、排除旅行有效期、班期 ；表明1/2RT 混舱计算时，取严还是各取各 2、默认值是 全部各取各
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRtType(_rtType string) error {
	r._rtType = _rtType
	r.Set("rtType", _rtType)
	return nil
}

// GetRtType RtType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRtType() string {
	return r._rtType
}

// SetCombinationFilecode is CombinationFilecode Setter
// 可组文件编号， 当需要多填入多个时，请以&#34;,&#34;分隔 1、标记可组文件的编号政策信息，可填写空白； 2、如果是否1/2RT 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCombinationFilecode(_combinationFilecode string) error {
	r._combinationFilecode = _combinationFilecode
	r.Set("combinationFilecode", _combinationFilecode)
	return nil
}

// GetCombinationFilecode CombinationFilecode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCombinationFilecode() string {
	return r._combinationFilecode
}

// SetIsAllowOj is IsAllowOj Setter
// （后期字段，预留）,是否允许缺口，1、为是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsAllowOj(_isAllowOj string) error {
	r._isAllowOj = _isAllowOj
	r.Set("isAllowOj", _isAllowOj)
	return nil
}

// GetIsAllowOj IsAllowOj Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsAllowOj() string {
	return r._isAllowOj
}

// SetOjType is OjType Setter
// （后期字段，预留）,缺口类型，1、可填单缺、双缺、始发地缺、目的地缺、或为空；默认为空（当允许缺口组合时，此项为必输项）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOjType(_ojType string) error {
	r._ojType = _ojType
	r.Set("ojType", _ojType)
	return nil
}

// GetOjType OjType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOjType() string {
	return r._ojType
}

// SetCombinationOjFilecode is CombinationOjFilecode Setter
// （后期字段，预留）,可组缺口文件编号,当需要多填入多个时，请以&#34;,&#34;分隔 1、标记政策信息，可填写空白； 2、如果是否缺口 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCombinationOjFilecode(_combinationOjFilecode string) error {
	r._combinationOjFilecode = _combinationOjFilecode
	r.Set("combinationOjFilecode", _combinationOjFilecode)
	return nil
}

// GetCombinationOjFilecode CombinationOjFilecode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCombinationOjFilecode() string {
	return r._combinationOjFilecode
}

// SetTicketingAirline is TicketingAirline Setter
// 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTicketingAirline(_ticketingAirline string) error {
	r._ticketingAirline = _ticketingAirline
	r.Set("ticketingAirline", _ticketingAirline)
	return nil
}

// GetTicketingAirline TicketingAirline Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTicketingAirline() string {
	return r._ticketingAirline
}

// SetSaleAirline is SaleAirline Setter
// 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetSaleAirline(_saleAirline string) error {
	r._saleAirline = _saleAirline
	r.Set("saleAirline", _saleAirline)
	return nil
}

// GetSaleAirline SaleAirline Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetSaleAirline() string {
	return r._saleAirline
}

// SetAddressOption is AddressOption Setter
// 城市/机场选项，默认为城市1、可以填写：“机场&#34;,“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAddressOption(_addressOption string) error {
	r._addressOption = _addressOption
	r.Set("addressOption", _addressOption)
	return nil
}

// GetAddressOption AddressOption Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAddressOption() string {
	return r._addressOption
}

// SetTripType is TripType Setter
// 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripType(_tripType string) error {
	r._tripType = _tripType
	r.Set("tripType", _tripType)
	return nil
}

// GetTripType TripType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripType() string {
	return r._tripType
}

// SetOriginLand is OriginLand Setter
// 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetOriginLand(_originLand string) error {
	r._originLand = _originLand
	r.Set("originLand", _originLand)
	return nil
}

// GetOriginLand OriginLand Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetOriginLand() string {
	return r._originLand
}

// SetDestination is Destination Setter
// 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDestination(_destination string) error {
	r._destination = _destination
	r.Set("destination", _destination)
	return nil
}

// GetDestination Destination Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDestination() string {
	return r._destination
}

// SetTransitLand is TransitLand Setter
// 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTransitLand(_transitLand string) error {
	r._transitLand = _transitLand
	r.Set("transitLand", _transitLand)
	return nil
}

// GetTransitLand TransitLand Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTransitLand() string {
	return r._transitLand
}

// SetCabin is Cabin Setter
// 舱位， 用&#34;,&#34;表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCabin(_cabin string) error {
	r._cabin = _cabin
	r.Set("cabin", _cabin)
	return nil
}

// GetCabin Cabin Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCabin() string {
	return r._cabin
}

// SetRestrictFlightNo is RestrictFlightNo Setter
// 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRestrictFlightNo(_restrictFlightNo string) error {
	r._restrictFlightNo = _restrictFlightNo
	r.Set("restrictFlightNo", _restrictFlightNo)
	return nil
}

// GetRestrictFlightNo RestrictFlightNo Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRestrictFlightNo() string {
	return r._restrictFlightNo
}

// SetExcludeFlightNo is ExcludeFlightNo Setter
// 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeFlightNo(_excludeFlightNo string) error {
	r._excludeFlightNo = _excludeFlightNo
	r.Set("excludeFlightNo", _excludeFlightNo)
	return nil
}

// GetExcludeFlightNo ExcludeFlightNo Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeFlightNo() string {
	return r._excludeFlightNo
}

// SetValidDate4dep is ValidDate4dep Setter
// 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetValidDate4dep(_validDate4dep string) error {
	r._validDate4dep = _validDate4dep
	r.Set("validDate4Dep", _validDate4dep)
	return nil
}

// GetValidDate4dep ValidDate4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetValidDate4dep() string {
	return r._validDate4dep
}

// SetExcludeDateRange4dep is ExcludeDateRange4dep Setter
// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeDateRange4dep(_excludeDateRange4dep string) error {
	r._excludeDateRange4dep = _excludeDateRange4dep
	r.Set("excludeDateRange4Dep", _excludeDateRange4dep)
	return nil
}

// GetExcludeDateRange4dep ExcludeDateRange4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeDateRange4dep() string {
	return r._excludeDateRange4dep
}

// SetTripDatePoint4dep is TripDatePoint4dep Setter
// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripDatePoint4dep(_tripDatePoint4dep string) error {
	r._tripDatePoint4dep = _tripDatePoint4dep
	r.Set("tripDatePoint4Dep", _tripDatePoint4dep)
	return nil
}

// GetTripDatePoint4dep TripDatePoint4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripDatePoint4dep() string {
	return r._tripDatePoint4dep
}

// SetTripExcludeDatePoint4dep is TripExcludeDatePoint4dep Setter
// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripExcludeDatePoint4dep(_tripExcludeDatePoint4dep string) error {
	r._tripExcludeDatePoint4dep = _tripExcludeDatePoint4dep
	r.Set("tripExcludeDatePoint4Dep", _tripExcludeDatePoint4dep)
	return nil
}

// GetTripExcludeDatePoint4dep TripExcludeDatePoint4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripExcludeDatePoint4dep() string {
	return r._tripExcludeDatePoint4dep
}

// SetFlightDateRestrict4dep is FlightDateRestrict4dep Setter
// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDateRestrict4dep(_flightDateRestrict4dep string) error {
	r._flightDateRestrict4dep = _flightDateRestrict4dep
	r.Set("flightDateRestrict4Dep", _flightDateRestrict4dep)
	return nil
}

// GetFlightDateRestrict4dep FlightDateRestrict4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDateRestrict4dep() string {
	return r._flightDateRestrict4dep
}

// SetFlightDatePoint4dep is FlightDatePoint4dep Setter
// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDatePoint4dep(_flightDatePoint4dep string) error {
	r._flightDatePoint4dep = _flightDatePoint4dep
	r.Set("flightDatePoint4Dep", _flightDatePoint4dep)
	return nil
}

// GetFlightDatePoint4dep FlightDatePoint4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDatePoint4dep() string {
	return r._flightDatePoint4dep
}

// SetValidDate4ret is ValidDate4ret Setter
// 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetValidDate4ret(_validDate4ret string) error {
	r._validDate4ret = _validDate4ret
	r.Set("validDate4Ret", _validDate4ret)
	return nil
}

// GetValidDate4ret ValidDate4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetValidDate4ret() string {
	return r._validDate4ret
}

// SetExcludeDateRange4ret is ExcludeDateRange4ret Setter
// 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeDateRange4ret(_excludeDateRange4ret string) error {
	r._excludeDateRange4ret = _excludeDateRange4ret
	r.Set("excludeDateRange4Ret", _excludeDateRange4ret)
	return nil
}

// GetExcludeDateRange4ret ExcludeDateRange4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeDateRange4ret() string {
	return r._excludeDateRange4ret
}

// SetTripDatePoint4ret is TripDatePoint4ret Setter
// 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripDatePoint4ret(_tripDatePoint4ret string) error {
	r._tripDatePoint4ret = _tripDatePoint4ret
	r.Set("tripDatePoint4Ret", _tripDatePoint4ret)
	return nil
}

// GetTripDatePoint4ret TripDatePoint4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripDatePoint4ret() string {
	return r._tripDatePoint4ret
}

// SetTripExcludeDatePoint4ret is TripExcludeDatePoint4ret Setter
// 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTripExcludeDatePoint4ret(_tripExcludeDatePoint4ret string) error {
	r._tripExcludeDatePoint4ret = _tripExcludeDatePoint4ret
	r.Set("tripExcludeDatePoint4Ret", _tripExcludeDatePoint4ret)
	return nil
}

// GetTripExcludeDatePoint4ret TripExcludeDatePoint4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTripExcludeDatePoint4ret() string {
	return r._tripExcludeDatePoint4ret
}

// SetFlightDateRestrict4ret is FlightDateRestrict4ret Setter
// 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDateRestrict4ret(_flightDateRestrict4ret string) error {
	r._flightDateRestrict4ret = _flightDateRestrict4ret
	r.Set("flightDateRestrict4Ret", _flightDateRestrict4ret)
	return nil
}

// GetFlightDateRestrict4ret FlightDateRestrict4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDateRestrict4ret() string {
	return r._flightDateRestrict4ret
}

// SetFlightDatePoint4ret is FlightDatePoint4ret Setter
// 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFlightDatePoint4ret(_flightDatePoint4ret string) error {
	r._flightDatePoint4ret = _flightDatePoint4ret
	r.Set("flightDatePoint4Ret", _flightDatePoint4ret)
	return nil
}

// GetFlightDatePoint4ret FlightDatePoint4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFlightDatePoint4ret() string {
	return r._flightDatePoint4ret
}

// SetSaleDate is SaleDate Setter
// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetSaleDate(_saleDate string) error {
	r._saleDate = _saleDate
	r.Set("saleDate", _saleDate)
	return nil
}

// GetSaleDate SaleDate Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetSaleDate() string {
	return r._saleDate
}

// SetMinStay is MinStay Setter
// 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMinStay(_minStay string) error {
	r._minStay = _minStay
	r.Set("minStay", _minStay)
	return nil
}

// GetMinStay MinStay Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMinStay() string {
	return r._minStay
}

// SetMaxStay is MaxStay Setter
// 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMaxStay(_maxStay string) error {
	r._maxStay = _maxStay
	r.Set("maxStay", _maxStay)
	return nil
}

// GetMaxStay MaxStay Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMaxStay() string {
	return r._maxStay
}

// SetAdultPassengerIdentity is AdultPassengerIdentity Setter
// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdultPassengerIdentity(_adultPassengerIdentity string) error {
	r._adultPassengerIdentity = _adultPassengerIdentity
	r.Set("adultPassengerIdentity", _adultPassengerIdentity)
	return nil
}

// GetAdultPassengerIdentity AdultPassengerIdentity Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdultPassengerIdentity() string {
	return r._adultPassengerIdentity
}

// SetGv2childRule is Gv2childRule Setter
// （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetGv2childRule(_gv2childRule string) error {
	r._gv2childRule = _gv2childRule
	r.Set("gv2ChildRule", _gv2childRule)
	return nil
}

// GetGv2childRule Gv2childRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetGv2childRule() string {
	return r._gv2childRule
}

// SetNationality is Nationality Setter
// 国籍，可录入多个用&#34;,&#34;隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNationality(_nationality string) error {
	r._nationality = _nationality
	r.Set("nationality", _nationality)
	return nil
}

// GetNationality Nationality Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNationality() string {
	return r._nationality
}

// SetExcludeNationality is ExcludeNationality Setter
// 除外国籍，可录入多个用&#34;,&#34;隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExcludeNationality(_excludeNationality string) error {
	r._excludeNationality = _excludeNationality
	r.Set("excludeNationality", _excludeNationality)
	return nil
}

// GetExcludeNationality ExcludeNationality Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExcludeNationality() string {
	return r._excludeNationality
}

// SetPassengerAge is PassengerAge Setter
// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetPassengerAge(_passengerAge string) error {
	r._passengerAge = _passengerAge
	r.Set("passengerAge", _passengerAge)
	return nil
}

// GetPassengerAge PassengerAge Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetPassengerAge() string {
	return r._passengerAge
}

// SetChildPrice is ChildPrice Setter
// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当&#34;乘客类型&#34;输入非“普通”（成人）时，此项输入无效。
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetChildPrice(_childPrice string) error {
	r._childPrice = _childPrice
	r.Set("childPrice", _childPrice)
	return nil
}

// GetChildPrice ChildPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetChildPrice() string {
	return r._childPrice
}

// SetRtCommissionFormula is RtCommissionFormula Setter
// 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRtCommissionFormula(_rtCommissionFormula string) error {
	r._rtCommissionFormula = _rtCommissionFormula
	r.Set("rtCommissionFormula", _rtCommissionFormula)
	return nil
}

// GetRtCommissionFormula RtCommissionFormula Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRtCommissionFormula() string {
	return r._rtCommissionFormula
}

// SetVipCode is VipCode Setter
// 大客户编码，文本框
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetVipCode(_vipCode string) error {
	r._vipCode = _vipCode
	r.Set("vipCode", _vipCode)
	return nil
}

// GetVipCode VipCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetVipCode() string {
	return r._vipCode
}

// SetFareSource is FareSource Setter
// （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareSource(_fareSource string) error {
	r._fareSource = _fareSource
	r.Set("fareSource", _fareSource)
	return nil
}

// GetFareSource FareSource Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareSource() string {
	return r._fareSource
}

// SetIsCreatePnr is IsCreatePnr Setter
// （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCreatePnr(_isCreatePnr string) error {
	r._isCreatePnr = _isCreatePnr
	r.Set("isCreatePnr", _isCreatePnr)
	return nil
}

// GetIsCreatePnr IsCreatePnr Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCreatePnr() string {
	return r._isCreatePnr
}

// SetBookingOffice is BookingOffice Setter
// 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetBookingOffice(_bookingOffice string) error {
	r._bookingOffice = _bookingOffice
	r.Set("bookingOffice", _bookingOffice)
	return nil
}

// GetBookingOffice BookingOffice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetBookingOffice() string {
	return r._bookingOffice
}

// SetReceipts is Receipts Setter
// 必填项 赋值范围 境外电子凭证,旅行发票,差额行程单发票,等额行程单
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReceipts(_receipts string) error {
	r._receipts = _receipts
	r.Set("receipts", _receipts)
	return nil
}

// GetReceipts Receipts Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReceipts() string {
	return r._receipts
}

// SetIsValidatPrice is IsValidatPrice Setter
// 是否校验票面价,1、可填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsValidatPrice(_isValidatPrice string) error {
	r._isValidatPrice = _isValidatPrice
	r.Set("isValidatPrice", _isValidatPrice)
	return nil
}

// GetIsValidatPrice IsValidatPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsValidatPrice() string {
	return r._isValidatPrice
}

// SetIsCanRefund4dep is IsCanRefund4dep Setter
// （已废除字段）,去程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanRefund4dep(_isCanRefund4dep string) error {
	r._isCanRefund4dep = _isCanRefund4dep
	r.Set("isCanRefund4Dep", _isCanRefund4dep)
	return nil
}

// GetIsCanRefund4dep IsCanRefund4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanRefund4dep() string {
	return r._isCanRefund4dep
}

// SetRefundPrice4dep is RefundPrice4dep Setter
// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPrice4dep(_refundPrice4dep string) error {
	r._refundPrice4dep = _refundPrice4dep
	r.Set("refundPrice4Dep", _refundPrice4dep)
	return nil
}

// GetRefundPrice4dep RefundPrice4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPrice4dep() string {
	return r._refundPrice4dep
}

// SetRefundPartPrice4dep is RefundPartPrice4dep Setter
// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPartPrice4dep(_refundPartPrice4dep string) error {
	r._refundPartPrice4dep = _refundPartPrice4dep
	r.Set("refundPartPrice4Dep", _refundPartPrice4dep)
	return nil
}

// GetRefundPartPrice4dep RefundPartPrice4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPartPrice4dep() string {
	return r._refundPartPrice4dep
}

// SetIsCanRefund4ret is IsCanRefund4ret Setter
// （已废除字段）,回程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanRefund4ret(_isCanRefund4ret string) error {
	r._isCanRefund4ret = _isCanRefund4ret
	r.Set("isCanRefund4Ret", _isCanRefund4ret)
	return nil
}

// GetIsCanRefund4ret IsCanRefund4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanRefund4ret() string {
	return r._isCanRefund4ret
}

// SetRefundPrice4ret is RefundPrice4ret Setter
// （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPrice4ret(_refundPrice4ret string) error {
	r._refundPrice4ret = _refundPrice4ret
	r.Set("refundPrice4Ret", _refundPrice4ret)
	return nil
}

// GetRefundPrice4ret RefundPrice4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPrice4ret() string {
	return r._refundPrice4ret
}

// SetRefundPartPrice4ret is RefundPartPrice4ret Setter
// （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundPartPrice4ret(_refundPartPrice4ret string) error {
	r._refundPartPrice4ret = _refundPartPrice4ret
	r.Set("refundPartPrice4Ret", _refundPartPrice4ret)
	return nil
}

// GetRefundPartPrice4ret RefundPartPrice4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundPartPrice4ret() string {
	return r._refundPartPrice4ret
}

// SetIsCanReissue4dep is IsCanReissue4dep Setter
// （已废除字段）,去程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanReissue4dep(_isCanReissue4dep string) error {
	r._isCanReissue4dep = _isCanReissue4dep
	r.Set("isCanReissue4Dep", _isCanReissue4dep)
	return nil
}

// GetIsCanReissue4dep IsCanReissue4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanReissue4dep() string {
	return r._isCanReissue4dep
}

// SetReissuePrice4dep is ReissuePrice4dep Setter
// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePrice4dep(_reissuePrice4dep string) error {
	r._reissuePrice4dep = _reissuePrice4dep
	r.Set("reissuePrice4Dep", _reissuePrice4dep)
	return nil
}

// GetReissuePrice4dep ReissuePrice4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePrice4dep() string {
	return r._reissuePrice4dep
}

// SetReissuePartPrice4dep is ReissuePartPrice4dep Setter
// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePartPrice4dep(_reissuePartPrice4dep string) error {
	r._reissuePartPrice4dep = _reissuePartPrice4dep
	r.Set("reissuePartPrice4Dep", _reissuePartPrice4dep)
	return nil
}

// GetReissuePartPrice4dep ReissuePartPrice4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePartPrice4dep() string {
	return r._reissuePartPrice4dep
}

// SetIsCanReissue4ret is IsCanReissue4ret Setter
// （已废除字段）,回程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanReissue4ret(_isCanReissue4ret string) error {
	r._isCanReissue4ret = _isCanReissue4ret
	r.Set("isCanReissue4Ret", _isCanReissue4ret)
	return nil
}

// GetIsCanReissue4ret IsCanReissue4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanReissue4ret() string {
	return r._isCanReissue4ret
}

// SetReissuePrice4ret is ReissuePrice4ret Setter
// （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePrice4ret(_reissuePrice4ret string) error {
	r._reissuePrice4ret = _reissuePrice4ret
	r.Set("reissuePrice4Ret", _reissuePrice4ret)
	return nil
}

// GetReissuePrice4ret ReissuePrice4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePrice4ret() string {
	return r._reissuePrice4ret
}

// SetReissuePartPrice4ret is ReissuePartPrice4ret Setter
// （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissuePartPrice4ret(_reissuePartPrice4ret string) error {
	r._reissuePartPrice4ret = _reissuePartPrice4ret
	r.Set("reissuePartPrice4Ret", _reissuePartPrice4ret)
	return nil
}

// GetReissuePartPrice4ret ReissuePartPrice4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissuePartPrice4ret() string {
	return r._reissuePartPrice4ret
}

// SetIsNoShowCanRefund4dep is IsNoShowCanRefund4dep Setter
// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanRefund4dep(_isNoShowCanRefund4dep string) error {
	r._isNoShowCanRefund4dep = _isNoShowCanRefund4dep
	r.Set("isNoShowCanRefund4Dep", _isNoShowCanRefund4dep)
	return nil
}

// GetIsNoShowCanRefund4dep IsNoShowCanRefund4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanRefund4dep() string {
	return r._isNoShowCanRefund4dep
}

// SetIsNoShowCanReissue4dep is IsNoShowCanReissue4dep Setter
// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanReissue4dep(_isNoShowCanReissue4dep string) error {
	r._isNoShowCanReissue4dep = _isNoShowCanReissue4dep
	r.Set("isNoShowCanReissue4Dep", _isNoShowCanReissue4dep)
	return nil
}

// GetIsNoShowCanReissue4dep IsNoShowCanReissue4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanReissue4dep() string {
	return r._isNoShowCanReissue4dep
}

// SetIsNoShowCanRefund4ret is IsNoShowCanRefund4ret Setter
// （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanRefund4ret(_isNoShowCanRefund4ret string) error {
	r._isNoShowCanRefund4ret = _isNoShowCanRefund4ret
	r.Set("isNoShowCanRefund4Ret", _isNoShowCanRefund4ret)
	return nil
}

// GetIsNoShowCanRefund4ret IsNoShowCanRefund4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanRefund4ret() string {
	return r._isNoShowCanRefund4ret
}

// SetIsNoShowCanReissue4ret is IsNoShowCanReissue4ret Setter
// （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsNoShowCanReissue4ret(_isNoShowCanReissue4ret string) error {
	r._isNoShowCanReissue4ret = _isNoShowCanReissue4ret
	r.Set("isNoShowCanReissue4Ret", _isNoShowCanReissue4ret)
	return nil
}

// GetIsNoShowCanReissue4ret IsNoShowCanReissue4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsNoShowCanReissue4ret() string {
	return r._isNoShowCanReissue4ret
}

// SetLuggageRule4dep is LuggageRule4dep Setter
// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用&#34;,&#34;隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule4dep(_luggageRule4dep string) error {
	r._luggageRule4dep = _luggageRule4dep
	r.Set("luggageRule4Dep", _luggageRule4dep)
	return nil
}

// GetLuggageRule4dep LuggageRule4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule4dep() string {
	return r._luggageRule4dep
}

// SetLuggageRule4ret is LuggageRule4ret Setter
// （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用&#34;,&#34;隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule4ret(_luggageRule4ret string) error {
	r._luggageRule4ret = _luggageRule4ret
	r.Set("luggageRule4Ret", _luggageRule4ret)
	return nil
}

// GetLuggageRule4ret LuggageRule4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule4ret() string {
	return r._luggageRule4ret
}

// SetRemark is Remark Setter
// 备注,出票备注文本
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRemark() string {
	return r._remark
}

// SetWorkingHours is WorkingHours Setter
// 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetWorkingHours(_workingHours string) error {
	r._workingHours = _workingHours
	r.Set("workingHours", _workingHours)
	return nil
}

// GetWorkingHours WorkingHours Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetWorkingHours() string {
	return r._workingHours
}

// SetRefundRule is RefundRule Setter
// （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundRule(_refundRule string) error {
	r._refundRule = _refundRule
	r.Set("refundRule", _refundRule)
	return nil
}

// GetRefundRule RefundRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundRule() string {
	return r._refundRule
}

// SetReissueRule is ReissueRule Setter
// （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReissueRule(_reissueRule string) error {
	r._reissueRule = _reissueRule
	r.Set("reissueRule", _reissueRule)
	return nil
}

// GetReissueRule ReissueRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReissueRule() string {
	return r._reissueRule
}

// SetNoshowRule is NoshowRule Setter
// （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRule(_noshowRule string) error {
	r._noshowRule = _noshowRule
	r.Set("noshowRule", _noshowRule)
	return nil
}

// GetNoshowRule NoshowRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRule() string {
	return r._noshowRule
}

// SetLuggageRule is LuggageRule Setter
// 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLuggageRule(_luggageRule string) error {
	r._luggageRule = _luggageRule
	r.Set("luggageRule", _luggageRule)
	return nil
}

// GetLuggageRule LuggageRule Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLuggageRule() string {
	return r._luggageRule
}

// SetApplyChannel is ApplyChannel Setter
// 运价渠道 可选listing宝贝  默认listing
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetApplyChannel(_applyChannel string) error {
	r._applyChannel = _applyChannel
	r.Set("applyChannel", _applyChannel)
	return nil
}

// GetApplyChannel ApplyChannel Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetApplyChannel() string {
	return r._applyChannel
}

// SetCommodityType is CommodityType Setter
// 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCommodityType(_commodityType string) error {
	r._commodityType = _commodityType
	r.Set("commodityType", _commodityType)
	return nil
}

// GetCommodityType CommodityType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCommodityType() string {
	return r._commodityType
}

// SetCodeSharingType is CodeSharingType Setter
// 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCodeSharingType(_codeSharingType string) error {
	r._codeSharingType = _codeSharingType
	r.Set("codeSharingType", _codeSharingType)
	return nil
}

// GetCodeSharingType CodeSharingType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCodeSharingType() string {
	return r._codeSharingType
}

// SetExtendAttributes is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetBuyTicketNotice is BuyTicketNotice Setter
// 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetBuyTicketNotice(_buyTicketNotice string) error {
	r._buyTicketNotice = _buyTicketNotice
	r.Set("buyTicketNotice", _buyTicketNotice)
	return nil
}

// GetBuyTicketNotice BuyTicketNotice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetBuyTicketNotice() string {
	return r._buyTicketNotice
}

// SetIsCanAllRefund is IsCanAllRefund Setter
// 必填项，全部未使用可否退票，可输入:是，否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanAllRefund(_isCanAllRefund string) error {
	r._isCanAllRefund = _isCanAllRefund
	r.Set("isCanAllRefund", _isCanAllRefund)
	return nil
}

// GetIsCanAllRefund IsCanAllRefund Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanAllRefund() string {
	return r._isCanAllRefund
}

// SetRefundFeeAllUnused is RefundFeeAllUnused Setter
// 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeAllUnused(_refundFeeAllUnused string) error {
	r._refundFeeAllUnused = _refundFeeAllUnused
	r.Set("refundFeeAllUnused", _refundFeeAllUnused)
	return nil
}

// GetRefundFeeAllUnused RefundFeeAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeAllUnused() string {
	return r._refundFeeAllUnused
}

// SetRefundCurrencyAllUnused is RefundCurrencyAllUnused Setter
// 全部未使用退票币种，只能录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundCurrencyAllUnused(_refundCurrencyAllUnused string) error {
	r._refundCurrencyAllUnused = _refundCurrencyAllUnused
	r.Set("refundCurrencyAllUnused", _refundCurrencyAllUnused)
	return nil
}

// GetRefundCurrencyAllUnused RefundCurrencyAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundCurrencyAllUnused() string {
	return r._refundCurrencyAllUnused
}

// SetRefundFeeTypeAllUnused is RefundFeeTypeAllUnused Setter
// 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeTypeAllUnused(_refundFeeTypeAllUnused string) error {
	r._refundFeeTypeAllUnused = _refundFeeTypeAllUnused
	r.Set("refundFeeTypeAllUnused", _refundFeeTypeAllUnused)
	return nil
}

// GetRefundFeeTypeAllUnused RefundFeeTypeAllUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeTypeAllUnused() string {
	return r._refundFeeTypeAllUnused
}

// SetIsCanPartRefund is IsCanPartRefund Setter
// 必填项，部分未使用可否退票，可输入：是，否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetIsCanPartRefund(_isCanPartRefund string) error {
	r._isCanPartRefund = _isCanPartRefund
	r.Set("isCanPartRefund", _isCanPartRefund)
	return nil
}

// GetIsCanPartRefund IsCanPartRefund Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetIsCanPartRefund() string {
	return r._isCanPartRefund
}

// SetRefundFeePartUnused is RefundFeePartUnused Setter
// 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeePartUnused(_refundFeePartUnused string) error {
	r._refundFeePartUnused = _refundFeePartUnused
	r.Set("refundFeePartUnused", _refundFeePartUnused)
	return nil
}

// GetRefundFeePartUnused RefundFeePartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeePartUnused() string {
	return r._refundFeePartUnused
}

// SetRefundCurrencyPartUnused is RefundCurrencyPartUnused Setter
// 部分未使用退票币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundCurrencyPartUnused(_refundCurrencyPartUnused string) error {
	r._refundCurrencyPartUnused = _refundCurrencyPartUnused
	r.Set("refundCurrencyPartUnused", _refundCurrencyPartUnused)
	return nil
}

// GetRefundCurrencyPartUnused RefundCurrencyPartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundCurrencyPartUnused() string {
	return r._refundCurrencyPartUnused
}

// SetRefundFeeTypePartUnused is RefundFeeTypePartUnused Setter
// 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRefundFeeTypePartUnused(_refundFeeTypePartUnused string) error {
	r._refundFeeTypePartUnused = _refundFeeTypePartUnused
	r.Set("refundFeeTypePartUnused", _refundFeeTypePartUnused)
	return nil
}

// GetRefundFeeTypePartUnused RefundFeeTypePartUnused Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRefundFeeTypePartUnused() string {
	return r._refundFeeTypePartUnused
}

// SetCanDepChange is CanDepChange Setter
// 必填项，去程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCanDepChange(_canDepChange string) error {
	r._canDepChange = _canDepChange
	r.Set("canDepChange", _canDepChange)
	return nil
}

// GetCanDepChange CanDepChange Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCanDepChange() string {
	return r._canDepChange
}

// SetDepChangeFee is DepChangeFee Setter
// 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeFee(_depChangeFee string) error {
	r._depChangeFee = _depChangeFee
	r.Set("depChangeFee", _depChangeFee)
	return nil
}

// GetDepChangeFee DepChangeFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeFee() string {
	return r._depChangeFee
}

// SetDepChangeCurrency is DepChangeCurrency Setter
// 去程改期币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeCurrency(_depChangeCurrency string) error {
	r._depChangeCurrency = _depChangeCurrency
	r.Set("depChangeCurrency", _depChangeCurrency)
	return nil
}

// GetDepChangeCurrency DepChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeCurrency() string {
	return r._depChangeCurrency
}

// SetDepChangeFeeType is DepChangeFeeType Setter
// 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetDepChangeFeeType(_depChangeFeeType string) error {
	r._depChangeFeeType = _depChangeFeeType
	r.Set("depChangeFeeType", _depChangeFeeType)
	return nil
}

// GetDepChangeFeeType DepChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetDepChangeFeeType() string {
	return r._depChangeFeeType
}

// SetCanRetChange is CanRetChange Setter
// 必填项，回程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetCanRetChange(_canRetChange string) error {
	r._canRetChange = _canRetChange
	r.Set("canRetChange", _canRetChange)
	return nil
}

// GetCanRetChange CanRetChange Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetCanRetChange() string {
	return r._canRetChange
}

// SetRetChangeFee is RetChangeFee Setter
// 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeFee(_retChangeFee string) error {
	r._retChangeFee = _retChangeFee
	r.Set("retChangeFee", _retChangeFee)
	return nil
}

// GetRetChangeFee RetChangeFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeFee() string {
	return r._retChangeFee
}

// SetRetChangeCurrency is RetChangeCurrency Setter
// 回程改期币种，可输入币种三字码，默认值CN
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeCurrency(_retChangeCurrency string) error {
	r._retChangeCurrency = _retChangeCurrency
	r.Set("retChangeCurrency", _retChangeCurrency)
	return nil
}

// GetRetChangeCurrency RetChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeCurrency() string {
	return r._retChangeCurrency
}

// SetRetChangeFeeType is RetChangeFeeType Setter
// 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRetChangeFeeType(_retChangeFeeType string) error {
	r._retChangeFeeType = _retChangeFeeType
	r.Set("retChangeFeeType", _retChangeFeeType)
	return nil
}

// GetRetChangeFeeType RetChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRetChangeFeeType() string {
	return r._retChangeFeeType
}

// SetNoshowRestrict is NoshowRestrict Setter
// 必填项，NOSHOW是否有限制，可输入是或否
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRestrict(_noshowRestrict string) error {
	r._noshowRestrict = _noshowRestrict
	r.Set("noshowRestrict", _noshowRestrict)
	return nil
}

// GetNoshowRestrict NoshowRestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRestrict() string {
	return r._noshowRestrict
}

// SetNoshowTimeRestrict is NoshowTimeRestrict Setter
// NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowTimeRestrict(_noshowTimeRestrict string) error {
	r._noshowTimeRestrict = _noshowTimeRestrict
	r.Set("noshowTimeRestrict", _noshowTimeRestrict)
	return nil
}

// GetNoshowTimeRestrict NoshowTimeRestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowTimeRestrict() string {
	return r._noshowTimeRestrict
}

// SetNoshowTimeRestrictUnit is NoshowTimeRestrictUnit Setter
// NOSHOW时限单位(小时/天, 默认为小时)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowTimeRestrictUnit(_noshowTimeRestrictUnit string) error {
	r._noshowTimeRestrictUnit = _noshowTimeRestrictUnit
	r.Set("noshowTimeRestrictUnit", _noshowTimeRestrictUnit)
	return nil
}

// GetNoshowTimeRestrictUnit NoshowTimeRestrictUnit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowTimeRestrictUnit() string {
	return r._noshowTimeRestrictUnit
}

// SetNoshowRuleType is NoshowRuleType Setter
// NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可改期,不可改期
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowRuleType(_noshowRuleType string) error {
	r._noshowRuleType = _noshowRuleType
	r.Set("noshowRuleType", _noshowRuleType)
	return nil
}

// GetNoshowRuleType NoshowRuleType Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowRuleType() string {
	return r._noshowRuleType
}

// SetNoshowFee is NoshowFee Setter
// NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowFee(_noshowFee string) error {
	r._noshowFee = _noshowFee
	r.Set("noshowFee", _noshowFee)
	return nil
}

// GetNoshowFee NoshowFee Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowFee() string {
	return r._noshowFee
}

// SetNoshowCurrency is NoshowCurrency Setter
// NOSHOW币种,可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoshowCurrency(_noshowCurrency string) error {
	r._noshowCurrency = _noshowCurrency
	r.Set("noshowCurrency", _noshowCurrency)
	return nil
}

// GetNoshowCurrency NoshowCurrency Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoshowCurrency() string {
	return r._noshowCurrency
}

// SetFarebasis is Farebasis Setter
// 运价基础，最大长度8
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFarebasis(_farebasis string) error {
	r._farebasis = _farebasis
	r.Set("farebasis", _farebasis)
	return nil
}

// GetFarebasis Farebasis Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFarebasis() string {
	return r._farebasis
}

// SetFareTypeCode is FareTypeCode Setter
// 运价类型，最大长度3
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareTypeCode(_fareTypeCode string) error {
	r._fareTypeCode = _fareTypeCode
	r.Set("fareTypeCode", _fareTypeCode)
	return nil
}

// GetFareTypeCode FareTypeCode Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareTypeCode() string {
	return r._fareTypeCode
}

// SetTariff is Tariff Setter
// 运价tariff，最大长度3
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTariff(_tariff string) error {
	r._tariff = _tariff
	r.Set("tariff", _tariff)
	return nil
}

// GetTariff Tariff Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTariff() string {
	return r._tariff
}

// SetRuleId is RuleId Setter
// 运价规则id，最大长度4
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetRuleId(_ruleId string) error {
	r._ruleId = _ruleId
	r.Set("ruleId", _ruleId)
	return nil
}

// GetRuleId RuleId Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetRuleId() string {
	return r._ruleId
}

// SetMinTravelPerson is MinTravelPerson Setter
// 最小出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMinTravelPerson(_minTravelPerson int64) error {
	r._minTravelPerson = _minTravelPerson
	r.Set("minTravelPerson", _minTravelPerson)
	return nil
}

// GetMinTravelPerson MinTravelPerson Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMinTravelPerson() int64 {
	return r._minTravelPerson
}

// SetMaxTravelPerson is MaxTravelPerson Setter
// 最大出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetMaxTravelPerson(_maxTravelPerson int64) error {
	r._maxTravelPerson = _maxTravelPerson
	r.Set("maxTravelPerson", _maxTravelPerson)
	return nil
}

// GetMaxTravelPerson MaxTravelPerson Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetMaxTravelPerson() int64 {
	return r._maxTravelPerson
}

// SetTicketPrice is TicketPrice Setter
// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetTicketPrice(_ticketPrice int64) error {
	r._ticketPrice = _ticketPrice
	r.Set("ticketPrice", _ticketPrice)
	return nil
}

// GetTicketPrice TicketPrice Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetTicketPrice() int64 {
	return r._ticketPrice
}

// SetAdultTax is AdultTax Setter
// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdultTax(_adultTax int64) error {
	r._adultTax = _adultTax
	r.Set("adultTax", _adultTax)
	return nil
}

// GetAdultTax AdultTax Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdultTax() int64 {
	return r._adultTax
}

// SetChildTax is ChildTax Setter
// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetChildTax(_childTax int64) error {
	r._childTax = _childTax
	r.Set("childTax", _childTax)
	return nil
}

// GetChildTax ChildTax Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetChildTax() int64 {
	return r._childTax
}

// SetReturnPoint is ReturnPoint Setter
// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetReturnPoint(_returnPoint float64) error {
	r._returnPoint = _returnPoint
	r.Set("returnPoint", _returnPoint)
	return nil
}

// GetReturnPoint ReturnPoint Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetReturnPoint() float64 {
	return r._returnPoint
}

// SetAdjustMoney is AdjustMoney Setter
// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetAdjustMoney(_adjustMoney int64) error {
	r._adjustMoney = _adjustMoney
	r.Set("adjustMoney", _adjustMoney)
	return nil
}

// GetAdjustMoney AdjustMoney Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetAdjustMoney() int64 {
	return r._adjustMoney
}

// SetEarlyTicketingTimeLimit is EarlyTicketingTimeLimit Setter
// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetEarlyTicketingTimeLimit(_earlyTicketingTimeLimit int64) error {
	r._earlyTicketingTimeLimit = _earlyTicketingTimeLimit
	r.Set("earlyTicketingTimeLimit", _earlyTicketingTimeLimit)
	return nil
}

// GetEarlyTicketingTimeLimit EarlyTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetEarlyTicketingTimeLimit() int64 {
	return r._earlyTicketingTimeLimit
}

// SetLateTicketingTimeLimit is LateTicketingTimeLimit Setter
// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetLateTicketingTimeLimit(_lateTicketingTimeLimit int64) error {
	r._lateTicketingTimeLimit = _lateTicketingTimeLimit
	r.Set("lateTicketingTimeLimit", _lateTicketingTimeLimit)
	return nil
}

// GetLateTicketingTimeLimit LateTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetLateTicketingTimeLimit() int64 {
	return r._lateTicketingTimeLimit
}

// SetNoShowTimeLimit4dep is NoShowTimeLimit4dep Setter
// （已废除字段）,去程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowTimeLimit4dep(_noShowTimeLimit4dep int64) error {
	r._noShowTimeLimit4dep = _noShowTimeLimit4dep
	r.Set("noShowTimeLimit4Dep", _noShowTimeLimit4dep)
	return nil
}

// GetNoShowTimeLimit4dep NoShowTimeLimit4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowTimeLimit4dep() int64 {
	return r._noShowTimeLimit4dep
}

// SetNoShowPenalty4dep is NoShowPenalty4dep Setter
// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowPenalty4dep(_noShowPenalty4dep int64) error {
	r._noShowPenalty4dep = _noShowPenalty4dep
	r.Set("noShowPenalty4Dep", _noShowPenalty4dep)
	return nil
}

// GetNoShowPenalty4dep NoShowPenalty4dep Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowPenalty4dep() int64 {
	return r._noShowPenalty4dep
}

// SetNoShowTimeLimit4ret is NoShowTimeLimit4ret Setter
// （已废除字段）,回程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowTimeLimit4ret(_noShowTimeLimit4ret int64) error {
	r._noShowTimeLimit4ret = _noShowTimeLimit4ret
	r.Set("noShowTimeLimit4Ret", _noShowTimeLimit4ret)
	return nil
}

// GetNoShowTimeLimit4ret NoShowTimeLimit4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowTimeLimit4ret() int64 {
	return r._noShowTimeLimit4ret
}

// SetNoShowPenalty4ret is NoShowPenalty4ret Setter
// （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetNoShowPenalty4ret(_noShowPenalty4ret int64) error {
	r._noShowPenalty4ret = _noShowPenalty4ret
	r.Set("noShowPenalty4Ret", _noShowPenalty4ret)
	return nil
}

// GetNoShowPenalty4ret NoShowPenalty4ret Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetNoShowPenalty4ret() int64 {
	return r._noShowPenalty4ret
}

// SetFareDirectDestrict is FareDirectDestrict Setter
// 运价组合适用方向,0(或者字段不存在):不限制/1:仅作用在去程/2:仅作用在回程
func (r *TaobaoAlitripItFareAddrtAPIRequest) SetFareDirectDestrict(_fareDirectDestrict int64) error {
	r._fareDirectDestrict = _fareDirectDestrict
	r.Set("fareDirectDestrict", _fareDirectDestrict)
	return nil
}

// GetFareDirectDestrict FareDirectDestrict Getter
func (r TaobaoAlitripItFareAddrtAPIRequest) GetFareDirectDestrict() int64 {
	return r._fareDirectDestrict
}
