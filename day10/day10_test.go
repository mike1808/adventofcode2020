package day10

import (
	"strings"
	"testing"

	"github.com/mike1808/adventofcode2020/util"
)

func TestDay10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day10()
}

func TestPart1(t *testing.T) {
	input, _ := util.ReadIntSlice(strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`))
	actual := Part1(input)

	if actual != 7*5 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 7*5)
	}

	input, _ = util.ReadIntSlice(strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`))
	actual = Part1(input)

	if actual != 22*10 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 22*10)
	}
}

func TestPart2(t *testing.T) {
	input, _ := util.ReadIntSlice(strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`))
	actual := Part2(input)

	if actual != 8 {
		t.Errorf("Part2() returned: %d, expected: %d", actual, 8)
	}

	input, _ = util.ReadIntSlice(strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`))
	actual = Part2(input)

	if actual != 19208 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 19208)
	}
}
