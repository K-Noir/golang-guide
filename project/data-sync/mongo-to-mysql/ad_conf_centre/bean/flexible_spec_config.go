package bean

var (
	EducationStatusesMap = map[int]string{
		1:  "HIGH_SCHOOL",
		2:  "UNDERGRAD",
		3:  "ALUM",
		4:  "HIGH_SCHOOL_GRAD",
		5:  "SOME_COLLEGE",
		6:  "ASSOCIATE_DEGREE",
		7:  "IN_GRAD_SCHOOL",
		8:  "SOME_GRAD_SCHOOL",
		9:  "MASTER_DEGREE",
		10: "PROFESSIONAL_DEGREE",
		11: "DOCTORATE_DEGREE",
		12: "UNSPECIFIED",
		13: "SOME_HIGH_SCHOOL",
	}
	EducationStatusesMapChinese = map[int]string{
		1:  "读过高中",
		2:  "在读大学",
		3:  "大学毕业",
		4:  "在读高中",
		5:  "某所大学",
		6:  "副学士学位",
		7:  "在读研究生",
		8:  "某研究生院",
		9:  "硕士学位",
		10: "专硕学位",
		11: "博士学位",
		12: "未指定",
		13: "高中毕业",
	}
	RelationshipStatusesMap = map[int]string{
		1:  "单身",
		2:  "恋爱中",
		3:  "已婚",
		4:  "已订婚",
		6:  "未指定",
		7:  "有同性伴侣",
		8:  "有同居伴侣",
		9:  "交往中，但保留交友空间",
		10: "关系复杂",
		11: "分居",
		12: "离异",
		13: "丧偶",

		//默认值：ALL（如果指定 Null 或不提供值）。
		//限制：不能使用 0。
	}
)
