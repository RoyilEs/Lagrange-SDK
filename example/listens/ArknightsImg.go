package listens

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/common/limiter"
	"Lagrange-SDK/events"
	"Lagrange-SDK/global"
	"Lagrange-SDK/models/pixiv"
	"Lagrange-SDK/utils"
	"Lagrange-SDK/utils/image"
	"context"
	"encoding/base64"
	"github.com/charmbracelet/log"
	"golang.org/x/time/rate"
	"io/ioutil"
	"strings"
	"time"
)

var (
	Limit     = limiter.NewLimiter(rate.Every(1*time.Second), 2, "")
	iMainFunc = apiBuilder.New(global.BotUrl)
)

func ArknightsImg(ctx context.Context, event events.IEvent) {
	if event.GetMessageType() == string(events.EventGroupMsg) {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()

		var ZhouTest = []string{"来点粥图", "来张粥图", "粥图一张"}

		query, err := pixiv.NewPixiv().Set().DoQuery()
		if err != nil {
			log.Error(err)
			return
		}
		iPixiv, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)

		if utils.IsInListToS(text[0], ZhouTest) {
			if err != nil {
				log.Error(err)
				return
			}
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				buf := i.UrlToBase64(pixiv.ModifyPixivImageUrl(url))
				log.Info(url)
				iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).
					ImgBase64Msg(buf).Do(ctx)
			}
		}

		var ZhouTest2 = []string{"明日方舟", "来点舟图", "舟图一张"}
		if utils.IsInListToS(text[0], ZhouTest2) {
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				log.Info(url)
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				if err != nil {
					iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).TextMsg(err.Error()).Do(ctx)
					log.Error(err)
					return
				}
				log.Info(url)
				iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).
					ImgBase64Msg(encodeToBase64).Do(ctx)
			}
		}
	}
}

func PixivImg(ctx context.Context, event events.IEvent) {
	if event.GetMessageType() == string(events.EventGroupMsg) {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		split := strings.Split(text[0], " ")
		if split[0] == "pixiv" {
			// 令牌桶限流器--防止大量请求
			if !Limit.Allow() {
				iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请求过于频繁").Do(ctx)
				return
			}

			query, err := pixiv.NewPixiv().Set().SetTag(split[1]).DoQuery()
			if err != nil {
				log.Error(err)
				return
			}
			iPixiv, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)
			for _, data := range iPixiv.GetData() {
				url := data.GetDataUrls().GetSize()
				log.Info(url)
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
					ImgBase64Msg(encodeToBase64).Do(ctx)
				if err != nil {
					log.Error(err)
					return
				}
			}
		}
	}
}

const mingImgPath = "img/"

var mingImg = map[int]string{
	0: "2226", 1: "ash", 2: "刺梅", 3: "安洁莉娜", 4: "博士", 5: "提丰",
	6: "IMG_005", 7: "IMG_010", 8: "IMG_011", 9: "IMG_012", 10: "IMG_013",
	11: "IMG_014", 12: "IMG_015", 13: "IMG_016", 14: "IMG_017", 15: "IMG_018",
	16: "IMG_019", 17: "IMG_020", 18: "IMG_021", 19: "IMG_022", 20: "IMG_023",
	21: "IMG_024", 22: "IMG_025", 23: "IMG_026", 24: "IMG_027", 25: "IMG_028",
	26: "IMG_029", 27: "IMG_030", 28: "IMG_031", 29: "IMG_032", 30: "IMG_033",
	31: "IMG_034", 32: "IMG_035", 33: "IMG_036", 34: "IMG_037", 35: "IMG_038",
	36: "IMG_039", 37: "IMG_040", 38: "IMG_041", 39: "IMG_042", 40: "IMG_043",
	41: "IMG_044", 42: "IMG_045", 43: "IMG_046", 44: "IMG_047", 45: "IMG_048",
	46: "IMG_049", 47: "IMG_050", 48: "IMG_051", 49: "IMG_052", 50: "IMG_053",
	51: "IMG_054", 52: "IMG_055", 53: "IMG_056", 54: "IMG_057", 55: "IMG_058",
	56: "IMG_059", 57: "IMG_060", 58: "IMG_061", 59: "IMG_062", 60: "IMG_063",
	61: "IMG_064", 62: "IMG_065", 63: "IMG_066", 64: "IMG_067", 65: "IMG_068",
	66: "IMG_069", 67: "IMG_070", 68: "IMG_071", 69: "IMG_072", 70: "IMG_073",
	71: "IMG_074", 72: "IMG_075", 73: "IMG_076", 74: "IMG_077", 75: "IMG_078",
	76: "IMG_079", 77: "IMG_080", 78: "P脸陈", 79: "U", 80: "VVAN", 81: "埃拉托",
	82: "艾丽妮皮肤", 83: "爱丽丝", 84: "安比尔", 85: "安塞尔", 86: "安哲拉",
	87: "暗索", 88: "暗索2", 89: "白面鸮", 90: "百炼嘉维尔", 91: "薄绿",
	92: "冰酿", 93: "波登可", 94: "布丁", 95: "澄闪皮肤", 96: "澄闪泳装",
	97: "澄闪泳装", 98: "初雪", 99: "伺夜", 100: "淬羽赫默", 101: "嵯峨",
	102: "地灵", 103: "杜林", 104: "多萝西", 105: "菲亚梅塔1", 106: "芬",
	107: "风丸", 108: "风丸_001", 109: "芙蓉", 110: "歌蕾蒂娅", 111: "格劳克斯",
	112: "古米", 113: "哈洛德", 114: "海沫", 115: "寒芒克洛丝", 116: "寒檀",
	117: "和弦", 118: "赫德雷", 119: "黑角", 120: "黑角_001", 121: "黑骑士",
	122: "红豆", 123: "煌", 124: "灰喉", 125: "火哨", 126: "极境",
	127: "惊蛰", 128: "九色鹿", 129: "卡涅利安", 130: "凯尔希皮肤", 131: "克洛丝",
	132: "空爆", 133: "空构", 134: "空弦", 135: "苦艾", 136: "莱伊",
	137: "老鲤", 138: "雷蛇", 139: "砾", 140: "烈夏", 141: "林",
	142: "琳琅诗怀雅", 143: "凛冬", 144: "灵知", 145: "铃兰皮肤", 146: "流明",
	147: "流星", 148: "罗宾", 149: "洛洛1", 150: "缪尔赛思", 151: "慕斯",
	152: "慕斯2", 153: "帕拉斯", 154: "泡普卡", 155: "麒麟夜刀", 156: "绮良",
	157: "琴柳", 158: "青枳", 159: "熔泉", 160: "桑葚",
	161: "闪灵3", 162: "蛇屠箱", 163: "深律", 164: "诗怀雅",
	165: "狮蝎", 166: "史尔特尔泳装", 167: "黍", 168: "霜叶",
	169: "桃金娘", 170: "推进之王", 171: "澄闪泳装76", 172: "苇草1",
	173: "温米", 174: "巫恋", 175: "稀音", 176: "小满",
	177: "晓歌", 178: "星极", 179: "星熊", 180: "星源",
	181: "絮雨3", 182: "雪雉", 183: "崖心", 184: "亚叶",
	185: "炎熔", 186: "宴", 187: "耶拉", 188: "夜魔",
	189: "伊芙利特", 190: "伊内丝", 191: "异客", 192: "异客3",
	193: "幽灵鲨", 194: "跃跃", 195: "早露", 196: "真理",
	197: "止颂", 198: "子月", 199: "梓兰", 200: "左乐",
}

func randGetMingImg() (string, int) {
	// 获取map的长度
	length := len(mingImg)
	randomIndex := utils.Random(0, length)
	// 使用随机索引从map中获取元素
	randomElement, ok := mingImg[randomIndex]
	if ok {
		return randomElement, randomIndex
	} else {
		return "", 0
	}
}

func MingImg(ctx context.Context, event events.IEvent) {
	if event.GetMessageType() == string(events.EventGroupMsg) {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		if text[0] != "大头" {
			return
		}
		imgPath, i := randGetMingImg()

		var (
			file []byte
			err  error
		)
		if i >= 6 && i <= 77 {
			file, err = ioutil.ReadFile(mingImgPath + imgPath + ".jpg")
			if err != nil {
				log.Error(err)
				return
			}
		} else {
			file, err = ioutil.ReadFile(mingImgPath + imgPath + ".png")
			if err != nil {
				log.Error(err)
				return
			}
		}
		iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).TextMsg("获取成功---请稍等").Do(ctx)

		log.Info("大头图片：" + imgPath)

		iMainFunc.SendGroupMsg(groupMsg.GetGroupID()).
			ImgBase64Msg(base64.StdEncoding.EncodeToString(file)).Do(ctx)
	}
}
