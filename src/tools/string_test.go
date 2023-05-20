package tools

import (
	"fmt"
	"testing"
)

func Test_StringBuilder(t *testing.T) {
	origin := "apple,iphone,apple"
	result := StringBuilder(origin).Replace("apple", "fruit").Replace("iphone", "phone").Build()
	if result != "fruit,phone,fruit" {
		t.Error()
	}
}

func Test_KeysOrder1(t *testing.T) {
	// abc
	s := "apple,iphone,cat"
	keywords := []string{"apple", "iphone", "cat"}
	result := StringKeywordsOrderCheck(&s, keywords)
	if result == false {
		t.Error()
	}
}

func Test_KeysOrder2(t *testing.T) {
	// aba
	s := "apple,iphone,apple"
	keywords := []string{"apple", "iphone", "apple"}
	result := StringKeywordsOrderCheck(&s, keywords)
	println(result)
}
func Test_KeysOrder3(t *testing.T) {
	// aba
	s := "靳辅（catIdx=J），字紫垣。盛京辽阳人，隶汉军镶黄旗，清朝政治人物，以戮力治河而知名。== 生平 ==靳辅生于明崇祯六年（1633年）。其祖先原是山东济南人，后于明朝初年从军戍卫，从此在辽阳落户。靳辅在顺治六年（1649年）出仕，顺治九年（1652年），由官学生考授国史院编修。康熙元年（1662年），升任郎中，康熙七年，升通政使司右通政。康熙十年（1671年）授安徽巡抚，加兵部尚书，在任共6年。康熙十六年（1677年）三月，靳辅从安徽巡抚升迁为河道总督。靳辅在这一年是45岁。从这时开始，直到60岁身故时，靳辅这一段的人生岁月致力于治河。任总督时，浙江钱塘人陈潢是他的僚属。靳辅对治河本无研究，所以凡治河之事，都向陈潢请教，同陈调查黄河、淮河形势，连上八疏，以为“治河之道，必当审其全局，将河道、运道为一体，彻首尾而合治之，而后可无弊也”。靳辅治水大致遵奉明代潘季驯“束水冲沙”之法。治河十几年所取得的成果，与陈潢的协助密不可分。靳辅出任河道总督时，正值黄河、淮河泛滥严重，使得江南的漕运米粮不能顺利运往北京，靳辅首先在淮阴以东当时黄河的两岸建造坚固的河堤，一直延伸到距离海岸线20里的地方，又在淮阴以西地区，加淝沿淮河与洪泽湖的一线的若干河堤、水坝，使之不再溃堤。如此一来，河水单纯向前流动，增大了冲刷河底淤的力量，使黄河自我除沙的功能大为加强。靳辅知道运河关系国家漕运，为了不让黄河改道夺路而使其中断，于是在黄河北岸开辟一条长达一百八十里，为了防止改道运河的黄河之水溃决大堤，靳辅在淮阴至江都治运河各县兴筑若干「减水坝」1=即在堤下开几个可以开闭的洞，洞外开凿引河，与附近的沟渠或小河相连，在必要时可开洞放水。」。不但防止运河泛滥或缺口，又可以使用在运河周围的农田灌溉。康熙二十四年（1685年），靳辅与安徽按察使小于成龙（小于成龙）备注|1=于成龙字振甲，直隶人，同慕天颜友好。与另一位“天下第一廉吏”于成龙不同人。不合，小于成龙主张开挖下-{游}-河道、疏通海口，康熙二十七年（1688年），两人多次激辩，小于成龙攻击靳辅党附纳兰明珠。康熙二十七年，御史郭琇告靳辅治河九年无功，靳辅被免职，陈潢病死狱中。康熙三十一年（1692年）复职，是年即病逝。卒谥文襄。著有《治河方略》八卷、《治河奏绩书》四卷、《靳文襄奏疏》八卷。== 评价与争议 ==靳辅为官清廉，除应得的俸禄外，治河公款分文不取，工程中的大小开支更是详细记载，也时常亲临施工现场监督指挥，得到治河民工和两岸百姓的一致称赞。靳辅治河虽然有所贡献，治河的方法确实踏实可靠，然而直隶巡抚小于成龙等人却认为靳辅治河的方法，斥资过大、旷日费时，靳辅以「治河无功」罪名，遭到免职。靳辅身故后，此后，王新命、小于成龙、张鹏翮相继任河道总督，基本上仍然沿用靳辅与陈潢的治河方法，受到康熙帝的嘲讽。== 参考文献 ==脚注引用书籍* 《图说清朝》，龚书铎，知书房出版社，ISBN 978-986-7151-63-6* 《康乾盛世》，童超，知书房出版社，ISBN 978-986-6344-28-2文献* (明–清)靳辅，生卒：崇祯6年(天聪7年)-康熙31年，中央研究院历史语言研究所== 参见 ==* 陈潢* 潘季驯*水利工程"
	keywords := []string{"水利工程", "李冰"}
	result := StringKeywordsOrderCheck(&s, keywords)
	println(result)
}

func Test_Pretext(t *testing.T) {
	// aba
	s := "`apple,iphone,apple"
	result := Pretext(s)
	fmt.Printf("%s\n", result)
}

func Test_JoinInt(t *testing.T) {
	// aba
	n := []int64{1, 2, 3}
	result := JoinInt(n, ",")
	fmt.Printf("%s\n", result)
}

func Benchmark_GetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(64)
	}
}
