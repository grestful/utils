package utils

const StringItemTyp = 1
const ListItemType = 2
const IntItemTyp = 3

type Item interface {
	CompareTo(item Item) int
	GetType() int
	IsNull() bool
}

type IntItem struct {
	val int
}

func (i *IntItem) GetType() int {
	return IntItemTyp
}

func (i *IntItem) IsNull() bool {
	return 0 == i.val
}

//switch ( item.getType() )
//{
//case INT_ITEM:
//int itemValue = ( (IntItem) item ).value;
//return Integer.compare( value, itemValue );
//case LONG_ITEM:
//case BIGINTEGER_ITEM:
//return -1;
//
//case STRING_ITEM:
//return 1; // 1.1 > 1-sp
//
//case LIST_ITEM:
//return 1; // 1.1 > 1-1
//
//default:
//throw new IllegalStateException( "invalid item: " + item.getClass() );
//}
func (i *IntItem) CompareTo(item Item) int {
	if item == nil {
		if i.val == 0 {
			return 0
		} else {
			return 1
		}
	}
	switch item.GetType() {
	case IntItemTyp:
		v, ok := item.(*IntItem)
		if ok {
			if i.val > v.val {
				return 1
			} else {
				return 0
			}
		}
	case ListItemType:
		return 1
	case StringItemTyp:
		return 1
	}

	return 0
}

// 解析 pom.xml 中的版本
func VersionParser(version string) ([]string, error) {
	return nil, nil
}

// 比较两个依赖的版本大小
// 如果成功，返回比较结果，如果失败，返回错误原因
// 如果两个版本相同，返回 0
// 如果第一个版本大于第二个版本，返回 1
// 如果第一个版本小于第二个版本，返回 -1
func VersionCompare(ver1, ver2 string) (int, error) {
	ver1Part, err := VersionParser(ver1)
	if err != nil {
		return 2, err
	}

	ver2Part, err := VersionParser(ver2)
	if err != nil {
		return 2, err
	}

	return versionCompare(ver1Part, ver2Part)
}

func versionCompare(ver1, ver2 []string) (int, error) {

	return 0, nil
}
