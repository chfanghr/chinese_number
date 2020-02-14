package chinese_number

import "errors"

type Node struct {
	value       int
	factor      int
	text        rune
	left, right *Node
}

var (
	zero = Node{
		value:  0,
		factor: 0,
		text:   '零',
	}
	one = Node{
		value:  1,
		factor: 1,
		text:   '一',
	}
	two = Node{
		value:  2,
		factor: 1,
		text:   '二',
	}
	three = Node{
		value:  3,
		factor: 1,
		text:   '三',
	}
	four = Node{
		value:  4,
		factor: 1,
		text:   '四',
	}
	five = Node{
		value:  5,
		factor: 1,
		text:   '五',
	}
	six = Node{
		value:  6,
		factor: 1,
		text:   '六',
	}
	seven = Node{
		value:  7,
		factor: 1,
		text:   '七',
	}
	eight = Node{
		value:  8,
		factor: 1,
		text:   '八',
	}
	nine = Node{
		value:  9,
		factor: 1,
		text:   '九',
	}
	ten = Node{
		factor: 10,
		text:   '十',
	}
	hundred = Node{
		factor: 100,
		text:   '百',
	}
	thousand = Node{
		factor: 1000,
		text:   '千',
	}
	tenThousand = Node{
		factor: 10000,
		text:   '万',
	}
	hundredMillion = Node{
		factor: 100000000,
		text:   '亿',
	}
)

func (c *Node) GetValue() int {
	if c.factor <= 1 {
		return c.value
	}

	var left, right int

	if c.left != nil {
		left = c.left.GetValue()
	} else {
		left = 1
	}

	left *= c.factor

	if c.right != nil {
		right = c.right.GetValue()
	} else {
		right = 0
	}

	if c.right != nil && c.right.factor == 1 && c.right.left == nil && c.factor > 10 {
		right = right * (c.factor / 10)
	}

	return left + right
}

func ParseChineseNumberCharacter(num rune) (Node, error) {
	switch num {
	case '零':
		return zero, nil
	case '一':
		return one, nil
	case '二':
		return two, nil
	case '三':
		return three, nil
	case '四':
		return four, nil
	case '五':
		return five, nil
	case '六':
		return six, nil
	case '七':
		return seven, nil
	case '八':
		return eight, nil
	case '九':
		return nine, nil
	case '十':
		return ten, nil
	case '百':
		return hundred, nil
	case '千':
		return thousand, nil
	case '万':
		return tenThousand, nil
	case '亿':
		return hundredMillion, nil
	default:
		return Node{}, errors.New("not a chinese number")
	}
}

func ToArabicNumber(chineseNumber string) (int, error) {
	if chineseNumber == "" {
		return 0, errors.New("empty string")
	}

	var root *Node
	runeChineseNumber := []rune(chineseNumber)

	for _, r := range runeChineseNumber {
		current, err := ParseChineseNumberCharacter(r)
		if err != nil {
			return 0, err
		}
		if root == nil {
			root = &current
		} else {
			if current.factor == root.factor {
				return 0, errors.New("not a chinese number")
			}
			root = buildChineseNumberTree(root, &current)
		}
	}

	//noinspection GoNilness
	return root.GetValue(), nil
}

func buildChineseNumberTree(root, current *Node) *Node {
	if current.factor > root.factor {
		current.left = root
		return current
	} else {
		if root.right == nil {
			root.right = current
			return root
		} else {
			root.right = buildChineseNumberTree(root.right, current)
			return root
		}
	}
}
