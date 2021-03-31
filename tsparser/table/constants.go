package table

var COMPONENT_TYPE map[byte]map[byte]string = map[byte]map[byte]string{
	0x01: {
		0x00: "将来使用のためリザーブ",
		0x01: "映像480i(525i)、アスペクト比4:3",
		0x02: "映像480i(525i)、アスペクト比16:9 パンベクトルあり",
		0x03: "映像480i(525i)、アスペクト比16:9 パンベクトルなし",
		0x04: "映像480i(525i)、アスペクト比16:9",
		0x91: "映像2160p、アスペクト比4:3",
		0x92: "映像2160p、アスペクト比16:9 パンベクトルあり",
		0x93: "映像2160p、アスペクト比16:9 パンベクトルなし",
		0x94: "映像2160p、アスペクト比16:9",
		0xA1: "映像480p(525p)、アスペクト比4:3",
		0xA2: "映像480p(525p)、アスペクト比16:9 パンベクトルあり",
		0xA3: "映像480p(525p)、アスペクト比16:9 パンベクトルなし",
		0xA4: "映像480p(525p)、アスペクト比16:9",
		0xB1: "映像1080i(1125i)、アスペクト比4:3",
		0xB2: "映像1080i(1125i)、アスペクト比16:9 パンベクトルあり",
		0xB3: "映像1080i(1125i)、アスペクト比16:9 パンベクトルなし",
		0xB4: "映像1080i(1125i)、アスペクト比16:9",
		0xC1: "映像720p(750p)、アスペクト比4:3",
		0xC2: "映像720p(750p)、アスペクト比16:9 パンベクトルあり",
		0xC3: "映像720p(750p)、アスペクト比16:9 パンベクトルなし",
		0xC4: "映像720p(750p)、アスペクト比16:9",
		0xD1: "映像240p アスペクト比4:3",
		0xD2: "映像240p アスペクト比16:9 パンベクトルあり",
		0xD3: "映像240p アスペクト比16:9 パンベクトルなし",
		0xD4: "映像240p アスペクト比16:9",
		0xE1: "映像1080p(1125p)、アスペクト比4:3",
		0xE2: "映像1080p(1125p)、アスペクト比16:9 パンベクトルあり",
		0xE3: "映像1080p(1125p)、アスペクト比16:9 パンベクトルなし",
		0xE4: "映像1080p(1125p)、アスペクト比16:9",
		0xF1: "映像180p アスペクト比4:3",
		0xF2: "映像180p アスペクト比16:9 パンベクトルあり",
		0xF3: "映像180p アスペクト比16:9 パンベクトルなし",
		0xF4: "映像180p アスペクト比16:9",
	},
	0x02: {
		0x00: "将来使用のためリザーブ",
		0x01: "1/0モード(シングルモノ)",
		0x02: "1/0+1/0モード(デュアルモノ)",
		0x03: "2/0モード(ステレオ)",
		0x04: "2/1モード",
		0x05: "3/0モード",
		0x06: "2/2モード",
		0x07: "3/1モード",
		0x08: "3/2モード",
		0x09: "3/2+LFEモード(3/2.1モード)",
		0x0A: "3/3.1モード",
		0x0B: "2/0/0-2/0/2-0.1モード",
		0x0C: "5/2.1モード",
		0x0D: "3/2/2.1モード",
		0x0E: "2/0/0-3/0/2-0.1モード",
		0x0F: "0/2/0-3/0/2-0.2モード",
		0x10: "2/0/0-3/2/3-0.2モード",
		0x11: "3/3/3-5/2/3-3/0/0.2モード",
		0x40: "視覚障害者用音声解説",
		0x41: "聴覚障害者用音声",
	},
	0x05: {
		0x01: "H.264|MPEG-4 AVC、 映像480i(525i)、アスペクト比4:3",
		0x02: "H.264|MPEG-4 AVC、 映像480i(525i)、アスペクト比16:9 パンベクトルあり",
		0x03: "H.264|MPEG-4 AVC、 映像480i(525i)、アスペクト比16:9 パンベクトルなし ",
		0x04: "H.264|MPEG-4 AVC、 映像480i(525i)、アスペクト比 > 16:9",
		0x91: "H.264|MPEG-4 AVC、 映像2160p、アスペクト比4:3",
		0x92: "H.264|MPEG-4 AVC、 映像2160p、アスペクト比16:9 パンベクトルあり",
		0x93: "H.264|MPEG-4 AVC、 映像2160p、アスペクト比16:9 パンベクトルなし",
		0x94: "H.264|MPEG-4 AVC、 映像2160p、アスペクト比 > 16:9",
		0xA1: "H.264|MPEG-4 AVC、 映像480p(525p)、アスペクト比4:3",
		0xA2: "H.264|MPEG-4 AVC、 映像480p(525p)、アスペクト比16:9 パンベクトルあり",
		0xA3: "H.264|MPEG-4 AVC、 映像480p(525p)、アスペクト比16:9 パンベクトルなし",
		0xA4: "H.264|MPEG-4 AVC、 映像480p(525p)、アスペクト比 > 16:9",
		0xB1: "H.264|MPEG-4 AVC、 映像1080i(1125i)、アスペクト比4:3",
		0xB2: "H.264|MPEG-4 AVC、 映像1080i(1125i)、アスペクト比16:9 パンベクトルあり",
		0xB3: "H.264|MPEG-4 AVC、 映像1080i(1125i)、アスペクト比16:9 パンベクトルなし",
		0xB4: "H.264|MPEG-4 AVC、 映像1080i(1125i)、アスペクト比 > 16:9",
		0xC1: "H.264|MPEG-4 AVC、 映像720p(750p)、アスペクト比4:3",
		0xC2: "H.264|MPEG-4 AVC、 映像720p(750p)、アスペクト比16:9 パンベクトルあり",
		0xC3: "H.264|MPEG-4 AVC、 映像720p(750p)、アスペクト比16:9 パンベクトルなし",
		0xC4: "H.264|MPEG-4 AVC、 映像720p(750p)、アスペクト比 > 16:9",
		0xD1: "H.264|MPEG-4 AVC、 映像240p アスペクト比4:3",
		0xD2: "H.264|MPEG-4 AVC、 映像240p アスペクト比16:9 パンベクトルあり",
		0xD3: "H.264|MPEG-4 AVC、 映像240p アスペクト比16:9 パンベクトルなし",
		0xD4: "H.264|MPEG-4 AVC、 映像240p アスペクト比 > 16:9",
		0xE1: "H.264|MPEG-4 AVC、 映像1080p(1125p)、アスペクト比4:3",
		0xE2: "H.264|MPEG-4 AVC、 映像1080p(1125p)、アスペクト比16:9 パンベクトルあり",
		0xE3: "H.264|MPEG-4 AVC、 映像1080p(1125p)、アスペクト比16:9 パンベクトルなし",
		0xE4: "H.264|MPEG-4 AVC、 映像1080p(1125p)、アスペクト比 > 16:9",
		0xF1: "H.264|MPEG-4 AVC、 映像180p アスペクト比4:3",
		0xF2: "H.264|MPEG-4 AVC、 映像180p アスペクト比16:9 パンベクトルあり",
		0xF3: "H.264|MPEG-4 AVC、 映像180p アスペクト比16:9 パンベクトルなし",
		0xF4: "H.264|MPEG-4 AVC、 映像180p アスペクト比 > 16:9",
	},
	0x09: {
		0x83: "H.265|MPEG-H HEVC、映像4320p、アスペクト比16:9",
		0x93: "H.265|MPEG-H HEVC、映像2160p、アスペクト比16:9",
		0xB3: "H.265|MPEG-H HEVC、映像1080i(1125i)、アスペクト比16:9",
		0xE3: "H.265|MPEG-H HEVC、映像1080p(1125p)、アスペクト比16:9",
	},
}

var CONTENT_TYPE map[byte]map[byte]string = map[byte]map[byte]string{
	0x0: {
		0xff: "ニュース／報道",
		0x0:  "定時・総合",
		0x1:  "天気",
		0x2:  "特集・ドキュメント",
		0x3:  "政治・国会",
		0x4:  "経済・市況",
		0x5:  "海外・国際",
		0x6:  "解説",
		0x7:  "討論・会談",
		0x8:  "報道特番",
		0x9:  "ローカル・地域",
		0xA:  "交通",
		0xF:  "その他",
	},
	0x1: {
		0xff: "スポーツ",
		0x0:  "スポーツニュース",
		0x1:  "野球",
		0x2:  "サッカー",
		0x3:  "ゴルフ",
		0x4:  "その他の球技",
		0x5:  "相撲・格闘技",
		0x6:  "オリンピック・国際大会",
		0x7:  "マラソン・陸上・水泳",
		0x8:  "モータースポーツ",
		0x9:  "マリン・ウィンタースポーツ",
		0xA:  "競馬・公営競技",
		0xF:  "その他",
	},
	0x2: {
		0xff: "情報／ワイドショー",
		0x0:  "芸能・ワイドショー",
		0x1:  "ファッション",
		0x2:  "暮らし・住まい",
		0x3:  "健康・医療",
		0x4:  "ショッピング・通販",
		0x5:  "グルメ・料理",
		0x6:  "イベント",
		0x7:  "番組紹介・お知らせ",
		0xF:  "その他",
	},
	0x3: {
		0xff: "ドラマ",
		0x0:  "国内ドラマ",
		0x1:  "海外ドラマ",
		0x2:  "時代劇",
		0xF:  "その他",
	},
	0x4: {
		0xff: "音楽",
		0x0:  "国内ロック・ポップス",
		0x1:  "海外ロック・ポップス",
		0x2:  "クラシック・オペラ",
		0x3:  "ジャズ・フュージョン",
		0x4:  "歌謡曲・演歌",
		0x5:  "ライブ・コンサート",
		0x6:  "ランキング・リクエスト",
		0x7:  "カラオケ・のど自慢",
		0x8:  "民謡・邦楽",
		0x9:  "童謡・キッズ",
		0xA:  "民族音楽・ワールドミュージック",
		0xF:  "その他",
	},
	0x5: {
		0xff: "バラエティ",
		0x0:  "クイズ",
		0x1:  "ゲーム",
		0x2:  "トークバラエティ",
		0x3:  "お笑い・コメディ",
		0x4:  "音楽バラエティ",
		0x5:  "旅バラエティ",
		0x6:  "料理バラエティ",
		0xF:  "その他",
	},
	0x6: {
		0xff: "映画",
		0x0:  "洋画",
		0x1:  "邦画",
		0x2:  "アニメ",
		0xF:  "その他",
	},
	0x7: {
		0xff: "アニメ／特撮",
		0x0:  "国内アニメ",
		0x1:  "海外アニメ",
		0x2:  "特撮",
		0xF:  "その他",
	},
	0x8: {
		0xff: "ドキュメンタリー／教養",
		0x0:  "社会・時事",
		0x1:  "歴史・紀行",
		0x2:  "自然・動物・環境",
		0x3:  "宇宙・科学・医学",
		0x4:  "カルチャー・伝統文化",
		0x5:  "文学・文芸",
		0x6:  "スポーツ",
		0x7:  "ドキュメンタリー全般",
		0x8:  "インタビュー・討論",
		0xF:  "その他",
	},
	0x9: {
		0xff: "劇場／公演",
		0x0:  "現代劇・新劇",
		0x1:  "ミュージカル",
		0x2:  "ダンス・バレエ",
		0x3:  "落語・演芸",
		0x4:  "歌舞伎・古典",
		0xF:  "その他",
	},
	0xA: {
		0xff: "趣味／教育",
		0x0:  "旅・釣り・アウトドア",
		0x1:  "園芸・ペット・手芸",
		0x2:  "音楽・美術・工芸",
		0x3:  "囲碁・将棋",
		0x4:  "麻雀・パチンコ",
		0x5:  "車・オートバイ",
		0x6:  "コンピュータ・ＴＶゲーム",
		0x7:  "会話・語学",
		0x8:  "幼児・小学生",
		0x9:  "中学生・高校生",
		0xA:  "大学生・受験",
		0xB:  "生涯教育・資格",
		0xC:  "教育問題",
		0xF:  "その他",
	},
	0xB: {
		0xff: "福祉",
		0x0:  "高齢者",
		0x1:  "障害者",
		0x2:  "社会福祉",
		0x3:  "ボランティア",
		0x4:  "手話",
		0x5:  "文字（字幕）",
		0x6:  "音声解説",
		0xF:  "その他",
	},
	0xE: {
		0xff: "拡張",
		0x0:  "BS/地上デジタル放送用番組付属情報",
		0x1:  "広帯域CS デジタル放送用拡張",
		0x2:  "衛星デジタル音声放送用拡張",
		0x3:  "サーバー型番組付属情報",
		0x4:  "IP 放送用番組付属情報",
	},
	0xF: {
		0xff: "その他",
		0xF:  "その他",
	},
}