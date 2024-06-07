package stream

import (
	"fmt"
	"testing"
)

type testItem struct {
	itemNum   int
	itemValue string
}

func TestReverse(t *testing.T) {
	resReverse := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 4, itemValue: "item4"},
	).Reverse().ToSlice()
	fmt.Println(resReverse)
}

func TestConcatenate(t *testing.T) {
	resConcat := Concat(Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
	), Of(
		testItem{itemNum: 5, itemValue: "item5"},
		testItem{itemNum: 6, itemValue: "item6"},
	), Of(
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 4, itemValue: "item4"},
	)).ToSlice()
	fmt.Println(resConcat)
}

func TestCount(t *testing.T) {
	res := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
	).Count()
	fmt.Println(res)
}

func TestFilter(t *testing.T) {
	Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 4, itemValue: "item4"},
		testItem{itemNum: 5, itemValue: "item5"},
		testItem{itemNum: 6, itemValue: "item6"},
		testItem{itemNum: 7, itemValue: "item7"},
		testItem{itemNum: 8, itemValue: "item8"},
		testItem{itemNum: 9, itemValue: "item9"},
	).Filter(func(item testItem) bool {
		if item.itemNum%2 != 0 {
			return true
		}
		return false
	}).ForEach(func(item testItem) {
		fmt.Print(item.itemNum)
		fmt.Println(item.itemValue)
	})
}

func TestPeek(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 5, itemValue: "item5"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	slice := FromSlice(items).Peek(func(item *testItem) {
		item.itemNum += 1
	}).ToSlice()
	fmt.Printf("%+v \n", slice)
}

func TestPeekP(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 5, itemValue: "item5"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	slice := FromSlice(items).PeekP(func(item testItem) {
		item.itemValue = item.itemValue + "peek"
	}).ToSlice()
	fmt.Printf("%+v \n", slice)
}

func TestLimit(t *testing.T) {
	resLimit := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 3, itemValue: "item4"},
		testItem{itemNum: 4, itemValue: "item4"},
	).Skip(1).Limit(7).ToSlice()
	fmt.Println(resLimit)
}

func TestSkip(t *testing.T) {
	resSkip := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 3, itemValue: "item4"},
		testItem{itemNum: 4, itemValue: "item4"},
	).Skip(1).ToSlice()
	fmt.Println(resSkip)
}

func TestDistinctBy(t *testing.T) {
	resReverse := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
		testItem{itemNum: 3, itemValue: "item4"},
		testItem{itemNum: 4, itemValue: "item4"},
	).DistinctBy(func(item testItem) any {
		return item.itemValue
	}).ToSlice()
	fmt.Println(resReverse)
}

func TestDistinct(t *testing.T) {
	resReverse := Of(
		1, 2, 3, 3, 5, 6, 5,
	).Distinct().ToSlice()
	fmt.Println(resReverse)
}

func TestSorted(t *testing.T) {
	resSorted := Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
	).Sorted(func(a, b testItem) bool {
		// 降序
		return a.itemNum > b.itemNum
	}).ToSlice()
	fmt.Println(resSorted)
}

func TestMax(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	max, _ := FromSlice(items).Max(func(newItem testItem, oldItem testItem) bool {
		return newItem.itemNum > oldItem.itemNum
	}).Get()
	fmt.Printf("%+v \n", max)
}

func TestMin(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	min, _ := FromSlice(items).Min(func(newItem testItem, oldItem testItem) bool {
		return newItem.itemNum < oldItem.itemNum
	}).Get()
	fmt.Printf("%+v \n", min)
}

func TestForEach(t *testing.T) {
	Of(
		testItem{itemNum: 1, itemValue: "item1"},
		testItem{itemNum: 2, itemValue: "item2"},
		testItem{itemNum: 3, itemValue: "item3"},
	).ForEach(func(item testItem) {
		fmt.Print(item.itemNum)
		fmt.Println(item.itemValue)
	})
}

func TestAllMatch(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	allMatch := FromSlice(items).AllMatch(func(item testItem) bool {
		// 返回此流中是否全都==1
		return item.itemNum == 1
	})
	fmt.Printf("%+v \n", allMatch)
}

func TestAnyMatch(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	anyMatch := FromSlice(items).AnyMatch(func(item testItem) bool {
		// 返回此流中是否存在 == 8的
		return item.itemNum == 8
	})
	fmt.Printf("%+v \n", anyMatch)
}

func TestNoneMatch(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"},
	}
	noneMatch := FromSlice(items).NoneMatch(func(item testItem) bool {
		// 返回此流中是否全部不等于8
		return item.itemNum == 8
	})
	fmt.Printf("%+v \n", noneMatch)
}

func TestFindFirst(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	findFirst := FromSlice(items).FindFirst()
	fmt.Println(findFirst.Get())
}

func TestFindLast(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	result := FromSlice(items).FindLast()
	fmt.Println(result.Get())
}

func TestToMapString(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	slice := FromSlice(items).ToMapString(func(item testItem) string {
		return item.itemValue
	}, func(item testItem) testItem {
		return item
	}, func(oldV, newV testItem) testItem {
		return newV
	})
	fmt.Printf("%+v \n", slice)
}

func TestToMapInt(t *testing.T) {
	items := []testItem{
		{itemNum: 7, itemValue: "item7"}, {itemNum: 6, itemValue: "item6"},
		{itemNum: 1, itemValue: "item1"}, {itemNum: 2, itemValue: "item2"},
		{itemNum: 3, itemValue: "item3"}, {itemNum: 4, itemValue: "item4"},
		{itemNum: 5, itemValue: "item5"}, {itemNum: 7, itemValue: "item5"},
		{itemNum: 6, itemValue: "item5"}, {itemNum: 8, itemValue: "item8"},
	}
	slice := FromSlice(items).ToMapInt(func(item testItem) int {
		return item.itemNum
	}, func(item testItem) testItem {
		return item
	}, func(oldV, newV testItem) testItem {
		return newV
	})
	fmt.Printf("%+v \n", slice)
}
