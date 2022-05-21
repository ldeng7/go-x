package collectionx

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRbtree(t *testing.T) {
	Convey("test rbtree", t, func() {
		tree := (&RBTree[int, int]{}).Init(func(a, b int) bool { return a < b })

		{
			So(tree.Len(), ShouldEqual, 0)
			So(tree.Head(), ShouldBeNil)
			So(tree.Tail(), ShouldBeNil)
			So(tree.LowerBound(0), ShouldBeNil)
			So(tree.UpperBound(0), ShouldBeNil)
			l, r := tree.EqualRange(0)
			So(l, ShouldBeNil)
			So(r, ShouldBeNil)
			So(tree.Count(0), ShouldEqual, 0)
			So(tree.Remove(0), ShouldEqual, 0)
		}

		l := 17
		nums, m := make([]int, 1, l), map[int]int{}
		{
			k := rand.Intn(l)
			nums[0], m[k] = k, 1
			node := tree.Add(k, 0)
			So(node.Key(), ShouldEqual, k)

			So(tree.Len(), ShouldEqual, 1)
			head := tree.Head()
			So(head.Prev(), ShouldBeNil)
			So(head.Key(), ShouldEqual, k)
			So(head.Next(), ShouldBeNil)
			So(head, ShouldEqual, tree.Tail())

			So(tree.LowerBound(k-1), ShouldEqual, head)
			So(tree.UpperBound(k-1), ShouldEqual, head)
			l, r := tree.EqualRange(k - 1)
			So(l, ShouldBeNil)
			So(r, ShouldBeNil)
			So(tree.Count(k-1), ShouldEqual, 0)
			So(tree.LowerBound(k), ShouldEqual, head)
			So(tree.UpperBound(k), ShouldBeNil)
			l, r = tree.EqualRange(k)
			So(l, ShouldEqual, head)
			So(r, ShouldBeNil)
			So(tree.Count(k), ShouldEqual, 1)
			So(tree.LowerBound(k+1), ShouldBeNil)
			So(tree.UpperBound(k+1), ShouldBeNil)
		}

		for i := 2; i <= l; i++ {
			k := rand.Intn(l)
			nums, m[k] = append(nums, k), m[k]+1
			sort.Ints(nums)
			tree.Add(k, 0)

			So(tree.Len(), ShouldEqual, i)
			head, tail := tree.Head(), tree.Tail()
			So(head.Prev(), ShouldBeNil)
			So(head.Key(), ShouldEqual, nums[0])
			So(head.Next().Key(), ShouldEqual, nums[1])
			for j, node := 1, head.Next(); j != i-1; node, j = node.Next(), j+1 {
				So(node.Prev().Key(), ShouldEqual, nums[j-1])
				So(node.Key(), ShouldEqual, nums[j])
				So(node.Next().Key(), ShouldEqual, nums[j+1])
			}
			So(tail.Prev().Key(), ShouldEqual, nums[i-2])
			So(tail.Key(), ShouldEqual, nums[i-1])
			So(tail.Next(), ShouldBeNil)

			So(tree.LowerBound(k).Key(), ShouldEqual, k)
			So(tree.UpperBound(tail.Key()), ShouldBeNil)
			So(tree.Count(k), ShouldEqual, m[k])
		}

		for len(nums) > 0 {
			l = len(nums)
			k := nums[rand.Intn(l)]
			nums1 := make([]int, 0, l-1)
			for _, num := range nums {
				if num != k {
					nums1 = append(nums1, num)
				}
			}
			nums, l = nums1, len(nums1)
			So(tree.Remove(k), ShouldEqual, m[k])
			delete(m, k)

			So(tree.Len(), ShouldEqual, l)
			head, tail := tree.Head(), tree.Tail()
			if l >= 1 {
				So(head.Prev(), ShouldBeNil)
				So(head.Key(), ShouldEqual, nums[0])
				if l >= 2 {
					So(head.Next().Key(), ShouldEqual, nums[1])
				} else {
					So(head.Next(), ShouldBeNil)
				}
				for j, node := 1, head.Next(); j < l-1; node, j = node.Next(), j+1 {
					So(node.Prev().Key(), ShouldEqual, nums[j-1])
					So(node.Key(), ShouldEqual, nums[j])
					So(node.Next().Key(), ShouldEqual, nums[j+1])
				}
				if l >= 2 {
					So(tail.Prev().Key(), ShouldEqual, nums[l-2])
				} else {
					So(head.Prev(), ShouldBeNil)
				}
				So(tail.Key(), ShouldEqual, nums[l-1])
				So(tail.Next(), ShouldBeNil)
			} else {
				So(head, ShouldBeNil)
				So(tail, ShouldBeNil)
			}
		}

		{
			node, ok := tree.AddUnique(1, 0)
			So(node.Key(), ShouldEqual, 1)
			So(ok, ShouldBeTrue)
			node, ok = tree.AddUnique(1, 0)
			So(tree.Len(), ShouldEqual, 1)
			So(node.Key(), ShouldEqual, 1)
			So(ok, ShouldBeFalse)
			tree.AddUnique(0, 0)
			node, _ = tree.AddUnique(3, 0)
			tree.AddUnique(2, 0)
			tree.RemoveAt(node)
			So(tree.Len(), ShouldEqual, 3)
			So(tree.Head().Key(), ShouldEqual, 0)
			So(tree.Tail().Key(), ShouldEqual, 2)
		}
	})
}
